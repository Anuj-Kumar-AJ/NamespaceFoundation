// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: framework/kubernetes/kubedef/storage.proto

package kubedef

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

type KubernetesEnvDiagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemInfo          *SystemInfo `protobuf:"bytes,1,opt,name=system_info,json=systemInfo,proto3" json:"system_info,omitempty"`
	SerializedEventList string      `protobuf:"bytes,2,opt,name=serialized_event_list,json=serializedEventList,proto3" json:"serialized_event_list,omitempty"`
}

func (x *KubernetesEnvDiagnostics) Reset() {
	*x = KubernetesEnvDiagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_framework_kubernetes_kubedef_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesEnvDiagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesEnvDiagnostics) ProtoMessage() {}

func (x *KubernetesEnvDiagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_framework_kubernetes_kubedef_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesEnvDiagnostics.ProtoReflect.Descriptor instead.
func (*KubernetesEnvDiagnostics) Descriptor() ([]byte, []int) {
	return file_framework_kubernetes_kubedef_storage_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesEnvDiagnostics) GetSystemInfo() *SystemInfo {
	if x != nil {
		return x.SystemInfo
	}
	return nil
}

func (x *KubernetesEnvDiagnostics) GetSerializedEventList() string {
	if x != nil {
		return x.SerializedEventList
	}
	return ""
}

var File_framework_kubernetes_kubedef_storage_proto protoreflect.FileDescriptor

var file_framework_kubernetes_kubedef_storage_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x64, 0x65, 0x66, 0x1a, 0x2d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65,
	0x66, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x18, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x76, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x52, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x3b, 0x5a, 0x39, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x64, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_framework_kubernetes_kubedef_storage_proto_rawDescOnce sync.Once
	file_framework_kubernetes_kubedef_storage_proto_rawDescData = file_framework_kubernetes_kubedef_storage_proto_rawDesc
)

func file_framework_kubernetes_kubedef_storage_proto_rawDescGZIP() []byte {
	file_framework_kubernetes_kubedef_storage_proto_rawDescOnce.Do(func() {
		file_framework_kubernetes_kubedef_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_framework_kubernetes_kubedef_storage_proto_rawDescData)
	})
	return file_framework_kubernetes_kubedef_storage_proto_rawDescData
}

var file_framework_kubernetes_kubedef_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_framework_kubernetes_kubedef_storage_proto_goTypes = []interface{}{
	(*KubernetesEnvDiagnostics)(nil), // 0: foundation.runtime.kubernetes.kubedef.KubernetesEnvDiagnostics
	(*SystemInfo)(nil),               // 1: foundation.runtime.kubernetes.kubedef.SystemInfo
}
var file_framework_kubernetes_kubedef_storage_proto_depIdxs = []int32{
	1, // 0: foundation.runtime.kubernetes.kubedef.KubernetesEnvDiagnostics.system_info:type_name -> foundation.runtime.kubernetes.kubedef.SystemInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_framework_kubernetes_kubedef_storage_proto_init() }
func file_framework_kubernetes_kubedef_storage_proto_init() {
	if File_framework_kubernetes_kubedef_storage_proto != nil {
		return
	}
	file_framework_kubernetes_kubedef_systeminfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_framework_kubernetes_kubedef_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesEnvDiagnostics); i {
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
			RawDescriptor: file_framework_kubernetes_kubedef_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_framework_kubernetes_kubedef_storage_proto_goTypes,
		DependencyIndexes: file_framework_kubernetes_kubedef_storage_proto_depIdxs,
		MessageInfos:      file_framework_kubernetes_kubedef_storage_proto_msgTypes,
	}.Build()
	File_framework_kubernetes_kubedef_storage_proto = out.File
	file_framework_kubernetes_kubedef_storage_proto_rawDesc = nil
	file_framework_kubernetes_kubedef_storage_proto_goTypes = nil
	file_framework_kubernetes_kubedef_storage_proto_depIdxs = nil
}