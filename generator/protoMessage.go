package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type goImportPath interface {
	String() string
	Ident(string) protogen.GoIdent
}

var (
	errorsPackage       goImportPath = protogen.GoImportPath("errors")
	protoifacePackage   goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")
	protoreflectPackage goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
)

type messageInfo struct {
	*protogen.Message

	genRawDescMethod  bool
	genExtRangeMethod bool

	isTracked bool
	hasWeak   bool
}

func genGetMethods(g *GeneratedFile, msg *messageInfo) {
	g.P("// returns the fast methods for the message")
	g.P("func (x ", msg.GoIdent.GoName, ") GetMethods() *", protoifacePackage.Ident("Methods"), " {")
	g.P("return &", protoifacePackage.Ident("Methods{"))
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Flags: 0,")

	// Size
	g.P("Size: func(input ", protoifacePackage.Ident("SizeInput"), ") ", protoifacePackage.Ident("SizeOutput"), " {")
	g.P("return ", protoifacePackage.Ident("SizeOutput"), "{")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Size: x.Size(),")
	g.P("}")
	g.P("},")

	// Marshal
	g.P("Marshal: func(input ", protoifacePackage.Ident("MarshalInput"), ") (", protoifacePackage.Ident("MarshalOutput"), ", error) {")
	g.P("v, ok := input.Message.(*", msg.GoIdent.GoName, ")")
	g.P("if !ok {")
	g.P("return ", protoifacePackage.Ident("MarshalOutput"), "{}, ", errorsPackage.Ident("New"), "(\"", msg.GoIdent.GoName, " does not implement the protoreflect.Message interface\")")
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

	// Unmarshal
	g.P("Unmarshal: func(input ", protoifacePackage.Ident("UnmarshalInput"), ") (", protoifacePackage.Ident("UnmarshalOutput"), ", error){")
	g.P("v, ok := input.Message.(*", msg.GoIdent.GoName, ")")
	g.P("if !ok {")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{}, ", errorsPackage.Ident("New"), "(\"", msg.GoIdent.GoName, " does not implement the protoreflect.Message interface\")")
	g.P("}")
	g.P()
	g.P("if len(input.Buf) < 1 {")
	g.P("return ", protoifacePackage.Ident("UnmarshalOutput"), "{}, ", errorsPackage.Ident("New"), "(\"unmarshal input did not contain any bytes to unmarshal\")")
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
