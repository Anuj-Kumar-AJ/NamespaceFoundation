// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/workspace.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Workspace definition.
type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string         `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Env        []*Environment `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// Package manager.
	Dep []*Workspace_Dependency `protobuf:"bytes,3,rep,name=dep,proto3" json:"dep,omitempty"`
	// Development options.
	Replace     []*Workspace_Replace `protobuf:"bytes,4,rep,name=replace,proto3" json:"replace,omitempty"`
	PrivateRepo []string             `protobuf:"bytes,5,rep,name=private_repo,json=privateRepo,proto3" json:"private_repo,omitempty"`
	// Workspace-wide pre-compiled binaries.
	PrebuiltBinary         []*Workspace_BinaryDigest         `protobuf:"bytes,6,rep,name=prebuilt_binary,json=prebuiltBinary,proto3" json:"prebuilt_binary,omitempty"`
	PrebuiltBaseRepository string                            `protobuf:"bytes,7,opt,name=prebuilt_base_repository,json=prebuiltBaseRepository,proto3" json:"prebuilt_base_repository,omitempty"`
	Foundation             *Workspace_FoundationRequirements `protobuf:"bytes,8,opt,name=foundation,proto3" json:"foundation,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *Workspace) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Workspace) GetEnv() []*Environment {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Workspace) GetDep() []*Workspace_Dependency {
	if x != nil {
		return x.Dep
	}
	return nil
}

func (x *Workspace) GetReplace() []*Workspace_Replace {
	if x != nil {
		return x.Replace
	}
	return nil
}

func (x *Workspace) GetPrivateRepo() []string {
	if x != nil {
		return x.PrivateRepo
	}
	return nil
}

func (x *Workspace) GetPrebuiltBinary() []*Workspace_BinaryDigest {
	if x != nil {
		return x.PrebuiltBinary
	}
	return nil
}

func (x *Workspace) GetPrebuiltBaseRepository() string {
	if x != nil {
		return x.PrebuiltBaseRepository
	}
	return ""
}

func (x *Workspace) GetFoundation() *Workspace_FoundationRequirements {
	if x != nil {
		return x.Foundation
	}
	return nil
}

// Configure a developer workstation.
type DevHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configure         []*DevHost_ConfigureEnvironment `protobuf:"bytes,1,rep,name=configure,proto3" json:"configure,omitempty"`
	ConfigurePlatform []*DevHost_ConfigurePlatform    `protobuf:"bytes,2,rep,name=configure_platform,json=configurePlatform,proto3" json:"configure_platform,omitempty"`
	ConfigureTools    []*anypb.Any                    `protobuf:"bytes,3,rep,name=configure_tools,json=configureTools,proto3" json:"configure_tools,omitempty"`
}

func (x *DevHost) Reset() {
	*x = DevHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevHost) ProtoMessage() {}

func (x *DevHost) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevHost.ProtoReflect.Descriptor instead.
func (*DevHost) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *DevHost) GetConfigure() []*DevHost_ConfigureEnvironment {
	if x != nil {
		return x.Configure
	}
	return nil
}

func (x *DevHost) GetConfigurePlatform() []*DevHost_ConfigurePlatform {
	if x != nil {
		return x.ConfigurePlatform
	}
	return nil
}

func (x *DevHost) GetConfigureTools() []*anypb.Any {
	if x != nil {
		return x.ConfigureTools
	}
	return nil
}

type Workspace_Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Workspace_Dependency) Reset() {
	*x = Workspace_Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace_Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace_Dependency) ProtoMessage() {}

func (x *Workspace_Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace_Dependency.ProtoReflect.Descriptor instead.
func (*Workspace_Dependency) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Workspace_Dependency) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Workspace_Dependency) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Workspace_Replace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	Path       string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Workspace_Replace) Reset() {
	*x = Workspace_Replace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace_Replace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace_Replace) ProtoMessage() {}

func (x *Workspace_Replace) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace_Replace.ProtoReflect.Descriptor instead.
func (*Workspace_Replace) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Workspace_Replace) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Workspace_Replace) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Workspace_BinaryDigest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Digest      string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Repository  string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *Workspace_BinaryDigest) Reset() {
	*x = Workspace_BinaryDigest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace_BinaryDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace_BinaryDigest) ProtoMessage() {}

func (x *Workspace_BinaryDigest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace_BinaryDigest.ProtoReflect.Descriptor instead.
func (*Workspace_BinaryDigest) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Workspace_BinaryDigest) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Workspace_BinaryDigest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Workspace_BinaryDigest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type Workspace_FoundationRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumApi int32 `protobuf:"varint,1,opt,name=minimum_api,json=minimumApi,proto3" json:"minimum_api,omitempty"`
}

func (x *Workspace_FoundationRequirements) Reset() {
	*x = Workspace_FoundationRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace_FoundationRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace_FoundationRequirements) ProtoMessage() {}

func (x *Workspace_FoundationRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace_FoundationRequirements.ProtoReflect.Descriptor instead.
func (*Workspace_FoundationRequirements) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Workspace_FoundationRequirements) GetMinimumApi() int32 {
	if x != nil {
		return x.MinimumApi
	}
	return 0
}

type DevHost_ConfigureEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Runtime       string              `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Purpose       Environment_Purpose `protobuf:"varint,3,opt,name=purpose,proto3,enum=foundation.schema.Environment_Purpose" json:"purpose,omitempty"`
	Configuration []*anypb.Any        `protobuf:"bytes,4,rep,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *DevHost_ConfigureEnvironment) Reset() {
	*x = DevHost_ConfigureEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevHost_ConfigureEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevHost_ConfigureEnvironment) ProtoMessage() {}

func (x *DevHost_ConfigureEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevHost_ConfigureEnvironment.ProtoReflect.Descriptor instead.
func (*DevHost_ConfigureEnvironment) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DevHost_ConfigureEnvironment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DevHost_ConfigureEnvironment) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *DevHost_ConfigureEnvironment) GetPurpose() Environment_Purpose {
	if x != nil {
		return x.Purpose
	}
	return Environment_PURPOSE_UNKNOWN
}

func (x *DevHost_ConfigureEnvironment) GetConfiguration() []*anypb.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type DevHost_ConfigurePlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Architecture  string       `protobuf:"bytes,1,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Os            string       `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Variant       string       `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	Configuration []*anypb.Any `protobuf:"bytes,4,rep,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *DevHost_ConfigurePlatform) Reset() {
	*x = DevHost_ConfigurePlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevHost_ConfigurePlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevHost_ConfigurePlatform) ProtoMessage() {}

func (x *DevHost_ConfigurePlatform) ProtoReflect() protoreflect.Message {
	mi := &file_schema_workspace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevHost_ConfigurePlatform.ProtoReflect.Descriptor instead.
func (*DevHost_ConfigurePlatform) Descriptor() ([]byte, []int) {
	return file_schema_workspace_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DevHost_ConfigurePlatform) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *DevHost_ConfigurePlatform) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *DevHost_ConfigurePlatform) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *DevHost_ConfigurePlatform) GetConfiguration() []*anypb.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

var File_schema_workspace_proto protoreflect.FileDescriptor

var file_schema_workspace_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x06, 0x0a, 0x09,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x39, 0x0a, 0x03,
	0x64, 0x65, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x03, 0x64, 0x65, 0x70, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x52, 0x0a, 0x0f, 0x70, 0x72,
	0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0e,
	0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x38,
	0x0a, 0x18, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x70, 0x72, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x0a, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x0a, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x47, 0x0a,
	0x0a, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x69, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x1a, 0x39, 0x0a, 0x16, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x41, 0x70, 0x69, 0x22, 0xd9, 0x04, 0x0a,
	0x07, 0x44, 0x65, 0x76, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x44, 0x65, 0x76, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x76, 0x48, 0x6f, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x54, 0x6f,
	0x6f, 0x6c, 0x73, 0x1a, 0xc2, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x70,
	0x6f, 0x73, 0x65, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x9d, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x25, 0x5a, 0x23, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_workspace_proto_rawDescOnce sync.Once
	file_schema_workspace_proto_rawDescData = file_schema_workspace_proto_rawDesc
)

func file_schema_workspace_proto_rawDescGZIP() []byte {
	file_schema_workspace_proto_rawDescOnce.Do(func() {
		file_schema_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_workspace_proto_rawDescData)
	})
	return file_schema_workspace_proto_rawDescData
}

var file_schema_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_schema_workspace_proto_goTypes = []interface{}{
	(*Workspace)(nil),                        // 0: foundation.schema.Workspace
	(*DevHost)(nil),                          // 1: foundation.schema.DevHost
	(*Workspace_Dependency)(nil),             // 2: foundation.schema.Workspace.Dependency
	(*Workspace_Replace)(nil),                // 3: foundation.schema.Workspace.Replace
	(*Workspace_BinaryDigest)(nil),           // 4: foundation.schema.Workspace.BinaryDigest
	(*Workspace_FoundationRequirements)(nil), // 5: foundation.schema.Workspace.FoundationRequirements
	(*DevHost_ConfigureEnvironment)(nil),     // 6: foundation.schema.DevHost.ConfigureEnvironment
	(*DevHost_ConfigurePlatform)(nil),        // 7: foundation.schema.DevHost.ConfigurePlatform
	(*Environment)(nil),                      // 8: foundation.schema.Environment
	(*anypb.Any)(nil),                        // 9: google.protobuf.Any
	(Environment_Purpose)(0),                 // 10: foundation.schema.Environment.Purpose
}
var file_schema_workspace_proto_depIdxs = []int32{
	8,  // 0: foundation.schema.Workspace.env:type_name -> foundation.schema.Environment
	2,  // 1: foundation.schema.Workspace.dep:type_name -> foundation.schema.Workspace.Dependency
	3,  // 2: foundation.schema.Workspace.replace:type_name -> foundation.schema.Workspace.Replace
	4,  // 3: foundation.schema.Workspace.prebuilt_binary:type_name -> foundation.schema.Workspace.BinaryDigest
	5,  // 4: foundation.schema.Workspace.foundation:type_name -> foundation.schema.Workspace.FoundationRequirements
	6,  // 5: foundation.schema.DevHost.configure:type_name -> foundation.schema.DevHost.ConfigureEnvironment
	7,  // 6: foundation.schema.DevHost.configure_platform:type_name -> foundation.schema.DevHost.ConfigurePlatform
	9,  // 7: foundation.schema.DevHost.configure_tools:type_name -> google.protobuf.Any
	10, // 8: foundation.schema.DevHost.ConfigureEnvironment.purpose:type_name -> foundation.schema.Environment.Purpose
	9,  // 9: foundation.schema.DevHost.ConfigureEnvironment.configuration:type_name -> google.protobuf.Any
	9,  // 10: foundation.schema.DevHost.ConfigurePlatform.configuration:type_name -> google.protobuf.Any
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_schema_workspace_proto_init() }
func file_schema_workspace_proto_init() {
	if File_schema_workspace_proto != nil {
		return
	}
	file_schema_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
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
		file_schema_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevHost); i {
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
		file_schema_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace_Dependency); i {
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
		file_schema_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace_Replace); i {
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
		file_schema_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace_BinaryDigest); i {
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
		file_schema_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace_FoundationRequirements); i {
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
		file_schema_workspace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevHost_ConfigureEnvironment); i {
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
		file_schema_workspace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevHost_ConfigurePlatform); i {
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
			RawDescriptor: file_schema_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_workspace_proto_goTypes,
		DependencyIndexes: file_schema_workspace_proto_depIdxs,
		MessageInfos:      file_schema_workspace_proto_msgTypes,
	}.Build()
	File_schema_workspace_proto = out.File
	file_schema_workspace_proto_rawDesc = nil
	file_schema_workspace_proto_goTypes = nil
	file_schema_workspace_proto_depIdxs = nil
}
