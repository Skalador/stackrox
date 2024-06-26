// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/administration_usage_service.proto

package v1

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// TimeRange allows for requesting data by a time range.
type TimeRange struct {
	From                 *types.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *types.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7d8a93856728d1, []int{0}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return m.Size()
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetFrom() *types.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TimeRange) GetTo() *types.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TimeRange) MessageClone() proto.Message {
	return m.Clone()
}
func (m *TimeRange) Clone() *TimeRange {
	if m == nil {
		return nil
	}
	cloned := new(TimeRange)
	*cloned = *m

	cloned.From = m.From.Clone()
	cloned.To = m.To.Clone()
	return cloned
}

// SecuredUnitsUsageResponse holds the values of the currently observable
// administration usage metrics.
type SecuredUnitsUsageResponse struct {
	NumNodes             int64    `protobuf:"varint,1,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	NumCpuUnits          int64    `protobuf:"varint,2,opt,name=num_cpu_units,json=numCpuUnits,proto3" json:"num_cpu_units,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecuredUnitsUsageResponse) Reset()         { *m = SecuredUnitsUsageResponse{} }
func (m *SecuredUnitsUsageResponse) String() string { return proto.CompactTextString(m) }
func (*SecuredUnitsUsageResponse) ProtoMessage()    {}
func (*SecuredUnitsUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7d8a93856728d1, []int{1}
}
func (m *SecuredUnitsUsageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecuredUnitsUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecuredUnitsUsageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecuredUnitsUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecuredUnitsUsageResponse.Merge(m, src)
}
func (m *SecuredUnitsUsageResponse) XXX_Size() int {
	return m.Size()
}
func (m *SecuredUnitsUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecuredUnitsUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecuredUnitsUsageResponse proto.InternalMessageInfo

func (m *SecuredUnitsUsageResponse) GetNumNodes() int64 {
	if m != nil {
		return m.NumNodes
	}
	return 0
}

func (m *SecuredUnitsUsageResponse) GetNumCpuUnits() int64 {
	if m != nil {
		return m.NumCpuUnits
	}
	return 0
}

func (m *SecuredUnitsUsageResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SecuredUnitsUsageResponse) Clone() *SecuredUnitsUsageResponse {
	if m == nil {
		return nil
	}
	cloned := new(SecuredUnitsUsageResponse)
	*cloned = *m

	return cloned
}

// MaxSecuredUnitsUsageResponse holds the maximum values of the secured nodes
// and CPU Units (as reported by Kubernetes) with the time at which these
// values were aggregated, with the aggregation period accuracy (1h).
type MaxSecuredUnitsUsageResponse struct {
	MaxNodesAt           *types.Timestamp `protobuf:"bytes,1,opt,name=max_nodes_at,json=maxNodesAt,proto3" json:"max_nodes_at,omitempty"`
	MaxNodes             int64            `protobuf:"varint,2,opt,name=max_nodes,json=maxNodes,proto3" json:"max_nodes,omitempty"`
	MaxCpuUnitsAt        *types.Timestamp `protobuf:"bytes,3,opt,name=max_cpu_units_at,json=maxCpuUnitsAt,proto3" json:"max_cpu_units_at,omitempty"`
	MaxCpuUnits          int64            `protobuf:"varint,4,opt,name=max_cpu_units,json=maxCpuUnits,proto3" json:"max_cpu_units,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MaxSecuredUnitsUsageResponse) Reset()         { *m = MaxSecuredUnitsUsageResponse{} }
func (m *MaxSecuredUnitsUsageResponse) String() string { return proto.CompactTextString(m) }
func (*MaxSecuredUnitsUsageResponse) ProtoMessage()    {}
func (*MaxSecuredUnitsUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e7d8a93856728d1, []int{2}
}
func (m *MaxSecuredUnitsUsageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MaxSecuredUnitsUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MaxSecuredUnitsUsageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MaxSecuredUnitsUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxSecuredUnitsUsageResponse.Merge(m, src)
}
func (m *MaxSecuredUnitsUsageResponse) XXX_Size() int {
	return m.Size()
}
func (m *MaxSecuredUnitsUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxSecuredUnitsUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaxSecuredUnitsUsageResponse proto.InternalMessageInfo

func (m *MaxSecuredUnitsUsageResponse) GetMaxNodesAt() *types.Timestamp {
	if m != nil {
		return m.MaxNodesAt
	}
	return nil
}

func (m *MaxSecuredUnitsUsageResponse) GetMaxNodes() int64 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

func (m *MaxSecuredUnitsUsageResponse) GetMaxCpuUnitsAt() *types.Timestamp {
	if m != nil {
		return m.MaxCpuUnitsAt
	}
	return nil
}

func (m *MaxSecuredUnitsUsageResponse) GetMaxCpuUnits() int64 {
	if m != nil {
		return m.MaxCpuUnits
	}
	return 0
}

func (m *MaxSecuredUnitsUsageResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *MaxSecuredUnitsUsageResponse) Clone() *MaxSecuredUnitsUsageResponse {
	if m == nil {
		return nil
	}
	cloned := new(MaxSecuredUnitsUsageResponse)
	*cloned = *m

	cloned.MaxNodesAt = m.MaxNodesAt.Clone()
	cloned.MaxCpuUnitsAt = m.MaxCpuUnitsAt.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*TimeRange)(nil), "v1.TimeRange")
	proto.RegisterType((*SecuredUnitsUsageResponse)(nil), "v1.SecuredUnitsUsageResponse")
	proto.RegisterType((*MaxSecuredUnitsUsageResponse)(nil), "v1.MaxSecuredUnitsUsageResponse")
}

func init() {
	proto.RegisterFile("api/v1/administration_usage_service.proto", fileDescriptor_4e7d8a93856728d1)
}

var fileDescriptor_4e7d8a93856728d1 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0x37, 0x15, 0x6a, 0x1c, 0x22, 0x21, 0x5f, 0x08, 0xdb, 0x12, 0xaa, 0x3d, 0x41, 0x05,
	0x5e, 0x36, 0x48, 0x70, 0x80, 0x4b, 0x88, 0x50, 0x4f, 0x70, 0xd8, 0xb6, 0x12, 0x42, 0x48, 0x2b,
	0x77, 0xe3, 0xae, 0x2c, 0xb0, 0xbd, 0xb2, 0xc7, 0xab, 0xe5, 0x02, 0x12, 0xbf, 0xc0, 0x85, 0x0b,
	0xff, 0xc3, 0x11, 0x89, 0x1f, 0x40, 0x81, 0x0f, 0x41, 0xb6, 0x93, 0x88, 0xaa, 0x94, 0xe6, 0x96,
	0xc9, 0x9b, 0x37, 0xef, 0x8d, 0xe7, 0x2d, 0xba, 0x4b, 0x1b, 0x9e, 0xb5, 0x79, 0x46, 0xe7, 0x82,
	0x4b, 0x6e, 0x40, 0x53, 0xe0, 0x4a, 0x96, 0xd6, 0xd0, 0x9a, 0x95, 0x86, 0xe9, 0x96, 0x57, 0x8c,
	0x34, 0x5a, 0x81, 0xc2, 0x71, 0x9b, 0x27, 0x78, 0xd9, 0xce, 0x44, 0x03, 0xef, 0xc3, 0xff, 0xc9,
	0x6e, 0xad, 0x54, 0xfd, 0x8e, 0x65, 0x0e, 0xa2, 0x52, 0x2a, 0xf0, 0x23, 0xcc, 0x12, 0xbd, 0xbd,
	0x44, 0x7d, 0x75, 0x62, 0x4f, 0x33, 0xe0, 0x82, 0x19, 0xa0, 0xa2, 0x09, 0x0d, 0x69, 0x8d, 0xfa,
	0x47, 0x5c, 0xb0, 0x82, 0xca, 0x9a, 0x61, 0x82, 0xb6, 0x4e, 0xb5, 0x12, 0xa3, 0x68, 0x2f, 0xba,
	0x33, 0x98, 0x24, 0x24, 0x90, 0xc9, 0x8a, 0x4c, 0x8e, 0x56, 0xe4, 0xc2, 0xf7, 0xe1, 0x7d, 0x14,
	0x83, 0x1a, 0xc5, 0x97, 0x76, 0xc7, 0xa0, 0xd2, 0x37, 0xe8, 0xe6, 0x21, 0xab, 0xac, 0x66, 0xf3,
	0x63, 0xc9, 0xc1, 0x1c, 0xbb, 0x15, 0x0b, 0x66, 0x1a, 0x25, 0x0d, 0xc3, 0x3b, 0xa8, 0x2f, 0xad,
	0x28, 0xa5, 0x9a, 0x33, 0xe3, 0xd5, 0x7b, 0xc5, 0xb6, 0xb4, 0xe2, 0xa5, 0xab, 0x71, 0x8a, 0x86,
	0x0e, 0xac, 0x1a, 0x5b, 0x5a, 0x47, 0xf5, 0x82, 0xbd, 0x62, 0x20, 0xad, 0x98, 0x35, 0xd6, 0x4f,
	0x4b, 0x17, 0x11, 0xda, 0x7d, 0x41, 0xbb, 0x8b, 0x15, 0x9e, 0xa2, 0x6b, 0x82, 0x76, 0x41, 0xa1,
	0xa4, 0xb0, 0xc1, 0x8a, 0x48, 0xd0, 0xce, 0x1b, 0x98, 0x82, 0xf3, 0xb7, 0x66, 0x2f, 0xe5, 0xb7,
	0x57, 0x30, 0x9e, 0xa1, 0xeb, 0x0e, 0x5c, 0xfb, 0x73, 0xe3, 0x7b, 0x97, 0x8e, 0x1f, 0x0a, 0xda,
	0xad, 0xec, 0x4f, 0xc1, 0x2d, 0x79, 0x66, 0xc8, 0x68, 0x2b, 0x2c, 0xf9, 0x57, 0xd7, 0xe4, 0x6b,
	0x8c, 0x92, 0xe9, 0x99, 0xa4, 0xf8, 0x1d, 0x0f, 0x43, 0x4e, 0xf0, 0x47, 0xb4, 0x73, 0xc0, 0x60,
	0x66, 0xb5, 0x66, 0x12, 0xce, 0xbd, 0x04, 0xee, 0x93, 0x36, 0x27, 0xcf, 0x5d, 0x72, 0x92, 0x5b,
	0xee, 0xe7, 0x85, 0x6f, 0x95, 0x3e, 0xfa, 0xf4, 0xe3, 0xf7, 0xe7, 0xf8, 0x01, 0x26, 0xe7, 0xa3,
	0x99, 0xf9, 0x68, 0x66, 0x26, 0x70, 0xef, 0x7b, 0xab, 0x59, 0x15, 0x24, 0xf1, 0x07, 0x74, 0xe3,
	0x80, 0xc1, 0xbf, 0xce, 0x80, 0x87, 0x4e, 0x71, 0x1d, 0xb4, 0x64, 0xcf, 0x95, 0xff, 0xbb, 0x57,
	0x3a, 0xf1, 0x1e, 0xee, 0xe1, 0xfd, 0x0d, 0x3d, 0x08, 0xda, 0x3d, 0x7b, 0xfc, 0x6d, 0x31, 0x8e,
	0xbe, 0x2f, 0xc6, 0xd1, 0xcf, 0xc5, 0x38, 0xfa, 0xf2, 0x6b, 0x7c, 0x05, 0x8d, 0xb8, 0x22, 0x06,
	0x68, 0xf5, 0x56, 0xab, 0x2e, 0x1c, 0x81, 0xd0, 0x86, 0x93, 0x36, 0x7f, 0x3d, 0x20, 0x59, 0xf8,
	0x98, 0x9e, 0xb4, 0xf9, 0xab, 0xe8, 0xe4, 0xaa, 0x07, 0x1f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0x9a, 0xbe, 0xbc, 0x91, 0x03, 0x00, 0x00,
}

func (m *TimeRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.To != nil {
		{
			size, err := m.To.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.From != nil {
		{
			size, err := m.From.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SecuredUnitsUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecuredUnitsUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecuredUnitsUsageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NumCpuUnits != 0 {
		i = encodeVarintAdministrationUsageService(dAtA, i, uint64(m.NumCpuUnits))
		i--
		dAtA[i] = 0x10
	}
	if m.NumNodes != 0 {
		i = encodeVarintAdministrationUsageService(dAtA, i, uint64(m.NumNodes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MaxSecuredUnitsUsageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MaxSecuredUnitsUsageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MaxSecuredUnitsUsageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxCpuUnits != 0 {
		i = encodeVarintAdministrationUsageService(dAtA, i, uint64(m.MaxCpuUnits))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxCpuUnitsAt != nil {
		{
			size, err := m.MaxCpuUnitsAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxNodes != 0 {
		i = encodeVarintAdministrationUsageService(dAtA, i, uint64(m.MaxNodes))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNodesAt != nil {
		{
			size, err := m.MaxNodesAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAdministrationUsageService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdministrationUsageService(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdministrationUsageService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimeRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovAdministrationUsageService(uint64(l))
	}
	if m.To != nil {
		l = m.To.Size()
		n += 1 + l + sovAdministrationUsageService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SecuredUnitsUsageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumNodes != 0 {
		n += 1 + sovAdministrationUsageService(uint64(m.NumNodes))
	}
	if m.NumCpuUnits != 0 {
		n += 1 + sovAdministrationUsageService(uint64(m.NumCpuUnits))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MaxSecuredUnitsUsageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNodesAt != nil {
		l = m.MaxNodesAt.Size()
		n += 1 + l + sovAdministrationUsageService(uint64(l))
	}
	if m.MaxNodes != 0 {
		n += 1 + sovAdministrationUsageService(uint64(m.MaxNodes))
	}
	if m.MaxCpuUnitsAt != nil {
		l = m.MaxCpuUnitsAt.Size()
		n += 1 + l + sovAdministrationUsageService(uint64(l))
	}
	if m.MaxCpuUnits != 0 {
		n += 1 + sovAdministrationUsageService(uint64(m.MaxCpuUnits))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdministrationUsageService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdministrationUsageService(x uint64) (n int) {
	return sovAdministrationUsageService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdministrationUsageService
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
			return fmt.Errorf("proto: TimeRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
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
				return ErrInvalidLengthAdministrationUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &types.Timestamp{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
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
				return ErrInvalidLengthAdministrationUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.To == nil {
				m.To = &types.Timestamp{}
			}
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdministrationUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SecuredUnitsUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdministrationUsageService
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
			return fmt.Errorf("proto: SecuredUnitsUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecuredUnitsUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumNodes", wireType)
			}
			m.NumNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumNodes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumCpuUnits", wireType)
			}
			m.NumCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumCpuUnits |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAdministrationUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MaxSecuredUnitsUsageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdministrationUsageService
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
			return fmt.Errorf("proto: MaxSecuredUnitsUsageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MaxSecuredUnitsUsageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodesAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
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
				return ErrInvalidLengthAdministrationUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxNodesAt == nil {
				m.MaxNodesAt = &types.Timestamp{}
			}
			if err := m.MaxNodesAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNodes", wireType)
			}
			m.MaxNodes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNodes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnitsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
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
				return ErrInvalidLengthAdministrationUsageService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxCpuUnitsAt == nil {
				m.MaxCpuUnitsAt = &types.Timestamp{}
			}
			if err := m.MaxCpuUnitsAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCpuUnits", wireType)
			}
			m.MaxCpuUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdministrationUsageService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCpuUnits |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAdministrationUsageService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdministrationUsageService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAdministrationUsageService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdministrationUsageService
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
					return 0, ErrIntOverflowAdministrationUsageService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdministrationUsageService
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
				return 0, ErrInvalidLengthAdministrationUsageService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdministrationUsageService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdministrationUsageService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdministrationUsageService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdministrationUsageService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdministrationUsageService = fmt.Errorf("proto: unexpected end of group")
)
