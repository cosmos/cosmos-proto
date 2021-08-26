package testpb

import "google.golang.org/protobuf/reflect/protoreflect"

type ListWrapper struct {
	protoreflect.List
}

var _ protoreflect.List = (*GenericList)(nil)

type GenericList struct {
	list *[]interface{}
}

func (x GenericList) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *GenericList) Get(i int) protoreflect.Value {
	return protoreflect.ValueOf((*x.list)[i])
}

func (x *GenericList) Set(i int, value protoreflect.Value) {
	(*x.list)[i] = value.Interface()
}

func (x *GenericList) Append(value protoreflect.Value) {
	*x.list = append(*x.list, value)
}

func (x *GenericList) AppendMutable() protoreflect.Value {
	v := new(protoreflect.Value)
	*x.list = append(*x.list, *v)
	return *v
}

func (x *GenericList) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *GenericList) NewElement() protoreflect.Value {
	v := new(protoreflect.Value)
	return *v
}

func (x *GenericList) IsValid() bool {
	return x != nil
}
