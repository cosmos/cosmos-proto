package test3

import (
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	reflect "reflect"
	sort "sort"
	sync "sync"
)

var (
	md_NestedMessage             protoreflect.MessageDescriptor
	fd_NestedMessage_a           protoreflect.FieldDescriptor
	fd_NestedMessage_corecursive protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_proto_init()
	md_NestedMessage = File_internal_testprotos_test3_test_proto.Messages().ByName("NestedMessage")
	fd_NestedMessage_a = md_NestedMessage.Fields().ByName("a")
	fd_NestedMessage_corecursive = md_NestedMessage.Fields().ByName("corecursive")
}

var _ protoreflect.Message = (*fastReflection_NestedMessage)(nil)

type fastReflection_NestedMessage NestedMessage

func (x *NestedMessage) ProtoReflect() protoreflect.Message {
	return (*fastReflection_NestedMessage)(x)
}

func (x *NestedMessage) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_NestedMessage_messageType fastReflection_NestedMessage_messageType
var _ protoreflect.MessageType = fastReflection_NestedMessage_messageType{}

type fastReflection_NestedMessage_messageType struct{}

func (x fastReflection_NestedMessage_messageType) Zero() protoreflect.Message {
	return (*fastReflection_NestedMessage)(nil)
}
func (x fastReflection_NestedMessage_messageType) New() protoreflect.Message {
	return new(fastReflection_NestedMessage)
}
func (x fastReflection_NestedMessage_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_NestedMessage
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_NestedMessage) Descriptor() protoreflect.MessageDescriptor {
	return md_NestedMessage
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_NestedMessage) Type() protoreflect.MessageType {
	return _fastReflection_NestedMessage_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_NestedMessage) New() protoreflect.Message {
	return new(fastReflection_NestedMessage)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_NestedMessage) Interface() protoreflect.ProtoMessage {
	return (*NestedMessage)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_NestedMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.A != int32(0) {
		value := protoreflect.ValueOfInt32(x.A)
		if !f(fd_NestedMessage_a, value) {
			return
		}
	}
	if x.Corecursive != nil {
		value := protoreflect.ValueOfMessage(x.Corecursive.ProtoReflect())
		if !f(fd_NestedMessage_corecursive, value) {
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
func (x *fastReflection_NestedMessage) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.NestedMessage.a":
		return x.A != int32(0)
	case "goproto.proto.test3.NestedMessage.corecursive":
		return x.Corecursive != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NestedMessage) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.NestedMessage.a":
		x.A = int32(0)
	case "goproto.proto.test3.NestedMessage.corecursive":
		x.Corecursive = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_NestedMessage) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.NestedMessage.a":
		value := x.A
		return protoreflect.ValueOfInt32(value)
	case "goproto.proto.test3.NestedMessage.corecursive":
		value := x.Corecursive
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_NestedMessage) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.NestedMessage.a":
		x.A = int32(value.Int())
	case "goproto.proto.test3.NestedMessage.corecursive":
		x.Corecursive = value.Message().Interface().(*TestAllTypes)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", fd.FullName()))
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
func (x *fastReflection_NestedMessage) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.NestedMessage.corecursive":
		if x.Corecursive == nil {
			x.Corecursive = new(TestAllTypes)
		}
		return protoreflect.ValueOfMessage(x.Corecursive.ProtoReflect())
	case "goproto.proto.test3.NestedMessage.a":
		panic(fmt.Errorf("field a of message goproto.proto.test3.NestedMessage is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_NestedMessage) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.NestedMessage.a":
		return protoreflect.ValueOfInt32(int32(0))
	case "goproto.proto.test3.NestedMessage.corecursive":
		m := new(TestAllTypes)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.NestedMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.NestedMessage does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_NestedMessage) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.NestedMessage", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_NestedMessage) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_NestedMessage) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_NestedMessage) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_NestedMessage) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*NestedMessage)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.A != 0 {
			n += 1 + runtime.Sov(uint64(x.A))
		}
		if x.Corecursive != nil {
			l = options.Size(x.Corecursive)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*NestedMessage)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Corecursive != nil {
			encoded, err := options.Marshal(x.Corecursive)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.A != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.A))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*NestedMessage)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NestedMessage: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: NestedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
				}
				x.A = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.A |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Corecursive", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Corecursive == nil {
					x.Corecursive = &TestAllTypes{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Corecursive); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_TestAllTypes_31_list)(nil)

type _TestAllTypes_31_list struct {
	list *[]int32
}

func (x *_TestAllTypes_31_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_31_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt32((*x.list)[i])
}

func (x *_TestAllTypes_31_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_31_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_31_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedInt32 as it is not of Message kind"))
}

func (x *_TestAllTypes_31_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_31_list) NewElement() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_31_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_32_list)(nil)

type _TestAllTypes_32_list struct {
	list *[]int64
}

func (x *_TestAllTypes_32_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_32_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt64((*x.list)[i])
}

func (x *_TestAllTypes_32_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_32_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_32_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedInt64 as it is not of Message kind"))
}

func (x *_TestAllTypes_32_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_32_list) NewElement() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_32_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_33_list)(nil)

type _TestAllTypes_33_list struct {
	list *[]uint32
}

func (x *_TestAllTypes_33_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_33_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint32((*x.list)[i])
}

func (x *_TestAllTypes_33_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_33_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_33_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedUint32 as it is not of Message kind"))
}

func (x *_TestAllTypes_33_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_33_list) NewElement() protoreflect.Value {
	v := uint32(0)
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_33_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_34_list)(nil)

type _TestAllTypes_34_list struct {
	list *[]uint64
}

func (x *_TestAllTypes_34_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_34_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint64((*x.list)[i])
}

func (x *_TestAllTypes_34_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_34_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_34_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedUint64 as it is not of Message kind"))
}

func (x *_TestAllTypes_34_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_34_list) NewElement() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_34_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_35_list)(nil)

type _TestAllTypes_35_list struct {
	list *[]int32
}

func (x *_TestAllTypes_35_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_35_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt32((*x.list)[i])
}

func (x *_TestAllTypes_35_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_35_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_35_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedSint32 as it is not of Message kind"))
}

func (x *_TestAllTypes_35_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_35_list) NewElement() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_35_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_36_list)(nil)

type _TestAllTypes_36_list struct {
	list *[]int64
}

func (x *_TestAllTypes_36_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_36_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt64((*x.list)[i])
}

func (x *_TestAllTypes_36_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_36_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_36_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedSint64 as it is not of Message kind"))
}

func (x *_TestAllTypes_36_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_36_list) NewElement() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_36_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_37_list)(nil)

type _TestAllTypes_37_list struct {
	list *[]uint32
}

func (x *_TestAllTypes_37_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_37_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint32((*x.list)[i])
}

func (x *_TestAllTypes_37_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_37_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_37_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedFixed32 as it is not of Message kind"))
}

func (x *_TestAllTypes_37_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_37_list) NewElement() protoreflect.Value {
	v := uint32(0)
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_37_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_38_list)(nil)

type _TestAllTypes_38_list struct {
	list *[]uint64
}

func (x *_TestAllTypes_38_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_38_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfUint64((*x.list)[i])
}

func (x *_TestAllTypes_38_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_38_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_38_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedFixed64 as it is not of Message kind"))
}

func (x *_TestAllTypes_38_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_38_list) NewElement() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_38_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_39_list)(nil)

type _TestAllTypes_39_list struct {
	list *[]int32
}

func (x *_TestAllTypes_39_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_39_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt32((*x.list)[i])
}

func (x *_TestAllTypes_39_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_39_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_39_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedSfixed32 as it is not of Message kind"))
}

func (x *_TestAllTypes_39_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_39_list) NewElement() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_39_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_40_list)(nil)

type _TestAllTypes_40_list struct {
	list *[]int64
}

func (x *_TestAllTypes_40_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_40_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt64((*x.list)[i])
}

func (x *_TestAllTypes_40_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_40_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_40_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedSfixed64 as it is not of Message kind"))
}

func (x *_TestAllTypes_40_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_40_list) NewElement() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_40_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_41_list)(nil)

type _TestAllTypes_41_list struct {
	list *[]float32
}

func (x *_TestAllTypes_41_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_41_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfFloat32((*x.list)[i])
}

func (x *_TestAllTypes_41_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Float()
	concreteValue := (float32)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_41_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Float()
	concreteValue := (float32)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_41_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedFloat as it is not of Message kind"))
}

func (x *_TestAllTypes_41_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_41_list) NewElement() protoreflect.Value {
	v := float32(0)
	return protoreflect.ValueOfFloat32(v)
}

func (x *_TestAllTypes_41_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_42_list)(nil)

type _TestAllTypes_42_list struct {
	list *[]float64
}

func (x *_TestAllTypes_42_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_42_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfFloat64((*x.list)[i])
}

func (x *_TestAllTypes_42_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Float()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_42_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Float()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_42_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedDouble as it is not of Message kind"))
}

func (x *_TestAllTypes_42_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_42_list) NewElement() protoreflect.Value {
	v := float64(0)
	return protoreflect.ValueOfFloat64(v)
}

func (x *_TestAllTypes_42_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_43_list)(nil)

type _TestAllTypes_43_list struct {
	list *[]bool
}

func (x *_TestAllTypes_43_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_43_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBool((*x.list)[i])
}

func (x *_TestAllTypes_43_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bool()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_43_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bool()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_43_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedBool as it is not of Message kind"))
}

func (x *_TestAllTypes_43_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_43_list) NewElement() protoreflect.Value {
	v := false
	return protoreflect.ValueOfBool(v)
}

func (x *_TestAllTypes_43_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_44_list)(nil)

type _TestAllTypes_44_list struct {
	list *[]string
}

func (x *_TestAllTypes_44_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_44_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_TestAllTypes_44_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_44_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_44_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedString as it is not of Message kind"))
}

func (x *_TestAllTypes_44_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_44_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_TestAllTypes_44_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_45_list)(nil)

type _TestAllTypes_45_list struct {
	list *[][]byte
}

func (x *_TestAllTypes_45_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_45_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_TestAllTypes_45_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_45_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_45_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedBytes as it is not of Message kind"))
}

func (x *_TestAllTypes_45_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_45_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_TestAllTypes_45_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_48_list)(nil)

type _TestAllTypes_48_list struct {
	list *[]*NestedMessage
}

func (x *_TestAllTypes_48_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_48_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TestAllTypes_48_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*NestedMessage)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_48_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*NestedMessage)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_48_list) AppendMutable() protoreflect.Value {
	v := new(NestedMessage)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_48_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_48_list) NewElement() protoreflect.Value {
	v := new(NestedMessage)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_48_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_49_list)(nil)

type _TestAllTypes_49_list struct {
	list *[]*ForeignMessage
}

func (x *_TestAllTypes_49_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_49_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TestAllTypes_49_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ForeignMessage)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_49_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ForeignMessage)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_49_list) AppendMutable() protoreflect.Value {
	v := new(ForeignMessage)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_49_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_49_list) NewElement() protoreflect.Value {
	v := new(ForeignMessage)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_49_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_50_list)(nil)

type _TestAllTypes_50_list struct {
	list *[]*ImportMessage
}

func (x *_TestAllTypes_50_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_50_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TestAllTypes_50_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ImportMessage)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_50_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ImportMessage)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_50_list) AppendMutable() protoreflect.Value {
	v := new(ImportMessage)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_50_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_50_list) NewElement() protoreflect.Value {
	v := new(ImportMessage)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_50_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_51_list)(nil)

type _TestAllTypes_51_list struct {
	list *[]NestedEnum
}

func (x *_TestAllTypes_51_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_51_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)((*x.list)[i]))
}

func (x *_TestAllTypes_51_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (NestedEnum)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_51_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (NestedEnum)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_51_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedNestedEnum as it is not of Message kind"))
}

func (x *_TestAllTypes_51_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_51_list) NewElement() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_TestAllTypes_51_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_52_list)(nil)

type _TestAllTypes_52_list struct {
	list *[]ForeignEnum
}

func (x *_TestAllTypes_52_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_52_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)((*x.list)[i]))
}

func (x *_TestAllTypes_52_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (ForeignEnum)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_52_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (ForeignEnum)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_52_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedForeignEnum as it is not of Message kind"))
}

func (x *_TestAllTypes_52_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_52_list) NewElement() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_TestAllTypes_52_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TestAllTypes_53_list)(nil)

type _TestAllTypes_53_list struct {
	list *[]ImportEnum
}

func (x *_TestAllTypes_53_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TestAllTypes_53_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)((*x.list)[i]))
}

func (x *_TestAllTypes_53_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (ImportEnum)(valueUnwrapped)
	(*x.list)[i] = concreteValue
}

func (x *_TestAllTypes_53_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Enum()
	concreteValue := (ImportEnum)(valueUnwrapped)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TestAllTypes_53_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message TestAllTypes at list field RepeatedImportenum as it is not of Message kind"))
}

func (x *_TestAllTypes_53_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_TestAllTypes_53_list) NewElement() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_TestAllTypes_53_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.Map = (*_TestAllTypes_56_map)(nil)

type _TestAllTypes_56_map struct {
	m *map[int32]int32
}

func (x *_TestAllTypes_56_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_56_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt32(k))
		mapValue := protoreflect.ValueOfInt32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_56_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := (int32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_56_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_56_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_56_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_56_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_56_map) NewValue() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_56_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_57_map)(nil)

type _TestAllTypes_57_map struct {
	m *map[int64]int64
}

func (x *_TestAllTypes_57_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_57_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt64(k))
		mapValue := protoreflect.ValueOfInt64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_57_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_57_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_57_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_57_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_57_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_57_map) NewValue() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_57_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_58_map)(nil)

type _TestAllTypes_58_map struct {
	m *map[uint32]uint32
}

func (x *_TestAllTypes_58_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_58_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfUint32(k))
		mapValue := protoreflect.ValueOfUint32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_58_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Uint()
	concreteValue := (uint32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_58_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_58_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_58_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_58_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_58_map) NewValue() protoreflect.Value {
	v := uint32(0)
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_58_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_59_map)(nil)

type _TestAllTypes_59_map struct {
	m *map[uint64]uint64
}

func (x *_TestAllTypes_59_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_59_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfUint64(k))
		mapValue := protoreflect.ValueOfUint64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_59_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Uint()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_59_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_59_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_59_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_59_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_59_map) NewValue() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_59_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_60_map)(nil)

type _TestAllTypes_60_map struct {
	m *map[int32]int32
}

func (x *_TestAllTypes_60_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_60_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt32(k))
		mapValue := protoreflect.ValueOfInt32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_60_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := (int32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_60_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_60_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_60_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_60_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_60_map) NewValue() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_60_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_61_map)(nil)

type _TestAllTypes_61_map struct {
	m *map[int64]int64
}

func (x *_TestAllTypes_61_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_61_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt64(k))
		mapValue := protoreflect.ValueOfInt64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_61_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_61_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_61_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_61_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_61_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_61_map) NewValue() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_61_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_62_map)(nil)

type _TestAllTypes_62_map struct {
	m *map[uint32]uint32
}

func (x *_TestAllTypes_62_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_62_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfUint32(k))
		mapValue := protoreflect.ValueOfUint32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_62_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Uint()
	concreteValue := (uint32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_62_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_62_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_62_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Uint()
	concreteKey := (uint32)(keyUnwrapped)
	valueUnwrapped := value.Uint()
	concreteValue := (uint32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_62_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_62_map) NewValue() protoreflect.Value {
	v := uint32(0)
	return protoreflect.ValueOfUint32(v)
}

func (x *_TestAllTypes_62_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_63_map)(nil)

type _TestAllTypes_63_map struct {
	m *map[uint64]uint64
}

func (x *_TestAllTypes_63_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_63_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfUint64(k))
		mapValue := protoreflect.ValueOfUint64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_63_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Uint()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_63_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_63_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_63_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Uint()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Uint()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_63_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_63_map) NewValue() protoreflect.Value {
	v := uint64(0)
	return protoreflect.ValueOfUint64(v)
}

func (x *_TestAllTypes_63_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_64_map)(nil)

type _TestAllTypes_64_map struct {
	m *map[int32]int32
}

func (x *_TestAllTypes_64_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_64_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt32(k))
		mapValue := protoreflect.ValueOfInt32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_64_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := (int32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_64_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_64_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_64_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	valueUnwrapped := value.Int()
	concreteValue := (int32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_64_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_64_map) NewValue() protoreflect.Value {
	v := int32(0)
	return protoreflect.ValueOfInt32(v)
}

func (x *_TestAllTypes_64_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_65_map)(nil)

type _TestAllTypes_65_map struct {
	m *map[int64]int64
}

func (x *_TestAllTypes_65_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_65_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt64(k))
		mapValue := protoreflect.ValueOfInt64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_65_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_65_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_65_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_65_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_65_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_65_map) NewValue() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_TestAllTypes_65_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_66_map)(nil)

type _TestAllTypes_66_map struct {
	m *map[int32]float32
}

func (x *_TestAllTypes_66_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_66_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt32(k))
		mapValue := protoreflect.ValueOfFloat32(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_66_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := (int32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_66_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_66_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfFloat32(v)
}

func (x *_TestAllTypes_66_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	valueUnwrapped := value.Float()
	concreteValue := (float32)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_66_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_66_map) NewValue() protoreflect.Value {
	v := float32(0)
	return protoreflect.ValueOfFloat32(v)
}

func (x *_TestAllTypes_66_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_67_map)(nil)

type _TestAllTypes_67_map struct {
	m *map[int32]float64
}

func (x *_TestAllTypes_67_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_67_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfInt32(k))
		mapValue := protoreflect.ValueOfFloat64(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_67_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Int()
	concreteValue := (int32)(keyUnwrapped)
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_67_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_67_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfFloat64(v)
}

func (x *_TestAllTypes_67_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Int()
	concreteKey := (int32)(keyUnwrapped)
	valueUnwrapped := value.Float()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_67_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_67_map) NewValue() protoreflect.Value {
	v := float64(0)
	return protoreflect.ValueOfFloat64(v)
}

func (x *_TestAllTypes_67_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_68_map)(nil)

type _TestAllTypes_68_map struct {
	m *map[bool]bool
}

func (x *_TestAllTypes_68_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_68_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfBool(k))
		mapValue := protoreflect.ValueOfBool(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_68_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.Bool()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_68_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.Bool()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_68_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.Bool()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfBool(v)
}

func (x *_TestAllTypes_68_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.Bool()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Bool()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_68_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_68_map) NewValue() protoreflect.Value {
	v := false
	return protoreflect.ValueOfBool(v)
}

func (x *_TestAllTypes_68_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_69_map)(nil)

type _TestAllTypes_69_map struct {
	m *map[string]string
}

func (x *_TestAllTypes_69_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_69_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfString(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_69_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_69_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_69_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfString(v)
}

func (x *_TestAllTypes_69_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_69_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_69_map) NewValue() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_TestAllTypes_69_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_70_map)(nil)

type _TestAllTypes_70_map struct {
	m *map[string][]byte
}

func (x *_TestAllTypes_70_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_70_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfBytes(v)
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_70_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_70_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_70_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfBytes(v)
}

func (x *_TestAllTypes_70_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_70_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_70_map) NewValue() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_TestAllTypes_70_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_71_map)(nil)

type _TestAllTypes_71_map struct {
	m *map[string]*NestedMessage
}

func (x *_TestAllTypes_71_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_71_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfMessage(v.ProtoReflect())
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_71_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_71_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_71_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_71_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*NestedMessage)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_71_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if ok {
		return protoreflect.ValueOfMessage(v.ProtoReflect())
	}
	newValue := new(NestedMessage)
	(*x.m)[concreteKey] = newValue
	return protoreflect.ValueOfMessage(newValue.ProtoReflect())
}

func (x *_TestAllTypes_71_map) NewValue() protoreflect.Value {
	v := new(NestedMessage)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TestAllTypes_71_map) IsValid() bool {
	return x.m != nil
}

var _ protoreflect.Map = (*_TestAllTypes_73_map)(nil)

type _TestAllTypes_73_map struct {
	m *map[string]NestedEnum
}

func (x *_TestAllTypes_73_map) Len() int {
	if x.m == nil {
		return 0
	}
	return len(*x.m)
}

func (x *_TestAllTypes_73_map) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	if x.m == nil {
		return
	}
	for k, v := range *x.m {
		mapKey := (protoreflect.MapKey)(protoreflect.ValueOfString(k))
		mapValue := protoreflect.ValueOfEnum(v.Number())
		if !f(mapKey, mapValue) {
			break
		}
	}
}

func (x *_TestAllTypes_73_map) Has(key protoreflect.MapKey) bool {
	if x.m == nil {
		return false
	}
	keyUnwrapped := key.String()
	concreteValue := keyUnwrapped
	_, ok := (*x.m)[concreteValue]
	return ok
}

func (x *_TestAllTypes_73_map) Clear(key protoreflect.MapKey) {
	if x.m == nil {
		return
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	delete(*x.m, concreteKey)
}

func (x *_TestAllTypes_73_map) Get(key protoreflect.MapKey) protoreflect.Value {
	if x.m == nil {
		return protoreflect.Value{}
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	v, ok := (*x.m)[concreteKey]
	if !ok {
		return protoreflect.Value{}
	}
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_TestAllTypes_73_map) Set(key protoreflect.MapKey, value protoreflect.Value) {
	if !key.IsValid() || !value.IsValid() {
		panic("invalid key or value provided")
	}
	keyUnwrapped := key.String()
	concreteKey := keyUnwrapped
	valueUnwrapped := value.Enum()
	concreteValue := (NestedEnum)(valueUnwrapped)
	(*x.m)[concreteKey] = concreteValue
}

func (x *_TestAllTypes_73_map) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("should not call Mutable on protoreflect.Map whose value is not of type protoreflect.Message")
}

func (x *_TestAllTypes_73_map) NewValue() protoreflect.Value {
	v := 0
	return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
}

func (x *_TestAllTypes_73_map) IsValid() bool {
	return x.m != nil
}

var (
	md_TestAllTypes                           protoreflect.MessageDescriptor
	fd_TestAllTypes_singular_int32            protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_int64            protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_uint32           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_uint64           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_sint32           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_sint64           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_fixed32          protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_fixed64          protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_sfixed32         protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_sfixed64         protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_float            protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_double           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_bool             protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_string           protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_bytes            protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_nested_message   protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_foreign_message  protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_import_message   protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_nested_enum      protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_foreign_enum     protoreflect.FieldDescriptor
	fd_TestAllTypes_singular_import_enum      protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_int32            protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_int64            protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_uint32           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_uint64           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_sint32           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_sint64           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_fixed32          protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_fixed64          protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_sfixed32         protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_sfixed64         protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_float            protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_double           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_bool             protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_string           protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_bytes            protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_nested_message   protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_foreign_message  protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_importmessage    protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_nested_enum      protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_foreign_enum     protoreflect.FieldDescriptor
	fd_TestAllTypes_repeated_importenum       protoreflect.FieldDescriptor
	fd_TestAllTypes_map_int32_int32           protoreflect.FieldDescriptor
	fd_TestAllTypes_map_int64_int64           protoreflect.FieldDescriptor
	fd_TestAllTypes_map_uint32_uint32         protoreflect.FieldDescriptor
	fd_TestAllTypes_map_uint64_uint64         protoreflect.FieldDescriptor
	fd_TestAllTypes_map_sint32_sint32         protoreflect.FieldDescriptor
	fd_TestAllTypes_map_sint64_sint64         protoreflect.FieldDescriptor
	fd_TestAllTypes_map_fixed32_fixed32       protoreflect.FieldDescriptor
	fd_TestAllTypes_map_fixed64_fixed64       protoreflect.FieldDescriptor
	fd_TestAllTypes_map_sfixed32_sfixed32     protoreflect.FieldDescriptor
	fd_TestAllTypes_map_sfixed64_sfixed64     protoreflect.FieldDescriptor
	fd_TestAllTypes_map_int32_float           protoreflect.FieldDescriptor
	fd_TestAllTypes_map_int32_double          protoreflect.FieldDescriptor
	fd_TestAllTypes_map_bool_bool             protoreflect.FieldDescriptor
	fd_TestAllTypes_map_string_string         protoreflect.FieldDescriptor
	fd_TestAllTypes_map_string_bytes          protoreflect.FieldDescriptor
	fd_TestAllTypes_map_string_nested_message protoreflect.FieldDescriptor
	fd_TestAllTypes_map_string_nested_enum    protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_uint32              protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_nested_message      protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_string              protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_bytes               protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_bool                protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_uint64              protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_float               protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_double              protoreflect.FieldDescriptor
	fd_TestAllTypes_oneof_enum                protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_proto_init()
	md_TestAllTypes = File_internal_testprotos_test3_test_proto.Messages().ByName("TestAllTypes")
	fd_TestAllTypes_singular_int32 = md_TestAllTypes.Fields().ByName("singular_int32")
	fd_TestAllTypes_singular_int64 = md_TestAllTypes.Fields().ByName("singular_int64")
	fd_TestAllTypes_singular_uint32 = md_TestAllTypes.Fields().ByName("singular_uint32")
	fd_TestAllTypes_singular_uint64 = md_TestAllTypes.Fields().ByName("singular_uint64")
	fd_TestAllTypes_singular_sint32 = md_TestAllTypes.Fields().ByName("singular_sint32")
	fd_TestAllTypes_singular_sint64 = md_TestAllTypes.Fields().ByName("singular_sint64")
	fd_TestAllTypes_singular_fixed32 = md_TestAllTypes.Fields().ByName("singular_fixed32")
	fd_TestAllTypes_singular_fixed64 = md_TestAllTypes.Fields().ByName("singular_fixed64")
	fd_TestAllTypes_singular_sfixed32 = md_TestAllTypes.Fields().ByName("singular_sfixed32")
	fd_TestAllTypes_singular_sfixed64 = md_TestAllTypes.Fields().ByName("singular_sfixed64")
	fd_TestAllTypes_singular_float = md_TestAllTypes.Fields().ByName("singular_float")
	fd_TestAllTypes_singular_double = md_TestAllTypes.Fields().ByName("singular_double")
	fd_TestAllTypes_singular_bool = md_TestAllTypes.Fields().ByName("singular_bool")
	fd_TestAllTypes_singular_string = md_TestAllTypes.Fields().ByName("singular_string")
	fd_TestAllTypes_singular_bytes = md_TestAllTypes.Fields().ByName("singular_bytes")
	fd_TestAllTypes_singular_nested_message = md_TestAllTypes.Fields().ByName("singular_nested_message")
	fd_TestAllTypes_singular_foreign_message = md_TestAllTypes.Fields().ByName("singular_foreign_message")
	fd_TestAllTypes_singular_import_message = md_TestAllTypes.Fields().ByName("singular_import_message")
	fd_TestAllTypes_singular_nested_enum = md_TestAllTypes.Fields().ByName("singular_nested_enum")
	fd_TestAllTypes_singular_foreign_enum = md_TestAllTypes.Fields().ByName("singular_foreign_enum")
	fd_TestAllTypes_singular_import_enum = md_TestAllTypes.Fields().ByName("singular_import_enum")
	fd_TestAllTypes_repeated_int32 = md_TestAllTypes.Fields().ByName("repeated_int32")
	fd_TestAllTypes_repeated_int64 = md_TestAllTypes.Fields().ByName("repeated_int64")
	fd_TestAllTypes_repeated_uint32 = md_TestAllTypes.Fields().ByName("repeated_uint32")
	fd_TestAllTypes_repeated_uint64 = md_TestAllTypes.Fields().ByName("repeated_uint64")
	fd_TestAllTypes_repeated_sint32 = md_TestAllTypes.Fields().ByName("repeated_sint32")
	fd_TestAllTypes_repeated_sint64 = md_TestAllTypes.Fields().ByName("repeated_sint64")
	fd_TestAllTypes_repeated_fixed32 = md_TestAllTypes.Fields().ByName("repeated_fixed32")
	fd_TestAllTypes_repeated_fixed64 = md_TestAllTypes.Fields().ByName("repeated_fixed64")
	fd_TestAllTypes_repeated_sfixed32 = md_TestAllTypes.Fields().ByName("repeated_sfixed32")
	fd_TestAllTypes_repeated_sfixed64 = md_TestAllTypes.Fields().ByName("repeated_sfixed64")
	fd_TestAllTypes_repeated_float = md_TestAllTypes.Fields().ByName("repeated_float")
	fd_TestAllTypes_repeated_double = md_TestAllTypes.Fields().ByName("repeated_double")
	fd_TestAllTypes_repeated_bool = md_TestAllTypes.Fields().ByName("repeated_bool")
	fd_TestAllTypes_repeated_string = md_TestAllTypes.Fields().ByName("repeated_string")
	fd_TestAllTypes_repeated_bytes = md_TestAllTypes.Fields().ByName("repeated_bytes")
	fd_TestAllTypes_repeated_nested_message = md_TestAllTypes.Fields().ByName("repeated_nested_message")
	fd_TestAllTypes_repeated_foreign_message = md_TestAllTypes.Fields().ByName("repeated_foreign_message")
	fd_TestAllTypes_repeated_importmessage = md_TestAllTypes.Fields().ByName("repeated_importmessage")
	fd_TestAllTypes_repeated_nested_enum = md_TestAllTypes.Fields().ByName("repeated_nested_enum")
	fd_TestAllTypes_repeated_foreign_enum = md_TestAllTypes.Fields().ByName("repeated_foreign_enum")
	fd_TestAllTypes_repeated_importenum = md_TestAllTypes.Fields().ByName("repeated_importenum")
	fd_TestAllTypes_map_int32_int32 = md_TestAllTypes.Fields().ByName("map_int32_int32")
	fd_TestAllTypes_map_int64_int64 = md_TestAllTypes.Fields().ByName("map_int64_int64")
	fd_TestAllTypes_map_uint32_uint32 = md_TestAllTypes.Fields().ByName("map_uint32_uint32")
	fd_TestAllTypes_map_uint64_uint64 = md_TestAllTypes.Fields().ByName("map_uint64_uint64")
	fd_TestAllTypes_map_sint32_sint32 = md_TestAllTypes.Fields().ByName("map_sint32_sint32")
	fd_TestAllTypes_map_sint64_sint64 = md_TestAllTypes.Fields().ByName("map_sint64_sint64")
	fd_TestAllTypes_map_fixed32_fixed32 = md_TestAllTypes.Fields().ByName("map_fixed32_fixed32")
	fd_TestAllTypes_map_fixed64_fixed64 = md_TestAllTypes.Fields().ByName("map_fixed64_fixed64")
	fd_TestAllTypes_map_sfixed32_sfixed32 = md_TestAllTypes.Fields().ByName("map_sfixed32_sfixed32")
	fd_TestAllTypes_map_sfixed64_sfixed64 = md_TestAllTypes.Fields().ByName("map_sfixed64_sfixed64")
	fd_TestAllTypes_map_int32_float = md_TestAllTypes.Fields().ByName("map_int32_float")
	fd_TestAllTypes_map_int32_double = md_TestAllTypes.Fields().ByName("map_int32_double")
	fd_TestAllTypes_map_bool_bool = md_TestAllTypes.Fields().ByName("map_bool_bool")
	fd_TestAllTypes_map_string_string = md_TestAllTypes.Fields().ByName("map_string_string")
	fd_TestAllTypes_map_string_bytes = md_TestAllTypes.Fields().ByName("map_string_bytes")
	fd_TestAllTypes_map_string_nested_message = md_TestAllTypes.Fields().ByName("map_string_nested_message")
	fd_TestAllTypes_map_string_nested_enum = md_TestAllTypes.Fields().ByName("map_string_nested_enum")
	fd_TestAllTypes_oneof_uint32 = md_TestAllTypes.Fields().ByName("oneof_uint32")
	fd_TestAllTypes_oneof_nested_message = md_TestAllTypes.Fields().ByName("oneof_nested_message")
	fd_TestAllTypes_oneof_string = md_TestAllTypes.Fields().ByName("oneof_string")
	fd_TestAllTypes_oneof_bytes = md_TestAllTypes.Fields().ByName("oneof_bytes")
	fd_TestAllTypes_oneof_bool = md_TestAllTypes.Fields().ByName("oneof_bool")
	fd_TestAllTypes_oneof_uint64 = md_TestAllTypes.Fields().ByName("oneof_uint64")
	fd_TestAllTypes_oneof_float = md_TestAllTypes.Fields().ByName("oneof_float")
	fd_TestAllTypes_oneof_double = md_TestAllTypes.Fields().ByName("oneof_double")
	fd_TestAllTypes_oneof_enum = md_TestAllTypes.Fields().ByName("oneof_enum")
}

var _ protoreflect.Message = (*fastReflection_TestAllTypes)(nil)

type fastReflection_TestAllTypes TestAllTypes

func (x *TestAllTypes) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TestAllTypes)(x)
}

func (x *TestAllTypes) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TestAllTypes_messageType fastReflection_TestAllTypes_messageType
var _ protoreflect.MessageType = fastReflection_TestAllTypes_messageType{}

type fastReflection_TestAllTypes_messageType struct{}

func (x fastReflection_TestAllTypes_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TestAllTypes)(nil)
}
func (x fastReflection_TestAllTypes_messageType) New() protoreflect.Message {
	return new(fastReflection_TestAllTypes)
}
func (x fastReflection_TestAllTypes_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TestAllTypes
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TestAllTypes) Descriptor() protoreflect.MessageDescriptor {
	return md_TestAllTypes
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TestAllTypes) Type() protoreflect.MessageType {
	return _fastReflection_TestAllTypes_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TestAllTypes) New() protoreflect.Message {
	return new(fastReflection_TestAllTypes)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TestAllTypes) Interface() protoreflect.ProtoMessage {
	return (*TestAllTypes)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TestAllTypes) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.SingularInt32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.SingularInt32)
		if !f(fd_TestAllTypes_singular_int32, value) {
			return
		}
	}
	if x.SingularInt64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.SingularInt64)
		if !f(fd_TestAllTypes_singular_int64, value) {
			return
		}
	}
	if x.SingularUint32 != uint32(0) {
		value := protoreflect.ValueOfUint32(x.SingularUint32)
		if !f(fd_TestAllTypes_singular_uint32, value) {
			return
		}
	}
	if x.SingularUint64 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SingularUint64)
		if !f(fd_TestAllTypes_singular_uint64, value) {
			return
		}
	}
	if x.SingularSint32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.SingularSint32)
		if !f(fd_TestAllTypes_singular_sint32, value) {
			return
		}
	}
	if x.SingularSint64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.SingularSint64)
		if !f(fd_TestAllTypes_singular_sint64, value) {
			return
		}
	}
	if x.SingularFixed32 != uint32(0) {
		value := protoreflect.ValueOfUint32(x.SingularFixed32)
		if !f(fd_TestAllTypes_singular_fixed32, value) {
			return
		}
	}
	if x.SingularFixed64 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SingularFixed64)
		if !f(fd_TestAllTypes_singular_fixed64, value) {
			return
		}
	}
	if x.SingularSfixed32 != int32(0) {
		value := protoreflect.ValueOfInt32(x.SingularSfixed32)
		if !f(fd_TestAllTypes_singular_sfixed32, value) {
			return
		}
	}
	if x.SingularSfixed64 != int64(0) {
		value := protoreflect.ValueOfInt64(x.SingularSfixed64)
		if !f(fd_TestAllTypes_singular_sfixed64, value) {
			return
		}
	}
	if x.SingularFloat != float32(0) || math.Signbit(float64(x.SingularFloat)) {
		value := protoreflect.ValueOfFloat32(x.SingularFloat)
		if !f(fd_TestAllTypes_singular_float, value) {
			return
		}
	}
	if x.SingularDouble != float64(0) || math.Signbit(x.SingularDouble) {
		value := protoreflect.ValueOfFloat64(x.SingularDouble)
		if !f(fd_TestAllTypes_singular_double, value) {
			return
		}
	}
	if x.SingularBool != false {
		value := protoreflect.ValueOfBool(x.SingularBool)
		if !f(fd_TestAllTypes_singular_bool, value) {
			return
		}
	}
	if x.SingularString != "" {
		value := protoreflect.ValueOfString(x.SingularString)
		if !f(fd_TestAllTypes_singular_string, value) {
			return
		}
	}
	if len(x.SingularBytes) != 0 {
		value := protoreflect.ValueOfBytes(x.SingularBytes)
		if !f(fd_TestAllTypes_singular_bytes, value) {
			return
		}
	}
	if x.SingularNestedMessage != nil {
		value := protoreflect.ValueOfMessage(x.SingularNestedMessage.ProtoReflect())
		if !f(fd_TestAllTypes_singular_nested_message, value) {
			return
		}
	}
	if x.SingularForeignMessage != nil {
		value := protoreflect.ValueOfMessage(x.SingularForeignMessage.ProtoReflect())
		if !f(fd_TestAllTypes_singular_foreign_message, value) {
			return
		}
	}
	if x.SingularImportMessage != nil {
		value := protoreflect.ValueOfMessage(x.SingularImportMessage.ProtoReflect())
		if !f(fd_TestAllTypes_singular_import_message, value) {
			return
		}
	}
	if x.SingularNestedEnum != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.SingularNestedEnum))
		if !f(fd_TestAllTypes_singular_nested_enum, value) {
			return
		}
	}
	if x.SingularForeignEnum != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.SingularForeignEnum))
		if !f(fd_TestAllTypes_singular_foreign_enum, value) {
			return
		}
	}
	if x.SingularImportEnum != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.SingularImportEnum))
		if !f(fd_TestAllTypes_singular_import_enum, value) {
			return
		}
	}
	if len(x.RepeatedInt32) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_31_list{list: &x.RepeatedInt32})
		if !f(fd_TestAllTypes_repeated_int32, value) {
			return
		}
	}
	if len(x.RepeatedInt64) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_32_list{list: &x.RepeatedInt64})
		if !f(fd_TestAllTypes_repeated_int64, value) {
			return
		}
	}
	if len(x.RepeatedUint32) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_33_list{list: &x.RepeatedUint32})
		if !f(fd_TestAllTypes_repeated_uint32, value) {
			return
		}
	}
	if len(x.RepeatedUint64) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_34_list{list: &x.RepeatedUint64})
		if !f(fd_TestAllTypes_repeated_uint64, value) {
			return
		}
	}
	if len(x.RepeatedSint32) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_35_list{list: &x.RepeatedSint32})
		if !f(fd_TestAllTypes_repeated_sint32, value) {
			return
		}
	}
	if len(x.RepeatedSint64) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_36_list{list: &x.RepeatedSint64})
		if !f(fd_TestAllTypes_repeated_sint64, value) {
			return
		}
	}
	if len(x.RepeatedFixed32) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_37_list{list: &x.RepeatedFixed32})
		if !f(fd_TestAllTypes_repeated_fixed32, value) {
			return
		}
	}
	if len(x.RepeatedFixed64) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_38_list{list: &x.RepeatedFixed64})
		if !f(fd_TestAllTypes_repeated_fixed64, value) {
			return
		}
	}
	if len(x.RepeatedSfixed32) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_39_list{list: &x.RepeatedSfixed32})
		if !f(fd_TestAllTypes_repeated_sfixed32, value) {
			return
		}
	}
	if len(x.RepeatedSfixed64) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_40_list{list: &x.RepeatedSfixed64})
		if !f(fd_TestAllTypes_repeated_sfixed64, value) {
			return
		}
	}
	if len(x.RepeatedFloat) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_41_list{list: &x.RepeatedFloat})
		if !f(fd_TestAllTypes_repeated_float, value) {
			return
		}
	}
	if len(x.RepeatedDouble) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_42_list{list: &x.RepeatedDouble})
		if !f(fd_TestAllTypes_repeated_double, value) {
			return
		}
	}
	if len(x.RepeatedBool) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_43_list{list: &x.RepeatedBool})
		if !f(fd_TestAllTypes_repeated_bool, value) {
			return
		}
	}
	if len(x.RepeatedString) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_44_list{list: &x.RepeatedString})
		if !f(fd_TestAllTypes_repeated_string, value) {
			return
		}
	}
	if len(x.RepeatedBytes) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_45_list{list: &x.RepeatedBytes})
		if !f(fd_TestAllTypes_repeated_bytes, value) {
			return
		}
	}
	if len(x.RepeatedNestedMessage) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_48_list{list: &x.RepeatedNestedMessage})
		if !f(fd_TestAllTypes_repeated_nested_message, value) {
			return
		}
	}
	if len(x.RepeatedForeignMessage) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_49_list{list: &x.RepeatedForeignMessage})
		if !f(fd_TestAllTypes_repeated_foreign_message, value) {
			return
		}
	}
	if len(x.RepeatedImportmessage) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_50_list{list: &x.RepeatedImportmessage})
		if !f(fd_TestAllTypes_repeated_importmessage, value) {
			return
		}
	}
	if len(x.RepeatedNestedEnum) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_51_list{list: &x.RepeatedNestedEnum})
		if !f(fd_TestAllTypes_repeated_nested_enum, value) {
			return
		}
	}
	if len(x.RepeatedForeignEnum) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_52_list{list: &x.RepeatedForeignEnum})
		if !f(fd_TestAllTypes_repeated_foreign_enum, value) {
			return
		}
	}
	if len(x.RepeatedImportenum) != 0 {
		value := protoreflect.ValueOfList(&_TestAllTypes_53_list{list: &x.RepeatedImportenum})
		if !f(fd_TestAllTypes_repeated_importenum, value) {
			return
		}
	}
	if len(x.MapInt32Int32) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_56_map{m: &x.MapInt32Int32})
		if !f(fd_TestAllTypes_map_int32_int32, value) {
			return
		}
	}
	if len(x.MapInt64Int64) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_57_map{m: &x.MapInt64Int64})
		if !f(fd_TestAllTypes_map_int64_int64, value) {
			return
		}
	}
	if len(x.MapUint32Uint32) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_58_map{m: &x.MapUint32Uint32})
		if !f(fd_TestAllTypes_map_uint32_uint32, value) {
			return
		}
	}
	if len(x.MapUint64Uint64) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_59_map{m: &x.MapUint64Uint64})
		if !f(fd_TestAllTypes_map_uint64_uint64, value) {
			return
		}
	}
	if len(x.MapSint32Sint32) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_60_map{m: &x.MapSint32Sint32})
		if !f(fd_TestAllTypes_map_sint32_sint32, value) {
			return
		}
	}
	if len(x.MapSint64Sint64) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_61_map{m: &x.MapSint64Sint64})
		if !f(fd_TestAllTypes_map_sint64_sint64, value) {
			return
		}
	}
	if len(x.MapFixed32Fixed32) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_62_map{m: &x.MapFixed32Fixed32})
		if !f(fd_TestAllTypes_map_fixed32_fixed32, value) {
			return
		}
	}
	if len(x.MapFixed64Fixed64) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_63_map{m: &x.MapFixed64Fixed64})
		if !f(fd_TestAllTypes_map_fixed64_fixed64, value) {
			return
		}
	}
	if len(x.MapSfixed32Sfixed32) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_64_map{m: &x.MapSfixed32Sfixed32})
		if !f(fd_TestAllTypes_map_sfixed32_sfixed32, value) {
			return
		}
	}
	if len(x.MapSfixed64Sfixed64) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_65_map{m: &x.MapSfixed64Sfixed64})
		if !f(fd_TestAllTypes_map_sfixed64_sfixed64, value) {
			return
		}
	}
	if len(x.MapInt32Float) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_66_map{m: &x.MapInt32Float})
		if !f(fd_TestAllTypes_map_int32_float, value) {
			return
		}
	}
	if len(x.MapInt32Double) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_67_map{m: &x.MapInt32Double})
		if !f(fd_TestAllTypes_map_int32_double, value) {
			return
		}
	}
	if len(x.MapBoolBool) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_68_map{m: &x.MapBoolBool})
		if !f(fd_TestAllTypes_map_bool_bool, value) {
			return
		}
	}
	if len(x.MapStringString) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_69_map{m: &x.MapStringString})
		if !f(fd_TestAllTypes_map_string_string, value) {
			return
		}
	}
	if len(x.MapStringBytes) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_70_map{m: &x.MapStringBytes})
		if !f(fd_TestAllTypes_map_string_bytes, value) {
			return
		}
	}
	if len(x.MapStringNestedMessage) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_71_map{m: &x.MapStringNestedMessage})
		if !f(fd_TestAllTypes_map_string_nested_message, value) {
			return
		}
	}
	if len(x.MapStringNestedEnum) != 0 {
		value := protoreflect.ValueOfMap(&_TestAllTypes_73_map{m: &x.MapStringNestedEnum})
		if !f(fd_TestAllTypes_map_string_nested_enum, value) {
			return
		}
	}
	if x.OneofField != nil {
		switch o := x.OneofField.(type) {
		case *TestAllTypes_OneofUint32:
			v := o.OneofUint32
			value := protoreflect.ValueOfUint32(v)
			if !f(fd_TestAllTypes_oneof_uint32, value) {
				return
			}
		case *TestAllTypes_OneofNestedMessage:
			v := o.OneofNestedMessage
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_TestAllTypes_oneof_nested_message, value) {
				return
			}
		case *TestAllTypes_OneofString:
			v := o.OneofString
			value := protoreflect.ValueOfString(v)
			if !f(fd_TestAllTypes_oneof_string, value) {
				return
			}
		case *TestAllTypes_OneofBytes:
			v := o.OneofBytes
			value := protoreflect.ValueOfBytes(v)
			if !f(fd_TestAllTypes_oneof_bytes, value) {
				return
			}
		case *TestAllTypes_OneofBool:
			v := o.OneofBool
			value := protoreflect.ValueOfBool(v)
			if !f(fd_TestAllTypes_oneof_bool, value) {
				return
			}
		case *TestAllTypes_OneofUint64:
			v := o.OneofUint64
			value := protoreflect.ValueOfUint64(v)
			if !f(fd_TestAllTypes_oneof_uint64, value) {
				return
			}
		case *TestAllTypes_OneofFloat:
			v := o.OneofFloat
			value := protoreflect.ValueOfFloat32(v)
			if !f(fd_TestAllTypes_oneof_float, value) {
				return
			}
		case *TestAllTypes_OneofDouble:
			v := o.OneofDouble
			value := protoreflect.ValueOfFloat64(v)
			if !f(fd_TestAllTypes_oneof_double, value) {
				return
			}
		case *TestAllTypes_OneofEnum:
			v := o.OneofEnum
			value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v))
			if !f(fd_TestAllTypes_oneof_enum, value) {
				return
			}
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
func (x *fastReflection_TestAllTypes) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		return x.SingularInt32 != int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		return x.SingularInt64 != int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		return x.SingularUint32 != uint32(0)
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		return x.SingularUint64 != uint64(0)
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		return x.SingularSint32 != int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		return x.SingularSint64 != int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		return x.SingularFixed32 != uint32(0)
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		return x.SingularFixed64 != uint64(0)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		return x.SingularSfixed32 != int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		return x.SingularSfixed64 != int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_float":
		return x.SingularFloat != float32(0) || math.Signbit(float64(x.SingularFloat))
	case "goproto.proto.test3.TestAllTypes.singular_double":
		return x.SingularDouble != float64(0) || math.Signbit(x.SingularDouble)
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		return x.SingularBool != false
	case "goproto.proto.test3.TestAllTypes.singular_string":
		return x.SingularString != ""
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		return len(x.SingularBytes) != 0
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		return x.SingularNestedMessage != nil
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		return x.SingularForeignMessage != nil
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		return x.SingularImportMessage != nil
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		return x.SingularNestedEnum != 0
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		return x.SingularForeignEnum != 0
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		return x.SingularImportEnum != 0
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		return len(x.RepeatedInt32) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		return len(x.RepeatedInt64) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		return len(x.RepeatedUint32) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		return len(x.RepeatedUint64) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		return len(x.RepeatedSint32) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		return len(x.RepeatedSint64) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		return len(x.RepeatedFixed32) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		return len(x.RepeatedFixed64) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		return len(x.RepeatedSfixed32) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		return len(x.RepeatedSfixed64) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		return len(x.RepeatedFloat) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		return len(x.RepeatedDouble) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		return len(x.RepeatedBool) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		return len(x.RepeatedString) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		return len(x.RepeatedBytes) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		return len(x.RepeatedNestedMessage) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		return len(x.RepeatedForeignMessage) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		return len(x.RepeatedImportmessage) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		return len(x.RepeatedNestedEnum) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		return len(x.RepeatedForeignEnum) != 0
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		return len(x.RepeatedImportenum) != 0
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		return len(x.MapInt32Int32) != 0
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		return len(x.MapInt64Int64) != 0
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		return len(x.MapUint32Uint32) != 0
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		return len(x.MapUint64Uint64) != 0
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		return len(x.MapSint32Sint32) != 0
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		return len(x.MapSint64Sint64) != 0
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		return len(x.MapFixed32Fixed32) != 0
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		return len(x.MapFixed64Fixed64) != 0
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		return len(x.MapSfixed32Sfixed32) != 0
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		return len(x.MapSfixed64Sfixed64) != 0
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		return len(x.MapInt32Float) != 0
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		return len(x.MapInt32Double) != 0
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		return len(x.MapBoolBool) != 0
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		return len(x.MapStringString) != 0
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		return len(x.MapStringBytes) != 0
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		return len(x.MapStringNestedMessage) != 0
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		return len(x.MapStringNestedEnum) != 0
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofUint32); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofNestedMessage); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofString); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofBytes); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofBool); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofUint64); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofFloat); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofDouble); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		if x.OneofField == nil {
			return false
		} else if _, ok := x.OneofField.(*TestAllTypes_OneofEnum); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TestAllTypes) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		x.SingularInt32 = int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		x.SingularInt64 = int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		x.SingularUint32 = uint32(0)
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		x.SingularUint64 = uint64(0)
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		x.SingularSint32 = int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		x.SingularSint64 = int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		x.SingularFixed32 = uint32(0)
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		x.SingularFixed64 = uint64(0)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		x.SingularSfixed32 = int32(0)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		x.SingularSfixed64 = int64(0)
	case "goproto.proto.test3.TestAllTypes.singular_float":
		x.SingularFloat = float32(0)
	case "goproto.proto.test3.TestAllTypes.singular_double":
		x.SingularDouble = float64(0)
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		x.SingularBool = false
	case "goproto.proto.test3.TestAllTypes.singular_string":
		x.SingularString = ""
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		x.SingularBytes = nil
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		x.SingularNestedMessage = nil
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		x.SingularForeignMessage = nil
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		x.SingularImportMessage = nil
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		x.SingularNestedEnum = 0
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		x.SingularForeignEnum = 0
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		x.SingularImportEnum = 0
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		x.RepeatedInt32 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		x.RepeatedInt64 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		x.RepeatedUint32 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		x.RepeatedUint64 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		x.RepeatedSint32 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		x.RepeatedSint64 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		x.RepeatedFixed32 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		x.RepeatedFixed64 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		x.RepeatedSfixed32 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		x.RepeatedSfixed64 = nil
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		x.RepeatedFloat = nil
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		x.RepeatedDouble = nil
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		x.RepeatedBool = nil
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		x.RepeatedString = nil
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		x.RepeatedBytes = nil
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		x.RepeatedNestedMessage = nil
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		x.RepeatedForeignMessage = nil
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		x.RepeatedImportmessage = nil
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		x.RepeatedNestedEnum = nil
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		x.RepeatedForeignEnum = nil
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		x.RepeatedImportenum = nil
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		x.MapInt32Int32 = nil
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		x.MapInt64Int64 = nil
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		x.MapUint32Uint32 = nil
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		x.MapUint64Uint64 = nil
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		x.MapSint32Sint32 = nil
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		x.MapSint64Sint64 = nil
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		x.MapFixed32Fixed32 = nil
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		x.MapFixed64Fixed64 = nil
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		x.MapSfixed32Sfixed32 = nil
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		x.MapSfixed64Sfixed64 = nil
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		x.MapInt32Float = nil
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		x.MapInt32Double = nil
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		x.MapBoolBool = nil
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		x.MapStringString = nil
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		x.MapStringBytes = nil
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		x.MapStringNestedMessage = nil
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		x.MapStringNestedEnum = nil
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		x.OneofField = nil
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		x.OneofField = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TestAllTypes) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		value := x.SingularInt32
		return protoreflect.ValueOfInt32(value)
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		value := x.SingularInt64
		return protoreflect.ValueOfInt64(value)
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		value := x.SingularUint32
		return protoreflect.ValueOfUint32(value)
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		value := x.SingularUint64
		return protoreflect.ValueOfUint64(value)
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		value := x.SingularSint32
		return protoreflect.ValueOfInt32(value)
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		value := x.SingularSint64
		return protoreflect.ValueOfInt64(value)
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		value := x.SingularFixed32
		return protoreflect.ValueOfUint32(value)
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		value := x.SingularFixed64
		return protoreflect.ValueOfUint64(value)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		value := x.SingularSfixed32
		return protoreflect.ValueOfInt32(value)
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		value := x.SingularSfixed64
		return protoreflect.ValueOfInt64(value)
	case "goproto.proto.test3.TestAllTypes.singular_float":
		value := x.SingularFloat
		return protoreflect.ValueOfFloat32(value)
	case "goproto.proto.test3.TestAllTypes.singular_double":
		value := x.SingularDouble
		return protoreflect.ValueOfFloat64(value)
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		value := x.SingularBool
		return protoreflect.ValueOfBool(value)
	case "goproto.proto.test3.TestAllTypes.singular_string":
		value := x.SingularString
		return protoreflect.ValueOfString(value)
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		value := x.SingularBytes
		return protoreflect.ValueOfBytes(value)
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		value := x.SingularNestedMessage
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		value := x.SingularForeignMessage
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		value := x.SingularImportMessage
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		value := x.SingularNestedEnum
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		value := x.SingularForeignEnum
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		value := x.SingularImportEnum
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		if len(x.RepeatedInt32) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_31_list{})
		}
		listValue := &_TestAllTypes_31_list{list: &x.RepeatedInt32}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		if len(x.RepeatedInt64) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_32_list{})
		}
		listValue := &_TestAllTypes_32_list{list: &x.RepeatedInt64}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		if len(x.RepeatedUint32) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_33_list{})
		}
		listValue := &_TestAllTypes_33_list{list: &x.RepeatedUint32}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		if len(x.RepeatedUint64) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_34_list{})
		}
		listValue := &_TestAllTypes_34_list{list: &x.RepeatedUint64}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		if len(x.RepeatedSint32) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_35_list{})
		}
		listValue := &_TestAllTypes_35_list{list: &x.RepeatedSint32}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		if len(x.RepeatedSint64) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_36_list{})
		}
		listValue := &_TestAllTypes_36_list{list: &x.RepeatedSint64}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		if len(x.RepeatedFixed32) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_37_list{})
		}
		listValue := &_TestAllTypes_37_list{list: &x.RepeatedFixed32}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		if len(x.RepeatedFixed64) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_38_list{})
		}
		listValue := &_TestAllTypes_38_list{list: &x.RepeatedFixed64}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		if len(x.RepeatedSfixed32) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_39_list{})
		}
		listValue := &_TestAllTypes_39_list{list: &x.RepeatedSfixed32}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		if len(x.RepeatedSfixed64) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_40_list{})
		}
		listValue := &_TestAllTypes_40_list{list: &x.RepeatedSfixed64}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		if len(x.RepeatedFloat) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_41_list{})
		}
		listValue := &_TestAllTypes_41_list{list: &x.RepeatedFloat}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		if len(x.RepeatedDouble) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_42_list{})
		}
		listValue := &_TestAllTypes_42_list{list: &x.RepeatedDouble}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		if len(x.RepeatedBool) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_43_list{})
		}
		listValue := &_TestAllTypes_43_list{list: &x.RepeatedBool}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		if len(x.RepeatedString) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_44_list{})
		}
		listValue := &_TestAllTypes_44_list{list: &x.RepeatedString}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		if len(x.RepeatedBytes) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_45_list{})
		}
		listValue := &_TestAllTypes_45_list{list: &x.RepeatedBytes}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		if len(x.RepeatedNestedMessage) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_48_list{})
		}
		listValue := &_TestAllTypes_48_list{list: &x.RepeatedNestedMessage}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		if len(x.RepeatedForeignMessage) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_49_list{})
		}
		listValue := &_TestAllTypes_49_list{list: &x.RepeatedForeignMessage}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		if len(x.RepeatedImportmessage) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_50_list{})
		}
		listValue := &_TestAllTypes_50_list{list: &x.RepeatedImportmessage}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		if len(x.RepeatedNestedEnum) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_51_list{})
		}
		listValue := &_TestAllTypes_51_list{list: &x.RepeatedNestedEnum}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		if len(x.RepeatedForeignEnum) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_52_list{})
		}
		listValue := &_TestAllTypes_52_list{list: &x.RepeatedForeignEnum}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		if len(x.RepeatedImportenum) == 0 {
			return protoreflect.ValueOfList(&_TestAllTypes_53_list{})
		}
		listValue := &_TestAllTypes_53_list{list: &x.RepeatedImportenum}
		return protoreflect.ValueOfList(listValue)
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		if len(x.MapInt32Int32) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_56_map{})
		}
		mapValue := &_TestAllTypes_56_map{m: &x.MapInt32Int32}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		if len(x.MapInt64Int64) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_57_map{})
		}
		mapValue := &_TestAllTypes_57_map{m: &x.MapInt64Int64}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		if len(x.MapUint32Uint32) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_58_map{})
		}
		mapValue := &_TestAllTypes_58_map{m: &x.MapUint32Uint32}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		if len(x.MapUint64Uint64) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_59_map{})
		}
		mapValue := &_TestAllTypes_59_map{m: &x.MapUint64Uint64}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		if len(x.MapSint32Sint32) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_60_map{})
		}
		mapValue := &_TestAllTypes_60_map{m: &x.MapSint32Sint32}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		if len(x.MapSint64Sint64) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_61_map{})
		}
		mapValue := &_TestAllTypes_61_map{m: &x.MapSint64Sint64}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		if len(x.MapFixed32Fixed32) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_62_map{})
		}
		mapValue := &_TestAllTypes_62_map{m: &x.MapFixed32Fixed32}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		if len(x.MapFixed64Fixed64) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_63_map{})
		}
		mapValue := &_TestAllTypes_63_map{m: &x.MapFixed64Fixed64}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		if len(x.MapSfixed32Sfixed32) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_64_map{})
		}
		mapValue := &_TestAllTypes_64_map{m: &x.MapSfixed32Sfixed32}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		if len(x.MapSfixed64Sfixed64) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_65_map{})
		}
		mapValue := &_TestAllTypes_65_map{m: &x.MapSfixed64Sfixed64}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		if len(x.MapInt32Float) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_66_map{})
		}
		mapValue := &_TestAllTypes_66_map{m: &x.MapInt32Float}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		if len(x.MapInt32Double) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_67_map{})
		}
		mapValue := &_TestAllTypes_67_map{m: &x.MapInt32Double}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		if len(x.MapBoolBool) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_68_map{})
		}
		mapValue := &_TestAllTypes_68_map{m: &x.MapBoolBool}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		if len(x.MapStringString) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_69_map{})
		}
		mapValue := &_TestAllTypes_69_map{m: &x.MapStringString}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		if len(x.MapStringBytes) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_70_map{})
		}
		mapValue := &_TestAllTypes_70_map{m: &x.MapStringBytes}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		if len(x.MapStringNestedMessage) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_71_map{})
		}
		mapValue := &_TestAllTypes_71_map{m: &x.MapStringNestedMessage}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		if len(x.MapStringNestedEnum) == 0 {
			return protoreflect.ValueOfMap(&_TestAllTypes_73_map{})
		}
		mapValue := &_TestAllTypes_73_map{m: &x.MapStringNestedEnum}
		return protoreflect.ValueOfMap(mapValue)
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		if x.OneofField == nil {
			return protoreflect.ValueOfUint32(uint32(0))
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofUint32); ok {
			return protoreflect.ValueOfUint32(v.OneofUint32)
		} else {
			return protoreflect.ValueOfUint32(uint32(0))
		}
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		if x.OneofField == nil {
			return protoreflect.ValueOfMessage((*NestedMessage)(nil).ProtoReflect())
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofNestedMessage); ok {
			return protoreflect.ValueOfMessage(v.OneofNestedMessage.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*NestedMessage)(nil).ProtoReflect())
		}
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		if x.OneofField == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofString); ok {
			return protoreflect.ValueOfString(v.OneofString)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		if x.OneofField == nil {
			return protoreflect.ValueOfBytes(nil)
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofBytes); ok {
			return protoreflect.ValueOfBytes(v.OneofBytes)
		} else {
			return protoreflect.ValueOfBytes(nil)
		}
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		if x.OneofField == nil {
			return protoreflect.ValueOfBool(false)
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofBool); ok {
			return protoreflect.ValueOfBool(v.OneofBool)
		} else {
			return protoreflect.ValueOfBool(false)
		}
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		if x.OneofField == nil {
			return protoreflect.ValueOfUint64(uint64(0))
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofUint64); ok {
			return protoreflect.ValueOfUint64(v.OneofUint64)
		} else {
			return protoreflect.ValueOfUint64(uint64(0))
		}
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		if x.OneofField == nil {
			return protoreflect.ValueOfFloat32(float32(0))
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofFloat); ok {
			return protoreflect.ValueOfFloat32(v.OneofFloat)
		} else {
			return protoreflect.ValueOfFloat32(float32(0))
		}
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		if x.OneofField == nil {
			return protoreflect.ValueOfFloat64(float64(0))
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofDouble); ok {
			return protoreflect.ValueOfFloat64(v.OneofDouble)
		} else {
			return protoreflect.ValueOfFloat64(float64(0))
		}
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		if x.OneofField == nil {
			return protoreflect.ValueOfEnum(0)
		} else if v, ok := x.OneofField.(*TestAllTypes_OneofEnum); ok {
			return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(v.OneofEnum))
		} else {
			return protoreflect.ValueOfEnum(0)
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TestAllTypes) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		x.SingularInt32 = int32(value.Int())
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		x.SingularInt64 = value.Int()
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		x.SingularUint32 = uint32(value.Uint())
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		x.SingularUint64 = value.Uint()
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		x.SingularSint32 = int32(value.Int())
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		x.SingularSint64 = value.Int()
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		x.SingularFixed32 = uint32(value.Uint())
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		x.SingularFixed64 = value.Uint()
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		x.SingularSfixed32 = int32(value.Int())
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		x.SingularSfixed64 = value.Int()
	case "goproto.proto.test3.TestAllTypes.singular_float":
		x.SingularFloat = float32(value.Float())
	case "goproto.proto.test3.TestAllTypes.singular_double":
		x.SingularDouble = value.Float()
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		x.SingularBool = value.Bool()
	case "goproto.proto.test3.TestAllTypes.singular_string":
		x.SingularString = value.Interface().(string)
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		x.SingularBytes = value.Bytes()
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		x.SingularNestedMessage = value.Message().Interface().(*NestedMessage)
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		x.SingularForeignMessage = value.Message().Interface().(*ForeignMessage)
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		x.SingularImportMessage = value.Message().Interface().(*ImportMessage)
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		x.SingularNestedEnum = (NestedEnum)(value.Enum())
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		x.SingularForeignEnum = (ForeignEnum)(value.Enum())
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		x.SingularImportEnum = (ImportEnum)(value.Enum())
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		lv := value.List()
		clv := lv.(*_TestAllTypes_31_list)
		x.RepeatedInt32 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		lv := value.List()
		clv := lv.(*_TestAllTypes_32_list)
		x.RepeatedInt64 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		lv := value.List()
		clv := lv.(*_TestAllTypes_33_list)
		x.RepeatedUint32 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		lv := value.List()
		clv := lv.(*_TestAllTypes_34_list)
		x.RepeatedUint64 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		lv := value.List()
		clv := lv.(*_TestAllTypes_35_list)
		x.RepeatedSint32 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		lv := value.List()
		clv := lv.(*_TestAllTypes_36_list)
		x.RepeatedSint64 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		lv := value.List()
		clv := lv.(*_TestAllTypes_37_list)
		x.RepeatedFixed32 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		lv := value.List()
		clv := lv.(*_TestAllTypes_38_list)
		x.RepeatedFixed64 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		lv := value.List()
		clv := lv.(*_TestAllTypes_39_list)
		x.RepeatedSfixed32 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		lv := value.List()
		clv := lv.(*_TestAllTypes_40_list)
		x.RepeatedSfixed64 = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		lv := value.List()
		clv := lv.(*_TestAllTypes_41_list)
		x.RepeatedFloat = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		lv := value.List()
		clv := lv.(*_TestAllTypes_42_list)
		x.RepeatedDouble = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		lv := value.List()
		clv := lv.(*_TestAllTypes_43_list)
		x.RepeatedBool = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		lv := value.List()
		clv := lv.(*_TestAllTypes_44_list)
		x.RepeatedString = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		lv := value.List()
		clv := lv.(*_TestAllTypes_45_list)
		x.RepeatedBytes = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		lv := value.List()
		clv := lv.(*_TestAllTypes_48_list)
		x.RepeatedNestedMessage = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		lv := value.List()
		clv := lv.(*_TestAllTypes_49_list)
		x.RepeatedForeignMessage = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		lv := value.List()
		clv := lv.(*_TestAllTypes_50_list)
		x.RepeatedImportmessage = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		lv := value.List()
		clv := lv.(*_TestAllTypes_51_list)
		x.RepeatedNestedEnum = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		lv := value.List()
		clv := lv.(*_TestAllTypes_52_list)
		x.RepeatedForeignEnum = *clv.list
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		lv := value.List()
		clv := lv.(*_TestAllTypes_53_list)
		x.RepeatedImportenum = *clv.list
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_56_map)
		x.MapInt32Int32 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_57_map)
		x.MapInt64Int64 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_58_map)
		x.MapUint32Uint32 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_59_map)
		x.MapUint64Uint64 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_60_map)
		x.MapSint32Sint32 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_61_map)
		x.MapSint64Sint64 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_62_map)
		x.MapFixed32Fixed32 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_63_map)
		x.MapFixed64Fixed64 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_64_map)
		x.MapSfixed32Sfixed32 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_65_map)
		x.MapSfixed64Sfixed64 = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_66_map)
		x.MapInt32Float = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_67_map)
		x.MapInt32Double = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_68_map)
		x.MapBoolBool = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_69_map)
		x.MapStringString = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_70_map)
		x.MapStringBytes = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_71_map)
		x.MapStringNestedMessage = *cmv.m
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		mv := value.Map()
		cmv := mv.(*_TestAllTypes_73_map)
		x.MapStringNestedEnum = *cmv.m
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		cv := uint32(value.Uint())
		x.OneofField = &TestAllTypes_OneofUint32{OneofUint32: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		cv := value.Message().Interface().(*NestedMessage)
		x.OneofField = &TestAllTypes_OneofNestedMessage{OneofNestedMessage: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		cv := value.Interface().(string)
		x.OneofField = &TestAllTypes_OneofString{OneofString: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		cv := value.Bytes()
		x.OneofField = &TestAllTypes_OneofBytes{OneofBytes: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		cv := value.Bool()
		x.OneofField = &TestAllTypes_OneofBool{OneofBool: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		cv := value.Uint()
		x.OneofField = &TestAllTypes_OneofUint64{OneofUint64: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		cv := float32(value.Float())
		x.OneofField = &TestAllTypes_OneofFloat{OneofFloat: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		cv := value.Float()
		x.OneofField = &TestAllTypes_OneofDouble{OneofDouble: cv}
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		cv := (NestedEnum)(value.Enum())
		x.OneofField = &TestAllTypes_OneofEnum{OneofEnum: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TestAllTypes) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		if x.SingularNestedMessage == nil {
			x.SingularNestedMessage = new(NestedMessage)
		}
		return protoreflect.ValueOfMessage(x.SingularNestedMessage.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		if x.SingularForeignMessage == nil {
			x.SingularForeignMessage = new(ForeignMessage)
		}
		return protoreflect.ValueOfMessage(x.SingularForeignMessage.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		if x.SingularImportMessage == nil {
			x.SingularImportMessage = new(ImportMessage)
		}
		return protoreflect.ValueOfMessage(x.SingularImportMessage.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		if x.RepeatedInt32 == nil {
			x.RepeatedInt32 = []int32{}
		}
		value := &_TestAllTypes_31_list{list: &x.RepeatedInt32}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		if x.RepeatedInt64 == nil {
			x.RepeatedInt64 = []int64{}
		}
		value := &_TestAllTypes_32_list{list: &x.RepeatedInt64}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		if x.RepeatedUint32 == nil {
			x.RepeatedUint32 = []uint32{}
		}
		value := &_TestAllTypes_33_list{list: &x.RepeatedUint32}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		if x.RepeatedUint64 == nil {
			x.RepeatedUint64 = []uint64{}
		}
		value := &_TestAllTypes_34_list{list: &x.RepeatedUint64}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		if x.RepeatedSint32 == nil {
			x.RepeatedSint32 = []int32{}
		}
		value := &_TestAllTypes_35_list{list: &x.RepeatedSint32}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		if x.RepeatedSint64 == nil {
			x.RepeatedSint64 = []int64{}
		}
		value := &_TestAllTypes_36_list{list: &x.RepeatedSint64}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		if x.RepeatedFixed32 == nil {
			x.RepeatedFixed32 = []uint32{}
		}
		value := &_TestAllTypes_37_list{list: &x.RepeatedFixed32}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		if x.RepeatedFixed64 == nil {
			x.RepeatedFixed64 = []uint64{}
		}
		value := &_TestAllTypes_38_list{list: &x.RepeatedFixed64}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		if x.RepeatedSfixed32 == nil {
			x.RepeatedSfixed32 = []int32{}
		}
		value := &_TestAllTypes_39_list{list: &x.RepeatedSfixed32}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		if x.RepeatedSfixed64 == nil {
			x.RepeatedSfixed64 = []int64{}
		}
		value := &_TestAllTypes_40_list{list: &x.RepeatedSfixed64}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		if x.RepeatedFloat == nil {
			x.RepeatedFloat = []float32{}
		}
		value := &_TestAllTypes_41_list{list: &x.RepeatedFloat}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		if x.RepeatedDouble == nil {
			x.RepeatedDouble = []float64{}
		}
		value := &_TestAllTypes_42_list{list: &x.RepeatedDouble}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		if x.RepeatedBool == nil {
			x.RepeatedBool = []bool{}
		}
		value := &_TestAllTypes_43_list{list: &x.RepeatedBool}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		if x.RepeatedString == nil {
			x.RepeatedString = []string{}
		}
		value := &_TestAllTypes_44_list{list: &x.RepeatedString}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		if x.RepeatedBytes == nil {
			x.RepeatedBytes = [][]byte{}
		}
		value := &_TestAllTypes_45_list{list: &x.RepeatedBytes}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		if x.RepeatedNestedMessage == nil {
			x.RepeatedNestedMessage = []*NestedMessage{}
		}
		value := &_TestAllTypes_48_list{list: &x.RepeatedNestedMessage}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		if x.RepeatedForeignMessage == nil {
			x.RepeatedForeignMessage = []*ForeignMessage{}
		}
		value := &_TestAllTypes_49_list{list: &x.RepeatedForeignMessage}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		if x.RepeatedImportmessage == nil {
			x.RepeatedImportmessage = []*ImportMessage{}
		}
		value := &_TestAllTypes_50_list{list: &x.RepeatedImportmessage}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		if x.RepeatedNestedEnum == nil {
			x.RepeatedNestedEnum = []NestedEnum{}
		}
		value := &_TestAllTypes_51_list{list: &x.RepeatedNestedEnum}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		if x.RepeatedForeignEnum == nil {
			x.RepeatedForeignEnum = []ForeignEnum{}
		}
		value := &_TestAllTypes_52_list{list: &x.RepeatedForeignEnum}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		if x.RepeatedImportenum == nil {
			x.RepeatedImportenum = []ImportEnum{}
		}
		value := &_TestAllTypes_53_list{list: &x.RepeatedImportenum}
		return protoreflect.ValueOfList(value)
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		if x.MapInt32Int32 == nil {
			x.MapInt32Int32 = make(map[int32]int32)
		}
		value := &_TestAllTypes_56_map{m: &x.MapInt32Int32}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		if x.MapInt64Int64 == nil {
			x.MapInt64Int64 = make(map[int64]int64)
		}
		value := &_TestAllTypes_57_map{m: &x.MapInt64Int64}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		if x.MapUint32Uint32 == nil {
			x.MapUint32Uint32 = make(map[uint32]uint32)
		}
		value := &_TestAllTypes_58_map{m: &x.MapUint32Uint32}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		if x.MapUint64Uint64 == nil {
			x.MapUint64Uint64 = make(map[uint64]uint64)
		}
		value := &_TestAllTypes_59_map{m: &x.MapUint64Uint64}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		if x.MapSint32Sint32 == nil {
			x.MapSint32Sint32 = make(map[int32]int32)
		}
		value := &_TestAllTypes_60_map{m: &x.MapSint32Sint32}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		if x.MapSint64Sint64 == nil {
			x.MapSint64Sint64 = make(map[int64]int64)
		}
		value := &_TestAllTypes_61_map{m: &x.MapSint64Sint64}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		if x.MapFixed32Fixed32 == nil {
			x.MapFixed32Fixed32 = make(map[uint32]uint32)
		}
		value := &_TestAllTypes_62_map{m: &x.MapFixed32Fixed32}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		if x.MapFixed64Fixed64 == nil {
			x.MapFixed64Fixed64 = make(map[uint64]uint64)
		}
		value := &_TestAllTypes_63_map{m: &x.MapFixed64Fixed64}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		if x.MapSfixed32Sfixed32 == nil {
			x.MapSfixed32Sfixed32 = make(map[int32]int32)
		}
		value := &_TestAllTypes_64_map{m: &x.MapSfixed32Sfixed32}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		if x.MapSfixed64Sfixed64 == nil {
			x.MapSfixed64Sfixed64 = make(map[int64]int64)
		}
		value := &_TestAllTypes_65_map{m: &x.MapSfixed64Sfixed64}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		if x.MapInt32Float == nil {
			x.MapInt32Float = make(map[int32]float32)
		}
		value := &_TestAllTypes_66_map{m: &x.MapInt32Float}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		if x.MapInt32Double == nil {
			x.MapInt32Double = make(map[int32]float64)
		}
		value := &_TestAllTypes_67_map{m: &x.MapInt32Double}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		if x.MapBoolBool == nil {
			x.MapBoolBool = make(map[bool]bool)
		}
		value := &_TestAllTypes_68_map{m: &x.MapBoolBool}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		if x.MapStringString == nil {
			x.MapStringString = make(map[string]string)
		}
		value := &_TestAllTypes_69_map{m: &x.MapStringString}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		if x.MapStringBytes == nil {
			x.MapStringBytes = make(map[string][]byte)
		}
		value := &_TestAllTypes_70_map{m: &x.MapStringBytes}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		if x.MapStringNestedMessage == nil {
			x.MapStringNestedMessage = make(map[string]*NestedMessage)
		}
		value := &_TestAllTypes_71_map{m: &x.MapStringNestedMessage}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		if x.MapStringNestedEnum == nil {
			x.MapStringNestedEnum = make(map[string]NestedEnum)
		}
		value := &_TestAllTypes_73_map{m: &x.MapStringNestedEnum}
		return protoreflect.ValueOfMap(value)
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		if x.OneofField == nil {
			value := &NestedMessage{}
			oneofValue := &TestAllTypes_OneofNestedMessage{OneofNestedMessage: value}
			x.OneofField = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.OneofField.(type) {
		case *TestAllTypes_OneofNestedMessage:
			return protoreflect.ValueOfMessage(m.OneofNestedMessage.ProtoReflect())
		default:
			value := &NestedMessage{}
			oneofValue := &TestAllTypes_OneofNestedMessage{OneofNestedMessage: value}
			x.OneofField = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		panic(fmt.Errorf("field singular_int32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		panic(fmt.Errorf("field singular_int64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		panic(fmt.Errorf("field singular_uint32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		panic(fmt.Errorf("field singular_uint64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		panic(fmt.Errorf("field singular_sint32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		panic(fmt.Errorf("field singular_sint64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		panic(fmt.Errorf("field singular_fixed32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		panic(fmt.Errorf("field singular_fixed64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		panic(fmt.Errorf("field singular_sfixed32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		panic(fmt.Errorf("field singular_sfixed64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_float":
		panic(fmt.Errorf("field singular_float of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_double":
		panic(fmt.Errorf("field singular_double of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		panic(fmt.Errorf("field singular_bool of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_string":
		panic(fmt.Errorf("field singular_string of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		panic(fmt.Errorf("field singular_bytes of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		panic(fmt.Errorf("field singular_nested_enum of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		panic(fmt.Errorf("field singular_foreign_enum of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		panic(fmt.Errorf("field singular_import_enum of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		panic(fmt.Errorf("field oneof_uint32 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		panic(fmt.Errorf("field oneof_string of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		panic(fmt.Errorf("field oneof_bytes of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		panic(fmt.Errorf("field oneof_bool of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		panic(fmt.Errorf("field oneof_uint64 of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		panic(fmt.Errorf("field oneof_float of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		panic(fmt.Errorf("field oneof_double of message goproto.proto.test3.TestAllTypes is not mutable"))
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		panic(fmt.Errorf("field oneof_enum of message goproto.proto.test3.TestAllTypes is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TestAllTypes) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.TestAllTypes.singular_int32":
		return protoreflect.ValueOfInt32(int32(0))
	case "goproto.proto.test3.TestAllTypes.singular_int64":
		return protoreflect.ValueOfInt64(int64(0))
	case "goproto.proto.test3.TestAllTypes.singular_uint32":
		return protoreflect.ValueOfUint32(uint32(0))
	case "goproto.proto.test3.TestAllTypes.singular_uint64":
		return protoreflect.ValueOfUint64(uint64(0))
	case "goproto.proto.test3.TestAllTypes.singular_sint32":
		return protoreflect.ValueOfInt32(int32(0))
	case "goproto.proto.test3.TestAllTypes.singular_sint64":
		return protoreflect.ValueOfInt64(int64(0))
	case "goproto.proto.test3.TestAllTypes.singular_fixed32":
		return protoreflect.ValueOfUint32(uint32(0))
	case "goproto.proto.test3.TestAllTypes.singular_fixed64":
		return protoreflect.ValueOfUint64(uint64(0))
	case "goproto.proto.test3.TestAllTypes.singular_sfixed32":
		return protoreflect.ValueOfInt32(int32(0))
	case "goproto.proto.test3.TestAllTypes.singular_sfixed64":
		return protoreflect.ValueOfInt64(int64(0))
	case "goproto.proto.test3.TestAllTypes.singular_float":
		return protoreflect.ValueOfFloat32(float32(0))
	case "goproto.proto.test3.TestAllTypes.singular_double":
		return protoreflect.ValueOfFloat64(float64(0))
	case "goproto.proto.test3.TestAllTypes.singular_bool":
		return protoreflect.ValueOfBool(false)
	case "goproto.proto.test3.TestAllTypes.singular_string":
		return protoreflect.ValueOfString("")
	case "goproto.proto.test3.TestAllTypes.singular_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "goproto.proto.test3.TestAllTypes.singular_nested_message":
		m := new(NestedMessage)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_foreign_message":
		m := new(ForeignMessage)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_import_message":
		m := new(ImportMessage)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.singular_nested_enum":
		return protoreflect.ValueOfEnum(0)
	case "goproto.proto.test3.TestAllTypes.singular_foreign_enum":
		return protoreflect.ValueOfEnum(0)
	case "goproto.proto.test3.TestAllTypes.singular_import_enum":
		return protoreflect.ValueOfEnum(0)
	case "goproto.proto.test3.TestAllTypes.repeated_int32":
		list := []int32{}
		return protoreflect.ValueOfList(&_TestAllTypes_31_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_int64":
		list := []int64{}
		return protoreflect.ValueOfList(&_TestAllTypes_32_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_uint32":
		list := []uint32{}
		return protoreflect.ValueOfList(&_TestAllTypes_33_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_uint64":
		list := []uint64{}
		return protoreflect.ValueOfList(&_TestAllTypes_34_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_sint32":
		list := []int32{}
		return protoreflect.ValueOfList(&_TestAllTypes_35_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_sint64":
		list := []int64{}
		return protoreflect.ValueOfList(&_TestAllTypes_36_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_fixed32":
		list := []uint32{}
		return protoreflect.ValueOfList(&_TestAllTypes_37_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_fixed64":
		list := []uint64{}
		return protoreflect.ValueOfList(&_TestAllTypes_38_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed32":
		list := []int32{}
		return protoreflect.ValueOfList(&_TestAllTypes_39_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_sfixed64":
		list := []int64{}
		return protoreflect.ValueOfList(&_TestAllTypes_40_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_float":
		list := []float32{}
		return protoreflect.ValueOfList(&_TestAllTypes_41_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_double":
		list := []float64{}
		return protoreflect.ValueOfList(&_TestAllTypes_42_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_bool":
		list := []bool{}
		return protoreflect.ValueOfList(&_TestAllTypes_43_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_string":
		list := []string{}
		return protoreflect.ValueOfList(&_TestAllTypes_44_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_bytes":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_TestAllTypes_45_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_nested_message":
		list := []*NestedMessage{}
		return protoreflect.ValueOfList(&_TestAllTypes_48_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_message":
		list := []*ForeignMessage{}
		return protoreflect.ValueOfList(&_TestAllTypes_49_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_importmessage":
		list := []*ImportMessage{}
		return protoreflect.ValueOfList(&_TestAllTypes_50_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_nested_enum":
		list := []NestedEnum{}
		return protoreflect.ValueOfList(&_TestAllTypes_51_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_foreign_enum":
		list := []ForeignEnum{}
		return protoreflect.ValueOfList(&_TestAllTypes_52_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.repeated_importenum":
		list := []ImportEnum{}
		return protoreflect.ValueOfList(&_TestAllTypes_53_list{list: &list})
	case "goproto.proto.test3.TestAllTypes.map_int32_int32":
		m := make(map[int32]int32)
		return protoreflect.ValueOfMap(&_TestAllTypes_56_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_int64_int64":
		m := make(map[int64]int64)
		return protoreflect.ValueOfMap(&_TestAllTypes_57_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_uint32_uint32":
		m := make(map[uint32]uint32)
		return protoreflect.ValueOfMap(&_TestAllTypes_58_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_uint64_uint64":
		m := make(map[uint64]uint64)
		return protoreflect.ValueOfMap(&_TestAllTypes_59_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_sint32_sint32":
		m := make(map[int32]int32)
		return protoreflect.ValueOfMap(&_TestAllTypes_60_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_sint64_sint64":
		m := make(map[int64]int64)
		return protoreflect.ValueOfMap(&_TestAllTypes_61_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_fixed32_fixed32":
		m := make(map[uint32]uint32)
		return protoreflect.ValueOfMap(&_TestAllTypes_62_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_fixed64_fixed64":
		m := make(map[uint64]uint64)
		return protoreflect.ValueOfMap(&_TestAllTypes_63_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32":
		m := make(map[int32]int32)
		return protoreflect.ValueOfMap(&_TestAllTypes_64_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64":
		m := make(map[int64]int64)
		return protoreflect.ValueOfMap(&_TestAllTypes_65_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_int32_float":
		m := make(map[int32]float32)
		return protoreflect.ValueOfMap(&_TestAllTypes_66_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_int32_double":
		m := make(map[int32]float64)
		return protoreflect.ValueOfMap(&_TestAllTypes_67_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_bool_bool":
		m := make(map[bool]bool)
		return protoreflect.ValueOfMap(&_TestAllTypes_68_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_string_string":
		m := make(map[string]string)
		return protoreflect.ValueOfMap(&_TestAllTypes_69_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_string_bytes":
		m := make(map[string][]byte)
		return protoreflect.ValueOfMap(&_TestAllTypes_70_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_string_nested_message":
		m := make(map[string]*NestedMessage)
		return protoreflect.ValueOfMap(&_TestAllTypes_71_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.map_string_nested_enum":
		m := make(map[string]NestedEnum)
		return protoreflect.ValueOfMap(&_TestAllTypes_73_map{m: &m})
	case "goproto.proto.test3.TestAllTypes.oneof_uint32":
		return protoreflect.ValueOfUint32(uint32(0))
	case "goproto.proto.test3.TestAllTypes.oneof_nested_message":
		value := &NestedMessage{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "goproto.proto.test3.TestAllTypes.oneof_string":
		return protoreflect.ValueOfString("")
	case "goproto.proto.test3.TestAllTypes.oneof_bytes":
		return protoreflect.ValueOfBytes(nil)
	case "goproto.proto.test3.TestAllTypes.oneof_bool":
		return protoreflect.ValueOfBool(false)
	case "goproto.proto.test3.TestAllTypes.oneof_uint64":
		return protoreflect.ValueOfUint64(uint64(0))
	case "goproto.proto.test3.TestAllTypes.oneof_float":
		return protoreflect.ValueOfFloat32(float32(0))
	case "goproto.proto.test3.TestAllTypes.oneof_double":
		return protoreflect.ValueOfFloat64(float64(0))
	case "goproto.proto.test3.TestAllTypes.oneof_enum":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.TestAllTypes"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.TestAllTypes does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TestAllTypes) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "goproto.proto.test3.TestAllTypes.oneof_field":
		if x.OneofField == nil {
			return nil
		}
		switch x.OneofField.(type) {
		case *TestAllTypes_OneofUint32:
			return x.Descriptor().Fields().ByName("oneof_uint32")
		case *TestAllTypes_OneofNestedMessage:
			return x.Descriptor().Fields().ByName("oneof_nested_message")
		case *TestAllTypes_OneofString:
			return x.Descriptor().Fields().ByName("oneof_string")
		case *TestAllTypes_OneofBytes:
			return x.Descriptor().Fields().ByName("oneof_bytes")
		case *TestAllTypes_OneofBool:
			return x.Descriptor().Fields().ByName("oneof_bool")
		case *TestAllTypes_OneofUint64:
			return x.Descriptor().Fields().ByName("oneof_uint64")
		case *TestAllTypes_OneofFloat:
			return x.Descriptor().Fields().ByName("oneof_float")
		case *TestAllTypes_OneofDouble:
			return x.Descriptor().Fields().ByName("oneof_double")
		case *TestAllTypes_OneofEnum:
			return x.Descriptor().Fields().ByName("oneof_enum")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.TestAllTypes", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TestAllTypes) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TestAllTypes) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TestAllTypes) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TestAllTypes) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TestAllTypes)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.SingularInt32 != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularInt32))
		}
		if x.SingularInt64 != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularInt64))
		}
		if x.SingularUint32 != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularUint32))
		}
		if x.SingularUint64 != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularUint64))
		}
		if x.SingularSint32 != 0 {
			n += 2 + runtime.Soz(uint64(x.SingularSint32))
		}
		if x.SingularSint64 != 0 {
			n += 2 + runtime.Soz(uint64(x.SingularSint64))
		}
		if x.SingularFixed32 != 0 {
			n += 6
		}
		if x.SingularFixed64 != 0 {
			n += 10
		}
		if x.SingularSfixed32 != 0 {
			n += 6
		}
		if x.SingularSfixed64 != 0 {
			n += 10
		}
		if x.SingularFloat != 0 || math.Signbit(float64(x.SingularFloat)) {
			n += 6
		}
		if x.SingularDouble != 0 || math.Signbit(x.SingularDouble) {
			n += 10
		}
		if x.SingularBool {
			n += 3
		}
		l = len(x.SingularString)
		if l > 0 {
			n += 2 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SingularBytes)
		if l > 0 {
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.SingularNestedMessage != nil {
			l = options.Size(x.SingularNestedMessage)
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.SingularForeignMessage != nil {
			l = options.Size(x.SingularForeignMessage)
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.SingularImportMessage != nil {
			l = options.Size(x.SingularImportMessage)
			n += 2 + l + runtime.Sov(uint64(l))
		}
		if x.SingularNestedEnum != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularNestedEnum))
		}
		if x.SingularForeignEnum != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularForeignEnum))
		}
		if x.SingularImportEnum != 0 {
			n += 2 + runtime.Sov(uint64(x.SingularImportEnum))
		}
		if len(x.RepeatedInt32) > 0 {
			l = 0
			for _, e := range x.RepeatedInt32 {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedInt64) > 0 {
			l = 0
			for _, e := range x.RepeatedInt64 {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedUint32) > 0 {
			l = 0
			for _, e := range x.RepeatedUint32 {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedUint64) > 0 {
			l = 0
			for _, e := range x.RepeatedUint64 {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedSint32) > 0 {
			l = 0
			for _, e := range x.RepeatedSint32 {
				l += runtime.Soz(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedSint64) > 0 {
			l = 0
			for _, e := range x.RepeatedSint64 {
				l += runtime.Soz(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedFixed32) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedFixed32)*4)) + len(x.RepeatedFixed32)*4
		}
		if len(x.RepeatedFixed64) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedFixed64)*8)) + len(x.RepeatedFixed64)*8
		}
		if len(x.RepeatedSfixed32) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedSfixed32)*4)) + len(x.RepeatedSfixed32)*4
		}
		if len(x.RepeatedSfixed64) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedSfixed64)*8)) + len(x.RepeatedSfixed64)*8
		}
		if len(x.RepeatedFloat) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedFloat)*4)) + len(x.RepeatedFloat)*4
		}
		if len(x.RepeatedDouble) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedDouble)*8)) + len(x.RepeatedDouble)*8
		}
		if len(x.RepeatedBool) > 0 {
			n += 2 + runtime.Sov(uint64(len(x.RepeatedBool))) + len(x.RepeatedBool)*1
		}
		if len(x.RepeatedString) > 0 {
			for _, s := range x.RepeatedString {
				l = len(s)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.RepeatedBytes) > 0 {
			for _, b := range x.RepeatedBytes {
				l = len(b)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.RepeatedNestedMessage) > 0 {
			for _, e := range x.RepeatedNestedMessage {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.RepeatedForeignMessage) > 0 {
			for _, e := range x.RepeatedForeignMessage {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.RepeatedImportmessage) > 0 {
			for _, e := range x.RepeatedImportmessage {
				l = options.Size(e)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.RepeatedNestedEnum) > 0 {
			l = 0
			for _, e := range x.RepeatedNestedEnum {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedForeignEnum) > 0 {
			l = 0
			for _, e := range x.RepeatedForeignEnum {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.RepeatedImportenum) > 0 {
			l = 0
			for _, e := range x.RepeatedImportenum {
				l += runtime.Sov(uint64(e))
			}
			n += 2 + runtime.Sov(uint64(l)) + l
		}
		if len(x.MapInt32Int32) > 0 {
			SiZeMaP := func(k int32, v int32) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int32, 0, len(x.MapInt32Int32))
				for k := range x.MapInt32Int32 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapInt32Int32[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapInt32Int32 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapInt64Int64) > 0 {
			SiZeMaP := func(k int64, v int64) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int64, 0, len(x.MapInt64Int64))
				for k := range x.MapInt64Int64 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapInt64Int64[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapInt64Int64 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapUint32Uint32) > 0 {
			SiZeMaP := func(k uint32, v uint32) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]uint32, 0, len(x.MapUint32Uint32))
				for k := range x.MapUint32Uint32 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapUint32Uint32[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapUint32Uint32 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapUint64Uint64) > 0 {
			SiZeMaP := func(k uint64, v uint64) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]uint64, 0, len(x.MapUint64Uint64))
				for k := range x.MapUint64Uint64 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapUint64Uint64[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapUint64Uint64 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapSint32Sint32) > 0 {
			SiZeMaP := func(k int32, v int32) {
				mapEntrySize := 1 + runtime.Soz(uint64(k)) + 1 + runtime.Soz(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int32, 0, len(x.MapSint32Sint32))
				for k := range x.MapSint32Sint32 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapSint32Sint32[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapSint32Sint32 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapSint64Sint64) > 0 {
			SiZeMaP := func(k int64, v int64) {
				mapEntrySize := 1 + runtime.Soz(uint64(k)) + 1 + runtime.Soz(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int64, 0, len(x.MapSint64Sint64))
				for k := range x.MapSint64Sint64 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapSint64Sint64[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapSint64Sint64 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapFixed32Fixed32) > 0 {
			SiZeMaP := func(k uint32, v uint32) {
				mapEntrySize := 1 + 4 + 1 + 4
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]uint32, 0, len(x.MapFixed32Fixed32))
				for k := range x.MapFixed32Fixed32 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapFixed32Fixed32[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapFixed32Fixed32 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapFixed64Fixed64) > 0 {
			SiZeMaP := func(k uint64, v uint64) {
				mapEntrySize := 1 + 8 + 1 + 8
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]uint64, 0, len(x.MapFixed64Fixed64))
				for k := range x.MapFixed64Fixed64 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapFixed64Fixed64[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapFixed64Fixed64 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapSfixed32Sfixed32) > 0 {
			SiZeMaP := func(k int32, v int32) {
				mapEntrySize := 1 + 4 + 1 + 4
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int32, 0, len(x.MapSfixed32Sfixed32))
				for k := range x.MapSfixed32Sfixed32 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapSfixed32Sfixed32[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapSfixed32Sfixed32 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapSfixed64Sfixed64) > 0 {
			SiZeMaP := func(k int64, v int64) {
				mapEntrySize := 1 + 8 + 1 + 8
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int64, 0, len(x.MapSfixed64Sfixed64))
				for k := range x.MapSfixed64Sfixed64 {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapSfixed64Sfixed64[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapSfixed64Sfixed64 {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapInt32Float) > 0 {
			SiZeMaP := func(k int32, v float32) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + 4
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int32, 0, len(x.MapInt32Float))
				for k := range x.MapInt32Float {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapInt32Float[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapInt32Float {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapInt32Double) > 0 {
			SiZeMaP := func(k int32, v float64) {
				mapEntrySize := 1 + runtime.Sov(uint64(k)) + 1 + 8
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]int32, 0, len(x.MapInt32Double))
				for k := range x.MapInt32Double {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return sortme[i] < sortme[j]
				})
				for _, k := range sortme {
					v := x.MapInt32Double[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapInt32Double {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapBoolBool) > 0 {
			SiZeMaP := func(k bool, v bool) {
				mapEntrySize := 1 + 1 + 1 + 1
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]bool, 0, len(x.MapBoolBool))
				for k := range x.MapBoolBool {
					sortme = append(sortme, k)
				}
				sort.Slice(sortme, func(i, j int) bool {
					return !sortme[i] && sortme[j]
				})
				for _, k := range sortme {
					v := x.MapBoolBool[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapBoolBool {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapStringString) > 0 {
			SiZeMaP := func(k string, v string) {
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + 1 + len(v) + runtime.Sov(uint64(len(v)))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]string, 0, len(x.MapStringString))
				for k := range x.MapStringString {
					sortme = append(sortme, k)
				}
				sort.Strings(sortme)
				for _, k := range sortme {
					v := x.MapStringString[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapStringString {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapStringBytes) > 0 {
			SiZeMaP := func(k string, v []byte) {
				l = 1 + len(v) + runtime.Sov(uint64(len(v)))
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + l
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]string, 0, len(x.MapStringBytes))
				for k := range x.MapStringBytes {
					sortme = append(sortme, k)
				}
				sort.Strings(sortme)
				for _, k := range sortme {
					v := x.MapStringBytes[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapStringBytes {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapStringNestedMessage) > 0 {
			SiZeMaP := func(k string, v *NestedMessage) {
				l := 0
				if v != nil {
					l = options.Size(v)
				}
				l += 1 + runtime.Sov(uint64(l))
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + l
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]string, 0, len(x.MapStringNestedMessage))
				for k := range x.MapStringNestedMessage {
					sortme = append(sortme, k)
				}
				sort.Strings(sortme)
				for _, k := range sortme {
					v := x.MapStringNestedMessage[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapStringNestedMessage {
					SiZeMaP(k, v)
				}
			}
		}
		if len(x.MapStringNestedEnum) > 0 {
			SiZeMaP := func(k string, v NestedEnum) {
				mapEntrySize := 1 + len(k) + runtime.Sov(uint64(len(k))) + 1 + runtime.Sov(uint64(v))
				n += mapEntrySize + 2 + runtime.Sov(uint64(mapEntrySize))
			}
			if options.Deterministic {
				sortme := make([]string, 0, len(x.MapStringNestedEnum))
				for k := range x.MapStringNestedEnum {
					sortme = append(sortme, k)
				}
				sort.Strings(sortme)
				for _, k := range sortme {
					v := x.MapStringNestedEnum[k]
					SiZeMaP(k, v)
				}
			} else {
				for k, v := range x.MapStringNestedEnum {
					SiZeMaP(k, v)
				}
			}
		}
		switch x := x.OneofField.(type) {
		case *TestAllTypes_OneofUint32:
			if x == nil {
				break
			}
			if x.OneofUint32 != 0 {
				n += 2 + runtime.Sov(uint64(x.OneofUint32))
			}
		case *TestAllTypes_OneofNestedMessage:
			if x == nil {
				break
			}
			if x.OneofNestedMessage != nil {
				l = options.Size(x.OneofNestedMessage)
				n += 2 + l + runtime.Sov(uint64(l))
			}
		case *TestAllTypes_OneofString:
			if x == nil {
				break
			}
			l = len(x.OneofString)
			if l > 0 {
				n += 2 + l + runtime.Sov(uint64(l))
			}
		case *TestAllTypes_OneofBytes:
			if x == nil {
				break
			}
			l = len(x.OneofBytes)
			if l > 0 {
				n += 2 + l + runtime.Sov(uint64(l))
			}
		case *TestAllTypes_OneofBool:
			if x == nil {
				break
			}
			if x.OneofBool {
				n += 3
			}
		case *TestAllTypes_OneofUint64:
			if x == nil {
				break
			}
			if x.OneofUint64 != 0 {
				n += 2 + runtime.Sov(uint64(x.OneofUint64))
			}
		case *TestAllTypes_OneofFloat:
			if x == nil {
				break
			}
			if x.OneofFloat != 0 || math.Signbit(float64(x.OneofFloat)) {
				n += 6
			}
		case *TestAllTypes_OneofDouble:
			if x == nil {
				break
			}
			if x.OneofDouble != 0 || math.Signbit(x.OneofDouble) {
				n += 10
			}
		case *TestAllTypes_OneofEnum:
			if x == nil {
				break
			}
			if x.OneofEnum != 0 {
				n += 2 + runtime.Sov(uint64(x.OneofEnum))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*TestAllTypes)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.OneofField.(type) {
		case *TestAllTypes_OneofUint32:
			if x.OneofUint32 != 0 {
				i = runtime.EncodeVarint(dAtA, i, uint64(x.OneofUint32))
				i--
				dAtA[i] = 0x6
				i--
				dAtA[i] = 0xf8
			}
		case *TestAllTypes_OneofNestedMessage:
			if x.OneofNestedMessage != nil {
				encoded, err := options.Marshal(x.OneofNestedMessage)
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0x82
			}
		case *TestAllTypes_OneofString:
			if len(x.OneofString) > 0 {
				i -= len(x.OneofString)
				copy(dAtA[i:], x.OneofString)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.OneofString)))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0x8a
			}
		case *TestAllTypes_OneofBytes:
			if len(x.OneofBytes) > 0 {
				i -= len(x.OneofBytes)
				copy(dAtA[i:], x.OneofBytes)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.OneofBytes)))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0x92
			}
		case *TestAllTypes_OneofBool:
			if x.OneofBool {
				i--
				if x.OneofBool {
					dAtA[i] = 1
				} else {
					dAtA[i] = 0
				}
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0x98
			}
		case *TestAllTypes_OneofUint64:
			if x.OneofUint64 != 0 {
				i = runtime.EncodeVarint(dAtA, i, uint64(x.OneofUint64))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0xa0
			}
		case *TestAllTypes_OneofFloat:
			if x.OneofFloat != 0 || math.Signbit(float64(x.OneofFloat)) {
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(x.OneofFloat))))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0xad
			}
		case *TestAllTypes_OneofDouble:
			if x.OneofDouble != 0 || math.Signbit(x.OneofDouble) {
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(x.OneofDouble))))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0xb1
			}
		case *TestAllTypes_OneofEnum:
			if x.OneofEnum != 0 {
				i = runtime.EncodeVarint(dAtA, i, uint64(x.OneofEnum))
				i--
				dAtA[i] = 0x7
				i--
				dAtA[i] = 0xb8
			}
		}
		if x.SingularImportEnum != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularImportEnum))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xb8
		}
		if x.SingularForeignEnum != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularForeignEnum))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xb0
		}
		if x.SingularNestedEnum != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularNestedEnum))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xa8
		}
		if x.SingularImportMessage != nil {
			encoded, err := options.Marshal(x.SingularImportMessage)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0xa2
		}
		if x.SingularForeignMessage != nil {
			encoded, err := options.Marshal(x.SingularForeignMessage)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0x9a
		}
		if x.SingularNestedMessage != nil {
			encoded, err := options.Marshal(x.SingularNestedMessage)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x6
			i--
			dAtA[i] = 0x92
		}
		if len(x.SingularBytes) > 0 {
			i -= len(x.SingularBytes)
			copy(dAtA[i:], x.SingularBytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SingularBytes)))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xfa
		}
		if len(x.SingularString) > 0 {
			i -= len(x.SingularString)
			copy(dAtA[i:], x.SingularString)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SingularString)))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xf2
		}
		if x.SingularBool {
			i--
			if x.SingularBool {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xe8
		}
		if x.SingularDouble != 0 || math.Signbit(x.SingularDouble) {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(x.SingularDouble))))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xe1
		}
		if x.SingularFloat != 0 || math.Signbit(float64(x.SingularFloat)) {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(x.SingularFloat))))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xdd
		}
		if x.SingularSfixed64 != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.SingularSfixed64))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xd1
		}
		if x.SingularSfixed32 != 0 {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.SingularSfixed32))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xcd
		}
		if x.SingularFixed64 != 0 {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.SingularFixed64))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xc1
		}
		if x.SingularFixed32 != 0 {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.SingularFixed32))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xbd
		}
		if x.SingularSint64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64((uint64(x.SingularSint64)<<1)^uint64((x.SingularSint64>>63))))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xb0
		}
		if x.SingularSint32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64((uint32(x.SingularSint32)<<1)^uint32((x.SingularSint32>>31))))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xa8
		}
		if x.SingularUint64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularUint64))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0xa0
		}
		if x.SingularUint32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularUint32))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0x98
		}
		if x.SingularInt64 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularInt64))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0x90
		}
		if x.SingularInt32 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SingularInt32))
			i--
			dAtA[i] = 0x5
			i--
			dAtA[i] = 0x88
		}
		if len(x.MapStringNestedEnum) > 0 {
			MaRsHaLmAp := func(k string, v NestedEnum) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i -= len(k)
				copy(dAtA[i:], k)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(k)))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0xca
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapStringNestedEnum := make([]string, 0, len(x.MapStringNestedEnum))
				for k := range x.MapStringNestedEnum {
					keysForMapStringNestedEnum = append(keysForMapStringNestedEnum, string(k))
				}
				sort.Slice(keysForMapStringNestedEnum, func(i, j int) bool {
					return keysForMapStringNestedEnum[i] < keysForMapStringNestedEnum[j]
				})
				for iNdEx := len(keysForMapStringNestedEnum) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapStringNestedEnum[string(keysForMapStringNestedEnum[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapStringNestedEnum[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapStringNestedEnum {
					v := x.MapStringNestedEnum[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapStringNestedMessage) > 0 {
			MaRsHaLmAp := func(k string, v *NestedMessage) (protoiface.MarshalOutput, error) {
				baseI := i
				encoded, err := options.Marshal(v)
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
				i -= len(k)
				copy(dAtA[i:], k)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(k)))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0xba
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapStringNestedMessage := make([]string, 0, len(x.MapStringNestedMessage))
				for k := range x.MapStringNestedMessage {
					keysForMapStringNestedMessage = append(keysForMapStringNestedMessage, string(k))
				}
				sort.Slice(keysForMapStringNestedMessage, func(i, j int) bool {
					return keysForMapStringNestedMessage[i] < keysForMapStringNestedMessage[j]
				})
				for iNdEx := len(keysForMapStringNestedMessage) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapStringNestedMessage[string(keysForMapStringNestedMessage[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapStringNestedMessage[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapStringNestedMessage {
					v := x.MapStringNestedMessage[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapStringBytes) > 0 {
			MaRsHaLmAp := func(k string, v []byte) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= len(v)
				copy(dAtA[i:], v)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
				i -= len(k)
				copy(dAtA[i:], k)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(k)))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0xb2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapStringBytes := make([]string, 0, len(x.MapStringBytes))
				for k := range x.MapStringBytes {
					keysForMapStringBytes = append(keysForMapStringBytes, string(k))
				}
				sort.Slice(keysForMapStringBytes, func(i, j int) bool {
					return keysForMapStringBytes[i] < keysForMapStringBytes[j]
				})
				for iNdEx := len(keysForMapStringBytes) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapStringBytes[string(keysForMapStringBytes[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapStringBytes[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapStringBytes {
					v := x.MapStringBytes[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapStringString) > 0 {
			MaRsHaLmAp := func(k string, v string) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= len(v)
				copy(dAtA[i:], v)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
				i -= len(k)
				copy(dAtA[i:], k)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(k)))
				i--
				dAtA[i] = 0xa
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0xaa
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapStringString := make([]string, 0, len(x.MapStringString))
				for k := range x.MapStringString {
					keysForMapStringString = append(keysForMapStringString, string(k))
				}
				sort.Slice(keysForMapStringString, func(i, j int) bool {
					return keysForMapStringString[i] < keysForMapStringString[j]
				})
				for iNdEx := len(keysForMapStringString) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapStringString[string(keysForMapStringString[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapStringString[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapStringString {
					v := x.MapStringString[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapBoolBool) > 0 {
			MaRsHaLmAp := func(k bool, v bool) (protoiface.MarshalOutput, error) {
				baseI := i
				i--
				if v {
					dAtA[i] = 1
				} else {
					dAtA[i] = 0
				}
				i--
				dAtA[i] = 0x10
				i--
				if k {
					dAtA[i] = 1
				} else {
					dAtA[i] = 0
				}
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0xa2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapBoolBool := make([]bool, 0, len(x.MapBoolBool))
				for k := range x.MapBoolBool {
					keysForMapBoolBool = append(keysForMapBoolBool, bool(k))
				}
				sort.Slice(keysForMapBoolBool, func(i, j int) bool {
					return !keysForMapBoolBool[i] && keysForMapBoolBool[j]
				})
				for iNdEx := len(keysForMapBoolBool) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapBoolBool[bool(keysForMapBoolBool[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapBoolBool[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapBoolBool {
					v := x.MapBoolBool[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapInt32Double) > 0 {
			MaRsHaLmAp := func(k int32, v float64) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(v))))
				i--
				dAtA[i] = 0x11
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0x9a
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapInt32Double := make([]int32, 0, len(x.MapInt32Double))
				for k := range x.MapInt32Double {
					keysForMapInt32Double = append(keysForMapInt32Double, int32(k))
				}
				sort.Slice(keysForMapInt32Double, func(i, j int) bool {
					return keysForMapInt32Double[i] < keysForMapInt32Double[j]
				})
				for iNdEx := len(keysForMapInt32Double) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapInt32Double[int32(keysForMapInt32Double[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapInt32Double[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapInt32Double {
					v := x.MapInt32Double[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapInt32Float) > 0 {
			MaRsHaLmAp := func(k int32, v float32) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(v))))
				i--
				dAtA[i] = 0x15
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0x92
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapInt32Float := make([]int32, 0, len(x.MapInt32Float))
				for k := range x.MapInt32Float {
					keysForMapInt32Float = append(keysForMapInt32Float, int32(k))
				}
				sort.Slice(keysForMapInt32Float, func(i, j int) bool {
					return keysForMapInt32Float[i] < keysForMapInt32Float[j]
				})
				for iNdEx := len(keysForMapInt32Float) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapInt32Float[int32(keysForMapInt32Float[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapInt32Float[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapInt32Float {
					v := x.MapInt32Float[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapSfixed64Sfixed64) > 0 {
			MaRsHaLmAp := func(k int64, v int64) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(v))
				i--
				dAtA[i] = 0x11
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(k))
				i--
				dAtA[i] = 0x9
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0x8a
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapSfixed64Sfixed64 := make([]int64, 0, len(x.MapSfixed64Sfixed64))
				for k := range x.MapSfixed64Sfixed64 {
					keysForMapSfixed64Sfixed64 = append(keysForMapSfixed64Sfixed64, int64(k))
				}
				sort.Slice(keysForMapSfixed64Sfixed64, func(i, j int) bool {
					return keysForMapSfixed64Sfixed64[i] < keysForMapSfixed64Sfixed64[j]
				})
				for iNdEx := len(keysForMapSfixed64Sfixed64) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapSfixed64Sfixed64[int64(keysForMapSfixed64Sfixed64[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapSfixed64Sfixed64[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapSfixed64Sfixed64 {
					v := x.MapSfixed64Sfixed64[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapSfixed32Sfixed32) > 0 {
			MaRsHaLmAp := func(k int32, v int32) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(v))
				i--
				dAtA[i] = 0x15
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(k))
				i--
				dAtA[i] = 0xd
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x4
				i--
				dAtA[i] = 0x82
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapSfixed32Sfixed32 := make([]int32, 0, len(x.MapSfixed32Sfixed32))
				for k := range x.MapSfixed32Sfixed32 {
					keysForMapSfixed32Sfixed32 = append(keysForMapSfixed32Sfixed32, int32(k))
				}
				sort.Slice(keysForMapSfixed32Sfixed32, func(i, j int) bool {
					return keysForMapSfixed32Sfixed32[i] < keysForMapSfixed32Sfixed32[j]
				})
				for iNdEx := len(keysForMapSfixed32Sfixed32) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapSfixed32Sfixed32[int32(keysForMapSfixed32Sfixed32[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapSfixed32Sfixed32[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapSfixed32Sfixed32 {
					v := x.MapSfixed32Sfixed32[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapFixed64Fixed64) > 0 {
			MaRsHaLmAp := func(k uint64, v uint64) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(v))
				i--
				dAtA[i] = 0x11
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(k))
				i--
				dAtA[i] = 0x9
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xfa
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapFixed64Fixed64 := make([]uint64, 0, len(x.MapFixed64Fixed64))
				for k := range x.MapFixed64Fixed64 {
					keysForMapFixed64Fixed64 = append(keysForMapFixed64Fixed64, uint64(k))
				}
				sort.Slice(keysForMapFixed64Fixed64, func(i, j int) bool {
					return keysForMapFixed64Fixed64[i] < keysForMapFixed64Fixed64[j]
				})
				for iNdEx := len(keysForMapFixed64Fixed64) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapFixed64Fixed64[uint64(keysForMapFixed64Fixed64[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapFixed64Fixed64[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapFixed64Fixed64 {
					v := x.MapFixed64Fixed64[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapFixed32Fixed32) > 0 {
			MaRsHaLmAp := func(k uint32, v uint32) (protoiface.MarshalOutput, error) {
				baseI := i
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(v))
				i--
				dAtA[i] = 0x15
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(k))
				i--
				dAtA[i] = 0xd
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xf2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapFixed32Fixed32 := make([]uint32, 0, len(x.MapFixed32Fixed32))
				for k := range x.MapFixed32Fixed32 {
					keysForMapFixed32Fixed32 = append(keysForMapFixed32Fixed32, uint32(k))
				}
				sort.Slice(keysForMapFixed32Fixed32, func(i, j int) bool {
					return keysForMapFixed32Fixed32[i] < keysForMapFixed32Fixed32[j]
				})
				for iNdEx := len(keysForMapFixed32Fixed32) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapFixed32Fixed32[uint32(keysForMapFixed32Fixed32[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapFixed32Fixed32[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapFixed32Fixed32 {
					v := x.MapFixed32Fixed32[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapSint64Sint64) > 0 {
			MaRsHaLmAp := func(k int64, v int64) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64((uint64(v)<<1)^uint64((v>>63))))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64((uint64(k)<<1)^uint64((k>>63))))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xea
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapSint64Sint64 := make([]int64, 0, len(x.MapSint64Sint64))
				for k := range x.MapSint64Sint64 {
					keysForMapSint64Sint64 = append(keysForMapSint64Sint64, int64(k))
				}
				sort.Slice(keysForMapSint64Sint64, func(i, j int) bool {
					return keysForMapSint64Sint64[i] < keysForMapSint64Sint64[j]
				})
				for iNdEx := len(keysForMapSint64Sint64) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapSint64Sint64[int64(keysForMapSint64Sint64[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapSint64Sint64[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapSint64Sint64 {
					v := x.MapSint64Sint64[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapSint32Sint32) > 0 {
			MaRsHaLmAp := func(k int32, v int32) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64((uint32(v)<<1)^uint32((v>>31))))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64((uint32(k)<<1)^uint32((k>>31))))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xe2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapSint32Sint32 := make([]int32, 0, len(x.MapSint32Sint32))
				for k := range x.MapSint32Sint32 {
					keysForMapSint32Sint32 = append(keysForMapSint32Sint32, int32(k))
				}
				sort.Slice(keysForMapSint32Sint32, func(i, j int) bool {
					return keysForMapSint32Sint32[i] < keysForMapSint32Sint32[j]
				})
				for iNdEx := len(keysForMapSint32Sint32) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapSint32Sint32[int32(keysForMapSint32Sint32[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapSint32Sint32[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapSint32Sint32 {
					v := x.MapSint32Sint32[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapUint64Uint64) > 0 {
			MaRsHaLmAp := func(k uint64, v uint64) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xda
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapUint64Uint64 := make([]uint64, 0, len(x.MapUint64Uint64))
				for k := range x.MapUint64Uint64 {
					keysForMapUint64Uint64 = append(keysForMapUint64Uint64, uint64(k))
				}
				sort.Slice(keysForMapUint64Uint64, func(i, j int) bool {
					return keysForMapUint64Uint64[i] < keysForMapUint64Uint64[j]
				})
				for iNdEx := len(keysForMapUint64Uint64) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapUint64Uint64[uint64(keysForMapUint64Uint64[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapUint64Uint64[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapUint64Uint64 {
					v := x.MapUint64Uint64[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapUint32Uint32) > 0 {
			MaRsHaLmAp := func(k uint32, v uint32) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xd2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapUint32Uint32 := make([]uint32, 0, len(x.MapUint32Uint32))
				for k := range x.MapUint32Uint32 {
					keysForMapUint32Uint32 = append(keysForMapUint32Uint32, uint32(k))
				}
				sort.Slice(keysForMapUint32Uint32, func(i, j int) bool {
					return keysForMapUint32Uint32[i] < keysForMapUint32Uint32[j]
				})
				for iNdEx := len(keysForMapUint32Uint32) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapUint32Uint32[uint32(keysForMapUint32Uint32[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapUint32Uint32[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapUint32Uint32 {
					v := x.MapUint32Uint32[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapInt64Int64) > 0 {
			MaRsHaLmAp := func(k int64, v int64) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xca
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapInt64Int64 := make([]int64, 0, len(x.MapInt64Int64))
				for k := range x.MapInt64Int64 {
					keysForMapInt64Int64 = append(keysForMapInt64Int64, int64(k))
				}
				sort.Slice(keysForMapInt64Int64, func(i, j int) bool {
					return keysForMapInt64Int64[i] < keysForMapInt64Int64[j]
				})
				for iNdEx := len(keysForMapInt64Int64) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapInt64Int64[int64(keysForMapInt64Int64[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapInt64Int64[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapInt64Int64 {
					v := x.MapInt64Int64[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.MapInt32Int32) > 0 {
			MaRsHaLmAp := func(k int32, v int32) (protoiface.MarshalOutput, error) {
				baseI := i
				i = runtime.EncodeVarint(dAtA, i, uint64(v))
				i--
				dAtA[i] = 0x10
				i = runtime.EncodeVarint(dAtA, i, uint64(k))
				i--
				dAtA[i] = 0x8
				i = runtime.EncodeVarint(dAtA, i, uint64(baseI-i))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0xc2
				return protoiface.MarshalOutput{}, nil
			}
			if options.Deterministic {
				keysForMapInt32Int32 := make([]int32, 0, len(x.MapInt32Int32))
				for k := range x.MapInt32Int32 {
					keysForMapInt32Int32 = append(keysForMapInt32Int32, int32(k))
				}
				sort.Slice(keysForMapInt32Int32, func(i, j int) bool {
					return keysForMapInt32Int32[i] < keysForMapInt32Int32[j]
				})
				for iNdEx := len(keysForMapInt32Int32) - 1; iNdEx >= 0; iNdEx-- {
					v := x.MapInt32Int32[int32(keysForMapInt32Int32[iNdEx])]
					out, err := MaRsHaLmAp(keysForMapInt32Int32[iNdEx], v)
					if err != nil {
						return out, err
					}
				}
			} else {
				for k := range x.MapInt32Int32 {
					v := x.MapInt32Int32[k]
					out, err := MaRsHaLmAp(k, v)
					if err != nil {
						return out, err
					}
				}
			}
		}
		if len(x.RepeatedImportenum) > 0 {
			var pksize2 int
			for _, num := range x.RepeatedImportenum {
				pksize2 += runtime.Sov(uint64(num))
			}
			i -= pksize2
			j1 := i
			for _, num1 := range x.RepeatedImportenum {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j1++
				}
				dAtA[j1] = uint8(num)
				j1++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize2))
			i--
			dAtA[i] = 0x3
			i--
			dAtA[i] = 0xaa
		}
		if len(x.RepeatedForeignEnum) > 0 {
			var pksize4 int
			for _, num := range x.RepeatedForeignEnum {
				pksize4 += runtime.Sov(uint64(num))
			}
			i -= pksize4
			j3 := i
			for _, num1 := range x.RepeatedForeignEnum {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j3] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j3++
				}
				dAtA[j3] = uint8(num)
				j3++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize4))
			i--
			dAtA[i] = 0x3
			i--
			dAtA[i] = 0xa2
		}
		if len(x.RepeatedNestedEnum) > 0 {
			var pksize6 int
			for _, num := range x.RepeatedNestedEnum {
				pksize6 += runtime.Sov(uint64(num))
			}
			i -= pksize6
			j5 := i
			for _, num1 := range x.RepeatedNestedEnum {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j5] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j5++
				}
				dAtA[j5] = uint8(num)
				j5++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize6))
			i--
			dAtA[i] = 0x3
			i--
			dAtA[i] = 0x9a
		}
		if len(x.RepeatedImportmessage) > 0 {
			for iNdEx := len(x.RepeatedImportmessage) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.RepeatedImportmessage[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0x92
			}
		}
		if len(x.RepeatedForeignMessage) > 0 {
			for iNdEx := len(x.RepeatedForeignMessage) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.RepeatedForeignMessage[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0x8a
			}
		}
		if len(x.RepeatedNestedMessage) > 0 {
			for iNdEx := len(x.RepeatedNestedMessage) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.RepeatedNestedMessage[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x3
				i--
				dAtA[i] = 0x82
			}
		}
		if len(x.RepeatedBytes) > 0 {
			for iNdEx := len(x.RepeatedBytes) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.RepeatedBytes[iNdEx])
				copy(dAtA[i:], x.RepeatedBytes[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedBytes[iNdEx])))
				i--
				dAtA[i] = 0x2
				i--
				dAtA[i] = 0xea
			}
		}
		if len(x.RepeatedString) > 0 {
			for iNdEx := len(x.RepeatedString) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.RepeatedString[iNdEx])
				copy(dAtA[i:], x.RepeatedString[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedString[iNdEx])))
				i--
				dAtA[i] = 0x2
				i--
				dAtA[i] = 0xe2
			}
		}
		if len(x.RepeatedBool) > 0 {
			for iNdEx := len(x.RepeatedBool) - 1; iNdEx >= 0; iNdEx-- {
				i--
				if x.RepeatedBool[iNdEx] {
					dAtA[i] = 1
				} else {
					dAtA[i] = 0
				}
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedBool)))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xda
		}
		if len(x.RepeatedDouble) > 0 {
			for iNdEx := len(x.RepeatedDouble) - 1; iNdEx >= 0; iNdEx-- {
				f7 := math.Float64bits(float64(x.RepeatedDouble[iNdEx]))
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(f7))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedDouble)*8))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xd2
		}
		if len(x.RepeatedFloat) > 0 {
			for iNdEx := len(x.RepeatedFloat) - 1; iNdEx >= 0; iNdEx-- {
				f8 := math.Float32bits(float32(x.RepeatedFloat[iNdEx]))
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(f8))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedFloat)*4))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xca
		}
		if len(x.RepeatedSfixed64) > 0 {
			for iNdEx := len(x.RepeatedSfixed64) - 1; iNdEx >= 0; iNdEx-- {
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.RepeatedSfixed64[iNdEx]))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedSfixed64)*8))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xc2
		}
		if len(x.RepeatedSfixed32) > 0 {
			for iNdEx := len(x.RepeatedSfixed32) - 1; iNdEx >= 0; iNdEx-- {
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.RepeatedSfixed32[iNdEx]))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedSfixed32)*4))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xba
		}
		if len(x.RepeatedFixed64) > 0 {
			for iNdEx := len(x.RepeatedFixed64) - 1; iNdEx >= 0; iNdEx-- {
				i -= 8
				binary.LittleEndian.PutUint64(dAtA[i:], uint64(x.RepeatedFixed64[iNdEx]))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedFixed64)*8))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xb2
		}
		if len(x.RepeatedFixed32) > 0 {
			for iNdEx := len(x.RepeatedFixed32) - 1; iNdEx >= 0; iNdEx-- {
				i -= 4
				binary.LittleEndian.PutUint32(dAtA[i:], uint32(x.RepeatedFixed32[iNdEx]))
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RepeatedFixed32)*4))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xaa
		}
		if len(x.RepeatedSint64) > 0 {
			var pksize10 int
			for _, num := range x.RepeatedSint64 {
				pksize10 += runtime.Soz(uint64(num))
			}
			i -= pksize10
			j9 := i
			for _, num := range x.RepeatedSint64 {
				x11 := (uint64(num) << 1) ^ uint64((num >> 63))
				for x11 >= 1<<7 {
					dAtA[j9] = uint8(uint64(x11)&0x7f | 0x80)
					j9++
					x11 >>= 7
				}
				dAtA[j9] = uint8(x11)
				j9++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize10))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0xa2
		}
		if len(x.RepeatedSint32) > 0 {
			var pksize13 int
			for _, num := range x.RepeatedSint32 {
				pksize13 += runtime.Soz(uint64(num))
			}
			i -= pksize13
			j12 := i
			for _, num := range x.RepeatedSint32 {
				x14 := (uint32(num) << 1) ^ uint32((num >> 31))
				for x14 >= 1<<7 {
					dAtA[j12] = uint8(uint64(x14)&0x7f | 0x80)
					j12++
					x14 >>= 7
				}
				dAtA[j12] = uint8(x14)
				j12++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize13))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0x9a
		}
		if len(x.RepeatedUint64) > 0 {
			var pksize16 int
			for _, num := range x.RepeatedUint64 {
				pksize16 += runtime.Sov(uint64(num))
			}
			i -= pksize16
			j15 := i
			for _, num := range x.RepeatedUint64 {
				for num >= 1<<7 {
					dAtA[j15] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j15++
				}
				dAtA[j15] = uint8(num)
				j15++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize16))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0x92
		}
		if len(x.RepeatedUint32) > 0 {
			var pksize18 int
			for _, num := range x.RepeatedUint32 {
				pksize18 += runtime.Sov(uint64(num))
			}
			i -= pksize18
			j17 := i
			for _, num := range x.RepeatedUint32 {
				for num >= 1<<7 {
					dAtA[j17] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j17++
				}
				dAtA[j17] = uint8(num)
				j17++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize18))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0x8a
		}
		if len(x.RepeatedInt64) > 0 {
			var pksize20 int
			for _, num := range x.RepeatedInt64 {
				pksize20 += runtime.Sov(uint64(num))
			}
			i -= pksize20
			j19 := i
			for _, num1 := range x.RepeatedInt64 {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j19] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j19++
				}
				dAtA[j19] = uint8(num)
				j19++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize20))
			i--
			dAtA[i] = 0x2
			i--
			dAtA[i] = 0x82
		}
		if len(x.RepeatedInt32) > 0 {
			var pksize22 int
			for _, num := range x.RepeatedInt32 {
				pksize22 += runtime.Sov(uint64(num))
			}
			i -= pksize22
			j21 := i
			for _, num1 := range x.RepeatedInt32 {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j21] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j21++
				}
				dAtA[j21] = uint8(num)
				j21++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize22))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xfa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*TestAllTypes)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TestAllTypes: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TestAllTypes: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 81:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularInt32", wireType)
				}
				x.SingularInt32 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularInt32 |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 82:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularInt64", wireType)
				}
				x.SingularInt64 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularInt64 |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 83:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularUint32", wireType)
				}
				x.SingularUint32 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularUint32 |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 84:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularUint64", wireType)
				}
				x.SingularUint64 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularUint64 |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 85:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularSint32", wireType)
				}
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
				x.SingularSint32 = v
			case 86:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularSint64", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
				x.SingularSint64 = int64(v)
			case 87:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularFixed32", wireType)
				}
				x.SingularFixed32 = 0
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularFixed32 = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
			case 88:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularFixed64", wireType)
				}
				x.SingularFixed64 = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularFixed64 = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			case 89:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularSfixed32", wireType)
				}
				x.SingularSfixed32 = 0
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularSfixed32 = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
			case 90:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularSfixed64", wireType)
				}
				x.SingularSfixed64 = 0
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularSfixed64 = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
			case 91:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularFloat", wireType)
				}
				var v uint32
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				x.SingularFloat = float32(math.Float32frombits(v))
			case 92:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularDouble", wireType)
				}
				var v uint64
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				x.SingularDouble = float64(math.Float64frombits(v))
			case 93:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularBool", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.SingularBool = bool(v != 0)
			case 94:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularString", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularString = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 95:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SingularBytes = append(x.SingularBytes[:0], dAtA[iNdEx:postIndex]...)
				if x.SingularBytes == nil {
					x.SingularBytes = []byte{}
				}
				iNdEx = postIndex
			case 98:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularNestedMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SingularNestedMessage == nil {
					x.SingularNestedMessage = &NestedMessage{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SingularNestedMessage); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 99:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularForeignMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SingularForeignMessage == nil {
					x.SingularForeignMessage = &ForeignMessage{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SingularForeignMessage); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 100:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularImportMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SingularImportMessage == nil {
					x.SingularImportMessage = &ImportMessage{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SingularImportMessage); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 101:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularNestedEnum", wireType)
				}
				x.SingularNestedEnum = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularNestedEnum |= NestedEnum(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 102:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularForeignEnum", wireType)
				}
				x.SingularForeignEnum = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularForeignEnum |= ForeignEnum(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 103:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SingularImportEnum", wireType)
				}
				x.SingularImportEnum = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SingularImportEnum |= ImportEnum(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 31:
				if wireType == 0 {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedInt32 = append(x.RepeatedInt32, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedInt32) == 0 {
						x.RepeatedInt32 = make([]int32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int32
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedInt32 = append(x.RepeatedInt32, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedInt32", wireType)
				}
			case 32:
				if wireType == 0 {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedInt64 = append(x.RepeatedInt64, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedInt64) == 0 {
						x.RepeatedInt64 = make([]int64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= int64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedInt64 = append(x.RepeatedInt64, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedInt64", wireType)
				}
			case 33:
				if wireType == 0 {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedUint32 = append(x.RepeatedUint32, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedUint32) == 0 {
						x.RepeatedUint32 = make([]uint32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint32
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= uint32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedUint32 = append(x.RepeatedUint32, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedUint32", wireType)
				}
			case 34:
				if wireType == 0 {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedUint64 = append(x.RepeatedUint64, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedUint64) == 0 {
						x.RepeatedUint64 = make([]uint64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedUint64 = append(x.RepeatedUint64, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedUint64", wireType)
				}
			case 35:
				if wireType == 0 {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
					x.RepeatedSint32 = append(x.RepeatedSint32, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedSint32) == 0 {
						x.RepeatedSint32 = make([]int32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int32
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
						x.RepeatedSint32 = append(x.RepeatedSint32, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedSint32", wireType)
				}
			case 36:
				if wireType == 0 {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
					x.RepeatedSint64 = append(x.RepeatedSint64, int64(v))
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.RepeatedSint64) == 0 {
						x.RepeatedSint64 = make([]int64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
						x.RepeatedSint64 = append(x.RepeatedSint64, int64(v))
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedSint64", wireType)
				}
			case 37:
				if wireType == 5 {
					var v uint32
					if (iNdEx + 4) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
					iNdEx += 4
					x.RepeatedFixed32 = append(x.RepeatedFixed32, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 4
					if elementCount != 0 && len(x.RepeatedFixed32) == 0 {
						x.RepeatedFixed32 = make([]uint32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint32
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
						x.RepeatedFixed32 = append(x.RepeatedFixed32, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedFixed32", wireType)
				}
			case 38:
				if wireType == 1 {
					var v uint64
					if (iNdEx + 8) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					x.RepeatedFixed64 = append(x.RepeatedFixed64, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 8
					if elementCount != 0 && len(x.RepeatedFixed64) == 0 {
						x.RepeatedFixed64 = make([]uint64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint64
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
						x.RepeatedFixed64 = append(x.RepeatedFixed64, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedFixed64", wireType)
				}
			case 39:
				if wireType == 5 {
					var v int32
					if (iNdEx + 4) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
					iNdEx += 4
					x.RepeatedSfixed32 = append(x.RepeatedSfixed32, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 4
					if elementCount != 0 && len(x.RepeatedSfixed32) == 0 {
						x.RepeatedSfixed32 = make([]int32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int32
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
						x.RepeatedSfixed32 = append(x.RepeatedSfixed32, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedSfixed32", wireType)
				}
			case 40:
				if wireType == 1 {
					var v int64
					if (iNdEx + 8) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					x.RepeatedSfixed64 = append(x.RepeatedSfixed64, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 8
					if elementCount != 0 && len(x.RepeatedSfixed64) == 0 {
						x.RepeatedSfixed64 = make([]int64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int64
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
						x.RepeatedSfixed64 = append(x.RepeatedSfixed64, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedSfixed64", wireType)
				}
			case 41:
				if wireType == 5 {
					var v uint32
					if (iNdEx + 4) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
					iNdEx += 4
					v2 := float32(math.Float32frombits(v))
					x.RepeatedFloat = append(x.RepeatedFloat, v2)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 4
					if elementCount != 0 && len(x.RepeatedFloat) == 0 {
						x.RepeatedFloat = make([]float32, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint32
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
						v2 := float32(math.Float32frombits(v))
						x.RepeatedFloat = append(x.RepeatedFloat, v2)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedFloat", wireType)
				}
			case 42:
				if wireType == 1 {
					var v uint64
					if (iNdEx + 8) > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					x.RepeatedDouble = append(x.RepeatedDouble, v2)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen / 8
					if elementCount != 0 && len(x.RepeatedDouble) == 0 {
						x.RepeatedDouble = make([]float64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v uint64
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
						v2 := float64(math.Float64frombits(v))
						x.RepeatedDouble = append(x.RepeatedDouble, v2)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedDouble", wireType)
				}
			case 43:
				if wireType == 0 {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedBool = append(x.RepeatedBool, bool(v != 0))
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					elementCount = packedLen
					if elementCount != 0 && len(x.RepeatedBool) == 0 {
						x.RepeatedBool = make([]bool, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= int(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedBool = append(x.RepeatedBool, bool(v != 0))
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedBool", wireType)
				}
			case 44:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedString", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RepeatedString = append(x.RepeatedString, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 45:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RepeatedBytes = append(x.RepeatedBytes, make([]byte, postIndex-iNdEx))
				copy(x.RepeatedBytes[len(x.RepeatedBytes)-1], dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 48:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedNestedMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RepeatedNestedMessage = append(x.RepeatedNestedMessage, &NestedMessage{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.RepeatedNestedMessage[len(x.RepeatedNestedMessage)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 49:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedForeignMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RepeatedForeignMessage = append(x.RepeatedForeignMessage, &ForeignMessage{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.RepeatedForeignMessage[len(x.RepeatedForeignMessage)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 50:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedImportmessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RepeatedImportmessage = append(x.RepeatedImportmessage, &ImportMessage{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.RepeatedImportmessage[len(x.RepeatedImportmessage)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 51:
				if wireType == 0 {
					var v NestedEnum
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= NestedEnum(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedNestedEnum = append(x.RepeatedNestedEnum, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					if elementCount != 0 && len(x.RepeatedNestedEnum) == 0 {
						x.RepeatedNestedEnum = make([]NestedEnum, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v NestedEnum
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= NestedEnum(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedNestedEnum = append(x.RepeatedNestedEnum, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedNestedEnum", wireType)
				}
			case 52:
				if wireType == 0 {
					var v ForeignEnum
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ForeignEnum(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedForeignEnum = append(x.RepeatedForeignEnum, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					if elementCount != 0 && len(x.RepeatedForeignEnum) == 0 {
						x.RepeatedForeignEnum = make([]ForeignEnum, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v ForeignEnum
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= ForeignEnum(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedForeignEnum = append(x.RepeatedForeignEnum, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedForeignEnum", wireType)
				}
			case 53:
				if wireType == 0 {
					var v ImportEnum
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ImportEnum(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.RepeatedImportenum = append(x.RepeatedImportenum, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					if elementCount != 0 && len(x.RepeatedImportenum) == 0 {
						x.RepeatedImportenum = make([]ImportEnum, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v ImportEnum
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= ImportEnum(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.RepeatedImportenum = append(x.RepeatedImportenum, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RepeatedImportenum", wireType)
				}
			case 56:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapInt32Int32", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapInt32Int32 == nil {
					x.MapInt32Int32 = make(map[int32]int32)
				}
				var mapkey int32
				var mapvalue int32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapInt32Int32[mapkey] = mapvalue
				iNdEx = postIndex
			case 57:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapInt64Int64", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapInt64Int64 == nil {
					x.MapInt64Int64 = make(map[int64]int64)
				}
				var mapkey int64
				var mapvalue int64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= int64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= int64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapInt64Int64[mapkey] = mapvalue
				iNdEx = postIndex
			case 58:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapUint32Uint32", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapUint32Uint32 == nil {
					x.MapUint32Uint32 = make(map[uint32]uint32)
				}
				var mapkey uint32
				var mapvalue uint32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= uint32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= uint32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapUint32Uint32[mapkey] = mapvalue
				iNdEx = postIndex
			case 59:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapUint64Uint64", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapUint64Uint64 == nil {
					x.MapUint64Uint64 = make(map[uint64]uint64)
				}
				var mapkey uint64
				var mapvalue uint64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapUint64Uint64[mapkey] = mapvalue
				iNdEx = postIndex
			case 60:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapSint32Sint32", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapSint32Sint32 == nil {
					x.MapSint32Sint32 = make(map[int32]int32)
				}
				var mapkey int32
				var mapvalue int32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						var mapkeytemp int32
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkeytemp |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapkeytemp = int32((uint32(mapkeytemp) >> 1) ^ uint32(((mapkeytemp&1)<<31)>>31))
						mapkey = int32(mapkeytemp)
					} else if fieldNum == 2 {
						var mapvaluetemp int32
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvaluetemp |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapvaluetemp = int32((uint32(mapvaluetemp) >> 1) ^ uint32(((mapvaluetemp&1)<<31)>>31))
						mapvalue = int32(mapvaluetemp)
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapSint32Sint32[mapkey] = mapvalue
				iNdEx = postIndex
			case 61:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapSint64Sint64", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapSint64Sint64 == nil {
					x.MapSint64Sint64 = make(map[int64]int64)
				}
				var mapkey int64
				var mapvalue int64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						var mapkeytemp uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkeytemp |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapkeytemp = (mapkeytemp >> 1) ^ uint64((int64(mapkeytemp&1)<<63)>>63)
						mapkey = int64(mapkeytemp)
					} else if fieldNum == 2 {
						var mapvaluetemp uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvaluetemp |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapvaluetemp = (mapvaluetemp >> 1) ^ uint64((int64(mapvaluetemp&1)<<63)>>63)
						mapvalue = int64(mapvaluetemp)
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapSint64Sint64[mapkey] = mapvalue
				iNdEx = postIndex
			case 62:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapFixed32Fixed32", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapFixed32Fixed32 == nil {
					x.MapFixed32Fixed32 = make(map[uint32]uint32)
				}
				var mapkey uint32
				var mapvalue uint32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
					} else if fieldNum == 2 {
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapFixed32Fixed32[mapkey] = mapvalue
				iNdEx = postIndex
			case 63:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapFixed64Fixed64", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapFixed64Fixed64 == nil {
					x.MapFixed64Fixed64 = make(map[uint64]uint64)
				}
				var mapkey uint64
				var mapvalue uint64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
					} else if fieldNum == 2 {
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapFixed64Fixed64[mapkey] = mapvalue
				iNdEx = postIndex
			case 64:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapSfixed32Sfixed32", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapSfixed32Sfixed32 == nil {
					x.MapSfixed32Sfixed32 = make(map[int32]int32)
				}
				var mapkey int32
				var mapvalue int32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
					} else if fieldNum == 2 {
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapSfixed32Sfixed32[mapkey] = mapvalue
				iNdEx = postIndex
			case 65:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapSfixed64Sfixed64", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapSfixed64Sfixed64 == nil {
					x.MapSfixed64Sfixed64 = make(map[int64]int64)
				}
				var mapkey int64
				var mapvalue int64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
					} else if fieldNum == 2 {
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapSfixed64Sfixed64[mapkey] = mapvalue
				iNdEx = postIndex
			case 66:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapInt32Float", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapInt32Float == nil {
					x.MapInt32Float = make(map[int32]float32)
				}
				var mapkey int32
				var mapvalue float32
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						var mapvaluetemp uint32
						if (iNdEx + 4) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvaluetemp = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
						iNdEx += 4
						mapvalue = math.Float32frombits(mapvaluetemp)
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapInt32Float[mapkey] = mapvalue
				iNdEx = postIndex
			case 67:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapInt32Double", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapInt32Double == nil {
					x.MapInt32Double = make(map[int32]float64)
				}
				var mapkey int32
				var mapvalue float64
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkey |= int32(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else if fieldNum == 2 {
						var mapvaluetemp uint64
						if (iNdEx + 8) > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvaluetemp = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
						iNdEx += 8
						mapvalue = math.Float64frombits(mapvaluetemp)
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapInt32Double[mapkey] = mapvalue
				iNdEx = postIndex
			case 68:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapBoolBool", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapBoolBool == nil {
					x.MapBoolBool = make(map[bool]bool)
				}
				var mapkey bool
				var mapvalue bool
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
						var mapkeytemp int
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapkeytemp |= int(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapkey = bool(mapkeytemp != 0)
					} else if fieldNum == 2 {
						var mapvaluetemp int
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvaluetemp |= int(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						mapvalue = bool(mapvaluetemp != 0)
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapBoolBool[mapkey] = mapvalue
				iNdEx = postIndex
			case 69:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapStringString", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapStringString == nil {
					x.MapStringString = make(map[string]string)
				}
				var mapkey string
				var mapvalue string
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						var stringLenmapvalue uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							stringLenmapvalue |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						intStringLenmapvalue := int(stringLenmapvalue)
						if intStringLenmapvalue < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapvalue := iNdEx + intStringLenmapvalue
						if postStringIndexmapvalue < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapvalue > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
						iNdEx = postStringIndexmapvalue
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapStringString[mapkey] = mapvalue
				iNdEx = postIndex
			case 70:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapStringBytes", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapStringBytes == nil {
					x.MapStringBytes = make(map[string][]byte)
				}
				var mapkey string
				var mapvalue []byte
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						var mapbyteLen uint64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapbyteLen |= uint64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						intMapbyteLen := int(mapbyteLen)
						if intMapbyteLen < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postbytesIndex := iNdEx + intMapbyteLen
						if postbytesIndex < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postbytesIndex > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = make([]byte, mapbyteLen)
						copy(mapvalue, dAtA[iNdEx:postbytesIndex])
						iNdEx = postbytesIndex
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapStringBytes[mapkey] = mapvalue
				iNdEx = postIndex
			case 71:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapStringNestedMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapStringNestedMessage == nil {
					x.MapStringNestedMessage = make(map[string]*NestedMessage)
				}
				var mapkey string
				var mapvalue *NestedMessage
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						var mapmsglen int
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapmsglen |= int(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						if mapmsglen < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postmsgIndex := iNdEx + mapmsglen
						if postmsgIndex < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postmsgIndex > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapvalue = &NestedMessage{}
						if err := options.Unmarshal(dAtA[iNdEx:postmsgIndex], mapvalue); err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						iNdEx = postmsgIndex
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapStringNestedMessage[mapkey] = mapvalue
				iNdEx = postIndex
			case 73:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MapStringNestedEnum", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MapStringNestedEnum == nil {
					x.MapStringNestedEnum = make(map[string]NestedEnum)
				}
				var mapkey string
				var mapvalue NestedEnum
				for iNdEx < postIndex {
					entryPreIndex := iNdEx
					var wire uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						postStringIndexmapkey := iNdEx + intStringLenmapkey
						if postStringIndexmapkey < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if postStringIndexmapkey > l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
						iNdEx = postStringIndexmapkey
					} else if fieldNum == 2 {
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							mapvalue |= NestedEnum(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
					} else {
						iNdEx = entryPreIndex
						skippy, err := runtime.Skip(dAtA[iNdEx:])
						if err != nil {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
						}
						if (skippy < 0) || (iNdEx+skippy) < 0 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
						}
						if (iNdEx + skippy) > postIndex {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						iNdEx += skippy
					}
				}
				x.MapStringNestedEnum[mapkey] = mapvalue
				iNdEx = postIndex
			case 111:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofUint32", wireType)
				}
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.OneofField = &TestAllTypes_OneofUint32{v}
			case 112:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofNestedMessage", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &NestedMessage{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.OneofField = &TestAllTypes_OneofNestedMessage{v}
				iNdEx = postIndex
			case 113:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofString", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.OneofField = &TestAllTypes_OneofString{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 114:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofBytes", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := make([]byte, postIndex-iNdEx)
				copy(v, dAtA[iNdEx:postIndex])
				x.OneofField = &TestAllTypes_OneofBytes{v}
				iNdEx = postIndex
			case 115:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofBool", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				b := bool(v != 0)
				x.OneofField = &TestAllTypes_OneofBool{b}
			case 116:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofUint64", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.OneofField = &TestAllTypes_OneofUint64{v}
			case 117:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofFloat", wireType)
				}
				var v uint32
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				x.OneofField = &TestAllTypes_OneofFloat{float32(math.Float32frombits(v))}
			case 118:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofDouble", wireType)
				}
				var v uint64
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				x.OneofField = &TestAllTypes_OneofDouble{float64(math.Float64frombits(v))}
			case 119:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OneofEnum", wireType)
				}
				var v NestedEnum
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= NestedEnum(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.OneofField = &TestAllTypes_OneofEnum{v}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ForeignMessage   protoreflect.MessageDescriptor
	fd_ForeignMessage_c protoreflect.FieldDescriptor
	fd_ForeignMessage_d protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_proto_init()
	md_ForeignMessage = File_internal_testprotos_test3_test_proto.Messages().ByName("ForeignMessage")
	fd_ForeignMessage_c = md_ForeignMessage.Fields().ByName("c")
	fd_ForeignMessage_d = md_ForeignMessage.Fields().ByName("d")
}

var _ protoreflect.Message = (*fastReflection_ForeignMessage)(nil)

type fastReflection_ForeignMessage ForeignMessage

func (x *ForeignMessage) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ForeignMessage)(x)
}

func (x *ForeignMessage) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ForeignMessage_messageType fastReflection_ForeignMessage_messageType
var _ protoreflect.MessageType = fastReflection_ForeignMessage_messageType{}

type fastReflection_ForeignMessage_messageType struct{}

func (x fastReflection_ForeignMessage_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ForeignMessage)(nil)
}
func (x fastReflection_ForeignMessage_messageType) New() protoreflect.Message {
	return new(fastReflection_ForeignMessage)
}
func (x fastReflection_ForeignMessage_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ForeignMessage
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ForeignMessage) Descriptor() protoreflect.MessageDescriptor {
	return md_ForeignMessage
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ForeignMessage) Type() protoreflect.MessageType {
	return _fastReflection_ForeignMessage_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ForeignMessage) New() protoreflect.Message {
	return new(fastReflection_ForeignMessage)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ForeignMessage) Interface() protoreflect.ProtoMessage {
	return (*ForeignMessage)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ForeignMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.C != int32(0) {
		value := protoreflect.ValueOfInt32(x.C)
		if !f(fd_ForeignMessage_c, value) {
			return
		}
	}
	if x.D != int32(0) {
		value := protoreflect.ValueOfInt32(x.D)
		if !f(fd_ForeignMessage_d, value) {
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
func (x *fastReflection_ForeignMessage) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		return x.C != int32(0)
	case "goproto.proto.test3.ForeignMessage.d":
		return x.D != int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ForeignMessage) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		x.C = int32(0)
	case "goproto.proto.test3.ForeignMessage.d":
		x.D = int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ForeignMessage) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		value := x.C
		return protoreflect.ValueOfInt32(value)
	case "goproto.proto.test3.ForeignMessage.d":
		value := x.D
		return protoreflect.ValueOfInt32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ForeignMessage) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		x.C = int32(value.Int())
	case "goproto.proto.test3.ForeignMessage.d":
		x.D = int32(value.Int())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ForeignMessage) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		panic(fmt.Errorf("field c of message goproto.proto.test3.ForeignMessage is not mutable"))
	case "goproto.proto.test3.ForeignMessage.d":
		panic(fmt.Errorf("field d of message goproto.proto.test3.ForeignMessage is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ForeignMessage) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.ForeignMessage.c":
		return protoreflect.ValueOfInt32(int32(0))
	case "goproto.proto.test3.ForeignMessage.d":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.ForeignMessage"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.ForeignMessage does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ForeignMessage) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.ForeignMessage", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ForeignMessage) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ForeignMessage) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ForeignMessage) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ForeignMessage) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ForeignMessage)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.C != 0 {
			n += 1 + runtime.Sov(uint64(x.C))
		}
		if x.D != 0 {
			n += 1 + runtime.Sov(uint64(x.D))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ForeignMessage)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.D != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.D))
			i--
			dAtA[i] = 0x10
		}
		if x.C != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.C))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ForeignMessage)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ForeignMessage: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ForeignMessage: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
				}
				x.C = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.C |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
				}
				x.D = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.D |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.7
// source: internal/testprotos/test3/test.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NestedEnum should eventually become nested when fast-reflection supports it.
type NestedEnum int32

const (
	NestedEnum_FOO NestedEnum = 0
	NestedEnum_BAR NestedEnum = 1
	NestedEnum_BAZ NestedEnum = 2
	NestedEnum_NEG NestedEnum = -1 // Intentionally negative.
)

// Enum value maps for NestedEnum.
var (
	NestedEnum_name = map[int32]string{
		0:  "FOO",
		1:  "BAR",
		2:  "BAZ",
		-1: "NEG",
	}
	NestedEnum_value = map[string]int32{
		"FOO": 0,
		"BAR": 1,
		"BAZ": 2,
		"NEG": -1,
	}
)

func (x NestedEnum) Enum() *NestedEnum {
	p := new(NestedEnum)
	*p = x
	return p
}

func (x NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_test3_test_proto_enumTypes[0].Descriptor()
}

func (NestedEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_test3_test_proto_enumTypes[0]
}

func (x NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NestedEnum.Descriptor instead.
func (NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_proto_rawDescGZIP(), []int{0}
}

type ForeignEnum int32

const (
	ForeignEnum_FOREIGN_ZERO ForeignEnum = 0
	ForeignEnum_FOREIGN_FOO  ForeignEnum = 4
	ForeignEnum_FOREIGN_BAR  ForeignEnum = 5
	ForeignEnum_FOREIGN_BAZ  ForeignEnum = 6
)

// Enum value maps for ForeignEnum.
var (
	ForeignEnum_name = map[int32]string{
		0: "FOREIGN_ZERO",
		4: "FOREIGN_FOO",
		5: "FOREIGN_BAR",
		6: "FOREIGN_BAZ",
	}
	ForeignEnum_value = map[string]int32{
		"FOREIGN_ZERO": 0,
		"FOREIGN_FOO":  4,
		"FOREIGN_BAR":  5,
		"FOREIGN_BAZ":  6,
	}
)

func (x ForeignEnum) Enum() *ForeignEnum {
	p := new(ForeignEnum)
	*p = x
	return p
}

func (x ForeignEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ForeignEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_test3_test_proto_enumTypes[1].Descriptor()
}

func (ForeignEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_test3_test_proto_enumTypes[1]
}

func (x ForeignEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ForeignEnum.Descriptor instead.
func (ForeignEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_proto_rawDescGZIP(), []int{1}
}

// NestedMessage should eventually become nested when fast-reflection supports it
type NestedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A           int32         `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	Corecursive *TestAllTypes `protobuf:"bytes,2,opt,name=corecursive,proto3" json:"corecursive,omitempty"`
}

func (x *NestedMessage) Reset() {
	*x = NestedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedMessage) ProtoMessage() {}

// Deprecated: Use NestedMessage.ProtoReflect.Descriptor instead.
func (*NestedMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_proto_rawDescGZIP(), []int{0}
}

func (x *NestedMessage) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *NestedMessage) GetCorecursive() *TestAllTypes {
	if x != nil {
		return x.Corecursive
	}
	return nil
}

type TestAllTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingularInt32          int32                     `protobuf:"varint,81,opt,name=singular_int32,json=singularInt32,proto3" json:"singular_int32,omitempty"`
	SingularInt64          int64                     `protobuf:"varint,82,opt,name=singular_int64,json=singularInt64,proto3" json:"singular_int64,omitempty"`
	SingularUint32         uint32                    `protobuf:"varint,83,opt,name=singular_uint32,json=singularUint32,proto3" json:"singular_uint32,omitempty"`
	SingularUint64         uint64                    `protobuf:"varint,84,opt,name=singular_uint64,json=singularUint64,proto3" json:"singular_uint64,omitempty"`
	SingularSint32         int32                     `protobuf:"zigzag32,85,opt,name=singular_sint32,json=singularSint32,proto3" json:"singular_sint32,omitempty"`
	SingularSint64         int64                     `protobuf:"zigzag64,86,opt,name=singular_sint64,json=singularSint64,proto3" json:"singular_sint64,omitempty"`
	SingularFixed32        uint32                    `protobuf:"fixed32,87,opt,name=singular_fixed32,json=singularFixed32,proto3" json:"singular_fixed32,omitempty"`
	SingularFixed64        uint64                    `protobuf:"fixed64,88,opt,name=singular_fixed64,json=singularFixed64,proto3" json:"singular_fixed64,omitempty"`
	SingularSfixed32       int32                     `protobuf:"fixed32,89,opt,name=singular_sfixed32,json=singularSfixed32,proto3" json:"singular_sfixed32,omitempty"`
	SingularSfixed64       int64                     `protobuf:"fixed64,90,opt,name=singular_sfixed64,json=singularSfixed64,proto3" json:"singular_sfixed64,omitempty"`
	SingularFloat          float32                   `protobuf:"fixed32,91,opt,name=singular_float,json=singularFloat,proto3" json:"singular_float,omitempty"`
	SingularDouble         float64                   `protobuf:"fixed64,92,opt,name=singular_double,json=singularDouble,proto3" json:"singular_double,omitempty"`
	SingularBool           bool                      `protobuf:"varint,93,opt,name=singular_bool,json=singularBool,proto3" json:"singular_bool,omitempty"`
	SingularString         string                    `protobuf:"bytes,94,opt,name=singular_string,json=singularString,proto3" json:"singular_string,omitempty"`
	SingularBytes          []byte                    `protobuf:"bytes,95,opt,name=singular_bytes,json=singularBytes,proto3" json:"singular_bytes,omitempty"`
	SingularNestedMessage  *NestedMessage            `protobuf:"bytes,98,opt,name=singular_nested_message,json=singularNestedMessage,proto3" json:"singular_nested_message,omitempty"`
	SingularForeignMessage *ForeignMessage           `protobuf:"bytes,99,opt,name=singular_foreign_message,json=singularForeignMessage,proto3" json:"singular_foreign_message,omitempty"`
	SingularImportMessage  *ImportMessage            `protobuf:"bytes,100,opt,name=singular_import_message,json=singularImportMessage,proto3" json:"singular_import_message,omitempty"`
	SingularNestedEnum     NestedEnum                `protobuf:"varint,101,opt,name=singular_nested_enum,json=singularNestedEnum,proto3,enum=goproto.proto.test3.NestedEnum" json:"singular_nested_enum,omitempty"`
	SingularForeignEnum    ForeignEnum               `protobuf:"varint,102,opt,name=singular_foreign_enum,json=singularForeignEnum,proto3,enum=goproto.proto.test3.ForeignEnum" json:"singular_foreign_enum,omitempty"`
	SingularImportEnum     ImportEnum                `protobuf:"varint,103,opt,name=singular_import_enum,json=singularImportEnum,proto3,enum=goproto.proto.test3.ImportEnum" json:"singular_import_enum,omitempty"`
	RepeatedInt32          []int32                   `protobuf:"varint,31,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedInt64          []int64                   `protobuf:"varint,32,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedUint32         []uint32                  `protobuf:"varint,33,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedUint64         []uint64                  `protobuf:"varint,34,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSint32         []int32                   `protobuf:"zigzag32,35,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedSint64         []int64                   `protobuf:"zigzag64,36,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedFixed32        []uint32                  `protobuf:"fixed32,37,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64        []uint64                  `protobuf:"fixed64,38,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32       []int32                   `protobuf:"fixed32,39,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64       []int64                   `protobuf:"fixed64,40,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedFloat          []float32                 `protobuf:"fixed32,41,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedDouble         []float64                 `protobuf:"fixed64,42,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedBool           []bool                    `protobuf:"varint,43,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedString         []string                  `protobuf:"bytes,44,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes          [][]byte                  `protobuf:"bytes,45,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	RepeatedNestedMessage  []*NestedMessage          `protobuf:"bytes,48,rep,name=repeated_nested_message,json=repeatedNestedMessage,proto3" json:"repeated_nested_message,omitempty"`
	RepeatedForeignMessage []*ForeignMessage         `protobuf:"bytes,49,rep,name=repeated_foreign_message,json=repeatedForeignMessage,proto3" json:"repeated_foreign_message,omitempty"`
	RepeatedImportmessage  []*ImportMessage          `protobuf:"bytes,50,rep,name=repeated_importmessage,json=repeatedImportmessage,proto3" json:"repeated_importmessage,omitempty"`
	RepeatedNestedEnum     []NestedEnum              `protobuf:"varint,51,rep,packed,name=repeated_nested_enum,json=repeatedNestedEnum,proto3,enum=goproto.proto.test3.NestedEnum" json:"repeated_nested_enum,omitempty"`
	RepeatedForeignEnum    []ForeignEnum             `protobuf:"varint,52,rep,packed,name=repeated_foreign_enum,json=repeatedForeignEnum,proto3,enum=goproto.proto.test3.ForeignEnum" json:"repeated_foreign_enum,omitempty"`
	RepeatedImportenum     []ImportEnum              `protobuf:"varint,53,rep,packed,name=repeated_importenum,json=repeatedImportenum,proto3,enum=goproto.proto.test3.ImportEnum" json:"repeated_importenum,omitempty"`
	MapInt32Int32          map[int32]int32           `protobuf:"bytes,56,rep,name=map_int32_int32,json=mapInt32Int32,proto3" json:"map_int32_int32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapInt64Int64          map[int64]int64           `protobuf:"bytes,57,rep,name=map_int64_int64,json=mapInt64Int64,proto3" json:"map_int64_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapUint32Uint32        map[uint32]uint32         `protobuf:"bytes,58,rep,name=map_uint32_uint32,json=mapUint32Uint32,proto3" json:"map_uint32_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapUint64Uint64        map[uint64]uint64         `protobuf:"bytes,59,rep,name=map_uint64_uint64,json=mapUint64Uint64,proto3" json:"map_uint64_uint64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapSint32Sint32        map[int32]int32           `protobuf:"bytes,60,rep,name=map_sint32_sint32,json=mapSint32Sint32,proto3" json:"map_sint32_sint32,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	MapSint64Sint64        map[int64]int64           `protobuf:"bytes,61,rep,name=map_sint64_sint64,json=mapSint64Sint64,proto3" json:"map_sint64_sint64,omitempty" protobuf_key:"zigzag64,1,opt,name=key,proto3" protobuf_val:"zigzag64,2,opt,name=value,proto3"`
	MapFixed32Fixed32      map[uint32]uint32         `protobuf:"bytes,62,rep,name=map_fixed32_fixed32,json=mapFixed32Fixed32,proto3" json:"map_fixed32_fixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	MapFixed64Fixed64      map[uint64]uint64         `protobuf:"bytes,63,rep,name=map_fixed64_fixed64,json=mapFixed64Fixed64,proto3" json:"map_fixed64_fixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	MapSfixed32Sfixed32    map[int32]int32           `protobuf:"bytes,64,rep,name=map_sfixed32_sfixed32,json=mapSfixed32Sfixed32,proto3" json:"map_sfixed32_sfixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	MapSfixed64Sfixed64    map[int64]int64           `protobuf:"bytes,65,rep,name=map_sfixed64_sfixed64,json=mapSfixed64Sfixed64,proto3" json:"map_sfixed64_sfixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	MapInt32Float          map[int32]float32         `protobuf:"bytes,66,rep,name=map_int32_float,json=mapInt32Float,proto3" json:"map_int32_float,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	MapInt32Double         map[int32]float64         `protobuf:"bytes,67,rep,name=map_int32_double,json=mapInt32Double,proto3" json:"map_int32_double,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	MapBoolBool            map[bool]bool             `protobuf:"bytes,68,rep,name=map_bool_bool,json=mapBoolBool,proto3" json:"map_bool_bool,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStringString        map[string]string         `protobuf:"bytes,69,rep,name=map_string_string,json=mapStringString,proto3" json:"map_string_string,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStringBytes         map[string][]byte         `protobuf:"bytes,70,rep,name=map_string_bytes,json=mapStringBytes,proto3" json:"map_string_bytes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStringNestedMessage map[string]*NestedMessage `protobuf:"bytes,71,rep,name=map_string_nested_message,json=mapStringNestedMessage,proto3" json:"map_string_nested_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStringNestedEnum    map[string]NestedEnum     `protobuf:"bytes,73,rep,name=map_string_nested_enum,json=mapStringNestedEnum,proto3" json:"map_string_nested_enum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=goproto.proto.test3.NestedEnum"`
	// Types that are assignable to OneofField:
	//	*TestAllTypes_OneofUint32
	//	*TestAllTypes_OneofNestedMessage
	//	*TestAllTypes_OneofString
	//	*TestAllTypes_OneofBytes
	//	*TestAllTypes_OneofBool
	//	*TestAllTypes_OneofUint64
	//	*TestAllTypes_OneofFloat
	//	*TestAllTypes_OneofDouble
	//	*TestAllTypes_OneofEnum
	OneofField isTestAllTypes_OneofField `protobuf_oneof:"oneof_field"`
}

func (x *TestAllTypes) Reset() {
	*x = TestAllTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAllTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAllTypes) ProtoMessage() {}

// Deprecated: Use TestAllTypes.ProtoReflect.Descriptor instead.
func (*TestAllTypes) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestAllTypes) GetSingularInt32() int32 {
	if x != nil {
		return x.SingularInt32
	}
	return 0
}

func (x *TestAllTypes) GetSingularInt64() int64 {
	if x != nil {
		return x.SingularInt64
	}
	return 0
}

func (x *TestAllTypes) GetSingularUint32() uint32 {
	if x != nil {
		return x.SingularUint32
	}
	return 0
}

func (x *TestAllTypes) GetSingularUint64() uint64 {
	if x != nil {
		return x.SingularUint64
	}
	return 0
}

func (x *TestAllTypes) GetSingularSint32() int32 {
	if x != nil {
		return x.SingularSint32
	}
	return 0
}

func (x *TestAllTypes) GetSingularSint64() int64 {
	if x != nil {
		return x.SingularSint64
	}
	return 0
}

func (x *TestAllTypes) GetSingularFixed32() uint32 {
	if x != nil {
		return x.SingularFixed32
	}
	return 0
}

func (x *TestAllTypes) GetSingularFixed64() uint64 {
	if x != nil {
		return x.SingularFixed64
	}
	return 0
}

func (x *TestAllTypes) GetSingularSfixed32() int32 {
	if x != nil {
		return x.SingularSfixed32
	}
	return 0
}

func (x *TestAllTypes) GetSingularSfixed64() int64 {
	if x != nil {
		return x.SingularSfixed64
	}
	return 0
}

func (x *TestAllTypes) GetSingularFloat() float32 {
	if x != nil {
		return x.SingularFloat
	}
	return 0
}

func (x *TestAllTypes) GetSingularDouble() float64 {
	if x != nil {
		return x.SingularDouble
	}
	return 0
}

func (x *TestAllTypes) GetSingularBool() bool {
	if x != nil {
		return x.SingularBool
	}
	return false
}

func (x *TestAllTypes) GetSingularString() string {
	if x != nil {
		return x.SingularString
	}
	return ""
}

func (x *TestAllTypes) GetSingularBytes() []byte {
	if x != nil {
		return x.SingularBytes
	}
	return nil
}

func (x *TestAllTypes) GetSingularNestedMessage() *NestedMessage {
	if x != nil {
		return x.SingularNestedMessage
	}
	return nil
}

func (x *TestAllTypes) GetSingularForeignMessage() *ForeignMessage {
	if x != nil {
		return x.SingularForeignMessage
	}
	return nil
}

func (x *TestAllTypes) GetSingularImportMessage() *ImportMessage {
	if x != nil {
		return x.SingularImportMessage
	}
	return nil
}

func (x *TestAllTypes) GetSingularNestedEnum() NestedEnum {
	if x != nil {
		return x.SingularNestedEnum
	}
	return NestedEnum_FOO
}

func (x *TestAllTypes) GetSingularForeignEnum() ForeignEnum {
	if x != nil {
		return x.SingularForeignEnum
	}
	return ForeignEnum_FOREIGN_ZERO
}

func (x *TestAllTypes) GetSingularImportEnum() ImportEnum {
	if x != nil {
		return x.SingularImportEnum
	}
	return ImportEnum_IMPORT_ZERO
}

func (x *TestAllTypes) GetRepeatedInt32() []int32 {
	if x != nil {
		return x.RepeatedInt32
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedInt64() []int64 {
	if x != nil {
		return x.RepeatedInt64
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedUint32() []uint32 {
	if x != nil {
		return x.RepeatedUint32
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedUint64() []uint64 {
	if x != nil {
		return x.RepeatedUint64
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedSint32() []int32 {
	if x != nil {
		return x.RepeatedSint32
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedSint64() []int64 {
	if x != nil {
		return x.RepeatedSint64
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedFixed32() []uint32 {
	if x != nil {
		return x.RepeatedFixed32
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedFixed64() []uint64 {
	if x != nil {
		return x.RepeatedFixed64
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedSfixed32() []int32 {
	if x != nil {
		return x.RepeatedSfixed32
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedSfixed64() []int64 {
	if x != nil {
		return x.RepeatedSfixed64
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedFloat() []float32 {
	if x != nil {
		return x.RepeatedFloat
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedDouble() []float64 {
	if x != nil {
		return x.RepeatedDouble
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedBool() []bool {
	if x != nil {
		return x.RepeatedBool
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedString() []string {
	if x != nil {
		return x.RepeatedString
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedBytes() [][]byte {
	if x != nil {
		return x.RepeatedBytes
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedNestedMessage() []*NestedMessage {
	if x != nil {
		return x.RepeatedNestedMessage
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedForeignMessage() []*ForeignMessage {
	if x != nil {
		return x.RepeatedForeignMessage
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedImportmessage() []*ImportMessage {
	if x != nil {
		return x.RepeatedImportmessage
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedNestedEnum() []NestedEnum {
	if x != nil {
		return x.RepeatedNestedEnum
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedForeignEnum() []ForeignEnum {
	if x != nil {
		return x.RepeatedForeignEnum
	}
	return nil
}

func (x *TestAllTypes) GetRepeatedImportenum() []ImportEnum {
	if x != nil {
		return x.RepeatedImportenum
	}
	return nil
}

func (x *TestAllTypes) GetMapInt32Int32() map[int32]int32 {
	if x != nil {
		return x.MapInt32Int32
	}
	return nil
}

func (x *TestAllTypes) GetMapInt64Int64() map[int64]int64 {
	if x != nil {
		return x.MapInt64Int64
	}
	return nil
}

func (x *TestAllTypes) GetMapUint32Uint32() map[uint32]uint32 {
	if x != nil {
		return x.MapUint32Uint32
	}
	return nil
}

func (x *TestAllTypes) GetMapUint64Uint64() map[uint64]uint64 {
	if x != nil {
		return x.MapUint64Uint64
	}
	return nil
}

func (x *TestAllTypes) GetMapSint32Sint32() map[int32]int32 {
	if x != nil {
		return x.MapSint32Sint32
	}
	return nil
}

func (x *TestAllTypes) GetMapSint64Sint64() map[int64]int64 {
	if x != nil {
		return x.MapSint64Sint64
	}
	return nil
}

func (x *TestAllTypes) GetMapFixed32Fixed32() map[uint32]uint32 {
	if x != nil {
		return x.MapFixed32Fixed32
	}
	return nil
}

func (x *TestAllTypes) GetMapFixed64Fixed64() map[uint64]uint64 {
	if x != nil {
		return x.MapFixed64Fixed64
	}
	return nil
}

func (x *TestAllTypes) GetMapSfixed32Sfixed32() map[int32]int32 {
	if x != nil {
		return x.MapSfixed32Sfixed32
	}
	return nil
}

func (x *TestAllTypes) GetMapSfixed64Sfixed64() map[int64]int64 {
	if x != nil {
		return x.MapSfixed64Sfixed64
	}
	return nil
}

func (x *TestAllTypes) GetMapInt32Float() map[int32]float32 {
	if x != nil {
		return x.MapInt32Float
	}
	return nil
}

func (x *TestAllTypes) GetMapInt32Double() map[int32]float64 {
	if x != nil {
		return x.MapInt32Double
	}
	return nil
}

func (x *TestAllTypes) GetMapBoolBool() map[bool]bool {
	if x != nil {
		return x.MapBoolBool
	}
	return nil
}

func (x *TestAllTypes) GetMapStringString() map[string]string {
	if x != nil {
		return x.MapStringString
	}
	return nil
}

func (x *TestAllTypes) GetMapStringBytes() map[string][]byte {
	if x != nil {
		return x.MapStringBytes
	}
	return nil
}

func (x *TestAllTypes) GetMapStringNestedMessage() map[string]*NestedMessage {
	if x != nil {
		return x.MapStringNestedMessage
	}
	return nil
}

func (x *TestAllTypes) GetMapStringNestedEnum() map[string]NestedEnum {
	if x != nil {
		return x.MapStringNestedEnum
	}
	return nil
}

func (x *TestAllTypes) GetOneofField() isTestAllTypes_OneofField {
	if x != nil {
		return x.OneofField
	}
	return nil
}

func (x *TestAllTypes) GetOneofUint32() uint32 {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofUint32); ok {
		return x.OneofUint32
	}
	return 0
}

func (x *TestAllTypes) GetOneofNestedMessage() *NestedMessage {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofNestedMessage); ok {
		return x.OneofNestedMessage
	}
	return nil
}

func (x *TestAllTypes) GetOneofString() string {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (x *TestAllTypes) GetOneofBytes() []byte {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofBytes); ok {
		return x.OneofBytes
	}
	return nil
}

func (x *TestAllTypes) GetOneofBool() bool {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofBool); ok {
		return x.OneofBool
	}
	return false
}

func (x *TestAllTypes) GetOneofUint64() uint64 {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofUint64); ok {
		return x.OneofUint64
	}
	return 0
}

func (x *TestAllTypes) GetOneofFloat() float32 {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofFloat); ok {
		return x.OneofFloat
	}
	return 0
}

func (x *TestAllTypes) GetOneofDouble() float64 {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofDouble); ok {
		return x.OneofDouble
	}
	return 0
}

func (x *TestAllTypes) GetOneofEnum() NestedEnum {
	if x, ok := x.GetOneofField().(*TestAllTypes_OneofEnum); ok {
		return x.OneofEnum
	}
	return NestedEnum_FOO
}

type isTestAllTypes_OneofField interface {
	isTestAllTypes_OneofField()
}

type TestAllTypes_OneofUint32 struct {
	OneofUint32 uint32 `protobuf:"varint,111,opt,name=oneof_uint32,json=oneofUint32,proto3,oneof"`
}

type TestAllTypes_OneofNestedMessage struct {
	OneofNestedMessage *NestedMessage `protobuf:"bytes,112,opt,name=oneof_nested_message,json=oneofNestedMessage,proto3,oneof"`
}

type TestAllTypes_OneofString struct {
	OneofString string `protobuf:"bytes,113,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type TestAllTypes_OneofBytes struct {
	OneofBytes []byte `protobuf:"bytes,114,opt,name=oneof_bytes,json=oneofBytes,proto3,oneof"`
}

type TestAllTypes_OneofBool struct {
	OneofBool bool `protobuf:"varint,115,opt,name=oneof_bool,json=oneofBool,proto3,oneof"`
}

type TestAllTypes_OneofUint64 struct {
	OneofUint64 uint64 `protobuf:"varint,116,opt,name=oneof_uint64,json=oneofUint64,proto3,oneof"`
}

type TestAllTypes_OneofFloat struct {
	OneofFloat float32 `protobuf:"fixed32,117,opt,name=oneof_float,json=oneofFloat,proto3,oneof"`
}

type TestAllTypes_OneofDouble struct {
	OneofDouble float64 `protobuf:"fixed64,118,opt,name=oneof_double,json=oneofDouble,proto3,oneof"`
}

type TestAllTypes_OneofEnum struct {
	OneofEnum NestedEnum `protobuf:"varint,119,opt,name=oneof_enum,json=oneofEnum,proto3,enum=goproto.proto.test3.NestedEnum,oneof"`
}

func (*TestAllTypes_OneofUint32) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofNestedMessage) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofString) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofBytes) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofBool) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofUint64) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofFloat) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofDouble) isTestAllTypes_OneofField() {}

func (*TestAllTypes_OneofEnum) isTestAllTypes_OneofField() {}

type ForeignMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C int32 `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	D int32 `protobuf:"varint,2,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *ForeignMessage) Reset() {
	*x = ForeignMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForeignMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForeignMessage) ProtoMessage() {}

// Deprecated: Use ForeignMessage.ProtoReflect.Descriptor instead.
func (*ForeignMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_proto_rawDescGZIP(), []int{2}
}

func (x *ForeignMessage) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *ForeignMessage) GetD() int32 {
	if x != nil {
		return x.D
	}
	return 0
}

var File_internal_testprotos_test3_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_test3_test_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x1a, 0x2b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0d, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x0b, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0xbf, 0x2c, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x51, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x52, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x69,
	0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x53,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x54, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x55, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x56, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x57, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x69,
	0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x58,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x59, 0x20, 0x01, 0x28, 0x0f,
	0x52, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x10, 0x73,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x62, 0x6f, 0x6f, 0x6c,
	0x18, 0x5d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x5e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x5f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x17, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x62, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x15, 0x73, 0x69, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x5d, 0x0a, 0x18, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x16, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x5a, 0x0a, 0x17, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x15, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x73,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x12, 0x73, 0x69, 0x6e, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x54,
	0x0a, 0x15, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x33, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x13, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x51, 0x0a, 0x14, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x12, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x18, 0x20, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x22, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x23, 0x20, 0x03, 0x28, 0x11,
	0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x24, 0x20, 0x03, 0x28, 0x12, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x25, 0x20,
	0x03, 0x28, 0x07, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x26, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0f,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x18, 0x27, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x2b, 0x0a, 0x11,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x18, 0x28, 0x20, 0x03, 0x28, 0x10, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x29, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x2b, 0x20, 0x03, 0x28, 0x08,
	0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x2c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x2d, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x5a,
	0x0a, 0x17, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x30, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a, 0x18, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x31, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x33, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x16, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x15, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x33, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x54, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x34, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x46, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x50, 0x0a,
	0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x75, 0x6d, 0x18, 0x35, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x12, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x75, 0x6d, 0x12,
	0x5c, 0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x18, 0x38, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x5c, 0x0a,
	0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x18, 0x39, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61,
	0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x62, 0x0a, 0x11, 0x6d,
	0x61, 0x70, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x3a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f,
	0x6d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x62, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x3b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x55, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x62, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x53, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x53, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x62, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x73,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x3d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x53,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x53,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x68, 0x0a, 0x13, 0x6d,
	0x61, 0x70, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x18, 0x3e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x11, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x68, 0x0a, 0x13, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x3f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x6d, 0x61,
	0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x6e, 0x0a, 0x15, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x40, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x53, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x6d, 0x61, 0x70, 0x53,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x6e, 0x0a, 0x15, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x41, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x53, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x6d, 0x61, 0x70, 0x53,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x5c, 0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x42, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x5f, 0x0a,
	0x10, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x18, 0x43, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x56,
	0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18,
	0x44, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c,
	0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x42, 0x6f,
	0x6f, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x62, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x45, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x5f, 0x0a, 0x10, 0x6d, 0x61,
	0x70, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x46,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6d, 0x61, 0x70,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x78, 0x0a, 0x19, 0x6d,
	0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x47, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x6d,
	0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6f, 0x0a, 0x16, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18,
	0x49, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x13, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x56, 0x0a, 0x14, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x12, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x73, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0c,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x74, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x75, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x18, 0x76, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x77, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00,
	0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x40, 0x0a, 0x12, 0x4d,
	0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a,
	0x12, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x42, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x53, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x4d,
	0x61, 0x70, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x44, 0x0a, 0x16, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x16, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x4d,
	0x61, 0x70, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x4d, 0x61, 0x70, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x10, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x4d,
	0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a,
	0x13, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6d, 0x0a, 0x1b, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x67, 0x0a, 0x18, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x2c,
	0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x63, 0x12, 0x0c,
	0x0a, 0x01, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x64, 0x2a, 0x39, 0x0a, 0x0a,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f,
	0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x41, 0x5a, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x45, 0x47, 0x10, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x2a, 0x52, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4f, 0x52, 0x45, 0x49, 0x47,
	0x4e, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x45,
	0x49, 0x47, 0x4e, 0x5f, 0x46, 0x4f, 0x4f, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52,
	0x45, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f,
	0x52, 0x45, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x41, 0x5a, 0x10, 0x06, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_testprotos_test3_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_test3_test_proto_rawDescData = file_internal_testprotos_test3_test_proto_rawDesc
)

func file_internal_testprotos_test3_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_test3_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_test3_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_test3_test_proto_rawDescData)
	})
	return file_internal_testprotos_test3_test_proto_rawDescData
}

var file_internal_testprotos_test3_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_testprotos_test3_test_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_internal_testprotos_test3_test_proto_goTypes = []interface{}{
	(NestedEnum)(0),        // 0: goproto.proto.test3.NestedEnum
	(ForeignEnum)(0),       // 1: goproto.proto.test3.ForeignEnum
	(*NestedMessage)(nil),  // 2: goproto.proto.test3.NestedMessage
	(*TestAllTypes)(nil),   // 3: goproto.proto.test3.TestAllTypes
	(*ForeignMessage)(nil), // 4: goproto.proto.test3.ForeignMessage
	nil,                    // 5: goproto.proto.test3.TestAllTypes.MapInt32Int32Entry
	nil,                    // 6: goproto.proto.test3.TestAllTypes.MapInt64Int64Entry
	nil,                    // 7: goproto.proto.test3.TestAllTypes.MapUint32Uint32Entry
	nil,                    // 8: goproto.proto.test3.TestAllTypes.MapUint64Uint64Entry
	nil,                    // 9: goproto.proto.test3.TestAllTypes.MapSint32Sint32Entry
	nil,                    // 10: goproto.proto.test3.TestAllTypes.MapSint64Sint64Entry
	nil,                    // 11: goproto.proto.test3.TestAllTypes.MapFixed32Fixed32Entry
	nil,                    // 12: goproto.proto.test3.TestAllTypes.MapFixed64Fixed64Entry
	nil,                    // 13: goproto.proto.test3.TestAllTypes.MapSfixed32Sfixed32Entry
	nil,                    // 14: goproto.proto.test3.TestAllTypes.MapSfixed64Sfixed64Entry
	nil,                    // 15: goproto.proto.test3.TestAllTypes.MapInt32FloatEntry
	nil,                    // 16: goproto.proto.test3.TestAllTypes.MapInt32DoubleEntry
	nil,                    // 17: goproto.proto.test3.TestAllTypes.MapBoolBoolEntry
	nil,                    // 18: goproto.proto.test3.TestAllTypes.MapStringStringEntry
	nil,                    // 19: goproto.proto.test3.TestAllTypes.MapStringBytesEntry
	nil,                    // 20: goproto.proto.test3.TestAllTypes.MapStringNestedMessageEntry
	nil,                    // 21: goproto.proto.test3.TestAllTypes.MapStringNestedEnumEntry
	(*ImportMessage)(nil),  // 22: goproto.proto.test3.ImportMessage
	(ImportEnum)(0),        // 23: goproto.proto.test3.ImportEnum
}
var file_internal_testprotos_test3_test_proto_depIdxs = []int32{
	3,  // 0: goproto.proto.test3.NestedMessage.corecursive:type_name -> goproto.proto.test3.TestAllTypes
	2,  // 1: goproto.proto.test3.TestAllTypes.singular_nested_message:type_name -> goproto.proto.test3.NestedMessage
	4,  // 2: goproto.proto.test3.TestAllTypes.singular_foreign_message:type_name -> goproto.proto.test3.ForeignMessage
	22, // 3: goproto.proto.test3.TestAllTypes.singular_import_message:type_name -> goproto.proto.test3.ImportMessage
	0,  // 4: goproto.proto.test3.TestAllTypes.singular_nested_enum:type_name -> goproto.proto.test3.NestedEnum
	1,  // 5: goproto.proto.test3.TestAllTypes.singular_foreign_enum:type_name -> goproto.proto.test3.ForeignEnum
	23, // 6: goproto.proto.test3.TestAllTypes.singular_import_enum:type_name -> goproto.proto.test3.ImportEnum
	2,  // 7: goproto.proto.test3.TestAllTypes.repeated_nested_message:type_name -> goproto.proto.test3.NestedMessage
	4,  // 8: goproto.proto.test3.TestAllTypes.repeated_foreign_message:type_name -> goproto.proto.test3.ForeignMessage
	22, // 9: goproto.proto.test3.TestAllTypes.repeated_importmessage:type_name -> goproto.proto.test3.ImportMessage
	0,  // 10: goproto.proto.test3.TestAllTypes.repeated_nested_enum:type_name -> goproto.proto.test3.NestedEnum
	1,  // 11: goproto.proto.test3.TestAllTypes.repeated_foreign_enum:type_name -> goproto.proto.test3.ForeignEnum
	23, // 12: goproto.proto.test3.TestAllTypes.repeated_importenum:type_name -> goproto.proto.test3.ImportEnum
	5,  // 13: goproto.proto.test3.TestAllTypes.map_int32_int32:type_name -> goproto.proto.test3.TestAllTypes.MapInt32Int32Entry
	6,  // 14: goproto.proto.test3.TestAllTypes.map_int64_int64:type_name -> goproto.proto.test3.TestAllTypes.MapInt64Int64Entry
	7,  // 15: goproto.proto.test3.TestAllTypes.map_uint32_uint32:type_name -> goproto.proto.test3.TestAllTypes.MapUint32Uint32Entry
	8,  // 16: goproto.proto.test3.TestAllTypes.map_uint64_uint64:type_name -> goproto.proto.test3.TestAllTypes.MapUint64Uint64Entry
	9,  // 17: goproto.proto.test3.TestAllTypes.map_sint32_sint32:type_name -> goproto.proto.test3.TestAllTypes.MapSint32Sint32Entry
	10, // 18: goproto.proto.test3.TestAllTypes.map_sint64_sint64:type_name -> goproto.proto.test3.TestAllTypes.MapSint64Sint64Entry
	11, // 19: goproto.proto.test3.TestAllTypes.map_fixed32_fixed32:type_name -> goproto.proto.test3.TestAllTypes.MapFixed32Fixed32Entry
	12, // 20: goproto.proto.test3.TestAllTypes.map_fixed64_fixed64:type_name -> goproto.proto.test3.TestAllTypes.MapFixed64Fixed64Entry
	13, // 21: goproto.proto.test3.TestAllTypes.map_sfixed32_sfixed32:type_name -> goproto.proto.test3.TestAllTypes.MapSfixed32Sfixed32Entry
	14, // 22: goproto.proto.test3.TestAllTypes.map_sfixed64_sfixed64:type_name -> goproto.proto.test3.TestAllTypes.MapSfixed64Sfixed64Entry
	15, // 23: goproto.proto.test3.TestAllTypes.map_int32_float:type_name -> goproto.proto.test3.TestAllTypes.MapInt32FloatEntry
	16, // 24: goproto.proto.test3.TestAllTypes.map_int32_double:type_name -> goproto.proto.test3.TestAllTypes.MapInt32DoubleEntry
	17, // 25: goproto.proto.test3.TestAllTypes.map_bool_bool:type_name -> goproto.proto.test3.TestAllTypes.MapBoolBoolEntry
	18, // 26: goproto.proto.test3.TestAllTypes.map_string_string:type_name -> goproto.proto.test3.TestAllTypes.MapStringStringEntry
	19, // 27: goproto.proto.test3.TestAllTypes.map_string_bytes:type_name -> goproto.proto.test3.TestAllTypes.MapStringBytesEntry
	20, // 28: goproto.proto.test3.TestAllTypes.map_string_nested_message:type_name -> goproto.proto.test3.TestAllTypes.MapStringNestedMessageEntry
	21, // 29: goproto.proto.test3.TestAllTypes.map_string_nested_enum:type_name -> goproto.proto.test3.TestAllTypes.MapStringNestedEnumEntry
	2,  // 30: goproto.proto.test3.TestAllTypes.oneof_nested_message:type_name -> goproto.proto.test3.NestedMessage
	0,  // 31: goproto.proto.test3.TestAllTypes.oneof_enum:type_name -> goproto.proto.test3.NestedEnum
	2,  // 32: goproto.proto.test3.TestAllTypes.MapStringNestedMessageEntry.value:type_name -> goproto.proto.test3.NestedMessage
	0,  // 33: goproto.proto.test3.TestAllTypes.MapStringNestedEnumEntry.value:type_name -> goproto.proto.test3.NestedEnum
	34, // [34:34] is the sub-list for method output_type
	34, // [34:34] is the sub-list for method input_type
	34, // [34:34] is the sub-list for extension type_name
	34, // [34:34] is the sub-list for extension extendee
	0,  // [0:34] is the sub-list for field type_name
}

func init() { file_internal_testprotos_test3_test_proto_init() }
func file_internal_testprotos_test3_test_proto_init() {
	if File_internal_testprotos_test3_test_proto != nil {
		return
	}
	file_internal_testprotos_test3_test_import_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_test3_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedMessage); i {
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
		file_internal_testprotos_test3_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAllTypes); i {
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
		file_internal_testprotos_test3_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForeignMessage); i {
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
	file_internal_testprotos_test3_test_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TestAllTypes_OneofUint32)(nil),
		(*TestAllTypes_OneofNestedMessage)(nil),
		(*TestAllTypes_OneofString)(nil),
		(*TestAllTypes_OneofBytes)(nil),
		(*TestAllTypes_OneofBool)(nil),
		(*TestAllTypes_OneofUint64)(nil),
		(*TestAllTypes_OneofFloat)(nil),
		(*TestAllTypes_OneofDouble)(nil),
		(*TestAllTypes_OneofEnum)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_test3_test_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_test3_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_test3_test_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_test3_test_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_test3_test_proto_msgTypes,
	}.Build()
	File_internal_testprotos_test3_test_proto = out.File
	file_internal_testprotos_test3_test_proto_rawDesc = nil
	file_internal_testprotos_test3_test_proto_goTypes = nil
	file_internal_testprotos_test3_test_proto_depIdxs = nil
}
