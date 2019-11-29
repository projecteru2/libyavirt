package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const MagicPrefix = "SHOPEE-YET-ANOTHER-VIRT-20190429"

func EruID(id string) string {
	return fmt.Sprintf("%s%032s", MagicPrefix, id)
}

type CreateGuestReq struct {
	Cpu       int
	Mem       int64
	ImageName string
	Volumes   map[string]int64
	DmiUuid   string
}

type GuestReq struct {
	ID string `uri:"id" binding:"required"`
}

func (r GuestReq) VirtID() string {
	var id = r.ID
	if strings.HasPrefix(id, MagicPrefix) {
		id = id[len(MagicPrefix):]
	}

	id = r.checkOldVersionID(id)

	return id
}

func (r GuestReq) checkOldVersionID(id string) string {
	var i64, err = strconv.ParseInt(id, 10, 64)
	switch {
	case err != nil:
		fallthrough
	case i64 > 999999:
		return id
	default:
		return fmt.Sprintf("guest-%06d", i64)
	}
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

type AttachGuestFlags struct {
	Safe  bool
	Force bool
}

type ExecuteGuestMessage struct {
	ID       string
	Data     []byte
	ExitCode int
}

type EruGuestStatus struct {
	ID         string
	EruGuestID string
	Running    bool
	TTL        time.Duration
	CIDRs      []string
}

func NewEruGuestStatus(id string) (st EruGuestStatus) {
	st.ID = id
	st.EruGuestID = EruID(id)
	return
}

func (s EruGuestStatus) GetIPAddrs() string {
	return strings.Join(s.CIDRs, ", ")
}
