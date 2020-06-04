// Code generated by protoc-gen-go.
// source: MSG_CenterGate.proto
// DO NOT EDIT!

/*
Package MSG_CenterGate is a generated protocol buffer package.

It is generated from these files:
	MSG_CenterGate.proto

It has these top-level messages:
	CS_PlayerOnline_Req
	SC_PlayerOnline_Rsp
	CS_PlayerOffline_Req
	SC_PlayerOffline_Rsp
	CS_GetBroadCastSessions_Req
	SC_GetBroadCastSessions_Rsp
*/
package MSG_CenterGate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// add by stefan
// server
type SUBMSG int32

const (
	SUBMSG_Begin                   SUBMSG = 0
	SUBMSG_CS_PlayerOnline         SUBMSG = 1
	SUBMSG_SC_PlayerOnline         SUBMSG = 2
	SUBMSG_CS_PlayerOffline        SUBMSG = 3
	SUBMSG_SC_PlayerOffline        SUBMSG = 4
	SUBMSG_CS_GetBroadCastSessions SUBMSG = 5
	SUBMSG_SC_GetBroadCastSessions SUBMSG = 6
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_PlayerOnline",
	2: "SC_PlayerOnline",
	3: "CS_PlayerOffline",
	4: "SC_PlayerOffline",
	5: "CS_GetBroadCastSessions",
	6: "SC_GetBroadCastSessions",
}
var SUBMSG_value = map[string]int32{
	"Begin":                   0,
	"CS_PlayerOnline":         1,
	"SC_PlayerOnline":         2,
	"CS_PlayerOffline":        3,
	"SC_PlayerOffline":        4,
	"CS_GetBroadCastSessions": 5,
	"SC_GetBroadCastSessions": 6,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorCode int32

const (
	ErrorCode_Invalid ErrorCode = 0
	ErrorCode_Success ErrorCode = 1
	ErrorCode_Fail    ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var ErrorCode_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CS_PlayerOnline
type CS_PlayerOnline_Req struct {
	PlayerIdentify string `protobuf:"bytes,1,opt,name=PlayerIdentify" json:"PlayerIdentify,omitempty"`
}

func (m *CS_PlayerOnline_Req) Reset()                    { *m = CS_PlayerOnline_Req{} }
func (m *CS_PlayerOnline_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerOnline_Req) ProtoMessage()               {}
func (*CS_PlayerOnline_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// SC_PlayerOnline
type SC_PlayerOnline_Rsp struct {
	Ret            ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentify string    `protobuf:"bytes,2,opt,name=PlayerIdentify" json:"PlayerIdentify,omitempty"`
}

func (m *SC_PlayerOnline_Rsp) Reset()                    { *m = SC_PlayerOnline_Rsp{} }
func (m *SC_PlayerOnline_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerOnline_Rsp) ProtoMessage()               {}
func (*SC_PlayerOnline_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// CS_PlayerOffline
type CS_PlayerOffline_Req struct {
	PlayerIdentify string `protobuf:"bytes,1,opt,name=PlayerIdentify" json:"PlayerIdentify,omitempty"`
}

func (m *CS_PlayerOffline_Req) Reset()                    { *m = CS_PlayerOffline_Req{} }
func (m *CS_PlayerOffline_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerOffline_Req) ProtoMessage()               {}
func (*CS_PlayerOffline_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// SC_PlayerOffline
type SC_PlayerOffline_Rsp struct {
	Ret            ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentify string    `protobuf:"bytes,2,opt,name=PlayerIdentify" json:"PlayerIdentify,omitempty"`
}

func (m *SC_PlayerOffline_Rsp) Reset()                    { *m = SC_PlayerOffline_Rsp{} }
func (m *SC_PlayerOffline_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerOffline_Rsp) ProtoMessage()               {}
func (*SC_PlayerOffline_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// CS_GetBroadCastSessions
type CS_GetBroadCastSessions_Req struct {
	PlayerIdentifys []string `protobuf:"bytes,1,rep,name=PlayerIdentifys" json:"PlayerIdentifys,omitempty"`
}

func (m *CS_GetBroadCastSessions_Req) Reset()                    { *m = CS_GetBroadCastSessions_Req{} }
func (m *CS_GetBroadCastSessions_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_GetBroadCastSessions_Req) ProtoMessage()               {}
func (*CS_GetBroadCastSessions_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// SC_GetBroadCastSessions
type SC_GetBroadCastSessions_Rsp struct {
	Ret             ErrorCode `protobuf:"varint,1,opt,name=Ret,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentifys []string  `protobuf:"bytes,2,rep,name=PlayerIdentifys" json:"PlayerIdentifys,omitempty"`
}

func (m *SC_GetBroadCastSessions_Rsp) Reset()                    { *m = SC_GetBroadCastSessions_Rsp{} }
func (m *SC_GetBroadCastSessions_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_GetBroadCastSessions_Rsp) ProtoMessage()               {}
func (*SC_GetBroadCastSessions_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CS_PlayerOnline_Req)(nil), "MSG_CenterGate.CS_PlayerOnline_Req")
	proto.RegisterType((*SC_PlayerOnline_Rsp)(nil), "MSG_CenterGate.SC_PlayerOnline_Rsp")
	proto.RegisterType((*CS_PlayerOffline_Req)(nil), "MSG_CenterGate.CS_PlayerOffline_Req")
	proto.RegisterType((*SC_PlayerOffline_Rsp)(nil), "MSG_CenterGate.SC_PlayerOffline_Rsp")
	proto.RegisterType((*CS_GetBroadCastSessions_Req)(nil), "MSG_CenterGate.CS_GetBroadCastSessions_Req")
	proto.RegisterType((*SC_GetBroadCastSessions_Rsp)(nil), "MSG_CenterGate.SC_GetBroadCastSessions_Rsp")
	proto.RegisterEnum("MSG_CenterGate.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_CenterGate.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("MSG_CenterGate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xbb, 0xdb, 0x3f, 0xba, 0x23, 0xb4, 0x21, 0x5d, 0xb0, 0xd2, 0x4b, 0xe9, 0x41, 0x4a,
	0x85, 0x0a, 0x7a, 0xd6, 0x43, 0x83, 0x2e, 0x3d, 0x14, 0x65, 0x83, 0xe7, 0x25, 0x76, 0xa7, 0x12,
	0x5d, 0x92, 0x9a, 0x44, 0xa1, 0x9f, 0xc7, 0x2f, 0x2a, 0x5b, 0xd1, 0xd2, 0xb0, 0x3d, 0xf4, 0xe0,
	0x31, 0x2f, 0x79, 0x79, 0xbf, 0xc9, 0x4c, 0x20, 0x9e, 0xf3, 0x24, 0x63, 0xa8, 0x1c, 0x9a, 0x44,
	0x38, 0x9c, 0xac, 0x8c, 0x76, 0x9a, 0xb6, 0x77, 0xd5, 0xe1, 0x0d, 0x74, 0x19, 0xcf, 0x1e, 0x0b,
	0xb1, 0x46, 0xf3, 0xa0, 0x0a, 0xa9, 0x30, 0x4b, 0xf1, 0x9d, 0x9e, 0x43, 0xfb, 0x47, 0x9b, 0xe5,
	0xa8, 0x9c, 0x5c, 0xae, 0x7b, 0xc1, 0x20, 0x18, 0x45, 0xa9, 0xa7, 0x0e, 0x5f, 0xa1, 0xcb, 0x99,
	0x67, 0xb7, 0x2b, 0x7a, 0x01, 0xf5, 0x14, 0xdd, 0xc6, 0xd3, 0xbe, 0x3a, 0x9b, 0x78, 0x24, 0x77,
	0xc6, 0x68, 0xc3, 0x74, 0x8e, 0x69, 0x79, 0xaa, 0x22, 0x2b, 0xac, 0xcc, 0xba, 0x85, 0x78, 0x8b,
	0xba, 0x5c, 0x1e, 0xcc, 0xfa, 0x06, 0xf1, 0x96, 0xf5, 0xd7, 0xff, 0x5f, 0xb0, 0x09, 0xf4, 0x19,
	0xcf, 0x12, 0x74, 0x53, 0xa3, 0x45, 0xce, 0x84, 0x75, 0x1c, 0xad, 0x95, 0x5a, 0xd9, 0x0d, 0xf3,
	0x08, 0x3a, 0xbb, 0x06, 0xdb, 0x0b, 0x06, 0xf5, 0x51, 0x94, 0xfa, 0xf2, 0xd0, 0x41, 0x9f, 0xb3,
	0x3d, 0x17, 0x1d, 0x0a, 0x5f, 0x91, 0x1a, 0x56, 0xa6, 0x8e, 0xbf, 0x02, 0x68, 0xf1, 0xa7, 0xe9,
	0x9c, 0x27, 0x34, 0x82, 0xe6, 0x14, 0x5f, 0xa4, 0x22, 0x35, 0xda, 0x85, 0x8e, 0x37, 0x2c, 0x24,
	0x28, 0x45, 0x6f, 0x04, 0x48, 0x48, 0x63, 0x20, 0x7e, 0xaf, 0x48, 0xbd, 0x54, 0xfd, 0x0e, 0x90,
	0x06, 0xed, 0xc3, 0xe9, 0x9e, 0xa7, 0x22, 0xcd, 0x72, 0x73, 0x4f, 0xf9, 0xa4, 0x35, 0xbe, 0x84,
	0xe8, 0xaf, 0x42, 0x7a, 0x02, 0x47, 0x33, 0xf5, 0x29, 0x0a, 0x99, 0x93, 0x5a, 0xb9, 0xe0, 0x1f,
	0x8b, 0x05, 0x5a, 0x4b, 0x02, 0x7a, 0x0c, 0x8d, 0x7b, 0x21, 0x0b, 0x12, 0x3e, 0xb7, 0x36, 0x9f,
	0xe0, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x22, 0x9c, 0xb5, 0x1c, 0x03, 0x00, 0x00,
}
