// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto

package frontendv2pb

import (
	context "context"
	fmt "fmt"
	stats "github.com/cortexproject/cortex/pkg/querier/stats"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryResultRequest struct {
	QueryID      uint64                 `protobuf:"varint,1,opt,name=queryID,proto3" json:"queryID,omitempty"`
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,2,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	Stats        *stats.Stats           `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *QueryResultRequest) Reset()      { *m = QueryResultRequest{} }
func (*QueryResultRequest) ProtoMessage() {}
func (*QueryResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a7e5cdf8261f06, []int{0}
}
func (m *QueryResultRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResultRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResultRequest.Merge(m, src)
}
func (m *QueryResultRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResultRequest proto.InternalMessageInfo

func (m *QueryResultRequest) GetQueryID() uint64 {
	if m != nil {
		return m.QueryID
	}
	return 0
}

func (m *QueryResultRequest) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *QueryResultRequest) GetStats() *stats.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type QueryResultResponse struct {
}

func (m *QueryResultResponse) Reset()      { *m = QueryResultResponse{} }
func (*QueryResultResponse) ProtoMessage() {}
func (*QueryResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a7e5cdf8261f06, []int{1}
}
func (m *QueryResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResultResponse.Merge(m, src)
}
func (m *QueryResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryResultRequest)(nil), "frontendv2pb.QueryResultRequest")
	proto.RegisterType((*QueryResultResponse)(nil), "frontendv2pb.QueryResultResponse")
}

func init() {
	proto.RegisterFile("pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto", fileDescriptor_85a7e5cdf8261f06)
}

var fileDescriptor_85a7e5cdf8261f06 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0xfe, 0x4a, 0x6e, 0x27, 0x23, 0x50, 0xd4, 0xc1, 0x2a, 0x99, 0xba, 0x10, 0x4b,
	0x81, 0x01, 0x21, 0xb1, 0x54, 0xa8, 0x82, 0x8d, 0x86, 0x4e, 0x6c, 0x4d, 0x70, 0xd3, 0xbf, 0xb9,
	0xd4, 0x71, 0x5a, 0xba, 0xf1, 0x04, 0x88, 0xc7, 0xe0, 0x51, 0x18, 0x3b, 0x76, 0xa4, 0xe9, 0xc2,
	0xd8, 0x47, 0x40, 0xb1, 0xdb, 0x28, 0x15, 0x12, 0x8b, 0x75, 0x9f, 0xee, 0xfb, 0xe9, 0xbe, 0x3b,
	0xe3, 0xeb, 0x68, 0x10, 0xb0, 0x21, 0x0c, 0x7a, 0x1d, 0x01, 0xa1, 0xe4, 0xe1, 0x0b, 0xcb, 0x8b,
	0x89, 0x93, 0xd7, 0x13, 0x27, 0xf2, 0x72, 0x61, 0x47, 0x02, 0x24, 0x90, 0x72, 0xb1, 0x59, 0xb9,
	0x08, 0x7a, 0xb2, 0x9b, 0x78, 0xb6, 0x0f, 0x23, 0x16, 0x40, 0x00, 0x4c, 0x99, 0xbc, 0xa4, 0xa3,
	0x94, 0x12, 0xaa, 0xd2, 0x70, 0xe5, 0xaa, 0x60, 0x9f, 0xf2, 0xf6, 0x84, 0x4f, 0x41, 0x0c, 0x62,
	0xe6, 0xc3, 0x68, 0x04, 0x21, 0xeb, 0x4a, 0x19, 0x05, 0x22, 0xf2, 0xf3, 0x62, 0x43, 0xdd, 0x16,
	0x28, 0x1f, 0x84, 0xe4, 0xaf, 0x91, 0x80, 0x3e, 0xf7, 0xe5, 0x46, 0xb1, 0x6c, 0x99, 0x71, 0xc2,
	0x45, 0x8f, 0x0b, 0x16, 0xcb, 0xb6, 0x8c, 0xf5, 0xab, 0x71, 0xeb, 0x1d, 0x61, 0xd2, 0x4c, 0xb8,
	0x98, 0xb9, 0x3c, 0x4e, 0x86, 0xd2, 0xe5, 0xe3, 0x84, 0xc7, 0x92, 0x98, 0xf8, 0x38, 0x63, 0x66,
	0x0f, 0x77, 0x26, 0xaa, 0xa2, 0xda, 0x81, 0xbb, 0x95, 0xe4, 0x06, 0x97, 0xb3, 0x04, 0x2e, 0x8f,
	0x23, 0x08, 0x63, 0x6e, 0xee, 0x55, 0x51, 0xad, 0xe4, 0x9c, 0xd9, 0x79, 0xac, 0xfb, 0x56, 0xeb,
	0x71, 0xdb, 0x75, 0x77, 0xbc, 0xc4, 0xc2, 0x87, 0x6a, 0xb6, 0xb9, 0xaf, 0xa0, 0xb2, 0xad, 0x93,
	0x3c, 0x65, 0xaf, 0xab, 0x5b, 0xd6, 0x29, 0x3e, 0xd9, 0xc9, 0xa3, 0x51, 0xa7, 0x8f, 0x49, 0x63,
	0x73, 0xdb, 0x06, 0x88, 0xa6, 0xde, 0x87, 0xb4, 0x70, 0xa9, 0x60, 0x26, 0x55, 0xbb, 0x78, 0x7f,
	0xfb, 0xef, 0x5e, 0x95, 0xf3, 0x7f, 0x1c, 0x7a, 0x92, 0x65, 0xd4, 0xeb, 0xf3, 0x25, 0x35, 0x16,
	0x4b, 0x6a, 0xac, 0x97, 0x14, 0xbd, 0xa5, 0x14, 0x7d, 0xa6, 0x14, 0x7d, 0xa5, 0x14, 0xcd, 0x53,
	0x8a, 0xbe, 0x53, 0x8a, 0x7e, 0x52, 0x6a, 0xac, 0x53, 0x8a, 0x3e, 0x56, 0xd4, 0x98, 0xaf, 0xa8,
	0xb1, 0x58, 0x51, 0xe3, 0x79, 0xe7, 0xef, 0xbd, 0x23, 0x75, 0xde, 0xcb, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0x8d, 0x41, 0x35, 0x4c, 0x02, 0x00, 0x00,
}

func (this *QueryResultRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResultRequest)
	if !ok {
		that2, ok := that.(QueryResultRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.QueryID != that1.QueryID {
		return false
	}
	if !this.HttpResponse.Equal(that1.HttpResponse) {
		return false
	}
	if !this.Stats.Equal(that1.Stats) {
		return false
	}
	return true
}
func (this *QueryResultResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResultResponse)
	if !ok {
		that2, ok := that.(QueryResultResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *QueryResultRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&frontendv2pb.QueryResultRequest{")
	s = append(s, "QueryID: "+fmt.Sprintf("%#v", this.QueryID)+",\n")
	if this.HttpResponse != nil {
		s = append(s, "HttpResponse: "+fmt.Sprintf("%#v", this.HttpResponse)+",\n")
	}
	if this.Stats != nil {
		s = append(s, "Stats: "+fmt.Sprintf("%#v", this.Stats)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryResultResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&frontendv2pb.QueryResultResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFrontend(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendForQuerierClient is the client API for FrontendForQuerier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendForQuerierClient interface {
	QueryResult(ctx context.Context, in *QueryResultRequest, opts ...grpc.CallOption) (*QueryResultResponse, error)
}

type frontendForQuerierClient struct {
	cc *grpc.ClientConn
}

func NewFrontendForQuerierClient(cc *grpc.ClientConn) FrontendForQuerierClient {
	return &frontendForQuerierClient{cc}
}

func (c *frontendForQuerierClient) QueryResult(ctx context.Context, in *QueryResultRequest, opts ...grpc.CallOption) (*QueryResultResponse, error) {
	out := new(QueryResultResponse)
	err := c.cc.Invoke(ctx, "/frontendv2pb.FrontendForQuerier/QueryResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendForQuerierServer is the server API for FrontendForQuerier service.
type FrontendForQuerierServer interface {
	QueryResult(context.Context, *QueryResultRequest) (*QueryResultResponse, error)
}

// UnimplementedFrontendForQuerierServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendForQuerierServer struct {
}

func (*UnimplementedFrontendForQuerierServer) QueryResult(ctx context.Context, req *QueryResultRequest) (*QueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResult not implemented")
}

func RegisterFrontendForQuerierServer(s *grpc.Server, srv FrontendForQuerierServer) {
	s.RegisterService(&_FrontendForQuerier_serviceDesc, srv)
}

func _FrontendForQuerier_QueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendForQuerierServer).QueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontendv2pb.FrontendForQuerier/QueryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendForQuerierServer).QueryResult(ctx, req.(*QueryResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FrontendForQuerier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontendv2pb.FrontendForQuerier",
	HandlerType: (*FrontendForQuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryResult",
			Handler:    _FrontendForQuerier_QueryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/lokifrontend/frontend/v2/frontendv2pb/frontend.proto",
}

func (m *QueryResultRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResultRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResultRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HttpResponse != nil {
		{
			size, err := m.HttpResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.QueryID != 0 {
		i = encodeVarintFrontend(dAtA, i, uint64(m.QueryID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintFrontend(dAtA []byte, offset int, v uint64) int {
	offset -= sovFrontend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryResultRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QueryID != 0 {
		n += 1 + sovFrontend(uint64(m.QueryID))
	}
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *QueryResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovFrontend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFrontend(x uint64) (n int) {
	return sovFrontend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueryResultRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryResultRequest{`,
		`QueryID:` + fmt.Sprintf("%v", this.QueryID) + `,`,
		`HttpResponse:` + strings.Replace(fmt.Sprintf("%v", this.HttpResponse), "HTTPResponse", "httpgrpc.HTTPResponse", 1) + `,`,
		`Stats:` + strings.Replace(fmt.Sprintf("%v", this.Stats), "Stats", "stats.Stats", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryResultResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryResultResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringFrontend(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueryResultRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResultRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResultRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryID", wireType)
			}
			m.QueryID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &stats.Stats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFrontend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFrontend
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFrontend(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFrontend
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFrontend = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrontend   = fmt.Errorf("proto: integer overflow")
)