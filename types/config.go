package types

import "context"

type Config struct {
	URI string
	CA  string
}

// BasicCredential for basic credential
type BasicCredential struct {
	username string
	password string
}

// NewBasicCredential new a basic credential
func NewBasicCredential(username, password string) *BasicCredential {
	return &BasicCredential{username, password}
}

// GetRequestMetadata for basic auth
func (c BasicCredential) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		c.username: c.password,
	}, nil
}

// RequireTransportSecurity for ssl require
func (c BasicCredential) RequireTransportSecurity() bool {
	return false
}
