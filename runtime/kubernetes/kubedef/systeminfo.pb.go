// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: runtime/kubernetes/kubedef/systeminfo.proto

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

type SystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodePlatform         []string            `protobuf:"bytes,1,rep,name=node_platform,json=nodePlatform,proto3" json:"node_platform,omitempty"`
	DetectedDistribution string              `protobuf:"bytes,2,opt,name=detected_distribution,json=detectedDistribution,proto3" json:"detected_distribution,omitempty"` // k3d, eks, etc.
	EksClusterName       string              `protobuf:"bytes,3,opt,name=eks_cluster_name,json=eksClusterName,proto3" json:"eks_cluster_name,omitempty"`                 // Only set if distribution is eks.
	Regions              []string            `protobuf:"bytes,4,rep,name=regions,proto3" json:"regions,omitempty"`                                                       // See region_distribution for details.
	Zones                []string            `protobuf:"bytes,5,rep,name=zones,proto3" json:"zones,omitempty"`                                                           // Set zone_distribution for details.
	RegionDistribution   []*NodeDistribution `protobuf:"bytes,6,rep,name=region_distribution,json=regionDistribution,proto3" json:"region_distribution,omitempty"`
	ZoneDistribution     []*NodeDistribution `protobuf:"bytes,7,rep,name=zone_distribution,json=zoneDistribution,proto3" json:"zone_distribution,omitempty"`
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescGZIP(), []int{0}
}

func (x *SystemInfo) GetNodePlatform() []string {
	if x != nil {
		return x.NodePlatform
	}
	return nil
}

func (x *SystemInfo) GetDetectedDistribution() string {
	if x != nil {
		return x.DetectedDistribution
	}
	return ""
}

func (x *SystemInfo) GetEksClusterName() string {
	if x != nil {
		return x.EksClusterName
	}
	return ""
}

func (x *SystemInfo) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *SystemInfo) GetZones() []string {
	if x != nil {
		return x.Zones
	}
	return nil
}

func (x *SystemInfo) GetRegionDistribution() []*NodeDistribution {
	if x != nil {
		return x.RegionDistribution
	}
	return nil
}

func (x *SystemInfo) GetZoneDistribution() []*NodeDistribution {
	if x != nil {
		return x.ZoneDistribution
	}
	return nil
}

type NodeDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Count    int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *NodeDistribution) Reset() {
	*x = NodeDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDistribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDistribution) ProtoMessage() {}

func (x *NodeDistribution) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDistribution.ProtoReflect.Descriptor instead.
func (*NodeDistribution) Descriptor() ([]byte, []int) {
	return file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescGZIP(), []int{1}
}

func (x *NodeDistribution) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *NodeDistribution) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_runtime_kubernetes_kubedef_systeminfo_proto protoreflect.FileDescriptor

var file_runtime_kubernetes_kubedef_systeminfo_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x64, 0x65, 0x66, 0x22, 0x90, 0x03, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x33, 0x0a, 0x15, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x10, 0x65, 0x6b, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6b, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x64, 0x0a, 0x11, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x64, 0x65, 0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x7a, 0x6f, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x39, 0x5a,
	0x37, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescOnce sync.Once
	file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescData = file_runtime_kubernetes_kubedef_systeminfo_proto_rawDesc
)

func file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescGZIP() []byte {
	file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescOnce.Do(func() {
		file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescData)
	})
	return file_runtime_kubernetes_kubedef_systeminfo_proto_rawDescData
}

var file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_runtime_kubernetes_kubedef_systeminfo_proto_goTypes = []interface{}{
	(*SystemInfo)(nil),       // 0: foundation.runtime.kubernetes.kubedef.SystemInfo
	(*NodeDistribution)(nil), // 1: foundation.runtime.kubernetes.kubedef.NodeDistribution
}
var file_runtime_kubernetes_kubedef_systeminfo_proto_depIdxs = []int32{
	1, // 0: foundation.runtime.kubernetes.kubedef.SystemInfo.region_distribution:type_name -> foundation.runtime.kubernetes.kubedef.NodeDistribution
	1, // 1: foundation.runtime.kubernetes.kubedef.SystemInfo.zone_distribution:type_name -> foundation.runtime.kubernetes.kubedef.NodeDistribution
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_runtime_kubernetes_kubedef_systeminfo_proto_init() }
func file_runtime_kubernetes_kubedef_systeminfo_proto_init() {
	if File_runtime_kubernetes_kubedef_systeminfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemInfo); i {
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
		file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDistribution); i {
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
			RawDescriptor: file_runtime_kubernetes_kubedef_systeminfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runtime_kubernetes_kubedef_systeminfo_proto_goTypes,
		DependencyIndexes: file_runtime_kubernetes_kubedef_systeminfo_proto_depIdxs,
		MessageInfos:      file_runtime_kubernetes_kubedef_systeminfo_proto_msgTypes,
	}.Build()
	File_runtime_kubernetes_kubedef_systeminfo_proto = out.File
	file_runtime_kubernetes_kubedef_systeminfo_proto_rawDesc = nil
	file_runtime_kubernetes_kubedef_systeminfo_proto_goTypes = nil
	file_runtime_kubernetes_kubedef_systeminfo_proto_depIdxs = nil
}
