// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_manhua_chapter.proto

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

type SaverManhuaChapter struct {
	Id uint64 `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	// @inject_tag: gorm:"index:site"
	Site string `protobuf:"bytes,2,opt,name=Site" json:"Site,omitempty" gorm:"index:site"`
	// @inject_tag: gorm:"index:number"
	Number int32 `protobuf:"varint,3,opt,name=Number" json:"Number,omitempty" gorm:"index:number"`
	// @inject_tag: gorm:"index:chapter"
	Chapter string `protobuf:"bytes,4,opt,name=Chapter" json:"Chapter,omitempty" gorm:"index:chapter"`
	Title   string `protobuf:"bytes,5,opt,name=Title" json:"Title,omitempty"`
	Url     string `protobuf:"bytes,6,opt,name=Url" json:"Url,omitempty"`
	// @inject_tag: gorm:"column:FinishedTime"
	FinishedTime int64 `protobuf:"varint,7,opt,name=FinishedTime" json:"FinishedTime,omitempty" gorm:"column:FinishedTime"`
	Available    bool  `protobuf:"varint,8,opt,name=Available" json:"Available,omitempty"`
	// @inject_tag: gorm:"column:UpdateTime"
	UpdateTime           int64    `protobuf:"varint,9,opt,name=UpdateTime" json:"UpdateTime,omitempty" gorm:"column:UpdateTime"`
	// @inject_tag: gorm:"-"
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	// @inject_tag: gorm:"-"
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	// @inject_tag: gorm:"-"
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *SaverManhuaChapter) Reset()         { *m = SaverManhuaChapter{} }
func (m *SaverManhuaChapter) String() string { return proto.CompactTextString(m) }
func (*SaverManhuaChapter) ProtoMessage()    {}
func (*SaverManhuaChapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_manhua_chapter_2d8c1347e3c7e8b3, []int{0}
}
func (m *SaverManhuaChapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverManhuaChapter.Unmarshal(m, b)
}
func (m *SaverManhuaChapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverManhuaChapter.Marshal(b, m, deterministic)
}
func (dst *SaverManhuaChapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverManhuaChapter.Merge(dst, src)
}
func (m *SaverManhuaChapter) XXX_Size() int {
	return xxx_messageInfo_SaverManhuaChapter.Size(m)
}
func (m *SaverManhuaChapter) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverManhuaChapter.DiscardUnknown(m)
}

var xxx_messageInfo_SaverManhuaChapter proto.InternalMessageInfo

func (m *SaverManhuaChapter) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverManhuaChapter) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *SaverManhuaChapter) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SaverManhuaChapter) GetChapter() string {
	if m != nil {
		return m.Chapter
	}
	return ""
}

func (m *SaverManhuaChapter) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaverManhuaChapter) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SaverManhuaChapter) GetFinishedTime() int64 {
	if m != nil {
		return m.FinishedTime
	}
	return 0
}

func (m *SaverManhuaChapter) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *SaverManhuaChapter) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*SaverManhuaChapter)(nil), "protobuf.Saver_manhua_chapter")
}

func init() {
	proto.RegisterFile("protobuf/saver_manhua_chapter.proto", fileDescriptor_saver_manhua_chapter_2d8c1347e3c7e8b3)
}

var fileDescriptor_saver_manhua_chapter_2d8c1347e3c7e8b3 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xc9, 0xfe, 0xeb, 0xee, 0x20, 0x22, 0x43, 0x91, 0x39, 0x88, 0x84, 0x7a, 0xc9, 0x49,
	0x0f, 0x7e, 0x02, 0x11, 0x84, 0x5e, 0x3c, 0xa4, 0xed, 0xb9, 0x64, 0xdd, 0x91, 0x0d, 0x64, 0xdb,
	0x25, 0xcd, 0xf6, 0xa3, 0x7b, 0x16, 0xa7, 0x16, 0x15, 0xbc, 0xbd, 0xf7, 0xcb, 0xef, 0xe5, 0x30,
	0x70, 0x37, 0xc6, 0x7d, 0xda, 0xb7, 0xd3, 0xfb, 0xc3, 0xc1, 0x1d, 0x39, 0x6e, 0x07, 0xb7, 0xeb,
	0x27, 0xb7, 0x7d, 0xeb, 0xdd, 0x98, 0x38, 0xde, 0xcb, 0x2b, 0xd6, 0x67, 0x69, 0xf1, 0xa1, 0x60,
	0xbe, 0xfa, 0x47, 0xc4, 0x4b, 0xc8, 0x96, 0x1d, 0x29, 0xad, 0x4c, 0x61, 0xb3, 0x65, 0x87, 0x08,
	0xc5, 0xca, 0x27, 0xa6, 0x4c, 0x2b, 0xd3, 0x58, 0xc9, 0x78, 0x0d, 0xd5, 0xeb, 0x34, 0xb4, 0x1c,
	0x29, 0xd7, 0xca, 0x94, 0xf6, 0xbb, 0x21, 0xc1, 0xec, 0xf9, 0xf4, 0x0d, 0x15, 0xa2, 0x9f, 0x2b,
	0xce, 0xa1, 0x5c, 0xfb, 0x14, 0x98, 0x4a, 0xe1, 0xa7, 0x82, 0x57, 0x90, 0x6f, 0x62, 0xa0, 0x4a,
	0xd8, 0x57, 0xc4, 0x05, 0x5c, 0xbc, 0xf8, 0x9d, 0x3f, 0xf4, 0xdc, 0xad, 0xfd, 0xc0, 0x34, 0xd3,
	0xca, 0xe4, 0xf6, 0x0f, 0xc3, 0x1b, 0x68, 0x9e, 0x8e, 0xce, 0x07, 0xd7, 0x06, 0xa6, 0x5a, 0x2b,
	0x53, 0xdb, 0x1f, 0x80, 0xb7, 0x00, 0x9b, 0xb1, 0x73, 0x89, 0x65, 0xdf, 0xc8, 0xfe, 0x17, 0x69,
	0x2b, 0x39, 0xc1, 0xe3, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xd1, 0xf2, 0x09, 0x30, 0x01,
	0x00, 0x00,
}
