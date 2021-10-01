package interfaceservice

import (
	cosmos_proto "github.com/cosmos/cosmos-proto"
	"github.com/cosmos/cosmos-proto/generator"
	"github.com/cosmos/cosmos-proto/internal/strs"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	protoImportPath = protogen.GoImportPath("google.golang.org/protobuf/proto")
	anyImportPath   = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
)

type gen struct {
	*generator.GeneratedFile
	svc    *protogen.Service
	file   *protogen.File
	plugin *protogen.Plugin
	desc   *cosmos_proto.InterfaceService
}

func (g *gen) generate() error {
	if err := g.genInterface(); err != nil {
		return err
	}

	if err := g.genAnyInterface(); err != nil {
		return err
	}

	return nil
}

func (g *gen) genInterface() error {
	// generate the interface
	g.Import(protoImportPath)
	g.P("type ", interfaceServiceName(g.svc), " interface {")
	g.P(protoImportPath.Ident("Message"))
	// gen interface methods
	for _, method := range g.svc.Methods {
		if g.svc.Desc.ParentFile() != method.Input.Desc.ParentFile() {
			g.Import(method.Input.GoIdent.GoImportPath)
		}
		if g.svc.Desc.ParentFile() != method.Output.Desc.ParentFile() {
			g.Import(method.Output.GoIdent.GoImportPath)
		}
		// TODO: comments in a proper way
		if len(method.Comments.Leading) != 0 {
			g.P(method.Comments.Leading.String())
		}

		g.P(method.GoName, "(*", g.QualifiedGoIdent(method.Input.GoIdent), ") (*", g.QualifiedGoIdent(method.Output.GoIdent), ", error)")
	}
	g.P("}")

	g.P()
	return nil
}

func (g *gen) genAnyInterface() error {
	name := AnyInterfaceType(g.svc.Desc)

	// interface assertion
	g.P("var _ ", interfaceServiceName(g.svc), " = (*", name, ")(nil)")
	// type creation
	g.P("type ", name, " struct {")
	g.P("any *", anyImportPath.Ident("Any"))
	g.P("iface ", g.svc.GoName)
	g.P("}")
	g.P()
	// type implementation
	for _, method := range g.svc.Methods {
		if g.svc.Desc.ParentFile() != method.Input.Desc.ParentFile() {
			g.Import(method.Input.GoIdent.GoImportPath)
		}
		if g.svc.Desc.ParentFile() != method.Output.Desc.ParentFile() {
			g.Import(method.Output.GoIdent.GoImportPath)
		}
		// TODO: comments in a proper way
		if len(method.Comments.Leading) != 0 {
			g.P(method.Comments.Leading.String())
		}

		g.P("func (x *", name, ")", method.GoName, "(in *", g.QualifiedGoIdent(method.Input.GoIdent), ") (*", g.QualifiedGoIdent(method.Output.GoIdent), ", error) {")
		g.P("return x.iface.", method.GoName, "(in)")
		g.P("}")
		g.P()
	}
	g.P()
	// type constructor
	return nil
}

func interfaceServiceName(svc *protogen.Service) string {
	return svc.GoName
}

func AnyInterfaceType(desc protoreflect.ServiceDescriptor) string {
	return strs.GoCamelCase(string(desc.Name()) + "Any")
}
