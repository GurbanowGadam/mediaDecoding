// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mediaDecoding.proto

package mediaDecoding

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

type VideoCuttingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoPath string `protobuf:"bytes,1,opt,name=VideoPath,proto3" json:"VideoPath,omitempty"`
}

func (x *VideoCuttingRequest) Reset() {
	*x = VideoCuttingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaDecoding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoCuttingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoCuttingRequest) ProtoMessage() {}

func (x *VideoCuttingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaDecoding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoCuttingRequest.ProtoReflect.Descriptor instead.
func (*VideoCuttingRequest) Descriptor() ([]byte, []int) {
	return file_mediaDecoding_proto_rawDescGZIP(), []int{0}
}

func (x *VideoCuttingRequest) GetVideoPath() string {
	if x != nil {
		return x.VideoPath
	}
	return ""
}

type VideoCuttingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *VideoCuttingResponse) Reset() {
	*x = VideoCuttingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaDecoding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoCuttingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoCuttingResponse) ProtoMessage() {}

func (x *VideoCuttingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaDecoding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoCuttingResponse.ProtoReflect.Descriptor instead.
func (*VideoCuttingResponse) Descriptor() ([]byte, []int) {
	return file_mediaDecoding_proto_rawDescGZIP(), []int{1}
}

func (x *VideoCuttingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_mediaDecoding_proto protoreflect.FileDescriptor

var file_mediaDecoding_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x75, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x14, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x43, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x76, 0x0a, 0x14, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x75, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x43, 0x75, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x6f, 0x77, 0x47, 0x61, 0x64, 0x61, 0x6d, 0x2f, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mediaDecoding_proto_rawDescOnce sync.Once
	file_mediaDecoding_proto_rawDescData = file_mediaDecoding_proto_rawDesc
)

func file_mediaDecoding_proto_rawDescGZIP() []byte {
	file_mediaDecoding_proto_rawDescOnce.Do(func() {
		file_mediaDecoding_proto_rawDescData = protoimpl.X.CompressGZIP(file_mediaDecoding_proto_rawDescData)
	})
	return file_mediaDecoding_proto_rawDescData
}

var file_mediaDecoding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mediaDecoding_proto_goTypes = []interface{}{
	(*VideoCuttingRequest)(nil),  // 0: mediaDecoding.VideoCuttingRequest
	(*VideoCuttingResponse)(nil), // 1: mediaDecoding.VideoCuttingResponse
}
var file_mediaDecoding_proto_depIdxs = []int32{
	0, // 0: mediaDecoding.mediaDecodingService.StartVideoCutting:input_type -> mediaDecoding.VideoCuttingRequest
	1, // 1: mediaDecoding.mediaDecodingService.StartVideoCutting:output_type -> mediaDecoding.VideoCuttingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mediaDecoding_proto_init() }
func file_mediaDecoding_proto_init() {
	if File_mediaDecoding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mediaDecoding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoCuttingRequest); i {
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
		file_mediaDecoding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoCuttingResponse); i {
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
			RawDescriptor: file_mediaDecoding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mediaDecoding_proto_goTypes,
		DependencyIndexes: file_mediaDecoding_proto_depIdxs,
		MessageInfos:      file_mediaDecoding_proto_msgTypes,
	}.Build()
	File_mediaDecoding_proto = out.File
	file_mediaDecoding_proto_rawDesc = nil
	file_mediaDecoding_proto_goTypes = nil
	file_mediaDecoding_proto_depIdxs = nil
}
