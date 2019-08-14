// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_article.proto

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

type SaverArticle struct {
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: gorm:"index:site,site_num"
	Site string `protobuf:"bytes,2,opt,name=Site,proto3" json:"Site,omitempty" gorm:"index:site,site_num"`
	// @inject_tag: gorm:"index:number,site_num"
	Number string `protobuf:"bytes,3,opt,name=Number,proto3" json:"Number,omitempty" gorm:"index:number,site_num"`
	// @inject_tag: gorm:"type:text"
	Face string `protobuf:"bytes,4,opt,name=Face,proto3" json:"Face,omitempty" gorm:"type:text"`
	// @inject_tag: gorm:"type:text"
	Title string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty" gorm:"type:text"`
	// @inject_tag: gorm:"column:ZhTitle;type:text"
	ZhTitle string `protobuf:"bytes,6,opt,name=ZhTitle,proto3" json:"ZhTitle,omitempty" gorm:"column:ZhTitle;type:text"`
	Url     string `protobuf:"bytes,7,opt,name=Url,proto3" json:"Url,omitempty"`
	// @inject_tag: gorm:"type:text"
	Contents string `protobuf:"bytes,8,opt,name=Contents,proto3" json:"Contents,omitempty" gorm:"type:text"`
	// @inject_tag: gorm:"column:ZhContents;type:text"
	ZhContents string `protobuf:"bytes,9,opt,name=ZhContents,proto3" json:"ZhContents,omitempty" gorm:"column:ZhContents;type:text"`
	Tags       string `protobuf:"bytes,10,opt,name=Tags,proto3" json:"Tags,omitempty"`
	// @inject_tag: gorm:"column:ZhTags"
	ZhTags string `protobuf:"bytes,11,opt,name=ZhTags,proto3" json:"ZhTags,omitempty" gorm:"column:ZhTags"`
	// @inject_tag: gorm:"column:updatetime"
	UpdateTime           string   `protobuf:"bytes,12,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty" gorm:"column:updatetime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *SaverArticle) Reset()         { *m = SaverArticle{} }
func (m *SaverArticle) String() string { return proto.CompactTextString(m) }
func (*SaverArticle) ProtoMessage()    {}
func (*SaverArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_article_ecf9d4f28de69101, []int{0}
}
func (m *SaverArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverArticle.Unmarshal(m, b)
}
func (m *SaverArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverArticle.Marshal(b, m, deterministic)
}
func (dst *SaverArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverArticle.Merge(dst, src)
}
func (m *SaverArticle) XXX_Size() int {
	return xxx_messageInfo_SaverArticle.Size(m)
}
func (m *SaverArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverArticle.DiscardUnknown(m)
}

var xxx_messageInfo_SaverArticle proto.InternalMessageInfo

func (m *SaverArticle) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverArticle) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *SaverArticle) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SaverArticle) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *SaverArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaverArticle) GetZhTitle() string {
	if m != nil {
		return m.ZhTitle
	}
	return ""
}

func (m *SaverArticle) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SaverArticle) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *SaverArticle) GetZhContents() string {
	if m != nil {
		return m.ZhContents
	}
	return ""
}

func (m *SaverArticle) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *SaverArticle) GetZhTags() string {
	if m != nil {
		return m.ZhTags
	}
	return ""
}

func (m *SaverArticle) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func init() {
	proto.RegisterType((*SaverArticle)(nil), "protobuf.Saver_article")
}

func init() {
	proto.RegisterFile("protobuf/saver_article.proto", fileDescriptor_saver_article_ecf9d4f28de69101)
}

var fileDescriptor_saver_article_ecf9d4f28de69101 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xc9, 0x36, 0x4d, 0xd3, 0xf1, 0x0f, 0x32, 0x88, 0x0c, 0x22, 0x52, 0x3c, 0xf5, 0xa4,
	0x07, 0x3f, 0x82, 0x20, 0xf4, 0xe2, 0xa1, 0x4d, 0x2f, 0xb9, 0xc8, 0xa6, 0x19, 0xcd, 0x42, 0xda,
	0x84, 0xcd, 0xc6, 0xef, 0xe4, 0xb7, 0x94, 0x9d, 0x35, 0x69, 0x6e, 0xef, 0xfd, 0xde, 0x63, 0x78,
	0x0c, 0x3c, 0xb4, 0xb6, 0x71, 0x4d, 0xd1, 0x7f, 0xbd, 0x74, 0xfa, 0x87, 0xed, 0xa7, 0xb6, 0xce,
	0x1c, 0x6a, 0x7e, 0x16, 0x8c, 0xe9, 0x90, 0x3e, 0xfd, 0x2a, 0xb8, 0xda, 0x4d, 0x1b, 0x78, 0x0d,
	0x6a, 0x53, 0x52, 0xb4, 0x8a, 0xd6, 0xf1, 0x56, 0x6d, 0x4a, 0x44, 0x88, 0x77, 0xc6, 0x31, 0xa9,
	0x55, 0xb4, 0x5e, 0x6e, 0x45, 0xe3, 0x1d, 0x24, 0x1f, 0xfd, 0xb1, 0x60, 0x4b, 0x33, 0xa1, 0xff,
	0xce, 0x77, 0xdf, 0xf5, 0x81, 0x29, 0x0e, 0x5d, 0xaf, 0xf1, 0x16, 0xe6, 0x99, 0x71, 0x35, 0xd3,
	0x5c, 0x60, 0x30, 0x48, 0xb0, 0xc8, 0xab, 0xc0, 0x13, 0xe1, 0x83, 0xc5, 0x1b, 0x98, 0xed, 0x6d,
	0x4d, 0x0b, 0xa1, 0x5e, 0xe2, 0x3d, 0xa4, 0x6f, 0xcd, 0xc9, 0xf1, 0xc9, 0x75, 0x94, 0x0a, 0x1e,
	0x3d, 0x3e, 0x02, 0xe4, 0xd5, 0x98, 0x2e, 0x25, 0x9d, 0x10, 0xbf, 0x28, 0xd3, 0xdf, 0x1d, 0x41,
	0x58, 0xe4, 0xb5, 0x5f, 0x9f, 0x57, 0x42, 0x2f, 0xc2, 0xfa, 0xe0, 0xfc, 0xad, 0x7d, 0x5b, 0x6a,
	0xc7, 0x99, 0x39, 0x32, 0x5d, 0x86, 0x5b, 0x67, 0x52, 0x24, 0xf2, 0xb5, 0xd7, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0xc0, 0x20, 0x38, 0x5c, 0x01, 0x00, 0x00,
}
