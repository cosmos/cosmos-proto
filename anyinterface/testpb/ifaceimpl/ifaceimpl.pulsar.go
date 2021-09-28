package ifaceimpl

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-proto/anyinterface/testpb/iface"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgSend              protoreflect.MessageDescriptor
	fd_MsgSend_from_address protoreflect.FieldDescriptor
	fd_MsgSend_to_address   protoreflect.FieldDescriptor
	fd_MsgSend_amount       protoreflect.FieldDescriptor
)

func init() {
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_init()
	md_MsgSend = File_anyinterface_testpb_ifaceimpl_ifaceimpl_proto.Messages().ByName("MsgSend")
	fd_MsgSend_from_address = md_MsgSend.Fields().ByName("from_address")
	fd_MsgSend_to_address = md_MsgSend.Fields().ByName("to_address")
	fd_MsgSend_amount = md_MsgSend.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_MsgSend)(nil)

type fastReflection_MsgSend MsgSend

func (x *MsgSend) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSend)(x)
}

func (x *MsgSend) slowProtoReflect() protoreflect.Message {
	mi := &file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSend_messageType fastReflection_MsgSend_messageType
var _ protoreflect.MessageType = fastReflection_MsgSend_messageType{}

type fastReflection_MsgSend_messageType struct{}

func (x fastReflection_MsgSend_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSend)(nil)
}
func (x fastReflection_MsgSend_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}
func (x fastReflection_MsgSend_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSend) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSend
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSend) Type() protoreflect.MessageType {
	return _fastReflection_MsgSend_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSend) New() protoreflect.Message {
	return new(fastReflection_MsgSend)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSend) Interface() protoreflect.ProtoMessage {
	return (*MsgSend)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSend) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.FromAddress != "" {
		value := protoreflect.ValueOfString(x.FromAddress)
		if !f(fd_MsgSend_from_address, value) {
			return
		}
	}
	if x.ToAddress != "" {
		value := protoreflect.ValueOfString(x.ToAddress)
		if !f(fd_MsgSend_to_address, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_MsgSend_amount, value) {
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
func (x *fastReflection_MsgSend) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		return x.FromAddress != ""
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		return x.ToAddress != ""
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		return x.Amount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		x.FromAddress = ""
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		x.ToAddress = ""
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		x.Amount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSend) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		value := x.FromAddress
		return protoreflect.ValueOfString(value)
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		value := x.ToAddress
		return protoreflect.ValueOfString(value)
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSend) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		x.FromAddress = value.String()
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		x.ToAddress = value.String()
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		x.Amount = value.String()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSend) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		panic(fmt.Errorf("field from_address of message anyinterface.testpb.ifaceimpl.MsgSend is not mutable"))
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		panic(fmt.Errorf("field to_address of message anyinterface.testpb.ifaceimpl.MsgSend is not mutable"))
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		panic(fmt.Errorf("field amount of message anyinterface.testpb.ifaceimpl.MsgSend is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSend) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "anyinterface.testpb.ifaceimpl.MsgSend.from_address":
		return protoreflect.ValueOfString("")
	case "anyinterface.testpb.ifaceimpl.MsgSend.to_address":
		return protoreflect.ValueOfString("")
	case "anyinterface.testpb.ifaceimpl.MsgSend.amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: anyinterface.testpb.ifaceimpl.MsgSend"))
		}
		panic(fmt.Errorf("message anyinterface.testpb.ifaceimpl.MsgSend does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSend) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in anyinterface.testpb.ifaceimpl.MsgSend", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSend) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSend) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSend) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSend) ProtoMethods() *protoiface.Methods {
	return nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.7
// source: anyinterface/testpb/ifaceimpl/ifaceimpl.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgSend) Reset() {
	*x = MsgSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend) ProtoMessage() {}

// Deprecated: Use MsgSend.ProtoReflect.Descriptor instead.
func (*MsgSend) Descriptor() ([]byte, []int) {
	return file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSend) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgSend) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgSend) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_anyinterface_testpb_ifaceimpl_ifaceimpl_proto protoreflect.FileDescriptor

var file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x6e, 0x79, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x66, 0x61, 0x63, 0x65, 0x69, 0x6d, 0x70, 0x6c, 0x2f,
	0x69, 0x66, 0x61, 0x63, 0x65, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x61, 0x6e, 0x79, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x2e, 0x69, 0x66, 0x61, 0x63, 0x65, 0x69, 0x6d, 0x70, 0x6c, 0x1a, 0x0c,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x6e,
	0x79, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x62, 0x2f, 0x69, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x23, 0xca, 0xb4, 0x2d, 0x1f, 0x0a,
	0x1d, 0x61, 0x6e, 0x79, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x2e, 0x69, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x6e, 0x79, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x66, 0x61, 0x63, 0x65, 0x69, 0x6d, 0x70, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescOnce sync.Once
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescData = file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDesc
)

func file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescGZIP() []byte {
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescOnce.Do(func() {
		file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescData = protoimpl.X.CompressGZIP(file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescData)
	})
	return file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDescData
}

var file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_goTypes = []interface{}{
	(*MsgSend)(nil), // 0: anyinterface.testpb.ifaceimpl.MsgSend
}
var file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_init() }
func file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_init() {
	if File_anyinterface_testpb_ifaceimpl_ifaceimpl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend); i {
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
			RawDescriptor: file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_goTypes,
		DependencyIndexes: file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_depIdxs,
		MessageInfos:      file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_msgTypes,
	}.Build()
	File_anyinterface_testpb_ifaceimpl_ifaceimpl_proto = out.File
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_rawDesc = nil
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_goTypes = nil
	file_anyinterface_testpb_ifaceimpl_ifaceimpl_proto_depIdxs = nil
}
