// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	const.proto
	gate.proto
	gs.proto

It has these top-level messages:
	NULL
	LOGIN_REQ
	LOGIN_ACK
	LOGOUT_REQ
	COMMON_ACK
	HEART_REQ
	TEST_GATE_REQ
	MSG
*/
package protocol

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

// 协议号
type PROTO int32

const (
	PROTO_TEST_GATE PROTO = 0
	PROTO_HEART     PROTO = 1
	PROTO_LOGIN     PROTO = 2
	PROTO_LOGOUT    PROTO = 3
	PROTO_COMMON    PROTO = 4
)

var PROTO_name = map[int32]string{
	0: "TEST_GATE",
	1: "HEART",
	2: "LOGIN",
	3: "LOGOUT",
	4: "COMMON",
}
var PROTO_value = map[string]int32{
	"TEST_GATE": 0,
	"HEART":     1,
	"LOGIN":     2,
	"LOGOUT":    3,
	"COMMON":    4,
}

func (x PROTO) String() string {
	return proto.EnumName(PROTO_name, int32(x))
}
func (PROTO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 错误码
type ERR int32

const (
	ERR_OK               ERR = 0
	ERR_USER_NOT_LOGIN   ERR = 1
	ERR_SERVER_NOT_EXIST ERR = 2
	ERR_DATA             ERR = 3
	ERR_USER_LOGOUT      ERR = 4
)

var ERR_name = map[int32]string{
	0: "OK",
	1: "USER_NOT_LOGIN",
	2: "SERVER_NOT_EXIST",
	3: "DATA",
	4: "USER_LOGOUT",
}
var ERR_value = map[string]int32{
	"OK":               0,
	"USER_NOT_LOGIN":   1,
	"SERVER_NOT_EXIST": 2,
	"DATA":             3,
	"USER_LOGOUT":      4,
}

func (x ERR) String() string {
	return proto.EnumName(ERR_name, int32(x))
}
func (ERR) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterEnum("protocol.PROTO", PROTO_name, PROTO_value)
	proto.RegisterEnum("protocol.ERR", ERR_name, ERR_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xcd, 0x0a, 0x82, 0x40,
	0x14, 0x85, 0xfd, 0x47, 0xaf, 0x54, 0x97, 0x4b, 0x4f, 0xe1, 0xa2, 0x4d, 0x4f, 0x30, 0xe4, 0x60,
	0x92, 0x7a, 0x63, 0xe6, 0x1a, 0xed, 0x84, 0xdc, 0x46, 0x13, 0xe5, 0xfb, 0x13, 0x9a, 0xab, 0xf3,
	0x71, 0xe0, 0x3b, 0x1c, 0xc8, 0x47, 0xf7, 0xfa, 0x4e, 0x87, 0xf7, 0xc7, 0x4d, 0x8e, 0xd2, 0x25,
	0x46, 0xf7, 0x2c, 0x4a, 0x88, 0xaf, 0x86, 0x85, 0x69, 0x03, 0x99, 0x68, 0x2b, 0x43, 0xa5, 0x44,
	0xa3, 0x47, 0x19, 0xc4, 0x67, 0xad, 0x8c, 0xa0, 0x3f, 0x63, 0xc3, 0x55, 0xdd, 0x61, 0x40, 0x00,
	0x49, 0xc3, 0x15, 0xf7, 0x82, 0xe1, 0xcc, 0x27, 0x6e, 0x5b, 0xee, 0x30, 0x2a, 0x0c, 0x84, 0xda,
	0x18, 0x4a, 0x20, 0xe0, 0x0b, 0x7a, 0x44, 0xb0, 0xed, 0xad, 0x36, 0x43, 0xc7, 0x32, 0xfc, 0x55,
	0x9f, 0xf6, 0x80, 0x56, 0x9b, 0xdb, 0xda, 0xea, 0x7b, 0x6d, 0x05, 0x03, 0x4a, 0x21, 0x2a, 0x95,
	0x28, 0x0c, 0x69, 0x07, 0xf9, 0xe2, 0xac, 0xfb, 0xd1, 0x23, 0x59, 0x3e, 0x1e, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x1e, 0x3c, 0x54, 0x4a, 0xb9, 0x00, 0x00, 0x00,
}
