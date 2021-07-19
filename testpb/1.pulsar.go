package testpb

import (
	binary "encoding/binary"
	"errors"
	fmt "fmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	bits "math/bits"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var (
	// Interface guards to verify each message implements proto message interface
	_ protoreflect.Message = &A{}
	_ protoreflect.Message = &B{}
)

type A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum        Enumeration `protobuf:"varint,1,opt,name=enum,proto3,enum=Enumeration" json:"enum,omitempty" yaml:"enum"`
	SomeBoolean bool        `protobuf:"varint,2,opt,name=some_boolean,json=someBoolean,proto3" json:"some_boolean,omitempty" yaml:"some_boolean"`
	INT32       int32       `protobuf:"varint,3,opt,name=INT32,proto3" json:"INT32,omitempty" yaml:"INT32"`
	SINT32      int32       `protobuf:"zigzag32,4,opt,name=SINT32,proto3" json:"SINT32,omitempty" yaml:"SINT32"`
	UINT32      uint32      `protobuf:"varint,5,opt,name=UINT32,proto3" json:"UINT32,omitempty" yaml:"UINT32"`
	INT64       int64       `protobuf:"varint,6,opt,name=INT64,proto3" json:"INT64,omitempty" yaml:"INT64"`
	SING64      int64       `protobuf:"zigzag64,7,opt,name=SING64,proto3" json:"SING64,omitempty" yaml:"SING64"`
	UINT64      uint64      `protobuf:"varint,8,opt,name=UINT64,proto3" json:"UINT64,omitempty" yaml:"UINT64"`
	SFIXED32    int32       `protobuf:"fixed32,9,opt,name=SFIXED32,proto3" json:"SFIXED32,omitempty" yaml:"SFIXED32"`
	FIXED32     uint32      `protobuf:"fixed32,10,opt,name=FIXED32,proto3" json:"FIXED32,omitempty" yaml:"FIXED32"`
	FLOAT       float32     `protobuf:"fixed32,11,opt,name=FLOAT,proto3" json:"FLOAT,omitempty" yaml:"FLOAT"`
	SFIXED64    int64       `protobuf:"fixed64,12,opt,name=SFIXED64,proto3" json:"SFIXED64,omitempty" yaml:"SFIXED64"`
	FIXED64     uint64      `protobuf:"fixed64,13,opt,name=FIXED64,proto3" json:"FIXED64,omitempty" yaml:"FIXED64"`
	DOUBLE      float64     `protobuf:"fixed64,14,opt,name=DOUBLE,proto3" json:"DOUBLE,omitempty" yaml:"DOUBLE"`
	STRING      string      `protobuf:"bytes,15,opt,name=STRING,proto3" json:"STRING,omitempty" yaml:"STRING"`
	BYTES       []byte      `protobuf:"bytes,16,opt,name=BYTES,proto3" json:"BYTES,omitempty" yaml:"BYTES"`
	MESSAGE     *B          `protobuf:"bytes,17,opt,name=MESSAGE,proto3" json:"MESSAGE,omitempty" yaml:"MESSAGE"`
	MAP         *A_MAPEntry `protobuf:"bytes,18,rep,name=MAP,proto3" json:"MAP,omitempty" yaml:"MAP"`
	LIST        []*B        `protobuf:"bytes,19,rep,name=LIST,proto3" json:"LIST,omitempty" yaml:"LIST"`
	// Types that are assignable to ONEOF:
	//	*A_ONEOF_B
	//	*A_ONEOF_STRING
	ONEOF     isA_ONEOF     `protobuf_oneof:"ONEOF"`
	LIST_ENUM []Enumeration `protobuf:"varint,22,rep,packed,name=LIST_ENUM,json=LISTENUM,proto3,enum=Enumeration" json:"LIST_ENUM,omitempty" yaml:"LIST_ENUM"`
}

func (x *A) Reset() {
	*x = A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github.com_cosmos_cosmos - proto_testpb_1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (x *A) ProtoMessage() {}

func (x *A) ProtoReflect() protoreflect.Message {
	mi := &file_github.com_cosmos_cosmos - proto_testpb_1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *A) GetEnum() Enumeration {
	if x != nil {
		return x.Enum
	}
	var y Enumeration
	return y
}

func (x *A) GetSomeBoolean() bool {
	if x != nil {
		return x.SomeBoolean
	}
	var y bool
	return y
}

func (x *A) GetINT32() int32 {
	if x != nil {
		return x.INT32
	}
	var y int32
	return y
}

func (x *A) GetSINT32() int32 {
	if x != nil {
		return x.SINT32
	}
	var y int32
	return y
}

func (x *A) GetUINT32() uint32 {
	if x != nil {
		return x.UINT32
	}
	var y uint32
	return y
}

func (x *A) GetINT64() int64 {
	if x != nil {
		return x.INT64
	}
	var y int64
	return y
}

func (x *A) GetSING64() int64 {
	if x != nil {
		return x.SING64
	}
	var y int64
	return y
}

func (x *A) GetUINT64() uint64 {
	if x != nil {
		return x.UINT64
	}
	var y uint64
	return y
}

func (x *A) GetSFIXED32() int32 {
	if x != nil {
		return x.SFIXED32
	}
	var y int32
	return y
}

func (x *A) GetFIXED32() uint32 {
	if x != nil {
		return x.FIXED32
	}
	var y uint32
	return y
}

func (x *A) GetFLOAT() float32 {
	if x != nil {
		return x.FLOAT
	}
	var y float32
	return y
}

func (x *A) GetSFIXED64() int64 {
	if x != nil {
		return x.SFIXED64
	}
	var y int64
	return y
}

func (x *A) GetFIXED64() uint64 {
	if x != nil {
		return x.FIXED64
	}
	var y uint64
	return y
}

func (x *A) GetDOUBLE() float64 {
	if x != nil {
		return x.DOUBLE
	}
	var y float64
	return y
}

func (x *A) GetSTRING() string {
	if x != nil {
		return x.STRING
	}
	var y string
	return y
}

func (x *A) GetBYTES() []byte {
	if x != nil {
		return x.BYTES
	}
	var y []byte
	return y
}

func (x *A) GetMESSAGE() *B {
	if x != nil && x.MESSAGE != nil {
		return x.MESSAGE
	}
	var y *B
	return y
}

func (x *A) GetMAP() *A_MAPEntry {
	if x != nil {
		return x.MAP
	}
	var y *A_MAPEntry
	return y
}

func (x *A) GetLIST() []*B {
	if x != nil {
		return x.LIST
	}
	var y []*B
	return y
}

type isA_ONEOF interface {
	isA_ONEOF()
}

func (m *A) GetONEOF() isA_ONEOF {
	if m != nil {
		return m.ONEOF
	}
	return nil
}

type A_ONEOF_B struct {
	ONEOF_B *B
}

type A_ONEOF_STRING struct {
	ONEOF_STRING *B
}

func (*A_ONEOF_B) isA_ONEOF() {}

func (*A_ONEOF_STRING) isA_ONEOF() {}

func (x *A) GetONEOF_B() *B {
	if x, ok := x.GetONEOF().(*A_ONEOF_B); ok {
		return x.ONEOF_B
	}
	var y *B
	return y
}

func (x *A) GetONEOF_STRING() string {
	if x, ok := x.GetONEOF().(*A_ONEOF_STRING); ok {
		return x.ONEOF_STRING
	}
	var y string
	return y
}

func (x *A) GetLIST_ENUM() []Enumeration {
	if x != nil {
		return x.LIST_ENUM
	}
	var y []Enumeration
	return y
}

// returns the fast methods for the message
func (x A) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(A)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("A does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*A)
			if !ok {
				return protoiface.UnmarshalOutput{}, errors.New("A does not implement the protoreflect.Message interface")
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x A) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x A) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x A) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x A) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x A) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
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
func (x A) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x A) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
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
func (x *A) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
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
		return protoreflect.ValueOfMessage(x.MESSAGE)
	case "MAP":
	case "LIST":
	case "ONEOF_B":
		return protoreflect.ValueOfMessage(x.ONEOF_B)
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
func (x A) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
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
func (x A) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x A) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x A) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x A) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x A) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x A) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x A) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

type B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty" yaml:"x"`
}

func (x *B) Reset() {
	*x = B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github.com_cosmos_cosmos - proto_testpb_1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (x *B) ProtoMessage() {}

func (x *B) ProtoReflect() protoreflect.Message {
	mi := &file_github.com_cosmos_cosmos - proto_testpb_1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *B) GetX() string {
	if x != nil {
		return x.X
	}
	var y string
	return y
}

// returns the fast methods for the message
func (x B) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(B)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("B does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*B)
			if !ok {
				return protoiface.UnmarshalOutput{}, errors.New("B does not implement the protoreflect.Message interface")
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x B) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x B) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x B) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x B) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x B) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
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
func (x B) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x B) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *B) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
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
func (x B) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
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
func (x B) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x B) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x B) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x B) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x B) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x B) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x B) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

var File_testpb_1_proto protoreflect.FileDescriptor

var file_testpb_1_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa5, 0x05, 0x0a, 0x01, 0x41, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x6d, 0x65,
	0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e,
	0x54, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x36,
	0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x36, 0x34, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x12, 0x14, 0x0a,
	0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x46, 0x4c,
	0x4f, 0x41, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12,
	0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x4f, 0x55,
	0x42, 0x4c, 0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c,
	0x45, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x12,
	0x1c, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x02, 0x2e, 0x42, 0x52, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x12, 0x1d, 0x0a,
	0x03, 0x4d, 0x41, 0x50, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x2e, 0x4d,
	0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x4d, 0x41, 0x50, 0x12, 0x16, 0x0a, 0x04,
	0x4c, 0x49, 0x53, 0x54, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x04,
	0x4c, 0x49, 0x53, 0x54, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x42, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x48, 0x00, 0x52, 0x06, 0x4f, 0x4e, 0x45,
	0x4f, 0x46, 0x42, 0x12, 0x23, 0x0a, 0x0c, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x4e, 0x45,
	0x4f, 0x46, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x55, 0x4d, 0x1a, 0x3a, 0x0a, 0x08, 0x4d, 0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x02, 0x2e, 0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x22, 0x11, 0x0a, 0x01, 0x42, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x2a, 0x1f, 0x0a, 0x0b, 0x45,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_testpb_1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testpb_1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_testpb_1_proto_goTypes = []interface{}{
	(Enumeration)(0), // 0: Enumeration
	(*A)(nil),        // 1: A
	(*B)(nil),        // 2: B
	nil,              // 3: A.MAPEntry
}
var file_testpb_1_proto_depIdxs = []int32{
	0, // 0: A.enum:type_name -> Enumeration
	2, // 1: A.MESSAGE:type_name -> B
	3, // 2: A.MAP:type_name -> A.MAPEntry
	2, // 3: A.LIST:type_name -> B
	2, // 4: A.ONEOF_B:type_name -> B
	0, // 5: A.LIST_ENUM:type_name -> Enumeration
	2, // 6: A.MAPEntry.value:type_name -> B
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_testpb_1_proto_init() }
func file_testpb_1_proto_init() {
	if File_testpb_1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
	}
	file_testpb_1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*A_ONEOF_B)(nil),
		(*A_ONEOF_STRING)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testpb_1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testpb_1_proto_goTypes,
		DependencyIndexes: file_testpb_1_proto_depIdxs,
		EnumInfos:         file_testpb_1_proto_enumTypes,
		MessageInfos:      file_testpb_1_proto_msgTypes,
	}.Build()
	File_testpb_1_proto = out.File
	file_testpb_1_proto_rawDesc = nil
	file_testpb_1_proto_goTypes = nil
	file_testpb_1_proto_depIdxs = nil
}
func (m *A) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *A) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.LIST_ENUM) > 0 {
		var pksize2 int
		for _, num := range m.LIST_ENUM {
			pksize2 += sov(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num1 := range m.LIST_ENUM {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = encodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	if vtmsg, ok := m.ONEOF.(interface {
		MarshalTo([]byte) (int, error)
		Size() int
	}); ok {
		{
			size := vtmsg.Size()
			i -= size
			if _, err := vtmsg.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.LIST) > 0 {
		for iNdEx := len(m.LIST) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.LIST[iNdEx].MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x9a
		}
	}
	if len(m.MAP) > 0 {
		for k := range m.MAP {
			v := m.MAP[k]
			baseI := i
			size, err := v.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x92
		}
	}
	if m.MESSAGE != nil {
		size, err := m.MESSAGE.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.BYTES) > 0 {
		i -= len(m.BYTES)
		copy(dAtA[i:], m.BYTES)
		i = encodeVarint(dAtA, i, uint64(len(m.BYTES)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.STRING) > 0 {
		i -= len(m.STRING)
		copy(dAtA[i:], m.STRING)
		i = encodeVarint(dAtA, i, uint64(len(m.STRING)))
		i--
		dAtA[i] = 0x7a
	}
	if m.DOUBLE != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DOUBLE))))
		i--
		dAtA[i] = 0x71
	}
	if m.FIXED64 != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.FIXED64))
		i--
		dAtA[i] = 0x69
	}
	if m.SFIXED64 != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SFIXED64))
		i--
		dAtA[i] = 0x61
	}
	if m.FLOAT != 0 {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FLOAT))))
		i--
		dAtA[i] = 0x5d
	}
	if m.FIXED32 != 0 {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.FIXED32))
		i--
		dAtA[i] = 0x55
	}
	if m.SFIXED32 != 0 {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.SFIXED32))
		i--
		dAtA[i] = 0x4d
	}
	if m.UINT64 != 0 {
		i = encodeVarint(dAtA, i, uint64(m.UINT64))
		i--
		dAtA[i] = 0x40
	}
	if m.SING64 != 0 {
		i = encodeVarint(dAtA, i, uint64((uint64(m.SING64)<<1)^uint64((m.SING64>>63))))
		i--
		dAtA[i] = 0x38
	}
	if m.INT64 != 0 {
		i = encodeVarint(dAtA, i, uint64(m.INT64))
		i--
		dAtA[i] = 0x30
	}
	if m.UINT32 != 0 {
		i = encodeVarint(dAtA, i, uint64(m.UINT32))
		i--
		dAtA[i] = 0x28
	}
	if m.SINT32 != 0 {
		i = encodeVarint(dAtA, i, uint64((uint32(m.SINT32)<<1)^uint32((m.SINT32>>31))))
		i--
		dAtA[i] = 0x20
	}
	if m.INT32 != 0 {
		i = encodeVarint(dAtA, i, uint64(m.INT32))
		i--
		dAtA[i] = 0x18
	}
	if m.SomeBoolean {
		i--
		if m.SomeBoolean {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Enum != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Enum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *A_ONEOF_B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A_ONEOF_B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ONEOF_B != nil {
		size, err := m.ONEOF_B.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	return len(dAtA) - i, nil
}
func (m *A_ONEOF_STRING) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *A_ONEOF_STRING) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ONEOF_STRING)
	copy(dAtA[i:], m.ONEOF_STRING)
	i = encodeVarint(dAtA, i, uint64(len(m.ONEOF_STRING)))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xaa
	return len(dAtA) - i, nil
}
func (m *B) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarint(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *A) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enum != 0 {
		n += 1 + sov(uint64(m.Enum))
	}
	if m.SomeBoolean {
		n += 2
	}
	if m.INT32 != 0 {
		n += 1 + sov(uint64(m.INT32))
	}
	if m.SINT32 != 0 {
		n += 1 + soz(uint64(m.SINT32))
	}
	if m.UINT32 != 0 {
		n += 1 + sov(uint64(m.UINT32))
	}
	if m.INT64 != 0 {
		n += 1 + sov(uint64(m.INT64))
	}
	if m.SING64 != 0 {
		n += 1 + soz(uint64(m.SING64))
	}
	if m.UINT64 != 0 {
		n += 1 + sov(uint64(m.UINT64))
	}
	if m.SFIXED32 != 0 {
		n += 5
	}
	if m.FIXED32 != 0 {
		n += 5
	}
	if m.FLOAT != 0 {
		n += 5
	}
	if m.SFIXED64 != 0 {
		n += 9
	}
	if m.FIXED64 != 0 {
		n += 9
	}
	if m.DOUBLE != 0 {
		n += 9
	}
	l = len(m.STRING)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.BYTES)
	if l > 0 {
		n += 2 + l + sov(uint64(l))
	}
	if m.MESSAGE != nil {
		l = m.MESSAGE.Size()
		n += 2 + l + sov(uint64(l))
	}
	if len(m.MAP) > 0 {
		for k, v := range m.MAP {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			l += 1 + sov(uint64(l))
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + l
			n += mapEntrySize + 2 + sov(uint64(mapEntrySize))
		}
	}
	if len(m.LIST) > 0 {
		for _, e := range m.LIST {
			l = e.Size()
			n += 2 + l + sov(uint64(l))
		}
	}
	if vtmsg, ok := m.ONEOF.(interface{ Size() int }); ok {
		n += vtmsg.Size()
	}
	if len(m.LIST_ENUM) > 0 {
		l = 0
		for _, e := range m.LIST_ENUM {
			l += sov(uint64(e))
		}
		n += 2 + sov(uint64(l)) + l
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func (m *A_ONEOF_B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ONEOF_B != nil {
		l = m.ONEOF_B.Size()
		n += 2 + l + sov(uint64(l))
	}
	return n
}
func (m *A_ONEOF_STRING) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ONEOF_STRING)
	n += 2 + l + sov(uint64(l))
	return n
}
func (m *B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *A) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: A: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: A: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enum", wireType)
			}
			m.Enum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Enum |= Enumeration(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SomeBoolean", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SomeBoolean = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field INT32", wireType)
			}
			m.INT32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.INT32 |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SINT32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.SINT32 = v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UINT32", wireType)
			}
			m.UINT32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UINT32 |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field INT64", wireType)
			}
			m.INT64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.INT64 |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SING64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.SING64 = int64(v)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UINT64", wireType)
			}
			m.UINT64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UINT64 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field SFIXED32", wireType)
			}
			m.SFIXED32 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.SFIXED32 = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 10:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FIXED32", wireType)
			}
			m.FIXED32 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.FIXED32 = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FLOAT", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FLOAT = float32(math.Float32frombits(v))
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SFIXED64", wireType)
			}
			m.SFIXED64 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SFIXED64 = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 13:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field FIXED64", wireType)
			}
			m.FIXED64 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.FIXED64 = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 14:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DOUBLE", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DOUBLE = float64(math.Float64frombits(v))
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field STRING", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.STRING = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BYTES", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BYTES = append(m.BYTES[:0], dAtA[iNdEx:postIndex]...)
			if m.BYTES == nil {
				m.BYTES = []byte{}
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MESSAGE", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MESSAGE == nil {
				m.MESSAGE = &B{}
			}
			if err := m.MESSAGE.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MAP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MAP == nil {
				m.MAP = make(map[string]*B)
			}
			var mapkey string
			var mapvalue *B
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLength
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLength
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLength
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLength
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &B{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skip(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLength
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MAP[mapkey] = mapvalue
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LIST", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LIST = append(m.LIST, &B{})
			if err := m.LIST[len(m.LIST)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ONEOF_B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.ONEOF.(*A_ONEOF_B); ok {
				if err := oneof.ONEOF_B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &B{}
				if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.ONEOF = &A_ONEOF_B{v}
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ONEOF_STRING", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ONEOF = &A_ONEOF_STRING{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 22:
			if wireType == 0 {
				var v Enumeration
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Enumeration(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.LIST_ENUM = append(m.LIST_ENUM, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.LIST_ENUM) == 0 {
					m.LIST_ENUM = make([]Enumeration, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Enumeration
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Enumeration(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.LIST_ENUM = append(m.LIST_ENUM, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field LIST_ENUM", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)
