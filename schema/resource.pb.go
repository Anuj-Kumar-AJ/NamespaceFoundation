// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/resource.proto

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

// A resource represents an instance of an entity class, i.e. a _thing_ that has
// been instantiated to be used by a server, or by other resources.
type ResourceInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string      `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"` // Computed, where this resource lives.
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                  // The name of the resource, scoped to the package where it lives.
	Class       *PackageRef `protobuf:"bytes,3,opt,name=class,proto3" json:"class,omitempty"`                                // The class of resource being instantiated.
	Provider    string      `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`                          // Package name of the provider that knows how to instantiate this resource.
	// A resource is instantiated from an intent, which can be either specified inline, or be computed.
	Intent     *anypb.Any  `protobuf:"bytes,5,opt,name=intent,proto3" json:"intent,omitempty"`
	IntentFrom *Invocation `protobuf:"bytes,6,opt,name=intent_from,json=intentFrom,proto3" json:"intent_from,omitempty"`
	// Set of resources passed to provider. Must match the requirements set by the provider.
	InputResource []*ResourceInstance_InputResource `protobuf:"bytes,7,rep,name=input_resource,json=inputResource,proto3" json:"input_resource,omitempty"`
}

func (x *ResourceInstance) Reset() {
	*x = ResourceInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInstance) ProtoMessage() {}

func (x *ResourceInstance) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInstance.ProtoReflect.Descriptor instead.
func (*ResourceInstance) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceInstance) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *ResourceInstance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceInstance) GetClass() *PackageRef {
	if x != nil {
		return x.Class
	}
	return nil
}

func (x *ResourceInstance) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ResourceInstance) GetIntent() *anypb.Any {
	if x != nil {
		return x.Intent
	}
	return nil
}

func (x *ResourceInstance) GetIntentFrom() *Invocation {
	if x != nil {
		return x.IntentFrom
	}
	return nil
}

func (x *ResourceInstance) GetInputResource() []*ResourceInstance_InputResource {
	if x != nil {
		return x.InputResource
	}
	return nil
}

type ResourceClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName     string              `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`             // Computed, where this resource class lives.
	Name            string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // The name of class. A resource class' package_name and name combine to form `class`.
	IntentType      *ResourceClass_Type `protobuf:"bytes,3,opt,name=intent_type,json=intentType,proto3" json:"intent_type,omitempty"`                // How is the resource defined.
	InstanceType    *ResourceClass_Type `protobuf:"bytes,4,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`          // And what generated configuration is attached to each instance.
	DefaultProvider string              `protobuf:"bytes,5,opt,name=default_provider,json=defaultProvider,proto3" json:"default_provider,omitempty"` // If set, a resource that refers to this class and doesn't specify a provider will default to this default.
}

func (x *ResourceClass) Reset() {
	*x = ResourceClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceClass) ProtoMessage() {}

func (x *ResourceClass) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceClass.ProtoReflect.Descriptor instead.
func (*ResourceClass) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceClass) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *ResourceClass) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceClass) GetIntentType() *ResourceClass_Type {
	if x != nil {
		return x.IntentType
	}
	return nil
}

func (x *ResourceClass) GetInstanceType() *ResourceClass_Type {
	if x != nil {
		return x.InstanceType
	}
	return nil
}

func (x *ResourceClass) GetDefaultProvider() string {
	if x != nil {
		return x.DefaultProvider
	}
	return ""
}

type ResourceProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName                            string                            `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`             // Computed, where this provider lives.
	ProvidesClass                          *PackageRef                       `protobuf:"bytes,2,opt,name=provides_class,json=providesClass,proto3" json:"provides_class,omitempty"`       // The resource this provider supports.
	InitializedWith                        *Invocation                       `protobuf:"bytes,3,opt,name=initialized_with,json=initializedWith,proto3" json:"initialized_with,omitempty"` // Run this invocation to create the resource. This yields an instantiation during the execution phase.
	PrepareWith                            *Invocation                       `protobuf:"bytes,5,opt,name=prepare_with,json=prepareWith,proto3" json:"prepare_with,omitempty"`             // Create the resource during planning phase.
	ResourceInput                          []*ResourceProvider_ResourceInput `protobuf:"bytes,8,rep,name=resource_input,json=resourceInput,proto3" json:"resource_input,omitempty"`
	ResourcePack                           *ResourcePack                     `protobuf:"bytes,4,opt,name=resource_pack,json=resourcePack,proto3" json:"resource_pack,omitempty"`                                                                                     // Resources this provider depends on in order to instantiate its own resource.
	ResourceInstanceFromAvailableClasses   []*PackageRef                     `protobuf:"bytes,6,rep,name=resource_instance_from_available_classes,json=resourceInstanceFromAvailableClasses,proto3" json:"resource_instance_from_available_classes,omitempty"`       // Resource classes that an invocation may produce (any instance of a different class will be rejected).
	ResourceInstanceFromAvailableProviders []*PackageRef                     `protobuf:"bytes,7,rep,name=resource_instance_from_available_providers,json=resourceInstanceFromAvailableProviders,proto3" json:"resource_instance_from_available_providers,omitempty"` // Resource providers that an invocation may reference (any instance of a different class will be rejected).
}

func (x *ResourceProvider) Reset() {
	*x = ResourceProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceProvider) ProtoMessage() {}

func (x *ResourceProvider) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceProvider.ProtoReflect.Descriptor instead.
func (*ResourceProvider) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceProvider) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *ResourceProvider) GetProvidesClass() *PackageRef {
	if x != nil {
		return x.ProvidesClass
	}
	return nil
}

func (x *ResourceProvider) GetInitializedWith() *Invocation {
	if x != nil {
		return x.InitializedWith
	}
	return nil
}

func (x *ResourceProvider) GetPrepareWith() *Invocation {
	if x != nil {
		return x.PrepareWith
	}
	return nil
}

func (x *ResourceProvider) GetResourceInput() []*ResourceProvider_ResourceInput {
	if x != nil {
		return x.ResourceInput
	}
	return nil
}

func (x *ResourceProvider) GetResourcePack() *ResourcePack {
	if x != nil {
		return x.ResourcePack
	}
	return nil
}

func (x *ResourceProvider) GetResourceInstanceFromAvailableClasses() []*PackageRef {
	if x != nil {
		return x.ResourceInstanceFromAvailableClasses
	}
	return nil
}

func (x *ResourceProvider) GetResourceInstanceFromAvailableProviders() []*PackageRef {
	if x != nil {
		return x.ResourceInstanceFromAvailableProviders
	}
	return nil
}

type ResourcePack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef      []*PackageRef       `protobuf:"bytes,1,rep,name=resource_ref,json=resourceRef,proto3" json:"resource_ref,omitempty"`
	ResourceInstance []*ResourceInstance `protobuf:"bytes,2,rep,name=resource_instance,json=resourceInstance,proto3" json:"resource_instance,omitempty"`
}

func (x *ResourcePack) Reset() {
	*x = ResourcePack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcePack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePack) ProtoMessage() {}

func (x *ResourcePack) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePack.ProtoReflect.Descriptor instead.
func (*ResourcePack) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{3}
}

func (x *ResourcePack) GetResourceRef() []*PackageRef {
	if x != nil {
		return x.ResourceRef
	}
	return nil
}

func (x *ResourcePack) GetResourceInstance() []*ResourceInstance {
	if x != nil {
		return x.ResourceInstance
	}
	return nil
}

type ResourceInstance_InputResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *PackageRef `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResourceRef *PackageRef `protobuf:"bytes,2,opt,name=resource_ref,json=resourceRef,proto3" json:"resource_ref,omitempty"`
}

func (x *ResourceInstance_InputResource) Reset() {
	*x = ResourceInstance_InputResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInstance_InputResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInstance_InputResource) ProtoMessage() {}

func (x *ResourceInstance_InputResource) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInstance_InputResource.ProtoReflect.Descriptor instead.
func (*ResourceInstance_InputResource) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourceInstance_InputResource) GetName() *PackageRef {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ResourceInstance_InputResource) GetResourceRef() *PackageRef {
	if x != nil {
		return x.ResourceRef
	}
	return nil
}

type ResourceClass_Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoSource string `protobuf:"bytes,1,opt,name=proto_source,json=protoSource,proto3" json:"proto_source,omitempty"`
	ProtoType   string `protobuf:"bytes,2,opt,name=proto_type,json=protoType,proto3" json:"proto_type,omitempty"`
}

func (x *ResourceClass_Type) Reset() {
	*x = ResourceClass_Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceClass_Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceClass_Type) ProtoMessage() {}

func (x *ResourceClass_Type) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceClass_Type.ProtoReflect.Descriptor instead.
func (*ResourceClass_Type) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResourceClass_Type) GetProtoSource() string {
	if x != nil {
		return x.ProtoSource
	}
	return ""
}

func (x *ResourceClass_Type) GetProtoType() string {
	if x != nil {
		return x.ProtoType
	}
	return ""
}

type ResourceProvider_ResourceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *PackageRef `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Class *PackageRef `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *ResourceProvider_ResourceInput) Reset() {
	*x = ResourceProvider_ResourceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceProvider_ResourceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceProvider_ResourceInput) ProtoMessage() {}

func (x *ResourceProvider_ResourceInput) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceProvider_ResourceInput.ProtoReflect.Descriptor instead.
func (*ResourceProvider_ResourceInput) Descriptor() ([]byte, []int) {
	return file_schema_resource_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ResourceProvider_ResourceInput) GetName() *PackageRef {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ResourceProvider_ResourceInput) GetClass() *PackageRef {
	if x != nil {
		return x.Class
	}
	return nil
}

var File_schema_resource_proto protoreflect.FileDescriptor

var file_schema_resource_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x05, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3e,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x58,
	0x0a, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x84, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x22,
	0xcf, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x4a, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x48, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x92, 0x06, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x48, 0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x0c, 0x70, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x57, 0x69, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x75, 0x0a, 0x28, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x24, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x79, 0x0a, 0x2a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x26, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x77, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x31,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x40, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x50, 0x0a, 0x11, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_resource_proto_rawDescOnce sync.Once
	file_schema_resource_proto_rawDescData = file_schema_resource_proto_rawDesc
)

func file_schema_resource_proto_rawDescGZIP() []byte {
	file_schema_resource_proto_rawDescOnce.Do(func() {
		file_schema_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_resource_proto_rawDescData)
	})
	return file_schema_resource_proto_rawDescData
}

var file_schema_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_schema_resource_proto_goTypes = []interface{}{
	(*ResourceInstance)(nil),               // 0: foundation.schema.ResourceInstance
	(*ResourceClass)(nil),                  // 1: foundation.schema.ResourceClass
	(*ResourceProvider)(nil),               // 2: foundation.schema.ResourceProvider
	(*ResourcePack)(nil),                   // 3: foundation.schema.ResourcePack
	(*ResourceInstance_InputResource)(nil), // 4: foundation.schema.ResourceInstance.InputResource
	(*ResourceClass_Type)(nil),             // 5: foundation.schema.ResourceClass.Type
	(*ResourceProvider_ResourceInput)(nil), // 6: foundation.schema.ResourceProvider.ResourceInput
	(*PackageRef)(nil),                     // 7: foundation.schema.PackageRef
	(*anypb.Any)(nil),                      // 8: google.protobuf.Any
	(*Invocation)(nil),                     // 9: foundation.schema.Invocation
}
var file_schema_resource_proto_depIdxs = []int32{
	7,  // 0: foundation.schema.ResourceInstance.class:type_name -> foundation.schema.PackageRef
	8,  // 1: foundation.schema.ResourceInstance.intent:type_name -> google.protobuf.Any
	9,  // 2: foundation.schema.ResourceInstance.intent_from:type_name -> foundation.schema.Invocation
	4,  // 3: foundation.schema.ResourceInstance.input_resource:type_name -> foundation.schema.ResourceInstance.InputResource
	5,  // 4: foundation.schema.ResourceClass.intent_type:type_name -> foundation.schema.ResourceClass.Type
	5,  // 5: foundation.schema.ResourceClass.instance_type:type_name -> foundation.schema.ResourceClass.Type
	7,  // 6: foundation.schema.ResourceProvider.provides_class:type_name -> foundation.schema.PackageRef
	9,  // 7: foundation.schema.ResourceProvider.initialized_with:type_name -> foundation.schema.Invocation
	9,  // 8: foundation.schema.ResourceProvider.prepare_with:type_name -> foundation.schema.Invocation
	6,  // 9: foundation.schema.ResourceProvider.resource_input:type_name -> foundation.schema.ResourceProvider.ResourceInput
	3,  // 10: foundation.schema.ResourceProvider.resource_pack:type_name -> foundation.schema.ResourcePack
	7,  // 11: foundation.schema.ResourceProvider.resource_instance_from_available_classes:type_name -> foundation.schema.PackageRef
	7,  // 12: foundation.schema.ResourceProvider.resource_instance_from_available_providers:type_name -> foundation.schema.PackageRef
	7,  // 13: foundation.schema.ResourcePack.resource_ref:type_name -> foundation.schema.PackageRef
	0,  // 14: foundation.schema.ResourcePack.resource_instance:type_name -> foundation.schema.ResourceInstance
	7,  // 15: foundation.schema.ResourceInstance.InputResource.name:type_name -> foundation.schema.PackageRef
	7,  // 16: foundation.schema.ResourceInstance.InputResource.resource_ref:type_name -> foundation.schema.PackageRef
	7,  // 17: foundation.schema.ResourceProvider.ResourceInput.name:type_name -> foundation.schema.PackageRef
	7,  // 18: foundation.schema.ResourceProvider.ResourceInput.class:type_name -> foundation.schema.PackageRef
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_schema_resource_proto_init() }
func file_schema_resource_proto_init() {
	if File_schema_resource_proto != nil {
		return
	}
	file_schema_package_proto_init()
	file_schema_provision_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInstance); i {
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
		file_schema_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceClass); i {
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
		file_schema_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceProvider); i {
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
		file_schema_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcePack); i {
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
		file_schema_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInstance_InputResource); i {
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
		file_schema_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceClass_Type); i {
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
		file_schema_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceProvider_ResourceInput); i {
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
			RawDescriptor: file_schema_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_resource_proto_goTypes,
		DependencyIndexes: file_schema_resource_proto_depIdxs,
		MessageInfos:      file_schema_resource_proto_msgTypes,
	}.Build()
	File_schema_resource_proto = out.File
	file_schema_resource_proto_rawDesc = nil
	file_schema_resource_proto_goTypes = nil
	file_schema_resource_proto_depIdxs = nil
}
