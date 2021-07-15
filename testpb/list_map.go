package testpb

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TODO the map value implementer will look like
type _A_1_map_string_B struct {
}

func (_ _A_1_map_string_B) Len() int {
	panic("implement me")
}

func (_ _A_1_map_string_B) Range(f func(protoreflect.MapKey, protoreflect.Value) bool) {
	panic("implement me")
}

func (_ _A_1_map_string_B) Has(key protoreflect.MapKey) bool {
	panic("implement me")
}

func (_ _A_1_map_string_B) Clear(key protoreflect.MapKey) {
	panic("implement me")
}

func (_ _A_1_map_string_B) Get(key protoreflect.MapKey) protoreflect.Value {
	panic("implement me")
}

func (_ _A_1_map_string_B) Set(key protoreflect.MapKey, value protoreflect.Value) {
	panic("implement me")
}

func (_ _A_1_map_string_B) Mutable(key protoreflect.MapKey) protoreflect.Value {
	panic("implement me")
}

func (_ _A_1_map_string_B) NewValue() protoreflect.Value {
	panic("implement me")
}

func (_ _A_1_map_string_B) IsValid() bool {
	panic("implement me")
}
