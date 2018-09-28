// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
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

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Token)(nil), "Token")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0xc9, 0x11, 0x1f, 0x61, 0x23, 0x28, 0x56, 0x14, 0xa7, 0x88, 0x02, 0x5c, 0x51, 0xb9,
	0x08, 0xe2, 0x11, 0x68, 0x10, 0x42, 0xc2, 0xfc, 0xf4, 0x26, 0x5e, 0x45, 0x16, 0xe4, 0x1c, 0xad,
	0x8f, 0xf0, 0xfa, 0x68, 0xd7, 0x77, 0xd0, 0xed, 0x37, 0x33, 0x1a, 0x8d, 0x0d, 0xf0, 0x5d, 0x88,
	0xdd, 0x9e, 0xf3, 0x90, 0xed, 0x13, 0xcc, 0xdf, 0x0a, 0x31, 0x9e, 0x43, 0x93, 0x62, 0x37, 0xbb,
	0x9a, 0xdd, 0x9c, 0xfa, 0x26, 0x45, 0x5c, 0xc1, 0x42, 0x52, 0x7d, 0xd8, 0x51, 0xd7, 0xa8, 0xfa,
	0xc7, 0xe2, 0xed, 0x43, 0x29, 0x3f, 0x99, 0x63, 0x77, 0x5c, 0xbd, 0x89, 0xed, 0x33, 0x98, 0xd7,
	0xfc, 0x49, 0x3d, 0x5e, 0x80, 0x19, 0xe4, 0x18, 0x3b, 0x2b, 0x88, 0x7a, 0x08, 0x5f, 0x29, 0x6a,
	0xe7, 0xc2, 0x57, 0xc0, 0x4b, 0x30, 0xc4, 0x9c, 0x59, 0xdb, 0x96, 0xeb, 0xd6, 0xdd, 0x0b, 0xf9,
	0x2a, 0xda, 0x3b, 0x30, 0xca, 0x88, 0x30, 0xdf, 0xe4, 0x48, 0xda, 0x68, 0xbc, 0xde, 0xd8, 0xc1,
	0xc9, 0x8e, 0x4a, 0x09, 0xdb, 0x69, 0xe6, 0x84, 0xeb, 0x07, 0x58, 0xca, 0xcb, 0x5e, 0x88, 0x0f,
	0x69, 0x23, 0x41, 0xf3, 0x98, 0xb7, 0xa9, 0x47, 0xe3, 0x44, 0x5e, 0xb5, 0x4e, 0x77, 0xda, 0x23,
	0xbc, 0x86, 0xb3, 0x77, 0x99, 0x11, 0x06, 0xaa, 0xd3, 0x47, 0xeb, 0x3f, 0xf2, 0xd1, 0xea, 0x67,
	0xdd, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x39, 0xeb, 0x29, 0xfa, 0x3a, 0x01, 0x00, 0x00,
}