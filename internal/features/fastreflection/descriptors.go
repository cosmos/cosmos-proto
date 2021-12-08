package fastreflection

import (
	"fmt"
	copied2 "github.com/cosmos/cosmos-proto/internal/features/fastreflection/copied"
	generator2 "github.com/cosmos/cosmos-proto/internal/generator"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/compiler/protogen"
)

type descGen struct {
	*generator2.GeneratedFile
	file    *protogen.File
	message *protogen.Message
}

func (g *descGen) generate() {
	// init is going to initialize descriptor information
	g.P("var (")
	g.P(messageDescriptorName(g.message), " ", protoreflectPkg.Ident("MessageDescriptor"))
	for _, field := range g.message.Fields {
		g.P(fieldDescriptorName(field), " ", protoreflectPkg.Ident("FieldDescriptor"))
	}
	g.P(")")
	g.P("func init() {")
	g.P(copied2.InitFunctionName(g.file), "()")
	parentMd, ok := g.message.Desc.Parent().(protoreflect.MessageDescriptor)
	switch ok {
	// case nested message
	case true:
		g.P(messageDescriptorName(g.message), " = ", g.file.GoDescriptorIdent.GoName, ".Messages().ByName(\"", parentMd.Name(), "\").Messages().ByName(\"", g.message.Desc.Name(), "\")")
	default:
		g.P(messageDescriptorName(g.message), " = ", g.file.GoDescriptorIdent.GoName, ".Messages().ByName(\"", g.message.Desc.Name(), "\")")
	}
	for _, field := range g.message.Fields {
		g.P(fieldDescriptorName(field), " = ", messageDescriptorName(g.message), ".Fields().ByName(\"", field.Desc.Name(), "\")")
	}
	g.P("}")
}

func fieldDescriptorName(field *protogen.Field) string {
	return fmt.Sprintf("fd_%s_%s", field.Parent.GoIdent.GoName, field.Desc.Name())
}
