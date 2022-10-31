// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package cosmos_proto

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	strings "strings"
	sync "sync"
)

var (
	md_Any          protoreflect.MessageDescriptor
	fd_Any_type_url protoreflect.FieldDescriptor
	fd_Any_value    protoreflect.FieldDescriptor
)

func init() {
	file_proto_cosmos_proto_any_proto_init()
	md_Any = File_proto_cosmos_proto_any_proto.Messages().ByName("Any")
	fd_Any_type_url = md_Any.Fields().ByName("type_url")
	fd_Any_value = md_Any.Fields().ByName("value")
}

var _ protoreflect.Message = (*fastReflection_Any)(nil)

type fastReflection_Any Any

func (x *Any) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Any)(x)
}

func (x *Any) slowProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_proto_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Any_messageType fastReflection_Any_messageType
var _ protoreflect.MessageType = fastReflection_Any_messageType{}

type fastReflection_Any_messageType struct{}

func (x fastReflection_Any_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Any)(nil)
}
func (x fastReflection_Any_messageType) New() protoreflect.Message {
	return new(fastReflection_Any)
}
func (x fastReflection_Any_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Any
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Any) Descriptor() protoreflect.MessageDescriptor {
	return md_Any
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Any) Type() protoreflect.MessageType {
	return _fastReflection_Any_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Any) New() protoreflect.Message {
	return new(fastReflection_Any)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Any) Interface() protoreflect.ProtoMessage {
	return (*Any)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Any) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TypeUrl != "" {
		value := protoreflect.ValueOfString(x.TypeUrl)
		if !f(fd_Any_type_url, value) {
			return
		}
	}
	if len(x.Value) != 0 {
		value := protoreflect.ValueOfBytes(x.Value)
		if !f(fd_Any_value, value) {
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
func (x *fastReflection_Any) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "google.protobuf.Any.type_url":
		return x.TypeUrl != ""
	case "google.protobuf.Any.value":
		return len(x.Value) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Any) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "google.protobuf.Any.type_url":
		x.TypeUrl = ""
	case "google.protobuf.Any.value":
		x.Value = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Any) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "google.protobuf.Any.type_url":
		value := x.TypeUrl
		return protoreflect.ValueOfString(value)
	case "google.protobuf.Any.value":
		value := x.Value
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Any) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "google.protobuf.Any.type_url":
		x.TypeUrl = value.Interface().(string)
	case "google.protobuf.Any.value":
		x.Value = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Any) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "google.protobuf.Any.type_url":
		panic(fmt.Errorf("field type_url of message google.protobuf.Any is not mutable"))
	case "google.protobuf.Any.value":
		panic(fmt.Errorf("field value of message google.protobuf.Any is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Any) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "google.protobuf.Any.type_url":
		return protoreflect.ValueOfString("")
	case "google.protobuf.Any.value":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: google.protobuf.Any"))
		}
		panic(fmt.Errorf("message google.protobuf.Any does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Any) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in google.protobuf.Any", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Any) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Any) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Any) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Any) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Any)
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
		l = len(x.TypeUrl)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Value)
		if l > 0 {
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
		x := input.Message.Interface().(*Any)
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
		if len(x.Value) > 0 {
			i -= len(x.Value)
			copy(dAtA[i:], x.Value)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Value)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.TypeUrl) > 0 {
			i -= len(x.TypeUrl)
			copy(dAtA[i:], x.TypeUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TypeUrl)))
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
		x := input.Message.Interface().(*Any)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Any: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Any: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
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
				x.TypeUrl = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
				x.Value = append(x.Value[:0], dAtA[iNdEx:postIndex]...)
				if x.Value == nil {
					x.Value = []byte{}
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

// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.21.5
// source: proto/cosmos_proto/any.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//	Foo foo = ...;
//	Any any;
//	any.PackFrom(foo);
//	...
//	if (any.UnpackTo(&foo)) {
//	  ...
//	}
//
// Example 2: Pack and unpack a message in Java.
//
//	Foo foo = ...;
//	Any any = Any.pack(foo);
//	...
//	if (any.is(Foo.class)) {
//	  foo = any.unpack(Foo.class);
//	}
//
// Example 3: Pack and unpack a message in Python.
//
//	foo = Foo(...)
//	any = Any()
//	any.Pack(foo)
//	...
//	if any.Is(Foo.DESCRIPTOR):
//	  any.Unpack(foo)
//	  ...
//
// Example 4: Pack and unpack a message in Go
//
//	foo := &pb.Foo{...}
//	any, err := anypb.New(foo)
//	if err != nil {
//	  ...
//	}
//	...
//	foo := &pb.Foo{}
//	if err := any.UnmarshalTo(foo); err != nil {
//	  ...
//	}
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
// # JSON
//
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//	package google.profile;
//	message Person {
//	  string first_name = 1;
//	  string last_name = 2;
//	}
//
//	{
//	  "@type": "type.googleapis.com/google.profile.Person",
//	  "firstName": <string>,
//	  "lastName": <string>
//	}
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//	{
//	  "@type": "type.googleapis.com/google.protobuf.Duration",
//	  "value": "1.212s"
//	}
type Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	//   - If no scheme is provided, `https` is assumed.
	//   - An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//     value in binary format, or produce an error.
	//   - Applications are allowed to cache lookup results based on the
	//     URL, or have them precompiled into a binary to avoid any
	//     lookup. Therefore, binary compatibility needs to be preserved
	//     on changes to types. (Use versioned type names to manage
	//     breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Must be a valid serialized protocol buffer of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

// New marshals src into a new Any instance.
func New(src proto.Message) (*Any, error) {
	dst := new(Any)
	if err := dst.MarshalFrom(src); err != nil {
		return nil, err
	}
	return dst, nil
}

// MarshalFrom marshals src into dst as the underlying message
// using the provided marshal options.
//
// If no options are specified, call dst.MarshalFrom instead.
func MarshalFrom(dst *Any, src proto.Message, opts proto.MarshalOptions) error {
	if src == nil {
		return protoimpl.X.NewError("invalid nil source message")
	}
	b, err := opts.Marshal(src)
	if err != nil {
		return err
	}
	dst.TypeUrl = string(src.ProtoReflect().Descriptor().FullName())
	dst.Value = b
	return nil
}

// UnmarshalTo unmarshals the underlying message from src into dst
// using the provided unmarshal options.
// It reports an error if dst is not of the right message type.
//
// If no options are specified, call src.UnmarshalTo instead.
func UnmarshalTo(src *Any, dst proto.Message, opts proto.UnmarshalOptions) error {
	if src == nil {
		return protoimpl.X.NewError("invalid nil source message")
	}
	if !src.MessageIs(dst) {
		got := dst.ProtoReflect().Descriptor().FullName()
		want := src.MessageName()
		return protoimpl.X.NewError("mismatched message type: got %q, want %q", got, want)
	}
	return opts.Unmarshal(src.GetValue(), dst)
}

// UnmarshalNew unmarshals the underlying message from src into dst,
// which is newly created message using a type resolved from the type URL.
// The message type is resolved according to opt.Resolver,
// which should implement protoregistry.MessageTypeResolver.
// It reports an error if the underlying message type could not be resolved.
//
// If no options are specified, call src.UnmarshalNew instead.
func UnmarshalNew(src *Any, opts proto.UnmarshalOptions) (dst proto.Message, err error) {
	if src.GetTypeUrl() == "" {
		return nil, protoimpl.X.NewError("invalid empty type URL")
	}
	if opts.Resolver == nil {
		opts.Resolver = protoregistry.GlobalTypes
	}
	r, ok := opts.Resolver.(protoregistry.MessageTypeResolver)
	if !ok {
		return nil, protoregistry.NotFound
	}
	mt, err := r.FindMessageByURL(src.GetTypeUrl())
	if err != nil {
		if err == protoregistry.NotFound {
			return nil, err
		}
		return nil, protoimpl.X.NewError("could not resolve %q: %v", src.GetTypeUrl(), err)
	}
	dst = mt.New().Interface()
	return dst, opts.Unmarshal(src.GetValue(), dst)
}

// MessageIs reports whether the underlying message is of the same type as m.
func (x *Any) MessageIs(m proto.Message) bool {
	if m == nil {
		return false
	}
	url := x.GetTypeUrl()
	name := string(m.ProtoReflect().Descriptor().FullName())
	if !strings.HasSuffix(url, name) {
		return false
	}
	return len(url) == len(name) || url[len(url)-len(name)-1] == '/'
}

// MessageName reports the full name of the underlying message,
// returning an empty string if invalid.
func (x *Any) MessageName() protoreflect.FullName {
	url := x.GetTypeUrl()
	name := protoreflect.FullName(url)
	if i := strings.LastIndexByte(url, '/'); i >= 0 {
		name = name[i+len("/"):]
	}
	if !name.IsValid() {
		return ""
	}
	return name
}

// MarshalFrom marshals m into x as the underlying message.
func (x *Any) MarshalFrom(m proto.Message) error {
	return MarshalFrom(x, m, proto.MarshalOptions{})
}

// UnmarshalTo unmarshals the contents of the underlying message of x into m.
// It resets m before performing the unmarshal operation.
// It reports an error if m is not of the right message type.
func (x *Any) UnmarshalTo(m proto.Message) error {
	return UnmarshalTo(x, m, proto.UnmarshalOptions{})
}

// UnmarshalNew unmarshals the contents of the underlying message of x into
// a newly allocated message of the specified type.
// It reports an error if the underlying message type could not be resolved.
func (x *Any) UnmarshalNew() (proto.Message, error) {
	return UnmarshalNew(x, proto.UnmarshalOptions{})
}

func (x *Any) Reset() {
	*x = Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_proto_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Any) ProtoMessage() {}

// Deprecated: Use Any.ProtoReflect.Descriptor instead.
func (*Any) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_proto_any_proto_rawDescGZIP(), []int{0}
}

func (x *Any) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Any) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_proto_cosmos_proto_any_proto protoreflect.FileDescriptor

var file_proto_cosmos_proto_any_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x36, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x68, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x08,
	0x41, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cosmos_proto_any_proto_rawDescOnce sync.Once
	file_proto_cosmos_proto_any_proto_rawDescData = file_proto_cosmos_proto_any_proto_rawDesc
)

func file_proto_cosmos_proto_any_proto_rawDescGZIP() []byte {
	file_proto_cosmos_proto_any_proto_rawDescOnce.Do(func() {
		file_proto_cosmos_proto_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cosmos_proto_any_proto_rawDescData)
	})
	return file_proto_cosmos_proto_any_proto_rawDescData
}

var file_proto_cosmos_proto_any_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_cosmos_proto_any_proto_goTypes = []interface{}{
	(*Any)(nil), // 0: google.protobuf.Any
}
var file_proto_cosmos_proto_any_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cosmos_proto_any_proto_init() }
func file_proto_cosmos_proto_any_proto_init() {
	if File_proto_cosmos_proto_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cosmos_proto_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Any); i {
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
			RawDescriptor: file_proto_cosmos_proto_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_cosmos_proto_any_proto_goTypes,
		DependencyIndexes: file_proto_cosmos_proto_any_proto_depIdxs,
		MessageInfos:      file_proto_cosmos_proto_any_proto_msgTypes,
	}.Build()
	File_proto_cosmos_proto_any_proto = out.File
	file_proto_cosmos_proto_any_proto_rawDesc = nil
	file_proto_cosmos_proto_any_proto_goTypes = nil
	file_proto_cosmos_proto_any_proto_depIdxs = nil
}
