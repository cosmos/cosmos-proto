package zpbgen

import "github.com/cosmos/cosmos-proto/zeropb"

func (g goGen) genFieldDescriptor(msgStructName string, field *zeropb.FieldDescriptor) error {
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

func (g goGen) genFdHas(fdStructName, msgStructName string, fd *zeropb.FieldDescriptor) error {
	g.P("func (fd *", fdStructName, ") has(msg *", msgStructName, ") bool { return ", g.hasExpr(fd, "msg"), " }")
	return nil
}

func (g goGen) genFdGet(fdStructName, msgStructName string, fd *zeropb.FieldDescriptor) error {
	g.P("func (fd *", fdStructName, ") get(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoReflectValueExpr(fd.Desc, g.getExpr(fd, "msg")), " }")
	return nil
}

func (g goGen) genFdSet(fdStructName, msgStructName string, field *zeropb.FieldDescriptor) error {
	g.P("func (fd *", fdStructName, ") set(msg *", msgStructName, ", v ", protoreflectImport.Ident("Value"), ") {}")
	return nil
}

func (g goGen) genFdClear(fdStructName, msgStructName string, field *zeropb.FieldDescriptor) error {
	g.P("func (fd *", fdStructName, ") clear(msg *", msgStructName, ") {}")
	return nil
}

func (g goGen) genFdMutable(fdStructName, msgStructName string, field *zeropb.FieldDescriptor) error {
	return nil
}

func (g goGen) genFdNewField(fdStructName, msgStructName string, field *zeropb.FieldDescriptor) error {
	g.P("func (fd *", fdStructName, ") newField(msg *", msgStructName, ") ", protoreflectImport.Ident("Value"), " { return ", protoreflectImport.Ident("Value"), "{} }")
	return nil
}
