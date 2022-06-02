// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/provider/v1/provider.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	types1 "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// CreateConsumerChainProposal is a governance proposal on the provider chain to spawn a new consumer chain.
// If it passes, then all validators on the provider chain are expected to validate the consumer chain at spawn time
// or get slashed. It is recommended that spawn time occurs after the proposal end time.
type CreateConsumerChainProposal struct {
	// the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// the proposed chain-id of the new consumer chain, must be different from all other consumer chain ids of the executing
	// provider chain.
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// the proposed initial height of new consumer chain.
	// For a completely new chain, this will be {0,1}. However, it may be different if this is a chain that is converting to a consumer chain.
	InitialHeight *types.Height `protobuf:"bytes,4,opt,name=initial_height,json=initialHeight,proto3" json:"initial_height,omitempty"`
	// genesis hash with no staking information included.
	GenesisHash []byte `protobuf:"bytes,5,opt,name=genesis_hash,json=genesisHash,proto3" json:"genesis_hash,omitempty"`
	// binary hash is the hash of the binary that should be used by validators on chain initialization.
	BinaryHash []byte `protobuf:"bytes,6,opt,name=binary_hash,json=binaryHash,proto3" json:"binary_hash,omitempty"`
	// spawn time is the time on the provider chain at which the consumer chain genesis is finalized and all validators
	// will be responsible for starting their consumer chain validator node.
	SpawnTime time.Time `protobuf:"bytes,7,opt,name=spawn_time,json=spawnTime,proto3,stdtime" json:"spawn_time"`
}

func (m *CreateConsumerChainProposal) Reset()      { *m = CreateConsumerChainProposal{} }
func (*CreateConsumerChainProposal) ProtoMessage() {}
func (*CreateConsumerChainProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22ec409a72b7b72, []int{0}
}
func (m *CreateConsumerChainProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateConsumerChainProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateConsumerChainProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateConsumerChainProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsumerChainProposal.Merge(m, src)
}
func (m *CreateConsumerChainProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateConsumerChainProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsumerChainProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsumerChainProposal proto.InternalMessageInfo

// StopConsumerProposal is a governance proposal on the provider chain to stop a consumer chain.
// If it passes, all the consumer chain's states are removed from the provider chain. The outstanding unbonding
// operation funds are released if the LockUnbondingOnTimeout parameter is set to false for the consumer chain ID.
type StopConsumerChainProposal struct {
	// the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// the chain-id of the consumer chain to be stopped
	ChainId string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// the time on the provider chain at which all validators are responsible to stop their consumer chain validator node
	StopTime time.Time `protobuf:"bytes,4,opt,name=stop_time,json=stopTime,proto3,stdtime" json:"stop_time"`
}

func (m *StopConsumerChainProposal) Reset()         { *m = StopConsumerChainProposal{} }
func (m *StopConsumerChainProposal) String() string { return proto.CompactTextString(m) }
func (*StopConsumerChainProposal) ProtoMessage()    {}
func (*StopConsumerChainProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22ec409a72b7b72, []int{1}
}
func (m *StopConsumerChainProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StopConsumerChainProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StopConsumerChainProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StopConsumerChainProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopConsumerChainProposal.Merge(m, src)
}
func (m *StopConsumerChainProposal) XXX_Size() int {
	return m.Size()
}
func (m *StopConsumerChainProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_StopConsumerChainProposal.DiscardUnknown(m)
}

var xxx_messageInfo_StopConsumerChainProposal proto.InternalMessageInfo

func (m *StopConsumerChainProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StopConsumerChainProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StopConsumerChainProposal) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *StopConsumerChainProposal) GetStopTime() time.Time {
	if m != nil {
		return m.StopTime
	}
	return time.Time{}
}

// Params defines the parameters for CCV Provider module
type Params struct {
	TemplateClient *types1.ClientState `protobuf:"bytes,1,opt,name=template_client,json=templateClient,proto3" json:"template_client,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22ec409a72b7b72, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTemplateClient() *types1.ClientState {
	if m != nil {
		return m.TemplateClient
	}
	return nil
}

type HandshakeMetadata struct {
	ProviderFeePoolAddr string `protobuf:"bytes,1,opt,name=provider_fee_pool_addr,json=providerFeePoolAddr,proto3" json:"provider_fee_pool_addr,omitempty"`
	Version             string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *HandshakeMetadata) Reset()         { *m = HandshakeMetadata{} }
func (m *HandshakeMetadata) String() string { return proto.CompactTextString(m) }
func (*HandshakeMetadata) ProtoMessage()    {}
func (*HandshakeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f22ec409a72b7b72, []int{3}
}
func (m *HandshakeMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakeMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeMetadata.Merge(m, src)
}
func (m *HandshakeMetadata) XXX_Size() int {
	return m.Size()
}
func (m *HandshakeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeMetadata proto.InternalMessageInfo

func (m *HandshakeMetadata) GetProviderFeePoolAddr() string {
	if m != nil {
		return m.ProviderFeePoolAddr
	}
	return ""
}

func (m *HandshakeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateConsumerChainProposal)(nil), "interchain_security.ccv.provider.v1.CreateConsumerChainProposal")
	proto.RegisterType((*StopConsumerChainProposal)(nil), "interchain_security.ccv.provider.v1.StopConsumerChainProposal")
	proto.RegisterType((*Params)(nil), "interchain_security.ccv.provider.v1.Params")
	proto.RegisterType((*HandshakeMetadata)(nil), "interchain_security.ccv.provider.v1.HandshakeMetadata")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/provider/v1/provider.proto", fileDescriptor_f22ec409a72b7b72)
}

var fileDescriptor_f22ec409a72b7b72 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xbf, 0x6f, 0xd4, 0x30,
	0x14, 0x4e, 0x4a, 0x7f, 0x5c, 0x7d, 0xa5, 0x88, 0x50, 0xa1, 0xb4, 0x48, 0xc9, 0x71, 0x2c, 0x95,
	0x10, 0xb6, 0xae, 0xdd, 0xba, 0xb5, 0x27, 0xa1, 0x32, 0x20, 0x55, 0xd7, 0x4e, 0x0c, 0x44, 0x8e,
	0xf3, 0x9a, 0x58, 0x24, 0x71, 0x64, 0xfb, 0x02, 0xfd, 0x0f, 0x18, 0x3b, 0x32, 0xf6, 0x5f, 0xe0,
	0x9f, 0x40, 0x1d, 0x3b, 0x32, 0x01, 0xea, 0xfd, 0x23, 0x28, 0x76, 0x72, 0x77, 0x48, 0x2c, 0x2c,
	0x6c, 0x7e, 0xdf, 0xfb, 0xbe, 0xe4, 0xf3, 0xf7, 0xfc, 0xd0, 0x01, 0x2f, 0x35, 0x48, 0x96, 0x51,
	0x5e, 0x46, 0x0a, 0xd8, 0x54, 0x72, 0x7d, 0x45, 0x18, 0xab, 0x49, 0x25, 0x45, 0xcd, 0x13, 0x90,
	0xa4, 0x1e, 0xcd, 0xcf, 0xb8, 0x92, 0x42, 0x0b, 0xef, 0xc5, 0x5f, 0x34, 0x98, 0xb1, 0x1a, 0xcf,
	0x79, 0xf5, 0x68, 0x6f, 0x27, 0x15, 0xa9, 0x30, 0x7c, 0xd2, 0x9c, 0xac, 0x74, 0x2f, 0x4c, 0x85,
	0x48, 0x73, 0x20, 0xa6, 0x8a, 0xa7, 0x97, 0x44, 0xf3, 0x02, 0x94, 0xa6, 0x45, 0xd5, 0x11, 0x78,
	0xcc, 0x08, 0x13, 0x12, 0x08, 0xcb, 0x39, 0x94, 0xba, 0xf9, 0xbd, 0x3d, 0xb5, 0x04, 0xd2, 0x10,
	0x72, 0x9e, 0x66, 0xda, 0xc2, 0x8a, 0x68, 0x28, 0x13, 0x90, 0x05, 0xb7, 0xe4, 0x45, 0x65, 0x05,
	0xc3, 0x6f, 0x2b, 0xe8, 0xd9, 0x58, 0x02, 0xd5, 0x30, 0x16, 0xa5, 0x9a, 0x16, 0x20, 0xc7, 0x8d,
	0xf3, 0x33, 0x29, 0x2a, 0xa1, 0x68, 0xee, 0xed, 0xa0, 0x35, 0xcd, 0x75, 0x0e, 0xbe, 0x3b, 0x70,
	0xf7, 0x37, 0x27, 0xb6, 0xf0, 0x06, 0xa8, 0x9f, 0x80, 0x62, 0x92, 0x57, 0x9a, 0x8b, 0xd2, 0x5f,
	0x31, 0xbd, 0x65, 0xc8, 0xdb, 0x45, 0x3d, 0x1b, 0x01, 0x4f, 0xfc, 0x07, 0xa6, 0xbd, 0x61, 0xea,
	0x37, 0x89, 0x77, 0x8c, 0xb6, 0x79, 0xc9, 0x35, 0xa7, 0x79, 0x94, 0x41, 0x63, 0xd5, 0x5f, 0x1d,
	0xb8, 0xfb, 0xfd, 0x83, 0x3d, 0xcc, 0x63, 0x86, 0x9b, 0xdb, 0xe1, 0xf6, 0x4e, 0xf5, 0x08, 0x9f,
	0x1a, 0xc6, 0xe4, 0x61, 0xab, 0xb0, 0xa5, 0xf7, 0x1c, 0x6d, 0xa5, 0x50, 0x82, 0xe2, 0x2a, 0xca,
	0xa8, 0xca, 0xfc, 0xb5, 0x81, 0xbb, 0xbf, 0x35, 0xe9, 0xb7, 0xd8, 0x29, 0x55, 0x99, 0x17, 0xa2,
	0x7e, 0xcc, 0x4b, 0x2a, 0xaf, 0x2c, 0x63, 0xdd, 0x30, 0x90, 0x85, 0x0c, 0x61, 0x8c, 0x90, 0xaa,
	0xe8, 0xc7, 0x32, 0x6a, 0x42, 0xf6, 0x37, 0x5a, 0x0b, 0x76, 0x02, 0xb8, 0x9b, 0x00, 0xbe, 0xe8,
	0x26, 0x70, 0xd2, 0xbb, 0xfd, 0x11, 0x3a, 0xd7, 0x3f, 0x43, 0x77, 0xb2, 0x69, 0x74, 0x4d, 0xe7,
	0xa8, 0xf7, 0xf9, 0x26, 0x74, 0xbe, 0xdc, 0x84, 0xce, 0xf0, 0xab, 0x8b, 0x76, 0xcf, 0xb5, 0xa8,
	0xfe, 0x63, 0x8c, 0x9b, 0x4a, 0x8b, 0xca, 0xda, 0x5f, 0xfd, 0x07, 0xfb, 0xbd, 0x46, 0xd6, 0x34,
	0x86, 0xef, 0xd1, 0xfa, 0x19, 0x95, 0xb4, 0x50, 0xde, 0x05, 0x7a, 0xa4, 0xa1, 0xa8, 0x72, 0xaa,
	0x21, 0xb2, 0xe1, 0x1b, 0xa7, 0xfd, 0x83, 0x97, 0x66, 0x28, 0xcb, 0x2f, 0x0a, 0x2f, 0xbd, 0xa1,
	0x7a, 0x84, 0xc7, 0x06, 0x3d, 0xd7, 0x54, 0xc3, 0x64, 0xbb, 0xfb, 0x86, 0x05, 0x87, 0x31, 0x7a,
	0x7c, 0x4a, 0xcb, 0x44, 0x65, 0xf4, 0x03, 0xbc, 0x05, 0x4d, 0x13, 0xaa, 0xa9, 0x77, 0x88, 0x9e,
	0x76, 0x9b, 0x10, 0x5d, 0x02, 0x44, 0x95, 0x10, 0x79, 0x44, 0x93, 0x44, 0xb6, 0xd9, 0x3c, 0xe9,
	0xba, 0xaf, 0x01, 0xce, 0x84, 0xc8, 0x8f, 0x93, 0x44, 0x7a, 0x3e, 0xda, 0xa8, 0x41, 0xaa, 0x45,
	0x4a, 0x5d, 0x79, 0x72, 0x71, 0x7b, 0x1f, 0xb8, 0x77, 0xf7, 0x81, 0xfb, 0xeb, 0x3e, 0x70, 0xaf,
	0x67, 0x81, 0x73, 0x37, 0x0b, 0x9c, 0xef, 0xb3, 0xc0, 0x79, 0x77, 0x94, 0x72, 0x9d, 0x4d, 0x63,
	0xcc, 0x44, 0x41, 0x98, 0x50, 0x85, 0x50, 0x64, 0xb1, 0x9a, 0xaf, 0xe6, 0xeb, 0xfc, 0xe9, 0xcf,
	0x85, 0xd6, 0x57, 0x15, 0xa8, 0x78, 0xdd, 0x24, 0x78, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x3e, 0x38, 0x24, 0x01, 0x04, 0x00, 0x00,
}

func (m *CreateConsumerChainProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateConsumerChainProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateConsumerChainProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SpawnTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SpawnTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProvider(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if len(m.BinaryHash) > 0 {
		i -= len(m.BinaryHash)
		copy(dAtA[i:], m.BinaryHash)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.BinaryHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.GenesisHash) > 0 {
		i -= len(m.GenesisHash)
		copy(dAtA[i:], m.GenesisHash)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.GenesisHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.InitialHeight != nil {
		{
			size, err := m.InitialHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProvider(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StopConsumerChainProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StopConsumerChainProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StopConsumerChainProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StopTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StopTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintProvider(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TemplateClient != nil {
		{
			size, err := m.TemplateClient.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProvider(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HandshakeMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakeMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ProviderFeePoolAddr) > 0 {
		i -= len(m.ProviderFeePoolAddr)
		copy(dAtA[i:], m.ProviderFeePoolAddr)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.ProviderFeePoolAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateConsumerChainProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if m.InitialHeight != nil {
		l = m.InitialHeight.Size()
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.GenesisHash)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.BinaryHash)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SpawnTime)
	n += 1 + l + sovProvider(uint64(l))
	return n
}

func (m *StopConsumerChainProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StopTime)
	n += 1 + l + sovProvider(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TemplateClient != nil {
		l = m.TemplateClient.Size()
		n += 1 + l + sovProvider(uint64(l))
	}
	return n
}

func (m *HandshakeMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProviderFeePoolAddr)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	return n
}

func sovProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvider(x uint64) (n int) {
	return sovProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateConsumerChainProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: CreateConsumerChainProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateConsumerChainProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InitialHeight == nil {
				m.InitialHeight = &types.Height{}
			}
			if err := m.InitialHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisHash = append(m.GenesisHash[:0], dAtA[iNdEx:postIndex]...)
			if m.GenesisHash == nil {
				m.GenesisHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BinaryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BinaryHash = append(m.BinaryHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BinaryHash == nil {
				m.BinaryHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpawnTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SpawnTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *StopConsumerChainProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: StopConsumerChainProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StopConsumerChainProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StopTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StopTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemplateClient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TemplateClient == nil {
				m.TemplateClient = &types1.ClientState{}
			}
			if err := m.TemplateClient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *HandshakeMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: HandshakeMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderFeePoolAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderFeePoolAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func skipProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
				return 0, ErrInvalidLengthProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvider = fmt.Errorf("proto: unexpected end of group")
)
