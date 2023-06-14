// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: only_id_service.proto

package onlyIdSrv

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

type ReqId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizTag string `protobuf:"bytes,1,opt,name=BizTag,proto3" json:"BizTag,omitempty"`
}

func (x *ReqId) Reset() {
	*x = ReqId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_only_id_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqId) ProtoMessage() {}

func (x *ReqId) ProtoReflect() protoreflect.Message {
	mi := &file_only_id_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqId.ProtoReflect.Descriptor instead.
func (*ReqId) Descriptor() ([]byte, []int) {
	return file_only_id_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReqId) GetBizTag() string {
	if x != nil {
		return x.BizTag
	}
	return ""
}

type ResId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResId) Reset() {
	*x = ResId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_only_id_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResId) ProtoMessage() {}

func (x *ResId) ProtoReflect() protoreflect.Message {
	mi := &file_only_id_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResId.ProtoReflect.Descriptor instead.
func (*ResId) Descriptor() ([]byte, []int) {
	return file_only_id_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResId) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_only_id_service_proto protoreflect.FileDescriptor

var file_only_id_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x64, 0x53,
	0x72, 0x76, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1f, 0x0a, 0x05, 0x52, 0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x69, 0x7a, 0x54,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x69, 0x7a, 0x54, 0x61, 0x67,
	0x22, 0x31, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xa3, 0x01, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x05, 0x47, 0x65, 0x74, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x64,
	0x53, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x71, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x6f, 0x6e, 0x6c, 0x79,
	0x49, 0x64, 0x53, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x6e, 0x6f, 0x77, 0x46, 0x6c, 0x61, 0x6b, 0x65, 0x49, 0x64, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x64, 0x53, 0x72,
	0x76, 0x2e, 0x52, 0x65, 0x73, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x64, 0x53, 0x72,
	0x76, 0x2e, 0x52, 0x65, 0x71, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x6f, 0x6e, 0x6c, 0x79, 0x49, 0x64,
	0x53, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x49, 0x64, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6f,
	0x6e, 0x6c, 0x79, 0x49, 0x64, 0x53, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_only_id_service_proto_rawDescOnce sync.Once
	file_only_id_service_proto_rawDescData = file_only_id_service_proto_rawDesc
)

func file_only_id_service_proto_rawDescGZIP() []byte {
	file_only_id_service_proto_rawDescOnce.Do(func() {
		file_only_id_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_only_id_service_proto_rawDescData)
	})
	return file_only_id_service_proto_rawDescData
}

var file_only_id_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_only_id_service_proto_goTypes = []interface{}{
	(*ReqId)(nil),         // 0: onlyIdSrv.ReqId
	(*ResId)(nil),         // 1: onlyIdSrv.ResId
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_only_id_service_proto_depIdxs = []int32{
	0, // 0: onlyIdSrv.OnlyId.GetId:input_type -> onlyIdSrv.ReqId
	2, // 1: onlyIdSrv.OnlyId.GetSnowFlakeId:input_type -> google.protobuf.Empty
	0, // 2: onlyIdSrv.OnlyId.GetRedisId:input_type -> onlyIdSrv.ReqId
	1, // 3: onlyIdSrv.OnlyId.GetId:output_type -> onlyIdSrv.ResId
	1, // 4: onlyIdSrv.OnlyId.GetSnowFlakeId:output_type -> onlyIdSrv.ResId
	1, // 5: onlyIdSrv.OnlyId.GetRedisId:output_type -> onlyIdSrv.ResId
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_only_id_service_proto_init() }
func file_only_id_service_proto_init() {
	if File_only_id_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_only_id_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqId); i {
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
		file_only_id_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResId); i {
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
			RawDescriptor: file_only_id_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_only_id_service_proto_goTypes,
		DependencyIndexes: file_only_id_service_proto_depIdxs,
		MessageInfos:      file_only_id_service_proto_msgTypes,
	}.Build()
	File_only_id_service_proto = out.File
	file_only_id_service_proto_rawDesc = nil
	file_only_id_service_proto_goTypes = nil
	file_only_id_service_proto_depIdxs = nil
}
