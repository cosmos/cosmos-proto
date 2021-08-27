package testpb

import "google.golang.org/protobuf/reflect/protoreflect"

var _ protoreflect.List = (*ProtoListWrapper)(nil)

type ProtoListWrapper struct {
	list *[]interface{}
}

func (x ProtoListWrapper) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *ProtoListWrapper) Get(i int) protoreflect.Value {
	switch v := (*x.list)[i].(type) {
	case protoreflect.ProtoMessage:
		return protoreflect.ValueOf(v.ProtoReflect())
	}

	return protoreflect.ValueOf((*x.list)[i])
}

func (x *ProtoListWrapper) Set(i int, value protoreflect.Value) {
	(*x.list)[i] = value.Interface()
}

func (x *ProtoListWrapper) Append(value protoreflect.Value) {
	*x.list = append(*x.list, value)
}

func (x *ProtoListWrapper) AppendMutable() protoreflect.Value {
	v := new(protoreflect.Value)
	*x.list = append(*x.list, *v)
	return *v
}

func (x *ProtoListWrapper) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *ProtoListWrapper) NewElement() protoreflect.Value {
	v := new(protoreflect.Value)
	return *v
}

func (x *ProtoListWrapper) IsValid() bool {
	return x != nil
}
