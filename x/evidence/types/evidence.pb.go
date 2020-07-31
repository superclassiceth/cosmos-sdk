// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/evidence/beta/evidence.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSubmitEvidence defines an sdk.Msg type that supports submitting arbitrary
// Evidence.
type MsgSubmitEvidence struct {
	Submitter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=submitter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"submitter,omitempty"`
	Evidence  *types.Any                                    `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
}

func (m *MsgSubmitEvidence) Reset()         { *m = MsgSubmitEvidence{} }
func (m *MsgSubmitEvidence) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEvidence) ProtoMessage()    {}
func (*MsgSubmitEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_60e6edd1b2dc8540, []int{0}
}
func (m *MsgSubmitEvidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEvidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEvidence.Merge(m, src)
}
func (m *MsgSubmitEvidence) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEvidence proto.InternalMessageInfo

// Equivocation implements the Evidence interface and defines evidence of double
// signing misbehavior.
type Equivocation struct {
	Height           int64                                          `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time             time.Time                                      `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	Power            int64                                          `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	ConsensusAddress github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,4,opt,name=consensus_address,json=consensusAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"consensus_address,omitempty" yaml:"consensus_address"`
}

func (m *Equivocation) Reset()      { *m = Equivocation{} }
func (*Equivocation) ProtoMessage() {}
func (*Equivocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_60e6edd1b2dc8540, []int{1}
}
func (m *Equivocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Equivocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Equivocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Equivocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equivocation.Merge(m, src)
}
func (m *Equivocation) XXX_Size() int {
	return m.Size()
}
func (m *Equivocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Equivocation.DiscardUnknown(m)
}

var xxx_messageInfo_Equivocation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitEvidence)(nil), "cosmos.evidence.beta.MsgSubmitEvidence")
	proto.RegisterType((*Equivocation)(nil), "cosmos.evidence.beta.Equivocation")
}

func init() {
	proto.RegisterFile("cosmos/evidence/beta/evidence.proto", fileDescriptor_60e6edd1b2dc8540)
}

var fileDescriptor_60e6edd1b2dc8540 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbb, 0x8e, 0xd3, 0x40,
	0x18, 0x85, 0x3d, 0x24, 0x44, 0x61, 0x48, 0x41, 0x2c, 0x0b, 0x99, 0x14, 0x9e, 0xc8, 0x34, 0x69,
	0x32, 0xe6, 0xd2, 0xa0, 0x74, 0x09, 0x4a, 0x85, 0x00, 0xc9, 0x50, 0xd1, 0x20, 0x5f, 0x06, 0xc7,
	0x22, 0xf6, 0x18, 0xcf, 0x38, 0x60, 0xf1, 0x02, 0x94, 0x29, 0x29, 0xb6, 0x48, 0xb9, 0x8f, 0x92,
	0x32, 0xe5, 0x56, 0xde, 0x95, 0xf3, 0x06, 0x5b, 0x46, 0x5a, 0x69, 0xe5, 0xf1, 0x25, 0xd2, 0x46,
	0x5a, 0x6d, 0x65, 0x9f, 0xdf, 0x9f, 0xcf, 0xfc, 0xe7, 0xd8, 0xf0, 0xa5, 0x43, 0x59, 0x40, 0x99,
	0x41, 0x56, 0xbe, 0x4b, 0x42, 0x87, 0x18, 0x36, 0xe1, 0x56, 0xa3, 0x70, 0x14, 0x53, 0x4e, 0x65,
	0xa5, 0x84, 0x70, 0x33, 0x2e, 0xa0, 0x81, 0xe2, 0x51, 0x8f, 0x0a, 0xc0, 0x28, 0xee, 0x4a, 0x76,
	0x80, 0x3c, 0x4a, 0xbd, 0x25, 0x31, 0x84, 0xb2, 0x93, 0x1f, 0x06, 0xf7, 0x03, 0xc2, 0xb8, 0x15,
	0x44, 0x15, 0xf0, 0xe2, 0x2e, 0x60, 0x85, 0x69, 0xf9, 0x48, 0x3f, 0x03, 0xb0, 0xff, 0x91, 0x79,
	0x5f, 0x12, 0x3b, 0xf0, 0xf9, 0xbc, 0x3a, 0x4c, 0xfe, 0x0c, 0x9f, 0x30, 0x31, 0xe1, 0x24, 0x56,
	0xc1, 0x10, 0x8c, 0x7a, 0xb3, 0xd7, 0x87, 0x0c, 0x8d, 0x3d, 0x9f, 0x2f, 0x12, 0x1b, 0x3b, 0x34,
	0x30, 0xaa, 0x10, 0xe5, 0x65, 0xcc, 0xdc, 0x9f, 0x06, 0x4f, 0x23, 0xc2, 0xf0, 0xd4, 0x71, 0xa6,
	0xae, 0x1b, 0x13, 0xc6, 0xcc, 0xa3, 0x87, 0xfc, 0x0a, 0x76, 0xeb, 0x24, 0xea, 0xa3, 0x21, 0x18,
	0x3d, 0x7d, 0xa3, 0xe0, 0x72, 0x29, 0x5c, 0x2f, 0x85, 0xa7, 0x61, 0x6a, 0x36, 0xd4, 0xa4, 0xfd,
	0x6f, 0x83, 0x24, 0xfd, 0x06, 0xc0, 0xde, 0xfc, 0x57, 0xe2, 0xaf, 0xa8, 0x63, 0x71, 0x9f, 0x86,
	0xf2, 0x73, 0xd8, 0x59, 0x10, 0xdf, 0x5b, 0x70, 0xb1, 0x56, 0xcb, 0xac, 0x94, 0xfc, 0x0e, 0xb6,
	0x8b, 0xd4, 0x95, 0xf9, 0xe0, 0xc4, 0xfc, 0x6b, 0x5d, 0xc9, 0xac, 0xbb, 0xcd, 0x90, 0xb4, 0xbe,
	0x44, 0xc0, 0x14, 0x6f, 0xc8, 0x0a, 0x7c, 0x1c, 0xd1, 0xdf, 0x24, 0x56, 0x5b, 0xc2, 0xb0, 0x14,
	0xf2, 0x5f, 0xd8, 0x77, 0x68, 0xc8, 0x48, 0xc8, 0x12, 0xf6, 0xdd, 0x2a, 0x03, 0xa9, 0x6d, 0xd1,
	0xc4, 0xa7, 0xeb, 0x0c, 0xa9, 0xa9, 0x15, 0x2c, 0x27, 0xfa, 0x09, 0xa2, 0x1f, 0x32, 0x84, 0x1f,
	0xd0, 0xd2, 0x7b, 0x1a, 0xb2, 0xba, 0xa6, 0x67, 0x8d, 0x4b, 0x35, 0x99, 0x74, 0x8b, 0xec, 0xff,
	0x37, 0x48, 0x9a, 0x7d, 0x38, 0xcf, 0x35, 0xb0, 0xcd, 0x35, 0xb0, 0xcb, 0x35, 0x70, 0x95, 0x6b,
	0x60, 0xbd, 0xd7, 0xa4, 0xdd, 0x5e, 0x93, 0x2e, 0xf6, 0x9a, 0xf4, 0xed, 0xfe, 0xef, 0xf1, 0xe7,
	0xf8, 0x87, 0x89, 0x43, 0xed, 0x8e, 0x68, 0xe3, 0xed, 0x6d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50,
	0x0b, 0x00, 0x18, 0x81, 0x02, 0x00, 0x00,
}

func (this *MsgSubmitEvidence) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgSubmitEvidence)
	if !ok {
		that2, ok := that.(MsgSubmitEvidence)
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
	if !bytes.Equal(this.Submitter, that1.Submitter) {
		return false
	}
	if !this.Evidence.Equal(that1.Evidence) {
		return false
	}
	return true
}
func (this *Equivocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Equivocation)
	if !ok {
		that2, ok := that.(Equivocation)
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
	if this.Height != that1.Height {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if !bytes.Equal(this.ConsensusAddress, that1.ConsensusAddress) {
		return false
	}
	return true
}
func (m *MsgSubmitEvidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintEvidence(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Equivocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Equivocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Equivocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintEvidence(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.Power != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEvidence(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvidence(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvidence(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func (m *Equivocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovEvidence(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovEvidence(uint64(l))
	if m.Power != 0 {
		n += 1 + sovEvidence(uint64(m.Power))
	}
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func sovEvidence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvidence(x uint64) (n int) {
	return sovEvidence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitEvidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: MsgSubmitEvidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = append(m.Submitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Submitter == nil {
				m.Submitter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &types.Any{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *Equivocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: Equivocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Equivocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusAddress = append(m.ConsensusAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusAddress == nil {
				m.ConsensusAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func skipEvidence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
				return 0, ErrInvalidLengthEvidence
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvidence
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvidence
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvidence        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvidence          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvidence = fmt.Errorf("proto: unexpected end of group")
)
