// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/proto/author_service.proto

package pb

import (
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

//
//=================
//Metadata
//=================
type AuthorMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *AuthorMeta) Reset() {
	*x = AuthorMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorMeta) ProtoMessage() {}

func (x *AuthorMeta) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorMeta.ProtoReflect.Descriptor instead.
func (*AuthorMeta) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AuthorMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *AuthorMeta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

//
//=================
//Parameter Request
//=================
type AuthorParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage         int32  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	OrderBy         string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderMethod     string `protobuf:"bytes,4,opt,name=order_method,json=orderMethod,proto3" json:"order_method,omitempty"`
	SearchCondition string `protobuf:"bytes,5,opt,name=search_condition,json=searchCondition,proto3" json:"search_condition,omitempty"`
	Equal           string `protobuf:"bytes,6,opt,name=equal,proto3" json:"equal,omitempty"`
	Not             string `protobuf:"bytes,7,opt,name=not,proto3" json:"not,omitempty"`
	Like            string `protobuf:"bytes,8,opt,name=like,proto3" json:"like,omitempty"`
	DateRangeBy     string `protobuf:"bytes,9,opt,name=date_range_by,json=dateRangeBy,proto3" json:"date_range_by,omitempty"`
	DateStart       string `protobuf:"bytes,10,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd         string `protobuf:"bytes,11,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *AuthorParameterRequest) Reset() {
	*x = AuthorParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorParameterRequest) ProtoMessage() {}

func (x *AuthorParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorParameterRequest.ProtoReflect.Descriptor instead.
func (*AuthorParameterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorParameterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AuthorParameterRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *AuthorParameterRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *AuthorParameterRequest) GetOrderMethod() string {
	if x != nil {
		return x.OrderMethod
	}
	return ""
}

func (x *AuthorParameterRequest) GetSearchCondition() string {
	if x != nil {
		return x.SearchCondition
	}
	return ""
}

func (x *AuthorParameterRequest) GetEqual() string {
	if x != nil {
		return x.Equal
	}
	return ""
}

func (x *AuthorParameterRequest) GetNot() string {
	if x != nil {
		return x.Not
	}
	return ""
}

func (x *AuthorParameterRequest) GetLike() string {
	if x != nil {
		return x.Like
	}
	return ""
}

func (x *AuthorParameterRequest) GetDateRangeBy() string {
	if x != nil {
		return x.DateRangeBy
	}
	return ""
}

func (x *AuthorParameterRequest) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *AuthorParameterRequest) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

type AuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AuthorResponse) Reset() {
	*x = AuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponse) ProtoMessage() {}

func (x *AuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponse.ProtoReflect.Descriptor instead.
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AuthorResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthorResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AuthorResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*AuthorResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta *AuthorMeta       `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *AuthorsResponse) Reset() {
	*x = AuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsResponse) ProtoMessage() {}

func (x *AuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsResponse.ProtoReflect.Descriptor instead.
func (*AuthorsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorsResponse) GetData() []*AuthorResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AuthorsResponse) GetMeta() *AuthorMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AuthorStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *AuthorResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthorStatusResponse) Reset() {
	*x = AuthorStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorStatusResponse) ProtoMessage() {}

func (x *AuthorStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorStatusResponse.ProtoReflect.Descriptor instead.
func (*AuthorStatusResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AuthorStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthorStatusResponse) GetData() *AuthorResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAuthorRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateAuthorRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindAuthorByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindAuthorByIdRequest) Reset() {
	*x = FindAuthorByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_author_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAuthorByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAuthorByIdRequest) ProtoMessage() {}

func (x *FindAuthorByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_author_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAuthorByIdRequest.ProtoReflect.Descriptor instead.
func (*FindAuthorByIdRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_author_service_proto_rawDescGZIP(), []int{6}
}

func (x *FindAuthorByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_grpc_proto_author_service_proto protoreflect.FileDescriptor

var file_grpc_proto_author_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x51, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xca, 0x02, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x73, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x14,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xd5, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x60, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_author_service_proto_rawDescOnce sync.Once
	file_grpc_proto_author_service_proto_rawDescData = file_grpc_proto_author_service_proto_rawDesc
)

func file_grpc_proto_author_service_proto_rawDescGZIP() []byte {
	file_grpc_proto_author_service_proto_rawDescOnce.Do(func() {
		file_grpc_proto_author_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_author_service_proto_rawDescData)
	})
	return file_grpc_proto_author_service_proto_rawDescData
}

var file_grpc_proto_author_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_proto_author_service_proto_goTypes = []interface{}{
	(*AuthorMeta)(nil),             // 0: authorService.AuthorMeta
	(*AuthorParameterRequest)(nil), // 1: authorService.AuthorParameterRequest
	(*AuthorResponse)(nil),         // 2: authorService.AuthorResponse
	(*AuthorsResponse)(nil),        // 3: authorService.AuthorsResponse
	(*AuthorStatusResponse)(nil),   // 4: authorService.AuthorStatusResponse
	(*CreateAuthorRequest)(nil),    // 5: authorService.CreateAuthorRequest
	(*FindAuthorByIdRequest)(nil),  // 6: authorService.FindAuthorByIdRequest
}
var file_grpc_proto_author_service_proto_depIdxs = []int32{
	2, // 0: authorService.AuthorsResponse.data:type_name -> authorService.AuthorResponse
	0, // 1: authorService.AuthorsResponse.meta:type_name -> authorService.AuthorMeta
	2, // 2: authorService.AuthorStatusResponse.data:type_name -> authorService.AuthorResponse
	5, // 3: authorService.AuthorService.ServiceRegisterAuthor:input_type -> authorService.CreateAuthorRequest
	6, // 4: authorService.AuthorService.ServiceFindAuthorById:input_type -> authorService.FindAuthorByIdRequest
	4, // 5: authorService.AuthorService.ServiceRegisterAuthor:output_type -> authorService.AuthorStatusResponse
	4, // 6: authorService.AuthorService.ServiceFindAuthorById:output_type -> authorService.AuthorStatusResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_proto_author_service_proto_init() }
func file_grpc_proto_author_service_proto_init() {
	if File_grpc_proto_author_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_author_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorMeta); i {
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
		file_grpc_proto_author_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorParameterRequest); i {
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
		file_grpc_proto_author_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorResponse); i {
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
		file_grpc_proto_author_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorsResponse); i {
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
		file_grpc_proto_author_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorStatusResponse); i {
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
		file_grpc_proto_author_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorRequest); i {
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
		file_grpc_proto_author_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAuthorByIdRequest); i {
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
			RawDescriptor: file_grpc_proto_author_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_author_service_proto_goTypes,
		DependencyIndexes: file_grpc_proto_author_service_proto_depIdxs,
		MessageInfos:      file_grpc_proto_author_service_proto_msgTypes,
	}.Build()
	File_grpc_proto_author_service_proto = out.File
	file_grpc_proto_author_service_proto_rawDesc = nil
	file_grpc_proto_author_service_proto_goTypes = nil
	file_grpc_proto_author_service_proto_depIdxs = nil
}
