// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: celestia/qgb/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BridgeValidator represents a validator's ETH address and its power
type BridgeValidator struct {
	// Voting power of the validator.
	Power uint64 `protobuf:"varint,1,opt,name=power,proto3" json:"power,omitempty"`
	// EVM address that will be used by the validator to sign messages.
	EvmAddress string `protobuf:"bytes,2,opt,name=evm_address,json=evmAddress,proto3" json:"evm_address,omitempty"`
}

func (m *BridgeValidator) Reset()         { *m = BridgeValidator{} }
func (m *BridgeValidator) String() string { return proto.CompactTextString(m) }
func (*BridgeValidator) ProtoMessage()    {}
func (*BridgeValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db0e6d49b998544, []int{0}
}
func (m *BridgeValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BridgeValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BridgeValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BridgeValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeValidator.Merge(m, src)
}
func (m *BridgeValidator) XXX_Size() int {
	return m.Size()
}
func (m *BridgeValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeValidator.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeValidator proto.InternalMessageInfo

func (m *BridgeValidator) GetPower() uint64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *BridgeValidator) GetEvmAddress() string {
	if m != nil {
		return m.EvmAddress
	}
	return ""
}

// Valset is the EVM Bridge Multsig Set, each qgb validator also
// maintains an ETH key to sign messages, these are used to check signatures on
// ETH because of the significant gas savings
type Valset struct {
	// Universal nonce defined under:
	// https://github.com/celestiaorg/celestia-app/pull/464
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// List of BridgeValidator containing the current validator set.
	Members []BridgeValidator `protobuf:"bytes,2,rep,name=members,proto3" json:"members"`
	// Current chain height
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Valset) Reset()         { *m = Valset{} }
func (m *Valset) String() string { return proto.CompactTextString(m) }
func (*Valset) ProtoMessage()    {}
func (*Valset) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db0e6d49b998544, []int{1}
}
func (m *Valset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Valset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Valset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Valset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Valset.Merge(m, src)
}
func (m *Valset) XXX_Size() int {
	return m.Size()
}
func (m *Valset) XXX_DiscardUnknown() {
	xxx_messageInfo_Valset.DiscardUnknown(m)
}

var xxx_messageInfo_Valset proto.InternalMessageInfo

func (m *Valset) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Valset) GetMembers() []BridgeValidator {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Valset) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// DataCommitment is the data commitment request message that will be signed
// using orchestrators.
// It does not contain a `commitment` field as this message will be created
// inside the state machine and it doesn't make sense to ask tendermint for the
// commitment there.
type DataCommitment struct {
	// Universal nonce defined under:
	// https://github.com/celestiaorg/celestia-app/pull/464
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// First block defining the ordered set of blocks used to create the
	// commitment.
	BeginBlock uint64 `protobuf:"varint,2,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
	// Last block defining the ordered set of blocks used to create the
	// commitment.
	EndBlock uint64 `protobuf:"varint,3,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (m *DataCommitment) Reset()         { *m = DataCommitment{} }
func (m *DataCommitment) String() string { return proto.CompactTextString(m) }
func (*DataCommitment) ProtoMessage()    {}
func (*DataCommitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db0e6d49b998544, []int{2}
}
func (m *DataCommitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataCommitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataCommitment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataCommitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCommitment.Merge(m, src)
}
func (m *DataCommitment) XXX_Size() int {
	return m.Size()
}
func (m *DataCommitment) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCommitment.DiscardUnknown(m)
}

var xxx_messageInfo_DataCommitment proto.InternalMessageInfo

func (m *DataCommitment) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *DataCommitment) GetBeginBlock() uint64 {
	if m != nil {
		return m.BeginBlock
	}
	return 0
}

func (m *DataCommitment) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*BridgeValidator)(nil), "celestia.qgb.v1.BridgeValidator")
	proto.RegisterType((*Valset)(nil), "celestia.qgb.v1.Valset")
	proto.RegisterType((*DataCommitment)(nil), "celestia.qgb.v1.DataCommitment")
}

func init() { proto.RegisterFile("celestia/qgb/v1/types.proto", fileDescriptor_5db0e6d49b998544) }

var fileDescriptor_5db0e6d49b998544 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0x86, 0xed, 0x83, 0xe3, 0x8e, 0x45, 0x3a, 0x24, 0x1f, 0x4a, 0x08, 0x48, 0x06, 0x51, 0xd1,
	0x60, 0x8b, 0xa4, 0x4b, 0x15, 0x9c, 0x14, 0xa1, 0x75, 0x41, 0x91, 0x06, 0xad, 0xed, 0xd1, 0xb2,
	0x8a, 0xd7, 0x6b, 0xbc, 0x8b, 0x93, 0x14, 0x79, 0x87, 0xe4, 0x5d, 0xf2, 0x10, 0x94, 0x28, 0x55,
	0xaa, 0x28, 0x82, 0x17, 0x89, 0xbc, 0x8b, 0x53, 0x20, 0xd1, 0xcd, 0x3f, 0x9f, 0x67, 0xe6, 0xf7,
	0xfe, 0xa8, 0x1b, 0x42, 0x0c, 0x42, 0x52, 0xec, 0x2e, 0x49, 0xe0, 0xe6, 0x63, 0x57, 0x3e, 0xa5,
	0x20, 0x9c, 0x34, 0xe3, 0x92, 0x5b, 0xcd, 0x12, 0x3a, 0x4b, 0x12, 0x38, 0xf9, 0xb8, 0xd3, 0x22,
	0x9c, 0x70, 0xc5, 0xdc, 0xa2, 0xd2, 0x9f, 0x75, 0xce, 0x42, 0x2e, 0x18, 0x17, 0x73, 0x0d, 0xb4,
	0xd0, 0x68, 0x70, 0x8b, 0x9a, 0x5e, 0x46, 0x23, 0x02, 0x33, 0x1c, 0xd3, 0x08, 0x4b, 0x9e, 0x59,
	0x2d, 0xf4, 0x3b, 0xe5, 0x0f, 0x90, 0xb5, 0xcd, 0xbe, 0x39, 0xac, 0xfa, 0x5a, 0x58, 0x3d, 0xd4,
	0x80, 0x9c, 0xcd, 0x71, 0x14, 0x65, 0x20, 0x44, 0xfb, 0x57, 0xdf, 0x1c, 0xd6, 0x7d, 0x04, 0x39,
	0x9b, 0xe8, 0xce, 0xe0, 0xd5, 0x44, 0xb5, 0x19, 0x8e, 0x05, 0xc8, 0x62, 0x43, 0xc2, 0x93, 0x10,
	0xca, 0x0d, 0x4a, 0x58, 0x57, 0xe8, 0x0f, 0x03, 0x16, 0x40, 0x56, 0x4c, 0x57, 0x86, 0x8d, 0xf3,
	0xbe, 0x73, 0x60, 0xdf, 0x39, 0xb0, 0xe2, 0x55, 0xd7, 0x9f, 0x3d, 0xc3, 0x2f, 0xc7, 0xac, 0x13,
	0x54, 0x5b, 0x00, 0x25, 0x0b, 0xd9, 0xae, 0xa8, 0xc5, 0x7b, 0x75, 0x79, 0xfa, 0xfe, 0x36, 0xfa,
	0x3f, 0x91, 0x12, 0x84, 0xc4, 0x92, 0xf2, 0xc4, 0x87, 0xe5, 0x0a, 0x84, 0x9c, 0x0e, 0x9e, 0xd1,
	0xbf, 0x1b, 0x2c, 0xf1, 0x35, 0x67, 0x8c, 0x4a, 0x06, 0xc9, 0x31, 0x6b, 0x3d, 0xd4, 0x08, 0x80,
	0xd0, 0x64, 0x1e, 0xc4, 0x3c, 0xbc, 0x57, 0x3f, 0x57, 0xf5, 0x91, 0x6a, 0x79, 0x45, 0xc7, 0xea,
	0xa2, 0x3a, 0x24, 0xd1, 0x1e, 0xeb, 0xe3, 0x7f, 0x21, 0x89, 0x14, 0x3c, 0x7a, 0xde, 0x9b, 0xae,
	0xb7, 0xb6, 0xb9, 0xd9, 0xda, 0xe6, 0xd7, 0xd6, 0x36, 0x5f, 0x76, 0xb6, 0xb1, 0xd9, 0xd9, 0xc6,
	0xc7, 0xce, 0x36, 0xee, 0x5c, 0x42, 0xe5, 0x62, 0x15, 0x38, 0x21, 0x67, 0x6e, 0xf9, 0x08, 0x3c,
	0x23, 0x3f, 0xf5, 0x08, 0xa7, 0xa9, 0xfb, 0xa8, 0x22, 0x57, 0x79, 0x07, 0x35, 0x15, 0xd7, 0xc5,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x1e, 0x84, 0xc3, 0x0f, 0x02, 0x00, 0x00,
}

func (m *BridgeValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BridgeValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BridgeValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EvmAddress) > 0 {
		i -= len(m.EvmAddress)
		copy(dAtA[i:], m.EvmAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EvmAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Valset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Valset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Valset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DataCommitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataCommitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataCommitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlock != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.BeginBlock != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BeginBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BridgeValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	l = len(m.EvmAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Valset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	return n
}

func (m *DataCommitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	if m.BeginBlock != 0 {
		n += 1 + sovTypes(uint64(m.BeginBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovTypes(uint64(m.EndBlock))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BridgeValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BridgeValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BridgeValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvmAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Valset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Valset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Valset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, BridgeValidator{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *DataCommitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: DataCommitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataCommitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			m.BeginBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
