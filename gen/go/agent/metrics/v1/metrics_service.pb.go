// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opentelemetry/proto/agent/metrics/v1/metrics_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
	v11 "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"
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

type ExportMetricsServiceRequest struct {
	// Metric data. An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceMetrics      []*ResourceMetrics `protobuf:"bytes,1,rep,name=resource_metrics,json=resourceMetrics,proto3" json:"resource_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportMetricsServiceRequest) Reset()         { *m = ExportMetricsServiceRequest{} }
func (m *ExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceRequest) ProtoMessage()    {}
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6663e4888d72f6f3, []int{0}
}

func (m *ExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceRequest.Unmarshal(m, b)
}
func (m *ExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceRequest.Merge(m, src)
}
func (m *ExportMetricsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceRequest.Size(m)
}
func (m *ExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceRequest proto.InternalMessageInfo

func (m *ExportMetricsServiceRequest) GetResourceMetrics() []*ResourceMetrics {
	if m != nil {
		return m.ResourceMetrics
	}
	return nil
}

// A collection of metrics from a Resource.
type ResourceMetrics struct {
	// A list of metrics that originate from a resource.
	Metrics []*v1.Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// The resource for the metrics in this message that do not have an explicit
	// Metric.resource field set. If neither this field nor Metric.resource are set then no
	// resource info is known.
	Resource             *v11.Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceMetrics) Reset()         { *m = ResourceMetrics{} }
func (m *ResourceMetrics) String() string { return proto.CompactTextString(m) }
func (*ResourceMetrics) ProtoMessage()    {}
func (*ResourceMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_6663e4888d72f6f3, []int{1}
}

func (m *ResourceMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMetrics.Unmarshal(m, b)
}
func (m *ResourceMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMetrics.Marshal(b, m, deterministic)
}
func (m *ResourceMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMetrics.Merge(m, src)
}
func (m *ResourceMetrics) XXX_Size() int {
	return xxx_messageInfo_ResourceMetrics.Size(m)
}
func (m *ResourceMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMetrics proto.InternalMessageInfo

func (m *ResourceMetrics) GetMetrics() []*v1.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ResourceMetrics) GetResource() *v11.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportMetricsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportMetricsServiceResponse) Reset()         { *m = ExportMetricsServiceResponse{} }
func (m *ExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceResponse) ProtoMessage()    {}
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6663e4888d72f6f3, []int{2}
}

func (m *ExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceResponse.Unmarshal(m, b)
}
func (m *ExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceResponse.Merge(m, src)
}
func (m *ExportMetricsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceResponse.Size(m)
}
func (m *ExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportMetricsServiceRequest)(nil), "opentelemetry.proto.agent.metrics.v1.ExportMetricsServiceRequest")
	proto.RegisterType((*ResourceMetrics)(nil), "opentelemetry.proto.agent.metrics.v1.ResourceMetrics")
	proto.RegisterType((*ExportMetricsServiceResponse)(nil), "opentelemetry.proto.agent.metrics.v1.ExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/agent/metrics/v1/metrics_service.proto", fileDescriptor_6663e4888d72f6f3)
}

var fileDescriptor_6663e4888d72f6f3 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x8d, 0xc2, 0x94, 0x0c, 0x9c, 0xc4, 0xcb, 0xa8, 0x22, 0xa3, 0x88, 0x4e, 0x70, 0x09,
	0xab, 0x78, 0xf1, 0xa4, 0x85, 0x1e, 0x85, 0x51, 0x6f, 0x5e, 0xa6, 0x2b, 0x1f, 0xb5, 0xe0, 0x9a,
	0x9a, 0xa4, 0x45, 0x4f, 0x1e, 0x3d, 0x7b, 0xf5, 0xec, 0x1f, 0x2a, 0x6d, 0xda, 0xb9, 0xb8, 0x20,
	0x85, 0xdd, 0xc2, 0x97, 0xf7, 0x7b, 0xef, 0xe3, 0x91, 0xe0, 0x2b, 0x9e, 0x41, 0xaa, 0xe0, 0x19,
	0xe6, 0xa0, 0xc4, 0x1b, 0xcb, 0x04, 0x57, 0x9c, 0x3d, 0xc6, 0x90, 0x2a, 0x56, 0x4e, 0x92, 0x48,
	0xb2, 0x62, 0xdc, 0x1c, 0xa7, 0x12, 0x44, 0x91, 0x44, 0x40, 0x2b, 0x19, 0x39, 0x36, 0x58, 0x3d,
	0xa4, 0x15, 0x4b, 0x6b, 0x80, 0x16, 0x63, 0xe7, 0xdc, 0x96, 0xb0, 0xea, 0xad, 0x71, 0x87, 0xda,
	0xd4, 0x02, 0x24, 0xcf, 0x45, 0x04, 0xa5, 0xbc, 0x39, 0x6b, 0xbd, 0xfb, 0x8e, 0x0f, 0x82, 0xd7,
	0x8c, 0x0b, 0x75, 0xab, 0x6d, 0xee, 0xf4, 0x86, 0x21, 0xbc, 0xe4, 0x20, 0x15, 0x79, 0xc0, 0x7b,
	0x0d, 0x30, 0xad, 0x83, 0xfa, 0x68, 0xb0, 0x35, 0xec, 0x7a, 0x97, 0xb4, 0xcd, 0xf6, 0x34, 0xac,
	0xe9, 0xda, 0x3e, 0xec, 0x09, 0x73, 0xe0, 0x7e, 0x21, 0xdc, 0xfb, 0x23, 0x22, 0xd7, 0x78, 0xdb,
	0x0c, 0x3b, 0xb1, 0x86, 0x2d, 0xc5, 0x68, 0x32, 0x6c, 0x30, 0x12, 0xe0, 0x9d, 0x26, 0xa8, 0xbf,
	0x39, 0x40, 0xc3, 0xae, 0x77, 0x66, 0xb5, 0x58, 0xb4, 0xb1, 0xb4, 0x6a, 0xb8, 0x40, 0xdd, 0x23,
	0x7c, 0x68, 0x6f, 0x47, 0x66, 0x3c, 0x95, 0xe0, 0x7d, 0x23, 0xbc, 0x6b, 0x5e, 0x91, 0x4f, 0x84,
	0x3b, 0x9a, 0x21, 0x37, 0xed, 0x2a, 0xfa, 0xa7, 0x7f, 0xc7, 0x5f, 0xc7, 0x42, 0x2f, 0xe9, 0x6e,
	0xf8, 0x1f, 0x08, 0x9f, 0x26, 0xbc, 0x95, 0x95, 0xbf, 0x6f, 0xba, 0x4c, 0x4a, 0xd5, 0x04, 0xdd,
	0x07, 0x71, 0xa2, 0x9e, 0xf2, 0x19, 0x8d, 0xf8, 0x9c, 0x95, 0x3e, 0xa3, 0xdf, 0x37, 0x66, 0xd8,
	0x8e, 0xf4, 0x8b, 0x8b, 0x21, 0x65, 0xf1, 0xea, 0x47, 0x98, 0x75, 0xaa, 0xeb, 0x8b, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0x9e, 0x3c, 0xc5, 0x37, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error) {
	out := new(ExportMetricsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.agent.metrics.v1.MetricsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error)
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) Export(ctx context.Context, req *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportMetricsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.agent.metrics.v1.MetricsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).Export(ctx, req.(*ExportMetricsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.agent.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _MetricsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/agent/metrics/v1/metrics_service.proto",
}
