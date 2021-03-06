// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: UpstreamPayload.proto

package acproto

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

type UpstreamPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command          string            `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	SeqId            int64             `protobuf:"varint,2,opt,name=seqId,proto3" json:"seqId,omitempty"`
	RetryCount       uint32            `protobuf:"varint,3,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	PayloadData      []byte            `protobuf:"bytes,4,opt,name=payloadData,proto3" json:"payloadData,omitempty"`
	UserInstance     *UserInstance     `protobuf:"bytes,5,opt,name=userInstance,proto3" json:"userInstance,omitempty"`
	ErrorCode        int32             `protobuf:"varint,6,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	SettingInfo      *SettingInfo      `protobuf:"bytes,7,opt,name=settingInfo,proto3" json:"settingInfo,omitempty"`
	RequestBasicInfo *RequestBasicInfo `protobuf:"bytes,8,opt,name=requestBasicInfo,proto3" json:"requestBasicInfo,omitempty"`
	SubBiz           string            `protobuf:"bytes,9,opt,name=subBiz,proto3" json:"subBiz,omitempty"`
	FrontendInfo     *FrontendInfo     `protobuf:"bytes,10,opt,name=frontendInfo,proto3" json:"frontendInfo,omitempty"`
	Kpn              string            `protobuf:"bytes,11,opt,name=kpn,proto3" json:"kpn,omitempty"`
	AnonymouseUser   bool              `protobuf:"varint,12,opt,name=anonymouseUser,proto3" json:"anonymouseUser,omitempty"`
}

func (x *UpstreamPayload) Reset() {
	*x = UpstreamPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpstreamPayload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamPayload) ProtoMessage() {}

func (x *UpstreamPayload) ProtoReflect() protoreflect.Message {
	mi := &file_UpstreamPayload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamPayload.ProtoReflect.Descriptor instead.
func (*UpstreamPayload) Descriptor() ([]byte, []int) {
	return file_UpstreamPayload_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamPayload) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *UpstreamPayload) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *UpstreamPayload) GetRetryCount() uint32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *UpstreamPayload) GetPayloadData() []byte {
	if x != nil {
		return x.PayloadData
	}
	return nil
}

func (x *UpstreamPayload) GetUserInstance() *UserInstance {
	if x != nil {
		return x.UserInstance
	}
	return nil
}

func (x *UpstreamPayload) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *UpstreamPayload) GetSettingInfo() *SettingInfo {
	if x != nil {
		return x.SettingInfo
	}
	return nil
}

func (x *UpstreamPayload) GetRequestBasicInfo() *RequestBasicInfo {
	if x != nil {
		return x.RequestBasicInfo
	}
	return nil
}

func (x *UpstreamPayload) GetSubBiz() string {
	if x != nil {
		return x.SubBiz
	}
	return ""
}

func (x *UpstreamPayload) GetFrontendInfo() *FrontendInfo {
	if x != nil {
		return x.FrontendInfo
	}
	return nil
}

func (x *UpstreamPayload) GetKpn() string {
	if x != nil {
		return x.Kpn
	}
	return ""
}

func (x *UpstreamPayload) GetAnonymouseUser() bool {
	if x != nil {
		return x.AnonymouseUser
	}
	return false
}

var File_UpstreamPayload_proto protoreflect.FileDescriptor

var file_UpstreamPayload_proto_rawDesc = []byte{
	0x0a, 0x15, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x1a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x03, 0x0a, 0x0f, 0x55, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x62, 0x42, 0x69, 0x7a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x75, 0x62, 0x42, 0x69, 0x7a, 0x12, 0x3c, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x70, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_UpstreamPayload_proto_rawDescOnce sync.Once
	file_UpstreamPayload_proto_rawDescData = file_UpstreamPayload_proto_rawDesc
)

func file_UpstreamPayload_proto_rawDescGZIP() []byte {
	file_UpstreamPayload_proto_rawDescOnce.Do(func() {
		file_UpstreamPayload_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpstreamPayload_proto_rawDescData)
	})
	return file_UpstreamPayload_proto_rawDescData
}

var file_UpstreamPayload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpstreamPayload_proto_goTypes = []interface{}{
	(*UpstreamPayload)(nil),  // 0: AcFunDanmu.UpstreamPayload
	(*UserInstance)(nil),     // 1: AcFunDanmu.UserInstance
	(*SettingInfo)(nil),      // 2: AcFunDanmu.SettingInfo
	(*RequestBasicInfo)(nil), // 3: AcFunDanmu.RequestBasicInfo
	(*FrontendInfo)(nil),     // 4: AcFunDanmu.FrontendInfo
}
var file_UpstreamPayload_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.UpstreamPayload.userInstance:type_name -> AcFunDanmu.UserInstance
	2, // 1: AcFunDanmu.UpstreamPayload.settingInfo:type_name -> AcFunDanmu.SettingInfo
	3, // 2: AcFunDanmu.UpstreamPayload.requestBasicInfo:type_name -> AcFunDanmu.RequestBasicInfo
	4, // 3: AcFunDanmu.UpstreamPayload.frontendInfo:type_name -> AcFunDanmu.FrontendInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_UpstreamPayload_proto_init() }
func file_UpstreamPayload_proto_init() {
	if File_UpstreamPayload_proto != nil {
		return
	}
	file_UserInstance_proto_init()
	file_SettingInfo_proto_init()
	file_RequestBasicInfo_proto_init()
	file_FrontendInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpstreamPayload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamPayload); i {
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
			RawDescriptor: file_UpstreamPayload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpstreamPayload_proto_goTypes,
		DependencyIndexes: file_UpstreamPayload_proto_depIdxs,
		MessageInfos:      file_UpstreamPayload_proto_msgTypes,
	}.Build()
	File_UpstreamPayload_proto = out.File
	file_UpstreamPayload_proto_rawDesc = nil
	file_UpstreamPayload_proto_goTypes = nil
	file_UpstreamPayload_proto_depIdxs = nil
}
