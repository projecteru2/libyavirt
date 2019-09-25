package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type headers = map[string][]string

func (c *httpClient) Post(ctx context.Context, path string, obj, reply interface{}) (*Resp, error) {
	return c.send(&queryOption{
		ctx:    ctx,
		path:   path,
		obj:    obj,
		method: http.MethodPost,
	}, reply)
}

func (c *httpClient) Get(ctx context.Context, path string, reply interface{}) (*Resp, error) {
	return c.send(&queryOption{
		ctx:    ctx,
		path:   path,
		method: http.MethodGet,
	}, reply)
}

func (c *httpClient) send(q *queryOption, reply interface{}) (*Resp, error) {
	req, err := c.buildReq(q)
	if err != nil {
		return nil, err
	}

	resp, err := c.req(q.context(), req)
	if err != nil {
		return nil, err
	}

	defer resp.close()

	if err := c.requireOK(resp, err); err != nil {
		return nil, err
	}

	return resp, decode(resp, reply)
}

func (c *httpClient) req(ctx context.Context, req *http.Request) (*Resp, error) {
	req = req.WithContext(ctx)

	var rawResp, err = c.http.Do(req)
	if err != nil {
		return nil, c.procReqErr(err)
	}

	var resp = &Resp{
		body:       rawResp.Body,
		header:     rawResp.Header,
		statusCode: rawResp.StatusCode,
	}

	return resp, nil
}

func (c *httpClient) buildReq(q *queryOption) (*http.Request, error) {
	body, err := q.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(q.method, c.getPath(q.path), body)
	if err != nil {
		return nil, err
	}

	c.withDefaultHeaders(req, q.headers)

	req.URL.Host = c.addr
	req.URL.Scheme = c.scheme

	return req, nil
}

func (c *httpClient) withDefaultHeaders(req *http.Request, hdrs headers) {
	for k, v := range c.defaultHeaders {
		req.Header.Set(k, v)
	}

	if hdrs != nil {
		for k, v := range hdrs {
			req.Header[k] = v
		}
	}
}

func (c *httpClient) requireOK(resp *Resp, err error) error {
	if err != nil {
		return err
	}

	if resp.statusCode == http.StatusOK {
		return nil
	}

	var buf bytes.Buffer
	io.Copy(&buf, resp.body)

	return fmt.Errorf("unexpected status code: %d (%s)", resp.statusCode, buf.Bytes())
}

func (c *httpClient) getPath(path string) string {
	return fmt.Sprintf("/%s/%s", c.ver, strings.Trim(path, "/"))
}
