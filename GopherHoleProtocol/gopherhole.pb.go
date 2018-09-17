// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gopherhole.proto

package GopherHoleProtocol

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Data struct {
	IsCompressed         bool     `protobuf:"varint,1,opt,name=isCompressed,proto3" json:"isCompressed,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	CompressionType      string   `protobuf:"bytes,3,opt,name=compressionType,proto3" json:"compressionType,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_74ff1ba11f7cdfa4, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetIsCompressed() bool {
	if m != nil {
		return m.IsCompressed
	}
	return false
}

func (m *Data) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Data) GetCompressionType() string {
	if m != nil {
		return m.CompressionType
	}
	return ""
}

func (m *Data) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "GopherHoleProtocol.Data")
}

func init() { proto.RegisterFile("gopherhole.proto", fileDescriptor_74ff1ba11f7cdfa4) }

var fileDescriptor_74ff1ba11f7cdfa4 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcf, 0x2f, 0xc8,
	0x48, 0x2d, 0xca, 0xc8, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x72, 0x07,
	0x8b, 0x78, 0xe4, 0xe7, 0xa4, 0x06, 0x80, 0x04, 0x92, 0xf3, 0x73, 0x94, 0xba, 0x18, 0xb9, 0x58,
	0x5c, 0x12, 0x4b, 0x12, 0x85, 0x94, 0xb8, 0x78, 0x32, 0x8b, 0x9d, 0xf3, 0x73, 0x0b, 0x8a, 0x52,
	0x8b, 0x8b, 0x53, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x50, 0xc4, 0x84, 0xa4, 0xb8,
	0x38, 0xd2, 0x32, 0x73, 0x52, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0xe0, 0x7c, 0x21, 0x0d, 0x2e, 0xfe, 0x64, 0xa8, 0xca, 0xcc, 0xfc, 0xbc, 0x90, 0xca, 0x82, 0x54,
	0x09, 0x66, 0xb0, 0x12, 0x74, 0x61, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4,
	0x14, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x37, 0x89, 0x0d, 0xec, 0x4e, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x4d, 0xbd, 0x3a, 0xbb, 0x00, 0x00, 0x00,
}