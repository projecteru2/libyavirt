package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
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

	if _, err := io.Copy(ioutil.Discard, r.body); err != nil {
		logrus.Errorf("[libyavirt] Copy resp body error: %v", err)
	}
}
