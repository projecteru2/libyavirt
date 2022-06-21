package httpclient

import (
	"context"
	"fmt"
	"net"
	"strings"
)

func (c *HTTPClient) procReqErr(err error) error {
	switch {
	case err == context.Canceled:
		fallthrough
	case c.isDeadlineExceededErr(err):
		return err
	}

	if err, ok := err.(net.Error); ok {
		switch {
		case err.Timeout():
			fallthrough
		case c.isConnRefusedErr(err):
			return fmt.Errorf("cannot connect to yavirtd at %s", c.addr)
		}
	}

	return fmt.Errorf("error during request: %v", err)
}

func (c *HTTPClient) isConnRefusedErr(err error) bool {
	return strings.Contains(err.Error(), "connection refused")
}

func (c *HTTPClient) isDeadlineExceededErr(err error) bool {
	return strings.HasSuffix(err.Error(), "context deadline exceeded")
}
