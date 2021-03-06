// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/saver_article_participle.proto

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

type SaverArticleParticiple struct {
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: gorm:"index:word"
	Word string `protobuf:"bytes,2,opt,name=Word,proto3" json:"Word,omitempty" gorm:"index:word"`
	// @inject_tag: gorm:"index:aId"
	Aid uint64 `protobuf:"varint,3,opt,name=aid,proto3" json:"aid,omitempty" gorm:"index:aId"`
	// @inject_tag: gorm:"index:checked"
	Checked              bool     `protobuf:"varint,4,opt,name=checked,proto3" json:"checked,omitempty" gorm:"index:checked"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" gorm:"-"`
}

func (m *SaverArticleParticiple) Reset()         { *m = SaverArticleParticiple{} }
func (m *SaverArticleParticiple) String() string { return proto.CompactTextString(m) }
func (*SaverArticleParticiple) ProtoMessage()    {}
func (*SaverArticleParticiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_saver_article_participle_76510ed5cc3f51e9, []int{0}
}
func (m *SaverArticleParticiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverArticleParticiple.Unmarshal(m, b)
}
func (m *SaverArticleParticiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverArticleParticiple.Marshal(b, m, deterministic)
}
func (dst *SaverArticleParticiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverArticleParticiple.Merge(dst, src)
}
func (m *SaverArticleParticiple) XXX_Size() int {
	return xxx_messageInfo_SaverArticleParticiple.Size(m)
}
func (m *SaverArticleParticiple) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverArticleParticiple.DiscardUnknown(m)
}

var xxx_messageInfo_SaverArticleParticiple proto.InternalMessageInfo

func (m *SaverArticleParticiple) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaverArticleParticiple) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *SaverArticleParticiple) GetAid() uint64 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *SaverArticleParticiple) GetChecked() bool {
	if m != nil {
		return m.Checked
	}
	return false
}

func init() {
	proto.RegisterType((*SaverArticleParticiple)(nil), "protobuf.Saver_article_participle")
}

func init() {
	proto.RegisterFile("protobuf/saver_article_participle.proto", fileDescriptor_saver_article_participle_76510ed5cc3f51e9)
}

var fileDescriptor_saver_article_participle_76510ed5cc3f51e9 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2c, 0x4b, 0x2d, 0x8a, 0x4f, 0x2c, 0x2a, 0xc9, 0x4c,
	0xce, 0x49, 0x8d, 0x2f, 0x00, 0x33, 0x32, 0x0b, 0x72, 0x52, 0xf5, 0xc0, 0x2a, 0x84, 0x38, 0x60,
	0x0a, 0x95, 0xb2, 0xb8, 0x24, 0x82, 0x71, 0xa8, 0x15, 0xe2, 0xe3, 0x62, 0xf2, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xf2, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x09, 0xcf, 0x2f, 0x4a,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0x33, 0x53,
	0x24, 0x98, 0xc1, 0x8a, 0x40, 0x4c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0x8c, 0xd4, 0xe4, 0xec, 0xd4,
	0x14, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0xab, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x59, 0xc0, 0x4f, 0xf3, 0xa7, 0x00, 0x00, 0x00,
}
