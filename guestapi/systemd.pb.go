// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: systemd.proto

package guestapi

import (
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

var File_systemd_proto protoreflect.FileDescriptor

var file_systemd_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x64, 0x65, 0x76, 0x2e, 0x6c, 0x61, 0x62, 0x34, 0x37, 0x2e, 0x69, 0x73, 0x6c, 0x65, 0x2e,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x42, 0x19, 0x5a, 0x17, 0x6c, 0x61, 0x62, 0x34,
	0x37, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x73, 0x6c, 0x65, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_systemd_proto_goTypes = []interface{}{}
var file_systemd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_systemd_proto_init() }
func file_systemd_proto_init() {
	if File_systemd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_systemd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_systemd_proto_goTypes,
		DependencyIndexes: file_systemd_proto_depIdxs,
	}.Build()
	File_systemd_proto = out.File
	file_systemd_proto_rawDesc = nil
	file_systemd_proto_goTypes = nil
	file_systemd_proto_depIdxs = nil
}
