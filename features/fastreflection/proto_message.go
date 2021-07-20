package fastreflection

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-proto/features/fastreflection/copied"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func GenProtoMessage(f *protogen.File, g *protogen.GeneratedFile, message *protogen.Message) {
	genProtoMessageFunctions(f, g, message)
}

func genProtoMessageFunctions(f *protogen.File, g *protogen.GeneratedFile, msg *protogen.Message) {
	g.P()
	typeName := genReflectionType(f, g, msg)
	g.P()
	genDescriptorProto(g, msg, typeName)
	g.P()
	genTypeProto(g, msg, typeName)
	g.P()
	genNewProto(g, msg, typeName)
	g.P()
	genInterfaceProto(g, msg, typeName)
	g.P()
	genRangeProto(g, msg, typeName)
	g.P()
	genHasProto(g, msg, typeName)
	g.P()
	genClearProto(g, msg, typeName)
	g.P()
	genGetProto(g, msg, typeName)
	g.P()
	genSetProto(g, msg, typeName)
	g.P()
	genMutableProto(g, msg, typeName)
	g.P()
	genNewFieldProto(g, msg, typeName)
	g.P()
	genWhichOneOfProto(g, msg, typeName)
	g.P()
	genGetUnkownProto(g, msg, typeName)
	g.P()
	genSetUnkownProto(g, msg, typeName)
	g.P()
	genIsValidProto(g, msg, typeName)
	g.P()
	genProtoMethodsProto(g, msg, typeName)
	g.P()
}

func genMessageReflectMethods(g *protogen.GeneratedFile, f *copied.FileInfo, m *protogen.Message) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	const primpl = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")
	idx := func() int {
		var id int
		var found bool
		for mInfo, index := range f.AllMessagesByPtr {
			if mInfo.Message.Desc.FullName() == m.Desc.FullName() {
				id = index
				found = true
			}
		}
		if !found {
			panic("not found")
		}
		return id
	}()
	typesVar := copied.MessageTypesVarName(f)

	// ProtoReflect method.
	g.P("func (x *", m.GoIdent, ") slowProtoReflect() ", pref.Ident("Message"), " {")
	g.P("mi := &", typesVar, "[", idx, "]")
	g.P("if ", primpl.Ident("UnsafeEnabled"), " && x != nil {")
	g.P("ms := ", primpl.Ident("X"), ".MessageStateOf(", primpl.Ident("Pointer"), "(x))")
	g.P("if ms.LoadMessageInfo() == nil {")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("return ms")
	g.P("}")
	g.P("return mi.MessageOf(x)")
	g.P("}")
	g.P()
}

// genReflectionType generates the reflection type for the message
func genReflectionType(file *protogen.File, g *protogen.GeneratedFile, msg *protogen.Message) string {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	typeName := fmt.Sprintf("fastReflection_%s", msg.GoIdent.GoName)
	genMessageReflectMethods(g, copied.NewFileInfo(file), msg)
	// gen interface assertion
	g.P("var _ ", pref.Ident("Message"), " = (*", typeName, ")(nil)") // TODO(fdymylja): pointer really required?
	g.P()
	// gen type
	g.P("type ", typeName, " ", msg.GoIdent.GoName)
	// gen msg implementation
	g.P("func (x *", msg.GoIdent.GoName, ") ProtoReflect() ", pref.Ident("Message"), "{") // TODO(fdymylja): replace me with ProtoReflect function
	g.P("return (*", typeName, ")(x)")
	g.P("}")
	g.P()
	return typeName
}

func genDescriptorProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// Descriptor returns message descriptor, which contains only the protobuf")
	g.P("// type information for the message.")
	g.P("func (x *", typeName, ") Descriptor() ", pref.Ident("MessageDescriptor"), " {")
	slowReflectionFallBack(g, msg, true, "Descriptor")
	g.P("}")
}

func genTypeProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// Type returns the message type, which encapsulates both Go and protobuf")
	g.P("// type information. If the Go type information is not needed,")
	g.P("// it is recommended that the message descriptor be used instead.")
	g.P("func (x *", typeName, ") Type() ", pref.Ident("MessageType"), " {")
	slowReflectionFallBack(g, msg, true, "Type")
	g.P("}")
}

func genNewProto(g *protogen.GeneratedFile, _ *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// New returns a newly allocated and mutable empty message.")
	g.P("func (x *", typeName, ") New() ", pref.Ident("Message"), " {")
	g.P("return new(", typeName, ")")
	g.P("}")
}

func genInterfaceProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// Interface unwraps the message reflection interface and")
	g.P("// returns the underlying ProtoMessage interface.")
	g.P("func (x *", typeName, ") Interface() ", pref.Ident("ProtoMessage"), " {")
	g.P("return (*", msg.GoIdent, ")(x)")
	g.P("}")
}

func genRangeProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// Range iterates over every populated field in an undefined order,")
	g.P("// calling f for each field descriptor and value encountered.")
	g.P("// Range returns immediately if f returns false.")
	g.P("// While iterating, mutating operations may only be performed")
	g.P("// on the current field descriptor.")
	g.P("func (x *", typeName, ") Range(f func(", pref.Ident("FieldDescriptor"), ", ", pref.Ident("Value"), ") bool) {")
	slowReflectionFallBack(g, msg, false, "Range", "f")
	g.P("}")
}

func genHasProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

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
	g.P("func (x *", typeName, ") Has(descriptor ", pref.Ident("FieldDescriptor"), ") bool {")
	slowReflectionFallBack(g, msg, true, "Has", "descriptor")
	g.P("}")
}

func genClearProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// Clear clears the field such that a subsequent Has call reports false.")
	g.P("//")
	g.P("// Clearing an extension field clears both the extension type and value")
	g.P("// associated with the given field number.")
	g.P("//")
	g.P("// Clear is a mutating operation and unsafe for concurrent use.")
	g.P("func (x *", typeName, ") Clear(descriptor ", pref.Ident("FieldDescriptor"), ") {")
	slowReflectionFallBack(g, msg, false, "Clear", "descriptor")
	g.P("}")
}

func genGetProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	// we check if there are map or list fields
	// so that we can generate fast reflection wrapper types.
	for _, genFd := range msg.Fields {
		switch {
		case genFd.Desc.IsList():
			GenList(g, genFd)
		}
	}
	g.P("// Get retrieves the value for a field.")
	g.P("//")
	g.P("// For unpopulated scalars, it returns the default value, where")
	g.P("// the default value of a bytes scalar is guaranteed to be a copy.")
	g.P("// For unpopulated composite types, it returns an empty, read-only view")
	g.P("// of the value; to obtain a mutable reference, use Mutable.")
	g.P("func (x *", typeName, ") Get(descriptor ", pref.Ident("FieldDescriptor"), ") ", pref.Ident("Value"), " {")
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

func getfuncForField(g *protogen.GeneratedFile, kind protoreflect.Kind, fieldName string, genFd *protogen.Field) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	fieldRef := "x." + fieldName

	switch {
	case genFd.Desc.IsMap():
		return
	case genFd.Desc.IsList():
		return
	}

	switch kind {
	case protoreflect.BoolKind:
		g.P("return ", pref.Ident("ValueOfBool"), "(", fieldRef, ")")
	case protoreflect.EnumKind:
		g.P("return ", pref.Ident("ValueOfEnum"), "((", pref.Ident("EnumNumber"), ")", "(", fieldRef, ")", ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		g.P("return ", pref.Ident("ValueOfInt32"), "(", fieldRef, ")")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		g.P("return ", pref.Ident("ValueOfUint32"), "(", fieldRef, ")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		g.P("return ", pref.Ident("ValueOfInt64"), "(", fieldRef, ")")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.P("return ", pref.Ident("ValueOfUint64"), "(", fieldRef, ")")
	case protoreflect.FloatKind:
		g.P("return ", pref.Ident("ValueOfFloat32"), "(", fieldRef, ")")
	case protoreflect.DoubleKind:
		g.P("return ", pref.Ident("ValueOfFloat64"), "(", fieldRef, ")")
	case protoreflect.StringKind:
		g.P("return ", pref.Ident("ValueOfString"), "(", fieldRef, ")")
	case protoreflect.BytesKind:
		g.P("return ", pref.Ident("ValueOfBytes"), "(", fieldRef, ")")
	case protoreflect.MessageKind, protoreflect.GroupKind:
		g.P("return ", pref.Ident("ValueOfMessage"), "(", fieldRef, ".ProtoReflect())")
	}
}

func genSetProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

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
	g.P("func (x *", typeName, ") Set(descriptor ", pref.Ident("FieldDescriptor"), ", value ", pref.Ident("Value"), ") {")
	slowReflectionFallBack(g, msg, false, "Set", "descriptor", "value")
	g.P("}")
}

func genMutableProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

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
	g.P("func (x *", typeName, ") Mutable(descriptor ", pref.Ident("FieldDescriptor"), ") ", pref.Ident("Value"), " {")
	slowReflectionFallBack(g, msg, true, "Mutable", "descriptor")
	g.P("}")
}

func genNewFieldProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// NewField returns a new value that is assignable to the field")
	g.P("// for the given descriptor. For scalars, this returns the default value.")
	g.P("// For lists, maps, and messages, this returns a new, empty, mutable value.")
	g.P("func (x *", typeName, ") NewField(descriptor ", pref.Ident("FieldDescriptor"), ") ", pref.Ident("Value"), " {")
	slowReflectionFallBack(g, msg, true, "NewField", "descriptor")
	g.P("}")
}

func genWhichOneOfProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// WhichOneof reports which field within the oneof is populated,")
	g.P("// returning nil if none are populated.")
	g.P("// It panics if the oneof descriptor does not belong to this message.")
	g.P("func (x *", typeName, ") WhichOneof(descriptor ", pref.Ident("OneofDescriptor"), ") ", pref.Ident("FieldDescriptor"), " {")
	slowReflectionFallBack(g, msg, true, "WhichOneof", "descriptor")
	g.P("}")
}

func genGetUnkownProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// GetUnknown retrieves the entire list of unknown fields.")
	g.P("// The caller may only mutate the contents of the RawFields")
	g.P("// if the mutated bytes are stored back into the message with SetUnknown.")
	g.P("func (x *", typeName, ") GetUnknown() ", pref.Ident("RawFields"), " {")
	slowReflectionFallBack(g, msg, true, "GetUnknown")
	g.P("}")
}

func genSetUnkownProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	g.P("// SetUnknown stores an entire list of unknown fields.")
	g.P("// The raw fields must be syntactically valid according to the wire format.")
	g.P("// An implementation may panic if this is not the case.")
	g.P("// Once stored, the caller must not mutate the content of the RawFields.")
	g.P("// An empty RawFields may be passed to clear the fields.")
	g.P("//")
	g.P("// SetUnknown is a mutating operation and unsafe for concurrent use.")
	g.P("func (x *", typeName, ") SetUnknown(fields ", pref.Ident("RawFields"), ") {")
	slowReflectionFallBack(g, msg, false, "SetUnknown", "fields")
	g.P("}")
}

func genIsValidProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	g.P("// IsValid reports whether the message is valid.")
	g.P("//")
	g.P("// An invalid message is an empty, read-only value.")
	g.P("//")
	g.P("// An invalid message often corresponds to a nil pointer of the concrete")
	g.P("// message type, but the details are implementation dependent.")
	g.P("// Validity is not part of the protobuf data model, and may not")
	g.P("// be preserved in marshaling or other operations.")

	g.P("func (x *", typeName, ") IsValid() bool {")
	slowReflectionFallBack(g, msg, true, "IsValid")
	g.P("}")
}

func genProtoMethodsProto(g *protogen.GeneratedFile, msg *protogen.Message, typeName string) {
	const protoiface = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")

	g.P("// ProtoMethods returns optional fast-path implementations of various operations.")
	g.P("// This method may return nil.")
	g.P("//")
	g.P("// The returned methods type is identical to")
	g.P(`// "google.golang.org/protobuf/runtime/protoiface".Methods.`)
	g.P("// Consult the protoiface package documentation for details.")
	g.P("func (x *", typeName, ") ProtoMethods() *", protoiface.Ident("Methods"), " {")
	slowReflectionFallBack(g, msg, true, "ProtoMethods")
	g.P("}")
}

// slowReflectionFallBack can be used to fallback on slow reflection methods
func slowReflectionFallBack(g *protogen.GeneratedFile, msg *protogen.Message, returns bool, method string, args ...string) {
	switch returns {
	case true:
		g.P("return (*", msg.GoIdent.GoName, ")(x).slowProtoReflect().", method, "(", strings.Join(args, ","), ")")
	case false:
		g.P("(*", msg.GoIdent.GoName, ")(x).slowProtoReflect().", method, "(", strings.Join(args, ","), ")")
	}
}
