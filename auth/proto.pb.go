// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

/*
	Package auth is a generated protocol buffer package.

	It is generated from these files:
		proto.proto

	It has these top-level messages:
		Token
		UserStatusResponse
		Version
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import core "github.com/kubaj/mmo-demo/core"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptorProto, []int{0} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserStatusResponse struct {
	Logged bool `protobuf:"varint,1,opt,name=logged,proto3" json:"logged,omitempty"`
}

func (m *UserStatusResponse) Reset()                    { *m = UserStatusResponse{} }
func (m *UserStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*UserStatusResponse) ProtoMessage()               {}
func (*UserStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptorProto, []int{1} }

func (m *UserStatusResponse) GetLogged() bool {
	if m != nil {
		return m.Logged
	}
	return false
}

type Version struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptorProto, []int{2} }

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*UserStatusResponse)(nil), "auth.UserStatusResponse")
	proto.RegisterType((*Version)(nil), "auth.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Version, error)
	Check(ctx context.Context, in *core.HealthCheckRequest, opts ...grpc.CallOption) (*core.HealthCheckResponse, error)
	IsUserLogged(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserStatusResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/auth.AuthService/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Check(ctx context.Context, in *core.HealthCheckRequest, opts ...grpc.CallOption) (*core.HealthCheckResponse, error) {
	out := new(core.HealthCheckResponse)
	err := grpc.Invoke(ctx, "/auth.AuthService/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsUserLogged(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserStatusResponse, error) {
	out := new(UserStatusResponse)
	err := grpc.Invoke(ctx, "/auth.AuthService/IsUserLogged", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	GetVersion(context.Context, *google_protobuf.Empty) (*Version, error)
	Check(context.Context, *core.HealthCheckRequest) (*core.HealthCheckResponse, error)
	IsUserLogged(context.Context, *Token) (*UserStatusResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetVersion(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Check(ctx, req.(*core.HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsUserLogged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsUserLogged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/IsUserLogged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsUserLogged(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _AuthService_GetVersion_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _AuthService_Check_Handler,
		},
		{
			MethodName: "IsUserLogged",
			Handler:    _AuthService_IsUserLogged_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *UserStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Logged {
		dAtA[i] = 0x8
		i++
		if m.Logged {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Token) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func (m *UserStatusResponse) Size() (n int) {
	var l int
	_ = l
	if m.Logged {
		n += 2
	}
	return n
}

func (m *Version) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	return n
}

func sovProto(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *UserStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logged", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Logged = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProto
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProto
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
				next, err := skipProto(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthProto = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto.proto", fileDescriptorProto) }

var fileDescriptorProto = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x9b, 0x3f, 0x6d, 0xff, 0x3a, 0xb5, 0x08, 0x83, 0x96, 0x36, 0xda, 0x20, 0x59, 0xb9,
	0xd0, 0x89, 0xe8, 0xc2, 0x8d, 0x1b, 0x15, 0xf1, 0x03, 0x57, 0xa9, 0xba, 0x70, 0x37, 0x49, 0xaf,
	0x49, 0x6c, 0x33, 0x37, 0x66, 0x66, 0x0a, 0x6e, 0x7d, 0x05, 0x37, 0x3e, 0x92, 0xcb, 0x82, 0x2f,
	0x20, 0xd5, 0x07, 0x91, 0x4c, 0x52, 0x28, 0xe8, 0x26, 0xdc, 0x8f, 0x73, 0x32, 0xf7, 0x77, 0x48,
	0x2b, 0xcb, 0x51, 0x21, 0x33, 0x5f, 0x5a, 0xe7, 0x5a, 0xc5, 0xf6, 0x46, 0x84, 0x18, 0x8d, 0xc1,
	0x33, 0xb3, 0x40, 0x3f, 0x78, 0x90, 0x66, 0xea, 0xb9, 0x94, 0xd8, 0xbd, 0x10, 0xf3, 0x85, 0xd5,
	0x82, 0xdb, 0xde, 0xac, 0x7c, 0x3c, 0x4b, 0x3c, 0x2e, 0x04, 0x2a, 0xae, 0x12, 0x14, 0xb2, 0xdc,
	0xba, 0x7d, 0xd2, 0xb8, 0xc1, 0x11, 0x08, 0xba, 0x46, 0x1a, 0xaa, 0x28, 0xba, 0xd6, 0x96, 0xb5,
	0xbd, 0xec, 0x97, 0x8d, 0xbb, 0x43, 0xe8, 0xad, 0x84, 0x7c, 0xa0, 0xb8, 0xd2, 0xd2, 0x07, 0x99,
	0xa1, 0x90, 0x40, 0x3b, 0xa4, 0x39, 0xc6, 0x28, 0x82, 0xa1, 0x11, 0x2f, 0xf9, 0x55, 0xe7, 0xf6,
	0xc9, 0xff, 0x3b, 0xc8, 0x65, 0x82, 0x82, 0x52, 0x52, 0x17, 0x3c, 0x85, 0xea, 0x6f, 0xa6, 0xde,
	0x9f, 0x5a, 0xa4, 0x75, 0xac, 0x55, 0x3c, 0x80, 0x7c, 0x92, 0x84, 0x40, 0xaf, 0x08, 0x39, 0x07,
	0x35, 0x77, 0x74, 0x58, 0x79, 0x28, 0x9b, 0x53, 0xb0, 0xb3, 0x02, 0xd0, 0x6e, 0xb3, 0x02, 0x9f,
	0x55, 0x32, 0x77, 0xfd, 0xe5, 0xe3, 0xfb, 0xf5, 0xdf, 0x2a, 0x6d, 0x7b, 0xc5, 0xd8, 0x9b, 0x54,
	0xee, 0x23, 0xd2, 0x38, 0x8d, 0x21, 0x1c, 0xd1, 0x2e, 0x2b, 0xa2, 0x60, 0x17, 0xc0, 0xc7, 0x2a,
	0x36, 0x23, 0x1f, 0x9e, 0x34, 0x48, 0x65, 0xf7, 0xfe, 0xd8, 0x54, 0x40, 0x87, 0x64, 0xe5, 0x52,
	0x16, 0xa0, 0xd7, 0x06, 0x84, 0xb6, 0xca, 0x37, 0x4d, 0x32, 0x76, 0xb7, 0x6c, 0x7e, 0xe7, 0xe0,
	0xd6, 0x4e, 0xf6, 0xde, 0x67, 0x8e, 0x35, 0x9d, 0x39, 0xd6, 0xe7, 0xcc, 0xb1, 0xde, 0xbe, 0x9c,
	0xda, 0xbd, 0x13, 0x25, 0x2a, 0xd6, 0x01, 0x0b, 0x31, 0xf5, 0x46, 0x3a, 0xe0, 0x8f, 0x5e, 0x9a,
	0xe2, 0xee, 0x10, 0x52, 0x34, 0x07, 0x07, 0x4d, 0x83, 0x77, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x78, 0x54, 0x3b, 0x48, 0xe2, 0x01, 0x00, 0x00,
}
