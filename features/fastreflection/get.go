package fastreflection

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genGetProto generates the implementation for protoreflect.Message.Get
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
	genGetExtension(g, msg)
	g.P("switch descriptor.FullName() {")
	// implement the fast Get function
	for _, genFd := range msg.Fields {
		fd := genFd.Desc
		g.P("case \"", fd.FullName(), "\":")
		getfuncForField(g, fd.Kind(), genFd.GoName, genFd)
	}
	// insert default case which panics
	g.P("default:")
	g.P("panic(fmt.Errorf(\"message ", msg.Desc.FullName(), " does not contain field %s\", descriptor.Name()))")
	g.P("}")
	g.P("}")
}

// genGetExtension generates the logic that handles protobuf extensions
func genGetExtension(g *protogen.GeneratedFile, msg *protogen.Message) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	const fmtPkg = protogen.GoImportPath("fmt")

	g.P("// handle extension logic")
	g.P("if descriptor.IsExtension() && descriptor.ContainingMessage().FullName() == \"", msg.Desc.FullName(), "\" {")
	g.P("if _, ok := descriptor.(", pref.Ident("ExtensionTypeDescriptor"), "); !ok {")
	g.P("panic(", fmtPkg.Ident("Errorf"), "(\"%s: extension field descriptor does not implement ExtensionTypeDescriptor\", descriptor.FullName()))")
	g.P("}")
	g.P("panic(\"implement xt logic\")") // TODO(fdymylja)
	g.P("}")
	g.P()
}

func getfuncForField(g *protogen.GeneratedFile, kind protoreflect.Kind, fieldName string, genFd *protogen.Field) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	if genFd.Oneof != nil {
		getOneOf(g, kind, genFd, genFd.Oneof)
		return
	}

	fieldRef := "x." + fieldName
	g.P("value := ", fieldRef)
	switch {
	case genFd.Desc.IsMap():
		g.P("_ = value")
		g.P("panic(\"not implemented\")")
		return
	case genFd.Desc.IsList():
		g.P("_ = value")
		g.P("panic(\"not implemented\")")
		return
	}

	switch kind {
	case protoreflect.BoolKind:
		g.P("return ", pref.Ident("ValueOfBool"), "(value)")
	case protoreflect.EnumKind:
		g.P("return ", pref.Ident("ValueOfEnum"), "((", pref.Ident("EnumNumber"), ")", "(value)", ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		g.P("return ", pref.Ident("ValueOfInt32"), "(value)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		g.P("return ", pref.Ident("ValueOfUint32"), "(value)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		g.P("return ", pref.Ident("ValueOfInt64"), "(value)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.P("return ", pref.Ident("ValueOfUint64"), "(value)")
	case protoreflect.FloatKind:
		g.P("return ", pref.Ident("ValueOfFloat32"), "(value)")
	case protoreflect.DoubleKind:
		g.P("return ", pref.Ident("ValueOfFloat64"), "(value)")
	case protoreflect.StringKind:
		g.P("return ", pref.Ident("ValueOfString"), "(value)")
	case protoreflect.BytesKind:
		g.P("return ", pref.Ident("ValueOfBytes"), "(value)")
	case protoreflect.MessageKind, protoreflect.GroupKind:
		g.P("return ", pref.Ident("ValueOfMessage"), "(value.ProtoReflect())")
	}
}

func getOneOf(g *protogen.GeneratedFile, kind protoreflect.Kind, fd *protogen.Field, oneof *protogen.Oneof) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")

	// handle the case in which the oneof field is not set
	g.P("if x.", oneof.GoName, " == nil {")
	switch fd.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(nil)")
	default:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(", zeroValueForField(g, fd), ")")
	}
	// handle the case in which oneof field is set and it matches our sub-onefield type
	g.P("} else if v, ok := x.", oneof.GoName, ".(*", fd.GoIdent, "); ok {")
	oneofTypeContainerFieldName := fd.GoName // field containing the oneof value
	switch fd.Desc.Kind() {
	case protoreflect.MessageKind: // it can be mutable
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(v.", oneofTypeContainerFieldName, ".ProtoReflect())")
	case protoreflect.EnumKind:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "((", pref.Ident("EnumNumber"), ")(v.", oneofTypeContainerFieldName, "))")
	default:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(v.", oneofTypeContainerFieldName, ")")
	}
	// handle the case in which the oneof field is set but it does not match our field type
	g.P("} else {")
	switch fd.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(nil)")
	default:
		g.P("return ", kindToValueConstructor(fd.Desc.Kind()), "(", zeroValueForField(g, fd), ")")
	}
	g.P("}")
}
