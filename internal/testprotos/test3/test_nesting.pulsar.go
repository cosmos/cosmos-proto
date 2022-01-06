package test3

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MultiLayeredNesting         protoreflect.MessageDescriptor
	fd_MultiLayeredNesting_nested1 protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_nesting_proto_init()
	md_MultiLayeredNesting = File_internal_testprotos_test3_test_nesting_proto.Messages().ByName("MultiLayeredNesting")
	fd_MultiLayeredNesting_nested1 = md_MultiLayeredNesting.Fields().ByName("nested1")
}

var _ protoreflect.Message = (*fastReflection_MultiLayeredNesting)(nil)

type fastReflection_MultiLayeredNesting MultiLayeredNesting

func (x *MultiLayeredNesting) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting)(x)
}

func (x *MultiLayeredNesting) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MultiLayeredNesting_messageType fastReflection_MultiLayeredNesting_messageType
var _ protoreflect.MessageType = fastReflection_MultiLayeredNesting_messageType{}

type fastReflection_MultiLayeredNesting_messageType struct{}

func (x fastReflection_MultiLayeredNesting_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting)(nil)
}
func (x fastReflection_MultiLayeredNesting_messageType) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting)
}
func (x fastReflection_MultiLayeredNesting_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MultiLayeredNesting) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MultiLayeredNesting) Type() protoreflect.MessageType {
	return _fastReflection_MultiLayeredNesting_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MultiLayeredNesting) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MultiLayeredNesting) Interface() protoreflect.ProtoMessage {
	return (*MultiLayeredNesting)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MultiLayeredNesting) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Nested1 != nil {
		value := protoreflect.ValueOfMessage(x.Nested1.ProtoReflect())
		if !f(fd_MultiLayeredNesting_nested1, value) {
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
func (x *fastReflection_MultiLayeredNesting) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		return x.Nested1 != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		x.Nested1 = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MultiLayeredNesting) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		value := x.Nested1
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MultiLayeredNesting) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		x.Nested1 = value.Message().Interface().(*MultiLayeredNesting_Nested1)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MultiLayeredNesting) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		if x.Nested1 == nil {
			x.Nested1 = new(MultiLayeredNesting_Nested1)
		}
		return protoreflect.ValueOfMessage(x.Nested1.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MultiLayeredNesting) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.nested1":
		m := new(MultiLayeredNesting_Nested1)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MultiLayeredNesting) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.MultiLayeredNesting", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MultiLayeredNesting) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MultiLayeredNesting) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MultiLayeredNesting) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MultiLayeredNesting)
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
		if x.Nested1 != nil {
			l = options.Size(x.Nested1)
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
		x := input.Message.Interface().(*MultiLayeredNesting)
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
		if x.Nested1 != nil {
			encoded, err := options.Marshal(x.Nested1)
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
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*MultiLayeredNesting)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Nested1", wireType)
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
				if x.Nested1 == nil {
					x.Nested1 = &MultiLayeredNesting_Nested1{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Nested1); err != nil {
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

var (
	md_MultiLayeredNesting_Nested1 protoreflect.MessageDescriptor
)

func init() {
	file_internal_testprotos_test3_test_nesting_proto_init()
	md_MultiLayeredNesting_Nested1 = File_internal_testprotos_test3_test_nesting_proto.Messages().ByName("MultiLayeredNesting").Messages().ByName("Nested1")
}

var _ protoreflect.Message = (*fastReflection_MultiLayeredNesting_Nested1)(nil)

type fastReflection_MultiLayeredNesting_Nested1 MultiLayeredNesting_Nested1

func (x *MultiLayeredNesting_Nested1) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1)(x)
}

func (x *MultiLayeredNesting_Nested1) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MultiLayeredNesting_Nested1_messageType fastReflection_MultiLayeredNesting_Nested1_messageType
var _ protoreflect.MessageType = fastReflection_MultiLayeredNesting_Nested1_messageType{}

type fastReflection_MultiLayeredNesting_Nested1_messageType struct{}

func (x fastReflection_MultiLayeredNesting_Nested1_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1)(nil)
}
func (x fastReflection_MultiLayeredNesting_Nested1_messageType) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1)
}
func (x fastReflection_MultiLayeredNesting_Nested1_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MultiLayeredNesting_Nested1) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MultiLayeredNesting_Nested1) Type() protoreflect.MessageType {
	return _fastReflection_MultiLayeredNesting_Nested1_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MultiLayeredNesting_Nested1) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MultiLayeredNesting_Nested1) Interface() protoreflect.ProtoMessage {
	return (*MultiLayeredNesting_Nested1)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MultiLayeredNesting_Nested1) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MultiLayeredNesting_Nested1) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MultiLayeredNesting_Nested1) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1 does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MultiLayeredNesting_Nested1) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.MultiLayeredNesting.Nested1", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MultiLayeredNesting_Nested1) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MultiLayeredNesting_Nested1) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1)
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
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_MultiLayeredNesting_Nested1_Nested2          protoreflect.MessageDescriptor
	fd_MultiLayeredNesting_Nested1_Nested2_nested_3 protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_nesting_proto_init()
	md_MultiLayeredNesting_Nested1_Nested2 = File_internal_testprotos_test3_test_nesting_proto.Messages().ByName("MultiLayeredNesting").Messages().ByName("Nested1").Messages().ByName("Nested2")
	fd_MultiLayeredNesting_Nested1_Nested2_nested_3 = md_MultiLayeredNesting_Nested1_Nested2.Fields().ByName("nested_3")
}

var _ protoreflect.Message = (*fastReflection_MultiLayeredNesting_Nested1_Nested2)(nil)

type fastReflection_MultiLayeredNesting_Nested1_Nested2 MultiLayeredNesting_Nested1_Nested2

func (x *MultiLayeredNesting_Nested1_Nested2) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1_Nested2)(x)
}

func (x *MultiLayeredNesting_Nested1_Nested2) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType
var _ protoreflect.MessageType = fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType{}

type fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType struct{}

func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1_Nested2)(nil)
}
func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1_Nested2)
}
func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1_Nested2
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1_Nested2
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Type() protoreflect.MessageType {
	return _fastReflection_MultiLayeredNesting_Nested1_Nested2_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1_Nested2)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Interface() protoreflect.ProtoMessage {
	return (*MultiLayeredNesting_Nested1_Nested2)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Nested_3 != nil {
		value := protoreflect.ValueOfMessage(x.Nested_3.ProtoReflect())
		if !f(fd_MultiLayeredNesting_Nested1_Nested2_nested_3, value) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		return x.Nested_3 != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		x.Nested_3 = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		value := x.Nested_3
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		x.Nested_3 = value.Message().Interface().(*MultiLayeredNesting_Nested1_Nested2_Nested3)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		if x.Nested_3 == nil {
			x.Nested_3 = new(MultiLayeredNesting_Nested1_Nested2_Nested3)
		}
		return protoreflect.ValueOfMessage(x.Nested_3.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3":
		m := new(MultiLayeredNesting_Nested1_Nested2_Nested3)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2 does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2)
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
		if x.Nested_3 != nil {
			l = options.Size(x.Nested_3)
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
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2)
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
		if x.Nested_3 != nil {
			encoded, err := options.Marshal(x.Nested_3)
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
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1_Nested2: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1_Nested2: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Nested_3", wireType)
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
				if x.Nested_3 == nil {
					x.Nested_3 = &MultiLayeredNesting_Nested1_Nested2_Nested3{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Nested_3); err != nil {
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

var (
	md_MultiLayeredNesting_Nested1_Nested2_Nested3                 protoreflect.MessageDescriptor
	fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_string protoreflect.FieldDescriptor
	fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_int32  protoreflect.FieldDescriptor
)

func init() {
	file_internal_testprotos_test3_test_nesting_proto_init()
	md_MultiLayeredNesting_Nested1_Nested2_Nested3 = File_internal_testprotos_test3_test_nesting_proto.Messages().ByName("MultiLayeredNesting").Messages().ByName("Nested1").Messages().ByName("Nested2").Messages().ByName("Nested3")
	fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_string = md_MultiLayeredNesting_Nested1_Nested2_Nested3.Fields().ByName("nested_3_string")
	fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_int32 = md_MultiLayeredNesting_Nested1_Nested2_Nested3.Fields().ByName("nested_3_int32")
}

var _ protoreflect.Message = (*fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3)(nil)

type fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3 MultiLayeredNesting_Nested1_Nested2_Nested3

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3)(x)
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) slowProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType
var _ protoreflect.MessageType = fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType{}

type fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType struct{}

func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3)(nil)
}
func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3)
}
func (x fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1_Nested2_Nested3
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiLayeredNesting_Nested1_Nested2_Nested3
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Type() protoreflect.MessageType {
	return _fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) New() protoreflect.Message {
	return new(fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Interface() protoreflect.ProtoMessage {
	return (*MultiLayeredNesting_Nested1_Nested2_Nested3)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Nested3Oneof != nil {
		switch o := x.Nested3Oneof.(type) {
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String:
			v := o.Nested_3String
			value := protoreflect.ValueOfString(v)
			if !f(fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_string, value) {
				return
			}
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32:
			v := o.Nested_3Int32
			value := protoreflect.ValueOfInt32(v)
			if !f(fd_MultiLayeredNesting_Nested1_Nested2_Nested3_nested_3_int32, value) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		if x.Nested3Oneof == nil {
			return false
		} else if _, ok := x.Nested3Oneof.(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String); ok {
			return true
		} else {
			return false
		}
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		if x.Nested3Oneof == nil {
			return false
		} else if _, ok := x.Nested3Oneof.(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		x.Nested3Oneof = nil
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		x.Nested3Oneof = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		if x.Nested3Oneof == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Nested3Oneof.(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String); ok {
			return protoreflect.ValueOfString(v.Nested_3String)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		if x.Nested3Oneof == nil {
			return protoreflect.ValueOfInt32(int32(0))
		} else if v, ok := x.Nested3Oneof.(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32); ok {
			return protoreflect.ValueOfInt32(v.Nested_3Int32)
		} else {
			return protoreflect.ValueOfInt32(int32(0))
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		cv := value.Interface().(string)
		x.Nested3Oneof = &MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String{Nested_3String: cv}
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		cv := int32(value.Int())
		x.Nested3Oneof = &MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32{Nested_3Int32: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		panic(fmt.Errorf("field nested_3_string of message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 is not mutable"))
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		panic(fmt.Errorf("field nested_3_int32 of message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_string":
		return protoreflect.ValueOfString("")
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested_3_int32":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3"))
		}
		panic(fmt.Errorf("message goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3 does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3.nested3_oneof":
		if x.Nested3Oneof == nil {
			return nil
		}
		switch x.Nested3Oneof.(type) {
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String:
			return x.Descriptor().Fields().ByName("nested_3_string")
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32:
			return x.Descriptor().Fields().ByName("nested_3_int32")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MultiLayeredNesting_Nested1_Nested2_Nested3) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2_Nested3)
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
		switch x := x.Nested3Oneof.(type) {
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String:
			if x == nil {
				break
			}
			l = len(x.Nested_3String)
			n += 1 + l + runtime.Sov(uint64(l))
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32:
			if x == nil {
				break
			}
			n += 1 + runtime.Sov(uint64(x.Nested_3Int32))
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
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2_Nested3)
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
		switch x := x.Nested3Oneof.(type) {
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String:
			i -= len(x.Nested_3String)
			copy(dAtA[i:], x.Nested_3String)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Nested_3String)))
			i--
			dAtA[i] = 0xa
		case *MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32:
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Nested_3Int32))
			i--
			dAtA[i] = 0x10
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
		x := input.Message.Interface().(*MultiLayeredNesting_Nested1_Nested2_Nested3)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1_Nested2_Nested3: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiLayeredNesting_Nested1_Nested2_Nested3: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Nested_3String", wireType)
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
				x.Nested3Oneof = &MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Nested_3Int32", wireType)
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
				x.Nested3Oneof = &MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32{v}
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.18.1
// source: internal/testprotos/test3/test_nesting.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MultiLayeredNesting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nested1 *MultiLayeredNesting_Nested1 `protobuf:"bytes,1,opt,name=nested1,proto3" json:"nested1,omitempty"`
}

func (x *MultiLayeredNesting) Reset() {
	*x = MultiLayeredNesting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLayeredNesting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLayeredNesting) ProtoMessage() {}

// Deprecated: Use MultiLayeredNesting.ProtoReflect.Descriptor instead.
func (*MultiLayeredNesting) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_nesting_proto_rawDescGZIP(), []int{0}
}

func (x *MultiLayeredNesting) GetNested1() *MultiLayeredNesting_Nested1 {
	if x != nil {
		return x.Nested1
	}
	return nil
}

type MultiLayeredNesting_Nested1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MultiLayeredNesting_Nested1) Reset() {
	*x = MultiLayeredNesting_Nested1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLayeredNesting_Nested1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLayeredNesting_Nested1) ProtoMessage() {}

// Deprecated: Use MultiLayeredNesting_Nested1.ProtoReflect.Descriptor instead.
func (*MultiLayeredNesting_Nested1) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_nesting_proto_rawDescGZIP(), []int{0, 0}
}

type MultiLayeredNesting_Nested1_Nested2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nested_3 *MultiLayeredNesting_Nested1_Nested2_Nested3 `protobuf:"bytes,1,opt,name=nested_3,json=nested3,proto3" json:"nested_3,omitempty"`
}

func (x *MultiLayeredNesting_Nested1_Nested2) Reset() {
	*x = MultiLayeredNesting_Nested1_Nested2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLayeredNesting_Nested1_Nested2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLayeredNesting_Nested1_Nested2) ProtoMessage() {}

// Deprecated: Use MultiLayeredNesting_Nested1_Nested2.ProtoReflect.Descriptor instead.
func (*MultiLayeredNesting_Nested1_Nested2) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_nesting_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *MultiLayeredNesting_Nested1_Nested2) GetNested_3() *MultiLayeredNesting_Nested1_Nested2_Nested3 {
	if x != nil {
		return x.Nested_3
	}
	return nil
}

type MultiLayeredNesting_Nested1_Nested2_Nested3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Nested3Oneof:
	//	*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String
	//	*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32
	Nested3Oneof isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof `protobuf_oneof:"nested3_oneof"`
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) Reset() {
	*x = MultiLayeredNesting_Nested1_Nested2_Nested3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test3_test_nesting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiLayeredNesting_Nested1_Nested2_Nested3) ProtoMessage() {}

// Deprecated: Use MultiLayeredNesting_Nested1_Nested2_Nested3.ProtoReflect.Descriptor instead.
func (*MultiLayeredNesting_Nested1_Nested2_Nested3) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_nesting_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) GetNested3Oneof() isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof {
	if x != nil {
		return x.Nested3Oneof
	}
	return nil
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) GetNested_3String() string {
	if x, ok := x.GetNested3Oneof().(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String); ok {
		return x.Nested_3String
	}
	return ""
}

func (x *MultiLayeredNesting_Nested1_Nested2_Nested3) GetNested_3Int32() int32 {
	if x, ok := x.GetNested3Oneof().(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32); ok {
		return x.Nested_3Int32
	}
	return 0
}

type isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof interface {
	isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof()
}

type MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String struct {
	Nested_3String string `protobuf:"bytes,1,opt,name=nested_3_string,json=nested3String,proto3,oneof"`
}

type MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32 struct {
	Nested_3Int32 int32 `protobuf:"varint,2,opt,name=nested_3_int32,json=nested3Int32,proto3,oneof"`
}

func (*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String) isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof() {
}

func (*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32) isMultiLayeredNesting_Nested1_Nested2_Nested3_Nested3Oneof() {
}

var File_internal_testprotos_test3_test_nesting_proto protoreflect.FileDescriptor

var file_internal_testprotos_test3_test_nesting_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x33, 0x22, 0xc4, 0x02, 0x0a, 0x13, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4a, 0x0a, 0x07, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x33, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x65, 0x64, 0x4e,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x31, 0x52, 0x07,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x31, 0x1a, 0xe0, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x31, 0x1a, 0xd4, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x32, 0x12,
	0x5b, 0x0a, 0x08, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x31, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x32, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x33, 0x52, 0x07, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x33, 0x1a, 0x6c, 0x0a, 0x07,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x33, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x33, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x33, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x33, 0x5f, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x33, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x0f, 0x0a, 0x0d, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x33, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_testprotos_test3_test_nesting_proto_rawDescOnce sync.Once
	file_internal_testprotos_test3_test_nesting_proto_rawDescData = file_internal_testprotos_test3_test_nesting_proto_rawDesc
)

func file_internal_testprotos_test3_test_nesting_proto_rawDescGZIP() []byte {
	file_internal_testprotos_test3_test_nesting_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_test3_test_nesting_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_test3_test_nesting_proto_rawDescData)
	})
	return file_internal_testprotos_test3_test_nesting_proto_rawDescData
}

var file_internal_testprotos_test3_test_nesting_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_test3_test_nesting_proto_goTypes = []interface{}{
	(*MultiLayeredNesting)(nil),                         // 0: goproto.proto.test3.MultiLayeredNesting
	(*MultiLayeredNesting_Nested1)(nil),                 // 1: goproto.proto.test3.MultiLayeredNesting.Nested1
	(*MultiLayeredNesting_Nested1_Nested2)(nil),         // 2: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2
	(*MultiLayeredNesting_Nested1_Nested2_Nested3)(nil), // 3: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3
}
var file_internal_testprotos_test3_test_nesting_proto_depIdxs = []int32{
	1, // 0: goproto.proto.test3.MultiLayeredNesting.nested1:type_name -> goproto.proto.test3.MultiLayeredNesting.Nested1
	3, // 1: goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.nested_3:type_name -> goproto.proto.test3.MultiLayeredNesting.Nested1.Nested2.Nested3
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_testprotos_test3_test_nesting_proto_init() }
func file_internal_testprotos_test3_test_nesting_proto_init() {
	if File_internal_testprotos_test3_test_nesting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_test3_test_nesting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLayeredNesting); i {
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
		file_internal_testprotos_test3_test_nesting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLayeredNesting_Nested1); i {
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
		file_internal_testprotos_test3_test_nesting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLayeredNesting_Nested1_Nested2); i {
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
		file_internal_testprotos_test3_test_nesting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiLayeredNesting_Nested1_Nested2_Nested3); i {
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
	file_internal_testprotos_test3_test_nesting_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3String)(nil),
		(*MultiLayeredNesting_Nested1_Nested2_Nested3_Nested_3Int32)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_test3_test_nesting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_test3_test_nesting_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_test3_test_nesting_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_test3_test_nesting_proto_msgTypes,
	}.Build()
	File_internal_testprotos_test3_test_nesting_proto = out.File
	file_internal_testprotos_test3_test_nesting_proto_rawDesc = nil
	file_internal_testprotos_test3_test_nesting_proto_goTypes = nil
	file_internal_testprotos_test3_test_nesting_proto_depIdxs = nil
}
