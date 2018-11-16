// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/cacher_image.proto

package protobuf

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

type CacherImage struct {
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: gorm:"index:code,code"
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CacherImage) Reset()         { *m = CacherImage{} }
func (m *CacherImage) String() string { return proto.CompactTextString(m) }
func (*CacherImage) ProtoMessage()    {}
func (*CacherImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cacher_image_23dd616b07b8d286, []int{0}
}
func (m *CacherImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacherImage.Unmarshal(m, b)
}
func (m *CacherImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacherImage.Marshal(b, m, deterministic)
}
func (dst *CacherImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacherImage.Merge(dst, src)
}
func (m *CacherImage) XXX_Size() int {
	return xxx_messageInfo_CacherImage.Size(m)
}
func (m *CacherImage) XXX_DiscardUnknown() {
	xxx_messageInfo_CacherImage.DiscardUnknown(m)
}

var xxx_messageInfo_CacherImage proto.InternalMessageInfo

func (m *CacherImage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CacherImage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CacherImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*CacherImage)(nil), "protobuf.Cacher_image")
}

func init() {
	proto.RegisterFile("protobuf/cacher_image.proto", fileDescriptor_cacher_image_23dd616b07b8d286)
}

var fileDescriptor_cacher_image_23dd616b07b8d286 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x4e, 0x4c, 0xce, 0x48, 0x2d, 0x8a, 0xcf, 0xcc, 0x4d, 0x4c,
	0x4f, 0xd5, 0x03, 0x8b, 0x0a, 0x71, 0xc0, 0x24, 0x95, 0x5c, 0xb8, 0x78, 0x9c, 0x91, 0xe4, 0x85,
	0xf8, 0xb8, 0x98, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x3c, 0x53, 0x84,
	0x84, 0xb8, 0x58, 0x9c, 0xf3, 0x53, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x21, 0x01, 0x2e, 0xe6, 0xd0, 0xa2, 0x1c, 0x09, 0x66, 0xb0, 0x10, 0x88, 0x99, 0xc4, 0x06, 0x36,
	0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xf6, 0xb7, 0x6c, 0x75, 0x00, 0x00, 0x00,
}
