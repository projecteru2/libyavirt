package httpclient

import (
	"context"

	"github.com/projecteru2/libyavirt/types"
)

func (c *httpClient) Info(ctx context.Context) (reply types.HostInfo, err error) {
	_, err = c.Get(ctx, "/info", &reply)
	return
}
