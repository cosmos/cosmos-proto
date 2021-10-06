package testpb

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"testing"
)

func TestSize(t *testing.T) {
	msg := B{X: "hello world!"}
	m := msg.ProtoReflect()

	methods := m.ProtoMethods()

	res := methods.Size(protoiface.SizeInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           m,
		Flags:             0,
	})

	expected := proto.Size(m.Interface())
	actual := res.Size

	require.Equal(t, expected, actual)
}

func TestMarshaler(t *testing.T) {
	msg := B{X: "hello cosmos!"}
	m := msg.ProtoReflect()

	bz := make([]byte, 0)
	methods := m.ProtoMethods()
	res, err := methods.Marshal(protoiface.MarshalInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           m,
		Buf:               bz,
		Flags:             0,
	})
	require.NoError(t, err)

	expected, err := proto.Marshal(m.Interface())
	require.NoError(t, err)

	require.Equal(t, res.Buf, expected)
}

func TestUnmarshaler(t *testing.T) {
	msg := B{X: "hello cosmos!"}
	m := msg.ProtoReflect()

	bz, err := proto.Marshal(m.Interface())
	require.NoError(t, err)

	methods := m.ProtoMethods()

	expected := B{}
	_, err = methods.Unmarshal(protoiface.UnmarshalInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           expected.ProtoReflect(),
		Buf:               bz,
		Flags:             0,
		Resolver:          nil,
	})
	require.NoError(t, err)

	require.Equal(t, msg.X, expected.X)
}
