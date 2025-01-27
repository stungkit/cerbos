// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: cerbos/effect/v1/effect.proto

package effectv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Effect int32

const (
	Effect_EFFECT_UNSPECIFIED Effect = 0
	Effect_EFFECT_ALLOW       Effect = 1
	Effect_EFFECT_DENY        Effect = 2
	Effect_EFFECT_NO_MATCH    Effect = 3
)

// Enum value maps for Effect.
var (
	Effect_name = map[int32]string{
		0: "EFFECT_UNSPECIFIED",
		1: "EFFECT_ALLOW",
		2: "EFFECT_DENY",
		3: "EFFECT_NO_MATCH",
	}
	Effect_value = map[string]int32{
		"EFFECT_UNSPECIFIED": 0,
		"EFFECT_ALLOW":       1,
		"EFFECT_DENY":        2,
		"EFFECT_NO_MATCH":    3,
	}
)

func (x Effect) Enum() *Effect {
	p := new(Effect)
	*p = x
	return p
}

func (x Effect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Effect) Descriptor() protoreflect.EnumDescriptor {
	return file_cerbos_effect_v1_effect_proto_enumTypes[0].Descriptor()
}

func (Effect) Type() protoreflect.EnumType {
	return &file_cerbos_effect_v1_effect_proto_enumTypes[0]
}

func (x Effect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Effect.Descriptor instead.
func (Effect) EnumDescriptor() ([]byte, []int) {
	return file_cerbos_effect_v1_effect_proto_rawDescGZIP(), []int{0}
}

var File_cerbos_effect_v1_effect_proto protoreflect.FileDescriptor

var file_cerbos_effect_v1_effect_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2a, 0x58, 0x0a, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x45,
	0x46, 0x46, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x5f,
	0x44, 0x45, 0x4e, 0x59, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54,
	0x5f, 0x4e, 0x4f, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x42, 0x6f, 0x0a, 0x18, 0x64,
	0x65, 0x76, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x76, 0x31, 0xaa, 0x02, 0x14, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x41,
	0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_cerbos_effect_v1_effect_proto_rawDescOnce sync.Once
	file_cerbos_effect_v1_effect_proto_rawDescData []byte
)

func file_cerbos_effect_v1_effect_proto_rawDescGZIP() []byte {
	file_cerbos_effect_v1_effect_proto_rawDescOnce.Do(func() {
		file_cerbos_effect_v1_effect_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cerbos_effect_v1_effect_proto_rawDesc), len(file_cerbos_effect_v1_effect_proto_rawDesc)))
	})
	return file_cerbos_effect_v1_effect_proto_rawDescData
}

var file_cerbos_effect_v1_effect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cerbos_effect_v1_effect_proto_goTypes = []any{
	(Effect)(0), // 0: cerbos.effect.v1.Effect
}
var file_cerbos_effect_v1_effect_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cerbos_effect_v1_effect_proto_init() }
func file_cerbos_effect_v1_effect_proto_init() {
	if File_cerbos_effect_v1_effect_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cerbos_effect_v1_effect_proto_rawDesc), len(file_cerbos_effect_v1_effect_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_effect_v1_effect_proto_goTypes,
		DependencyIndexes: file_cerbos_effect_v1_effect_proto_depIdxs,
		EnumInfos:         file_cerbos_effect_v1_effect_proto_enumTypes,
	}.Build()
	File_cerbos_effect_v1_effect_proto = out.File
	file_cerbos_effect_v1_effect_proto_goTypes = nil
	file_cerbos_effect_v1_effect_proto_depIdxs = nil
}
