package testpb

import "google.golang.org/protobuf/reflect/protoreflect"

var _ protoreflect.List = (*GenericList)(nil)

type GenericList []protoreflect.Value

func (x *GenericList) Len() int {
	if x == nil {
		return 0
	}
	return len(*x)
}

func (x *GenericList) Get(i int) protoreflect.Value {
	return protoreflect.ValueOf((*x)[i])
}

func (x *GenericList) Set(i int, value protoreflect.Value) {
	(*x)[i] = value
}

func (x *GenericList) Append(value protoreflect.Value) {
	*x = append(*x, value)
}

func (x *GenericList) AppendMutable() protoreflect.Value {
	v := new(protoreflect.Value)
	*x = append(*x, *v)
	return *v
}

func (x *GenericList) Truncate(n int) {
	*x = (*x)[:n]
}

func (x *GenericList) NewElement() protoreflect.Value {
	v := new(protoreflect.Value)
	return *v
}

func (x *GenericList) IsValid() bool {
	return x != nil
}
