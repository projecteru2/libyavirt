package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Resp .
type Resp struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
}

func (r *Resp) close() {
	if r.body == nil {
		return
	}

	defer func() {
		r.body.Close()
		r.body = nil
	}()

	io.Copy(ioutil.Discard, r.body)
}
