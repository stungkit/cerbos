// Copyright 2021 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cerbos/engine/v1/engine.proto

package enginev1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Resource  *Resource  `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Principal *Principal `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Actions   []string   `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *CheckInput) Reset() {
	*x = CheckInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_engine_v1_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInput) ProtoMessage() {}

func (x *CheckInput) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_engine_v1_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInput.ProtoReflect.Descriptor instead.
func (*CheckInput) Descriptor() ([]byte, []int) {
	return file_cerbos_engine_v1_engine_proto_rawDescGZIP(), []int{0}
}

func (x *CheckInput) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CheckInput) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *CheckInput) GetPrincipal() *Principal {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *CheckInput) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type CheckOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId             string                               `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ResourceId            string                               `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Actions               map[string]*CheckOutput_ActionEffect `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EffectiveDerivedRoles []string                             `protobuf:"bytes,4,rep,name=effective_derived_roles,json=effectiveDerivedRoles,proto3" json:"effective_derived_roles,omitempty"`
}

func (x *CheckOutput) Reset() {
	*x = CheckOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_engine_v1_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOutput) ProtoMessage() {}

func (x *CheckOutput) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_engine_v1_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOutput.ProtoReflect.Descriptor instead.
func (*CheckOutput) Descriptor() ([]byte, []int) {
	return file_cerbos_engine_v1_engine_proto_rawDescGZIP(), []int{1}
}

func (x *CheckOutput) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CheckOutput) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CheckOutput) GetActions() map[string]*CheckOutput_ActionEffect {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *CheckOutput) GetEffectiveDerivedRoles() []string {
	if x != nil {
		return x.EffectiveDerivedRoles
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind          string                     `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	PolicyVersion string                     `protobuf:"bytes,2,opt,name=policy_version,json=policyVersion,proto3" json:"policy_version,omitempty"`
	Id            string                     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Attr          map[string]*structpb.Value `protobuf:"bytes,4,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_engine_v1_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_engine_v1_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_cerbos_engine_v1_engine_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Resource) GetPolicyVersion() string {
	if x != nil {
		return x.PolicyVersion
	}
	return ""
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetAttr() map[string]*structpb.Value {
	if x != nil {
		return x.Attr
	}
	return nil
}

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PolicyVersion string                     `protobuf:"bytes,2,opt,name=policy_version,json=policyVersion,proto3" json:"policy_version,omitempty"`
	Roles         []string                   `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Attr          map[string]*structpb.Value `protobuf:"bytes,4,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_engine_v1_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_engine_v1_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_cerbos_engine_v1_engine_proto_rawDescGZIP(), []int{3}
}

func (x *Principal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Principal) GetPolicyVersion() string {
	if x != nil {
		return x.PolicyVersion
	}
	return ""
}

func (x *Principal) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Principal) GetAttr() map[string]*structpb.Value {
	if x != nil {
		return x.Attr
	}
	return nil
}

type CheckOutput_ActionEffect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Effect v1.Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=cerbos.effect.v1.Effect" json:"effect,omitempty"`
	Policy string    `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *CheckOutput_ActionEffect) Reset() {
	*x = CheckOutput_ActionEffect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerbos_engine_v1_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOutput_ActionEffect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOutput_ActionEffect) ProtoMessage() {}

func (x *CheckOutput_ActionEffect) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_engine_v1_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOutput_ActionEffect.ProtoReflect.Descriptor instead.
func (*CheckOutput_ActionEffect) Descriptor() ([]byte, []int) {
	return file_cerbos_engine_v1_engine_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CheckOutput_ActionEffect) GetEffect() v1.Effect {
	if x != nil {
		return x.Effect
	}
	return v1.Effect(0)
}

func (x *CheckOutput_ActionEffect) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

var File_cerbos_engine_v1_engine_proto protoreflect.FileDescriptor

var file_cerbos_engine_v1_engine_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x0c, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x0c, 0xe2, 0x41,
	0x01, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x12, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x0b, 0x92,
	0x01, 0x08, 0x18, 0x01, 0x22, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x8d, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x1a, 0x58, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x66, 0x0a, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xae, 0x06, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0xe2, 0x01, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0xcd, 0x01, 0x92, 0x41, 0x7c, 0x32, 0x29, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6b, 0x69, 0x6e, 0x64,
	0x20, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x2e,
	0x4a, 0x0d, 0x22, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x3a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x3f, 0x5e, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b,
	0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x5d, 0x2a, 0x28,
	0x5c, 0x3a, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b, 0x3a,
	0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x5d, 0x2a, 0x29, 0x2a,
	0x24, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x47, 0x72, 0x45, 0x10, 0x01, 0x32, 0x41, 0x5e, 0x5b,
	0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72,
	0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x2f, 0x5d, 0x2a, 0x28, 0x5c, 0x3a, 0x5b,
	0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72,
	0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x2f, 0x5d, 0x2a, 0x29, 0x2a, 0x24, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xdd, 0x01, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb5,
	0x01, 0x92, 0x41, 0x99, 0x01, 0x32, 0x7c, 0x54, 0x68, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65,
	0x20, 0x74, 0x6f, 0x20, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2c, 0x20, 0x77, 0x69, 0x6c,
	0x6c, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4a, 0x09, 0x22, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x8a, 0x01,
	0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24, 0xe2, 0x41,
	0x01, 0x01, 0xfa, 0x42, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72,
	0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x34, 0x92, 0x41, 0x26, 0x32, 0x1b, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4a, 0x07, 0x22, 0x58, 0x58, 0x31, 0x32, 0x35, 0x22, 0xe2, 0x41, 0x01, 0x02,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0xc5, 0x01, 0x0a, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x8a, 0x01, 0x92, 0x41, 0x7f, 0x32, 0x64, 0x4b, 0x61, 0x79, 0x2d, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x70, 0x61, 0x69, 0x72, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x75, 0x61, 0x6c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x74,
	0x68, 0x61, 0x74, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x64, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x20, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x17, 0x7b, 0x22,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x62, 0x75, 0x67, 0x73, 0x5f, 0x62, 0x75,
	0x6e, 0x6e, 0x79, 0x22, 0x7d, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x18, 0x01, 0x52, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x1a, 0x4f, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xd5, 0x06, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x12, 0x41, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31,
	0x92, 0x41, 0x23, 0x32, 0x13, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x4a, 0x0c, 0x22, 0x62, 0x75, 0x67, 0x73, 0x5f,
	0x62, 0x75, 0x6e, 0x6e, 0x79, 0x22, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0xdd, 0x01, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb5,
	0x01, 0x92, 0x41, 0x99, 0x01, 0x32, 0x7c, 0x54, 0x68, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65,
	0x20, 0x74, 0x6f, 0x20, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2c, 0x20, 0x77, 0x69, 0x6c,
	0x6c, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4a, 0x09, 0x22, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x8a, 0x01,
	0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24, 0xe2, 0x41,
	0x01, 0x01, 0xfa, 0x42, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72,
	0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xb0, 0x01, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x99, 0x01, 0x92, 0x41, 0x6f, 0x32, 0x46, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x20, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x66, 0x72,
	0x6f, 0x6d, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x4a, 0x08, 0x5b, 0x22, 0x75, 0x73, 0x65, 0x72, 0x22, 0x5d, 0x8a, 0x01, 0x11,
	0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x2d, 0x5c, 0x2e, 0x5d, 0x2b,
	0x24, 0xa0, 0x01, 0x14, 0xa8, 0x01, 0x01, 0xb0, 0x01, 0x01, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42,
	0x20, 0x92, 0x01, 0x1d, 0x08, 0x01, 0x10, 0x14, 0x18, 0x01, 0x22, 0x15, 0x72, 0x13, 0x32, 0x11,
	0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x2d, 0x5c, 0x2e, 0x5d, 0x2b,
	0x24, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0xc5, 0x01, 0x0a, 0x04, 0x61, 0x74, 0x74,
	0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x89,
	0x01, 0x92, 0x41, 0x7e, 0x32, 0x65, 0x4b, 0x65, 0x79, 0x2d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20,
	0x70, 0x61, 0x69, 0x72, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x74, 0x68,
	0x61, 0x74, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x64, 0x20, 0x64, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x15, 0x7b, 0x22, 0x62,
	0x65, 0x74, 0x61, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75,
	0x65, 0x7d, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x18, 0x01, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72,
	0x1a, 0x4f, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x59, 0x92, 0x41, 0x56, 0x0a, 0x54, 0x32, 0x52, 0x41, 0x20, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20,
	0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x0a, 0x18,
	0x64, 0x65, 0x76, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x76, 0x31, 0xaa, 0x02, 0x14, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e,
	0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cerbos_engine_v1_engine_proto_rawDescOnce sync.Once
	file_cerbos_engine_v1_engine_proto_rawDescData = file_cerbos_engine_v1_engine_proto_rawDesc
)

func file_cerbos_engine_v1_engine_proto_rawDescGZIP() []byte {
	file_cerbos_engine_v1_engine_proto_rawDescOnce.Do(func() {
		file_cerbos_engine_v1_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerbos_engine_v1_engine_proto_rawDescData)
	})
	return file_cerbos_engine_v1_engine_proto_rawDescData
}

var file_cerbos_engine_v1_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cerbos_engine_v1_engine_proto_goTypes = []interface{}{
	(*CheckInput)(nil),               // 0: cerbos.engine.v1.CheckInput
	(*CheckOutput)(nil),              // 1: cerbos.engine.v1.CheckOutput
	(*Resource)(nil),                 // 2: cerbos.engine.v1.Resource
	(*Principal)(nil),                // 3: cerbos.engine.v1.Principal
	(*CheckOutput_ActionEffect)(nil), // 4: cerbos.engine.v1.CheckOutput.ActionEffect
	nil,                              // 5: cerbos.engine.v1.CheckOutput.ActionsEntry
	nil,                              // 6: cerbos.engine.v1.Resource.AttrEntry
	nil,                              // 7: cerbos.engine.v1.Principal.AttrEntry
	(v1.Effect)(0),                   // 8: cerbos.effect.v1.Effect
	(*structpb.Value)(nil),           // 9: google.protobuf.Value
}
var file_cerbos_engine_v1_engine_proto_depIdxs = []int32{
	2, // 0: cerbos.engine.v1.CheckInput.resource:type_name -> cerbos.engine.v1.Resource
	3, // 1: cerbos.engine.v1.CheckInput.principal:type_name -> cerbos.engine.v1.Principal
	5, // 2: cerbos.engine.v1.CheckOutput.actions:type_name -> cerbos.engine.v1.CheckOutput.ActionsEntry
	6, // 3: cerbos.engine.v1.Resource.attr:type_name -> cerbos.engine.v1.Resource.AttrEntry
	7, // 4: cerbos.engine.v1.Principal.attr:type_name -> cerbos.engine.v1.Principal.AttrEntry
	8, // 5: cerbos.engine.v1.CheckOutput.ActionEffect.effect:type_name -> cerbos.effect.v1.Effect
	4, // 6: cerbos.engine.v1.CheckOutput.ActionsEntry.value:type_name -> cerbos.engine.v1.CheckOutput.ActionEffect
	9, // 7: cerbos.engine.v1.Resource.AttrEntry.value:type_name -> google.protobuf.Value
	9, // 8: cerbos.engine.v1.Principal.AttrEntry.value:type_name -> google.protobuf.Value
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cerbos_engine_v1_engine_proto_init() }
func file_cerbos_engine_v1_engine_proto_init() {
	if File_cerbos_engine_v1_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cerbos_engine_v1_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInput); i {
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
		file_cerbos_engine_v1_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOutput); i {
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
		file_cerbos_engine_v1_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_cerbos_engine_v1_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
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
		file_cerbos_engine_v1_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOutput_ActionEffect); i {
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
			RawDescriptor: file_cerbos_engine_v1_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_engine_v1_engine_proto_goTypes,
		DependencyIndexes: file_cerbos_engine_v1_engine_proto_depIdxs,
		MessageInfos:      file_cerbos_engine_v1_engine_proto_msgTypes,
	}.Build()
	File_cerbos_engine_v1_engine_proto = out.File
	file_cerbos_engine_v1_engine_proto_rawDesc = nil
	file_cerbos_engine_v1_engine_proto_goTypes = nil
	file_cerbos_engine_v1_engine_proto_depIdxs = nil
}
