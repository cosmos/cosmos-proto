package proto_test

import (
	"fmt"
	examples "github.com/cosmos/cosmos-proto/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
	"testing"
)

//func TestUnmarshal(t *testing.T) {
//	p.
//	hey := p.Hello{World: "hey!!"}
//	bz, err := hey.Marshal()
//	require.NoError(t, err)
//	require.Greater(t, len(bz), 0)
//
//	var x p.Hello
//	x.Unmarshal(bz)
//	require.NotNil(t, x)
//	require.Equal(t, x.World, hey.World)
//
//	r := x.ProtoReflect()
//	fmt.Printf("\"%+v\n\"", r)
//}
//
//
//func TestProtoMethods(t *testing.T) {
//	tm := p.Tender{
//		Mint:    "consensus",
//		Version: 1,
//	}
//
//	msg := tm.ProtoReflect()
//	methods := msg.ProtoMethods()
//	require.NotNil(t, methods)
//}

func TestTest(t *testing.T) {
	//b := examples.Bar{Baz: "hello"}
	//methods := b.ProtoMethods()
	//marshal := methods.Marshal
	//bz, err := marshal(protoiface.MarshalInput{
	//	NoUnkeyedLiterals: struct{}{},
	//	Message:           nil,
	//	Buf:               nil,
	//	Flags:             0,
	//})
	//require.NoError(t, err)
	//require.NotEmpty(t, bz.Buf)
	//fmt.Println(bz.Buf)
	//var m examples.Bar
	//var bytz []byte
	//z, err := methods.Unmarshal(protoiface.UnmarshalInput{
	//	NoUnkeyedLiterals: struct{}{},
	//	Message: &m,
	//	Buf:               bytz,
	//	Flags:             0,
	//	Resolver:          nil,
	//})
	//require.NoError(t, err)
	//prettyPrintStruct(z)
}

func prettyPrintStruct(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func TestUnmarshalProtoMethod(t *testing.T) {
	foo := examples.Bar{Baz: "test"}
	bz, err := foo.Marshal()
	require.NoError(t, err)
	require.NotEmpty(t, bz)

	methods := foo.ProtoMethods()

	var qux examples.Bar
	out, err := methods.Unmarshal(protoiface.UnmarshalInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           &qux,
		Buf:               bz,
		Flags:             0,
		Resolver:          nil,
	})
	require.NoError(t, err)
	prettyPrintStruct(out)
	prettyPrintStruct(qux)
	prettyPrintStruct(foo)
}

func TestReflectMethodGet(t *testing.T) {
	foo := examples.Bar{Baz: "Qux"}
	bz, err := foo.Marshal()
	require.NoError(t, err)
	var m protoreflect.Message = foo
	methods := m.ProtoMethods()
	out, err := methods.Marshal(protoiface.MarshalInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           foo,
		Buf:               nil,
		Flags:             0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, out.Buf)
	require.Equal(t, bz, out.Buf)
}

func TestProtoReflectMethod(t *testing.T) {
	foo := examples.Bar{Baz: "hi"}
	msg := foo.ProtoReflect()
	require.NotNil(t, msg.ProtoMethods())
}

func TestSizeMethod(t *testing.T) {
	foo := examples.Bar{Baz: "hi"}
	size := func(msg protoreflect.Message) int {
		methods := msg.ProtoMethods()
		return methods.Size(protoiface.SizeInput{}).Size
	}
	require.Equal(t, 4, size(foo))
}
