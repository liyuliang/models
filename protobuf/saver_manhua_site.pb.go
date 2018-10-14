// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_manhua_site.proto

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

type SaverManhuaSite struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Site                 string   `protobuf:"bytes,2,opt,name=Site" json:"Site,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=Url" json:"Url,omitempty"`
	Number               int32    `protobuf:"varint,5,opt,name=Number" json:"Number,omitempty"`
	Categories           string   `protobuf:"bytes,6,opt,name=Categories" json:"Categories,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=Description" json:"Description,omitempty"`
	Face                 string   `protobuf:"bytes,8,opt,name=Face" json:"Face,omitempty"`
	UpdateTime           int64    `protobuf:"varint,9,opt,name=UpdateTime" json:"UpdateTime,omitempty"`
	CheckTime            int64    `protobuf:"varint,10,opt,name=CheckTime" json:"CheckTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `gorm:"-" json:"-"`
	XXX_unrecognized     []byte   `gorm:"-" json:"-"`
	XXX_sizecache        int32    `gorm:"-" json:"-"`
}

func (m *SaverManhuaSite) Reset()         { *m = SaverManhuaSite{} }
func (m *SaverManhuaSite) String() string { return proto.CompactTextString(m) }
func (*SaverManhuaSite) ProtoMessage()    {}
func (*SaverManhuaSite) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_manhua_site_b24d01f978e884f3, []int{0}
}
func (m *SaverManhuaSite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverManhuaSite.Unmarshal(m, b)
}
func (m *SaverManhuaSite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverManhuaSite.Marshal(b, m, deterministic)
}
func (dst *SaverManhuaSite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverManhuaSite.Merge(dst, src)
}
func (m *SaverManhuaSite) XXX_Size() int {
	return xxx_messageInfo_SaverManhuaSite.Size(m)
}
func (m *SaverManhuaSite) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverManhuaSite.DiscardUnknown(m)
}

var xxx_messageInfo_SaverManhuaSite proto.InternalMessageInfo

func (m *SaverManhuaSite) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverManhuaSite) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *SaverManhuaSite) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SaverManhuaSite) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SaverManhuaSite) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SaverManhuaSite) GetCategories() string {
	if m != nil {
		return m.Categories
	}
	return ""
}

func (m *SaverManhuaSite) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SaverManhuaSite) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *SaverManhuaSite) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SaverManhuaSite) GetCheckTime() int64 {
	if m != nil {
		return m.CheckTime
	}
	return 0
}

func init() {
	proto.RegisterType((*SaverManhuaSite)(nil), "protobuf.Saver_manhua_site")
}

func init() {
	proto.RegisterFile("protobuf/saver_manhua_site.proto", fileDescriptor_saver_manhua_site_b24d01f978e884f3)
}

var fileDescriptor_saver_manhua_site_b24d01f978e884f3 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x24, 0x0d, 0xcd, 0x21, 0x21, 0xb8, 0x01, 0xdd, 0x80, 0x90, 0xc5, 0x94, 0x09,
	0x06, 0x1e, 0xa1, 0x08, 0xa9, 0x4b, 0x87, 0x94, 0xce, 0x95, 0x93, 0x1c, 0xd4, 0x82, 0xd4, 0x91,
	0xed, 0xf0, 0x0e, 0xbc, 0x35, 0xf2, 0x15, 0x44, 0xc4, 0xf6, 0xf9, 0xfb, 0x7f, 0xfb, 0x97, 0x0c,
	0x7a, 0xf4, 0x2e, 0xba, 0x76, 0x7a, 0x7d, 0x08, 0xe6, 0x93, 0xfd, 0x7e, 0x30, 0xc7, 0xc3, 0x64,
	0xf6, 0xc1, 0x46, 0xbe, 0x97, 0x08, 0x97, 0xbf, 0x8d, 0xbb, 0xaf, 0x0c, 0xae, 0xb6, 0xff, 0x5b,
	0x78, 0x01, 0xd9, 0xba, 0x27, 0xa5, 0x55, 0x5d, 0x34, 0xd9, 0xba, 0x47, 0x84, 0x62, 0x6b, 0x23,
	0x53, 0xa6, 0x55, 0x5d, 0x35, 0xc2, 0xc9, 0x6d, 0xcc, 0xc0, 0x94, 0x9f, 0x5c, 0x62, 0xbc, 0x84,
	0x7c, 0xe7, 0x3f, 0xa8, 0x10, 0x95, 0x10, 0xaf, 0xa1, 0xdc, 0x4c, 0x43, 0xcb, 0x9e, 0x16, 0x5a,
	0xd5, 0x8b, 0xe6, 0xe7, 0x84, 0xb7, 0x00, 0x2b, 0x13, 0xf9, 0xcd, 0x79, 0xcb, 0x81, 0x4a, 0xb9,
	0x30, 0x33, 0xa8, 0xe1, 0xfc, 0x89, 0x43, 0xe7, 0xed, 0x18, 0xad, 0x3b, 0xd2, 0x99, 0x14, 0xe6,
	0x2a, 0xed, 0x3f, 0x9b, 0x8e, 0x69, 0x79, 0xda, 0x4f, 0x9c, 0x5e, 0xdd, 0x8d, 0xbd, 0x89, 0xfc,
	0x62, 0x07, 0xa6, 0x4a, 0xab, 0x3a, 0x6f, 0x66, 0x06, 0x6f, 0xa0, 0x5a, 0x1d, 0xb8, 0x7b, 0x97,
	0x18, 0x24, 0xfe, 0x13, 0x6d, 0x29, 0xbf, 0xf2, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xf1,
	0xca, 0x03, 0x40, 0x01, 0x00, 0x00,
}