package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	yavpb "github.com/projecteru2/libyavirt/grpc/gen"
)

const (
	// MagicPrefix .
	MagicPrefix = "00-ERU-YET-ANOTHER-VIRT-20230421"

	// BufferSize for grpc send file 1 MB
	BufferSize = 1024 * 1024
)

// Guest Operations
const (
	// OpStart .
	OpStart = "start"
	// OpStop .
	OpStop = "stop"
	// OpDestroy .
	OpDestroy = "destroy"
)

// EruID .
func EruID(id string) string {
	id = strings.TrimPrefix(id, "guest-")
	return fmt.Sprintf("%s%032s", MagicPrefix, id)
}

// Volume .
type Volume struct {
	Mount    string
	Capacity int64
	IO       string
}

// CreateGuestReq .
type CreateGuestReq struct {
	CPU        int
	Mem        int64
	ImageName  string
	ImageUser  string
	Volumes    []Volume
	DmiUUID    string
	Labels     map[string]string
	AncestorID string
	Cmd        []string
	Lambda     bool
	Stdin      bool
	Resources  map[string][]byte
}

// AncestorVirtID .
func (r CreateGuestReq) AncestorVirtID() string {
	return (GuestReq{ID: r.AncestorID}).VirtID()
}

// GetGrpcOpts .
func (r CreateGuestReq) GetGrpcOpts() *yavpb.CreateGuestOptions {
	ret := &yavpb.CreateGuestOptions{
		Cpu:        int64(r.CPU),
		Memory:     r.Mem,
		ImageName:  r.ImageName,
		ImageUser:  r.ImageUser,
		DmiUuid:    r.DmiUUID,
		Labels:     r.Labels,
		AncestorId: r.AncestorID,
		Cmd:        r.Cmd,
		Lambda:     r.Lambda,
		Stdin:      r.Stdin,
		Resources:  r.Resources,
	}
	ret.Volumes = make([]*yavpb.Volume, len(r.Volumes))
	for i, vol := range r.Volumes {
		ret.Volumes[i].Mount = vol.Mount
		ret.Volumes[i].Capacity = vol.Capacity
		ret.Volumes[i].Io = vol.IO
	}
	return ret
}

// GuestReq .
type GuestReq struct {
	ID    string `uri:"id" binding:"required"`
	Force bool   `uri:"force"`
}

// VirtID .
func (r GuestReq) VirtID() string {
	var id = r.ID
	id = strings.TrimPrefix(id, MagicPrefix)

	id = r.checkOldVersionID(id)

	return id
}

func (r GuestReq) checkOldVersionID(id string) string {
	var i64, err = strconv.ParseInt(id, 10, 64) //nolint
	switch {
	case err != nil:
		fallthrough
	case i64 > 999999: //nolint
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
	ImageUser string
	Networks  map[string]string
	Labels    map[string]string
	IPs       []string
	Hostname  string
	Running   bool
}

// AttachGuestFlags .
type AttachGuestFlags struct {
	Safe  bool
	Force bool
}

// ExecuteGuestMessage .
type ExecuteGuestMessage struct {
	Pid      int
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

// ResizeConsoleWindowReq .
type ResizeConsoleWindowReq struct {
	GuestReq
	Height uint
	Width  uint
}

// ExecuteGuestReq .
type ExecuteGuestReq struct {
	GuestReq
	Commands []string
}

// CaptureGuestReq .
type CaptureGuestReq struct {
	GuestReq
	User       string
	Name       string
	Overridden bool
}

// GetGrpcOpts .
func (r CaptureGuestReq) GetGrpcOpts() *yavpb.CaptureGuestOptions {
	return &yavpb.CaptureGuestOptions{
		Id:         r.ID,
		Name:       r.Name,
		User:       r.User,
		Overridden: r.Overridden,
	}
}

// ResizeGuestReq .
type ResizeGuestReq struct {
	GuestReq
	CPU       int
	Mem       int64
	Volumes   []Volume
	Resources map[string][]byte
}

// GetGrpcOpts .
func (r ResizeGuestReq) GetGrpcOpts() *yavpb.ResizeGuestOptions {
	ret := &yavpb.ResizeGuestOptions{
		Id:        r.ID,
		Cpu:       int64(r.CPU),
		Memory:    r.Mem,
		Resources: r.Resources,
	}
	ret.Volumes = make([]*yavpb.Volume, len(r.Volumes))
	for i, vol := range r.Volumes {
		ret.Volumes[i].Mount = vol.Mount
		ret.Volumes[i].Capacity = vol.Capacity
		ret.Volumes[i].Io = vol.IO
	}
	return ret
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

// GetGuestIDListReq .
type GetGuestIDListReq struct {
	Filters map[string]string
}

type Snapshot struct {
	VolID       string
	VolMountDir string
	SnapID      string
	CreatedTime int64
}

type Snapshots []*Snapshot

// ListSnapshotReq .
type ListSnapshotReq struct {
	ID    string `uri:"id" binding:"required"`
	VolID string
}

// CreateSnapshotReq .
type CreateSnapshotReq struct {
	ID    string `uri:"id" binding:"required"`
	VolID string
}

// CommitSnapshotReq .
type CommitSnapshotReq struct {
	ID     string `uri:"id" binding:"required"`
	VolID  string
	SnapID string
}

// RestoreSnapshotReq .
type RestoreSnapshotReq struct {
	ID     string `uri:"id" binding:"required"`
	VolID  string
	SnapID string
}
