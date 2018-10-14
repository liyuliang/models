// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/indexer_manhua_chapter.proto

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

type IndexManhuaChapter struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Site                 string   `protobuf:"bytes,2,opt,name=Site" json:"Site,omitempty"`
	Number               int32    `protobuf:"varint,3,opt,name=Number" json:"Number,omitempty"`
	Chapter              string   `protobuf:"bytes,4,opt,name=Chapter" json:"Chapter,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=Title" json:"Title,omitempty"`
	Update               int64    `protobuf:"varint,6,opt,name=Update" json:"Update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `gorm:"-" json:"-"`
	XXX_unrecognized     []byte   `gorm:"-" json:"-"`
	XXX_sizecache        int32    `gorm:"-" json:"-"`
}

func (m *IndexManhuaChapter) Reset()         { *m = IndexManhuaChapter{} }
func (m *IndexManhuaChapter) String() string { return proto.CompactTextString(m) }
func (*IndexManhuaChapter) ProtoMessage()    {}
func (*IndexManhuaChapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_indexer_manhua_chapter_81eae89eeed5eb78, []int{0}
}
func (m *IndexManhuaChapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexManhuaChapter.Unmarshal(m, b)
}
func (m *IndexManhuaChapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexManhuaChapter.Marshal(b, m, deterministic)
}
func (dst *IndexManhuaChapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexManhuaChapter.Merge(dst, src)
}
func (m *IndexManhuaChapter) XXX_Size() int {
	return xxx_messageInfo_IndexManhuaChapter.Size(m)
}
func (m *IndexManhuaChapter) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexManhuaChapter.DiscardUnknown(m)
}

var xxx_messageInfo_IndexManhuaChapter proto.InternalMessageInfo

func (m *IndexManhuaChapter) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IndexManhuaChapter) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *IndexManhuaChapter) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *IndexManhuaChapter) GetChapter() string {
	if m != nil {
		return m.Chapter
	}
	return ""
}

func (m *IndexManhuaChapter) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *IndexManhuaChapter) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

func init() {
	proto.RegisterType((*IndexManhuaChapter)(nil), "protobuf.Index_manhua_chapter")
}

func init() {
	proto.RegisterFile("protobuf/indexer_manhua_chapter.proto", fileDescriptor_indexer_manhua_chapter_81eae89eeed5eb78)
}

var fileDescriptor_indexer_manhua_chapter_81eae89eeed5eb78 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x48, 0x2d, 0x8a, 0xcf, 0x4d, 0xcc,
	0xcb, 0x28, 0x4d, 0x8c, 0x4f, 0xce, 0x48, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0x03, 0xcb, 0x0b, 0x71,
	0xc0, 0x94, 0x29, 0xcd, 0x62, 0xe4, 0x12, 0xf1, 0x04, 0x29, 0x45, 0x53, 0x28, 0xc4, 0xc7, 0xc5,
	0xe4, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0xe4, 0x99, 0x22, 0x24, 0xc4, 0xc5,
	0x12, 0x9c, 0x59, 0x92, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x71,
	0xb1, 0xf9, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79,
	0x42, 0x12, 0x5c, 0xec, 0xce, 0x10, 0x63, 0x24, 0x58, 0xc0, 0xca, 0x61, 0x5c, 0x21, 0x11, 0x2e,
	0xd6, 0x90, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x56, 0xb0, 0x38, 0x84, 0x03, 0x32, 0x27, 0xb4, 0x20,
	0x25, 0xb1, 0x24, 0x55, 0x82, 0x4d, 0x81, 0x51, 0x83, 0x39, 0x08, 0xca, 0x4b, 0x62, 0x03, 0x3b,
	0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x94, 0xab, 0x0e, 0xd6, 0x00, 0x00, 0x00,
}