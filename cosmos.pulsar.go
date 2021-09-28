package cosmos_proto

import (
	fmt "fmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

var (
	md_InterfaceService protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_proto_init()
	md_InterfaceService = File_cosmos_proto.Messages().ByName("InterfaceService")
}

var _ protoreflect.Message = (*fastReflection_InterfaceService)(nil)

type fastReflection_InterfaceService InterfaceService

func (x *InterfaceService) ProtoReflect() protoreflect.Message {
	return (*fastReflection_InterfaceService)(x)
}

func (x *InterfaceService) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_InterfaceService_messageType fastReflection_InterfaceService_messageType
var _ protoreflect.MessageType = fastReflection_InterfaceService_messageType{}

type fastReflection_InterfaceService_messageType struct{}

func (x fastReflection_InterfaceService_messageType) Zero() protoreflect.Message {
	return (*fastReflection_InterfaceService)(nil)
}
func (x fastReflection_InterfaceService_messageType) New() protoreflect.Message {
	return new(fastReflection_InterfaceService)
}
func (x fastReflection_InterfaceService_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceService
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_InterfaceService) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceService
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_InterfaceService) Type() protoreflect.MessageType {
	return _fastReflection_InterfaceService_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_InterfaceService) New() protoreflect.Message {
	return new(fastReflection_InterfaceService)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_InterfaceService) Interface() protoreflect.ProtoMessage {
	return (*InterfaceService)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_InterfaceService) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_InterfaceService) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceService) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_InterfaceService) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceService) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceService) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_InterfaceService) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.InterfaceService"))
		}
		panic(fmt.Errorf("message cosmos_proto.InterfaceService does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_InterfaceService) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos_proto.InterfaceService", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_InterfaceService) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceService) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_InterfaceService) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_InterfaceService) ProtoMethods() *protoiface.Methods {
	return nil
}

var (
	md_ImplementsInterface          protoreflect.MessageDescriptor
	fd_ImplementsInterface_fullname protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_proto_init()
	md_ImplementsInterface = File_cosmos_proto.Messages().ByName("ImplementsInterface")
	fd_ImplementsInterface_fullname = md_ImplementsInterface.Fields().ByName("fullname")
}

var _ protoreflect.Message = (*fastReflection_ImplementsInterface)(nil)

type fastReflection_ImplementsInterface ImplementsInterface

func (x *ImplementsInterface) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ImplementsInterface)(x)
}

func (x *ImplementsInterface) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ImplementsInterface_messageType fastReflection_ImplementsInterface_messageType
var _ protoreflect.MessageType = fastReflection_ImplementsInterface_messageType{}

type fastReflection_ImplementsInterface_messageType struct{}

func (x fastReflection_ImplementsInterface_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ImplementsInterface)(nil)
}
func (x fastReflection_ImplementsInterface_messageType) New() protoreflect.Message {
	return new(fastReflection_ImplementsInterface)
}
func (x fastReflection_ImplementsInterface_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ImplementsInterface
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ImplementsInterface) Descriptor() protoreflect.MessageDescriptor {
	return md_ImplementsInterface
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ImplementsInterface) Type() protoreflect.MessageType {
	return _fastReflection_ImplementsInterface_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ImplementsInterface) New() protoreflect.Message {
	return new(fastReflection_ImplementsInterface)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ImplementsInterface) Interface() protoreflect.ProtoMessage {
	return (*ImplementsInterface)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ImplementsInterface) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_ImplementsInterface_fullname, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ImplementsInterface) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		return x.Fullname != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ImplementsInterface) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		x.Fullname = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ImplementsInterface) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ImplementsInterface) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		x.Fullname = value.String()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ImplementsInterface) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		panic(fmt.Errorf("field fullname of message cosmos_proto.ImplementsInterface is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ImplementsInterface) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos_proto.ImplementsInterface.fullname":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos_proto.ImplementsInterface"))
		}
		panic(fmt.Errorf("message cosmos_proto.ImplementsInterface does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ImplementsInterface) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos_proto.ImplementsInterface", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ImplementsInterface) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ImplementsInterface) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ImplementsInterface) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ImplementsInterface) ProtoMethods() *protoiface.Methods {
	return nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.7
// source: cosmos.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InterfaceService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InterfaceService) Reset() {
	*x = InterfaceService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceService) ProtoMessage() {}

// Deprecated: Use InterfaceService.ProtoReflect.Descriptor instead.
func (*InterfaceService) Descriptor() ([]byte, []int) {
	return file_cosmos_proto_rawDescGZIP(), []int{0}
}

type ImplementsInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
}

func (x *ImplementsInterface) Reset() {
	*x = ImplementsInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplementsInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplementsInterface) ProtoMessage() {}

// Deprecated: Use ImplementsInterface.ProtoReflect.Descriptor instead.
func (*ImplementsInterface) Descriptor() ([]byte, []int) {
	return file_cosmos_proto_rawDescGZIP(), []int{1}
}

func (x *ImplementsInterface) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

var file_cosmos_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*ImplementsInterface)(nil),
		Field:         93001,
		Name:          "cosmos_proto.implements_interface",
		Tag:           "bytes,93001,opt,name=implements_interface",
		Filename:      "cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*InterfaceService)(nil),
		Field:         93001,
		Name:          "cosmos_proto.interface_service",
		Tag:           "bytes,93001,opt,name=interface_service",
		Filename:      "cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93001,
		Name:          "cosmos_proto.accepts_interface",
		Tag:           "bytes,93001,opt,name=accepts_interface",
		Filename:      "cosmos.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         93002,
		Name:          "cosmos_proto.scalar",
		Tag:           "bytes,93002,opt,name=scalar",
		Filename:      "cosmos.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// implements_interface is used to annotate interface implementations
	//
	// optional cosmos_proto.ImplementsInterface implements_interface = 93001;
	E_ImplementsInterface = &file_cosmos_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional cosmos_proto.InterfaceService interface_service = 93001;
	E_InterfaceService = &file_cosmos_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// accepts_interface is used to annote fields that accept interfaces
	//
	// optional string accepts_interface = 93001;
	E_AcceptsInterface = &file_cosmos_proto_extTypes[2]
	// scalar is used to define scalar types
	//
	// optional string scalar = 93002;
	E_Scalar = &file_cosmos_proto_extTypes[3]
)

var File_cosmos_proto protoreflect.FileDescriptor

var file_cosmos_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x77, 0x0a, 0x14, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9,
	0xd6, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x13, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x3a, 0x6e,
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x10, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x4c,
	0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xc9, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x3a, 0x37, 0x0a, 0x06,
	0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0xd6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_proto_rawDescOnce sync.Once
	file_cosmos_proto_rawDescData = file_cosmos_proto_rawDesc
)

func file_cosmos_proto_rawDescGZIP() []byte {
	file_cosmos_proto_rawDescOnce.Do(func() {
		file_cosmos_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_proto_rawDescData)
	})
	return file_cosmos_proto_rawDescData
}

var file_cosmos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_proto_goTypes = []interface{}{
	(*InterfaceService)(nil),            // 0: cosmos_proto.InterfaceService
	(*ImplementsInterface)(nil),         // 1: cosmos_proto.ImplementsInterface
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_cosmos_proto_depIdxs = []int32{
	2, // 0: cosmos_proto.implements_interface:extendee -> google.protobuf.MessageOptions
	3, // 1: cosmos_proto.interface_service:extendee -> google.protobuf.ServiceOptions
	4, // 2: cosmos_proto.accepts_interface:extendee -> google.protobuf.FieldOptions
	4, // 3: cosmos_proto.scalar:extendee -> google.protobuf.FieldOptions
	1, // 4: cosmos_proto.implements_interface:type_name -> cosmos_proto.ImplementsInterface
	0, // 5: cosmos_proto.interface_service:type_name -> cosmos_proto.InterfaceService
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	4, // [4:6] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_proto_init() }
func file_cosmos_proto_init() {
	if File_cosmos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceService); i {
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
		file_cosmos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplementsInterface); i {
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
			RawDescriptor: file_cosmos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_proto_goTypes,
		DependencyIndexes: file_cosmos_proto_depIdxs,
		MessageInfos:      file_cosmos_proto_msgTypes,
		ExtensionInfos:    file_cosmos_proto_extTypes,
	}.Build()
	File_cosmos_proto = out.File
	file_cosmos_proto_rawDesc = nil
	file_cosmos_proto_goTypes = nil
	file_cosmos_proto_depIdxs = nil
}
