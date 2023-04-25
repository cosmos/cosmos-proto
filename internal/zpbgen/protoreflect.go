package zpbgen

import "google.golang.org/protobuf/compiler/protogen"

func (g goGen) genProtoReflect(msg *protogen.Message) error {
	msgStructName := "msg_" + msg.GoIdent.GoName
	mdStructName := "md_" + msg.GoIdent.GoName
	mdInstanceName := "md_" + msg.GoIdent.GoName + "_Instance"
	fdsInstanceName := "fds_" + msg.GoIdent.GoName + "_Instance"
	fdInterfaceName := "fd_" + msg.GoIdent.GoName + "_Interface"

	g.P("func (m *", msg.GoIdent, ") ProtoReflect() ", protoreflectImport.Ident("Message"), " {")
	g.P("return &", msgStructName, "{ctx: m.ctx}")
	g.P("}")
	g.P()

	g.P("type ", msgStructName, " struct {")
	g.P("ctx *", zeropbImport.Ident("BufferContext"))
	g.P("slowMsg ", protoreflectImport.Ident("Message"))
	g.P("}")

	g.P("func (m *", msgStructName, ") Descriptor() ", protoreflectImport.Ident("MessageDescriptor"), " {")
	g.P("return ", mdInstanceName)
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Type() ", protoreflectImport.Ident("Message"), "Type {")
	g.P("panic(\"TODO\")")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") New() ", protoreflectImport.Ident("Message"), " {")
	g.P("return &", msgStructName, "{ctx: m.ctx}")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Interface() ", protoreflectImport.Ident("ProtoMessage"), " {")
	g.P("return &", msg.GoIdent.GoName, "{ctx: m.ctx}")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Range(f func(", protoreflectImport.Ident("FieldDescriptor"), ", ", protoreflectImport.Ident("Value"), ") bool) {")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Has(fd ", protoreflectImport.Ident("FieldDescriptor"), ") bool {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("return f.has(m)")
	g.P("}")
	g.P("return m.slow().Has(fd)")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Get(fd ", protoreflectImport.Ident("FieldDescriptor"), ") ", protoreflectImport.Ident("Value"), " {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("return f.get(m)")
	g.P("}")
	g.P("return m.slow().Get(fd)")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Set(fd ", protoreflectImport.Ident("FieldDescriptor"), ", v ", protoreflectImport.Ident("Value"), ") {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("f.set(m, v)")
	g.P("} else {")
	g.P("m.slow().Set(fd, v)")
	g.P("}")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Clear(fd ", protoreflectImport.Ident("FieldDescriptor"), ") {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("f.clear(m)")
	g.P("} else {")
	g.P("m.slow().Clear(fd)")
	g.P("}")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") Mutable(fd ", protoreflectImport.Ident("FieldDescriptor"), ") ", protoreflectImport.Ident("Value"), " {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("return f.mutable(m)")
	g.P("}")
	g.P("return m.slow().Mutable(fd)")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") NewField(fd ", protoreflectImport.Ident("FieldDescriptor"), ") ", protoreflectImport.Ident("Value"), " {")
	g.P("if f, ok := fd.(", fdInterfaceName, "); ok {")
	g.P("return f.newField(m)")
	g.P("}")
	g.P("return m.slow().NewField(fd)")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") WhichOneof(od ", protoreflectImport.Ident("OneofDescriptor"), ") ", protoreflectImport.Ident("FieldDescriptor"), " {")
	g.P("panic(\"TODO\")")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") GetUnknown() ", protoreflectImport.Ident("RawFields"), " {")
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") SetUnknown(u ", protoreflectImport.Ident("RawFields"), ") {")
	g.P("panic(\"unsupported\")")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") IsValid() bool {")
	g.P("panic(\"TODO\")")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") ProtoMethods() *", protoifaceImport.Ident("Methods"), " {")
	g.P("return nil")
	g.P("}")
	g.P()

	g.P("func (m *", msgStructName, ") slow() ", protoreflectImport.Ident("Message"), " {")
	g.P("if m.slowMsg == nil {")
	g.P("panic(\"TODO\")")
	g.P("}")
	g.P("return m.slowMsg")
	g.P("}")
	g.P()

	g.P("type ", mdStructName, " struct {", protoreflectImport.Ident("MessageDescriptor"), "}")

	g.P("func (md *", mdStructName, ") Fields() ", protoreflectImport.Ident("FieldDescriptors"), " {")
	g.P("return ", fdsInstanceName)
	g.P("}")

	g.P("var ", mdInstanceName, " ", protoreflectImport.Ident("MessageDescriptor"))

	g.P("type ", fdInterfaceName, " interface {")
	g.P("has(*", msgStructName, ") bool")
	g.P("get(*", msgStructName, ") ", protoreflectImport.Ident("Value"))
	g.P("set(*", msgStructName, ", ", protoreflectImport.Ident("Value"), ")")
	g.P("clear(*", msgStructName, ")")
	g.P("mutable(*", msgStructName, ") ", protoreflectImport.Ident("Value"))
	g.P("newField(*", msgStructName, ") ", protoreflectImport.Ident("Value"))
	g.P("}")
	g.P()

	for _, field := range msg.Fields {
		err := g.genFieldDescriptor(msgStructName, field)
		if err != nil {
			return err
		}
	}

	g.P("var ", fdsInstanceName, " ", protoreflectImport.Ident("FieldDescriptors"))

	g.P()

	return nil
}
