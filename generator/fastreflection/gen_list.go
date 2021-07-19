package fastreflection

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GenList generates the reflection fast paths over repeated fields
func GenList(g *protogen.GeneratedFile, field *protogen.Field) {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	const fmtPkg = protogen.GoImportPath("fmt")
	typeName := listTypeName(field)
	g.P("type ", typeName, " struct {")
	g.P("list []", getType(g, field))
	g.P("}")
	g.P()

	// implement protoreflect.List
	// type assertion
	g.P("var _ ", pref.Ident("List"), " = (*", typeName, ")(nil)")

	// Len
	g.P("func (x *", typeName, ") Len() int {")
	g.P("return len(x.list)")
	g.P("}")
	g.P()

	// Get
	g.P("func (x *", typeName, ") Get(i int) ", pref.Ident("Value"), " {")
	constructor := kindToValueConstructor(field.Desc.Kind())
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P("return ", constructor, "(x.list[i].ProtoReflect())")
	case protoreflect.EnumKind:
		g.P("return ", constructor, "((", pref.Ident("EnumNumber"), ")(x.list[i]))")
	default:
		g.P("return ", constructor, "(x.list[i])")
	}
	g.P("}")
	g.P()

	// Set
	g.P("func (x *", typeName, ") Set(i int, value ", pref.Ident("Value"), ") {")
	concreteValueName := genPrefValueToGoValue(g, field)
	g.P("x.list[i] = ", concreteValueName)
	g.P("}")
	g.P()

	// Append
	g.P("func (x *", typeName, ") Append(value ", pref.Ident("Value"), ") {")
	concreteValueName = genPrefValueToGoValue(g, field)
	g.P("x.list = append(x.list, ", concreteValueName, ")")
	g.P("}")
	g.P()

	// AppendMutable
	g.P("func (x *", typeName, ") AppendMutable() ", pref.Ident("Value"), " {")
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P("v := new(", g.QualifiedGoIdent(field.Message.GoIdent), ")")
		g.P("x.list = append(x.list, v)")
		g.P("return ", pref.Ident("ValueOfMessage"), "(v.ProtoReflect())")
	default:
		panicMsg := fmt.Sprintf("AppendMutable can not be called on message %s at list field %s as it is not of Message kind", field.Parent.GoIdent.GoName, field.GoName)
		g.P("panic(", fmtPkg.Ident("Errorf"), "(\"", panicMsg, "\"))")
	}
	g.P("}")
	g.P()

	// Truncate
	g.P("func (x *", typeName, ") Truncate(n int)", "{")

	switch field.Desc.Kind() {
	case protoreflect.MessageKind: // zero message kinds to avoid keeping data alive
		g.P("for i := n; i < len(x.list); i++ {")
		g.P("x.list[i] = nil")
		g.P("}")
	}
	g.P("x.list = x.list[:n]") // truncate
	g.P("}")
	g.P()

	// NewElement
	g.P("func (x *", typeName, ") NewElement() ", pref.Ident("Value"), "{")
	zeroValue := zeroValueForField(g, field)
	g.P("v := ", zeroValue)
	switch field.Desc.Kind() {
	case protoreflect.MessageKind: // it can be mutable
		g.P("x.list = append(x.list, v)")
		g.P("return ", kindToValueConstructor(field.Desc.Kind()), "(v.ProtoReflect())")
	case protoreflect.EnumKind:
		g.P("return ", kindToValueConstructor(field.Desc.Kind()), "((", pref.Ident("EnumNumber"), ")(v))")
	default:
		g.P("return ", kindToValueConstructor(field.Desc.Kind()), "(v)")
	}
	g.P("}")
	g.P()

	// IsValid
	g.P("func (x *", typeName, ") IsValid() bool {")
	g.P("return x.list == nil || len(x.list) == 0") // TODO(fdymylja) len(list) validity??
	g.P("}")
	g.P()
}

func listTypeName(field *protogen.Field) string {
	return fmt.Sprintf("_%s_%d_list", field.Parent.GoIdent.GoName, field.Desc.Number())
}

func getType(g *protogen.GeneratedFile, field *protogen.Field) (goType string) {
	if field.Desc.IsWeak() {
		return "struct{}"
	}

	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
	}
	return goType
}

func kindToValueConstructor(kind protoreflect.Kind) protogen.GoIdent {
	const pref = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	switch kind {
	case protoreflect.BoolKind:
		return pref.Ident("ValueOfBool")
	case protoreflect.EnumKind:
		return pref.Ident("ValueOfEnum")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return pref.Ident("ValueOfInt32")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return pref.Ident("ValueOfUint32")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return pref.Ident("ValueOfInt64")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return pref.Ident("ValueOfUint64")
	case protoreflect.FloatKind:
		return pref.Ident("ValueOfFloat32")
	case protoreflect.DoubleKind:
		return pref.Ident("ValueOfFloat64")
	case protoreflect.StringKind:
		return pref.Ident("ValueOfString")
	case protoreflect.BytesKind:
		return pref.Ident("ValueOfBytes")
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return pref.Ident("ValueOfMessage")
	default:
		panic("should not reach here")
	}
}

// valueUnwrapper provides the function to call on value
// in order to get the concrete underlying type
func valueUnwrapper(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return "Bool"
	case protoreflect.EnumKind:
		return "Enum"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "Int"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "Uint"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "Int"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "Uint"
	case protoreflect.FloatKind:
		return "Float"
	case protoreflect.DoubleKind:
		return "Float"
	case protoreflect.StringKind:
		return "String"
	case protoreflect.BytesKind:
		return "Bytes"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return "Message"
	default:
		panic("should not reach here")
	}
}

func genPrefValueToGoValue(g *protogen.GeneratedFile, field *protogen.Field) string {
	const concreteValueName = "concreteValue"

	unwrapperFunc := valueUnwrapper(field.Desc.Kind())
	g.P("unwrapped := value.", unwrapperFunc, "()")
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		g.P(concreteValueName, " := unwrapped.Interface().(*", g.QualifiedGoIdent(field.Message.GoIdent), ")")
	case protoreflect.EnumKind:
		g.P(concreteValueName, " := (", g.QualifiedGoIdent(field.Enum.GoIdent), ")(unwrapped)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		g.P(concreteValueName, " := (int32)(unwrapped)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		g.P(concreteValueName, " := (uint32)(unwrapped)")
	case protoreflect.FloatKind:
		g.P(concreteValueName, " := (float32)(unwrapped)")
	default:
		g.P(concreteValueName, " := unwrapped")
	}

	return concreteValueName
}

func zeroValueForField(g *protogen.GeneratedFile, field *protogen.Field) string {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.EnumKind:
		return fmt.Sprintf("%d", field.Enum.Desc.Values().Get(0).Number())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32(0)"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32(0)"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64(0)"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64(0)"
	case protoreflect.FloatKind:
		return "float32(0)"
	case protoreflect.DoubleKind:
		return "float64(0)"
	case protoreflect.StringKind:
		return "\"\""
	case protoreflect.BytesKind:
		return "nil"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return fmt.Sprintf("new(%s)", g.QualifiedGoIdent(field.Message.GoIdent))
	default:
		panic("should not reach")
	}
}
