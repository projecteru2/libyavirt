package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"time"
)

// HTTPClient .
type HTTPClient struct {
	addr   string
	scheme string
	ver    string

	http *http.Client

	defaultHeaders map[string]string
}

// CopyToGuest .
func (c *HTTPClient) CopyToGuest(ctx context.Context, ID, dest string, content io.Reader, AllowOverwriteDirWithFile, CopyUIDGID bool) error {
	return errors.New("not and will not implemented")
}

// New .
func New(addr, ver string) (*HTTPClient, error) {
	if len(addr) < 1 {
		return nil, fmt.Errorf("invalid addr")
	}
	if len(ver) < 1 {
		return nil, fmt.Errorf("invalid ver")
	}

	return &HTTPClient{
		addr:   addr,
		scheme: "http",
		ver:    ver,

		http: defaultHttphttpClient(),

		defaultHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

// Close .
func (c *HTTPClient) Close() error {
	if tr, ok := c.http.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}

	return fmt.Errorf("invalid *http.Transport")
}

func defaultHttphttpClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if via[0].Method == http.MethodGet {
				return http.ErrUseLastResponse
			}
			return fmt.Errorf("unexpected redirect")
		},

		Transport: defaultPoolTransport(),
	}
}

func defaultPoolTransport() *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			KeepAlive: time.Second * 30,
			Timeout:   time.Second * 30,
		}).DialContext,

		ExpectContinueTimeout: time.Second,
		IdleConnTimeout:       time.Second * 90,
		TLSHandshakeTimeout:   time.Second * 10,

		MaxIdleConnsPerHost: runtime.GOMAXPROCS(0) + 1,
		MaxIdleConns:        16,

		Proxy: http.ProxyFromEnvironment,
	}
}
