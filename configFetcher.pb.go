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

type Probes struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Probes               []*Probe `protobuf:"bytes,2,rep,name=probes,proto3" json:"probes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probes) Reset()         { *m = Probes{} }
func (m *Probes) String() string { return proto.CompactTextString(m) }
func (*Probes) ProtoMessage()    {}
func (*Probes) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{0}
}

func (m *Probes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probes.Unmarshal(m, b)
}
func (m *Probes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probes.Marshal(b, m, deterministic)
}
func (m *Probes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probes.Merge(m, src)
}
func (m *Probes) XXX_Size() int {
	return xxx_messageInfo_Probes.Size(m)
}
func (m *Probes) XXX_DiscardUnknown() {
	xxx_messageInfo_Probes.DiscardUnknown(m)
}

var xxx_messageInfo_Probes proto.InternalMessageInfo

func (m *Probes) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Probes) GetProbes() []*Probe {
	if m != nil {
		return m.Probes
	}
	return nil
}

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
	return fileDescriptor_401a461fcce453dc, []int{1}
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
	return fileDescriptor_401a461fcce453dc, []int{2}
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
	Probes               *Probes  `protobuf:"bytes,1,opt,name=probes,proto3" json:"probes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_401a461fcce453dc, []int{3}
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

func (m *ConfigResponse) GetProbes() *Probes {
	if m != nil {
		return m.Probes
	}
	return nil
}

// Results from Host
type PingResponse struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Port                 int64    `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	Rtt                  int64    `protobuf:"varint,4,opt,name=Rtt,proto3" json:"Rtt,omitempty"`
	Err                  string   `protobuf:"bytes,5,opt,name=err,proto3" json:"err,omitempty"`
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

func (m *PingResponse) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PingResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PingResponse) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
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
	return fileDescriptor_401a461fcce453dc, []int{5}
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

func init() {
	proto.RegisterType((*Probes)(nil), "configfetcher.Probes")
	proto.RegisterType((*Probe)(nil), "configfetcher.Probe")
	proto.RegisterType((*ConfigRequest)(nil), "configfetcher.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "configfetcher.ConfigResponse")
	proto.RegisterType((*PingResponse)(nil), "configfetcher.PingResponse")
	proto.RegisterType((*ReceivedAck)(nil), "configfetcher.ReceivedAck")
}

func init() { proto.RegisterFile("configFetcher.proto", fileDescriptor_401a461fcce453dc) }

var fileDescriptor_401a461fcce453dc = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x6f, 0x6b, 0xd4, 0x40,
	0x10, 0x87, 0x7b, 0xcd, 0x5d, 0xda, 0xcc, 0x79, 0x56, 0xc7, 0x0a, 0x4b, 0xac, 0x52, 0x82, 0x2f,
	0x2a, 0xe8, 0x09, 0xe7, 0x07, 0x90, 0xaa, 0x58, 0x14, 0x91, 0xb0, 0x8a, 0x2f, 0x7c, 0x97, 0x4b,
	0xa6, 0x6d, 0xa8, 0xdd, 0x8d, 0xbb, 0x7b, 0x01, 0x3f, 0x92, 0x9f, 0xc5, 0x2f, 0x25, 0x3b, 0xf9,
	0x73, 0xc9, 0xa1, 0xef, 0x66, 0x9e, 0x99, 0xcc, 0xf2, 0x7b, 0x20, 0xf0, 0x20, 0xd7, 0xea, 0xb2,
	0xbc, 0x7a, 0x4f, 0x2e, 0xbf, 0x26, 0xb3, 0xac, 0x8c, 0x76, 0x1a, 0x17, 0x0d, 0xbc, 0x6c, 0x60,
	0x92, 0x42, 0x98, 0x1a, 0xbd, 0x26, 0x8b, 0x02, 0x0e, 0x6a, 0x32, 0xb6, 0xd4, 0x4a, 0x4c, 0x4e,
	0x27, 0x67, 0x33, 0xd9, 0xb5, 0xf8, 0x1c, 0xc2, 0x8a, 0x77, 0xc4, 0xfe, 0x69, 0x70, 0x36, 0x5f,
	0x1d, 0x2f, 0x47, 0x37, 0x96, 0x7c, 0x40, 0xb6, 0x3b, 0xc9, 0x9f, 0x09, 0xcc, 0x98, 0x20, 0xc2,
	0x54, 0x65, 0xb7, 0xc4, 0xe7, 0x22, 0xc9, 0xb5, 0x7f, 0x85, 0xf7, 0x3e, 0xbc, 0x13, 0xfb, 0xcd,
	0x2b, 0x6d, 0x8b, 0xc7, 0x30, 0x2b, 0xab, 0x6f, 0x64, 0x44, 0xc0, 0xbc, 0x69, 0xf0, 0x04, 0x22,
	0x5e, 0xf8, 0xfa, 0xab, 0x22, 0x31, 0xe5, 0x43, 0x5b, 0x80, 0x4f, 0x61, 0xe1, 0x32, 0x73, 0x45,
	0xee, 0xbc, 0x28, 0x0c, 0x59, 0x2b, 0x66, 0xbc, 0x31, 0x86, 0xf8, 0x04, 0x40, 0x6d, 0x6e, 0xd3,
	0x2c, 0xbf, 0x21, 0x67, 0x45, 0xc8, 0xe7, 0x07, 0x04, 0x63, 0x38, 0x2c, 0x95, 0x23, 0x53, 0x67,
	0x3f, 0xc4, 0x01, 0x4f, 0xfb, 0x3e, 0x79, 0x09, 0x8b, 0xb7, 0x1c, 0x56, 0xd2, 0xcf, 0x0d, 0x59,
	0xe7, 0x8f, 0x15, 0x54, 0x97, 0x39, 0x7d, 0xde, 0x46, 0x1b, 0x90, 0xe4, 0x35, 0xdc, 0xed, 0x3e,
	0xb0, 0x95, 0x56, 0x96, 0xf0, 0x45, 0xaf, 0xcf, 0x6f, 0xcf, 0x57, 0x0f, 0xff, 0xa5, 0xcf, 0xf6,
	0xfe, 0x14, 0xdc, 0x49, 0x4b, 0xb5, 0xfd, 0x1c, 0x61, 0x7a, 0xad, 0xad, 0xeb, 0x2c, 0xfa, 0xda,
	0x33, 0x1f, 0x8e, 0x15, 0x46, 0x92, 0x6b, 0xcf, 0x52, 0x6d, 0x1c, 0xeb, 0x0b, 0x24, 0xd7, 0x78,
	0x0f, 0x02, 0xe9, 0x1c, 0x7b, 0x0b, 0xa4, 0x2f, 0x3d, 0x21, 0x63, 0x5a, 0x4f, 0xbe, 0x4c, 0x9e,
	0xc1, 0x5c, 0x52, 0x4e, 0x65, 0x4d, 0xc5, 0x79, 0x7e, 0xe3, 0x65, 0x98, 0xb6, 0xe5, 0x27, 0x0f,
	0x65, 0xdf, 0xaf, 0x7e, 0x4f, 0x3a, 0x1b, 0x5f, 0xc8, 0xf8, 0xc4, 0xf8, 0x11, 0xa2, 0x0b, 0x72,
	0x0d, 0xc3, 0x93, 0x9d, 0x60, 0x23, 0x71, 0xf1, 0xe3, 0xff, 0x4c, 0x9b, 0x98, 0xc9, 0x1e, 0x7e,
	0x82, 0xa3, 0x0b, 0x72, 0xa3, 0xec, 0x8f, 0x76, 0x55, 0x0d, 0x86, 0x71, 0xbc, 0x33, 0x1c, 0xa4,
	0x48, 0xf6, 0xde, 0xdc, 0xff, 0x7e, 0x34, 0x1a, 0x57, 0xeb, 0x75, 0xc8, 0x7f, 0xc0, 0xab, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x86, 0xdc, 0xff, 0x18, 0x03, 0x00, 0x00,
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
	GetPingResponse(ctx context.Context, in *PingResponse, opts ...grpc.CallOption) (*ReceivedAck, error)
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

func (c *configServiceClient) GetPingResponse(ctx context.Context, in *PingResponse, opts ...grpc.CallOption) (*ReceivedAck, error) {
	out := new(ReceivedAck)
	err := c.cc.Invoke(ctx, "/configfetcher.ConfigService/GetPingResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	GetConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	GetPingResponse(context.Context, *PingResponse) (*ReceivedAck, error)
}

// UnimplementedConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (*UnimplementedConfigServiceServer) GetConfig(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (*UnimplementedConfigServiceServer) GetPingResponse(ctx context.Context, req *PingResponse) (*ReceivedAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPingResponse not implemented")
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

func _ConfigService_GetPingResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetPingResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configfetcher.ConfigService/GetPingResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetPingResponse(ctx, req.(*PingResponse))
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
			MethodName: "GetPingResponse",
			Handler:    _ConfigService_GetPingResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configFetcher.proto",
}
