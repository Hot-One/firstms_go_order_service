// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: courier_service.proto

package courier_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

var File_courier_service_proto protoreflect.FileDescriptor

var file_courier_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x18, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x05,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a,
	0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_courier_service_proto_goTypes = []interface{}{
	(*CourierCreate)(nil),          // 0: courier_service.CourierCreate
	(*CourierPrimaryKey)(nil),      // 1: courier_service.CourierPrimaryKey
	(*CourierGetListRequest)(nil),  // 2: courier_service.CourierGetListRequest
	(*CourierUpdate)(nil),          // 3: courier_service.CourierUpdate
	(*CourierUpdatePatch)(nil),     // 4: courier_service.CourierUpdatePatch
	(*Courier)(nil),                // 5: courier_service.Courier
	(*CourierGetListResponse)(nil), // 6: courier_service.CourierGetListResponse
	(*empty.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_courier_service_proto_depIdxs = []int32{
	0, // 0: courier_service.CourierService.Create:input_type -> courier_service.CourierCreate
	1, // 1: courier_service.CourierService.GetById:input_type -> courier_service.CourierPrimaryKey
	2, // 2: courier_service.CourierService.GetList:input_type -> courier_service.CourierGetListRequest
	3, // 3: courier_service.CourierService.Update:input_type -> courier_service.CourierUpdate
	4, // 4: courier_service.CourierService.Patch:input_type -> courier_service.CourierUpdatePatch
	1, // 5: courier_service.CourierService.Delete:input_type -> courier_service.CourierPrimaryKey
	5, // 6: courier_service.CourierService.Create:output_type -> courier_service.Courier
	5, // 7: courier_service.CourierService.GetById:output_type -> courier_service.Courier
	6, // 8: courier_service.CourierService.GetList:output_type -> courier_service.CourierGetListResponse
	5, // 9: courier_service.CourierService.Update:output_type -> courier_service.Courier
	5, // 10: courier_service.CourierService.Patch:output_type -> courier_service.Courier
	7, // 11: courier_service.CourierService.Delete:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_courier_service_proto_init() }
func file_courier_service_proto_init() {
	if File_courier_service_proto != nil {
		return
	}
	file_courier_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_courier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courier_service_proto_goTypes,
		DependencyIndexes: file_courier_service_proto_depIdxs,
	}.Build()
	File_courier_service_proto = out.File
	file_courier_service_proto_rawDesc = nil
	file_courier_service_proto_goTypes = nil
	file_courier_service_proto_depIdxs = nil
}
