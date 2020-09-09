// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: CommonStateSignalChatEnd.proto

package acproto

import (
	proto "github.com/golang/protobuf/proto"
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

type CommonStateSignalChatEnd_EndType int32

const (
	CommonStateSignalChatEnd_UNKNOWN                  CommonStateSignalChatEnd_EndType = 0
	CommonStateSignalChatEnd_CANCEL_BY_AUTHOR         CommonStateSignalChatEnd_EndType = 1
	CommonStateSignalChatEnd_END_BY_AUTHOR            CommonStateSignalChatEnd_EndType = 2
	CommonStateSignalChatEnd_END_BY_GUEST             CommonStateSignalChatEnd_EndType = 3
	CommonStateSignalChatEnd_GUEST_REJECT             CommonStateSignalChatEnd_EndType = 4
	CommonStateSignalChatEnd_GUEST_TIMEOUT            CommonStateSignalChatEnd_EndType = 5
	CommonStateSignalChatEnd_GUEST_HEARTBEAT_TIMEOUT  CommonStateSignalChatEnd_EndType = 6
	CommonStateSignalChatEnd_AUTHOR_HEARTBEAT_TIMEOUT CommonStateSignalChatEnd_EndType = 7
)

// Enum value maps for CommonStateSignalChatEnd_EndType.
var (
	CommonStateSignalChatEnd_EndType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CANCEL_BY_AUTHOR",
		2: "END_BY_AUTHOR",
		3: "END_BY_GUEST",
		4: "GUEST_REJECT",
		5: "GUEST_TIMEOUT",
		6: "GUEST_HEARTBEAT_TIMEOUT",
		7: "AUTHOR_HEARTBEAT_TIMEOUT",
	}
	CommonStateSignalChatEnd_EndType_value = map[string]int32{
		"UNKNOWN":                  0,
		"CANCEL_BY_AUTHOR":         1,
		"END_BY_AUTHOR":            2,
		"END_BY_GUEST":             3,
		"GUEST_REJECT":             4,
		"GUEST_TIMEOUT":            5,
		"GUEST_HEARTBEAT_TIMEOUT":  6,
		"AUTHOR_HEARTBEAT_TIMEOUT": 7,
	}
)

func (x CommonStateSignalChatEnd_EndType) Enum() *CommonStateSignalChatEnd_EndType {
	p := new(CommonStateSignalChatEnd_EndType)
	*p = x
	return p
}

func (x CommonStateSignalChatEnd_EndType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonStateSignalChatEnd_EndType) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonStateSignalChatEnd_proto_enumTypes[0].Descriptor()
}

func (CommonStateSignalChatEnd_EndType) Type() protoreflect.EnumType {
	return &file_CommonStateSignalChatEnd_proto_enumTypes[0]
}

func (x CommonStateSignalChatEnd_EndType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonStateSignalChatEnd_EndType.Descriptor instead.
func (CommonStateSignalChatEnd_EndType) EnumDescriptor() ([]byte, []int) {
	return file_CommonStateSignalChatEnd_proto_rawDescGZIP(), []int{0, 0}
}

type CommonStateSignalChatEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  string                           `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	EndType CommonStateSignalChatEnd_EndType `protobuf:"varint,2,opt,name=endType,proto3,enum=AcFunDanmu.CommonStateSignalChatEnd_EndType" json:"endType,omitempty"`
}

func (x *CommonStateSignalChatEnd) Reset() {
	*x = CommonStateSignalChatEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalChatEnd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalChatEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalChatEnd) ProtoMessage() {}

func (x *CommonStateSignalChatEnd) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalChatEnd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalChatEnd.ProtoReflect.Descriptor instead.
func (*CommonStateSignalChatEnd) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalChatEnd_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalChatEnd) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *CommonStateSignalChatEnd) GetEndType() CommonStateSignalChatEnd_EndType {
	if x != nil {
		return x.EndType
	}
	return CommonStateSignalChatEnd_UNKNOWN
}

var File_CommonStateSignalChatEnd_proto protoreflect.FileDescriptor

var file_CommonStateSignalChatEnd_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xae, 0x02, 0x0a,
	0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x64, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x42, 0x59, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x44, 0x5f,
	0x42, 0x59, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45,
	0x4e, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x52,
	0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x06, 0x12,
	0x1c, 0x0a, 0x18, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42,
	0x45, 0x41, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x07, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalChatEnd_proto_rawDescOnce sync.Once
	file_CommonStateSignalChatEnd_proto_rawDescData = file_CommonStateSignalChatEnd_proto_rawDesc
)

func file_CommonStateSignalChatEnd_proto_rawDescGZIP() []byte {
	file_CommonStateSignalChatEnd_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalChatEnd_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalChatEnd_proto_rawDescData)
	})
	return file_CommonStateSignalChatEnd_proto_rawDescData
}

var file_CommonStateSignalChatEnd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonStateSignalChatEnd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalChatEnd_proto_goTypes = []interface{}{
	(CommonStateSignalChatEnd_EndType)(0), // 0: AcFunDanmu.CommonStateSignalChatEnd.EndType
	(*CommonStateSignalChatEnd)(nil),      // 1: AcFunDanmu.CommonStateSignalChatEnd
}
var file_CommonStateSignalChatEnd_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonStateSignalChatEnd.endType:type_name -> AcFunDanmu.CommonStateSignalChatEnd.EndType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalChatEnd_proto_init() }
func file_CommonStateSignalChatEnd_proto_init() {
	if File_CommonStateSignalChatEnd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalChatEnd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalChatEnd); i {
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
			RawDescriptor: file_CommonStateSignalChatEnd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalChatEnd_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalChatEnd_proto_depIdxs,
		EnumInfos:         file_CommonStateSignalChatEnd_proto_enumTypes,
		MessageInfos:      file_CommonStateSignalChatEnd_proto_msgTypes,
	}.Build()
	File_CommonStateSignalChatEnd_proto = out.File
	file_CommonStateSignalChatEnd_proto_rawDesc = nil
	file_CommonStateSignalChatEnd_proto_goTypes = nil
	file_CommonStateSignalChatEnd_proto_depIdxs = nil
}
