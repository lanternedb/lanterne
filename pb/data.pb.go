// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: data.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_OK                    Status = 0
	Status_INTERNAL_SERVER_ERROR Status = 1
	Status_INVALID_REQUEST       Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "INTERNAL_SERVER_ERROR",
		2: "INVALID_REQUEST",
	}
	Status_value = map[string]int32{
		"OK":                    0,
		"INTERNAL_SERVER_ERROR": 1,
		"INVALID_REQUEST":       2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_data_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_data_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Value:
	//	*Vertex_Float64
	//	*Vertex_Float32
	//	*Vertex_Int32
	//	*Vertex_Int64
	//	*Vertex_Uint32
	//	*Vertex_Uint64
	//	*Vertex_Bool
	//	*Vertex_String_
	//	*Vertex_Bytes
	//	*Vertex_Timestamp
	//	*Vertex_Nil
	Value isVertex_Value `protobuf_oneof:"value"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *Vertex) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *Vertex) GetValue() isVertex_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Vertex) GetFloat64() float64 {
	if x, ok := x.GetValue().(*Vertex_Float64); ok {
		return x.Float64
	}
	return 0
}

func (x *Vertex) GetFloat32() float32 {
	if x, ok := x.GetValue().(*Vertex_Float32); ok {
		return x.Float32
	}
	return 0
}

func (x *Vertex) GetInt32() int32 {
	if x, ok := x.GetValue().(*Vertex_Int32); ok {
		return x.Int32
	}
	return 0
}

func (x *Vertex) GetInt64() int64 {
	if x, ok := x.GetValue().(*Vertex_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *Vertex) GetUint32() uint32 {
	if x, ok := x.GetValue().(*Vertex_Uint32); ok {
		return x.Uint32
	}
	return 0
}

func (x *Vertex) GetUint64() uint64 {
	if x, ok := x.GetValue().(*Vertex_Uint64); ok {
		return x.Uint64
	}
	return 0
}

func (x *Vertex) GetBool() bool {
	if x, ok := x.GetValue().(*Vertex_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Vertex) GetString_() string {
	if x, ok := x.GetValue().(*Vertex_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Vertex) GetBytes() []byte {
	if x, ok := x.GetValue().(*Vertex_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *Vertex) GetTimestamp() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*Vertex_Timestamp); ok {
		return x.Timestamp
	}
	return nil
}

func (x *Vertex) GetNil() bool {
	if x, ok := x.GetValue().(*Vertex_Nil); ok {
		return x.Nil
	}
	return false
}

type isVertex_Value interface {
	isVertex_Value()
}

type Vertex_Float64 struct {
	Float64 float64 `protobuf:"fixed64,2,opt,name=float64,proto3,oneof"`
}

type Vertex_Float32 struct {
	Float32 float32 `protobuf:"fixed32,3,opt,name=float32,proto3,oneof"`
}

type Vertex_Int32 struct {
	Int32 int32 `protobuf:"varint,4,opt,name=int32,proto3,oneof"`
}

type Vertex_Int64 struct {
	Int64 int64 `protobuf:"varint,5,opt,name=int64,proto3,oneof"`
}

type Vertex_Uint32 struct {
	Uint32 uint32 `protobuf:"varint,6,opt,name=uint32,proto3,oneof"`
}

type Vertex_Uint64 struct {
	Uint64 uint64 `protobuf:"varint,7,opt,name=uint64,proto3,oneof"`
}

type Vertex_Bool struct {
	Bool bool `protobuf:"varint,8,opt,name=bool,proto3,oneof"`
}

type Vertex_String_ struct {
	String_ string `protobuf:"bytes,9,opt,name=string,proto3,oneof"`
}

type Vertex_Bytes struct {
	Bytes []byte `protobuf:"bytes,10,opt,name=bytes,proto3,oneof"`
}

type Vertex_Timestamp struct {
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=timestamp,proto3,oneof"`
}

type Vertex_Nil struct {
	Nil bool `protobuf:"varint,30,opt,name=nil,proto3,oneof"`
}

func (*Vertex_Float64) isVertex_Value() {}

func (*Vertex_Float32) isVertex_Value() {}

func (*Vertex_Int32) isVertex_Value() {}

func (*Vertex_Int64) isVertex_Value() {}

func (*Vertex_Uint32) isVertex_Value() {}

func (*Vertex_Uint64) isVertex_Value() {}

func (*Vertex_Bool) isVertex_Value() {}

func (*Vertex_String_) isVertex_Value() {}

func (*Vertex_Bytes) isVertex_Value() {}

func (*Vertex_Timestamp) isVertex_Value() {}

func (*Vertex_Nil) isVertex_Value() {}

type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tail   *Vertex `protobuf:"bytes,1,opt,name=tail,proto3" json:"tail,omitempty"`
	Head   *Vertex `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	Weight float32 `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *Edge) GetTail() *Vertex {
	if x != nil {
		return x.Tail
	}
	return nil
}

func (x *Edge) GetHead() *Vertex {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *Edge) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Neighbor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeightMap map[string]float32 `protobuf:"bytes,1,rep,name=weight_map,json=weightMap,proto3" json:"weight_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Neighbor) Reset() {
	*x = Neighbor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Neighbor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Neighbor) ProtoMessage() {}

func (x *Neighbor) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Neighbor.ProtoReflect.Descriptor instead.
func (*Neighbor) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *Neighbor) GetWeightMap() map[string]float32 {
	if x != nil {
		return x.WeightMap
	}
	return nil
}

type Graph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VertexMap   map[string]*Vertex   `protobuf:"bytes,1,rep,name=vertex_map,json=vertexMap,proto3" json:"vertex_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NeighborMap map[string]*Neighbor `protobuf:"bytes,2,rep,name=neighbor_map,json=neighborMap,proto3" json:"neighbor_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Graph) Reset() {
	*x = Graph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Graph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Graph) ProtoMessage() {}

func (x *Graph) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Graph.ProtoReflect.Descriptor instead.
func (*Graph) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *Graph) GetVertexMap() map[string]*Vertex {
	if x != nil {
		return x.VertexMap
	}
	return nil
}

func (x *Graph) GetNeighborMap() map[string]*Neighbor {
	if x != nil {
		return x.NeighborMap
	}
	return nil
}

type IlluminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed      *Vertex `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Step      uint32  `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	MinWeight float32 `protobuf:"fixed32,3,opt,name=min_weight,json=minWeight,proto3" json:"min_weight,omitempty"`
	MaxWeight float32 `protobuf:"fixed32,4,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
}

func (x *IlluminateRequest) Reset() {
	*x = IlluminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IlluminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IlluminateRequest) ProtoMessage() {}

func (x *IlluminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IlluminateRequest.ProtoReflect.Descriptor instead.
func (*IlluminateRequest) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *IlluminateRequest) GetSeed() *Vertex {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *IlluminateRequest) GetStep() uint32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *IlluminateRequest) GetMinWeight() float32 {
	if x != nil {
		return x.MinWeight
	}
	return 0
}

func (x *IlluminateRequest) GetMaxWeight() float32 {
	if x != nil {
		return x.MaxWeight
	}
	return 0
}

type IlluminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Graph  *Graph `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *IlluminateResponse) Reset() {
	*x = IlluminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IlluminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IlluminateResponse) ProtoMessage() {}

func (x *IlluminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IlluminateResponse.ProtoReflect.Descriptor instead.
func (*IlluminateResponse) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{5}
}

func (x *IlluminateResponse) GetGraph() *Graph {
	if x != nil {
		return x.Graph
	}
	return nil
}

func (x *IlluminateResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

type DumpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *DumpResponse) Reset() {
	*x = DumpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpResponse) ProtoMessage() {}

func (x *DumpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpResponse.ProtoReflect.Descriptor instead.
func (*DumpResponse) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{6}
}

func (x *DumpResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02,
	0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x33, 0x32, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x18, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x06,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x03, 0x6e, 0x69,
	0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x6e, 0x69, 0x6c, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x58, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x12, 0x37,
	0x0a, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x2e, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x3c, 0x0a, 0x0e, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8b, 0x02, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12,
	0x34, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x4d, 0x61, 0x70, 0x12, 0x3a, 0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x4d, 0x61,
	0x70, 0x1a, 0x45, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x49, 0x0a, 0x10, 0x4e, 0x65, 0x69, 0x67,
	0x68, 0x62, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x49, 0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x73, 0x65, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69,
	0x6e, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x6d, 0x69, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6d,
	0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x53, 0x0a, 0x12, 0x49, 0x6c, 0x6c, 0x75,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1f, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a,
	0x0c, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x40,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02,
	0x32, 0x8f, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x49, 0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x49, 0x6c,
	0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x49, 0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x70, 0x56, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x12, 0x07, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x1a, 0x0d, 0x2e,
	0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x22,
	0x0a, 0x08, 0x44, 0x75, 0x6d, 0x70, 0x45, 0x64, 0x67, 0x65, 0x12, 0x05, 0x2e, 0x45, 0x64, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x64, 0x62, 0x2f, 0x6c, 0x61, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_data_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: Status
	(*Vertex)(nil),                // 1: Vertex
	(*Edge)(nil),                  // 2: Edge
	(*Neighbor)(nil),              // 3: Neighbor
	(*Graph)(nil),                 // 4: Graph
	(*IlluminateRequest)(nil),     // 5: IlluminateRequest
	(*IlluminateResponse)(nil),    // 6: IlluminateResponse
	(*DumpResponse)(nil),          // 7: DumpResponse
	nil,                           // 8: Neighbor.WeightMapEntry
	nil,                           // 9: Graph.VertexMapEntry
	nil,                           // 10: Graph.NeighborMapEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_data_proto_depIdxs = []int32{
	11, // 0: Vertex.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 1: Edge.tail:type_name -> Vertex
	1,  // 2: Edge.head:type_name -> Vertex
	8,  // 3: Neighbor.weight_map:type_name -> Neighbor.WeightMapEntry
	9,  // 4: Graph.vertex_map:type_name -> Graph.VertexMapEntry
	10, // 5: Graph.neighbor_map:type_name -> Graph.NeighborMapEntry
	1,  // 6: IlluminateRequest.seed:type_name -> Vertex
	4,  // 7: IlluminateResponse.graph:type_name -> Graph
	0,  // 8: IlluminateResponse.status:type_name -> Status
	0,  // 9: DumpResponse.status:type_name -> Status
	1,  // 10: Graph.VertexMapEntry.value:type_name -> Vertex
	3,  // 11: Graph.NeighborMapEntry.value:type_name -> Neighbor
	5,  // 12: Lanterne.Illuminate:input_type -> IlluminateRequest
	1,  // 13: Lanterne.DumpVertex:input_type -> Vertex
	2,  // 14: Lanterne.DumpEdge:input_type -> Edge
	6,  // 15: Lanterne.Illuminate:output_type -> IlluminateResponse
	7,  // 16: Lanterne.DumpVertex:output_type -> DumpResponse
	7,  // 17: Lanterne.DumpEdge:output_type -> DumpResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Neighbor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Graph); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IlluminateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IlluminateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DumpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_data_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Vertex_Float64)(nil),
		(*Vertex_Float32)(nil),
		(*Vertex_Int32)(nil),
		(*Vertex_Int64)(nil),
		(*Vertex_Uint32)(nil),
		(*Vertex_Uint64)(nil),
		(*Vertex_Bool)(nil),
		(*Vertex_String_)(nil),
		(*Vertex_Bytes)(nil),
		(*Vertex_Timestamp)(nil),
		(*Vertex_Nil)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		EnumInfos:         file_data_proto_enumTypes,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
