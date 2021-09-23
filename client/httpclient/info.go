package httpclient

import (
	"context"
	"errors"

	"github.com/projecteru2/libyavirt/types"
)

// Info .
func (c *HTTPClient) Info(ctx context.Context) (reply types.HostInfo, err error) {
	_, err = c.Get(ctx, "/info", &reply)
	return
}

// NetworkList .
func (c *HTTPClient) NetworkList(ctx context.Context, drivers []string) ([]*types.Network, error) {
	return nil, errors.New("does not and will not implemented")
}
