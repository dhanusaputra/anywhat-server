// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: anywhat.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Anything ...
type Anything struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Anything) Reset() {
	*x = Anything{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Anything) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Anything) ProtoMessage() {}

func (x *Anything) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Anything.ProtoReflect.Descriptor instead.
func (*Anything) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{0}
}

func (x *Anything) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Anything) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Anything) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Anything) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// GetAnythingRequest ...
type GetAnythingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAnythingRequest) Reset() {
	*x = GetAnythingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnythingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnythingRequest) ProtoMessage() {}

func (x *GetAnythingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnythingRequest.ProtoReflect.Descriptor instead.
func (*GetAnythingRequest) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnythingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetAnythingResponse ...
type GetAnythingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anything *Anything `protobuf:"bytes,1,opt,name=Anything,proto3" json:"Anything,omitempty"`
}

func (x *GetAnythingResponse) Reset() {
	*x = GetAnythingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnythingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnythingResponse) ProtoMessage() {}

func (x *GetAnythingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnythingResponse.ProtoReflect.Descriptor instead.
func (*GetAnythingResponse) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{2}
}

func (x *GetAnythingResponse) GetAnything() *Anything {
	if x != nil {
		return x.Anything
	}
	return nil
}

// UpdateAnythingRequest ...
type UpdateAnythingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anything *Anything `protobuf:"bytes,1,opt,name=Anything,proto3" json:"Anything,omitempty"`
}

func (x *UpdateAnythingRequest) Reset() {
	*x = UpdateAnythingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnythingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnythingRequest) ProtoMessage() {}

func (x *UpdateAnythingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnythingRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnythingRequest) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAnythingRequest) GetAnything() *Anything {
	if x != nil {
		return x.Anything
	}
	return nil
}

// UpdateAnythingResponse ...
type UpdateAnythingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated bool `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateAnythingResponse) Reset() {
	*x = UpdateAnythingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnythingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnythingResponse) ProtoMessage() {}

func (x *UpdateAnythingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnythingResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnythingResponse) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAnythingResponse) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

// CreateAnythingRequest ...
type CreateAnythingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anything *Anything `protobuf:"bytes,1,opt,name=Anything,proto3" json:"Anything,omitempty"`
}

func (x *CreateAnythingRequest) Reset() {
	*x = CreateAnythingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnythingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnythingRequest) ProtoMessage() {}

func (x *CreateAnythingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnythingRequest.ProtoReflect.Descriptor instead.
func (*CreateAnythingRequest) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAnythingRequest) GetAnything() *Anything {
	if x != nil {
		return x.Anything
	}
	return nil
}

// CreateAnythingResponse ...
type CreateAnythingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateAnythingResponse) Reset() {
	*x = CreateAnythingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnythingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnythingResponse) ProtoMessage() {}

func (x *CreateAnythingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnythingResponse.ProtoReflect.Descriptor instead.
func (*CreateAnythingResponse) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAnythingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteAnythingRequest ...
type DeleteAnythingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAnythingRequest) Reset() {
	*x = DeleteAnythingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAnythingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAnythingRequest) ProtoMessage() {}

func (x *DeleteAnythingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAnythingRequest.ProtoReflect.Descriptor instead.
func (*DeleteAnythingRequest) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAnythingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteAnythingResponse ...
type DeleteAnythingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteAnythingResponse) Reset() {
	*x = DeleteAnythingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAnythingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAnythingResponse) ProtoMessage() {}

func (x *DeleteAnythingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAnythingResponse.ProtoReflect.Descriptor instead.
func (*DeleteAnythingResponse) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAnythingResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// ListAnythingRequest ...
type ListAnythingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAnythingRequest) Reset() {
	*x = ListAnythingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAnythingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAnythingRequest) ProtoMessage() {}

func (x *ListAnythingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAnythingRequest.ProtoReflect.Descriptor instead.
func (*ListAnythingRequest) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{9}
}

// ListAnythingResponse ...
type ListAnythingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anythings []*Anything `protobuf:"bytes,1,rep,name=Anythings,proto3" json:"Anythings,omitempty"`
}

func (x *ListAnythingResponse) Reset() {
	*x = ListAnythingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anywhat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAnythingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAnythingResponse) ProtoMessage() {}

func (x *ListAnythingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anywhat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAnythingResponse.ProtoReflect.Descriptor instead.
func (*ListAnythingResponse) Descriptor() ([]byte, []int) {
	return file_anywhat_proto_rawDescGZIP(), []int{10}
}

func (x *ListAnythingResponse) GetAnythings() []*Anything {
	if x != nil {
		return x.Anythings
	}
	return nil
}

var File_anywhat_proto protoreflect.FileDescriptor

var file_anywhat_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x79, 0x77, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x08, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x32, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x42, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x41, 0x6e, 0x79, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x41, 0x6e, 0x79, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x07, 0x41, 0x6e, 0x79, 0x77, 0x68, 0x61, 0x74,
	0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x79,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anywhat_proto_rawDescOnce sync.Once
	file_anywhat_proto_rawDescData = file_anywhat_proto_rawDesc
)

func file_anywhat_proto_rawDescGZIP() []byte {
	file_anywhat_proto_rawDescOnce.Do(func() {
		file_anywhat_proto_rawDescData = protoimpl.X.CompressGZIP(file_anywhat_proto_rawDescData)
	})
	return file_anywhat_proto_rawDescData
}

var file_anywhat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_anywhat_proto_goTypes = []interface{}{
	(*Anything)(nil),               // 0: pb.Anything
	(*GetAnythingRequest)(nil),     // 1: pb.GetAnythingRequest
	(*GetAnythingResponse)(nil),    // 2: pb.GetAnythingResponse
	(*UpdateAnythingRequest)(nil),  // 3: pb.UpdateAnythingRequest
	(*UpdateAnythingResponse)(nil), // 4: pb.UpdateAnythingResponse
	(*CreateAnythingRequest)(nil),  // 5: pb.CreateAnythingRequest
	(*CreateAnythingResponse)(nil), // 6: pb.CreateAnythingResponse
	(*DeleteAnythingRequest)(nil),  // 7: pb.DeleteAnythingRequest
	(*DeleteAnythingResponse)(nil), // 8: pb.DeleteAnythingResponse
	(*ListAnythingRequest)(nil),    // 9: pb.ListAnythingRequest
	(*ListAnythingResponse)(nil),   // 10: pb.ListAnythingResponse
	(*timestamp.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_anywhat_proto_depIdxs = []int32{
	11, // 0: pb.Anything.createdAt:type_name -> google.protobuf.Timestamp
	0,  // 1: pb.GetAnythingResponse.Anything:type_name -> pb.Anything
	0,  // 2: pb.UpdateAnythingRequest.Anything:type_name -> pb.Anything
	0,  // 3: pb.CreateAnythingRequest.Anything:type_name -> pb.Anything
	0,  // 4: pb.ListAnythingResponse.Anythings:type_name -> pb.Anything
	1,  // 5: pb.Anywhat.GetAnything:input_type -> pb.GetAnythingRequest
	3,  // 6: pb.Anywhat.UpdateAnything:input_type -> pb.UpdateAnythingRequest
	5,  // 7: pb.Anywhat.CreateAnything:input_type -> pb.CreateAnythingRequest
	7,  // 8: pb.Anywhat.DeleteAnything:input_type -> pb.DeleteAnythingRequest
	9,  // 9: pb.Anywhat.ListAnything:input_type -> pb.ListAnythingRequest
	2,  // 10: pb.Anywhat.GetAnything:output_type -> pb.GetAnythingResponse
	4,  // 11: pb.Anywhat.UpdateAnything:output_type -> pb.UpdateAnythingResponse
	6,  // 12: pb.Anywhat.CreateAnything:output_type -> pb.CreateAnythingResponse
	8,  // 13: pb.Anywhat.DeleteAnything:output_type -> pb.DeleteAnythingResponse
	10, // 14: pb.Anywhat.ListAnything:output_type -> pb.ListAnythingResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_anywhat_proto_init() }
func file_anywhat_proto_init() {
	if File_anywhat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anywhat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Anything); i {
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
		file_anywhat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnythingRequest); i {
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
		file_anywhat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnythingResponse); i {
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
		file_anywhat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnythingRequest); i {
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
		file_anywhat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnythingResponse); i {
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
		file_anywhat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnythingRequest); i {
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
		file_anywhat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnythingResponse); i {
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
		file_anywhat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAnythingRequest); i {
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
		file_anywhat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAnythingResponse); i {
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
		file_anywhat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAnythingRequest); i {
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
		file_anywhat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAnythingResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_anywhat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anywhat_proto_goTypes,
		DependencyIndexes: file_anywhat_proto_depIdxs,
		MessageInfos:      file_anywhat_proto_msgTypes,
	}.Build()
	File_anywhat_proto = out.File
	file_anywhat_proto_rawDesc = nil
	file_anywhat_proto_goTypes = nil
	file_anywhat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnywhatClient is the client API for Anywhat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnywhatClient interface {
	// GetAnything ...
	GetAnything(ctx context.Context, in *GetAnythingRequest, opts ...grpc.CallOption) (*GetAnythingResponse, error)
	// UpdateAnything ...
	UpdateAnything(ctx context.Context, in *UpdateAnythingRequest, opts ...grpc.CallOption) (*UpdateAnythingResponse, error)
	// CreateAnything ...
	CreateAnything(ctx context.Context, in *CreateAnythingRequest, opts ...grpc.CallOption) (*CreateAnythingResponse, error)
	// DeleteAnything ...
	DeleteAnything(ctx context.Context, in *DeleteAnythingRequest, opts ...grpc.CallOption) (*DeleteAnythingResponse, error)
	// ListAnything ...
	ListAnything(ctx context.Context, in *ListAnythingRequest, opts ...grpc.CallOption) (*ListAnythingResponse, error)
}

type anywhatClient struct {
	cc grpc.ClientConnInterface
}

func NewAnywhatClient(cc grpc.ClientConnInterface) AnywhatClient {
	return &anywhatClient{cc}
}

func (c *anywhatClient) GetAnything(ctx context.Context, in *GetAnythingRequest, opts ...grpc.CallOption) (*GetAnythingResponse, error) {
	out := new(GetAnythingResponse)
	err := c.cc.Invoke(ctx, "/pb.Anywhat/GetAnything", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anywhatClient) UpdateAnything(ctx context.Context, in *UpdateAnythingRequest, opts ...grpc.CallOption) (*UpdateAnythingResponse, error) {
	out := new(UpdateAnythingResponse)
	err := c.cc.Invoke(ctx, "/pb.Anywhat/UpdateAnything", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anywhatClient) CreateAnything(ctx context.Context, in *CreateAnythingRequest, opts ...grpc.CallOption) (*CreateAnythingResponse, error) {
	out := new(CreateAnythingResponse)
	err := c.cc.Invoke(ctx, "/pb.Anywhat/CreateAnything", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anywhatClient) DeleteAnything(ctx context.Context, in *DeleteAnythingRequest, opts ...grpc.CallOption) (*DeleteAnythingResponse, error) {
	out := new(DeleteAnythingResponse)
	err := c.cc.Invoke(ctx, "/pb.Anywhat/DeleteAnything", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anywhatClient) ListAnything(ctx context.Context, in *ListAnythingRequest, opts ...grpc.CallOption) (*ListAnythingResponse, error) {
	out := new(ListAnythingResponse)
	err := c.cc.Invoke(ctx, "/pb.Anywhat/ListAnything", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnywhatServer is the server API for Anywhat service.
type AnywhatServer interface {
	// GetAnything ...
	GetAnything(context.Context, *GetAnythingRequest) (*GetAnythingResponse, error)
	// UpdateAnything ...
	UpdateAnything(context.Context, *UpdateAnythingRequest) (*UpdateAnythingResponse, error)
	// CreateAnything ...
	CreateAnything(context.Context, *CreateAnythingRequest) (*CreateAnythingResponse, error)
	// DeleteAnything ...
	DeleteAnything(context.Context, *DeleteAnythingRequest) (*DeleteAnythingResponse, error)
	// ListAnything ...
	ListAnything(context.Context, *ListAnythingRequest) (*ListAnythingResponse, error)
}

// UnimplementedAnywhatServer can be embedded to have forward compatible implementations.
type UnimplementedAnywhatServer struct {
}

func (*UnimplementedAnywhatServer) GetAnything(context.Context, *GetAnythingRequest) (*GetAnythingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnything not implemented")
}
func (*UnimplementedAnywhatServer) UpdateAnything(context.Context, *UpdateAnythingRequest) (*UpdateAnythingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnything not implemented")
}
func (*UnimplementedAnywhatServer) CreateAnything(context.Context, *CreateAnythingRequest) (*CreateAnythingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnything not implemented")
}
func (*UnimplementedAnywhatServer) DeleteAnything(context.Context, *DeleteAnythingRequest) (*DeleteAnythingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnything not implemented")
}
func (*UnimplementedAnywhatServer) ListAnything(context.Context, *ListAnythingRequest) (*ListAnythingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnything not implemented")
}

func RegisterAnywhatServer(s *grpc.Server, srv AnywhatServer) {
	s.RegisterService(&_Anywhat_serviceDesc, srv)
}

func _Anywhat_GetAnything_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnythingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnywhatServer).GetAnything(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Anywhat/GetAnything",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnywhatServer).GetAnything(ctx, req.(*GetAnythingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anywhat_UpdateAnything_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnythingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnywhatServer).UpdateAnything(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Anywhat/UpdateAnything",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnywhatServer).UpdateAnything(ctx, req.(*UpdateAnythingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anywhat_CreateAnything_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnythingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnywhatServer).CreateAnything(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Anywhat/CreateAnything",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnywhatServer).CreateAnything(ctx, req.(*CreateAnythingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anywhat_DeleteAnything_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnythingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnywhatServer).DeleteAnything(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Anywhat/DeleteAnything",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnywhatServer).DeleteAnything(ctx, req.(*DeleteAnythingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anywhat_ListAnything_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnythingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnywhatServer).ListAnything(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Anywhat/ListAnything",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnywhatServer).ListAnything(ctx, req.(*ListAnythingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Anywhat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Anywhat",
	HandlerType: (*AnywhatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnything",
			Handler:    _Anywhat_GetAnything_Handler,
		},
		{
			MethodName: "UpdateAnything",
			Handler:    _Anywhat_UpdateAnything_Handler,
		},
		{
			MethodName: "CreateAnything",
			Handler:    _Anywhat_CreateAnything_Handler,
		},
		{
			MethodName: "DeleteAnything",
			Handler:    _Anywhat_DeleteAnything_Handler,
		},
		{
			MethodName: "ListAnything",
			Handler:    _Anywhat_ListAnything_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anywhat.proto",
}
