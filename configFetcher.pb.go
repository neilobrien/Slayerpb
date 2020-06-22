// Code generated by protoc-gen-go. DO NOT EDIT.
// source: configFetcher.proto

package configfetcherpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Probe struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProbeID              int32    `protobuf:"varint,2,opt,name=probeID,proto3" json:"probeID,omitempty"`
	IpVer                int32    `protobuf:"varint,3,opt,name=ipVer,proto3" json:"ipVer,omitempty"`
	ProbeType            string   `protobuf:"bytes,4,opt,name=probeType,proto3" json:"probeType,omitempty"`
	TargetAddress        string   `protobuf:"bytes,5,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	NumPackets           int32    `protobuf:"varint,6,opt,name=numPackets,proto3" json:"numPackets,omitempty"`
	Interval             int32    `protobuf:"varint,7,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}
func (*Probe) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{0}
}

func (m *Probe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe.Unmarshal(m, b)
}
func (m *Probe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe.Marshal(b, m, deterministic)
}
func (m *Probe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe.Merge(m, src)
}
func (m *Probe) XXX_Size() int {
	return xxx_messageInfo_Probe.Size(m)
}
func (m *Probe) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe.DiscardUnknown(m)
}

var xxx_messageInfo_Probe proto.InternalMessageInfo

func (m *Probe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Probe) GetProbeID() int32 {
	if m != nil {
		return m.ProbeID
	}
	return 0
}

func (m *Probe) GetIpVer() int32 {
	if m != nil {
		return m.IpVer
	}
	return 0
}

func (m *Probe) GetProbeType() string {
	if m != nil {
		return m.ProbeType
	}
	return ""
}

func (m *Probe) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *Probe) GetNumPackets() int32 {
	if m != nil {
		return m.NumPackets
	}
	return 0
}

func (m *Probe) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// Config Requests and Replies
type ConfigRequest struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{1}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

type ConfigResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Probes               []*Probe `protobuf:"bytes,2,rep,name=probes,proto3" json:"probes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{2}
}

func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ConfigResponse) GetProbes() []*Probe {
	if m != nil {
		return m.Probes
	}
	return nil
}

type ConfigVersionResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigVersionResponse) Reset()         { *m = ConfigVersionResponse{} }
func (m *ConfigVersionResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigVersionResponse) ProtoMessage()    {}
func (*ConfigVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{3}
}

func (m *ConfigVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigVersionResponse.Unmarshal(m, b)
}
func (m *ConfigVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigVersionResponse.Marshal(b, m, deterministic)
}
func (m *ConfigVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigVersionResponse.Merge(m, src)
}
func (m *ConfigVersionResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigVersionResponse.Size(m)
}
func (m *ConfigVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigVersionResponse proto.InternalMessageInfo

func (m *ConfigVersionResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// Results from Host
type PingResponse struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Rtt                  int64    `protobuf:"varint,5,opt,name=Rtt,proto3" json:"Rtt,omitempty"`
	Err                  string   `protobuf:"bytes,6,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{4}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *PingResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PingResponse) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *PingResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PingResponse) GetRtt() int64 {
	if m != nil {
		return m.Rtt
	}
	return 0
}

func (m *PingResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UDPResponse struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Pktssent             int64    `protobuf:"varint,5,opt,name=pktssent,proto3" json:"pktssent,omitempty"`
	Pktsrcvd             int64    `protobuf:"varint,6,opt,name=pktsrcvd,proto3" json:"pktsrcvd,omitempty"`
	Loss                 float64  `protobuf:"fixed64,7,opt,name=loss,proto3" json:"loss,omitempty"`
	Direction            string   `protobuf:"bytes,8,opt,name=direction,proto3" json:"direction,omitempty"`
	Err                  string   `protobuf:"bytes,9,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UDPResponse) Reset()         { *m = UDPResponse{} }
func (m *UDPResponse) String() string { return proto.CompactTextString(m) }
func (*UDPResponse) ProtoMessage()    {}
func (*UDPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{5}
}

func (m *UDPResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UDPResponse.Unmarshal(m, b)
}
func (m *UDPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UDPResponse.Marshal(b, m, deterministic)
}
func (m *UDPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UDPResponse.Merge(m, src)
}
func (m *UDPResponse) XXX_Size() int {
	return xxx_messageInfo_UDPResponse.Size(m)
}
func (m *UDPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UDPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UDPResponse proto.InternalMessageInfo

func (m *UDPResponse) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *UDPResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UDPResponse) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *UDPResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *UDPResponse) GetPktssent() int64 {
	if m != nil {
		return m.Pktssent
	}
	return 0
}

func (m *UDPResponse) GetPktsrcvd() int64 {
	if m != nil {
		return m.Pktsrcvd
	}
	return 0
}

func (m *UDPResponse) GetLoss() float64 {
	if m != nil {
		return m.Loss
	}
	return 0
}

func (m *UDPResponse) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *UDPResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReceivedAck struct {
	Received             bool     `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceivedAck) Reset()         { *m = ReceivedAck{} }
func (m *ReceivedAck) String() string { return proto.CompactTextString(m) }
func (*ReceivedAck) ProtoMessage()    {}
func (*ReceivedAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{6}
}

func (m *ReceivedAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedAck.Unmarshal(m, b)
}
func (m *ReceivedAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedAck.Marshal(b, m, deterministic)
}
func (m *ReceivedAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedAck.Merge(m, src)
}
func (m *ReceivedAck) XXX_Size() int {
	return xxx_messageInfo_ReceivedAck.Size(m)
}
func (m *ReceivedAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedAck.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedAck proto.InternalMessageInfo

func (m *ReceivedAck) GetReceived() bool {
	if m != nil {
		return m.Received
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PktCounter struct {
	Counter              int32    `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PktCounter) Reset()         { *m = PktCounter{} }
func (m *PktCounter) String() string { return proto.CompactTextString(m) }
func (*PktCounter) ProtoMessage()    {}
func (*PktCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{8}
}

func (m *PktCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PktCounter.Unmarshal(m, b)
}
func (m *PktCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PktCounter.Marshal(b, m, deterministic)
}
func (m *PktCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PktCounter.Merge(m, src)
}
func (m *PktCounter) XXX_Size() int {
	return xxx_messageInfo_PktCounter.Size(m)
}
func (m *PktCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_PktCounter.DiscardUnknown(m)
}

var xxx_messageInfo_PktCounter proto.InternalMessageInfo

func (m *PktCounter) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*Probe)(nil), "configfetcher.Probe")
	proto.RegisterType((*ConfigRequest)(nil), "configfetcher.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "configfetcher.ConfigResponse")
	proto.RegisterType((*ConfigVersionResponse)(nil), "configfetcher.ConfigVersionResponse")
	proto.RegisterType((*PingResponse)(nil), "configfetcher.PingResponse")
	proto.RegisterType((*UDPResponse)(nil), "configfetcher.UDPResponse")
	proto.RegisterType((*ReceivedAck)(nil), "configfetcher.ReceivedAck")
	proto.RegisterType((*Empty)(nil), "configfetcher.Empty")
	proto.RegisterType((*PktCounter)(nil), "configfetcher.PktCounter")
}

func init() { proto.RegisterFile("configFetcher.proto", fileDescriptor_401a461fcce453dc) }

var fileDescriptor_401a461fcce453dc = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x9b, 0x3a, 0x8d, 0x27, 0xf4, 0x87, 0xa5, 0x20, 0x63, 0x0a, 0xaa, 0x56, 0x15, 0x2a,
	0x12, 0x2a, 0x22, 0xbc, 0x23, 0xf5, 0x8f, 0x0a, 0x84, 0x50, 0xb4, 0x94, 0xaa, 0xe2, 0xcd, 0xb5,
	0xa7, 0xc5, 0x4a, 0x6b, 0x9b, 0xdd, 0x4d, 0xa4, 0x5c, 0x80, 0x13, 0x70, 0x05, 0x8e, 0xc0, 0x29,
	0x38, 0x03, 0x77, 0x41, 0x3b, 0xeb, 0xbf, 0x58, 0x81, 0x3e, 0xf1, 0x36, 0xf3, 0x7d, 0xbb, 0xe3,
	0x99, 0x6f, 0xbe, 0x35, 0xdc, 0x8b, 0xb2, 0xf4, 0x32, 0xb9, 0x7a, 0x83, 0x3a, 0xfa, 0x82, 0x72,
	0x2f, 0x97, 0x99, 0xce, 0xd8, 0xaa, 0x05, 0x2f, 0x2d, 0xc8, 0x7f, 0x39, 0xe0, 0x8e, 0x64, 0x76,
	0x81, 0x8c, 0xc1, 0x72, 0x1a, 0xde, 0xa0, 0xef, 0x6c, 0x3b, 0xbb, 0x9e, 0xa0, 0x98, 0xf9, 0xb0,
	0x92, 0x1b, 0xf2, 0xed, 0x91, 0xbf, 0xb4, 0xed, 0xec, 0xba, 0xa2, 0x4c, 0xd9, 0x26, 0xb8, 0x49,
	0x7e, 0x86, 0xd2, 0xef, 0x12, 0x6e, 0x13, 0xb6, 0x05, 0x1e, 0x1d, 0x38, 0x9d, 0xe5, 0xe8, 0x2f,
	0x53, 0xa1, 0x1a, 0x60, 0x3b, 0xb0, 0xaa, 0x43, 0x79, 0x85, 0x7a, 0x3f, 0x8e, 0x25, 0x2a, 0xe5,
	0xbb, 0x74, 0x62, 0x1e, 0x64, 0x4f, 0x00, 0xd2, 0xc9, 0xcd, 0x28, 0x8c, 0xc6, 0xa8, 0x95, 0xdf,
	0xa3, 0xf2, 0x0d, 0x84, 0x05, 0xd0, 0x4f, 0x52, 0x8d, 0x72, 0x1a, 0x5e, 0xfb, 0x2b, 0xc4, 0x56,
	0x39, 0x7f, 0x01, 0xab, 0x87, 0x34, 0x9e, 0xc0, 0xaf, 0x13, 0x54, 0xda, 0x14, 0x8b, 0x71, 0x9a,
	0x44, 0xf8, 0xa1, 0x1e, 0xad, 0x81, 0xf0, 0x73, 0x58, 0x2b, 0x2f, 0xa8, 0x3c, 0x4b, 0x15, 0x8d,
	0x3c, 0x45, 0xa9, 0x92, 0x2c, 0xa5, 0xe3, 0xae, 0x28, 0x53, 0xf6, 0x1c, 0x7a, 0x34, 0x8b, 0xf2,
	0x97, 0xb6, 0xbb, 0xbb, 0x83, 0xe1, 0xe6, 0xde, 0x9c, 0x94, 0x7b, 0x24, 0xa3, 0x28, 0xce, 0xf0,
	0x97, 0x70, 0xdf, 0x56, 0x3e, 0xb3, 0xd7, 0x6f, 0xff, 0x00, 0xff, 0xe6, 0xc0, 0x9d, 0x51, 0x92,
	0xd6, 0xbd, 0x98, 0x95, 0x64, 0x71, 0xbd, 0x92, 0x2c, 0x26, 0x4c, 0x1b, 0x75, 0x97, 0x2c, 0x66,
	0x62, 0xf6, 0x00, 0x7a, 0x56, 0x43, 0xda, 0x86, 0x27, 0x8a, 0xcc, 0x9c, 0x0d, 0xe3, 0x58, 0x16,
	0x9b, 0xa0, 0x98, 0x6d, 0x40, 0x57, 0x68, 0x4d, 0xd2, 0x77, 0x85, 0x09, 0x0d, 0x82, 0x52, 0x92,
	0xd2, 0x9e, 0x30, 0x21, 0xff, 0xed, 0xc0, 0xe0, 0xd3, 0xd1, 0xe8, 0xbf, 0xf6, 0x11, 0x40, 0x3f,
	0x1f, 0x6b, 0xa5, 0x30, 0x2d, 0x9b, 0xa9, 0xf2, 0x92, 0x93, 0xd1, 0x34, 0xa6, 0xb6, 0x0a, 0xce,
	0xe4, 0xa6, 0xd6, 0x75, 0xa6, 0x14, 0xad, 0xde, 0x11, 0x14, 0x1b, 0xdb, 0xc5, 0x89, 0xc4, 0x48,
	0x1b, 0x51, 0xfb, 0xd6, 0x76, 0x15, 0x50, 0xce, 0xe7, 0xd5, 0xf3, 0x3d, 0x83, 0x81, 0xc0, 0x08,
	0x93, 0x29, 0xc6, 0xfb, 0xd1, 0xd8, 0x7c, 0x4e, 0x16, 0x29, 0x8d, 0xd8, 0x17, 0x55, 0xce, 0x57,
	0xc0, 0x3d, 0xbe, 0xc9, 0xf5, 0x8c, 0x3f, 0x05, 0x18, 0x8d, 0xf5, 0x61, 0x36, 0x31, 0x5e, 0x33,
	0x4b, 0x8c, 0x6c, 0x58, 0x2e, 0xb1, 0x48, 0x87, 0x3f, 0x9d, 0xd2, 0x83, 0x1f, 0x51, 0x1a, 0x9f,
	0xb1, 0x77, 0xe0, 0x9d, 0xa0, 0xb6, 0x18, 0xdb, 0x6a, 0x99, 0x66, 0xce, 0xae, 0xc1, 0xe3, 0xbf,
	0xb0, 0x76, 0x0f, 0xbc, 0xc3, 0xce, 0x61, 0xa3, 0xaa, 0x55, 0x18, 0xeb, 0x96, 0x92, 0x3b, 0x0b,
	0xd9, 0x96, 0x29, 0x79, 0x67, 0xf8, 0xdd, 0x81, 0x81, 0x31, 0x5f, 0xd9, 0xf5, 0x6b, 0x58, 0x13,
	0x18, 0xc6, 0xb3, 0xd3, 0xac, 0x90, 0x8a, 0xb5, 0xfd, 0x4e, 0xba, 0x04, 0x0b, 0x51, 0xde, 0x61,
	0xc7, 0xd4, 0xa9, 0x7d, 0xb4, 0xa5, 0x6a, 0x8b, 0x2b, 0x3c, 0x6c, 0xbf, 0xa3, 0x4a, 0x66, 0xde,
	0x19, 0xfe, 0x70, 0xc0, 0x3b, 0x3a, 0xa8, 0xa5, 0x5c, 0x3b, 0x41, 0xdd, 0xb4, 0x66, 0xd0, 0xba,
	0xdc, 0xe0, 0x82, 0x36, 0xd7, 0xd8, 0x39, 0xef, 0xb0, 0xf7, 0xb0, 0x6e, 0x1a, 0x6c, 0xbe, 0xb7,
	0x47, 0xed, 0x4e, 0x1a, 0xe4, 0xbf, 0xab, 0x1d, 0xdc, 0xfd, 0xbc, 0x3e, 0x47, 0xe7, 0x17, 0x17,
	0x3d, 0xfa, 0xe1, 0xbe, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x74, 0xbd, 0x07, 0xdd, 0x87, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	GetConfigVersion(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigVersionResponse, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/configfetcher.ConfigService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetConfigVersion(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigVersionResponse, error) {
	out := new(ConfigVersionResponse)
	err := c.cc.Invoke(ctx, "/configfetcher.ConfigService/GetConfigVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	GetConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	GetConfigVersion(context.Context, *ConfigRequest) (*ConfigVersionResponse, error)
}

// UnimplementedConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (*UnimplementedConfigServiceServer) GetConfig(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (*UnimplementedConfigServiceServer) GetConfigVersion(ctx context.Context, req *ConfigRequest) (*ConfigVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigVersion not implemented")
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.ConfigService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetConfigVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfigVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.ConfigService/GetConfigVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfigVersion(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configfetcher.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ConfigService_GetConfig_Handler,
		},
		{
			MethodName: "GetConfigVersion",
			Handler:    _ConfigService_GetConfigVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configFetcher.proto",
}

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingServiceClient interface {
	ReadyToReceive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetPacketCounter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PktCounter, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingServiceClient(cc *grpc.ClientConn) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) ReadyToReceive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/configfetcher.PingService/ReadyToReceive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingServiceClient) GetPacketCounter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PktCounter, error) {
	out := new(PktCounter)
	err := c.cc.Invoke(ctx, "/configfetcher.PingService/GetPacketCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
type PingServiceServer interface {
	ReadyToReceive(context.Context, *Empty) (*Empty, error)
	GetPacketCounter(context.Context, *Empty) (*PktCounter, error)
}

// UnimplementedPingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (*UnimplementedPingServiceServer) ReadyToReceive(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadyToReceive not implemented")
}
func (*UnimplementedPingServiceServer) GetPacketCounter(ctx context.Context, req *Empty) (*PktCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPacketCounter not implemented")
}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_ReadyToReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).ReadyToReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.PingService/ReadyToReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).ReadyToReceive(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingService_GetPacketCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).GetPacketCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.PingService/GetPacketCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).GetPacketCounter(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configfetcher.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadyToReceive",
			Handler:    _PingService_ReadyToReceive_Handler,
		},
		{
			MethodName: "GetPacketCounter",
			Handler:    _PingService_GetPacketCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configFetcher.proto",
}

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBServiceClient interface {
	GetUDPResponse(ctx context.Context, in *UDPResponse, opts ...grpc.CallOption) (*ReceivedAck, error)
	GetPingResponse(ctx context.Context, in *PingResponse, opts ...grpc.CallOption) (*ReceivedAck, error)
}

type dBServiceClient struct {
	cc *grpc.ClientConn
}

func NewDBServiceClient(cc *grpc.ClientConn) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) GetUDPResponse(ctx context.Context, in *UDPResponse, opts ...grpc.CallOption) (*ReceivedAck, error) {
	out := new(ReceivedAck)
	err := c.cc.Invoke(ctx, "/configfetcher.DBService/GetUDPResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetPingResponse(ctx context.Context, in *PingResponse, opts ...grpc.CallOption) (*ReceivedAck, error) {
	out := new(ReceivedAck)
	err := c.cc.Invoke(ctx, "/configfetcher.DBService/GetPingResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
type DBServiceServer interface {
	GetUDPResponse(context.Context, *UDPResponse) (*ReceivedAck, error)
	GetPingResponse(context.Context, *PingResponse) (*ReceivedAck, error)
}

// UnimplementedDBServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (*UnimplementedDBServiceServer) GetUDPResponse(ctx context.Context, req *UDPResponse) (*ReceivedAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUDPResponse not implemented")
}
func (*UnimplementedDBServiceServer) GetPingResponse(ctx context.Context, req *PingResponse) (*ReceivedAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPingResponse not implemented")
}

func RegisterDBServiceServer(s *grpc.Server, srv DBServiceServer) {
	s.RegisterService(&_DBService_serviceDesc, srv)
}

func _DBService_GetUDPResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UDPResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetUDPResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.DBService/GetUDPResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetUDPResponse(ctx, req.(*UDPResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetPingResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetPingResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.DBService/GetPingResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetPingResponse(ctx, req.(*PingResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configfetcher.DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUDPResponse",
			Handler:    _DBService_GetUDPResponse_Handler,
		},
		{
			MethodName: "GetPingResponse",
			Handler:    _DBService_GetPingResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configFetcher.proto",
}
