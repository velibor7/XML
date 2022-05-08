// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: connection_service.proto

package connection

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type AddConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddConnectionDTO *UserAction `protobuf:"bytes,1,opt,name=addConnectionDTO,proto3" json:"addConnectionDTO,omitempty"`
}

func (x *AddConnectionRequest) Reset() {
	*x = AddConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConnectionRequest) ProtoMessage() {}

func (x *AddConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConnectionRequest.ProtoReflect.Descriptor instead.
func (*AddConnectionRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddConnectionRequest) GetAddConnectionDTO() *UserAction {
	if x != nil {
		return x.AddConnectionDTO
	}
	return nil
}

type ApproveConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApproveConnectionDTO *UserAction `protobuf:"bytes,1,opt,name=approveConnectionDTO,proto3" json:"approveConnectionDTO,omitempty"`
}

func (x *ApproveConnectionRequest) Reset() {
	*x = ApproveConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveConnectionRequest) ProtoMessage() {}

func (x *ApproveConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveConnectionRequest.ProtoReflect.Descriptor instead.
func (*ApproveConnectionRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{2}
}

func (x *ApproveConnectionRequest) GetApproveConnectionDTO() *UserAction {
	if x != nil {
		return x.ApproveConnectionDTO
	}
	return nil
}

type RejectConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RejectConnectionDTO *UserAction `protobuf:"bytes,1,opt,name=rejectConnectionDTO,proto3" json:"rejectConnectionDTO,omitempty"`
}

func (x *RejectConnectionRequest) Reset() {
	*x = RejectConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectConnectionRequest) ProtoMessage() {}

func (x *RejectConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectConnectionRequest.ProtoReflect.Descriptor instead.
func (*RejectConnectionRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{3}
}

func (x *RejectConnectionRequest) GetRejectConnectionDTO() *UserAction {
	if x != nil {
		return x.RejectConnectionDTO
	}
	return nil
}

type AddBlockUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddBlockUserDTO *UserAction `protobuf:"bytes,1,opt,name=addBlockUserDTO,proto3" json:"addBlockUserDTO,omitempty"`
}

func (x *AddBlockUserRequest) Reset() {
	*x = AddBlockUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlockUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlockUserRequest) ProtoMessage() {}

func (x *AddBlockUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlockUserRequest.ProtoReflect.Descriptor instead.
func (*AddBlockUserRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddBlockUserRequest) GetAddBlockUserDTO() *UserAction {
	if x != nil {
		return x.AddBlockUserDTO
	}
	return nil
}

type ActionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ActionResult) Reset() {
	*x = ActionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResult) ProtoMessage() {}

func (x *ActionResult) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResult.ProtoReflect.Descriptor instead.
func (*ActionResult) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{5}
}

func (x *ActionResult) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ActionResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDa string `protobuf:"bytes,1,opt,name=userIDa,proto3" json:"userIDa,omitempty"`
	UserIDb string `protobuf:"bytes,2,opt,name=userIDb,proto3" json:"userIDb,omitempty"`
}

func (x *UserAction) Reset() {
	*x = UserAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAction) ProtoMessage() {}

func (x *UserAction) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAction.ProtoReflect.Descriptor instead.
func (*UserAction) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{7}
}

func (x *UserAction) GetUserIDa() string {
	if x != nil {
		return x.UserIDa
	}
	return ""
}

func (x *UserAction) GetUserIDb() string {
	if x != nil {
		return x.UserIDb
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	IsPublic bool   `protobuf:"varint,2,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{9}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UsersID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID []string `protobuf:"bytes,1,rep,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UsersID) Reset() {
	*x = UsersID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connection_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersID) ProtoMessage() {}

func (x *UsersID) ProtoReflect() protoreflect.Message {
	mi := &file_connection_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersID.ProtoReflect.Descriptor instead.
func (*UsersID) Descriptor() ([]byte, []int) {
	return file_connection_service_proto_rawDescGZIP(), []int{10}
}

func (x *UsersID) GetUserID() []string {
	if x != nil {
		return x.UserID
	}
	return nil
}

var File_connection_service_proto protoreflect.FileDescriptor

var file_connection_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x62,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x10, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x54, 0x4f, 0x22, 0x6e, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52,
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x54, 0x4f, 0x22, 0x6b, 0x0a, 0x17, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a,
	0x13, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x54, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x22,
	0x5f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x61, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f,
	0x22, 0x38, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x40, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x62, 0x22, 0x3a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x37,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x21, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x32, 0xb5, 0x05, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x7d,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x75, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x14, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x82, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x0b, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x54, 0x4f, 0x12, 0x92, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x1a, 0x12, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x3a, 0x13, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x12, 0x96, 0x01, 0x0a, 0x11, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x1a, 0x13, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x3a, 0x14, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x54, 0x4f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x65, 0x6c, 0x69, 0x62, 0x6f, 0x72, 0x37, 0x2f, 0x58, 0x4d, 0x4c, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connection_service_proto_rawDescOnce sync.Once
	file_connection_service_proto_rawDescData = file_connection_service_proto_rawDesc
)

func file_connection_service_proto_rawDescGZIP() []byte {
	file_connection_service_proto_rawDescOnce.Do(func() {
		file_connection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_connection_service_proto_rawDescData)
	})
	return file_connection_service_proto_rawDescData
}

var file_connection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_connection_service_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),          // 0: connection_service.RegisterRequest
	(*AddConnectionRequest)(nil),     // 1: connection_service.AddConnectionRequest
	(*ApproveConnectionRequest)(nil), // 2: connection_service.ApproveConnectionRequest
	(*RejectConnectionRequest)(nil),  // 3: connection_service.RejectConnectionRequest
	(*AddBlockUserRequest)(nil),      // 4: connection_service.AddBlockUserRequest
	(*ActionResult)(nil),             // 5: connection_service.ActionResult
	(*GetRequest)(nil),               // 6: connection_service.GetRequest
	(*UserAction)(nil),               // 7: connection_service.UserAction
	(*User)(nil),                     // 8: connection_service.User
	(*Users)(nil),                    // 9: connection_service.Users
	(*UsersID)(nil),                  // 10: connection_service.UsersID
}
var file_connection_service_proto_depIdxs = []int32{
	8,  // 0: connection_service.RegisterRequest.user:type_name -> connection_service.User
	7,  // 1: connection_service.AddConnectionRequest.addConnectionDTO:type_name -> connection_service.UserAction
	7,  // 2: connection_service.ApproveConnectionRequest.approveConnectionDTO:type_name -> connection_service.UserAction
	7,  // 3: connection_service.RejectConnectionRequest.rejectConnectionDTO:type_name -> connection_service.UserAction
	7,  // 4: connection_service.AddBlockUserRequest.addBlockUserDTO:type_name -> connection_service.UserAction
	8,  // 5: connection_service.Users.users:type_name -> connection_service.User
	6,  // 6: connection_service.ConnectionService.GetConnections:input_type -> connection_service.GetRequest
	0,  // 7: connection_service.ConnectionService.Register:input_type -> connection_service.RegisterRequest
	1,  // 8: connection_service.ConnectionService.AddConnection:input_type -> connection_service.AddConnectionRequest
	3,  // 9: connection_service.ConnectionService.RejectConnection:input_type -> connection_service.RejectConnectionRequest
	2,  // 10: connection_service.ConnectionService.ApproveConnection:input_type -> connection_service.ApproveConnectionRequest
	9,  // 11: connection_service.ConnectionService.GetConnections:output_type -> connection_service.Users
	5,  // 12: connection_service.ConnectionService.Register:output_type -> connection_service.ActionResult
	5,  // 13: connection_service.ConnectionService.AddConnection:output_type -> connection_service.ActionResult
	5,  // 14: connection_service.ConnectionService.RejectConnection:output_type -> connection_service.ActionResult
	5,  // 15: connection_service.ConnectionService.ApproveConnection:output_type -> connection_service.ActionResult
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_connection_service_proto_init() }
func file_connection_service_proto_init() {
	if File_connection_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connection_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_connection_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddConnectionRequest); i {
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
		file_connection_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveConnectionRequest); i {
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
		file_connection_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectConnectionRequest); i {
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
		file_connection_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlockUserRequest); i {
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
		file_connection_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResult); i {
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
		file_connection_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_connection_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAction); i {
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
		file_connection_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_connection_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_connection_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersID); i {
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
			RawDescriptor: file_connection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connection_service_proto_goTypes,
		DependencyIndexes: file_connection_service_proto_depIdxs,
		MessageInfos:      file_connection_service_proto_msgTypes,
	}.Build()
	File_connection_service_proto = out.File
	file_connection_service_proto_rawDesc = nil
	file_connection_service_proto_goTypes = nil
	file_connection_service_proto_depIdxs = nil
}
