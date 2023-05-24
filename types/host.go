package types

import "fmt"

// HostInfo .
type HostInfo struct {
	ID        string
	CPU       int
	Mem       int64
	Storage   int64
	Resources map[string][]byte
}

// String .
func (h HostInfo) String() string {
	return fmt.Sprintf("%s, cpu: %d, memory: %d, storage: %d, resources: %v",
		h.ID, h.CPU, h.Mem, h.Storage, h.Resources)
}
