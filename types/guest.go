package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
)

// MagicPrefix .
const MagicPrefix = "SHOPEE-YET-ANOTHER-VIRT-20190429"

// EruID .
func EruID(id string) string {
	if strings.HasPrefix(id, "guest-") {
		id = id[6:]
	}
	return fmt.Sprintf("%s%032s", MagicPrefix, id)
}

// CreateGuestReq .
type CreateGuestReq struct {
	CPU       int
	Mem       int64
	ImageName string
	Volumes   map[string]int64
	DmiUUID   string
	Labels    map[string]string
}

// GetGrpcOpts .
func (r CreateGuestReq) GetGrpcOpts() *yavpb.CreateGuestOptions {
	return &yavpb.CreateGuestOptions{
		Cpu:       int64(r.CPU),
		Memory:    r.Mem,
		ImageName: r.ImageName,
		Volumes:   r.Volumes,
		DmiUuid:   r.DmiUUID,
		Labels:    r.Labels,
	}
}

// GuestReq .
type GuestReq struct {
	ID    string `uri:"id" binding:"required"`
	Force bool   `uri:"force"`
}

// VirtID .
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

// Guest .
type Guest struct {
	Resource
	CPU       int
	Mem       int64
	Storage   int64
	ImageID   int64
	ImageName string
	Networks  map[string]string
}

// AttachGuestFlags .
type AttachGuestFlags struct {
	Safe  bool
	Force bool
}

// ExecuteGuestMessage .
type ExecuteGuestMessage struct {
	ID       string
	Data     []byte
	ExitCode int
}

// EruGuestStatus .
type EruGuestStatus struct {
	ID         string
	EruGuestID string
	Running    bool
	Healthy    bool
	TTL        time.Duration
	CIDRs      []string
}

// NewEruGuestStatus .
func NewEruGuestStatus(id string) (st EruGuestStatus) {
	st.ID = id
	st.EruGuestID = EruID(id)
	return
}

// GetIPAddrs .
func (s EruGuestStatus) GetIPAddrs() string {
	return strings.Join(s.CIDRs, ", ")
}

// ExecuteGuestReq .
type ExecuteGuestReq struct {
	GuestReq
	Commands []string
}

// CaptureGuestReq .
type CaptureGuestReq struct {
	GuestReq
	Name string
}

// GetGrpcOpts .
func (r CaptureGuestReq) GetGrpcOpts() *yavpb.CaptureGuestOptions {
	return &yavpb.CaptureGuestOptions{
		Id:   r.ID,
		Name: r.Name,
	}
}

// ResizeGuestReq .
type ResizeGuestReq struct {
	GuestReq
	CPU     int
	Mem     int64
	Volumes map[string]int64
}

// GetGrpcOpts .
func (r ResizeGuestReq) GetGrpcOpts() *yavpb.ResizeGuestOptions {
	return &yavpb.ResizeGuestOptions{
		Id:      r.ID,
		Cpu:     int64(r.CPU),
		Memory:  r.Mem,
		Volumes: r.Volumes,
	}
}

// ConnectNetworkReq .
type ConnectNetworkReq struct {
	GuestReq
	Network string
	IPv4    string
}

// GetGrpcOpts .
func (r ConnectNetworkReq) GetGrpcOpts() *yavpb.ConnectNetworkOptions {
	return &yavpb.ConnectNetworkOptions{
		Id:      r.ID,
		Network: r.Network,
		Ipv4:    r.IPv4,
	}
}

// DisconnectNetworkReq .
type DisconnectNetworkReq ConnectNetworkReq

// GetGrpcOpts .
func (r DisconnectNetworkReq) GetGrpcOpts() *yavpb.DisconnectNetworkOptions {
	return &yavpb.DisconnectNetworkOptions{
		Id:      r.ID,
		Network: r.Network,
	}
}
