// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/mesh.proto

package types

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The Mesh types known to Service Mesh Hub. Note that there is some duplication between this enum and the discovery.MeshSpec message
type MeshType int32

const (
	MeshType_ISTIO   MeshType = 0
	MeshType_APPMESH MeshType = 1
	MeshType_CONSUL  MeshType = 2
	MeshType_LINKERD MeshType = 3
)

var MeshType_name = map[int32]string{
	0: "ISTIO",
	1: "APPMESH",
	2: "CONSUL",
	3: "LINKERD",
}

var MeshType_value = map[string]int32{
	"ISTIO":   0,
	"APPMESH": 1,
	"CONSUL":  2,
	"LINKERD": 3,
}

func (x MeshType) String() string {
	return proto.EnumName(MeshType_name, int32(x))
}

func (MeshType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{0}
}

func init() {
	proto.RegisterEnum("core.zephyr.solo.io.MeshType", MeshType_name, MeshType_value)
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/core/v1alpha1/mesh.proto", fileDescriptor_4d4089e58fa1f863)
}

var fileDescriptor_4d4089e58fa1f863 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x29, 0x88, 0x02, 0x66, 0x89, 0x02, 0x53, 0x87, 0x32, 0x83, 0xa8, 0xad, 0x8a, 0x11,
	0x31, 0xf0, 0x51, 0x89, 0xa8, 0x9f, 0x90, 0xb2, 0xb0, 0x39, 0xe1, 0x6a, 0x1b, 0x12, 0xce, 0xb2,
	0x2f, 0x48, 0xe1, 0x17, 0xf1, 0xbb, 0xf8, 0x25, 0x28, 0x69, 0x44, 0xa5, 0x2e, 0xd9, 0xce, 0x7e,
	0x9e, 0xb3, 0xfc, 0xbe, 0xec, 0x46, 0x19, 0xd2, 0x45, 0xc2, 0x53, 0xcc, 0x85, 0xc7, 0x0c, 0x07,
	0x06, 0x45, 0x0e, 0x5e, 0x0f, 0xac, 0xc3, 0x77, 0x48, 0xc9, 0x0b, 0x69, 0x8d, 0x48, 0xd1, 0x81,
	0xf8, 0x1a, 0xca, 0xcc, 0x6a, 0x39, 0xac, 0x39, 0xb7, 0x0e, 0x09, 0xc3, 0x93, 0x8a, 0xf0, 0x6f,
	0xb0, 0xba, 0x74, 0xbc, 0xda, 0xe7, 0x06, 0x7b, 0xe7, 0x6d, 0x0f, 0x38, 0x58, 0xad, 0xf7, 0x7b,
	0x97, 0x6d, 0xaa, 0x27, 0x49, 0x85, 0x6f, 0xec, 0x33, 0x85, 0xa8, 0x32, 0x10, 0xf5, 0x29, 0x29,
	0x56, 0x82, 0x4c, 0x0e, 0x9e, 0x64, 0x6e, 0x1b, 0xa1, 0xbf, 0x2d, 0xbc, 0x15, 0x4e, 0x92, 0xc1,
	0xcf, 0x86, 0x9f, 0x2a, 0x54, 0x58, 0x8f, 0xa2, 0x9a, 0xd6, 0xb7, 0x17, 0xd7, 0xec, 0x70, 0x0a,
	0x5e, 0x2f, 0x4b, 0x0b, 0xe1, 0x11, 0xdb, 0x8f, 0xe2, 0x65, 0x34, 0x0f, 0x76, 0xc2, 0x63, 0x76,
	0x70, 0xbb, 0x58, 0x4c, 0x47, 0xf1, 0x63, 0xd0, 0x09, 0x19, 0xeb, 0xde, 0xcf, 0x67, 0xf1, 0xcb,
	0x24, 0xd8, 0xad, 0xc0, 0x24, 0x9a, 0x8d, 0x47, 0xcf, 0x0f, 0xc1, 0xde, 0xdd, 0xd3, 0xcf, 0x6f,
	0xbf, 0xf3, 0x3a, 0x6e, 0xad, 0xd1, 0x7e, 0xa8, 0xff, 0x78, 0x5b, 0x85, 0x6d, 0xd2, 0x52, 0x69,
	0xc1, 0x27, 0xdd, 0xfa, 0x5b, 0x57, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x38, 0xb3, 0x25,
	0x9c, 0x01, 0x00, 0x00,
}
