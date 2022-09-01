// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: workspace/source/op.proto

package source

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	schema "namespacelabs.dev/foundation/schema"
	protos "namespacelabs.dev/foundation/workspace/source/protos"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpProtoGen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string                           `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Protos      *protos.FileDescriptorSetAndDeps `protobuf:"bytes,3,opt,name=protos,proto3" json:"protos,omitempty"`
	Framework   schema.Framework                 `protobuf:"varint,5,opt,name=framework,proto3,enum=foundation.schema.Framework" json:"framework,omitempty"`
}

func (x *OpProtoGen) Reset() {
	*x = OpProtoGen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_source_op_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpProtoGen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpProtoGen) ProtoMessage() {}

func (x *OpProtoGen) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_source_op_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpProtoGen.ProtoReflect.Descriptor instead.
func (*OpProtoGen) Descriptor() ([]byte, []int) {
	return file_workspace_source_op_proto_rawDescGZIP(), []int{0}
}

func (x *OpProtoGen) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *OpProtoGen) GetProtos() *protos.FileDescriptorSetAndDeps {
	if x != nil {
		return x.Protos
	}
	return nil
}

func (x *OpProtoGen) GetFramework() schema.Framework {
	if x != nil {
		return x.Framework
	}
	return schema.Framework(0)
}

var File_workspace_source_op_proto protoreflect.FileDescriptor

var file_workspace_source_op_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x29, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x65, 0x73, 0x63, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x70, 0x73, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x12, 0x3a, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x2f, 0x5a, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_workspace_source_op_proto_rawDescOnce sync.Once
	file_workspace_source_op_proto_rawDescData = file_workspace_source_op_proto_rawDesc
)

func file_workspace_source_op_proto_rawDescGZIP() []byte {
	file_workspace_source_op_proto_rawDescOnce.Do(func() {
		file_workspace_source_op_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_source_op_proto_rawDescData)
	})
	return file_workspace_source_op_proto_rawDescData
}

var file_workspace_source_op_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_workspace_source_op_proto_goTypes = []interface{}{
	(*OpProtoGen)(nil),                      // 0: foundation.workspace.source.OpProtoGen
	(*protos.FileDescriptorSetAndDeps)(nil), // 1: foundation.workspace.source.protos.FileDescriptorSetAndDeps
	(schema.Framework)(0),                   // 2: foundation.schema.Framework
}
var file_workspace_source_op_proto_depIdxs = []int32{
	1, // 0: foundation.workspace.source.OpProtoGen.protos:type_name -> foundation.workspace.source.protos.FileDescriptorSetAndDeps
	2, // 1: foundation.workspace.source.OpProtoGen.framework:type_name -> foundation.schema.Framework
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_workspace_source_op_proto_init() }
func file_workspace_source_op_proto_init() {
	if File_workspace_source_op_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workspace_source_op_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpProtoGen); i {
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
			RawDescriptor: file_workspace_source_op_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workspace_source_op_proto_goTypes,
		DependencyIndexes: file_workspace_source_op_proto_depIdxs,
		MessageInfos:      file_workspace_source_op_proto_msgTypes,
	}.Build()
	File_workspace_source_op_proto = out.File
	file_workspace_source_op_proto_rawDesc = nil
	file_workspace_source_op_proto_goTypes = nil
	file_workspace_source_op_proto_depIdxs = nil
}
