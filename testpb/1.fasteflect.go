package testpb

import (
	fmt "fmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

func (x *A) slowProtoReflect() protoreflect.Message {
	mi := &file_testpb_1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _ protoreflect.Message = (*fastReflection_A)(nil)

type fastReflection_A A

func (x *A) ProtoReflect() protoreflect.Message {
	return (*fastReflection_A)(x)
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_A) Descriptor() protoreflect.MessageDescriptor {
	return (*A)(x).slowProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_A) Type() protoreflect.MessageType {
	return (*A)(x).slowProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_A) New() protoreflect.Message {
	return (*A)(x).slowProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_A) Interface() protoreflect.ProtoMessage {
	return (*A)(x).slowProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_A) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	(*A)(x).slowProtoReflect().Range(f)
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
func (x *fastReflection_A) Has(descriptor protoreflect.FieldDescriptor) bool {
	return (*A)(x).slowProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) Clear(descriptor protoreflect.FieldDescriptor) {
	(*A)(x).slowProtoReflect().Clear(descriptor)
}

type _A_19_list struct {
	list []*B
}

var _ protoreflect.List = (*_A_19_list)(nil)

func (x *_A_19_list) Len() int {
	return len(x.list)
}

func (x *_A_19_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage(x.list[i].ProtoReflect())
}

func (x *_A_19_list) Set(i int, value protoreflect.Value) {
	unwrapped := value.Message()
	concreteValue := unwrapped.Interface().(*B)
	x.list[i] = concreteValue
}

func (x *_A_19_list) Append(value protoreflect.Value) {
	unwrapped := value.Message()
	concreteValue := unwrapped.Interface().(*B)
	x.list = append(x.list, concreteValue)
}

func (x *_A_19_list) AppendMutable() protoreflect.Value {
	v := new(B)
	x.list = append(x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_19_list) Truncate(n int) {
	for i := n; i < len(x.list); i++ {
		x.list[i] = nil
	}
	x.list = x.list[:n]
}

func (x *_A_19_list) NewElement() protoreflect.Value {
	v := new(B)
	x.list = append(x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_A_19_list) IsValid() bool {
	return x.list == nil || len(x.list) == 0
}

type _A_22_list struct {
	list []Enumeration
}

var _ protoreflect.List = (*_A_22_list)(nil)

func (x *_A_22_list) Len() int {
	return len(x.list)
}

func (x *_A_22_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.list[i]))
}

func (x *_A_22_list) Set(i int, value protoreflect.Value) {
	unwrapped := value.Enum()
	concreteValue := (Enumeration)(unwrapped)
	x.list[i] = concreteValue
}

func (x *_A_22_list) Append(value protoreflect.Value) {
	unwrapped := value.Enum()
	concreteValue := (Enumeration)(unwrapped)
	x.list = append(x.list, concreteValue)
}

func (x *_A_22_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message A at list field LIST_ENUM as it is not of Message kind"))
}

func (x *_A_22_list) Truncate(n int) {
	x.list = x.list[:n]
}

func (x *_A_22_list) NewElement() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_A_22_list) IsValid() bool {
	return x.list == nil || len(x.list) == 0
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_A) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.Name() {
	case "enum":
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Enum))
	case "some_boolean":
		return protoreflect.ValueOfBool(x.SomeBoolean)
	case "INT32":
		return protoreflect.ValueOfInt32(x.INT32)
	case "SINT32":
		return protoreflect.ValueOfInt32(x.SINT32)
	case "UINT32":
		return protoreflect.ValueOfUint32(x.UINT32)
	case "INT64":
		return protoreflect.ValueOfInt64(x.INT64)
	case "SING64":
		return protoreflect.ValueOfInt64(x.SING64)
	case "UINT64":
		return protoreflect.ValueOfUint64(x.UINT64)
	case "SFIXED32":
		return protoreflect.ValueOfInt32(x.SFIXED32)
	case "FIXED32":
		return protoreflect.ValueOfUint32(x.FIXED32)
	case "FLOAT":
		return protoreflect.ValueOfFloat32(x.FLOAT)
	case "SFIXED64":
		return protoreflect.ValueOfInt64(x.SFIXED64)
	case "FIXED64":
		return protoreflect.ValueOfUint64(x.FIXED64)
	case "DOUBLE":
		return protoreflect.ValueOfFloat64(x.DOUBLE)
	case "STRING":
		return protoreflect.ValueOfString(x.STRING)
	case "BYTES":
		return protoreflect.ValueOfBytes(x.BYTES)
	case "MESSAGE":
		return protoreflect.ValueOfMessage(x.MESSAGE.ProtoReflect())
	case "MAP":
	case "LIST":
	case "ONEOF_B":
		return protoreflect.ValueOfMessage(x.ONEOF_B.ProtoReflect())
	case "ONEOF_STRING":
		return protoreflect.ValueOfString(x.ONEOF_STRING)
	case "LIST_ENUM":
	default:
		panic(fmt.Errorf("message A does not contain field %s", descriptor.Name()))
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
func (x *fastReflection_A) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	(*A)(x).slowProtoReflect().Set(descriptor, value)
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
func (x *fastReflection_A) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return (*A)(x).slowProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_A) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return (*A)(x).slowProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_A) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return (*A)(x).slowProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_A) GetUnknown() protoreflect.RawFields {
	return (*A)(x).slowProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_A) SetUnknown(fields protoreflect.RawFields) {
	(*A)(x).slowProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_A) IsValid() bool {
	return (*A)(x).slowProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_A) ProtoMethods() *protoiface.Methods {
	return (*A)(x).slowProtoReflect().ProtoMethods()
}

func (x *B) slowProtoReflect() protoreflect.Message {
	mi := &file_testpb_1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _ protoreflect.Message = (*fastReflection_B)(nil)

type fastReflection_B B

func (x *B) ProtoReflect() protoreflect.Message {
	return (*fastReflection_B)(x)
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_B) Descriptor() protoreflect.MessageDescriptor {
	return (*B)(x).slowProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_B) Type() protoreflect.MessageType {
	return (*B)(x).slowProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_B) New() protoreflect.Message {
	return (*B)(x).slowProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_B) Interface() protoreflect.ProtoMessage {
	return (*B)(x).slowProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_B) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	(*B)(x).slowProtoReflect().Range(f)
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
func (x *fastReflection_B) Has(descriptor protoreflect.FieldDescriptor) bool {
	return (*B)(x).slowProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) Clear(descriptor protoreflect.FieldDescriptor) {
	(*B)(x).slowProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_B) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.Name() {
	case "x":
		return protoreflect.ValueOfString(x.X)
	default:
		panic(fmt.Errorf("message B does not contain field %s", descriptor.Name()))
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
func (x *fastReflection_B) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	(*B)(x).slowProtoReflect().Set(descriptor, value)
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
func (x *fastReflection_B) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return (*B)(x).slowProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_B) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return (*B)(x).slowProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_B) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return (*B)(x).slowProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_B) GetUnknown() protoreflect.RawFields {
	return (*B)(x).slowProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_B) SetUnknown(fields protoreflect.RawFields) {
	(*B)(x).slowProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_B) IsValid() bool {
	return (*B)(x).slowProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_B) ProtoMethods() *protoiface.Methods {
	return (*B)(x).slowProtoReflect().ProtoMethods()
}
