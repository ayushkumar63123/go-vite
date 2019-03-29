// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus_point.proto

package vitepb

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

type PointContent struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	FNum                 uint32   `protobuf:"varint,2,opt,name=fNum,proto3" json:"fNum,omitempty"`
	ENum                 uint32   `protobuf:"varint,3,opt,name=eNum,proto3" json:"eNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointContent) Reset()         { *m = PointContent{} }
func (m *PointContent) String() string { return proto.CompactTextString(m) }
func (*PointContent) ProtoMessage()    {}
func (*PointContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_528a02e43f476ca6, []int{0}
}

func (m *PointContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointContent.Unmarshal(m, b)
}
func (m *PointContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointContent.Marshal(b, m, deterministic)
}
func (m *PointContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointContent.Merge(m, src)
}
func (m *PointContent) XXX_Size() int {
	return xxx_messageInfo_PointContent.Size(m)
}
func (m *PointContent) XXX_DiscardUnknown() {
	xxx_messageInfo_PointContent.DiscardUnknown(m)
}

var xxx_messageInfo_PointContent proto.InternalMessageInfo

func (m *PointContent) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PointContent) GetFNum() uint32 {
	if m != nil {
		return m.FNum
	}
	return 0
}

func (m *PointContent) GetENum() uint32 {
	if m != nil {
		return m.ENum
	}
	return 0
}

type ConsensusPoint struct {
	PrevHash             []byte          `protobuf:"bytes,1,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Hash                 []byte          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Contents             []*PointContent `protobuf:"bytes,3,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConsensusPoint) Reset()         { *m = ConsensusPoint{} }
func (m *ConsensusPoint) String() string { return proto.CompactTextString(m) }
func (*ConsensusPoint) ProtoMessage()    {}
func (*ConsensusPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_528a02e43f476ca6, []int{1}
}

func (m *ConsensusPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusPoint.Unmarshal(m, b)
}
func (m *ConsensusPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusPoint.Marshal(b, m, deterministic)
}
func (m *ConsensusPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusPoint.Merge(m, src)
}
func (m *ConsensusPoint) XXX_Size() int {
	return xxx_messageInfo_ConsensusPoint.Size(m)
}
func (m *ConsensusPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusPoint.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusPoint proto.InternalMessageInfo

func (m *ConsensusPoint) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *ConsensusPoint) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ConsensusPoint) GetContents() []*PointContent {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*PointContent)(nil), "vitepb.PointContent")
	proto.RegisterType((*ConsensusPoint)(nil), "vitepb.ConsensusPoint")
}

func init() { proto.RegisterFile("consensus_point.proto", fileDescriptor_528a02e43f476ca6) }

var fileDescriptor_528a02e43f476ca6 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2b,
	0x4e, 0xcd, 0x2b, 0x2e, 0x2d, 0x8e, 0x2f, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2b, 0xcb, 0x2c, 0x49, 0x2d, 0x48, 0x52, 0x0a, 0xe0, 0xe2, 0x09, 0x00, 0x09,
	0x3b, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x08, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5,
	0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x42, 0x42, 0x5c, 0x2c, 0x69,
	0x7e, 0xa5, 0xb9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0x36, 0x48, 0x2c, 0x15, 0x24,
	0xc6, 0x0c, 0x11, 0x03, 0xb1, 0x95, 0x8a, 0xb8, 0xf8, 0x9c, 0x61, 0x56, 0x82, 0x8d, 0x16, 0x92,
	0xe2, 0xe2, 0x28, 0x28, 0x4a, 0x2d, 0xf3, 0x48, 0x2c, 0xce, 0x80, 0x1a, 0x0a, 0xe7, 0x83, 0x4c,
	0xc8, 0x00, 0x89, 0x33, 0x81, 0xc5, 0xc1, 0x6c, 0x21, 0x03, 0x2e, 0x8e, 0x64, 0x88, 0x73, 0x8a,
	0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x20, 0xce, 0xd5, 0x43, 0x76, 0x6b, 0x10,
	0x5c, 0x55, 0x12, 0x1b, 0xd8, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6d, 0x3f,
	0x39, 0xed, 0x00, 0x00, 0x00,
}