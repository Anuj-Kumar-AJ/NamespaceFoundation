// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: framework/kubernetes/kubedef/podreference.proto

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

type ContainerPodReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName   string `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Container string `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *ContainerPodReference) Reset() {
	*x = ContainerPodReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_framework_kubernetes_kubedef_podreference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerPodReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerPodReference) ProtoMessage() {}

func (x *ContainerPodReference) ProtoReflect() protoreflect.Message {
	mi := &file_framework_kubernetes_kubedef_podreference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerPodReference.ProtoReflect.Descriptor instead.
func (*ContainerPodReference) Descriptor() ([]byte, []int) {
	return file_framework_kubernetes_kubedef_podreference_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerPodReference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ContainerPodReference) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ContainerPodReference) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

var File_framework_kubernetes_kubedef_podreference_proto protoreflect.FileDescriptor

var file_framework_kubernetes_kubedef_podreference_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x2f, 0x70,
	0x6f, 0x64, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x22, 0x6e, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x3b, 0x5a, 0x39, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x64, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_framework_kubernetes_kubedef_podreference_proto_rawDescOnce sync.Once
	file_framework_kubernetes_kubedef_podreference_proto_rawDescData = file_framework_kubernetes_kubedef_podreference_proto_rawDesc
)

func file_framework_kubernetes_kubedef_podreference_proto_rawDescGZIP() []byte {
	file_framework_kubernetes_kubedef_podreference_proto_rawDescOnce.Do(func() {
		file_framework_kubernetes_kubedef_podreference_proto_rawDescData = protoimpl.X.CompressGZIP(file_framework_kubernetes_kubedef_podreference_proto_rawDescData)
	})
	return file_framework_kubernetes_kubedef_podreference_proto_rawDescData
}

var file_framework_kubernetes_kubedef_podreference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_framework_kubernetes_kubedef_podreference_proto_goTypes = []interface{}{
	(*ContainerPodReference)(nil), // 0: foundation.runtime.kubernetes.kubedef.ContainerPodReference
}
var file_framework_kubernetes_kubedef_podreference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_framework_kubernetes_kubedef_podreference_proto_init() }
func file_framework_kubernetes_kubedef_podreference_proto_init() {
	if File_framework_kubernetes_kubedef_podreference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_framework_kubernetes_kubedef_podreference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerPodReference); i {
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
			RawDescriptor: file_framework_kubernetes_kubedef_podreference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_framework_kubernetes_kubedef_podreference_proto_goTypes,
		DependencyIndexes: file_framework_kubernetes_kubedef_podreference_proto_depIdxs,
		MessageInfos:      file_framework_kubernetes_kubedef_podreference_proto_msgTypes,
	}.Build()
	File_framework_kubernetes_kubedef_podreference_proto = out.File
	file_framework_kubernetes_kubedef_podreference_proto_rawDesc = nil
	file_framework_kubernetes_kubedef_podreference_proto_goTypes = nil
	file_framework_kubernetes_kubedef_podreference_proto_depIdxs = nil
}
