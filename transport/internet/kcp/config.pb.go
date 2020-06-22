package kcp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	serial "v2ray.com/core/common/serial"
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

// Maximum Transmission Unit, in bytes.
type MTU struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MTU) Reset()         { *m = MTU{} }
func (m *MTU) String() string { return proto.CompactTextString(m) }
func (*MTU) ProtoMessage()    {}
func (*MTU) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{0}
}

func (m *MTU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MTU.Unmarshal(m, b)
}
func (m *MTU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MTU.Marshal(b, m, deterministic)
}
func (m *MTU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTU.Merge(m, src)
}
func (m *MTU) XXX_Size() int {
	return xxx_messageInfo_MTU.Size(m)
}
func (m *MTU) XXX_DiscardUnknown() {
	xxx_messageInfo_MTU.DiscardUnknown(m)
}

var xxx_messageInfo_MTU proto.InternalMessageInfo

func (m *MTU) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Transmission Time Interview, in milli-sec.
type TTI struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TTI) Reset()         { *m = TTI{} }
func (m *TTI) String() string { return proto.CompactTextString(m) }
func (*TTI) ProtoMessage()    {}
func (*TTI) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{1}
}

func (m *TTI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TTI.Unmarshal(m, b)
}
func (m *TTI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TTI.Marshal(b, m, deterministic)
}
func (m *TTI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TTI.Merge(m, src)
}
func (m *TTI) XXX_Size() int {
	return xxx_messageInfo_TTI.Size(m)
}
func (m *TTI) XXX_DiscardUnknown() {
	xxx_messageInfo_TTI.DiscardUnknown(m)
}

var xxx_messageInfo_TTI proto.InternalMessageInfo

func (m *TTI) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Uplink capacity, in MB.
type UplinkCapacity struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkCapacity) Reset()         { *m = UplinkCapacity{} }
func (m *UplinkCapacity) String() string { return proto.CompactTextString(m) }
func (*UplinkCapacity) ProtoMessage()    {}
func (*UplinkCapacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{2}
}

func (m *UplinkCapacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkCapacity.Unmarshal(m, b)
}
func (m *UplinkCapacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkCapacity.Marshal(b, m, deterministic)
}
func (m *UplinkCapacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkCapacity.Merge(m, src)
}
func (m *UplinkCapacity) XXX_Size() int {
	return xxx_messageInfo_UplinkCapacity.Size(m)
}
func (m *UplinkCapacity) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkCapacity.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkCapacity proto.InternalMessageInfo

func (m *UplinkCapacity) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Downlink capacity, in MB.
type DownlinkCapacity struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkCapacity) Reset()         { *m = DownlinkCapacity{} }
func (m *DownlinkCapacity) String() string { return proto.CompactTextString(m) }
func (*DownlinkCapacity) ProtoMessage()    {}
func (*DownlinkCapacity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{3}
}

func (m *DownlinkCapacity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkCapacity.Unmarshal(m, b)
}
func (m *DownlinkCapacity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkCapacity.Marshal(b, m, deterministic)
}
func (m *DownlinkCapacity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkCapacity.Merge(m, src)
}
func (m *DownlinkCapacity) XXX_Size() int {
	return xxx_messageInfo_DownlinkCapacity.Size(m)
}
func (m *DownlinkCapacity) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkCapacity.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkCapacity proto.InternalMessageInfo

func (m *DownlinkCapacity) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type WriteBuffer struct {
	// Buffer size in bytes.
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBuffer) Reset()         { *m = WriteBuffer{} }
func (m *WriteBuffer) String() string { return proto.CompactTextString(m) }
func (*WriteBuffer) ProtoMessage()    {}
func (*WriteBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{4}
}

func (m *WriteBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteBuffer.Unmarshal(m, b)
}
func (m *WriteBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteBuffer.Marshal(b, m, deterministic)
}
func (m *WriteBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBuffer.Merge(m, src)
}
func (m *WriteBuffer) XXX_Size() int {
	return xxx_messageInfo_WriteBuffer.Size(m)
}
func (m *WriteBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBuffer proto.InternalMessageInfo

func (m *WriteBuffer) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReadBuffer struct {
	// Buffer size in bytes.
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBuffer) Reset()         { *m = ReadBuffer{} }
func (m *ReadBuffer) String() string { return proto.CompactTextString(m) }
func (*ReadBuffer) ProtoMessage()    {}
func (*ReadBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{5}
}

func (m *ReadBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBuffer.Unmarshal(m, b)
}
func (m *ReadBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBuffer.Marshal(b, m, deterministic)
}
func (m *ReadBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBuffer.Merge(m, src)
}
func (m *ReadBuffer) XXX_Size() int {
	return xxx_messageInfo_ReadBuffer.Size(m)
}
func (m *ReadBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBuffer proto.InternalMessageInfo

func (m *ReadBuffer) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ConnectionReuse struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionReuse) Reset()         { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()    {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{6}
}

func (m *ConnectionReuse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionReuse.Unmarshal(m, b)
}
func (m *ConnectionReuse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionReuse.Marshal(b, m, deterministic)
}
func (m *ConnectionReuse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionReuse.Merge(m, src)
}
func (m *ConnectionReuse) XXX_Size() int {
	return xxx_messageInfo_ConnectionReuse.Size(m)
}
func (m *ConnectionReuse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionReuse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionReuse proto.InternalMessageInfo

func (m *ConnectionReuse) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

// Maximum Transmission Unit, in bytes.
type EncryptionSeed struct {
	Seed                 string   `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptionSeed) Reset()         { *m = EncryptionSeed{} }
func (m *EncryptionSeed) String() string { return proto.CompactTextString(m) }
func (*EncryptionSeed) ProtoMessage()    {}
func (*EncryptionSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{7}
}

func (m *EncryptionSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionSeed.Unmarshal(m, b)
}
func (m *EncryptionSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionSeed.Marshal(b, m, deterministic)
}
func (m *EncryptionSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionSeed.Merge(m, src)
}
func (m *EncryptionSeed) XXX_Size() int {
	return xxx_messageInfo_EncryptionSeed.Size(m)
}
func (m *EncryptionSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionSeed.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionSeed proto.InternalMessageInfo

func (m *EncryptionSeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type Config struct {
	Mtu                  *MTU                 `protobuf:"bytes,1,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Tti                  *TTI                 `protobuf:"bytes,2,opt,name=tti,proto3" json:"tti,omitempty"`
	UplinkCapacity       *UplinkCapacity      `protobuf:"bytes,3,opt,name=uplink_capacity,json=uplinkCapacity,proto3" json:"uplink_capacity,omitempty"`
	DownlinkCapacity     *DownlinkCapacity    `protobuf:"bytes,4,opt,name=downlink_capacity,json=downlinkCapacity,proto3" json:"downlink_capacity,omitempty"`
	Congestion           bool                 `protobuf:"varint,5,opt,name=congestion,proto3" json:"congestion,omitempty"`
	WriteBuffer          *WriteBuffer         `protobuf:"bytes,6,opt,name=write_buffer,json=writeBuffer,proto3" json:"write_buffer,omitempty"`
	ReadBuffer           *ReadBuffer          `protobuf:"bytes,7,opt,name=read_buffer,json=readBuffer,proto3" json:"read_buffer,omitempty"`
	HeaderConfig         *serial.TypedMessage `protobuf:"bytes,8,opt,name=header_config,json=headerConfig,proto3" json:"header_config,omitempty"`
	Seed                 *EncryptionSeed      `protobuf:"bytes,10,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3746d5d763e81577, []int{8}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetMtu() *MTU {
	if m != nil {
		return m.Mtu
	}
	return nil
}

func (m *Config) GetTti() *TTI {
	if m != nil {
		return m.Tti
	}
	return nil
}

func (m *Config) GetUplinkCapacity() *UplinkCapacity {
	if m != nil {
		return m.UplinkCapacity
	}
	return nil
}

func (m *Config) GetDownlinkCapacity() *DownlinkCapacity {
	if m != nil {
		return m.DownlinkCapacity
	}
	return nil
}

func (m *Config) GetCongestion() bool {
	if m != nil {
		return m.Congestion
	}
	return false
}

func (m *Config) GetWriteBuffer() *WriteBuffer {
	if m != nil {
		return m.WriteBuffer
	}
	return nil
}

func (m *Config) GetReadBuffer() *ReadBuffer {
	if m != nil {
		return m.ReadBuffer
	}
	return nil
}

func (m *Config) GetHeaderConfig() *serial.TypedMessage {
	if m != nil {
		return m.HeaderConfig
	}
	return nil
}

func (m *Config) GetSeed() *EncryptionSeed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func init() {
	proto.RegisterType((*MTU)(nil), "v2ray.core.transport.internet.kcp.MTU")
	proto.RegisterType((*TTI)(nil), "v2ray.core.transport.internet.kcp.TTI")
	proto.RegisterType((*UplinkCapacity)(nil), "v2ray.core.transport.internet.kcp.UplinkCapacity")
	proto.RegisterType((*DownlinkCapacity)(nil), "v2ray.core.transport.internet.kcp.DownlinkCapacity")
	proto.RegisterType((*WriteBuffer)(nil), "v2ray.core.transport.internet.kcp.WriteBuffer")
	proto.RegisterType((*ReadBuffer)(nil), "v2ray.core.transport.internet.kcp.ReadBuffer")
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.kcp.ConnectionReuse")
	proto.RegisterType((*EncryptionSeed)(nil), "v2ray.core.transport.internet.kcp.EncryptionSeed")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.kcp.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/kcp/config.proto", fileDescriptor_3746d5d763e81577)
}

var fileDescriptor_3746d5d763e81577 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0xf5, 0x0f, 0xdd, 0xe9, 0xd6, 0x95, 0x08, 0xa1, 0x08, 0x24, 0xb4, 0x56, 0x30,
	0x8d, 0x0b, 0x1c, 0xe8, 0x6e, 0xb8, 0x5e, 0xd9, 0x45, 0x99, 0x8a, 0xc0, 0xa4, 0x20, 0xed, 0xa6,
	0xb8, 0xce, 0x69, 0x89, 0xda, 0xd8, 0x96, 0xe3, 0xac, 0x2a, 0x2f, 0xc2, 0x3b, 0xf0, 0x94, 0x28,
	0x4e, 0xb2, 0xfe, 0x41, 0x63, 0xb9, 0xb3, 0x73, 0xbe, 0xf3, 0xb3, 0x74, 0xbe, 0xef, 0x04, 0xfa,
	0xb7, 0x7d, 0xcd, 0xd6, 0x84, 0xcb, 0xc8, 0xe3, 0x52, 0xa3, 0x67, 0x34, 0x13, 0xb1, 0x92, 0xda,
	0x78, 0xa1, 0x30, 0xa8, 0x05, 0x1a, 0x6f, 0xc1, 0x95, 0xc7, 0xa5, 0x98, 0x85, 0x73, 0xa2, 0xb4,
	0x34, 0xd2, 0xe9, 0x16, 0x3d, 0x1a, 0xc9, 0x9d, 0x9e, 0x14, 0x7a, 0xb2, 0xe0, 0xea, 0xd9, 0xdb,
	0x3d, 0x2c, 0x97, 0x51, 0x24, 0x85, 0x17, 0xa3, 0x0e, 0xd9, 0xd2, 0x33, 0x6b, 0x85, 0xc1, 0x24,
	0xc2, 0x38, 0x66, 0x73, 0xcc, 0xa0, 0xbd, 0xe7, 0x50, 0x1d, 0xf9, 0x63, 0xe7, 0x09, 0xd4, 0x6f,
	0xd9, 0x32, 0x41, 0xb7, 0x72, 0x5a, 0x39, 0x3f, 0xa6, 0xd9, 0x25, 0x2d, 0xfa, 0xfe, 0xf0, 0x9e,
	0xe2, 0x19, 0xb4, 0xc7, 0x6a, 0x19, 0x8a, 0xc5, 0x80, 0x29, 0xc6, 0x43, 0xb3, 0xbe, 0x47, 0x77,
	0x0e, 0x9d, 0x0f, 0x72, 0x25, 0x4a, 0x28, 0xbb, 0xd0, 0xfa, 0xae, 0x43, 0x83, 0x97, 0xc9, 0x6c,
	0x86, 0xda, 0x71, 0xa0, 0x16, 0x87, 0xbf, 0x0a, 0x8d, 0x3d, 0xf7, 0x4e, 0x01, 0x28, 0xb2, 0xe0,
	0x3f, 0x8a, 0xd7, 0x70, 0x32, 0x90, 0x42, 0x20, 0x37, 0xa1, 0x14, 0x14, 0x93, 0x18, 0x9d, 0xa7,
	0xd0, 0x40, 0xc1, 0xa6, 0xcb, 0x4c, 0xd8, 0xa4, 0xf9, 0xad, 0xf7, 0x12, 0xda, 0x57, 0x82, 0xeb,
	0xb5, 0x4a, 0xa5, 0x5f, 0x11, 0x03, 0x0b, 0x44, 0x0c, 0xac, 0xee, 0x90, 0xda, 0x73, 0xef, 0x77,
	0x1d, 0x1a, 0x03, 0xeb, 0x83, 0xf3, 0x1e, 0xaa, 0x91, 0x49, 0x6c, 0xb5, 0xd5, 0x3f, 0x23, 0x0f,
	0xfa, 0x41, 0x46, 0xfe, 0x98, 0xa6, 0x2d, 0x69, 0xa7, 0x31, 0xa1, 0x7b, 0x50, 0xba, 0xd3, 0xf7,
	0x87, 0x34, 0x6d, 0x71, 0x6e, 0xe0, 0x24, 0xb1, 0x63, 0x9e, 0xf0, 0x7c, 0x7a, 0x6e, 0xd5, 0x52,
	0xde, 0x95, 0xa0, 0xec, 0x1a, 0x44, 0xdb, 0xc9, 0xae, 0x61, 0x3f, 0xe0, 0x71, 0x90, 0x5b, 0xb3,
	0xa1, 0xd7, 0x2c, 0xfd, 0xa2, 0x04, 0x7d, 0xdf, 0x56, 0xda, 0x09, 0xf6, 0x8d, 0x7e, 0x01, 0xc0,
	0xa5, 0x98, 0x63, 0x9c, 0x8e, 0xd8, 0xad, 0xdb, 0xf1, 0x6f, 0x7d, 0x71, 0xbe, 0xc0, 0xd1, 0x2a,
	0xb5, 0x7c, 0x32, 0xb5, 0x8e, 0xba, 0x0d, 0xfb, 0x38, 0x29, 0xf1, 0xf8, 0x56, 0x52, 0x68, 0x6b,
	0xb5, 0x15, 0x9b, 0x4f, 0xd0, 0xd2, 0xc8, 0x82, 0x82, 0xf8, 0xc8, 0x12, 0xdf, 0x94, 0x20, 0x6e,
	0x82, 0x45, 0x41, 0x6f, 0x42, 0x76, 0x0d, 0xc7, 0x3f, 0x91, 0x05, 0xa8, 0x27, 0xd9, 0x36, 0xba,
	0xcd, 0x7f, 0x4d, 0xcc, 0xf6, 0x8c, 0x64, 0x7b, 0x46, 0xfc, 0x74, 0xcf, 0x46, 0xd9, 0x9a, 0xd1,
	0xa3, 0xac, 0x39, 0x4f, 0xd0, 0x55, 0x1e, 0x30, 0x28, 0x6d, 0xe1, 0x6e, 0x42, 0xb3, 0x4c, 0x7e,
	0xac, 0x35, 0x0f, 0x3b, 0x70, 0x49, 0xe1, 0x15, 0x97, 0xd1, 0xc3, 0x8c, 0xcf, 0x95, 0x9b, 0xea,
	0x82, 0xab, 0x3f, 0x07, 0xdd, 0x6f, 0x7d, 0xca, 0xd6, 0x64, 0x90, 0x4a, 0xfd, 0x3b, 0xe9, 0xb0,
	0x90, 0x5e, 0x73, 0x35, 0x6d, 0xd8, 0xdf, 0xc2, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x9b, 0x97, 0xb7, 0xa1, 0x04, 0x00, 0x00,
}
