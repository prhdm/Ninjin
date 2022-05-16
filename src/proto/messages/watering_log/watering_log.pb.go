// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: messages/watering_log.proto

package watering_log

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

type CreateWateringLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	DeviceSerial string               `protobuf:"bytes,2,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	WaterAmount  float32              `protobuf:"fixed32,3,opt,name=water_amount,json=waterAmount,proto3" json:"water_amount,omitempty"`
}

func (x *CreateWateringLogRequest) Reset() {
	*x = CreateWateringLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_watering_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWateringLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWateringLogRequest) ProtoMessage() {}

func (x *CreateWateringLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_watering_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWateringLogRequest.ProtoReflect.Descriptor instead.
func (*CreateWateringLogRequest) Descriptor() ([]byte, []int) {
	return file_messages_watering_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWateringLogRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *CreateWateringLogRequest) GetDeviceSerial() string {
	if x != nil {
		return x.DeviceSerial
	}
	return ""
}

func (x *CreateWateringLogRequest) GetWaterAmount() float32 {
	if x != nil {
		return x.WaterAmount
	}
	return 0
}

type CreateWateringLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CreateWateringLogResponse) Reset() {
	*x = CreateWateringLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_watering_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWateringLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWateringLogResponse) ProtoMessage() {}

func (x *CreateWateringLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_watering_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWateringLogResponse.ProtoReflect.Descriptor instead.
func (*CreateWateringLogResponse) Descriptor() ([]byte, []int) {
	return file_messages_watering_log_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWateringLogResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreateWateringLogResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_messages_watering_log_proto protoreflect.FileDescriptor

var file_messages_watering_log_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x77, 0x61, 0x74, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x58, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x66,
	0x61, 0x72, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_watering_log_proto_rawDescOnce sync.Once
	file_messages_watering_log_proto_rawDescData = file_messages_watering_log_proto_rawDesc
)

func file_messages_watering_log_proto_rawDescGZIP() []byte {
	file_messages_watering_log_proto_rawDescOnce.Do(func() {
		file_messages_watering_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_watering_log_proto_rawDescData)
	})
	return file_messages_watering_log_proto_rawDescData
}

var file_messages_watering_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_watering_log_proto_goTypes = []interface{}{
	(*CreateWateringLogRequest)(nil),  // 0: messages.log.CreateWateringLogRequest
	(*CreateWateringLogResponse)(nil), // 1: messages.log.CreateWateringLogResponse
	(*timestamp.Timestamp)(nil),       // 2: google.protobuf.Timestamp
}
var file_messages_watering_log_proto_depIdxs = []int32{
	2, // 0: messages.log.CreateWateringLogRequest.time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_watering_log_proto_init() }
func file_messages_watering_log_proto_init() {
	if File_messages_watering_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_watering_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWateringLogRequest); i {
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
		file_messages_watering_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWateringLogResponse); i {
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
			RawDescriptor: file_messages_watering_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_watering_log_proto_goTypes,
		DependencyIndexes: file_messages_watering_log_proto_depIdxs,
		MessageInfos:      file_messages_watering_log_proto_msgTypes,
	}.Build()
	File_messages_watering_log_proto = out.File
	file_messages_watering_log_proto_rawDesc = nil
	file_messages_watering_log_proto_goTypes = nil
	file_messages_watering_log_proto_depIdxs = nil
}
