// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/myproto.proto

package proto

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

// The request message containing the user's name.
type GRPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GRPCRequest) Reset() {
	*x = GRPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_myproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCRequest) ProtoMessage() {}

func (x *GRPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_myproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCRequest.ProtoReflect.Descriptor instead.
func (*GRPCRequest) Descriptor() ([]byte, []int) {
	return file_proto_myproto_proto_rawDescGZIP(), []int{0}
}

func (x *GRPCRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GRPCRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GRPCRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type GRPCReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccesOrNot string `protobuf:"bytes,1,opt,name=succesOrNot,proto3" json:"succesOrNot,omitempty"`
}

func (x *GRPCReply) Reset() {
	*x = GRPCReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_myproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCReply) ProtoMessage() {}

func (x *GRPCReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_myproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCReply.ProtoReflect.Descriptor instead.
func (*GRPCReply) Descriptor() ([]byte, []int) {
	return file_proto_myproto_proto_rawDescGZIP(), []int{1}
}

func (x *GRPCReply) GetSuccesOrNot() string {
	if x != nil {
		return x.SuccesOrNot
	}
	return ""
}

var File_proto_myproto_proto protoreflect.FileDescriptor

var file_proto_myproto_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x0b, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x09, 0x47, 0x52, 0x50,
	0x43, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x4f, 0x72, 0x4e, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x4f, 0x72, 0x4e, 0x6f, 0x74, 0x32, 0x88, 0x01, 0x0a, 0x15, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x56, 0x32, 0x72, 0x61, 0x79, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x67, 0x52,
	0x50, 0x43, 0x12, 0x35, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x38, 0x30, 0x31, 0x33, 0x2f, 0x6c, 0x6f, 0x67,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x66, 0x75, 0x6c, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_myproto_proto_rawDescOnce sync.Once
	file_proto_myproto_proto_rawDescData = file_proto_myproto_proto_rawDesc
)

func file_proto_myproto_proto_rawDescGZIP() []byte {
	file_proto_myproto_proto_rawDescOnce.Do(func() {
		file_proto_myproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_myproto_proto_rawDescData)
	})
	return file_proto_myproto_proto_rawDescData
}

var file_proto_myproto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_myproto_proto_goTypes = []interface{}{
	(*GRPCRequest)(nil), // 0: myproto.GRPCRequest
	(*GRPCReply)(nil),   // 1: myproto.GRPCReply
}
var file_proto_myproto_proto_depIdxs = []int32{
	0, // 0: myproto.manageV2rayUserBygRPC.AddUser:input_type -> myproto.GRPCRequest
	0, // 1: myproto.manageV2rayUserBygRPC.DeleteUser:input_type -> myproto.GRPCRequest
	1, // 2: myproto.manageV2rayUserBygRPC.AddUser:output_type -> myproto.GRPCReply
	1, // 3: myproto.manageV2rayUserBygRPC.DeleteUser:output_type -> myproto.GRPCReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_myproto_proto_init() }
func file_proto_myproto_proto_init() {
	if File_proto_myproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_myproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCRequest); i {
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
		file_proto_myproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCReply); i {
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
			RawDescriptor: file_proto_myproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_myproto_proto_goTypes,
		DependencyIndexes: file_proto_myproto_proto_depIdxs,
		MessageInfos:      file_proto_myproto_proto_msgTypes,
	}.Build()
	File_proto_myproto_proto = out.File
	file_proto_myproto_proto_rawDesc = nil
	file_proto_myproto_proto_goTypes = nil
	file_proto_myproto_proto_depIdxs = nil
}
