// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: helloworld/v1/greeter.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type DingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationID            string              `protobuf:"bytes,1,opt,name=conversationID,proto3" json:"conversationID,omitempty"`
	AtUsers                   []*DingRequest_User `protobuf:"bytes,2,rep,name=atUsers,proto3" json:"atUsers,omitempty"`
	ChatbotUserID             string              `protobuf:"bytes,3,opt,name=chatbotUserID,proto3" json:"chatbotUserID,omitempty"`
	MsgID                     string              `protobuf:"bytes,4,opt,name=msgID,proto3" json:"msgID,omitempty"`
	SenderNick                string              `protobuf:"bytes,5,opt,name=senderNick,proto3" json:"senderNick,omitempty"`
	IsAdmin                   bool                `protobuf:"varint,6,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	SessionWebhookExpiredTime int64               `protobuf:"varint,7,opt,name=sessionWebhookExpiredTime,proto3" json:"sessionWebhookExpiredTime,omitempty"`
	CreateAt                  int64               `protobuf:"varint,8,opt,name=createAt,proto3" json:"createAt,omitempty"`
	ConversationType          string              `protobuf:"bytes,9,opt,name=conversationType,proto3" json:"conversationType,omitempty"`
	SenderID                  string              `protobuf:"bytes,10,opt,name=senderID,proto3" json:"senderID,omitempty"`
	ConversationTitle         string              `protobuf:"bytes,11,opt,name=conversationTitle,proto3" json:"conversationTitle,omitempty"`
	IsInAtList                bool                `protobuf:"varint,12,opt,name=isInAtList,proto3" json:"isInAtList,omitempty"`
	SessionWebhook            string              `protobuf:"bytes,13,opt,name=sessionWebhook,proto3" json:"sessionWebhook,omitempty"`
	Text                      *Text               `protobuf:"bytes,14,opt,name=text,proto3" json:"text,omitempty"`
	RobotCode                 string              `protobuf:"bytes,15,opt,name=robotCode,proto3" json:"robotCode,omitempty"`
	Msgtype                   string              `protobuf:"bytes,16,opt,name=msgtype,proto3" json:"msgtype,omitempty"`
}

func (x *DingRequest) Reset() {
	*x = DingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingRequest) ProtoMessage() {}

func (x *DingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingRequest.ProtoReflect.Descriptor instead.
func (*DingRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *DingRequest) GetConversationID() string {
	if x != nil {
		return x.ConversationID
	}
	return ""
}

func (x *DingRequest) GetAtUsers() []*DingRequest_User {
	if x != nil {
		return x.AtUsers
	}
	return nil
}

func (x *DingRequest) GetChatbotUserID() string {
	if x != nil {
		return x.ChatbotUserID
	}
	return ""
}

func (x *DingRequest) GetMsgID() string {
	if x != nil {
		return x.MsgID
	}
	return ""
}

func (x *DingRequest) GetSenderNick() string {
	if x != nil {
		return x.SenderNick
	}
	return ""
}

func (x *DingRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *DingRequest) GetSessionWebhookExpiredTime() int64 {
	if x != nil {
		return x.SessionWebhookExpiredTime
	}
	return 0
}

func (x *DingRequest) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *DingRequest) GetConversationType() string {
	if x != nil {
		return x.ConversationType
	}
	return ""
}

func (x *DingRequest) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *DingRequest) GetConversationTitle() string {
	if x != nil {
		return x.ConversationTitle
	}
	return ""
}

func (x *DingRequest) GetIsInAtList() bool {
	if x != nil {
		return x.IsInAtList
	}
	return false
}

func (x *DingRequest) GetSessionWebhook() string {
	if x != nil {
		return x.SessionWebhook
	}
	return ""
}

func (x *DingRequest) GetText() *Text {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *DingRequest) GetRobotCode() string {
	if x != nil {
		return x.RobotCode
	}
	return ""
}

func (x *DingRequest) GetMsgtype() string {
	if x != nil {
		return x.Msgtype
	}
	return ""
}

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *Text) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// The response message containing the greetings
type DingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DingReply) Reset() {
	*x = DingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingReply) ProtoMessage() {}

func (x *DingReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingReply.ProtoReflect.Descriptor instead.
func (*DingReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{2}
}

func (x *DingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DingRequest_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DingtalkID string `protobuf:"bytes,1,opt,name=dingtalkID,proto3" json:"dingtalkID,omitempty"`
}

func (x *DingRequest_User) Reset() {
	*x = DingRequest_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingRequest_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingRequest_User) ProtoMessage() {}

func (x *DingRequest_User) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingRequest_User.ProtoReflect.Descriptor instead.
func (*DingRequest_User) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DingRequest_User) GetDingtalkID() string {
	if x != nil {
		return x.DingtalkID
	}
	return ""
}

var File_helloworld_v1_greeter_proto protoreflect.FileDescriptor

var file_helloworld_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x05, 0x0a, 0x0b, 0x44,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x19, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x49, 0x6e, 0x41, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x49, 0x6e, 0x41, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x27, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x26, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x61, 0x6c, 0x6b,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x61,
	0x6c, 0x6b, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x44, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5c, 0x0a,
	0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x22, 0x05, 0x2f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x51, 0x0a, 0x1c, 0x64,
	0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01,
	0x5a, 0x1c, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_greeter_proto_rawDescOnce sync.Once
	file_helloworld_v1_greeter_proto_rawDescData = file_helloworld_v1_greeter_proto_rawDesc
)

func file_helloworld_v1_greeter_proto_rawDescGZIP() []byte {
	file_helloworld_v1_greeter_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_greeter_proto_rawDescData)
	})
	return file_helloworld_v1_greeter_proto_rawDescData
}

var file_helloworld_v1_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_helloworld_v1_greeter_proto_goTypes = []interface{}{
	(*DingRequest)(nil),      // 0: helloworld.v1.DingRequest
	(*Text)(nil),             // 1: helloworld.v1.Text
	(*DingReply)(nil),        // 2: helloworld.v1.DingReply
	(*DingRequest_User)(nil), // 3: helloworld.v1.DingRequest.User
}
var file_helloworld_v1_greeter_proto_depIdxs = []int32{
	3, // 0: helloworld.v1.DingRequest.atUsers:type_name -> helloworld.v1.DingRequest.User
	1, // 1: helloworld.v1.DingRequest.text:type_name -> helloworld.v1.Text
	0, // 2: helloworld.v1.Greeter.ChatMsg:input_type -> helloworld.v1.DingRequest
	2, // 3: helloworld.v1.Greeter.ChatMsg:output_type -> helloworld.v1.DingReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_helloworld_v1_greeter_proto_init() }
func file_helloworld_v1_greeter_proto_init() {
	if File_helloworld_v1_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_v1_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingRequest); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingReply); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingRequest_User); i {
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
			RawDescriptor: file_helloworld_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_greeter_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_greeter_proto_depIdxs,
		MessageInfos:      file_helloworld_v1_greeter_proto_msgTypes,
	}.Build()
	File_helloworld_v1_greeter_proto = out.File
	file_helloworld_v1_greeter_proto_rawDesc = nil
	file_helloworld_v1_greeter_proto_goTypes = nil
	file_helloworld_v1_greeter_proto_depIdxs = nil
}
