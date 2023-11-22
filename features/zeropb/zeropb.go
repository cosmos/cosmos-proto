package zeropb

import (
	"github.com/cosmos/cosmos-proto/features/protoc"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	errorsPackage = protogen.GoImportPath("errors")
	mathPackage   = protogen.GoImportPath("math")
	binaryPackage = protogen.GoImportPath("encoding/binary")
)

func init() {
	generator.RegisterFeature("zeropb", func(gen *generator.GeneratedFile, _ *protogen.Plugin) generator.FeatureGenerator {
		return zeropbFeature{
			gen: gen,
		}
	})
}

type zeropbFeature struct {
	gen *generator.GeneratedFile
}

func (g zeropbFeature) GenerateFile(file *protogen.File, _ *protogen.Plugin) bool {
	for _, m := range file.Messages {
		g.generateMessage(file, m)
	}
	return true // only do this once
}

func (g zeropbFeature) GenerateHelpers() {}

func (g zeropbFeature) generateMessage(f *protogen.File, m *protogen.Message) {
	g.generateMarshal(f, m)
	g.generateUnmarshal(f, m)
}

func (g zeropbFeature) generateMarshal(f *protogen.File, m *protogen.Message) {
	g.gen.P("func (x *", m.GoIdent, ") MarshalZeroPB(buf []byte) (n int, err error) {")
	g.gen.P("defer func() {")
	g.gen.P("    if e := recover(); e != nil {")
	g.gen.P("        err = ", errorsPackage.Ident("New"), `("buffer overflow")`)
	g.gen.P("    }")
	g.gen.P("}()")
	for _, f := range m.Fields {
		g.generateMarshalField(f)
	}
	g.gen.P("return n, nil")
	g.gen.P("}")
}

func (g zeropbFeature) generateMarshalField(f *protogen.Field) {
	d := f.Desc
	switch {
	case d.IsList():
		g.gen.P("len_", d.Index(), " := uint16(len(x.", f.GoName, "))")
		g.gen.P("if len(x.", f.GoName, ") != int(len_", d.Index(), ") {")
		g.gen.P("    return n, ", errorsPackage.Ident("New"), `("field `, f.GoName, ` is too long")`)
		g.gen.P("}")
		g.gen.P(binaryPackage.Ident("LittleEndian"), ".PutUint16(buf[n:], len_", d.Index(), ")")
		g.gen.P("n += 2")
		g.gen.P("for _, e := range x.", f.GoName, " {")
		g.generateMarshalPrimitive(d, "e")
		g.gen.P("}")
	case d.IsMap():
		g.gen.P("len_", d.Index(), " := uint16(len(x.", f.GoName, "))")
		g.gen.P("if len(x.", f.GoName, ") != int(len_", d.Index(), ") {")
		g.gen.P("    return n, ", errorsPackage.Ident("New"), `("field `, f.GoName, ` is too long")`)
		g.gen.P("}")
		g.gen.P("binary.LittleEndian.PutUint16(buf[n:], len_", d.Index(), ")")
		g.gen.P("n += 2")
		g.gen.P("for k, v := range x.", f.GoName, " {")
		g.generateMarshalPrimitive(d.MapKey(), "k")
		g.generateMarshalPrimitive(d.MapValue(), "v")
		g.gen.P("}")
	case d.ContainingOneof() != nil:
		g.gen.P("// TODO: field ", f.GoName)
		return
	default:
		g.generateMarshalPrimitive(d, "x."+f.GoName)
	}
}

func (g zeropbFeature) generateMarshalPrimitive(d protoreflect.FieldDescriptor, name string) {
	switch d.Kind() {
	case protoreflect.FloatKind:
		g.gen.P("binary.LittleEndian.PutUint32(buf[n:], ", mathPackage.Ident("Float32bits"), "(", name, "))")
		g.gen.P("n += 4")
	case protoreflect.DoubleKind:
		g.gen.P("binary.LittleEndian.PutUint64(buf[n:], ", mathPackage.Ident("Float64bits"), "(", name, "))")
		g.gen.P("n += 8")
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.EnumKind:
		g.gen.P("binary.LittleEndian.PutUint32(buf[n:], uint32(", name, "))")
		g.gen.P("n += 4")
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind:
		g.gen.P("binary.LittleEndian.PutUint64(buf[n:], uint64(", name, "))")
		g.gen.P("n += 8")
	case protoreflect.BoolKind:
		g.gen.P("bool_", d.Index(), " := uint32(0)")
		g.gen.P("if ", name, " {")
		g.gen.P("    bool_", d.Index(), " = 1")
		g.gen.P("}")
		g.gen.P("binary.LittleEndian.PutUint32(buf[n:], bool_", d.Index(), ")")
		g.gen.P("n += 4")
	case protoreflect.StringKind, protoreflect.BytesKind:
		g.gen.P("len_", d.Index(), " := uint16(len(", name, "))")
		g.gen.P("if len(", name, ") != int(len_", d.Index(), ") {")
		g.gen.P("    return n, ", errorsPackage.Ident("New"), `("field `, name, ` is too long")`)
		g.gen.P("}")
		g.gen.P("binary.LittleEndian.PutUint16(buf[n:], len_", d.Index(), ")")
		g.gen.P("n += 2")
		// Reslice buf to convert a truncated write into a buffer overflow error.
		g.gen.P("copy(buf[n:n+len(", name, ")], ", name, ")")
		g.gen.P("n += len(", name, ")")
	case protoreflect.MessageKind:
		g.gen.P("n_", d.Index(), ", err := ", name, ".MarshalZeroPB(buf[n:])")
		g.gen.P("n += n_", d.Index())
		g.gen.P("if err != nil {")
		g.gen.P("    return n, err")
		g.gen.P("}")
	default:
		g.gen.P("// TODO: field ", name)
		g.gen.P("_ = ", name)
	}
}

func (g zeropbFeature) generateUnmarshal(f *protogen.File, m *protogen.Message) {
	g.gen.P("func (x *", m.GoIdent, ") UnmarshalZeroPB(buf []byte) (n int, err error) {")
	g.gen.P("defer func() {")
	g.gen.P("    if e := recover(); e != nil {")
	g.gen.P("        err = ", errorsPackage.Ident("New"), `("buffer underflow")`)
	g.gen.P("    }")
	g.gen.P("}()")
	for _, f := range m.Fields {
		g.generateUnmarshalField(f)
	}
	g.gen.P("return n, nil")
	g.gen.P("}")
}

func (g zeropbFeature) generateUnmarshalField(f *protogen.Field) {
	d := f.Desc
	switch {
	case d.IsList():
		g.gen.P("len_", d.Index(), " := int(binary.LittleEndian.Uint16(buf[n:]))")
		g.gen.P("n += 2")
		typ, pointer := protoc.FieldGoType(g.gen, f)
		if pointer {
			typ = "*" + typ
		}
		g.gen.P("x.", f.GoName, " = make(", typ, ", len_", d.Index(), ")")
		g.gen.P("for i := range x.", f.GoName, "{")
		g.generateUnmarshalPrimitive(f, "x."+f.GoName+"[i]")
		g.gen.P("}")
	case d.IsMap():
		g.gen.P("len_", d.Index(), " := int(", binaryPackage.Ident("LittleEndian"), ".Uint16(buf[n:]))")
		g.gen.P("n += 2")
		typ, _ := protoc.FieldGoType(g.gen, f)
		g.gen.P("x.", f.GoName, " = make(", typ, ", len_", d.Index(), ")")
		keyType, _ := protoc.FieldGoType(g.gen, f.Message.Fields[0])
		valType, _ := protoc.FieldGoType(g.gen, f.Message.Fields[1])
		g.gen.P("for i := 0; i < len_", d.Index(), "; i++ {")
		g.gen.P("var k ", keyType)
		g.gen.P("var v ", valType)
		g.generateUnmarshalPrimitive(f.Message.Fields[0], "k")
		g.generateUnmarshalPrimitive(f.Message.Fields[1], "v")
		g.gen.P("    x.", f.GoName, "[k] = v")
		g.gen.P("}")
	case d.ContainingOneof() != nil:
		g.gen.P("// TODO: field ", f.GoName)
	default:
		g.generateUnmarshalPrimitive(f, "x."+f.GoName)
	}
}

func (g zeropbFeature) generateUnmarshalPrimitive(f *protogen.Field, name string) {
	switch d := f.Desc; d.Kind() {
	case protoreflect.FloatKind:
		g.gen.P(name, " = float32(", mathPackage.Ident("Float32frombits"), "(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[n:])))")
		g.gen.P("n += 4")
	case protoreflect.DoubleKind:
		g.gen.P(name, " = float64(", mathPackage.Ident("Float64frombits"), "(", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[n:])))")
		g.gen.P("n += 8")
	case protoreflect.Sfixed32Kind, protoreflect.Int32Kind, protoreflect.Sint32Kind:
		g.gen.P(name, " = int32(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[n:]))")
		g.gen.P("n += 4")
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		g.gen.P(name, " = ", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[n:])")
		g.gen.P("n += 4")
	case protoreflect.Sfixed64Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind:
		g.gen.P(name, " = int64(", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[n:]))")
		g.gen.P("n += 8")
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		g.gen.P(name, " = ", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[n:])")
		g.gen.P("n += 8")
	case protoreflect.EnumKind:
		typ := g.gen.QualifiedGoIdent(f.Enum.GoIdent)
		g.gen.P(name, " = ", typ, "(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[n:]))")
		g.gen.P("n += 4")
	case protoreflect.BoolKind:
		g.gen.P("bool_", d.Index(), " := ", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[n:])")
		g.gen.P(name, " = false")
		g.gen.P("if bool_", d.Index(), " != 0 {")
		g.gen.P("    ", name, " = true")
		g.gen.P("}")
		g.gen.P("n += 4")
	case protoreflect.StringKind:
		g.gen.P("len_", d.Index(), " := int(", binaryPackage.Ident("LittleEndian"), ".Uint16(buf[n:]))")
		g.gen.P("n += 2")
		g.gen.P(name, " = string(buf[n:n+len_", d.Index(), "])")
		g.gen.P("n += len_", d.Index())
	case protoreflect.BytesKind:
		g.gen.P("len_", d.Index(), " := int(", binaryPackage.Ident("LittleEndian"), ".Uint16(buf[n:]))")
		g.gen.P("n += 2")
		g.gen.P(name, " = append([]byte{}, buf[n:n+len_", d.Index(), "]...)")
		g.gen.P("n += len_", d.Index())
	case protoreflect.MessageKind:
		typ := g.gen.QualifiedGoIdent(f.Message.GoIdent)
		g.gen.P(name, " = new(", typ, ")")
		g.gen.P("n_", d.Index(), ", err := ", name, ".UnmarshalZeroPB(buf[n:])")
		g.gen.P("n += n_", d.Index())
		g.gen.P("if err != nil {")
		g.gen.P("    return n, err")
		g.gen.P("}")
	default:
		g.gen.P("// TODO: field ", name)
		g.gen.P("_ = ", name)
	}
}
