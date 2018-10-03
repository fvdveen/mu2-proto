// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/encode/encode.proto

package encode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_Type int32

const (
	Error_NIL Error_Type = 0
)

var Error_Type_name = map[int32]string{
	0: "NIL",
}

var Error_Type_value = map[string]int32{
	"NIL": 0,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}

func (Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc5415be466f50e3, []int{0, 0}
}

type Error struct {
	Type                 Error_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Error_Type" json:"type,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Cause                *Error     `protobuf:"bytes,3,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5415be466f50e3, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_NIL
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCause() *Error {
	if m != nil {
		return m.Cause
	}
	return nil
}

type EncodeRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodeRequest) Reset()         { *m = EncodeRequest{} }
func (m *EncodeRequest) String() string { return proto.CompactTextString(m) }
func (*EncodeRequest) ProtoMessage()    {}
func (*EncodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5415be466f50e3, []int{1}
}

func (m *EncodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodeRequest.Unmarshal(m, b)
}
func (m *EncodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodeRequest.Marshal(b, m, deterministic)
}
func (m *EncodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodeRequest.Merge(m, src)
}
func (m *EncodeRequest) XXX_Size() int {
	return xxx_messageInfo_EncodeRequest.Size(m)
}
func (m *EncodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncodeRequest proto.InternalMessageInfo

func (m *EncodeRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type EncodeResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Opus                 []byte   `protobuf:"bytes,2,opt,name=opus,proto3" json:"opus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodeResponse) Reset()         { *m = EncodeResponse{} }
func (m *EncodeResponse) String() string { return proto.CompactTextString(m) }
func (*EncodeResponse) ProtoMessage()    {}
func (*EncodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5415be466f50e3, []int{2}
}

func (m *EncodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodeResponse.Unmarshal(m, b)
}
func (m *EncodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodeResponse.Marshal(b, m, deterministic)
}
func (m *EncodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodeResponse.Merge(m, src)
}
func (m *EncodeResponse) XXX_Size() int {
	return xxx_messageInfo_EncodeResponse.Size(m)
}
func (m *EncodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncodeResponse proto.InternalMessageInfo

func (m *EncodeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *EncodeResponse) GetOpus() []byte {
	if m != nil {
		return m.Opus
	}
	return nil
}

func init() {
	proto.RegisterEnum("Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*EncodeRequest)(nil), "EncodeRequest")
	proto.RegisterType((*EncodeResponse)(nil), "EncodeResponse")
}

func init() { proto.RegisterFile("proto/encode/encode.proto", fileDescriptor_cc5415be466f50e3) }

var fileDescriptor_cc5415be466f50e3 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x37, 0xb6, 0xdb, 0x65, 0x67, 0xb5, 0xbb, 0xcc, 0x29, 0x8a, 0x60, 0xcd, 0xa9, 0x17,
	0xa3, 0xd4, 0xbb, 0x07, 0x61, 0x0f, 0x82, 0x78, 0x88, 0xbe, 0xc0, 0x5a, 0x07, 0x11, 0xb4, 0x49,
	0x93, 0x46, 0xe8, 0xdb, 0x4b, 0xa7, 0x16, 0xb6, 0xa7, 0x4c, 0xbe, 0x4c, 0xe6, 0xfb, 0x19, 0x38,
	0x77, 0xde, 0x76, 0xf6, 0x96, 0x9a, 0xda, 0x7e, 0xd0, 0xff, 0xa1, 0x99, 0xa9, 0x16, 0x96, 0x7b,
	0xef, 0xad, 0xc7, 0x2b, 0x48, 0xbb, 0xde, 0x91, 0x14, 0x85, 0x28, 0xf3, 0x6a, 0xa3, 0x99, 0xea,
	0xb7, 0xde, 0x91, 0xe1, 0x07, 0x94, 0xb0, 0xfa, 0xa1, 0x10, 0x0e, 0x9f, 0x24, 0x4f, 0x0a, 0x51,
	0xae, 0xcd, 0x74, 0xc5, 0x4b, 0x58, 0xd6, 0x87, 0x18, 0x48, 0x26, 0x85, 0x28, 0x37, 0x55, 0x36,
	0xfe, 0x35, 0x23, 0x54, 0x5b, 0x48, 0x87, 0x29, 0xb8, 0x82, 0xe4, 0xe5, 0xe9, 0x79, 0xb7, 0x50,
	0xd7, 0x70, 0xb6, 0xe7, 0x08, 0x86, 0xda, 0x48, 0xa1, 0xc3, 0x1d, 0x24, 0xd1, 0x7f, 0xb3, 0x79,
	0x6d, 0x86, 0x52, 0x3d, 0x42, 0x3e, 0xb5, 0x04, 0x67, 0x9b, 0xc0, 0x0e, 0x1a, 0xa6, 0x72, 0xd7,
	0x91, 0x83, 0x21, 0x22, 0xa4, 0xd6, 0xc5, 0xc0, 0xc1, 0x4e, 0x0d, 0xd7, 0xd5, 0xc3, 0xa4, 0x79,
	0x25, 0xff, 0xfb, 0x55, 0x13, 0xde, 0x40, 0x36, 0x02, 0xcc, 0xf5, 0x2c, 0xc0, 0xc5, 0x56, 0xcf,
	0x6d, 0x6a, 0x71, 0x27, 0xde, 0x33, 0x5e, 0xd0, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xc6, 0x4e, 0x43, 0x3d, 0x01, 0x00, 0x00,
}