// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: event.proto

package binaryCoder

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

type EventIdentifierProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventHash []byte `protobuf:"bytes,1,opt,name=event_hash,json=eventHash,proto3" json:"event_hash,omitempty"`
	EventType int32  `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	FastMeta  []byte `protobuf:"bytes,3,opt,name=fast_meta,json=fastMeta,proto3" json:"fast_meta,omitempty"`
}

func (x *EventIdentifierProto) Reset() {
	*x = EventIdentifierProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventIdentifierProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventIdentifierProto) ProtoMessage() {}

func (x *EventIdentifierProto) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventIdentifierProto.ProtoReflect.Descriptor instead.
func (*EventIdentifierProto) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventIdentifierProto) GetEventHash() []byte {
	if x != nil {
		return x.EventHash
	}
	return nil
}

func (x *EventIdentifierProto) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

func (x *EventIdentifierProto) GetFastMeta() []byte {
	if x != nil {
		return x.FastMeta
	}
	return nil
}

type EventProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventIdentifier *EventIdentifierProto `protobuf:"bytes,1,opt,name=event_identifier,json=eventIdentifier,proto3" json:"event_identifier,omitempty"`
	Level           int64                 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	ContentHashes   [][]byte              `protobuf:"bytes,3,rep,name=content_hashes,json=contentHashes,proto3" json:"content_hashes,omitempty"`
	MetadataHashes  [][]byte              `protobuf:"bytes,4,rep,name=metadata_hashes,json=metadataHashes,proto3" json:"metadata_hashes,omitempty"`
	ParentEvent     []byte                `protobuf:"bytes,5,opt,name=parent_event,json=parentEvent,proto3" json:"parent_event,omitempty"`
	RootEvent       []byte                `protobuf:"bytes,6,opt,name=root_event,json=rootEvent,proto3" json:"root_event,omitempty"`
	Temporary       bool                  `protobuf:"varint,7,opt,name=temporary,proto3" json:"temporary,omitempty"`
	FullTextSearch  bool                  `protobuf:"varint,8,opt,name=full_text_search,json=fullTextSearch,proto3" json:"full_text_search,omitempty"`
}

func (x *EventProto) Reset() {
	*x = EventProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventProto) ProtoMessage() {}

func (x *EventProto) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventProto.ProtoReflect.Descriptor instead.
func (*EventProto) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventProto) GetEventIdentifier() *EventIdentifierProto {
	if x != nil {
		return x.EventIdentifier
	}
	return nil
}

func (x *EventProto) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *EventProto) GetContentHashes() [][]byte {
	if x != nil {
		return x.ContentHashes
	}
	return nil
}

func (x *EventProto) GetMetadataHashes() [][]byte {
	if x != nil {
		return x.MetadataHashes
	}
	return nil
}

func (x *EventProto) GetParentEvent() []byte {
	if x != nil {
		return x.ParentEvent
	}
	return nil
}

func (x *EventProto) GetRootEvent() []byte {
	if x != nil {
		return x.RootEvent
	}
	return nil
}

func (x *EventProto) GetTemporary() bool {
	if x != nil {
		return x.Temporary
	}
	return false
}

func (x *EventProto) GetFullTextSearch() bool {
	if x != nil {
		return x.FullTextSearch
	}
	return false
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x14, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x22, 0xca, 0x02,
	0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4c, 0x0a, 0x10,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79,
	0x12, 0x28, 0x0a, 0x10, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x75, 0x6c, 0x6c,
	0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x35, 0x68, 0x65, 0x75, 0x2f, 0x6f,
	0x75, 0x72, 0x6f, 0x62, 0x6f, 0x72, 0x6f, 0x73, 0x2d, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_proto_goTypes = []interface{}{
	(*EventIdentifierProto)(nil), // 0: binaryCoder.EventIdentifierProto
	(*EventProto)(nil),           // 1: binaryCoder.EventProto
}
var file_event_proto_depIdxs = []int32{
	0, // 0: binaryCoder.EventProto.event_identifier:type_name -> binaryCoder.EventIdentifierProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventIdentifierProto); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventProto); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}