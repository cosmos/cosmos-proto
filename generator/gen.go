package generator

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"math"
	"runtime/debug"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

/*
	╭━━━┳╮╱╭┳╮╱╱╭━━━┳━━━┳━━━╮
	┃╭━╮┃┃╱┃┃┃╱╱┃╭━╮┃╭━╮┃╭━╮┃
	┃╰━╯┃┃╱┃┃┃╱╱┃╰━━┫┃╱┃┃╰━╯┃
	┃╭━━┫┃╱┃┃┃╱╭╋━━╮┃╰━╯┃╭╮╭╯
	┃┃╱╱┃╰━╯┃╰━╯┃╰━╯┃╭━╮┃┃┃╰╮
	╰╯╱╱╰━━━┻━━━┻━━━┻╯╱╰┻╯╰━╯

	- many bytes, such speed. -
*/

var gotrackTags = structTags{{"go", "track"}}

type structTags [][2]string

func (tags structTags) String() string {
	if len(tags) == 0 {
		return ""
	}
	var ss []string
	for _, tag := range tags {
		// NOTE: When quoting the value, we need to make sure the backtick
		// character does not appear. Convert all cases to the escaped hex form.
		key := tag[0]
		val := strings.Replace(strconv.Quote(tag[1]), "`", `\x60`, -1)
		ss = append(ss, fmt.Sprintf("%s:%s", key, val))
	}
	return "`" + strings.Join(ss, " ") + "`"
}

// Protobuf library dependencies.
//
// These are declared as an interface type so that they can be more easily
// patched to support unique build environments that impose restrictions
// on the dependencies of generated source code.
var (
	protoPackage         goImportPath = protogen.GoImportPath("google.golang.org/protobuf/proto")
	protoifacePackage    goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")
	protoimplPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")
	protojsonPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	protoreflectPackage  goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	protoregistryPackage goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoregistry")
	errorsPackage        goImportPath = protogen.GoImportPath("errors")
)

type goImportPath interface {
	String() string
	Ident(string) protogen.GoIdent
}

func GenerateProtocGenGo(plugin *protogen.Plugin, g *GeneratedFile, file *protogen.File) *GeneratedFile {
	f := newFileInfo(file)
	genPackage(g, file.GoPackageName, file)
	genTopLevelVars(g, file.Messages)
	for _, enum := range f.allEnums {
		genEnum(g, f, enum)
	}
	for _, msg := range f.allMessages {
		genMessage(g, f, msg)
		g.P()
	}
	// genFileProtoTypes(g, file)
	genReflectFileDescriptor(plugin, g, f)
	return g
}

// trailingComment is like protogen.Comments, but lacks a trailing newline.
type trailingComment protogen.Comments

func genEnum(g *GeneratedFile, f *fileInfo, e *enumInfo) {
	// Enum type declaration.
	g.Annotate(e.GoIdent.GoName, e.Location)
	leadingComments := appendDeprecationSuffix(e.Comments.Leading,
		e.Desc.Options().(*descriptorpb.EnumOptions).GetDeprecated())
	g.P(leadingComments,
		"type ", e.GoIdent, " int32")

	// Enum value constants.
	g.P("const (")
	for _, value := range e.Values {
		g.Annotate(value.GoIdent.GoName, value.Location)
		leadingComments := appendDeprecationSuffix(value.Comments.Leading,
			value.Desc.Options().(*descriptorpb.EnumValueOptions).GetDeprecated())
		g.P(leadingComments,
			value.GoIdent, " ", e.GoIdent, " = ", value.Desc.Number(),
			trailingComment(value.Comments.Trailing))
	}
	g.P(")")
	g.P()

	// Enum value maps.
	g.P("// Enum value maps for ", e.GoIdent, ".")
	g.P("var (")
	g.P(e.GoIdent.GoName+"_name", " = map[int32]string{")
	for _, value := range e.Values {
		duplicate := ""
		if value.Desc != e.Desc.Values().ByNumber(value.Desc.Number()) {
			duplicate = "// Duplicate value: "
		}
		g.P(duplicate, value.Desc.Number(), ": ", strconv.Quote(string(value.Desc.Name())), ",")
	}
	g.P("}")
	g.P(e.GoIdent.GoName+"_value", " = map[string]int32{")
	for _, value := range e.Values {
		g.P(strconv.Quote(string(value.Desc.Name())), ": ", value.Desc.Number(), ",")
	}
	g.P("}")
	g.P(")")
	g.P()

	// Enum method.
	//
	// NOTE: A pointer value is needed to represent presence in proto2.
	// Since a proto2 message can reference a proto3 enum, it is useful to
	// always generate this method (even on proto3 enums) to support that case.
	g.P("func (x ", e.GoIdent, ") Enum() *", e.GoIdent, " {")
	g.P("p := new(", e.GoIdent, ")")
	g.P("*p = x")
	g.P("return p")
	g.P("}")
	g.P()

	// String method.
	g.P("func (x ", e.GoIdent, ") String() string {")
	g.P("return ", protoimplPackage.Ident("X"), ".EnumStringOf(x.Descriptor(), ", protoreflectPackage.Ident("EnumNumber"), "(x))")
	g.P("}")
	g.P()

	genEnumReflectMethods(g, f, e)

	// UnmarshalJSON method.
	if e.genJSONMethod && e.Desc.Syntax() == protoreflect.Proto2 {
		g.P("// Deprecated: Do not use.")
		g.P("func (x *", e.GoIdent, ") UnmarshalJSON(b []byte) error {")
		g.P("num, err := ", protoimplPackage.Ident("X"), ".UnmarshalJSONEnum(x.Descriptor(), b)")
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P("*x = ", e.GoIdent, "(num)")
		g.P("return nil")
		g.P("}")
		g.P()
	}

	// EnumDescriptor method.
	if e.genRawDescMethod {
		var indexes []string
		for i := 1; i < len(e.Location.Path); i += 2 {
			indexes = append(indexes, strconv.Itoa(int(e.Location.Path[i])))
		}
		g.P("// Deprecated: Use ", e.GoIdent, ".Descriptor instead.")
		g.P("func (", e.GoIdent, ") EnumDescriptor() ([]byte, []int) {")
		g.P("return ", rawDescVarName(f), "GZIP(), []int{", strings.Join(indexes, ","), "}")
		g.P("}")
		g.P()
		f.needRawDesc = true
	}
}

func genEnumReflectMethods(g *GeneratedFile, f *fileInfo, e *enumInfo) {
	idx := f.allEnumsByPtr[e]
	typesVar := enumTypesVarName(f)

	// Descriptor method.
	g.P("func (", e.GoIdent, ") Descriptor() ", protoreflectPackage.Ident("EnumDescriptor"), " {")
	g.P("return ", typesVar, "[", idx, "].Descriptor()")
	g.P("}")
	g.P()

	// Type method.
	g.P("func (", e.GoIdent, ") Type() ", protoreflectPackage.Ident("EnumType"), " {")
	g.P("return &", typesVar, "[", idx, "]")
	g.P("}")
	g.P()

	// Number method.
	g.P("func (x ", e.GoIdent, ") Number() ", protoreflectPackage.Ident("EnumNumber"), " {")
	g.P("return ", protoreflectPackage.Ident("EnumNumber"), "(x)")
	g.P("}")
	g.P()
}

func genPackage(g *GeneratedFile, packageName protogen.GoPackageName, file *protogen.File) {
	g.P("// Code generated by Pulsar \U0001FA90. DO NOT EDIT.")
	if bi, ok := debug.ReadBuildInfo(); ok {
		g.P("// Pulsar version: ", bi.Main.Version)
	}
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", packageName)
	g.P()
}

func genTopLevelVars(g *GeneratedFile, msgs []*protogen.Message) {
	g.P("const (")
	g.P("// Verify that this generated code is sufficiently up-to-date.")
	g.P("_ = ", protoimplPackage.Ident("EnforceVersion"), "(", protoimpl.GenVersion, " - ", protoimplPackage.Ident("MinVersion"), ")")
	g.P("// Verify that runtime/protoimpl is sufficiently up-to-date.")
	g.P("_ = ", protoimplPackage.Ident("EnforceVersion"), "(", protoimplPackage.Ident("MaxVersion"), " - ", protoimpl.GenVersion, ")")
	g.P()
	g.P(")")
	g.P()
	g.P("var (")
	g.P("// Interface guards to verify each message implements proto message interface")
	for _, msg := range msgs {
		g.P("_ ", protoreflectPackage.Ident("Message"), " = &", msg.GoIdent.GoName, "{}")
	}
	g.P(")")
	g.P()
}

func genMessage(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	if m.Desc.IsMapEntry() {
		return
	}

	// Message type declaration.
	g.Annotate(m.GoIdent.GoName, m.Location)
	leadingComments := appendDeprecationSuffix(m.Comments.Leading,
		m.Desc.Options().(*descriptorpb.MessageOptions).GetDeprecated())
	g.P(leadingComments,
		"type ", m.GoIdent, " struct {")
	genMessageFields(g, f, m)
	g.P("}")
	g.P()

	genMessageKnownFunctions(g, f, m)
	genMessageDefaultDecls(g, f, m)
	genMessageMethods(g, f, m)
	genMessageOneofWrapperTypes(g, f, m)
	genProtoMessageFunctions(g, m)
}

func genMessageMethods(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	genMessageBaseMethods(g, f, m)
	genMessageGetterMethods(g, f, m)
	genMessageSetterMethods(g, f, m)
}

func genMessageSetterMethods(g *GeneratedFile, _ *fileInfo, m *messageInfo) {
	for _, field := range m.Fields {
		if !field.Desc.IsWeak() {
			continue
		}

		genNoInterfacePragma(g, m.isTracked)

		g.Annotate(m.GoIdent.GoName+".Set"+field.GoName, field.Location)
		leadingComments := appendDeprecationSuffix("",
			field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated())
		g.P(leadingComments, "func (x *", m.GoIdent, ") Set", field.GoName, "(v ", protoPackage.Ident("Message"), ") {")
		g.P("var w *", protoimplPackage.Ident("WeakFields"))
		g.P("if x != nil {")
		g.P("w = &x.", WeakFields_goname)
		if m.isTracked {
			g.P("_ = x.", WeakFieldPrefix_goname+field.GoName)
		}
		g.P("}")
		g.P(protoimplPackage.Ident("X"), ".SetWeak(w, ", field.Desc.Number(), ", ", strconv.Quote(string(field.Message.Desc.FullName())), ", v)")
		g.P("}")
		g.P()
	}
}

func genMessageBaseMethods(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	// Reset method.
	g.P("func (x *", m.GoIdent, ") Reset() {")
	g.P("*x = ", m.GoIdent, "{}")
	g.P("if ", protoimplPackage.Ident("UnsafeEnabled"), " {")
	g.P("mi := &", messageTypesVarName(f), "[", f.allMessagesByPtr[m], "]")
	g.P("ms := ", protoimplPackage.Ident("X"), ".MessageStateOf(", protoimplPackage.Ident("Pointer"), "(x))")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("}")
	g.P()

	// String method.
	g.P("func (x *", m.GoIdent, ") String() string {")
	g.P("return ", protoimplPackage.Ident("X"), ".MessageStringOf(x)")
	g.P("}")
	g.P()

	// ProtoMessage method.
	g.P("func (*", m.GoIdent, ") ProtoMessage() {}")
	g.P()

	// ProtoReflect method.
	genMessageReflectMethods(g, f, m)

	// Descriptor method.
	if m.genRawDescMethod {
		var indexes []string
		for i := 1; i < len(m.Location.Path); i += 2 {
			indexes = append(indexes, strconv.Itoa(int(m.Location.Path[i])))
		}
		// override this - using our implementation on protoreflect.Message instead
		//
		//g.P("// Deprecated: Use ", m.GoIdent, ".ProtoReflect.Descriptor instead.")
		//g.P("func (*", m.GoIdent, ") Descriptor() ([]byte, []int) {")
		//g.P("return ", rawDescVarName(f), "GZIP(), []int{", strings.Join(indexes, ","), "}")
		//g.P("}")
		//g.P()
		f.needRawDesc = true
	}

	// ExtensionRangeArray method.
	extRanges := m.Desc.ExtensionRanges()
	if m.genExtRangeMethod && extRanges.Len() > 0 {
		protoExtRange := protoifacePackage.Ident("ExtensionRangeV1")
		extRangeVar := "extRange_" + m.GoIdent.GoName
		g.P("var ", extRangeVar, " = []", protoExtRange, " {")
		for i := 0; i < extRanges.Len(); i++ {
			r := extRanges.Get(i)
			g.P("{Start:", r[0], ", End:", r[1]-1 /* inclusive */, "},")
		}
		g.P("}")
		g.P()
		g.P("// Deprecated: Use ", m.GoIdent, ".ProtoReflect.Descriptor.ExtensionRanges instead.")
		g.P("func (*", m.GoIdent, ") ExtensionRangeArray() []", protoExtRange, " {")
		g.P("return ", extRangeVar)
		g.P("}")
		g.P()
	}
}

// genMessageDefaultDecls generates consts and vars holding the default
// values of fields.
func genMessageDefaultDecls(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	var consts, vars []string
	for _, field := range m.Fields {
		if !field.Desc.HasDefault() {
			continue
		}
		name := "Default_" + m.GoIdent.GoName + "_" + field.GoName
		goType, _ := fieldGoType(g, f, field)
		defVal := field.Desc.Default()
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			consts = append(consts, fmt.Sprintf("%s = %s(%q)", name, goType, defVal.String()))
		case protoreflect.BytesKind:
			vars = append(vars, fmt.Sprintf("%s = %s(%q)", name, goType, defVal.Bytes()))
		case protoreflect.EnumKind:
			idx := field.Desc.DefaultEnumValue().Index()
			val := field.Enum.Values[idx]
			consts = append(consts, fmt.Sprintf("%s = %s", name, g.QualifiedGoIdent(val.GoIdent)))
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			if f := defVal.Float(); math.IsNaN(f) || math.IsInf(f, 0) {
				var fn, arg string
				switch f := defVal.Float(); {
				case math.IsInf(f, -1):
					fn, arg = g.QualifiedGoIdent(mathPackage.Ident("Inf")), "-1"
				case math.IsInf(f, +1):
					fn, arg = g.QualifiedGoIdent(mathPackage.Ident("Inf")), "+1"
				case math.IsNaN(f):
					fn, arg = g.QualifiedGoIdent(mathPackage.Ident("NaN")), ""
				}
				vars = append(vars, fmt.Sprintf("%s = %s(%s(%s))", name, goType, fn, arg))
			} else {
				consts = append(consts, fmt.Sprintf("%s = %s(%v)", name, goType, f))
			}
		default:
			consts = append(consts, fmt.Sprintf("%s = %s(%v)", name, goType, defVal.Interface()))
		}
	}
	if len(consts) > 0 {
		g.P("// Default values for ", m.GoIdent, " fields.")
		g.P("const (")
		for _, s := range consts {
			g.P(s)
		}
		g.P(")")
	}
	if len(vars) > 0 {
		g.P("// Default values for ", m.GoIdent, " fields.")
		g.P("var (")
		for _, s := range vars {
			g.P(s)
		}
		g.P(")")
	}
	g.P()
}

func genMessageFields(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	sf := f.allMessageFieldsByPtr[m]
	genMessageInternalFields(g, f, m, sf)
	for _, field := range m.Fields {
		genMessageField(g, f, m, field, sf)
	}
}

func genMessageReflectMethods(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	idx := f.allMessagesByPtr[m]
	typesVar := messageTypesVarName(f)

	// ProtoReflect method.
	g.P("func (x *", m.GoIdent, ") ProtoReflect() ", protoreflectPackage.Ident("Message"), " {")
	g.P("mi := &", typesVar, "[", idx, "]")
	g.P("if ", protoimplPackage.Ident("UnsafeEnabled"), " && x != nil {")
	g.P("ms := ", protoimplPackage.Ident("X"), ".MessageStateOf(", protoimplPackage.Ident("Pointer"), "(x))")
	g.P("if ms.LoadMessageInfo() == nil {")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("return ms")
	g.P("}")
	g.P("return mi.MessageOf(x)")
	g.P("}")
	g.P()
}

func genMessageInternalFields(g *GeneratedFile, _ *fileInfo, m *messageInfo, sf *structFields) {
	g.P(State_goname, " ", protoimplPackage.Ident("MessageState"))
	sf.append(State_goname)
	g.P(SizeCache_goname, " ", protoimplPackage.Ident("SizeCache"))
	sf.append(SizeCache_goname)
	if m.hasWeak {
		g.P(WeakFields_goname, " ", protoimplPackage.Ident("WeakFields"))
		sf.append(WeakFields_goname)
	}
	g.P(UnknownFields_goname, " ", protoimplPackage.Ident("UnknownFields"))
	sf.append(UnknownFields_goname)
	if m.Desc.ExtensionRanges().Len() > 0 {
		g.P(ExtensionFields_goname, " ", protoimplPackage.Ident("ExtensionFields"))
		sf.append(ExtensionFields_goname)
	}
	if sf.count > 0 {
		g.P()
	}
}

// genMessageOneofWrapperTypes generates the oneof wrapper types and
// associates the types with the parent message type.
func genMessageOneofWrapperTypes(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	for _, oneof := range m.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}
		ifName := oneofInterfaceName(oneof)
		g.P("type ", ifName, " interface {")
		g.P(ifName, "()")
		g.P("}")
		g.P()
		for _, field := range oneof.Fields {
			g.Annotate(field.GoIdent.GoName, field.Location)
			g.Annotate(field.GoIdent.GoName+"."+field.GoName, field.Location)
			g.P("type ", field.GoIdent, " struct {")
			goType, _ := fieldGoType(g, f, field)
			tags := structTags{
				{"protobuf", fieldProtobufTagValue(field)},
			}
			if m.isTracked {
				tags = append(tags, gotrackTags...)
			}
			leadingComments := appendDeprecationSuffix(field.Comments.Leading,
				field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated())
			g.P(leadingComments,
				field.GoName, " ", goType, tags,
				trailingComment(field.Comments.Trailing))
			g.P("}")
			g.P()
		}
		for _, field := range oneof.Fields {
			g.P("func (*", field.GoIdent, ") ", ifName, "() {}")
			g.P()
		}
	}
}

func genMessageField(g *GeneratedFile, f *fileInfo, m *messageInfo, field *protogen.Field, sf *structFields) {
	if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() {
		// It would be a bit simpler to iterate over the oneofs below,
		// but generating the field here keeps the contents of the Go
		// struct in the same order as the contents of the source
		// .proto file.
		if oneof.Fields[0] != field {
			return // only generate for first appearance
		}

		tags := structTags{
			{"protobuf_oneof", string(oneof.Desc.Name())},
		}
		if m.isTracked {
			tags = append(tags, gotrackTags...)
		}

		g.Annotate(m.GoIdent.GoName+"."+oneof.GoName, oneof.Location)
		leadingComments := oneof.Comments.Leading
		if leadingComments != "" {
			leadingComments += "\n"
		}
		ss := []string{fmt.Sprintf(" Types that are assignable to %s:\n", oneof.GoName)}
		for _, field := range oneof.Fields {
			ss = append(ss, "\t*"+field.GoIdent.GoName+"\n")
		}
		leadingComments += protogen.Comments(strings.Join(ss, ""))
		g.P(leadingComments,
			oneof.GoName, " ", oneofInterfaceName(oneof), tags)
		sf.append(oneof.GoName)
		return
	}
	goType, pointer := fieldGoType(g, f, field)
	if pointer {
		goType = "*" + goType
	}
	tags := structTags{
		{"protobuf", fieldProtobufTagValue(field)},
		{"json", fieldJSONTagValue(field)},
	}
	if field.Desc.IsMap() {
		key := field.Message.Fields[0]
		val := field.Message.Fields[1]
		tags = append(tags, structTags{
			{"protobuf_key", fieldProtobufTagValue(key)},
			{"protobuf_val", fieldProtobufTagValue(val)},
		}...)
	}
	if m.isTracked {
		tags = append(tags, gotrackTags...)
	}

	name := field.GoName
	if field.Desc.IsWeak() {
		name = WeakFieldPrefix_goname + name
	}
	g.Annotate(m.GoIdent.GoName+"."+name, field.Location)
	leadingComments := appendDeprecationSuffix(field.Comments.Leading,
		field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated())
	g.P(leadingComments,
		name, " ", goType, tags,
		trailingComment(field.Comments.Trailing))
	sf.append(field.GoName)
}

func fieldProtobufTagValue(field *protogen.Field) string {
	var enumName string
	if field.Desc.Kind() == protoreflect.EnumKind {
		enumName = protoimpl.X.LegacyEnumName(field.Enum.Desc)
	}
	return Marshal(field.Desc, enumName)
}

func fieldJSONTagValue(field *protogen.Field) string {
	return string(field.Desc.Name()) + ",omitempty"
}

// genNoInterfacePragma generates a standalone "nointerface" pragma to
// decorate methods with field-tracking support.
func genNoInterfacePragma(g *GeneratedFile, tracked bool) {
	if tracked {
		g.P("//go:nointerface")
		g.P()
	}
}

// oneofInterfaceName returns the name of the interface type implemented by
// the oneof field value types.
func oneofInterfaceName(oneof *protogen.Oneof) string {
	return "is" + oneof.GoIdent.GoName
}

// fieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGoType(g *GeneratedFile, f *fileInfo, field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
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
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := fieldGoType(g, f, field.Message.Fields[0])
		valType, _ := fieldGoType(g, f, field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func fieldDefaultValue(g *GeneratedFile, m *messageInfo, field *protogen.Field) string {
	if field.Desc.IsList() {
		return "nil"
	}
	if field.Desc.HasDefault() {
		defVarName := "Default_" + m.GoIdent.GoName + "_" + field.GoName
		if field.Desc.Kind() == protoreflect.BytesKind {
			return "append([]byte(nil), " + defVarName + "...)"
		}
		return defVarName
	}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.StringKind:
		return `""`
	case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.BytesKind:
		return "nil"
	case protoreflect.EnumKind:
		return g.QualifiedGoIdent(field.Enum.Values[0].GoIdent)
	default:
		return "0"
	}
}

func genMessageGetterMethods(g *GeneratedFile, f *fileInfo, m *messageInfo) {
	for _, field := range m.Fields {
		genNoInterfacePragma(g, m.isTracked)

		// Getter for parent oneof.
		if oneof := field.Oneof; oneof != nil && oneof.Fields[0] == field && !oneof.Desc.IsSynthetic() {
			g.Annotate(m.GoIdent.GoName+".Get"+oneof.GoName, oneof.Location)
			g.P("func (m *", m.GoIdent.GoName, ") Get", oneof.GoName, "() ", oneofInterfaceName(oneof), " {")
			g.P("if m != nil {")
			g.P("return m.", oneof.GoName)
			g.P("}")
			g.P("return nil")
			g.P("}")
			g.P()
		}

		// Getter for message field.
		goType, pointer := fieldGoType(g, f, field)
		defaultValue := fieldDefaultValue(g, m, field)
		g.Annotate(m.GoIdent.GoName+".Get"+field.GoName, field.Location)
		leadingComments := appendDeprecationSuffix("",
			field.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated())
		switch {
		case field.Desc.IsWeak():
			g.P(leadingComments, "func (x *", m.GoIdent, ") Get", field.GoName, "() ", protoPackage.Ident("Message"), "{")
			g.P("var w ", protoimplPackage.Ident("WeakFields"))
			g.P("if x != nil {")
			g.P("w = x.", WeakFields_goname)
			if m.isTracked {
				g.P("_ = x.", WeakFieldPrefix_goname+field.GoName)
			}
			g.P("}")
			g.P("return ", protoimplPackage.Ident("X"), ".GetWeak(w, ", field.Desc.Number(), ", ", strconv.Quote(string(field.Message.Desc.FullName())), ")")
			g.P("}")
		case field.Oneof != nil && !field.Oneof.Desc.IsSynthetic():
			g.P(leadingComments, "func (x *", m.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			g.P("if x, ok := x.Get", field.Oneof.GoName, "().(*", field.GoIdent, "); ok {")
			g.P("return x.", field.GoName)
			g.P("}")
			g.P("return ", defaultValue)
			g.P("}")
		default:
			g.P(leadingComments, "func (x *", m.GoIdent, ") Get", field.GoName, "() ", goType, " {")
			if !field.Desc.HasPresence() || defaultValue == "nil" {
				g.P("if x != nil {")
			} else {
				g.P("if x != nil && x.", field.GoName, " != nil {")
			}
			star := ""
			if pointer {
				star = "*"
			}
			g.P("return ", star, " x.", field.GoName)
			g.P("}")
			g.P("return ", defaultValue)
			g.P("}")
		}
		g.P()
	}
}

// Format is the serialization format used to represent the default value.
type Format int

const (
	_ Format = iota

	// GoTag uses the historical serialization format in Go struct field tags.
	GoTag
)

// Marshal encodes the protoreflect.FieldDescriptor as a tag.
//
// The enumName must be provided if the kind is an enum.
// Historically, the formulation of the enum "name" was the proto package
// dot-concatenated with the generated Go identifier for the enum type.
// Depending on the context on how Marshal is called, there are different ways
// through which that information is determined. As such it is the caller's
// responsibility to provide a function to obtain that information.
func Marshal(fd protoreflect.FieldDescriptor, enumName string) string {
	var tag []string
	switch fd.Kind() {
	case protoreflect.BoolKind, protoreflect.EnumKind, protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Int64Kind, protoreflect.Uint64Kind:
		tag = append(tag, "varint")
	case protoreflect.Sint32Kind:
		tag = append(tag, "zigzag32")
	case protoreflect.Sint64Kind:
		tag = append(tag, "zigzag64")
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.FloatKind:
		tag = append(tag, "fixed32")
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.DoubleKind:
		tag = append(tag, "fixed64")
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind:
		tag = append(tag, "bytes")
	case protoreflect.GroupKind:
		tag = append(tag, "group")
	}
	tag = append(tag, strconv.Itoa(int(fd.Number())))
	switch fd.Cardinality() {
	case protoreflect.Optional:
		tag = append(tag, "opt")
	case protoreflect.Required:
		tag = append(tag, "req")
	case protoreflect.Repeated:
		tag = append(tag, "rep")
	}
	if fd.IsPacked() {
		tag = append(tag, "packed")
	}
	name := string(fd.Name())
	if fd.Kind() == protoreflect.GroupKind {
		// The name of the FieldDescriptor for a group field is
		// lowercased. To find the original capitalization, we
		// look in the field's MessageType.
		name = string(fd.Message().Name())
	}
	tag = append(tag, "name="+name)
	if jsonName := fd.JSONName(); jsonName != "" && jsonName != name && !fd.IsExtension() {
		// NOTE: The jsonName != name condition is suspect, but it preserve
		// the exact same semantics from the previous generator.
		tag = append(tag, "json="+jsonName)
	}
	if fd.IsWeak() {
		tag = append(tag, "weak="+string(fd.Message().FullName()))
	}
	// The previous implementation does not tag extension fields as proto3,
	// even when the field is defined in a proto3 file. Match that behavior
	// for consistency.
	if fd.Syntax() == protoreflect.Proto3 && !fd.IsExtension() {
		tag = append(tag, "proto3")
	}
	if fd.Kind() == protoreflect.EnumKind && enumName != "" {
		tag = append(tag, "enum="+enumName)
	}
	if fd.ContainingOneof() != nil {
		tag = append(tag, "oneof")
	}
	// This must appear last in the tag, since commas in strings aren't escaped.
	if fd.HasDefault() {
		def, _ := DefValMarshal(fd.Default(), fd.DefaultEnumValue(), fd.Kind(), GoTag)
		tag = append(tag, "def="+def)
	}
	return strings.Join(tag, ",")
}

// DefValMarshal serializes v as the default string according to the given kind k.
// When specifying the Descriptor format for an enum kind, the associated
// enum value descriptor must be provided.
func DefValMarshal(v protoreflect.Value, ev protoreflect.EnumValueDescriptor, k protoreflect.Kind, f Format) (string, error) {
	switch k {
	case protoreflect.BoolKind:
		if f == GoTag {
			if v.Bool() {
				return "1", nil
			} else {
				return "0", nil
			}
		} else {
			if v.Bool() {
				return "true", nil
			} else {
				return "false", nil
			}
		}
	case protoreflect.EnumKind:
		if f == GoTag {
			return strconv.FormatInt(int64(v.Enum()), 10), nil
		} else {
			return string(ev.Name()), nil
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return strconv.FormatInt(v.Int(), 10), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return strconv.FormatUint(v.Uint(), 10), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		f := v.Float()
		switch {
		case math.IsInf(f, -1):
			return "-inf", nil
		case math.IsInf(f, +1):
			return "inf", nil
		case math.IsNaN(f):
			return "nan", nil
		default:
			if k == protoreflect.FloatKind {
				return strconv.FormatFloat(f, 'g', -1, 32), nil
			} else {
				return strconv.FormatFloat(f, 'g', -1, 64), nil
			}
		}
	case protoreflect.StringKind:
		// String values are serialized as is without any escaping.
		return v.String(), nil
	case protoreflect.BytesKind:
		if s, ok := marshalBytes(v.Bytes()); ok {
			return s, nil
		}
	}
	return "", errors.New(fmt.Sprintf("could not format value for %v: %v", k, v))
}

// marshalBytes serializes bytes by using C escaping.
// To match the exact output of protoc, this is identical to the
// CEscape function in strutil.cc of the protoc source code.
func marshalBytes(b []byte) (string, bool) {
	var s []byte
	for _, c := range b {
		switch c {
		case '\n':
			s = append(s, `\n`...)
		case '\r':
			s = append(s, `\r`...)
		case '\t':
			s = append(s, `\t`...)
		case '"':
			s = append(s, `\"`...)
		case '\'':
			s = append(s, `\'`...)
		case '\\':
			s = append(s, `\\`...)
		default:
			if printableASCII := c >= 0x20 && c <= 0x7e; printableASCII {
				s = append(s, c)
			} else {
				s = append(s, fmt.Sprintf(`\%03o`, c)...)
			}
		}
	}
	return string(s), true
}

type fileInfo struct {
	*protogen.File

	allEnums      []*enumInfo
	allMessages   []*messageInfo
	allExtensions []*extensionInfo

	allEnumsByPtr         map[*enumInfo]int    // value is index into allEnums
	allMessagesByPtr      map[*messageInfo]int // value is index into allMessages
	allMessageFieldsByPtr map[*messageInfo]*structFields

	// needRawDesc specifies whether the generator should emit logic to provide
	// the legacy raw descriptor in GZIP'd form.
	// This is updated by enum and message generation logic as necessary,
	// and checked at the end of file generation.
	needRawDesc bool
}
type enumInfo struct {
	*protogen.Enum
	genJSONMethod    bool
	genRawDescMethod bool
}

type extensionInfo struct {
	*protogen.Extension
}

type structFields struct {
	count      int
	unexported map[int]string
}

func (sf *structFields) append(name string) {
	if r, _ := utf8.DecodeRuneInString(name); !unicode.IsUpper(r) {
		if sf.unexported == nil {
			sf.unexported = make(map[int]string)
		}
		sf.unexported[sf.count] = name
	}
	sf.count++
}

func fileVarName(f *protogen.File, suffix string) string {
	prefix := f.GoDescriptorIdent.GoName
	_, n := utf8.DecodeRuneInString(prefix)
	prefix = strings.ToLower(prefix[:n]) + prefix[n:]
	return prefix + "_" + suffix
}
func rawDescVarName(f *fileInfo) string {
	return fileVarName(f.File, "rawDesc")
}
func goTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "goTypes")
}
func depIdxsVarName(f *fileInfo) string {
	return fileVarName(f.File, "depIdxs")
}
func enumTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "enumTypes")
}
func messageTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "msgTypes")
}
func extensionTypesVarName(f *fileInfo) string {
	return fileVarName(f.File, "extTypes")
}
func initFuncName(f *protogen.File) string {
	return fileVarName(f, "init")
}

func genReflectFileDescriptor(gen *protogen.Plugin, g *GeneratedFile, f *fileInfo) {
	g.P("var ", f.GoDescriptorIdent, " ", protoreflectPackage.Ident("FileDescriptor"))
	g.P()

	genFileDescriptor(gen, g, f)
	if len(f.allEnums) > 0 {
		g.P("var ", enumTypesVarName(f), " = make([]", protoimplPackage.Ident("EnumInfo"), ",", len(f.allEnums), ")")
	}
	if len(f.allMessages) > 0 {
		g.P("var ", messageTypesVarName(f), " = make([]", protoimplPackage.Ident("MessageInfo"), ",", len(f.allMessages), ")")
	}

	// Generate a unique list of Go types for all declarations and dependencies,
	// and the associated index into the type list for all dependencies.
	var goTypes []string
	var depIdxs []string
	seen := map[protoreflect.FullName]int{}
	genDep := func(name protoreflect.FullName, depSource string) {
		if depSource != "" {
			line := fmt.Sprintf("%d, // %d: %s -> %s", seen[name], len(depIdxs), depSource, name)
			depIdxs = append(depIdxs, line)
		}
	}
	genEnum := func(e *protogen.Enum, depSource string) {
		if e != nil {
			name := e.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(%s)(0), // %d: %s", g.QualifiedGoIdent(e.GoIdent), len(goTypes), name)
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}
	genMessage := func(m *protogen.Message, depSource string) {
		if m != nil {
			name := m.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(*%s)(nil), // %d: %s", g.QualifiedGoIdent(m.GoIdent), len(goTypes), name)
				if m.Desc.IsMapEntry() {
					// Map entry messages have no associated Go type.
					line = fmt.Sprintf("nil, // %d: %s", len(goTypes), name)
				}
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}

	// This ordering is significant.
	// See filetype.TypeBuilder.DependencyIndexes.
	type offsetEntry struct {
		start int
		name  string
	}
	var depOffsets []offsetEntry
	for _, enum := range f.allEnums {
		genEnum(enum.Enum, "")
	}
	for _, message := range f.allMessages {
		genMessage(message.Message, "")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "field type_name"})
	for _, message := range f.allMessages {
		for _, field := range message.Fields {
			if field.Desc.IsWeak() {
				continue
			}
			source := string(field.Desc.FullName())
			genEnum(field.Enum, source+":type_name")
			genMessage(field.Message, source+":type_name")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "extension extendee"})
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genMessage(extension.Extendee, source+":extendee")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "extension type_name"})
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genEnum(extension.Enum, source+":type_name")
		genMessage(extension.Message, source+":type_name")
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "method input_type"})
	for _, service := range f.Services {
		for _, method := range service.Methods {
			source := string(method.Desc.FullName())
			genMessage(method.Input, source+":input_type")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), "method output_type"})
	for _, service := range f.Services {
		for _, method := range service.Methods {
			source := string(method.Desc.FullName())
			genMessage(method.Output, source+":output_type")
		}
	}
	depOffsets = append(depOffsets, offsetEntry{len(depIdxs), ""})
	for i := len(depOffsets) - 2; i >= 0; i-- {
		curr, next := depOffsets[i], depOffsets[i+1]
		depIdxs = append(depIdxs, fmt.Sprintf("%d, // [%d:%d] is the sub-list for %s",
			curr.start, curr.start, next.start, curr.name))
	}
	if len(depIdxs) > math.MaxInt32 {
		panic("too many dependencies") // sanity check
	}

	g.P("var ", goTypesVarName(f), " = []interface{}{")
	for _, s := range goTypes {
		g.P(s)
	}
	g.P("}")

	g.P("var ", depIdxsVarName(f), " = []int32{")
	for _, s := range depIdxs {
		g.P(s)
	}
	g.P("}")

	g.P("func init() { ", initFuncName(f.File), "() }")

	g.P("func ", initFuncName(f.File), "() {")
	g.P("if ", f.GoDescriptorIdent, " != nil {")
	g.P("return")
	g.P("}")

	// Ensure that initialization functions for different files in the same Go
	// package run in the correct order: Call the init funcs for every .proto file
	// imported by this one that is in the same Go package.
	for i, imps := 0, f.Desc.Imports(); i < imps.Len(); i++ {
		impFile := gen.FilesByPath[imps.Get(i).Path()]
		if impFile.GoImportPath != f.GoImportPath {
			continue
		}
		g.P(initFuncName(impFile), "()")
	}

	if len(f.allMessages) > 0 {
		// Populate MessageInfo.Exporters.
		g.P("if !", protoimplPackage.Ident("UnsafeEnabled"), " {")
		for _, message := range f.allMessages {
			if sf := f.allMessageFieldsByPtr[message]; len(sf.unexported) > 0 {
				idx := f.allMessagesByPtr[message]
				typesVar := messageTypesVarName(f)

				g.P(typesVar, "[", idx, "].Exporter = func(v interface{}, i int) interface{} {")
				g.P("switch v := v.(*", message.GoIdent, "); i {")
				for i := 0; i < sf.count; i++ {
					if name := sf.unexported[i]; name != "" {
						g.P("case ", i, ": return &v.", name)
					}
				}
				g.P("default: return nil")
				g.P("}")
				g.P("}")
			}
		}
		g.P("}")

		// Populate MessageInfo.OneofWrappers.
		for _, message := range f.allMessages {
			if len(message.Oneofs) > 0 {
				idx := f.allMessagesByPtr[message]
				typesVar := messageTypesVarName(f)

				// Associate the wrapper types by directly passing them to the MessageInfo.
				g.P(typesVar, "[", idx, "].OneofWrappers = []interface{} {")
				for _, oneof := range message.Oneofs {
					if !oneof.Desc.IsSynthetic() {
						for _, field := range oneof.Fields {
							g.P("(*", field.GoIdent, ")(nil),")
						}
					}
				}
				g.P("}")
			}
		}
	}

	g.P("type x struct{}")
	g.P("out := ", protoimplPackage.Ident("TypeBuilder"), "{")
	g.P("File: ", protoimplPackage.Ident("DescBuilder"), "{")
	g.P("GoPackagePath: ", reflectPackage.Ident("TypeOf"), "(x{}).PkgPath(),")
	g.P("RawDescriptor: ", rawDescVarName(f), ",")
	g.P("NumEnums: ", len(f.allEnums), ",")
	g.P("NumMessages: ", len(f.allMessages), ",")
	g.P("NumExtensions: ", len(f.allExtensions), ",")
	g.P("NumServices: ", len(f.Services), ",")
	g.P("},")
	g.P("GoTypes: ", goTypesVarName(f), ",")
	g.P("DependencyIndexes: ", depIdxsVarName(f), ",")
	if len(f.allEnums) > 0 {
		g.P("EnumInfos: ", enumTypesVarName(f), ",")
	}
	if len(f.allMessages) > 0 {
		g.P("MessageInfos: ", messageTypesVarName(f), ",")
	}
	if len(f.allExtensions) > 0 {
		g.P("ExtensionInfos: ", extensionTypesVarName(f), ",")
	}
	g.P("}.Build()")
	g.P(f.GoDescriptorIdent, " = out.File")

	// Set inputs to nil to allow GC to reclaim resources.
	g.P(rawDescVarName(f), " = nil")
	g.P(goTypesVarName(f), " = nil")
	g.P(depIdxsVarName(f), " = nil")
	g.P("}")
}

func genFileDescriptor(gen *protogen.Plugin, g *GeneratedFile, f *fileInfo) {
	descProto := proto.Clone(f.Proto).(*descriptorpb.FileDescriptorProto)
	descProto.SourceCodeInfo = nil // drop source code information

	b, err := proto.MarshalOptions{AllowPartial: true, Deterministic: true}.Marshal(descProto)
	if err != nil {
		gen.Error(err)
		return
	}

	g.P("var ", rawDescVarName(f), " = []byte{")
	for len(b) > 0 {
		n := 16
		if n > len(b) {
			n = len(b)
		}

		s := ""
		for _, c := range b[:n] {
			s += fmt.Sprintf("0x%02x,", c)
		}
		g.P(s)

		b = b[n:]
	}
	g.P("}")
	g.P()

	if f.needRawDesc {
		onceVar := rawDescVarName(f) + "Once"
		dataVar := rawDescVarName(f) + "Data"
		g.P("var (")
		g.P(onceVar, " ", syncPackage.Ident("Once"))
		g.P(dataVar, " = ", rawDescVarName(f))
		g.P(")")
		g.P()

		g.P("func ", rawDescVarName(f), "GZIP() []byte {")
		g.P(onceVar, ".Do(func() {")
		g.P(dataVar, " = ", protoimplPackage.Ident("X"), ".CompressGZIP(", dataVar, ")")
		g.P("})")
		g.P("return ", dataVar)
		g.P("}")
		g.P()
	}
}

func newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}

	// Collect all enums, messages, and extensions in "flattened ordering".
	// See filetype.TypeBuilder.
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}
	initEnumInfos := func(enums []*protogen.Enum) {
		for _, enum := range enums {
			f.allEnums = append(f.allEnums, newEnumInfo(f, enum))
		}
	}
	initMessageInfos := func(messages []*protogen.Message) {
		for _, message := range messages {
			f.allMessages = append(f.allMessages, newMessageInfo(f, message))
		}
	}
	initExtensionInfos := func(extensions []*protogen.Extension) {
		for _, extension := range extensions {
			f.allExtensions = append(f.allExtensions, newExtensionInfo(f, extension))
		}
	}
	initEnumInfos(f.Enums)
	initMessageInfos(f.Messages)
	initExtensionInfos(f.Extensions)
	walkMessages(f.Messages, func(m *protogen.Message) {
		initEnumInfos(m.Enums)
		initMessageInfos(m.Messages)
		initExtensionInfos(m.Extensions)
	})

	// Derive a reverse mapping of enum and message pointers to their index
	// in allEnums and allMessages.
	if len(f.allEnums) > 0 {
		f.allEnumsByPtr = make(map[*enumInfo]int)
		for i, e := range f.allEnums {
			f.allEnumsByPtr[e] = i
		}
	}
	if len(f.allMessages) > 0 {
		f.allMessagesByPtr = make(map[*messageInfo]int)
		f.allMessageFieldsByPtr = make(map[*messageInfo]*structFields)
		for i, m := range f.allMessages {
			f.allMessagesByPtr[m] = i
			f.allMessageFieldsByPtr[m] = new(structFields)
		}
	}

	return f
}

func newEnumInfo(_ *fileInfo, enum *protogen.Enum) *enumInfo {
	e := &enumInfo{Enum: enum}
	e.genJSONMethod = true
	e.genRawDescMethod = true
	return e
}

func newMessageInfo(_ *fileInfo, message *protogen.Message) *messageInfo {
	m := &messageInfo{Message: message}
	m.genRawDescMethod = true
	m.genExtRangeMethod = true
	m.isTracked = isTrackedMessage(m)
	for _, field := range m.Fields {
		m.hasWeak = m.hasWeak || field.Desc.IsWeak()
	}
	return m
}

// isTrackedMessage reports whether field tracking is enabled on the message.
func isTrackedMessage(m *messageInfo) (tracked bool) {
	const trackFieldUse_fieldNumber = 37383685

	// Decode the option from unknown fields to avoid a dependency on the
	// annotation proto from protoc-gen-go.
	b := m.Desc.Options().(*descriptorpb.MessageOptions).ProtoReflect().GetUnknown()
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		b = b[n:]
		if num == trackFieldUse_fieldNumber && typ == protowire.VarintType {
			v, _ := protowire.ConsumeVarint(b)
			tracked = protowire.DecodeBool(v)
		}
		m := protowire.ConsumeFieldValue(num, typ, b)
		b = b[m:]
	}
	return tracked
}

func newExtensionInfo(_ *fileInfo, extension *protogen.Extension) *extensionInfo {
	x := &extensionInfo{Extension: extension}
	return x
}

// appendDeprecationSuffix optionally appends a deprecation notice as a suffix.
func appendDeprecationSuffix(prefix protogen.Comments, deprecated bool) protogen.Comments {
	if !deprecated {
		return prefix
	}
	if prefix != "" {
		prefix += "\n"
	}
	return prefix + " Deprecated: Do not use.\n"
}

type messageInfo struct {
	*protogen.Message

	genRawDescMethod  bool
	genExtRangeMethod bool

	isTracked bool
	hasWeak   bool
}
