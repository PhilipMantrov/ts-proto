// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package pb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos                int64    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type Duration struct {
	Nanos                int64    `protobuf:"varint,1,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

// Empty is only used as a message for rpc calls that
// return no data.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// An optional string value used for RPCs that update a record.
// If the Optional is undefined, the updating code will not change
// the field.
type OptString struct {
	Val                  string   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OptString) Reset()         { *m = OptString{} }
func (m *OptString) String() string { return proto.CompactTextString(m) }
func (*OptString) ProtoMessage()    {}
func (*OptString) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *OptString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptString.Unmarshal(m, b)
}
func (m *OptString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptString.Marshal(b, m, deterministic)
}
func (m *OptString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptString.Merge(m, src)
}
func (m *OptString) XXX_Size() int {
	return xxx_messageInfo_OptString.Size(m)
}
func (m *OptString) XXX_DiscardUnknown() {
	xxx_messageInfo_OptString.DiscardUnknown(m)
}

var xxx_messageInfo_OptString proto.InternalMessageInfo

func (m *OptString) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type OptInt64 struct {
	Val                  int64    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OptInt64) Reset()         { *m = OptInt64{} }
func (m *OptInt64) String() string { return proto.CompactTextString(m) }
func (*OptInt64) ProtoMessage()    {}
func (*OptInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *OptInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptInt64.Unmarshal(m, b)
}
func (m *OptInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptInt64.Marshal(b, m, deterministic)
}
func (m *OptInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptInt64.Merge(m, src)
}
func (m *OptInt64) XXX_Size() int {
	return xxx_messageInfo_OptInt64.Size(m)
}
func (m *OptInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_OptInt64.DiscardUnknown(m)
}

var xxx_messageInfo_OptInt64 proto.InternalMessageInfo

func (m *OptInt64) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

type OptBool struct {
	Val                  bool     `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OptBool) Reset()         { *m = OptBool{} }
func (m *OptBool) String() string { return proto.CompactTextString(m) }
func (*OptBool) ProtoMessage()    {}
func (*OptBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}

func (m *OptBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptBool.Unmarshal(m, b)
}
func (m *OptBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptBool.Marshal(b, m, deterministic)
}
func (m *OptBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptBool.Merge(m, src)
}
func (m *OptBool) XXX_Size() int {
	return xxx_messageInfo_OptBool.Size(m)
}
func (m *OptBool) XXX_DiscardUnknown() {
	xxx_messageInfo_OptBool.DiscardUnknown(m)
}

var xxx_messageInfo_OptBool proto.InternalMessageInfo

func (m *OptBool) GetVal() bool {
	if m != nil {
		return m.Val
	}
	return false
}

type IPNet struct {
	Ip                   []byte   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Mask                 []byte   `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPNet) Reset()         { *m = IPNet{} }
func (m *IPNet) String() string { return proto.CompactTextString(m) }
func (*IPNet) ProtoMessage()    {}
func (*IPNet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6}
}

func (m *IPNet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPNet.Unmarshal(m, b)
}
func (m *IPNet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPNet.Marshal(b, m, deterministic)
}
func (m *IPNet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPNet.Merge(m, src)
}
func (m *IPNet) XXX_Size() int {
	return xxx_messageInfo_IPNet.Size(m)
}
func (m *IPNet) XXX_DiscardUnknown() {
	xxx_messageInfo_IPNet.DiscardUnknown(m)
}

var xxx_messageInfo_IPNet proto.InternalMessageInfo

func (m *IPNet) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *IPNet) GetMask() []byte {
	if m != nil {
		return m.Mask
	}
	return nil
}

type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{7}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "pb.Timestamp")
	proto.RegisterType((*Duration)(nil), "pb.Duration")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*OptString)(nil), "pb.OptString")
	proto.RegisterType((*OptInt64)(nil), "pb.OptInt64")
	proto.RegisterType((*OptBool)(nil), "pb.OptBool")
	proto.RegisterType((*IPNet)(nil), "pb.IPNet")
	proto.RegisterType((*ID)(nil), "pb.ID")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0x6a, 0x4d, 0x3b, 0x2e, 0x22, 0x61, 0x0f, 0x05, 0x15, 0x96, 0x9c, 0x04, 0xa1,
	0x1e, 0x14, 0x2f, 0xde, 0x96, 0xf5, 0xd0, 0x8b, 0x95, 0xe8, 0xc9, 0x5b, 0x6a, 0x83, 0x04, 0x37,
	0xc9, 0xd0, 0x8c, 0x42, 0xff, 0xbd, 0x98, 0x2a, 0xdb, 0xdb, 0x7b, 0xf3, 0x7d, 0x30, 0xc3, 0xc0,
	0x09, 0x4d, 0x68, 0x62, 0x83, 0x63, 0xa0, 0x20, 0x18, 0xf6, 0xf2, 0x01, 0xaa, 0x57, 0xeb, 0x4c,
	0x24, 0xed, 0x50, 0xd4, 0xc0, 0xa3, 0x79, 0x0f, 0x7e, 0x88, 0x75, 0xb6, 0xc9, 0xae, 0x72, 0xf5,
	0x5f, 0xc5, 0x1a, 0x0a, 0xaf, 0x7d, 0x88, 0x35, 0x4b, 0xf3, 0xb9, 0xc8, 0x0d, 0x94, 0xbb, 0xaf,
	0x51, 0x93, 0x0d, 0xfe, 0x60, 0x64, 0x4b, 0x83, 0x43, 0xf1, 0xe8, 0x90, 0x26, 0x79, 0x09, 0x55,
	0x87, 0xf4, 0x42, 0xa3, 0xf5, 0x1f, 0xe2, 0x0c, 0xf2, 0x6f, 0xbd, 0x4f, 0x66, 0xa5, 0x7e, 0xa3,
	0xbc, 0x80, 0xb2, 0x43, 0x6a, 0x3d, 0xdd, 0xdf, 0x2d, 0x69, 0x3e, 0xd3, 0x73, 0xe0, 0x1d, 0xd2,
	0x36, 0x84, 0xfd, 0x12, 0x96, 0x33, 0xbc, 0x86, 0xa2, 0x7d, 0x7e, 0x32, 0x24, 0x4e, 0x81, 0x59,
	0x4c, 0x64, 0xa5, 0x98, 0x45, 0x21, 0xe0, 0xc8, 0xe9, 0xf8, 0x99, 0x4e, 0x5e, 0xa9, 0x94, 0xe5,
	0x1a, 0x58, 0xbb, 0x4b, 0xe6, 0xf0, 0xb7, 0x9e, 0xd9, 0x61, 0xcb, 0xdf, 0x8a, 0xa6, 0xb9, 0xc1,
	0xbe, 0x3f, 0x4e, 0x8f, 0xb9, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x60, 0x40, 0x76, 0x82, 0x27,
	0x01, 0x00, 0x00,
}
