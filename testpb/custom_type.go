package testpb

import (
	"github.com/cosmos/cosmos-proto/runtime"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

var _ runtime.CustomType = (*Int)(nil)

type Int struct {
}

func (i Int) UnmarshalBytes(in protoiface.UnmarshalInput, b []byte) (out protoiface.UnmarshalOutput, err error) {
	panic("implement me")
}

func (i Int) MarshalBytes(in protoiface.MarshalOutput) (out protoiface.MarshalOutput, err error) {
	panic("implement me")
}

func (i Int) Size(in protoiface.SizeInput) (out protoiface.SizeOutput) {
	panic("implement me")
}

func (i Int) Set(value protoreflect.Value) {
	panic("implement me")
}

func (i Int) Get() protoreflect.Value {
	panic("implement me")
}

func (i Int) Clear() {
	panic("implement me")
}

func (i Int) IsSet() bool {
	panic("implement me")
}

func (i Int) Mutable() protoreflect.Value {
	panic("implement me")
}
