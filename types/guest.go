package types

import (
	"fmt"
	"strconv"
	"strings"
)

const MagicPrefix = "SHOPEE-YET-ANOTHER-VIRT-20190429"

func EruID(id int64) string {
	return fmt.Sprintf("%s%032d", MagicPrefix, id)
}

type CreateGuestReq struct {
	Cpu       int
	Mem       int64
	ImageName string
	Volumes   map[string]int64
}

type GuestReq struct {
	ID string `uri:"id" binding:"required"`
}

func (r GuestReq) VirtID() (int64, error) {
	var id = r.ID
	if strings.HasPrefix(id, MagicPrefix) {
		id = id[len(MagicPrefix):]
	}
	return strconv.ParseInt(id, 10, 0)
}

type Guest struct {
	Resource
	Cpu       int
	Mem       int64
	Storage   int64
	ImageID   int64
	ImageName string
	Networks  map[string]string
}
