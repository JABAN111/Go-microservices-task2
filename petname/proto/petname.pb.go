// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: proto/petname.proto

package petname

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PetnameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Words         int64                  `protobuf:"varint,1,opt,name=words,proto3" json:"words,omitempty"`
	Separator     string                 `protobuf:"bytes,2,opt,name=separator,proto3" json:"separator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PetnameRequest) Reset() {
	*x = PetnameRequest{}
	mi := &file_proto_petname_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PetnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetnameRequest) ProtoMessage() {}

func (x *PetnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_petname_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetnameRequest.ProtoReflect.Descriptor instead.
func (*PetnameRequest) Descriptor() ([]byte, []int) {
	return file_proto_petname_proto_rawDescGZIP(), []int{0}
}

func (x *PetnameRequest) GetWords() int64 {
	if x != nil {
		return x.Words
	}
	return 0
}

func (x *PetnameRequest) GetSeparator() string {
	if x != nil {
		return x.Separator
	}
	return ""
}

type PetnameStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Words         int64                  `protobuf:"varint,1,opt,name=words,proto3" json:"words,omitempty"`
	Separator     string                 `protobuf:"bytes,2,opt,name=separator,proto3" json:"separator,omitempty"`
	Names         int64                  `protobuf:"varint,3,opt,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PetnameStreamRequest) Reset() {
	*x = PetnameStreamRequest{}
	mi := &file_proto_petname_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PetnameStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetnameStreamRequest) ProtoMessage() {}

func (x *PetnameStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_petname_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetnameStreamRequest.ProtoReflect.Descriptor instead.
func (*PetnameStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_petname_proto_rawDescGZIP(), []int{1}
}

func (x *PetnameStreamRequest) GetWords() int64 {
	if x != nil {
		return x.Words
	}
	return 0
}

func (x *PetnameStreamRequest) GetSeparator() string {
	if x != nil {
		return x.Separator
	}
	return ""
}

func (x *PetnameStreamRequest) GetNames() int64 {
	if x != nil {
		return x.Names
	}
	return 0
}

type PetnameResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PetnameResponse) Reset() {
	*x = PetnameResponse{}
	mi := &file_proto_petname_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PetnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetnameResponse) ProtoMessage() {}

func (x *PetnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_petname_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetnameResponse.ProtoReflect.Descriptor instead.
func (*PetnameResponse) Descriptor() ([]byte, []int) {
	return file_proto_petname_proto_rawDescGZIP(), []int{2}
}

func (x *PetnameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_petname_proto protoreflect.FileDescriptor

var file_proto_petname_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0e, 0x50,
	0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x60, 0x0a, 0x14, 0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x10, 0x50,
	0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x38, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x2e,
	0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x65, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x50, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x79, 0x61, 0x64, 0x72, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_petname_proto_rawDescOnce sync.Once
	file_proto_petname_proto_rawDescData = file_proto_petname_proto_rawDesc
)

func file_proto_petname_proto_rawDescGZIP() []byte {
	file_proto_petname_proto_rawDescOnce.Do(func() {
		file_proto_petname_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_petname_proto_rawDescData)
	})
	return file_proto_petname_proto_rawDescData
}

var file_proto_petname_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_petname_proto_goTypes = []any{
	(*PetnameRequest)(nil),       // 0: petname.PetnameRequest
	(*PetnameStreamRequest)(nil), // 1: petname.PetnameStreamRequest
	(*PetnameResponse)(nil),      // 2: petname.PetnameResponse
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_proto_petname_proto_depIdxs = []int32{
	3, // 0: petname.PetnameGenerator.Ping:input_type -> google.protobuf.Empty
	0, // 1: petname.PetnameGenerator.Generate:input_type -> petname.PetnameRequest
	1, // 2: petname.PetnameGenerator.GenerateMany:input_type -> petname.PetnameStreamRequest
	3, // 3: petname.PetnameGenerator.Ping:output_type -> google.protobuf.Empty
	2, // 4: petname.PetnameGenerator.Generate:output_type -> petname.PetnameResponse
	2, // 5: petname.PetnameGenerator.GenerateMany:output_type -> petname.PetnameResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_petname_proto_init() }
func file_proto_petname_proto_init() {
	if File_proto_petname_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_petname_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_petname_proto_goTypes,
		DependencyIndexes: file_proto_petname_proto_depIdxs,
		MessageInfos:      file_proto_petname_proto_msgTypes,
	}.Build()
	File_proto_petname_proto = out.File
	file_proto_petname_proto_rawDesc = nil
	file_proto_petname_proto_goTypes = nil
	file_proto_petname_proto_depIdxs = nil
}
