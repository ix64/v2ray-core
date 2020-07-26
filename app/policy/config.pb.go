package policy

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

type Second struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Second) Reset()         { *m = Second{} }
func (m *Second) String() string { return proto.CompactTextString(m) }
func (*Second) ProtoMessage()    {}
func (*Second) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{0}
}

func (m *Second) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Second.Unmarshal(m, b)
}
func (m *Second) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Second.Marshal(b, m, deterministic)
}
func (m *Second) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Second.Merge(m, src)
}
func (m *Second) XXX_Size() int {
	return xxx_messageInfo_Second.Size(m)
}
func (m *Second) XXX_DiscardUnknown() {
	xxx_messageInfo_Second.DiscardUnknown(m)
}

var xxx_messageInfo_Second proto.InternalMessageInfo

func (m *Second) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Policy struct {
	Timeout              *Policy_Timeout `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Stats                *Policy_Stats   `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Buffer               *Policy_Buffer  `protobuf:"bytes,3,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetTimeout() *Policy_Timeout {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Policy) GetStats() *Policy_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Policy) GetBuffer() *Policy_Buffer {
	if m != nil {
		return m.Buffer
	}
	return nil
}

// Timeout is a message for timeout settings in various stages, in seconds.
type Policy_Timeout struct {
	Handshake            *Second  `protobuf:"bytes,1,opt,name=handshake,proto3" json:"handshake,omitempty"`
	ConnectionIdle       *Second  `protobuf:"bytes,2,opt,name=connection_idle,json=connectionIdle,proto3" json:"connection_idle,omitempty"`
	UplinkOnly           *Second  `protobuf:"bytes,3,opt,name=uplink_only,json=uplinkOnly,proto3" json:"uplink_only,omitempty"`
	DownlinkOnly         *Second  `protobuf:"bytes,4,opt,name=downlink_only,json=downlinkOnly,proto3" json:"downlink_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy_Timeout) Reset()         { *m = Policy_Timeout{} }
func (m *Policy_Timeout) String() string { return proto.CompactTextString(m) }
func (*Policy_Timeout) ProtoMessage()    {}
func (*Policy_Timeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{1, 0}
}

func (m *Policy_Timeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy_Timeout.Unmarshal(m, b)
}
func (m *Policy_Timeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy_Timeout.Marshal(b, m, deterministic)
}
func (m *Policy_Timeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy_Timeout.Merge(m, src)
}
func (m *Policy_Timeout) XXX_Size() int {
	return xxx_messageInfo_Policy_Timeout.Size(m)
}
func (m *Policy_Timeout) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy_Timeout.DiscardUnknown(m)
}

var xxx_messageInfo_Policy_Timeout proto.InternalMessageInfo

func (m *Policy_Timeout) GetHandshake() *Second {
	if m != nil {
		return m.Handshake
	}
	return nil
}

func (m *Policy_Timeout) GetConnectionIdle() *Second {
	if m != nil {
		return m.ConnectionIdle
	}
	return nil
}

func (m *Policy_Timeout) GetUplinkOnly() *Second {
	if m != nil {
		return m.UplinkOnly
	}
	return nil
}

func (m *Policy_Timeout) GetDownlinkOnly() *Second {
	if m != nil {
		return m.DownlinkOnly
	}
	return nil
}

type Policy_Stats struct {
	UserUplink           bool     `protobuf:"varint,1,opt,name=user_uplink,json=userUplink,proto3" json:"user_uplink,omitempty"`
	UserDownlink         bool     `protobuf:"varint,2,opt,name=user_downlink,json=userDownlink,proto3" json:"user_downlink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy_Stats) Reset()         { *m = Policy_Stats{} }
func (m *Policy_Stats) String() string { return proto.CompactTextString(m) }
func (*Policy_Stats) ProtoMessage()    {}
func (*Policy_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{1, 1}
}

func (m *Policy_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy_Stats.Unmarshal(m, b)
}
func (m *Policy_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy_Stats.Marshal(b, m, deterministic)
}
func (m *Policy_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy_Stats.Merge(m, src)
}
func (m *Policy_Stats) XXX_Size() int {
	return xxx_messageInfo_Policy_Stats.Size(m)
}
func (m *Policy_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Policy_Stats proto.InternalMessageInfo

func (m *Policy_Stats) GetUserUplink() bool {
	if m != nil {
		return m.UserUplink
	}
	return false
}

func (m *Policy_Stats) GetUserDownlink() bool {
	if m != nil {
		return m.UserDownlink
	}
	return false
}

type Policy_Buffer struct {
	// Buffer size per connection, in bytes. -1 for unlimited buffer.
	Connection           int32    `protobuf:"varint,1,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy_Buffer) Reset()         { *m = Policy_Buffer{} }
func (m *Policy_Buffer) String() string { return proto.CompactTextString(m) }
func (*Policy_Buffer) ProtoMessage()    {}
func (*Policy_Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{1, 2}
}

func (m *Policy_Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy_Buffer.Unmarshal(m, b)
}
func (m *Policy_Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy_Buffer.Marshal(b, m, deterministic)
}
func (m *Policy_Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy_Buffer.Merge(m, src)
}
func (m *Policy_Buffer) XXX_Size() int {
	return xxx_messageInfo_Policy_Buffer.Size(m)
}
func (m *Policy_Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Policy_Buffer proto.InternalMessageInfo

func (m *Policy_Buffer) GetConnection() int32 {
	if m != nil {
		return m.Connection
	}
	return 0
}

type SystemPolicy struct {
	Stats                *SystemPolicy_Stats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SystemPolicy) Reset()         { *m = SystemPolicy{} }
func (m *SystemPolicy) String() string { return proto.CompactTextString(m) }
func (*SystemPolicy) ProtoMessage()    {}
func (*SystemPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{2}
}

func (m *SystemPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemPolicy.Unmarshal(m, b)
}
func (m *SystemPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemPolicy.Marshal(b, m, deterministic)
}
func (m *SystemPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemPolicy.Merge(m, src)
}
func (m *SystemPolicy) XXX_Size() int {
	return xxx_messageInfo_SystemPolicy.Size(m)
}
func (m *SystemPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_SystemPolicy proto.InternalMessageInfo

func (m *SystemPolicy) GetStats() *SystemPolicy_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type SystemPolicy_Stats struct {
	InboundUplink        bool     `protobuf:"varint,1,opt,name=inbound_uplink,json=inboundUplink,proto3" json:"inbound_uplink,omitempty"`
	InboundDownlink      bool     `protobuf:"varint,2,opt,name=inbound_downlink,json=inboundDownlink,proto3" json:"inbound_downlink,omitempty"`
	OutboundUplink       bool     `protobuf:"varint,3,opt,name=outbound_uplink,json=outboundUplink,proto3" json:"outbound_uplink,omitempty"`
	OutboundDownlink     bool     `protobuf:"varint,4,opt,name=outbound_downlink,json=outboundDownlink,proto3" json:"outbound_downlink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemPolicy_Stats) Reset()         { *m = SystemPolicy_Stats{} }
func (m *SystemPolicy_Stats) String() string { return proto.CompactTextString(m) }
func (*SystemPolicy_Stats) ProtoMessage()    {}
func (*SystemPolicy_Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{2, 0}
}

func (m *SystemPolicy_Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemPolicy_Stats.Unmarshal(m, b)
}
func (m *SystemPolicy_Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemPolicy_Stats.Marshal(b, m, deterministic)
}
func (m *SystemPolicy_Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemPolicy_Stats.Merge(m, src)
}
func (m *SystemPolicy_Stats) XXX_Size() int {
	return xxx_messageInfo_SystemPolicy_Stats.Size(m)
}
func (m *SystemPolicy_Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemPolicy_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_SystemPolicy_Stats proto.InternalMessageInfo

func (m *SystemPolicy_Stats) GetInboundUplink() bool {
	if m != nil {
		return m.InboundUplink
	}
	return false
}

func (m *SystemPolicy_Stats) GetInboundDownlink() bool {
	if m != nil {
		return m.InboundDownlink
	}
	return false
}

func (m *SystemPolicy_Stats) GetOutboundUplink() bool {
	if m != nil {
		return m.OutboundUplink
	}
	return false
}

func (m *SystemPolicy_Stats) GetOutboundDownlink() bool {
	if m != nil {
		return m.OutboundDownlink
	}
	return false
}

type Config struct {
	Level                map[uint32]*Policy `protobuf:"bytes,1,rep,name=level,proto3" json:"level,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	System               *SystemPolicy      `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_48f54a345c1316d1, []int{3}
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

func (m *Config) GetLevel() map[uint32]*Policy {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *Config) GetSystem() *SystemPolicy {
	if m != nil {
		return m.System
	}
	return nil
}

func init() {
	proto.RegisterType((*Second)(nil), "v2ray.core.app.policy.Second")
	proto.RegisterType((*Policy)(nil), "v2ray.core.app.policy.Policy")
	proto.RegisterType((*Policy_Timeout)(nil), "v2ray.core.app.policy.Policy.Timeout")
	proto.RegisterType((*Policy_Stats)(nil), "v2ray.core.app.policy.Policy.Stats")
	proto.RegisterType((*Policy_Buffer)(nil), "v2ray.core.app.policy.Policy.Buffer")
	proto.RegisterType((*SystemPolicy)(nil), "v2ray.core.app.policy.SystemPolicy")
	proto.RegisterType((*SystemPolicy_Stats)(nil), "v2ray.core.app.policy.SystemPolicy.Stats")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.policy.Config")
	proto.RegisterMapType((map[uint32]*Policy)(nil), "v2ray.core.app.policy.Config.LevelEntry")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/policy/config.proto", fileDescriptor_48f54a345c1316d1)
}

var fileDescriptor_48f54a345c1316d1 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x8a, 0xd3, 0x4c,
	0x18, 0xc7, 0x49, 0xda, 0x64, 0xf7, 0x7d, 0xfa, 0xf9, 0x0e, 0x2e, 0xd4, 0x80, 0xeb, 0xd2, 0x75,
	0xb5, 0x8b, 0x90, 0x42, 0xf7, 0x44, 0x5d, 0x5d, 0xb1, 0x7e, 0x80, 0xa0, 0xb8, 0x4c, 0xfd, 0x00,
	0x4f, 0x4a, 0x9a, 0x4c, 0xdd, 0xd0, 0x74, 0x26, 0x24, 0x93, 0x4a, 0x0e, 0xbd, 0x05, 0x2f, 0xc3,
	0x13, 0x6f, 0xc5, 0x2b, 0xf0, 0x5a, 0x24, 0xf3, 0xd1, 0xb4, 0xb2, 0x6d, 0x3d, 0x4b, 0x9f, 0xfe,
	0xfe, 0x3f, 0xfa, 0x7f, 0x3a, 0x19, 0xb8, 0xbb, 0x18, 0x24, 0x5e, 0xee, 0xfa, 0x6c, 0xde, 0xf7,
	0x59, 0x42, 0xfa, 0x5e, 0x1c, 0xf7, 0x63, 0x16, 0x85, 0x7e, 0xde, 0xf7, 0x19, 0x9d, 0x86, 0x5f,
	0xdc, 0x38, 0x61, 0x9c, 0xa1, 0x03, 0xcd, 0x25, 0xc4, 0xf5, 0xe2, 0xd8, 0x95, 0x4c, 0xf7, 0x10,
	0xec, 0x11, 0xf1, 0x19, 0x0d, 0xd0, 0x0d, 0xb0, 0x16, 0x5e, 0x94, 0x91, 0x8e, 0x71, 0x64, 0xf4,
	0x1a, 0x58, 0x7e, 0xe8, 0xfe, 0xaa, 0x82, 0x7d, 0x29, 0x50, 0xf4, 0x14, 0xf6, 0x78, 0x38, 0x27,
	0x2c, 0xe3, 0x02, 0xa9, 0x0d, 0x4e, 0xdc, 0x6b, 0x9d, 0xae, 0xe4, 0xdd, 0xf7, 0x12, 0xc6, 0x3a,
	0x85, 0x1e, 0x82, 0x95, 0x72, 0x8f, 0xa7, 0x1d, 0x53, 0xc4, 0x8f, 0xb7, 0xc7, 0x47, 0x05, 0x8a,
	0x65, 0x02, 0x3d, 0x06, 0x7b, 0x92, 0x4d, 0xa7, 0x24, 0xe9, 0x54, 0x44, 0xf6, 0xce, 0xf6, 0xec,
	0x50, 0xb0, 0x58, 0x65, 0x9c, 0xef, 0x26, 0xec, 0xa9, 0x5f, 0x83, 0xce, 0xe1, 0xbf, 0x2b, 0x8f,
	0x06, 0xe9, 0x95, 0x37, 0x23, 0xaa, 0xc7, 0xad, 0x0d, 0x32, 0xb9, 0x18, 0x5c, 0xf2, 0xe8, 0x15,
	0xb4, 0x7c, 0x46, 0x29, 0xf1, 0x79, 0xc8, 0xe8, 0x38, 0x0c, 0x22, 0xa2, 0xba, 0xec, 0x50, 0x34,
	0xcb, 0xd4, 0xeb, 0x20, 0x22, 0xe8, 0x02, 0x6a, 0x59, 0x1c, 0x85, 0x74, 0x36, 0x66, 0x34, 0xca,
	0x55, 0xa7, 0x1d, 0x0e, 0x90, 0x89, 0x77, 0x34, 0xca, 0xd1, 0x10, 0x1a, 0x01, 0xfb, 0x4a, 0x4b,
	0x43, 0xf5, 0x5f, 0x0c, 0x75, 0x9d, 0x29, 0x1c, 0xce, 0x5b, 0xb0, 0xc4, 0x8a, 0xd1, 0x6d, 0xa8,
	0x65, 0x29, 0x49, 0xc6, 0xd2, 0x2f, 0x76, 0xb2, 0x8f, 0xa1, 0x18, 0x7d, 0x10, 0x13, 0x74, 0x0c,
	0x0d, 0x01, 0xe8, 0xb8, 0xe8, 0xbc, 0x8f, 0xeb, 0xc5, 0xf0, 0x85, 0x9a, 0x39, 0x3d, 0xb0, 0xe5,
	0xd6, 0xd1, 0x21, 0x40, 0x59, 0x57, 0xe8, 0x2c, 0xbc, 0x32, 0xe9, 0x7e, 0x33, 0xa1, 0x3e, 0xca,
	0x53, 0x4e, 0xe6, 0xcb, 0x83, 0xa5, 0xce, 0x85, 0xfc, 0x3b, 0x4e, 0x37, 0xb5, 0x58, 0xc9, 0xac,
	0x9d, 0x0e, 0xe7, 0xa7, 0xa1, 0xbb, 0x9c, 0x40, 0x33, 0xa4, 0x13, 0x96, 0xd1, 0x60, 0xbd, 0x4e,
	0x43, 0x4d, 0x55, 0xa3, 0x53, 0x68, 0x6b, 0xec, 0xaf, 0x52, 0x2d, 0x35, 0xd7, 0xbd, 0xd0, 0x3d,
	0x68, 0xb1, 0x8c, 0xaf, 0x29, 0x2b, 0x82, 0x6c, 0xea, 0xb1, 0x72, 0xde, 0x87, 0xff, 0x97, 0xe0,
	0x52, 0x5a, 0x15, 0x68, 0x5b, 0x7f, 0xa1, 0xad, 0xdd, 0xdf, 0x06, 0xd8, 0xcf, 0xc5, 0xeb, 0x89,
	0x2e, 0xc0, 0x8a, 0xc8, 0x82, 0x44, 0x1d, 0xe3, 0xa8, 0xd2, 0xab, 0x0d, 0x7a, 0x1b, 0xda, 0x4b,
	0xda, 0x7d, 0x53, 0xa0, 0x2f, 0x29, 0x4f, 0x72, 0x2c, 0x63, 0xe8, 0x1c, 0xec, 0x54, 0x6c, 0x66,
	0xc7, 0x6b, 0xb5, 0xba, 0x3e, 0xac, 0x22, 0xce, 0x27, 0x80, 0xd2, 0x88, 0xda, 0x50, 0x99, 0x91,
	0x5c, 0x5d, 0x00, 0xc5, 0x23, 0x3a, 0xd3, 0x97, 0xc2, 0xf6, 0x63, 0xae, 0xac, 0x92, 0x7d, 0x64,
	0x3e, 0x30, 0x86, 0x4f, 0xe0, 0xa6, 0xcf, 0xe6, 0xd7, 0xe3, 0x97, 0xc6, 0x67, 0x5b, 0x3e, 0xfd,
	0x30, 0x0f, 0x3e, 0x0e, 0xb0, 0x57, 0xb4, 0x4b, 0x88, 0xfb, 0x2c, 0x8e, 0x95, 0x69, 0x62, 0x8b,
	0x4b, 0xeb, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x33, 0x33, 0xa6, 0xde, 0x04, 0x00,
	0x00,
}
