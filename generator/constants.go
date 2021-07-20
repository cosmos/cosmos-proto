package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Standard library dependencies.
const (
	base64Package  = protogen.GoImportPath("encoding/base64")
	mathPackage    = protogen.GoImportPath("math")
	reflectPackage = protogen.GoImportPath("reflect")
	sortPackage    = protogen.GoImportPath("sort")
	stringsPackage = protogen.GoImportPath("strings")
	syncPackage    = protogen.GoImportPath("sync")
	timePackage    = protogen.GoImportPath("time")
	utf8Package    = protogen.GoImportPath("unicode/utf8")
)

// from protobuf - protoc-gen-go internal/genid/goname.go
// Go names of implementation-specific struct fields in generated messages.
const (
	State_goname = "state"

	SizeCache_goname  = "sizeCache"
	SizeCacheA_goname = "XXX_sizecache"

	WeakFields_goname  = "weakFields"
	WeakFieldsA_goname = "XXX_weak"

	UnknownFields_goname  = "unknownFields"
	UnknownFieldsA_goname = "XXX_unrecognized"

	ExtensionFields_goname  = "extensionFields"
	ExtensionFieldsA_goname = "XXX_InternalExtensions"
	ExtensionFieldsB_goname = "XXX_extensions"

	WeakFieldPrefix_goname = "XXX_weak_"
)

const File_google_protobuf_field_mask_proto = "google/protobuf/field_mask.proto"

// Names for google.protobuf.FieldMask.
const (
	FieldMask_message_name     protoreflect.Name     = "FieldMask"
	FieldMask_message_fullname protoreflect.FullName = "google.protobuf.FieldMask"
)

// Field names for google.protobuf.FieldMask.
const (
	FieldMask_Paths_field_name protoreflect.Name = "paths"

	FieldMask_Paths_field_fullname protoreflect.FullName = "google.protobuf.FieldMask.paths"
)

// Field numbers for google.protobuf.FieldMask.
const (
	FieldMask_Paths_field_number protoreflect.FieldNumber = 1
)

const File_google_protobuf_struct_proto = "google/protobuf/struct.proto"

// Full and short names for google.protobuf.NullValue.
const (
	NullValue_enum_fullname = "google.protobuf.NullValue"
	NullValue_enum_name     = "NullValue"
)

// Names for google.protobuf.Struct.
const (
	Struct_message_name     protoreflect.Name     = "Struct"
	Struct_message_fullname protoreflect.FullName = "google.protobuf.Struct"
)

// Field names for google.protobuf.Struct.
const (
	Struct_Fields_field_name protoreflect.Name = "fields"

	Struct_Fields_field_fullname protoreflect.FullName = "google.protobuf.Struct.fields"
)

// Field numbers for google.protobuf.Struct.
const (
	Struct_Fields_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_message_name     protoreflect.Name     = "FieldsEntry"
	Struct_FieldsEntry_message_fullname protoreflect.FullName = "google.protobuf.Struct.FieldsEntry"
)

// Field names for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_Key_field_name   protoreflect.Name = "key"
	Struct_FieldsEntry_Value_field_name protoreflect.Name = "value"

	Struct_FieldsEntry_Key_field_fullname   protoreflect.FullName = "google.protobuf.Struct.FieldsEntry.key"
	Struct_FieldsEntry_Value_field_fullname protoreflect.FullName = "google.protobuf.Struct.FieldsEntry.value"
)

// Field numbers for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_Key_field_number   protoreflect.FieldNumber = 1
	Struct_FieldsEntry_Value_field_number protoreflect.FieldNumber = 2
)

// Names for google.protobuf.Value.
const (
	Value_message_name     protoreflect.Name     = "Value"
	Value_message_fullname protoreflect.FullName = "google.protobuf.Value"
)

// Field names for google.protobuf.Value.
const (
	Value_NullValue_field_name   protoreflect.Name = "null_value"
	Value_NumberValue_field_name protoreflect.Name = "number_value"
	Value_StringValue_field_name protoreflect.Name = "string_value"
	Value_BoolValue_field_name   protoreflect.Name = "bool_value"
	Value_StructValue_field_name protoreflect.Name = "struct_value"
	Value_ListValue_field_name   protoreflect.Name = "list_value"

	Value_NullValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.null_value"
	Value_NumberValue_field_fullname protoreflect.FullName = "google.protobuf.Value.number_value"
	Value_StringValue_field_fullname protoreflect.FullName = "google.protobuf.Value.string_value"
	Value_BoolValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.bool_value"
	Value_StructValue_field_fullname protoreflect.FullName = "google.protobuf.Value.struct_value"
	Value_ListValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.list_value"
)

// Field numbers for google.protobuf.Value.
const (
	Value_NullValue_field_number   protoreflect.FieldNumber = 1
	Value_NumberValue_field_number protoreflect.FieldNumber = 2
	Value_StringValue_field_number protoreflect.FieldNumber = 3
	Value_BoolValue_field_number   protoreflect.FieldNumber = 4
	Value_StructValue_field_number protoreflect.FieldNumber = 5
	Value_ListValue_field_number   protoreflect.FieldNumber = 6
)

// Oneof names for google.protobuf.Value.
const (
	Value_Kind_oneof_name protoreflect.Name = "kind"

	Value_Kind_oneof_fullname protoreflect.FullName = "google.protobuf.Value.kind"
)

// Names for google.protobuf.ListValue.
const (
	ListValue_message_name     protoreflect.Name     = "ListValue"
	ListValue_message_fullname protoreflect.FullName = "google.protobuf.ListValue"
)

// Field names for google.protobuf.ListValue.
const (
	ListValue_Values_field_name protoreflect.Name = "values"

	ListValue_Values_field_fullname protoreflect.FullName = "google.protobuf.ListValue.values"
)

// Field numbers for google.protobuf.ListValue.
const (
	ListValue_Values_field_number protoreflect.FieldNumber = 1
)

const File_google_protobuf_duration_proto = "google/protobuf/duration.proto"

// Names for google.protobuf.Duration.
const (
	Duration_message_name     protoreflect.Name     = "Duration"
	Duration_message_fullname protoreflect.FullName = "google.protobuf.Duration"
)

// Field names for google.protobuf.Duration.
const (
	Duration_Seconds_field_name protoreflect.Name = "seconds"
	Duration_Nanos_field_name   protoreflect.Name = "nanos"

	Duration_Seconds_field_fullname protoreflect.FullName = "google.protobuf.Duration.seconds"
	Duration_Nanos_field_fullname   protoreflect.FullName = "google.protobuf.Duration.nanos"
)

// Field numbers for google.protobuf.Duration.
const (
	Duration_Seconds_field_number protoreflect.FieldNumber = 1
	Duration_Nanos_field_number   protoreflect.FieldNumber = 2
)

const File_google_protobuf_timestamp_proto = "google/protobuf/timestamp.proto"

// Names for google.protobuf.Timestamp.
const (
	Timestamp_message_name     protoreflect.Name     = "Timestamp"
	Timestamp_message_fullname protoreflect.FullName = "google.protobuf.Timestamp"
)

// Field names for google.protobuf.Timestamp.
const (
	Timestamp_Seconds_field_name protoreflect.Name = "seconds"
	Timestamp_Nanos_field_name   protoreflect.Name = "nanos"

	Timestamp_Seconds_field_fullname protoreflect.FullName = "google.protobuf.Timestamp.seconds"
	Timestamp_Nanos_field_fullname   protoreflect.FullName = "google.protobuf.Timestamp.nanos"
)

// Field numbers for google.protobuf.Timestamp.
const (
	Timestamp_Seconds_field_number protoreflect.FieldNumber = 1
	Timestamp_Nanos_field_number   protoreflect.FieldNumber = 2
)

const File_google_protobuf_any_proto = "google/protobuf/any.proto"

// Names for google.protobuf.Any.
const (
	Any_message_name     protoreflect.Name     = "Any"
	Any_message_fullname protoreflect.FullName = "google.protobuf.Any"
)

// Field names for google.protobuf.Any.
const (
	Any_TypeUrl_field_name protoreflect.Name = "type_url"
	Any_Value_field_name   protoreflect.Name = "value"

	Any_TypeUrl_field_fullname protoreflect.FullName = "google.protobuf.Any.type_url"
	Any_Value_field_fullname   protoreflect.FullName = "google.protobuf.Any.value"
)

// Field numbers for google.protobuf.Any.
const (
	Any_TypeUrl_field_number protoreflect.FieldNumber = 1
	Any_Value_field_number   protoreflect.FieldNumber = 2
)

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

const File_google_protobuf_wrappers_proto = "google/protobuf/wrappers.proto"

// Names for google.protobuf.DoubleValue.
const (
	DoubleValue_message_name     protoreflect.Name     = "DoubleValue"
	DoubleValue_message_fullname protoreflect.FullName = "google.protobuf.DoubleValue"
)

// Field names for google.protobuf.DoubleValue.
const (
	DoubleValue_Value_field_name protoreflect.Name = "value"

	DoubleValue_Value_field_fullname protoreflect.FullName = "google.protobuf.DoubleValue.value"
)

// Field numbers for google.protobuf.DoubleValue.
const (
	DoubleValue_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.FloatValue.
const (
	FloatValue_message_name     protoreflect.Name     = "FloatValue"
	FloatValue_message_fullname protoreflect.FullName = "google.protobuf.FloatValue"
)

// Field names for google.protobuf.FloatValue.
const (
	FloatValue_Value_field_name protoreflect.Name = "value"

	FloatValue_Value_field_fullname protoreflect.FullName = "google.protobuf.FloatValue.value"
)

// Field numbers for google.protobuf.FloatValue.
const (
	FloatValue_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.Int64Value.
const (
	Int64Value_message_name     protoreflect.Name     = "Int64Value"
	Int64Value_message_fullname protoreflect.FullName = "google.protobuf.Int64Value"
)

// Field names for google.protobuf.Int64Value.
const (
	Int64Value_Value_field_name protoreflect.Name = "value"

	Int64Value_Value_field_fullname protoreflect.FullName = "google.protobuf.Int64Value.value"
)

// Field numbers for google.protobuf.Int64Value.
const (
	Int64Value_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.UInt64Value.
const (
	UInt64Value_message_name     protoreflect.Name     = "UInt64Value"
	UInt64Value_message_fullname protoreflect.FullName = "google.protobuf.UInt64Value"
)

// Field names for google.protobuf.UInt64Value.
const (
	UInt64Value_Value_field_name protoreflect.Name = "value"

	UInt64Value_Value_field_fullname protoreflect.FullName = "google.protobuf.UInt64Value.value"
)

// Field numbers for google.protobuf.UInt64Value.
const (
	UInt64Value_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.Int32Value.
const (
	Int32Value_message_name     protoreflect.Name     = "Int32Value"
	Int32Value_message_fullname protoreflect.FullName = "google.protobuf.Int32Value"
)

// Field names for google.protobuf.Int32Value.
const (
	Int32Value_Value_field_name protoreflect.Name = "value"

	Int32Value_Value_field_fullname protoreflect.FullName = "google.protobuf.Int32Value.value"
)

// Field numbers for google.protobuf.Int32Value.
const (
	Int32Value_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.UInt32Value.
const (
	UInt32Value_message_name     protoreflect.Name     = "UInt32Value"
	UInt32Value_message_fullname protoreflect.FullName = "google.protobuf.UInt32Value"
)

// Field names for google.protobuf.UInt32Value.
const (
	UInt32Value_Value_field_name protoreflect.Name = "value"

	UInt32Value_Value_field_fullname protoreflect.FullName = "google.protobuf.UInt32Value.value"
)

// Field numbers for google.protobuf.UInt32Value.
const (
	UInt32Value_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.BoolValue.
const (
	BoolValue_message_name     protoreflect.Name     = "BoolValue"
	BoolValue_message_fullname protoreflect.FullName = "google.protobuf.BoolValue"
)

// Field names for google.protobuf.BoolValue.
const (
	BoolValue_Value_field_name protoreflect.Name = "value"

	BoolValue_Value_field_fullname protoreflect.FullName = "google.protobuf.BoolValue.value"
)

// Field numbers for google.protobuf.BoolValue.
const (
	BoolValue_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.StringValue.
const (
	StringValue_message_name     protoreflect.Name     = "StringValue"
	StringValue_message_fullname protoreflect.FullName = "google.protobuf.StringValue"
)

// Field names for google.protobuf.StringValue.
const (
	StringValue_Value_field_name protoreflect.Name = "value"

	StringValue_Value_field_fullname protoreflect.FullName = "google.protobuf.StringValue.value"
)

// Field numbers for google.protobuf.StringValue.
const (
	StringValue_Value_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.BytesValue.
const (
	BytesValue_message_name     protoreflect.Name     = "BytesValue"
	BytesValue_message_fullname protoreflect.FullName = "google.protobuf.BytesValue"
)

// Field names for google.protobuf.BytesValue.
const (
	BytesValue_Value_field_name protoreflect.Name = "value"

	BytesValue_Value_field_fullname protoreflect.FullName = "google.protobuf.BytesValue.value"
)

// Field numbers for google.protobuf.BytesValue.
const (
	BytesValue_Value_field_number protoreflect.FieldNumber = 1
)
