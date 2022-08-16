// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: google/actions/sdk/v2/interactionmodel/type/type.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Declaration of a custom type, as opposed to built-in types. Types can be
// assigned to slots in a scene or parameters of an intent's training phrases.
// Practically, Types can be thought of as enums.
// Note, type name is specified in the name of the file.
type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selection of sub type based on the type of matching to be done.
	//
	// Types that are assignable to SubType:
	//	*Type_Synonym
	//	*Type_RegularExpression
	//	*Type_FreeText
	SubType isType_SubType `protobuf_oneof:"sub_type"`
	// Set of exceptional words/phrases that shouldn't be matched by type.
	// Note: If word/phrase is matched by the type but listed as an exclusion it
	// won't be returned in parameter extraction result.
	// **This field is localizable.**
	Exclusions []string `protobuf:"bytes,4,rep,name=exclusions,proto3" json:"exclusions,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescGZIP(), []int{0}
}

func (m *Type) GetSubType() isType_SubType {
	if m != nil {
		return m.SubType
	}
	return nil
}

func (x *Type) GetSynonym() *SynonymType {
	if x, ok := x.GetSubType().(*Type_Synonym); ok {
		return x.Synonym
	}
	return nil
}

func (x *Type) GetRegularExpression() *RegularExpressionType {
	if x, ok := x.GetSubType().(*Type_RegularExpression); ok {
		return x.RegularExpression
	}
	return nil
}

func (x *Type) GetFreeText() *FreeTextType {
	if x, ok := x.GetSubType().(*Type_FreeText); ok {
		return x.FreeText
	}
	return nil
}

func (x *Type) GetExclusions() []string {
	if x != nil {
		return x.Exclusions
	}
	return nil
}

type isType_SubType interface {
	isType_SubType()
}

type Type_Synonym struct {
	// Synonyms type, which is essentially an enum.
	Synonym *SynonymType `protobuf:"bytes,1,opt,name=synonym,proto3,oneof"`
}

type Type_RegularExpression struct {
	// Regex type, allows regular expression matching.
	RegularExpression *RegularExpressionType `protobuf:"bytes,2,opt,name=regular_expression,json=regularExpression,proto3,oneof"`
}

type Type_FreeText struct {
	// FreeText type.
	FreeText *FreeTextType `protobuf:"bytes,3,opt,name=free_text,json=freeText,proto3,oneof"`
}

func (*Type_Synonym) isType_SubType() {}

func (*Type_RegularExpression) isType_SubType() {}

func (*Type_FreeText) isType_SubType() {}

var File_google_actions_sdk_v2_interactionmodel_type_type_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x73,
	0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x07, 0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x12, 0x73, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76,
	0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x0a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x8a, 0x01, 0x0a,
	0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x09, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_interactionmodel_type_type_proto_goTypes = []interface{}{
	(*Type)(nil),                  // 0: google.actions.sdk.v2.interactionmodel.type.Type
	(*SynonymType)(nil),           // 1: google.actions.sdk.v2.interactionmodel.type.SynonymType
	(*RegularExpressionType)(nil), // 2: google.actions.sdk.v2.interactionmodel.type.RegularExpressionType
	(*FreeTextType)(nil),          // 3: google.actions.sdk.v2.interactionmodel.type.FreeTextType
}
var file_google_actions_sdk_v2_interactionmodel_type_type_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.interactionmodel.type.Type.synonym:type_name -> google.actions.sdk.v2.interactionmodel.type.SynonymType
	2, // 1: google.actions.sdk.v2.interactionmodel.type.Type.regular_expression:type_name -> google.actions.sdk.v2.interactionmodel.type.RegularExpressionType
	3, // 2: google.actions.sdk.v2.interactionmodel.type.Type.free_text:type_name -> google.actions.sdk.v2.interactionmodel.type.FreeTextType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_type_type_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_type_type_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_type_type_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_type_free_text_type_proto_init()
	file_google_actions_sdk_v2_interactionmodel_type_regular_expression_type_proto_init()
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Type_Synonym)(nil),
		(*Type_RegularExpression)(nil),
		(*Type_FreeText)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_type_type_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_type_type_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_type_type_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_type_type_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_type_type_proto_depIdxs = nil
}
