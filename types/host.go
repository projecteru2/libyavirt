package types

import "fmt"

type HostInfo struct {
	ID      string
	Cpu     int
	Mem     int64
	Storage int64
}

func (h HostInfo) String() string {
	return fmt.Sprintf("%s, cpu: %d, memory: %d, storage: %d",
		h.ID, h.Cpu, h.Mem, h.Storage)
}
