package protoc

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type GoType interface {
	Display() interface{}
}

type StringType struct {
	s string
}

func (s StringType) Display() interface{} {
	return s.s
}

type IdentType struct {
	s protogen.GoIdent
}

func (it IdentType) Display() interface{} {
	return it.s
}

func NewGoType(i interface{}) GoType {
	switch i.(type) {
	case string:
		return StringType{s: i.(string)}
	case protogen.GoIdent:
		return IdentType{s: i.(protogen.GoIdent)}
	case GoType:
		return i.(GoType)
	default:
		return StringType{s: ""}
		//panic(fmt.Sprintf("unrecognized type: %+v", i))
	}
}
