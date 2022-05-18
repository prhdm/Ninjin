// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: messages/warning.proto

package warning

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type WarningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WarningRequest) Reset() {
	*x = WarningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_warning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarningRequest) ProtoMessage() {}

func (x *WarningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_warning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarningRequest.ProtoReflect.Descriptor instead.
func (*WarningRequest) Descriptor() ([]byte, []int) {
	return file_messages_warning_proto_rawDescGZIP(), []int{0}
}

type WarningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warnings []*Warning `protobuf:"bytes,1,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *WarningResponse) Reset() {
	*x = WarningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_warning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarningResponse) ProtoMessage() {}

func (x *WarningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_warning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarningResponse.ProtoReflect.Descriptor instead.
func (*WarningResponse) Descriptor() ([]byte, []int) {
	return file_messages_warning_proto_rawDescGZIP(), []int{1}
}

func (x *WarningResponse) GetWarnings() []*Warning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

type Warning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerial string               `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	Difference   float32              `protobuf:"fixed32,2,opt,name=difference,proto3" json:"difference,omitempty"`
	Time         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Warning) Reset() {
	*x = Warning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_warning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warning) ProtoMessage() {}

func (x *Warning) ProtoReflect() protoreflect.Message {
	mi := &file_messages_warning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warning.ProtoReflect.Descriptor instead.
func (*Warning) Descriptor() ([]byte, []int) {
	return file_messages_warning_proto_rawDescGZIP(), []int{2}
}

func (x *Warning) GetDeviceSerial() string {
	if x != nil {
		return x.DeviceSerial
	}
	return ""
}

func (x *Warning) GetDifference() float32 {
	if x != nil {
		return x.Difference
	}
	return 0
}

func (x *Warning) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_messages_warning_proto protoreflect.FileDescriptor

var file_messages_warning_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x57,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a,
	0x0f, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x77,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x7e, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x64, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x66, 0x61, 0x72, 0x6d, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_messages_warning_proto_rawDescOnce sync.Once
	file_messages_warning_proto_rawDescData = file_messages_warning_proto_rawDesc
)

func file_messages_warning_proto_rawDescGZIP() []byte {
	file_messages_warning_proto_rawDescOnce.Do(func() {
		file_messages_warning_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_warning_proto_rawDescData)
	})
	return file_messages_warning_proto_rawDescData
}

var file_messages_warning_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_warning_proto_goTypes = []interface{}{
	(*WarningRequest)(nil),      // 0: messages.warning.WarningRequest
	(*WarningResponse)(nil),     // 1: messages.warning.WarningResponse
	(*Warning)(nil),             // 2: messages.warning.Warning
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_messages_warning_proto_depIdxs = []int32{
	2, // 0: messages.warning.WarningResponse.warnings:type_name -> messages.warning.Warning
	3, // 1: messages.warning.Warning.time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_warning_proto_init() }
func file_messages_warning_proto_init() {
	if File_messages_warning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_warning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarningRequest); i {
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
		file_messages_warning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarningResponse); i {
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
		file_messages_warning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warning); i {
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
			RawDescriptor: file_messages_warning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_warning_proto_goTypes,
		DependencyIndexes: file_messages_warning_proto_depIdxs,
		MessageInfos:      file_messages_warning_proto_msgTypes,
	}.Build()
	File_messages_warning_proto = out.File
	file_messages_warning_proto_rawDesc = nil
	file_messages_warning_proto_goTypes = nil
	file_messages_warning_proto_depIdxs = nil
}
