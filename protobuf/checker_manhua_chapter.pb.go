// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/checker_manhua_chapter.proto

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

type CheckerManhuaChapter struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Site                 string   `protobuf:"bytes,2,opt,name=Site,proto3" json:"Site,omitempty"`
	Number               string   `protobuf:"bytes,3,opt,name=Number,proto3" json:"Number,omitempty"`
	Chapter              string   `protobuf:"bytes,4,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=Url,proto3" json:"Url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckerManhuaChapter) Reset()         { *m = CheckerManhuaChapter{} }
func (m *CheckerManhuaChapter) String() string { return proto.CompactTextString(m) }
func (*CheckerManhuaChapter) ProtoMessage()    {}
func (*CheckerManhuaChapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_checker_manhua_chapter_f19ed4e0bb73529d, []int{0}
}
func (m *CheckerManhuaChapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckerManhuaChapter.Unmarshal(m, b)
}
func (m *CheckerManhuaChapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckerManhuaChapter.Marshal(b, m, deterministic)
}
func (dst *CheckerManhuaChapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckerManhuaChapter.Merge(dst, src)
}
func (m *CheckerManhuaChapter) XXX_Size() int {
	return xxx_messageInfo_CheckerManhuaChapter.Size(m)
}
func (m *CheckerManhuaChapter) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckerManhuaChapter.DiscardUnknown(m)
}

var xxx_messageInfo_CheckerManhuaChapter proto.InternalMessageInfo

func (m *CheckerManhuaChapter) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CheckerManhuaChapter) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *CheckerManhuaChapter) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *CheckerManhuaChapter) GetChapter() string {
	if m != nil {
		return m.Chapter
	}
	return ""
}

func (m *CheckerManhuaChapter) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckerManhuaChapter)(nil), "protobuf.Checker_manhua_chapter")
}

func init() {
	proto.RegisterFile("protobuf/checker_manhua_chapter.proto", fileDescriptor_checker_manhua_chapter_f19ed4e0bb73529d)
}

var fileDescriptor_checker_manhua_chapter_f19ed4e0bb73529d = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0x8a, 0xcf, 0x4d, 0xcc,
	0xcb, 0x28, 0x4d, 0x8c, 0x4f, 0xce, 0x48, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0x03, 0xcb, 0x0b, 0x71,
	0xc0, 0x94, 0x29, 0x35, 0x30, 0x72, 0x89, 0x39, 0x63, 0x55, 0x2a, 0xc4, 0xc7, 0xc5, 0xe4, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0xe4, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x12, 0x9c,
	0x59, 0x92, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x71, 0xb1, 0xf9,
	0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x83, 0x45, 0xa1, 0x3c, 0x21, 0x09, 0x2e, 0x76, 0x67,
	0x88, 0x31, 0x12, 0x2c, 0x60, 0x09, 0x18, 0x57, 0x48, 0x80, 0x8b, 0x39, 0xb4, 0x28, 0x47, 0x82,
	0x15, 0x2c, 0x0a, 0x62, 0x26, 0xb1, 0x81, 0x1d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0x39, 0x6d, 0x9a, 0xbc, 0x00, 0x00, 0x00,
}