package client

import (
	"context"
	"io"
	"net/http"
)

type queryOption struct {
	ctx     context.Context
	method  string
	path    string
	headers headers
	obj     interface{}
}

func (q *queryOption) body() (body io.Reader, err error) {
	if q.method != http.MethodPost && q.method != http.MethodPut {
		return
	}

	body, err = encode(q.obj)

	return
}

func (q *queryOption) context() context.Context {
	if q.ctx != nil {
		return q.ctx
	}
	return context.Background()
}
