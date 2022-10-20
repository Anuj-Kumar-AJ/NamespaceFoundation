// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: framework/kubernetes/kubetool/hostenv.proto

package kubetool

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

type KubernetesEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *KubernetesEnv) Reset() {
	*x = KubernetesEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesEnv) ProtoMessage() {}

func (x *KubernetesEnv) ProtoReflect() protoreflect.Message {
	mi := &file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesEnv.ProtoReflect.Descriptor instead.
func (*KubernetesEnv) Descriptor() ([]byte, []int) {
	return file_framework_kubernetes_kubetool_hostenv_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesEnv) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type KubernetesToolContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace           string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CanSetNamespace     bool   `protobuf:"varint,2,opt,name=can_set_namespace,json=canSetNamespace,proto3" json:"can_set_namespace,omitempty"`
	HasApplyRoleBinding bool   `protobuf:"varint,3,opt,name=has_apply_role_binding,json=hasApplyRoleBinding,proto3" json:"has_apply_role_binding,omitempty"`
}

func (x *KubernetesToolContext) Reset() {
	*x = KubernetesToolContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesToolContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesToolContext) ProtoMessage() {}

func (x *KubernetesToolContext) ProtoReflect() protoreflect.Message {
	mi := &file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesToolContext.ProtoReflect.Descriptor instead.
func (*KubernetesToolContext) Descriptor() ([]byte, []int) {
	return file_framework_kubernetes_kubetool_hostenv_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesToolContext) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesToolContext) GetCanSetNamespace() bool {
	if x != nil {
		return x.CanSetNamespace
	}
	return false
}

func (x *KubernetesToolContext) GetHasApplyRoleBinding() bool {
	if x != nil {
		return x.HasApplyRoleBinding
	}
	return false
}

var File_framework_kubernetes_kubetool_hostenv_proto protoreflect.FileDescriptor

var file_framework_kubernetes_kubetool_hostenv_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x74, 0x6f, 0x6f, 0x6c, 0x2f,
	0x68, 0x6f, 0x73, 0x74, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x74, 0x6f, 0x6f, 0x6c, 0x22, 0x2d, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x76, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x15, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x54, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x63, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x68, 0x61, 0x73, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x68, 0x61, 0x73, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x3c, 0x5a,
	0x3a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x74, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_framework_kubernetes_kubetool_hostenv_proto_rawDescOnce sync.Once
	file_framework_kubernetes_kubetool_hostenv_proto_rawDescData = file_framework_kubernetes_kubetool_hostenv_proto_rawDesc
)

func file_framework_kubernetes_kubetool_hostenv_proto_rawDescGZIP() []byte {
	file_framework_kubernetes_kubetool_hostenv_proto_rawDescOnce.Do(func() {
		file_framework_kubernetes_kubetool_hostenv_proto_rawDescData = protoimpl.X.CompressGZIP(file_framework_kubernetes_kubetool_hostenv_proto_rawDescData)
	})
	return file_framework_kubernetes_kubetool_hostenv_proto_rawDescData
}

var file_framework_kubernetes_kubetool_hostenv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_framework_kubernetes_kubetool_hostenv_proto_goTypes = []interface{}{
	(*KubernetesEnv)(nil),         // 0: foundation.runtime.kubernetes.kubetool.KubernetesEnv
	(*KubernetesToolContext)(nil), // 1: foundation.runtime.kubernetes.kubetool.KubernetesToolContext
}
var file_framework_kubernetes_kubetool_hostenv_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_framework_kubernetes_kubetool_hostenv_proto_init() }
func file_framework_kubernetes_kubetool_hostenv_proto_init() {
	if File_framework_kubernetes_kubetool_hostenv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesEnv); i {
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
		file_framework_kubernetes_kubetool_hostenv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesToolContext); i {
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
			RawDescriptor: file_framework_kubernetes_kubetool_hostenv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_framework_kubernetes_kubetool_hostenv_proto_goTypes,
		DependencyIndexes: file_framework_kubernetes_kubetool_hostenv_proto_depIdxs,
		MessageInfos:      file_framework_kubernetes_kubetool_hostenv_proto_msgTypes,
	}.Build()
	File_framework_kubernetes_kubetool_hostenv_proto = out.File
	file_framework_kubernetes_kubetool_hostenv_proto_rawDesc = nil
	file_framework_kubernetes_kubetool_hostenv_proto_goTypes = nil
	file_framework_kubernetes_kubetool_hostenv_proto_depIdxs = nil
}
