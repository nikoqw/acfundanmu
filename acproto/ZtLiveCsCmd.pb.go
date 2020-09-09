// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: ZtLiveCsCmd.proto

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

type ZtLiveCsCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdType string `protobuf:"bytes,1,opt,name=cmdType,proto3" json:"cmdType,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Ticket  string `protobuf:"bytes,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	LiveId  string `protobuf:"bytes,4,opt,name=liveId,proto3" json:"liveId,omitempty"`
}

func (x *ZtLiveCsCmd) Reset() {
	*x = ZtLiveCsCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveCsCmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveCsCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveCsCmd) ProtoMessage() {}

func (x *ZtLiveCsCmd) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveCsCmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveCsCmd.ProtoReflect.Descriptor instead.
func (*ZtLiveCsCmd) Descriptor() ([]byte, []int) {
	return file_ZtLiveCsCmd_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveCsCmd) GetCmdType() string {
	if x != nil {
		return x.CmdType
	}
	return ""
}

func (x *ZtLiveCsCmd) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ZtLiveCsCmd) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *ZtLiveCsCmd) GetLiveId() string {
	if x != nil {
		return x.LiveId
	}
	return ""
}

type ZtLiveCsCmdAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdAckType string `protobuf:"bytes,1,opt,name=cmdAckType,proto3" json:"cmdAckType,omitempty"`
	ErrorCode  int64  `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMsg   string `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	Payload    []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ZtLiveCsCmdAck) Reset() {
	*x = ZtLiveCsCmdAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveCsCmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveCsCmdAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveCsCmdAck) ProtoMessage() {}

func (x *ZtLiveCsCmdAck) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveCsCmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveCsCmdAck.ProtoReflect.Descriptor instead.
func (*ZtLiveCsCmdAck) Descriptor() ([]byte, []int) {
	return file_ZtLiveCsCmd_proto_rawDescGZIP(), []int{1}
}

func (x *ZtLiveCsCmdAck) GetCmdAckType() string {
	if x != nil {
		return x.CmdAckType
	}
	return ""
}

func (x *ZtLiveCsCmdAck) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ZtLiveCsCmdAck) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *ZtLiveCsCmdAck) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_ZtLiveCsCmd_proto protoreflect.FileDescriptor

var file_ZtLiveCsCmd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x73, 0x43, 0x6d, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22,
	0x71, 0x0a, 0x0b, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x73, 0x43, 0x6d, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69,
	0x76, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x76, 0x65,
	0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x73, 0x43,
	0x6d, 0x64, 0x41, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6d, 0x64, 0x41, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6d, 0x64, 0x41, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ZtLiveCsCmd_proto_rawDescOnce sync.Once
	file_ZtLiveCsCmd_proto_rawDescData = file_ZtLiveCsCmd_proto_rawDesc
)

func file_ZtLiveCsCmd_proto_rawDescGZIP() []byte {
	file_ZtLiveCsCmd_proto_rawDescOnce.Do(func() {
		file_ZtLiveCsCmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveCsCmd_proto_rawDescData)
	})
	return file_ZtLiveCsCmd_proto_rawDescData
}

var file_ZtLiveCsCmd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ZtLiveCsCmd_proto_goTypes = []interface{}{
	(*ZtLiveCsCmd)(nil),    // 0: AcFunDanmu.ZtLiveCsCmd
	(*ZtLiveCsCmdAck)(nil), // 1: AcFunDanmu.ZtLiveCsCmdAck
}
var file_ZtLiveCsCmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ZtLiveCsCmd_proto_init() }
func file_ZtLiveCsCmd_proto_init() {
	if File_ZtLiveCsCmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveCsCmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveCsCmd); i {
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
		file_ZtLiveCsCmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveCsCmdAck); i {
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
			RawDescriptor: file_ZtLiveCsCmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveCsCmd_proto_goTypes,
		DependencyIndexes: file_ZtLiveCsCmd_proto_depIdxs,
		MessageInfos:      file_ZtLiveCsCmd_proto_msgTypes,
	}.Build()
	File_ZtLiveCsCmd_proto = out.File
	file_ZtLiveCsCmd_proto_rawDesc = nil
	file_ZtLiveCsCmd_proto_goTypes = nil
	file_ZtLiveCsCmd_proto_depIdxs = nil
}
