package types

import "fmt"

// HostInfo .
type HostInfo struct {
	ID      string
	CPU     int
	Mem     int64
	Storage int64
}

// String .
func (h HostInfo) String() string {
	return fmt.Sprintf("%s, cpu: %d, memory: %d, storage: %d",
		h.ID, h.CPU, h.Mem, h.Storage)
}
