// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: link.proto

package links

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

type LinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *LinkRequest) Reset() {
	*x = LinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkRequest) ProtoMessage() {}

func (x *LinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkRequest.ProtoReflect.Descriptor instead.
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return file_link_proto_rawDescGZIP(), []int{0}
}

func (x *LinkRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type LinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *LinkResponse) Reset() {
	*x = LinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkResponse) ProtoMessage() {}

func (x *LinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkResponse.ProtoReflect.Descriptor instead.
func (*LinkResponse) Descriptor() ([]byte, []int) {
	return file_link_proto_rawDescGZIP(), []int{1}
}

func (x *LinkResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_link_proto protoreflect.FileDescriptor

var file_link_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x22, 0x0a,
	0x0c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x32, 0xb3, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x56, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x7b, 0x6c, 0x69, 0x6e, 0x6b, 0x7d, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_proto_rawDescOnce sync.Once
	file_link_proto_rawDescData = file_link_proto_rawDesc
)

func file_link_proto_rawDescGZIP() []byte {
	file_link_proto_rawDescOnce.Do(func() {
		file_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_proto_rawDescData)
	})
	return file_link_proto_rawDescData
}

var file_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_link_proto_goTypes = []interface{}{
	(*LinkRequest)(nil),  // 0: shortLink.LinkRequest
	(*LinkResponse)(nil), // 1: shortLink.LinkResponse
}
var file_link_proto_depIdxs = []int32{
	0, // 0: shortLink.Link.CreateShortLink:input_type -> shortLink.LinkRequest
	0, // 1: shortLink.Link.GetFullLink:input_type -> shortLink.LinkRequest
	1, // 2: shortLink.Link.CreateShortLink:output_type -> shortLink.LinkResponse
	1, // 3: shortLink.Link.GetFullLink:output_type -> shortLink.LinkResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_link_proto_init() }
func file_link_proto_init() {
	if File_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkRequest); i {
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
		file_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkResponse); i {
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
			RawDescriptor: file_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_proto_goTypes,
		DependencyIndexes: file_link_proto_depIdxs,
		MessageInfos:      file_link_proto_msgTypes,
	}.Build()
	File_link_proto = out.File
	file_link_proto_rawDesc = nil
	file_link_proto_goTypes = nil
	file_link_proto_depIdxs = nil
}