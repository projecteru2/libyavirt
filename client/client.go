package client

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

type Client struct {
	addr   string
	scheme string
	ver    string

	http *http.Client

	defaultHeaders map[string]string
}

func New(addr, ver string) (*Client, error) {
	if len(addr) < 1 {
		return nil, fmt.Errorf("invalid addr")
	}
	if len(ver) < 1 {
		return nil, fmt.Errorf("invalid ver")
	}

	return &Client{
		addr:   addr,
		scheme: "http",
		ver:    ver,

		http: defaultHttpClient(),

		defaultHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func (c *Client) Close() error {
	if tr, ok := c.http.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}

	return fmt.Errorf("invalid *http.Transport")
}

func defaultHttpClient() *http.Client {
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
