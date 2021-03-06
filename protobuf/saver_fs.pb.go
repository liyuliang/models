// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_fs.proto

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

type SaverFs struct {
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: gorm:"not null;unique"
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty" gorm:"not null;unique"`
	Url                  string   `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *SaverFs) Reset()         { *m = SaverFs{} }
func (m *SaverFs) String() string { return proto.CompactTextString(m) }
func (*SaverFs) ProtoMessage()    {}
func (*SaverFs) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_fs_9b5eb01f32a65b7a, []int{0}
}
func (m *SaverFs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverFs.Unmarshal(m, b)
}
func (m *SaverFs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverFs.Marshal(b, m, deterministic)
}
func (dst *SaverFs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverFs.Merge(dst, src)
}
func (m *SaverFs) XXX_Size() int {
	return xxx_messageInfo_SaverFs.Size(m)
}
func (m *SaverFs) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverFs.DiscardUnknown(m)
}

var xxx_messageInfo_SaverFs proto.InternalMessageInfo

func (m *SaverFs) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverFs) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SaverFs) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*SaverFs)(nil), "protobuf.Saver_fs")
}

func init() { proto.RegisterFile("protobuf/saver_fs.proto", fileDescriptor_saver_fs_9b5eb01f32a65b7a) }

var fileDescriptor_saver_fs_9b5eb01f32a65b7a = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2c, 0x4b, 0x2d, 0x8a, 0x4f, 0x2b, 0xd6, 0x03, 0x8b,
	0x08, 0x71, 0xc0, 0x24, 0x94, 0x1c, 0xb8, 0x38, 0x82, 0xa1, 0x72, 0x42, 0x7c, 0x5c, 0x4c, 0x9e,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x9e, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0xce,
	0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73,
	0x68, 0x51, 0x8e, 0x04, 0x33, 0x58, 0x08, 0xc4, 0x4c, 0x62, 0x03, 0x9b, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xac, 0x45, 0x66, 0x6d, 0x00, 0x00, 0x00,
}
