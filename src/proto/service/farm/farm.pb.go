// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: services/farm.proto

package farm

import (
	context "context"
	delete_device "farm/src/proto/messages/delete_device"
	device "farm/src/proto/messages/device"
	device_list "farm/src/proto/messages/device_list"
	device_log "farm/src/proto/messages/device_log"
	set_device_humidity "farm/src/proto/messages/set_device_humidity"
	update_device_humidity "farm/src/proto/messages/update_device_humidity"
	user "farm/src/proto/messages/user"
	warning "farm/src/proto/messages/warning"
	watering_log "farm/src/proto/messages/watering_log"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_farm_proto protoreflect.FileDescriptor

var file_services_farm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x86, 0x0b, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x55, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x22, 0x07, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8b,
	0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x15, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4c,
	0x6f, 0x67, 0x12, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x2d, 0x6e, 0x61,
	0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xb7, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x3c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75, 0x6d,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x3a, 0x01, 0x2a, 0x12,
	0xa5, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75, 0x6d,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x36, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1d, 0x5a, 0x1b, 0x66,
	0x61, 0x72, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_services_farm_proto_goTypes = []interface{}{
	(*user.LoginRequest)(nil),                                   // 0: messages.user.LoginRequest
	(*device.CreateDeviceRequest)(nil),                          // 1: messages.device.CreateDeviceRequest
	(*device_log.GetDeviceLogRequest)(nil),                      // 2: messages.device_log.GetDeviceLogRequest
	(*device_list.GetDeviceListRequest)(nil),                    // 3: messages.device_list.GetDeviceListRequest
	(*user.LogoutRequest)(nil),                                  // 4: messages.user.LogoutRequest
	(*delete_device.DeleteDeviceRequest)(nil),                   // 5: messages.delete_device.DeleteDeviceRequest
	(*watering_log.CreateWateringLogRequest)(nil),               // 6: messages.log.CreateWateringLogRequest
	(*device.EditDeviceNameRequest)(nil),                        // 7: messages.device.EditDeviceNameRequest
	(*update_device_humidity.UpdateDeviceHumidityRequest)(nil),  // 8: messages.update_device_humidity.UpdateDeviceHumidityRequest
	(*set_device_humidity.SetDeviceHumidityRequest)(nil),        // 9: messages.set_device_humidity.SetDeviceHumidityRequest
	(*warning.WarningRequest)(nil),                              // 10: messages.warning.WarningRequest
	(*user.LoginResponse)(nil),                                  // 11: messages.user.LoginResponse
	(*device.CreateDeviceResponse)(nil),                         // 12: messages.device.CreateDeviceResponse
	(*device_log.GetDeviceLogResponse)(nil),                     // 13: messages.device_log.GetDeviceLogResponse
	(*device_list.GetDeviceListResponse)(nil),                   // 14: messages.device_list.GetDeviceListResponse
	(*user.LogoutResponse)(nil),                                 // 15: messages.user.LogoutResponse
	(*delete_device.DeleteDeviceResponse)(nil),                  // 16: messages.delete_device.DeleteDeviceResponse
	(*watering_log.CreateWateringLogResponse)(nil),              // 17: messages.log.CreateWateringLogResponse
	(*device.EditDeviceNameResponse)(nil),                       // 18: messages.device.EditDeviceNameResponse
	(*update_device_humidity.UpdateDeviceHumidityResponse)(nil), // 19: messages.update_device_humidity.UpdateDeviceHumidityResponse
	(*set_device_humidity.SetDeviceHumidityResponse)(nil),       // 20: messages.set_device_humidity.SetDeviceHumidityResponse
	(*warning.WarningResponse)(nil),                             // 21: messages.warning.WarningResponse
}
var file_services_farm_proto_depIdxs = []int32{
	0,  // 0: service.farm.Farm.Login:input_type -> messages.user.LoginRequest
	1,  // 1: service.farm.Farm.CreateDevice:input_type -> messages.device.CreateDeviceRequest
	2,  // 2: service.farm.Farm.GetDeviceLog:input_type -> messages.device_log.GetDeviceLogRequest
	3,  // 3: service.farm.Farm.GetDeviceList:input_type -> messages.device_list.GetDeviceListRequest
	4,  // 4: service.farm.Farm.Logout:input_type -> messages.user.LogoutRequest
	5,  // 5: service.farm.Farm.DeleteDevice:input_type -> messages.delete_device.DeleteDeviceRequest
	6,  // 6: service.farm.Farm.CreateWateringLog:input_type -> messages.log.CreateWateringLogRequest
	7,  // 7: service.farm.Farm.EditDeviceName:input_type -> messages.device.EditDeviceNameRequest
	8,  // 8: service.farm.Farm.UpdateDeviceHumidity:input_type -> messages.update_device_humidity.UpdateDeviceHumidityRequest
	9,  // 9: service.farm.Farm.SetDeviceHumidity:input_type -> messages.set_device_humidity.SetDeviceHumidityRequest
	10, // 10: service.farm.Farm.Warning:input_type -> messages.warning.WarningRequest
	11, // 11: service.farm.Farm.Login:output_type -> messages.user.LoginResponse
	12, // 12: service.farm.Farm.CreateDevice:output_type -> messages.device.CreateDeviceResponse
	13, // 13: service.farm.Farm.GetDeviceLog:output_type -> messages.device_log.GetDeviceLogResponse
	14, // 14: service.farm.Farm.GetDeviceList:output_type -> messages.device_list.GetDeviceListResponse
	15, // 15: service.farm.Farm.Logout:output_type -> messages.user.LogoutResponse
	16, // 16: service.farm.Farm.DeleteDevice:output_type -> messages.delete_device.DeleteDeviceResponse
	17, // 17: service.farm.Farm.CreateWateringLog:output_type -> messages.log.CreateWateringLogResponse
	18, // 18: service.farm.Farm.EditDeviceName:output_type -> messages.device.EditDeviceNameResponse
	19, // 19: service.farm.Farm.UpdateDeviceHumidity:output_type -> messages.update_device_humidity.UpdateDeviceHumidityResponse
	20, // 20: service.farm.Farm.SetDeviceHumidity:output_type -> messages.set_device_humidity.SetDeviceHumidityResponse
	21, // 21: service.farm.Farm.Warning:output_type -> messages.warning.WarningResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_farm_proto_init() }
func file_services_farm_proto_init() {
	if File_services_farm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_farm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_farm_proto_goTypes,
		DependencyIndexes: file_services_farm_proto_depIdxs,
	}.Build()
	File_services_farm_proto = out.File
	file_services_farm_proto_rawDesc = nil
	file_services_farm_proto_goTypes = nil
	file_services_farm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FarmClient is the client API for Farm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FarmClient interface {
	Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error)
	CreateDevice(ctx context.Context, in *device.CreateDeviceRequest, opts ...grpc.CallOption) (*device.CreateDeviceResponse, error)
	GetDeviceLog(ctx context.Context, in *device_log.GetDeviceLogRequest, opts ...grpc.CallOption) (*device_log.GetDeviceLogResponse, error)
	GetDeviceList(ctx context.Context, in *device_list.GetDeviceListRequest, opts ...grpc.CallOption) (*device_list.GetDeviceListResponse, error)
	Logout(ctx context.Context, in *user.LogoutRequest, opts ...grpc.CallOption) (*user.LogoutResponse, error)
	DeleteDevice(ctx context.Context, in *delete_device.DeleteDeviceRequest, opts ...grpc.CallOption) (*delete_device.DeleteDeviceResponse, error)
	CreateWateringLog(ctx context.Context, in *watering_log.CreateWateringLogRequest, opts ...grpc.CallOption) (*watering_log.CreateWateringLogResponse, error)
	EditDeviceName(ctx context.Context, in *device.EditDeviceNameRequest, opts ...grpc.CallOption) (*device.EditDeviceNameResponse, error)
	UpdateDeviceHumidity(ctx context.Context, in *update_device_humidity.UpdateDeviceHumidityRequest, opts ...grpc.CallOption) (*update_device_humidity.UpdateDeviceHumidityResponse, error)
	SetDeviceHumidity(ctx context.Context, in *set_device_humidity.SetDeviceHumidityRequest, opts ...grpc.CallOption) (*set_device_humidity.SetDeviceHumidityResponse, error)
	Warning(ctx context.Context, in *warning.WarningRequest, opts ...grpc.CallOption) (Farm_WarningClient, error)
}

type farmClient struct {
	cc grpc.ClientConnInterface
}

func NewFarmClient(cc grpc.ClientConnInterface) FarmClient {
	return &farmClient{cc}
}

func (c *farmClient) Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error) {
	out := new(user.LoginResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) CreateDevice(ctx context.Context, in *device.CreateDeviceRequest, opts ...grpc.CallOption) (*device.CreateDeviceResponse, error) {
	out := new(device.CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) GetDeviceLog(ctx context.Context, in *device_log.GetDeviceLogRequest, opts ...grpc.CallOption) (*device_log.GetDeviceLogResponse, error) {
	out := new(device_log.GetDeviceLogResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/GetDeviceLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) GetDeviceList(ctx context.Context, in *device_list.GetDeviceListRequest, opts ...grpc.CallOption) (*device_list.GetDeviceListResponse, error) {
	out := new(device_list.GetDeviceListResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/GetDeviceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) Logout(ctx context.Context, in *user.LogoutRequest, opts ...grpc.CallOption) (*user.LogoutResponse, error) {
	out := new(user.LogoutResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) DeleteDevice(ctx context.Context, in *delete_device.DeleteDeviceRequest, opts ...grpc.CallOption) (*delete_device.DeleteDeviceResponse, error) {
	out := new(delete_device.DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) CreateWateringLog(ctx context.Context, in *watering_log.CreateWateringLogRequest, opts ...grpc.CallOption) (*watering_log.CreateWateringLogResponse, error) {
	out := new(watering_log.CreateWateringLogResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/CreateWateringLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) EditDeviceName(ctx context.Context, in *device.EditDeviceNameRequest, opts ...grpc.CallOption) (*device.EditDeviceNameResponse, error) {
	out := new(device.EditDeviceNameResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/EditDeviceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) UpdateDeviceHumidity(ctx context.Context, in *update_device_humidity.UpdateDeviceHumidityRequest, opts ...grpc.CallOption) (*update_device_humidity.UpdateDeviceHumidityResponse, error) {
	out := new(update_device_humidity.UpdateDeviceHumidityResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/UpdateDeviceHumidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) SetDeviceHumidity(ctx context.Context, in *set_device_humidity.SetDeviceHumidityRequest, opts ...grpc.CallOption) (*set_device_humidity.SetDeviceHumidityResponse, error) {
	out := new(set_device_humidity.SetDeviceHumidityResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/SetDeviceHumidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) Warning(ctx context.Context, in *warning.WarningRequest, opts ...grpc.CallOption) (Farm_WarningClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Farm_serviceDesc.Streams[0], "/service.farm.Farm/Warning", opts...)
	if err != nil {
		return nil, err
	}
	x := &farmWarningClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Farm_WarningClient interface {
	Recv() (*warning.WarningResponse, error)
	grpc.ClientStream
}

type farmWarningClient struct {
	grpc.ClientStream
}

func (x *farmWarningClient) Recv() (*warning.WarningResponse, error) {
	m := new(warning.WarningResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FarmServer is the server API for Farm service.
type FarmServer interface {
	Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error)
	CreateDevice(context.Context, *device.CreateDeviceRequest) (*device.CreateDeviceResponse, error)
	GetDeviceLog(context.Context, *device_log.GetDeviceLogRequest) (*device_log.GetDeviceLogResponse, error)
	GetDeviceList(context.Context, *device_list.GetDeviceListRequest) (*device_list.GetDeviceListResponse, error)
	Logout(context.Context, *user.LogoutRequest) (*user.LogoutResponse, error)
	DeleteDevice(context.Context, *delete_device.DeleteDeviceRequest) (*delete_device.DeleteDeviceResponse, error)
	CreateWateringLog(context.Context, *watering_log.CreateWateringLogRequest) (*watering_log.CreateWateringLogResponse, error)
	EditDeviceName(context.Context, *device.EditDeviceNameRequest) (*device.EditDeviceNameResponse, error)
	UpdateDeviceHumidity(context.Context, *update_device_humidity.UpdateDeviceHumidityRequest) (*update_device_humidity.UpdateDeviceHumidityResponse, error)
	SetDeviceHumidity(context.Context, *set_device_humidity.SetDeviceHumidityRequest) (*set_device_humidity.SetDeviceHumidityResponse, error)
	Warning(*warning.WarningRequest, Farm_WarningServer) error
}

// UnimplementedFarmServer can be embedded to have forward compatible implementations.
type UnimplementedFarmServer struct {
}

func (*UnimplementedFarmServer) Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedFarmServer) CreateDevice(context.Context, *device.CreateDeviceRequest) (*device.CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (*UnimplementedFarmServer) GetDeviceLog(context.Context, *device_log.GetDeviceLogRequest) (*device_log.GetDeviceLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceLog not implemented")
}
func (*UnimplementedFarmServer) GetDeviceList(context.Context, *device_list.GetDeviceListRequest) (*device_list.GetDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceList not implemented")
}
func (*UnimplementedFarmServer) Logout(context.Context, *user.LogoutRequest) (*user.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedFarmServer) DeleteDevice(context.Context, *delete_device.DeleteDeviceRequest) (*delete_device.DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (*UnimplementedFarmServer) CreateWateringLog(context.Context, *watering_log.CreateWateringLogRequest) (*watering_log.CreateWateringLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWateringLog not implemented")
}
func (*UnimplementedFarmServer) EditDeviceName(context.Context, *device.EditDeviceNameRequest) (*device.EditDeviceNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditDeviceName not implemented")
}
func (*UnimplementedFarmServer) UpdateDeviceHumidity(context.Context, *update_device_humidity.UpdateDeviceHumidityRequest) (*update_device_humidity.UpdateDeviceHumidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceHumidity not implemented")
}
func (*UnimplementedFarmServer) SetDeviceHumidity(context.Context, *set_device_humidity.SetDeviceHumidityRequest) (*set_device_humidity.SetDeviceHumidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceHumidity not implemented")
}
func (*UnimplementedFarmServer) Warning(*warning.WarningRequest, Farm_WarningServer) error {
	return status.Errorf(codes.Unimplemented, "method Warning not implemented")
}

func RegisterFarmServer(s *grpc.Server, srv FarmServer) {
	s.RegisterService(&_Farm_serviceDesc, srv)
}

func _Farm_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).Login(ctx, req.(*user.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(device.CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).CreateDevice(ctx, req.(*device.CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_GetDeviceLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(device_log.GetDeviceLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).GetDeviceLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/GetDeviceLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).GetDeviceLog(ctx, req.(*device_log.GetDeviceLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_GetDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(device_list.GetDeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).GetDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/GetDeviceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).GetDeviceList(ctx, req.(*device_list.GetDeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).Logout(ctx, req.(*user.LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(delete_device.DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).DeleteDevice(ctx, req.(*delete_device.DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_CreateWateringLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(watering_log.CreateWateringLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).CreateWateringLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/CreateWateringLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).CreateWateringLog(ctx, req.(*watering_log.CreateWateringLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_EditDeviceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(device.EditDeviceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).EditDeviceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/EditDeviceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).EditDeviceName(ctx, req.(*device.EditDeviceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_UpdateDeviceHumidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(update_device_humidity.UpdateDeviceHumidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).UpdateDeviceHumidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/UpdateDeviceHumidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).UpdateDeviceHumidity(ctx, req.(*update_device_humidity.UpdateDeviceHumidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_SetDeviceHumidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(set_device_humidity.SetDeviceHumidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).SetDeviceHumidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/SetDeviceHumidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).SetDeviceHumidity(ctx, req.(*set_device_humidity.SetDeviceHumidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_Warning_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(warning.WarningRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FarmServer).Warning(m, &farmWarningServer{stream})
}

type Farm_WarningServer interface {
	Send(*warning.WarningResponse) error
	grpc.ServerStream
}

type farmWarningServer struct {
	grpc.ServerStream
}

func (x *farmWarningServer) Send(m *warning.WarningResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Farm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.farm.Farm",
	HandlerType: (*FarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Farm_Login_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _Farm_CreateDevice_Handler,
		},
		{
			MethodName: "GetDeviceLog",
			Handler:    _Farm_GetDeviceLog_Handler,
		},
		{
			MethodName: "GetDeviceList",
			Handler:    _Farm_GetDeviceList_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Farm_Logout_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Farm_DeleteDevice_Handler,
		},
		{
			MethodName: "CreateWateringLog",
			Handler:    _Farm_CreateWateringLog_Handler,
		},
		{
			MethodName: "EditDeviceName",
			Handler:    _Farm_EditDeviceName_Handler,
		},
		{
			MethodName: "UpdateDeviceHumidity",
			Handler:    _Farm_UpdateDeviceHumidity_Handler,
		},
		{
			MethodName: "SetDeviceHumidity",
			Handler:    _Farm_SetDeviceHumidity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Warning",
			Handler:       _Farm_Warning_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/farm.proto",
}
