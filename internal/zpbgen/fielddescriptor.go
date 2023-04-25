package zpbgen

import "google.golang.org/protobuf/compiler/protogen"

func (g goGen) genFieldDescriptor(msgStructName string, field *protogen.Field) error {
	fdStructName := "fd_" + field.GoName
	g.P("type ", fdStructName, " struct {", protoreflectImport.Ident("FieldDescriptor"), "}")
	g.P("func (fd *", fdStructName, ") has(msg *", msgStructName, ") bool { return false}")
	g.P("func (fd *", fdStructName, ") get(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")
	g.P("func (fd *", fdStructName, ") set(msg *", msgStructName, ", v ", protoreflectImport.Ident("Value"), ") {}")
	g.P("func (fd *", fdStructName, ") clear(msg *", msgStructName, ") {}")
	g.P("func (fd *", fdStructName, ") mutable(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")
	g.P("func (fd *", fdStructName, ") newField(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")

	return nil
}
