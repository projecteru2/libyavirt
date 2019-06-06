package client

import (
	"bytes"
	"encoding/json"
	"io"
)

func decode(resp *Resp, obj interface{}) error {
	return json.NewDecoder(resp.body).Decode(obj)
}

func encode(obj interface{}) (io.Reader, error) {
	if obj == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	var err = json.NewEncoder(&buf).Encode(obj)

	return &buf, err
}
