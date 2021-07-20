package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func genProtoMessageFunctions(g *GeneratedFile, msg *protogen.Message) {
	genGetMethods(g, msg)
	g.P()
	genDescriptorProto(g, msg)
	g.P()
	genTypeProto(g, msg)
	g.P()
	genNewProto(g, msg)
	g.P()
	genInterfaceProto(g, msg)
	g.P()
	genRangeProto(g, msg)
	g.P()
	genHasProto(g, msg)
	g.P()
	genClearProto(g, msg)
	g.P()
	genGetProto(g, msg)
	g.P()
	genSetProto(g, msg)
	g.P()
	genMutableProto(g, msg)
	g.P()
	genNewFieldProto(g, msg)
	g.P()
	genWhichOneOfProto(g, msg)
	g.P()
	genGetUnkownProto(g, msg)
	g.P()
	genSetUnkownProto(g, msg)
	g.P()
	genIsValidProto(g, msg)
	g.P()
	genProtoMethodsProto(g, msg)
	g.P()
}

func genGetMethods(g *GeneratedFile, msg *protogen.Message) {
	g.P("// returns the fast methods for the message")
	g.P("func (x ", msg.GoIdent.GoName, ") GetMethods() *", protoifacePackage.Ident("Methods"), " {")
	g.P("return &", protoifacePackage.Ident("Methods{"))
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Flags: 0,")
	g.P("Size: func(input ", protoifacePackage.Ident("SizeInput"), ") ", protoifacePackage.Ident("SizeOutput"), " {")
	g.P("return ", protoifacePackage.Ident("SizeOutput"), "{")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Size: x.Size(),")
	g.P("}")
	g.P("},")
	g.P("Marshal: func(input ", protoifacePackage.Ident("MarshalInput"), ") (", protoifacePackage.Ident("MarshalOutput"), ", error) {")
	g.P("v, ok := input.Message.(", msg.GoIdent.GoName, ")")
	g.P("if !ok {")
	g.P("return ", protoifacePackage.Ident("MarshalOutput"), "{}, errors.New(\"", msg.GoIdent.GoName, " does not implement the protoreflect.Message interface\")")
	g.P("}")
	g.P()
	g.P("bz, err := v.Marshal()")
	g.P("if err != nil {")
	g.P("return ", protoifacePackage.Ident("MarshalOutput"), "{}, err")
	g.P("}")
	g.P("return ", protoifacePackage.Ident("MarshalOutput"), "{")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Buf: bz,")
	g.P("}, nil")
	g.P("},")
	g.P("Unmarshal: func(input ", protoifacePackage.Ident("UnmarshalInput"), ") (", protoifacePackage.Ident("UnmarshalOutput"), ", error){")
	g.P("v, ok := input.Message.(*", msg.GoIdent.GoName, ")")
	g.P("if !ok {")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{}, errors.New(\"", msg.GoIdent.GoName, " does not implement the protoreflect.Message interface\")")
	g.P("}")
	g.P()
	g.P("if len(input.Buf) < 1 {")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{}, errors.New(\"unmarshal input did not contain any bytes to unmarshal\")")
	g.P("}")
	g.P("err := v.Unmarshal(input.Buf)")
	g.P("if err != nil {")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{}, err")
	g.P("}")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Flags: 0,")
	g.P("}, nil")
	g.P("},")
	g.P("Merge: nil,")
	g.P("CheckInitialized: nil,")
	g.P("}")
	g.P("}")
}

func genDescriptorProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Descriptor returns message descriptor, which contains only the protobuf")
	g.P("// type information for the message.")
	g.P("func (x ", msg.GoIdent.GoName, ") Descriptor() ", protoreflectPackage.Ident("MessageDescriptor"), " {")
	g.P("return x.ProtoReflect().Descriptor()")
	g.P("}")
}

func genTypeProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Type returns the message type, which encapsulates both Go and protobuf")
	g.P("// type information. If the Go type information is not needed,")
	g.P("// it is recommended that the message descriptor be used instead.")
	g.P("func (x ", msg.GoIdent.GoName, ") Type() ", protoreflectPackage.Ident("MessageType"), " {")
	g.P("return x.ProtoReflect().Type()")
	g.P("}")
}

func genNewProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// New returns a newly allocated and mutable empty message.")
	g.P("func (x ", msg.GoIdent.GoName, ") New() ", protoreflectPackage.Ident("Message"), " {")
	g.P("return x.ProtoReflect().New()")
	g.P("}")
}

func genInterfaceProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Interface unwraps the message reflection interface and")
	g.P("// returns the underlying ProtoMessage interface.")
	g.P("func (x ", msg.GoIdent.GoName, ") Interface() ", protoreflectPackage.Ident("ProtoMessage"), " {")
	g.P("return x.ProtoReflect().Interface()")
	g.P("}")
}

func genRangeProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Range iterates over every populated field in an undefined order,")
	g.P("// calling f for each field descriptor and value encountered.")
	g.P("// Range returns immediately if f returns false.")
	g.P("// While iterating, mutating operations may only be performed")
	g.P("// on the current field descriptor.")
	g.P("func (x ", msg.GoIdent.GoName, ") Range(f func(", protoreflectPackage.Ident("FieldDescriptor"), ", ", protoreflectPackage.Ident("Value"), ") bool) {")
	g.P("x.ProtoReflect().Range(f)")
	g.P("}")
}

func genHasProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Has reports whether a field is populated.")
	g.P("//")
	g.P("// Some fields have the property of nullability where it is possible to")
	g.P("// distinguish between the default value of a field and whether the field")
	g.P("// was explicitly populated with the default value. Singular message fields,")
	g.P("// member fields of a oneof, and proto2 scalar fields are nullable. Such")
	g.P("// fields are populated only if explicitly set.")
	g.P("//")
	g.P("// In other cases (aside from the nullable cases above),")
	g.P("// a proto3 scalar field is populated if it contains a non-zero value, and")
	g.P("// a repeated field is populated if it is non-empty.")
	g.P("func (x ", msg.GoIdent.GoName, ") Has(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ") bool {")
	g.P("return x.ProtoReflect().Has(descriptor)")
	g.P("}")
}

func genClearProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Clear clears the field such that a subsequent Has call reports false.")
	g.P("//")
	g.P("// Clearing an extension field clears both the extension type and value")
	g.P("// associated with the given field number.")
	g.P("//")
	g.P("// Clear is a mutating operation and unsafe for concurrent use.")
	g.P("func (x ", msg.GoIdent.GoName, ") Clear(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ") {")
	g.P("x.ProtoReflect().Clear(descriptor)")
	g.P("}")
}

func genGetProto(g *GeneratedFile, msg *protogen.Message) {
	// we check if there are map or list fields
	// so that we can generate fast reflection wrapper types.
	g.P("// Get retrieves the value for a field.")
	g.P("//")
	g.P("// For unpopulated scalars, it returns the default value, where")
	g.P("// the default value of a bytes scalar is guaranteed to be a copy.")
	g.P("// For unpopulated composite types, it returns an empty, read-only view")
	g.P("// of the value; to obtain a mutable reference, use Mutable.")
	g.P("func (x *", msg.GoIdent.GoName, ") Get(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ") ", protoreflectPackage.Ident("Value"), " {")
	g.P("switch descriptor.Name() {")
	// implement the fast Get function
	for _, genFd := range msg.Fields {
		fd := genFd.Desc
		g.P("case \"", fd.Name(), "\":")
		getfuncForField(g, fd.Kind(), genFd.GoName, genFd)
	}
	// insert default case which panics
	g.P("default:")
	g.P("panic(fmt.Errorf(\"message ", msg.Desc.FullName(), " does not contain field %s\", descriptor.Name()))")
	g.P("}")
	g.P("}")
}

func getfuncForField(g *GeneratedFile, kind protoreflect.Kind, fieldName string, genFd *protogen.Field) {
	pkg := protoreflectPackage
	fieldRef := "x." + fieldName

	switch {
	case genFd.Desc.IsMap():
		return
	case genFd.Desc.IsList():
		return
	}

	switch kind {
	case protoreflect.BoolKind:
		g.P("return ", pkg.Ident("ValueOfBool"), "(", fieldRef, ")")
	case protoreflect.EnumKind:
		g.P("return ", pkg.Ident("ValueOfEnum"), "((", pkg.Ident("EnumNumber"), ")", "(", fieldRef, ")", ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		g.P("return ", pkg.Ident("ValueOfInt32"), "(", fieldRef, ")")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		g.P("return ", pkg.Ident("ValueOfUint32"), "(", fieldRef, ")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		g.P("return ", pkg.Ident("ValueOfInt64"), "(", fieldRef, ")")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.P("return ", pkg.Ident("ValueOfUint64"), "(", fieldRef, ")")
	case protoreflect.FloatKind:
		g.P("return ", pkg.Ident("ValueOfFloat32"), "(", fieldRef, ")")
	case protoreflect.DoubleKind:
		g.P("return ", pkg.Ident("ValueOfFloat64"), "(", fieldRef, ")")
	case protoreflect.StringKind:
		g.P("return ", pkg.Ident("ValueOfString"), "(", fieldRef, ")")
	case protoreflect.BytesKind:
		g.P("return ", pkg.Ident("ValueOfBytes"), "(", fieldRef, ")")
	case protoreflect.MessageKind, protoreflect.GroupKind:
		g.P("return ", pkg.Ident("ValueOfMessage"), "(", fieldRef, ")")
	}
}

func genSetProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Set stores the value for a field.")
	g.P("//")
	g.P("// For a field belonging to a oneof, it implicitly clears any other field")
	g.P("// that may be currently set within the same oneof.")
	g.P("// For extension fields, it implicitly stores the provided ExtensionType.")
	g.P("// When setting a composite type, it is unspecified whether the stored value")
	g.P("// aliases the source's memory in any way. If the composite value is an")
	g.P("// empty, read-only value, then it panics.")
	g.P("//")
	g.P("// Set is a mutating operation and unsafe for concurrent use.")
	g.P("func (x ", msg.GoIdent.GoName, ") Set(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ", value ", protoreflectPackage.Ident("Value"), ") {")
	g.P("x.ProtoReflect().Set(descriptor, value)")
	g.P("}")
}

func genMutableProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// Mutable returns a mutable reference to a composite type.")
	g.P("//")
	g.P("// If the field is unpopulated, it may allocate a composite value.")
	g.P("// For a field belonging to a oneof, it implicitly clears any other field")
	g.P("// that may be currently set within the same oneof.")
	g.P("// For extension fields, it implicitly stores the provided ExtensionType")
	g.P("// if not already stored.")
	g.P("// It panics if the field does not contain a composite type.")
	g.P("//")
	g.P("// Mutable is a mutating operation and unsafe for concurrent use.")
	g.P("func (x ", msg.GoIdent.GoName, ") Mutable(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ") ", protoreflectPackage.Ident("Value"), " {")
	g.P("return x.ProtoReflect().Mutable(descriptor)")
	g.P("}")
}

func genNewFieldProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// NewField returns a new value that is assignable to the field")
	g.P("// for the given descriptor. For scalars, this returns the default value.")
	g.P("// For lists, maps, and messages, this returns a new, empty, mutable value.")
	g.P("func (x ", msg.GoIdent.GoName, ") NewField(descriptor ", protoreflectPackage.Ident("FieldDescriptor"), ") ", protoreflectPackage.Ident("Value"), " {")
	g.P("return x.ProtoReflect().NewField(descriptor)")
	g.P("}")
}

func genWhichOneOfProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// WhichOneof reports which field within the oneof is populated,")
	g.P("// returning nil if none are populated.")
	g.P("// It panics if the oneof descriptor does not belong to this message.")
	g.P("func (x ", msg.GoIdent.GoName, ") WhichOneof(descriptor ", protoreflectPackage.Ident("OneofDescriptor"), ") ", protoreflectPackage.Ident("FieldDescriptor"), " {")
	g.P("return x.ProtoReflect().WhichOneof(descriptor)")
	g.P("}")
}

func genGetUnkownProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// GetUnknown retrieves the entire list of unknown fields.")
	g.P("// The caller may only mutate the contents of the RawFields")
	g.P("// if the mutated bytes are stored back into the message with SetUnknown.")
	g.P("func (x ", msg.GoIdent.GoName, ") GetUnknown() ", protoreflectPackage.Ident("RawFields"), " {")
	g.P("return x.ProtoReflect().GetUnknown()")
	g.P("}")
}

func genSetUnkownProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// SetUnknown stores an entire list of unknown fields.")
	g.P("// The raw fields must be syntactically valid according to the wire format.")
	g.P("// An implementation may panic if this is not the case.")
	g.P("// Once stored, the caller must not mutate the content of the RawFields.")
	g.P("// An empty RawFields may be passed to clear the fields.")
	g.P("//")
	g.P("// SetUnknown is a mutating operation and unsafe for concurrent use.")
	g.P("func (x ", msg.GoIdent.GoName, ") SetUnknown(fields ", protoreflectPackage.Ident("RawFields"), ") {")
	g.P("x.ProtoReflect().SetUnknown(fields)")
	g.P("}")
}

func genIsValidProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// IsValid reports whether the message is valid.")
	g.P("//")
	g.P("// An invalid message is an empty, read-only value.")
	g.P("//")
	g.P("// An invalid message often corresponds to a nil pointer of the concrete")
	g.P("// message type, but the details are implementation dependent.")
	g.P("// Validity is not part of the protobuf data model, and may not")
	g.P("// be preserved in marshaling or other operations.")
	g.P("func (x ", msg.GoIdent.GoName, ") IsValid() bool {")
	g.P("return x.ProtoReflect().IsValid()")
	g.P("}")
}

func genProtoMethodsProto(g *GeneratedFile, msg *protogen.Message) {
	g.P("// ProtoMethods returns optional fast-path implementations of various operations.")
	g.P("// This method may return nil.")
	g.P("//")
	g.P("// The returned methods type is identical to")
	g.P(`// "google.golang.org/protobuf/runtime/protoiface".Methods.`)
	g.P("// Consult the protoiface package documentation for details.")
	g.P("func (x ", msg.GoIdent.GoName, ") ProtoMethods() *", protoifacePackage.Ident("Methods"), " {")
	g.P("return x.GetMethods()")
	g.P("}")
}
