package httpclient

import (
	"io"
	"io/ioutil" //nolint
	"net/http"
)

// Resp .
type Resp struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
}

func (r *Resp) close() error {
	if r.body == nil {
		return nil
	}

	defer func() {
		r.body.Close()
		r.body = nil
	}()

	_, err := io.Copy(ioutil.Discard, r.body)
	return err
}
