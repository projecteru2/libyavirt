package httpclient

import (
	"context"

	"github.com/projecteru2/libyavirt/types"
)

// Info .
func (c *HTTPClient) Info(ctx context.Context) (reply types.HostInfo, err error) {
	_, err = c.Get(ctx, "/info", &reply)
	return
}
