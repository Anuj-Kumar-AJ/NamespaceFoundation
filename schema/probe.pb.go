// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/probe.proto

package schema

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

type Probe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string      `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Http *Probe_Http `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	Exec *Probe_Exec `protobuf:"bytes,3,opt,name=exec,proto3" json:"exec,omitempty"`
}

func (x *Probe) Reset() {
	*x = Probe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_probe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Probe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Probe) ProtoMessage() {}

func (x *Probe) ProtoReflect() protoreflect.Message {
	mi := &file_schema_probe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Probe.ProtoReflect.Descriptor instead.
func (*Probe) Descriptor() ([]byte, []int) {
	return file_schema_probe_proto_rawDescGZIP(), []int{0}
}

func (x *Probe) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Probe) GetHttp() *Probe_Http {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Probe) GetExec() *Probe_Exec {
	if x != nil {
		return x.Exec
	}
	return nil
}

type Probe_Exec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *Probe_Exec) Reset() {
	*x = Probe_Exec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_probe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Probe_Exec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Probe_Exec) ProtoMessage() {}

func (x *Probe_Exec) ProtoReflect() protoreflect.Message {
	mi := &file_schema_probe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Probe_Exec.ProtoReflect.Descriptor instead.
func (*Probe_Exec) Descriptor() ([]byte, []int) {
	return file_schema_probe_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Probe_Exec) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type Probe_Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path          string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ContainerPort int32  `protobuf:"varint,2,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
}

func (x *Probe_Http) Reset() {
	*x = Probe_Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_probe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Probe_Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Probe_Http) ProtoMessage() {}

func (x *Probe_Http) ProtoReflect() protoreflect.Message {
	mi := &file_schema_probe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Probe_Http.ProtoReflect.Descriptor instead.
func (*Probe_Http) Descriptor() ([]byte, []int) {
	return file_schema_probe_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Probe_Http) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Probe_Http) GetContainerPort() int32 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

var File_schema_probe_proto protoreflect.FileDescriptor

var file_schema_probe_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xe6, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x31, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x1a, 0x20, 0x0a, 0x04, 0x45,
	0x78, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x41, 0x0a,
	0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_probe_proto_rawDescOnce sync.Once
	file_schema_probe_proto_rawDescData = file_schema_probe_proto_rawDesc
)

func file_schema_probe_proto_rawDescGZIP() []byte {
	file_schema_probe_proto_rawDescOnce.Do(func() {
		file_schema_probe_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_probe_proto_rawDescData)
	})
	return file_schema_probe_proto_rawDescData
}

var file_schema_probe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_probe_proto_goTypes = []interface{}{
	(*Probe)(nil),      // 0: foundation.schema.Probe
	(*Probe_Exec)(nil), // 1: foundation.schema.Probe.Exec
	(*Probe_Http)(nil), // 2: foundation.schema.Probe.Http
}
var file_schema_probe_proto_depIdxs = []int32{
	2, // 0: foundation.schema.Probe.http:type_name -> foundation.schema.Probe.Http
	1, // 1: foundation.schema.Probe.exec:type_name -> foundation.schema.Probe.Exec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_schema_probe_proto_init() }
func file_schema_probe_proto_init() {
	if File_schema_probe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_probe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Probe); i {
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
		file_schema_probe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Probe_Exec); i {
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
		file_schema_probe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Probe_Http); i {
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
			RawDescriptor: file_schema_probe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_probe_proto_goTypes,
		DependencyIndexes: file_schema_probe_proto_depIdxs,
		MessageInfos:      file_schema_probe_proto_msgTypes,
	}.Build()
	File_schema_probe_proto = out.File
	file_schema_probe_proto_rawDesc = nil
	file_schema_probe_proto_goTypes = nil
	file_schema_probe_proto_depIdxs = nil
}