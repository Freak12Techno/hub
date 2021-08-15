// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sentinel-official/hub/types"
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

type EventModule struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
}

func (m *EventModule) Reset()         { *m = EventModule{} }
func (m *EventModule) String() string { return proto.CompactTextString(m) }
func (*EventModule) ProtoMessage()    {}
func (*EventModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{0}
}
func (m *EventModule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventModule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventModule.Merge(m, src)
}
func (m *EventModule) XXX_Size() int {
	return m.Size()
}
func (m *EventModule) XXX_DiscardUnknown() {
	xxx_messageInfo_EventModule.DiscardUnknown(m)
}

var xxx_messageInfo_EventModule proto.InternalMessageInfo

type EventAdd struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
}

func (m *EventAdd) Reset()         { *m = EventAdd{} }
func (m *EventAdd) String() string { return proto.CompactTextString(m) }
func (*EventAdd) ProtoMessage()    {}
func (*EventAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{1}
}
func (m *EventAdd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAdd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAdd.Merge(m, src)
}
func (m *EventAdd) XXX_Size() int {
	return m.Size()
}
func (m *EventAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAdd.DiscardUnknown(m)
}

var xxx_messageInfo_EventAdd proto.InternalMessageInfo

type EventSetStatus struct {
	Id       uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Provider string       `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
	Status   types.Status `protobuf:"varint,3,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventSetStatus) Reset()         { *m = EventSetStatus{} }
func (m *EventSetStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetStatus) ProtoMessage()    {}
func (*EventSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{2}
}
func (m *EventSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetStatus.Merge(m, src)
}
func (m *EventSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetStatus proto.InternalMessageInfo

type EventAddNode struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node     string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
}

func (m *EventAddNode) Reset()         { *m = EventAddNode{} }
func (m *EventAddNode) String() string { return proto.CompactTextString(m) }
func (*EventAddNode) ProtoMessage()    {}
func (*EventAddNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{3}
}
func (m *EventAddNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddNode.Merge(m, src)
}
func (m *EventAddNode) XXX_Size() int {
	return m.Size()
}
func (m *EventAddNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddNode proto.InternalMessageInfo

type EventRemoveNode struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Node     string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty" yaml:"node"`
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty" yaml:"provider"`
}

func (m *EventRemoveNode) Reset()         { *m = EventRemoveNode{} }
func (m *EventRemoveNode) String() string { return proto.CompactTextString(m) }
func (*EventRemoveNode) ProtoMessage()    {}
func (*EventRemoveNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed806c731daf270, []int{4}
}
func (m *EventRemoveNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRemoveNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRemoveNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRemoveNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRemoveNode.Merge(m, src)
}
func (m *EventRemoveNode) XXX_Size() int {
	return m.Size()
}
func (m *EventRemoveNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRemoveNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventRemoveNode proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventModule)(nil), "sentinel.plan.v1.EventModule")
	proto.RegisterType((*EventAdd)(nil), "sentinel.plan.v1.EventAdd")
	proto.RegisterType((*EventSetStatus)(nil), "sentinel.plan.v1.EventSetStatus")
	proto.RegisterType((*EventAddNode)(nil), "sentinel.plan.v1.EventAddNode")
	proto.RegisterType((*EventRemoveNode)(nil), "sentinel.plan.v1.EventRemoveNode")
}

func init() { proto.RegisterFile("sentinel/plan/v1/events.proto", fileDescriptor_fed806c731daf270) }

var fileDescriptor_fed806c731daf270 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xb1, 0xae, 0xda, 0x30,
	0x18, 0x85, 0x63, 0x40, 0x08, 0x4c, 0x81, 0x36, 0xed, 0x40, 0x91, 0x70, 0x90, 0xbb, 0xb0, 0x10,
	0x0b, 0xba, 0x75, 0x6b, 0xd4, 0x8e, 0x6d, 0xa5, 0xb0, 0xb1, 0x05, 0x6c, 0xc0, 0x52, 0x12, 0x47,
	0xc4, 0x89, 0xca, 0x5c, 0xa9, 0x73, 0x9f, 0xa2, 0xea, 0xa3, 0x30, 0x32, 0x76, 0x8a, 0xee, 0x0d,
	0x6f, 0x90, 0x27, 0xb8, 0x8a, 0x4d, 0x10, 0xdb, 0xbd, 0xc3, 0x95, 0xee, 0x16, 0xe5, 0x9c, 0xff,
	0xfc, 0x27, 0x5f, 0x7e, 0x38, 0x8a, 0x59, 0x28, 0x79, 0xc8, 0x7c, 0x12, 0xf9, 0x5e, 0x48, 0xd2,
	0x19, 0x61, 0x29, 0x0b, 0x65, 0x6c, 0x47, 0x7b, 0x21, 0x85, 0xf9, 0xba, 0x92, 0xed, 0x52, 0xb6,
	0xd3, 0xd9, 0xf0, 0xdd, 0x56, 0x6c, 0x85, 0x12, 0x49, 0xf9, 0xa4, 0x7d, 0x43, 0x74, 0x8d, 0x91,
	0x87, 0x88, 0xc5, 0x65, 0x4e, 0x2c, 0x3d, 0x99, 0x5c, 0x72, 0xf0, 0x1c, 0x76, 0xbe, 0x96, 0xb9,
	0xdf, 0x04, 0x4d, 0x7c, 0x66, 0x7e, 0x80, 0x8d, 0xd0, 0x0b, 0xd8, 0x00, 0x8c, 0xc1, 0xa4, 0xed,
	0xf4, 0x8b, 0xcc, 0xea, 0x1c, 0xbc, 0xc0, 0xff, 0x84, 0xcb, 0xb7, 0xd8, 0x55, 0x22, 0x5e, 0xc2,
	0x96, 0x9a, 0xf9, 0x4c, 0xa9, 0x39, 0x82, 0x35, 0x4e, 0x95, 0xbd, 0xe1, 0x74, 0x8b, 0xcc, 0x6a,
	0x6b, 0x3b, 0xa7, 0xd8, 0xad, 0x71, 0x6a, 0x12, 0xd8, 0x8a, 0xf6, 0x22, 0xe5, 0x94, 0xed, 0x07,
	0x35, 0x95, 0xf9, 0xb6, 0xc8, 0xac, 0xbe, 0x36, 0x55, 0x0a, 0x76, 0xaf, 0x26, 0xfc, 0x17, 0xc0,
	0x9e, 0x0a, 0x5f, 0x30, 0xb9, 0x50, 0x45, 0x9f, 0x7b, 0x85, 0xf9, 0x05, 0x36, 0x35, 0x82, 0x41,
	0x7d, 0x0c, 0x26, 0xbd, 0xf9, 0x7b, 0xfb, 0xca, 0x52, 0x31, 0xb2, 0xd3, 0x99, 0xad, 0x57, 0x3b,
	0x6f, 0x8a, 0xcc, 0xea, 0xea, 0x24, 0x3d, 0x82, 0xdd, 0xcb, 0x2c, 0xfe, 0x05, 0xe0, 0xab, 0x8a,
	0xc2, 0x77, 0x41, 0xd9, 0x63, 0x35, 0x4b, 0xb2, 0x82, 0xb2, 0x4b, 0xc5, 0x5b, 0xb2, 0x82, 0x2a,
	0xb2, 0x65, 0xc6, 0xed, 0xb7, 0xd4, 0x9f, 0x82, 0xeb, 0x37, 0x80, 0x7d, 0xd5, 0xc2, 0x65, 0x81,
	0x48, 0xd9, 0x8b, 0x15, 0x71, 0x7e, 0x1c, 0xef, 0x91, 0xf1, 0x2f, 0x47, 0xc6, 0x31, 0x47, 0xe0,
	0x94, 0x23, 0x70, 0x97, 0x23, 0xf0, 0xe7, 0x8c, 0x8c, 0xd3, 0x19, 0x19, 0xff, 0xcf, 0xc8, 0x58,
	0x4e, 0xb7, 0x5c, 0xee, 0x92, 0x95, 0xbd, 0x16, 0x01, 0xa9, 0x80, 0x4f, 0xc5, 0x66, 0xc3, 0xd7,
	0xdc, 0xf3, 0xc9, 0x2e, 0x59, 0x91, 0x9f, 0xfa, 0xd4, 0xd5, 0x4f, 0x58, 0x35, 0xd5, 0x7d, 0x7e,
	0x7c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x64, 0x1f, 0x9d, 0x08, 0x03, 0x00, 0x00,
}

func (m *EventModule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventModule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventModule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAdd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAdd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAdd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventAddNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRemoveNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRemoveNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRemoveNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Node) > 0 {
		i -= len(m.Node)
		copy(dAtA[i:], m.Node)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Node)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventModule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventAdd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventAddNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRemoveNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Node)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventModule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventModule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventModule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAdd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAdd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAdd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAddNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAddNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRemoveNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRemoveNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRemoveNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
