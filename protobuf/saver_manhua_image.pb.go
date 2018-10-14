// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_manhua_image.proto

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

type SaverManhuaImage struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Site                 string   `protobuf:"bytes,2,opt,name=Site" json:"Site,omitempty"`
	Number               int32    `protobuf:"varint,3,opt,name=Number" json:"Number,omitempty"`
	Chapter              string   `protobuf:"bytes,4,opt,name=Chapter" json:"Chapter,omitempty"`
	Sequence             string   `protobuf:"bytes,5,opt,name=Sequence" json:"Sequence,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=Url" json:"Url,omitempty"`
	Uuid                 string   `protobuf:"bytes,7,opt,name=Uuid" json:"Uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `gorm:"-" json:"-"`
	XXX_unrecognized     []byte   `gorm:"-" json:"-"`
	XXX_sizecache        int32    `gorm:"-" json:"-"`
}

func (m *SaverManhuaImage) Reset()         { *m = SaverManhuaImage{} }
func (m *SaverManhuaImage) String() string { return proto.CompactTextString(m) }
func (*SaverManhuaImage) ProtoMessage()    {}
func (*SaverManhuaImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_manhua_image_4e4ac23122b2427a, []int{0}
}
func (m *SaverManhuaImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverManhuaImage.Unmarshal(m, b)
}
func (m *SaverManhuaImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverManhuaImage.Marshal(b, m, deterministic)
}
func (dst *SaverManhuaImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverManhuaImage.Merge(dst, src)
}
func (m *SaverManhuaImage) XXX_Size() int {
	return xxx_messageInfo_SaverManhuaImage.Size(m)
}
func (m *SaverManhuaImage) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverManhuaImage.DiscardUnknown(m)
}

var xxx_messageInfo_SaverManhuaImage proto.InternalMessageInfo

func (m *SaverManhuaImage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverManhuaImage) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *SaverManhuaImage) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SaverManhuaImage) GetChapter() string {
	if m != nil {
		return m.Chapter
	}
	return ""
}

func (m *SaverManhuaImage) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

func (m *SaverManhuaImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SaverManhuaImage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*SaverManhuaImage)(nil), "protobuf.Saver_manhua_image")
}

func init() {
	proto.RegisterFile("protobuf/saver_manhua_image.proto", fileDescriptor_saver_manhua_image_4e4ac23122b2427a)
}

var fileDescriptor_saver_manhua_image_4e4ac23122b2427a = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2c, 0x4b, 0x2d, 0x8a, 0xcf, 0x4d, 0xcc, 0xcb, 0x28,
	0x4d, 0x8c, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x03, 0xcb, 0x09, 0x71, 0xc0, 0x94, 0x28, 0xad,
	0x61, 0xe4, 0x12, 0x0a, 0xc6, 0x50, 0x26, 0xc4, 0xc7, 0xc5, 0xe4, 0x99, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x12, 0xc4, 0xe4, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x12, 0x9c, 0x59, 0x92, 0x2a, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x71, 0xb1, 0xf9, 0x95, 0xe6, 0x26, 0xa5,
	0x16, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xce, 0x19,
	0x89, 0x05, 0x25, 0xa9, 0x45, 0x12, 0x2c, 0x60, 0xe5, 0x30, 0xae, 0x90, 0x14, 0x17, 0x47, 0x70,
	0x6a, 0x61, 0x69, 0x6a, 0x5e, 0x72, 0xaa, 0x04, 0x2b, 0x58, 0x0a, 0xce, 0x17, 0x12, 0xe0, 0x62,
	0x0e, 0x2d, 0xca, 0x91, 0x60, 0x03, 0x0b, 0x83, 0x98, 0x20, 0x3b, 0x43, 0x4b, 0x33, 0x53, 0x24,
	0xd8, 0x21, 0x76, 0x82, 0xd8, 0x49, 0x6c, 0x60, 0x87, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x27, 0xa9, 0xa2, 0xbe, 0xe4, 0x00, 0x00, 0x00,
}