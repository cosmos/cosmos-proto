package testpb

import (
	"testing"
)

func TestOneOf(t *testing.T) {
	x := &A{}
	oneOf1 := x.ProtoReflect().Descriptor().Fields().ByName("ONEOF_STRING")
	oneOf2 := x.ProtoReflect().Descriptor().Fields().ByName("ONEOF_B")
	t.Logf("%s", oneOf1)

	v := x.ProtoReflect().Get(oneOf1)
	t.Logf("%s", v.Interface()) // empty string

	v = x.ProtoReflect().Get(oneOf2)
	t.Logf("%v", v.Message().IsValid()) // nil object, is valid
}
