// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: fastMeta.proto

package protoFastMeta

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

type FastMetaParameterProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameter []byte `protobuf:"bytes,1,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (x *FastMetaParameterProto) Reset() {
	*x = FastMetaParameterProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fastMeta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FastMetaParameterProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastMetaParameterProto) ProtoMessage() {}

func (x *FastMetaParameterProto) ProtoReflect() protoreflect.Message {
	mi := &file_fastMeta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastMetaParameterProto.ProtoReflect.Descriptor instead.
func (*FastMetaParameterProto) Descriptor() ([]byte, []int) {
	return file_fastMeta_proto_rawDescGZIP(), []int{0}
}

func (x *FastMetaParameterProto) GetParameter() []byte {
	if x != nil {
		return x.Parameter
	}
	return nil
}

type FastMetaProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters []*FastMetaParameterProto `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *FastMetaProto) Reset() {
	*x = FastMetaProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fastMeta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FastMetaProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastMetaProto) ProtoMessage() {}

func (x *FastMetaProto) ProtoReflect() protoreflect.Message {
	mi := &file_fastMeta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastMetaProto.ProtoReflect.Descriptor instead.
func (*FastMetaProto) Descriptor() ([]byte, []int) {
	return file_fastMeta_proto_rawDescGZIP(), []int{1}
}

func (x *FastMetaProto) GetParameters() []*FastMetaParameterProto {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_fastMeta_proto protoreflect.FileDescriptor

var file_fastMeta_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x22,
	0x36, 0x0a, 0x16, 0x46, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x0d, 0x46, 0x61, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x46, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x61, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x35,
	0x68, 0x65, 0x75, 0x2f, 0x6f, 0x75, 0x72, 0x6f, 0x62, 0x6f, 0x72, 0x6f, 0x73, 0x2d, 0x64, 0x62,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x46,
	0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fastMeta_proto_rawDescOnce sync.Once
	file_fastMeta_proto_rawDescData = file_fastMeta_proto_rawDesc
)

func file_fastMeta_proto_rawDescGZIP() []byte {
	file_fastMeta_proto_rawDescOnce.Do(func() {
		file_fastMeta_proto_rawDescData = protoimpl.X.CompressGZIP(file_fastMeta_proto_rawDescData)
	})
	return file_fastMeta_proto_rawDescData
}

var file_fastMeta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fastMeta_proto_goTypes = []interface{}{
	(*FastMetaParameterProto)(nil), // 0: protoFastMeta.FastMetaParameterProto
	(*FastMetaProto)(nil),          // 1: protoFastMeta.FastMetaProto
}
var file_fastMeta_proto_depIdxs = []int32{
	0, // 0: protoFastMeta.FastMetaProto.parameters:type_name -> protoFastMeta.FastMetaParameterProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fastMeta_proto_init() }
func file_fastMeta_proto_init() {
	if File_fastMeta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fastMeta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FastMetaParameterProto); i {
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
		file_fastMeta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FastMetaProto); i {
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
			RawDescriptor: file_fastMeta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fastMeta_proto_goTypes,
		DependencyIndexes: file_fastMeta_proto_depIdxs,
		MessageInfos:      file_fastMeta_proto_msgTypes,
	}.Build()
	File_fastMeta_proto = out.File
	file_fastMeta_proto_rawDesc = nil
	file_fastMeta_proto_goTypes = nil
	file_fastMeta_proto_depIdxs = nil
}