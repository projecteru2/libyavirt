package types

import yavpb "github.com/projecteru2/libyavirt/grpc/gen"

// SysImage .
type SysImage struct {
	Name   string
	User   string
	Distro string
	ID     string
	Type   string
}

func ToGRPCImageItem(img SysImage) *yavpb.ImageItem {
	return &yavpb.ImageItem{
		Name:   img.Name,
		User:   img.User,
		Distro: img.Distro,
		Id:     img.ID,
		Type:   img.Type,
	}
}

// UserImage .
type UserImage struct {
	ID            string
	Name          string
	Distro        string
	LatestVersion int64
	Size          int64
}
