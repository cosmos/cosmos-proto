package zpbgen

import "google.golang.org/protobuf/compiler/protogen"

func (g goGen) genFieldDescriptor(msgStructName string, field *protogen.Field) error {
	fdStructName := "fd_" + field.GoName

	g.P("type ", fdStructName, " struct {", protoreflectImport.Ident("FieldDescriptor"), "}")

	err := g.genFdHas(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	err = g.genFdGet(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	err = g.genFdSet(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	err = g.genFdClear(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	err = g.genFdNewField(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	err = g.genFdMutable(fdStructName, msgStructName, field)
	if err != nil {
		return err
	}

	return nil
}

func (g goGen) genFdHas(fdStructName, msgStructName string, field *protogen.Field) error {
	g.P("func (fd *", fdStructName, ") has(msg *", msgStructName, ") bool {")

	switch field.Desc.Kind() {

	}

	g.P("}")
	return nil
}

func (g goGen) genFdGet(fdStructName, msgStructName string, field *protogen.Field) error {
	g.P("func (fd *", fdStructName, ") get(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")
	return nil
}

func (g goGen) genFdSet(fdStructName, msgStructName string, field *protogen.Field) error {
	g.P("func (fd *", fdStructName, ") set(msg *", msgStructName, ", v ", protoreflectImport.Ident("Value"), ") {}")
	return nil
}

func (g goGen) genFdClear(fdStructName, msgStructName string, field *protogen.Field) error {
	g.P("func (fd *", fdStructName, ") clear(msg *", msgStructName, ") {}")
	return nil
}

func (g goGen) genFdMutable(fdStructName, msgStructName string, field *protogen.Field) error {
	return nil
}

func (g goGen) genFdNewField(fdStructName, msgStructName string, field *protogen.Field) error {
	g.P("func (fd *", fdStructName, ") newField(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")
	return nil
}
