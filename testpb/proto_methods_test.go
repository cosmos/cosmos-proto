package testpb

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/runtime/protoimpl"
	"pgregory.net/rapid"
	"testing"
)

func TestProtoMethods(t *testing.T) {
	t.Run("testSize", rapid.MakeCheck(testSize))
	t.Run("testMarshal", rapid.MakeCheck(testMarshal))
	t.Run("testUnmarshal", rapid.MakeCheck(testUnmarshal))
}

func testSize(t *rapid.T) {
	slowMsg := getRapidMsg(t)
	fastMsg := slowMsg.ProtoReflect()

	methods := fastMsg.ProtoMethods()

	result := methods.Size(protoiface.SizeInput{Message: fastMsg})
	expected := proto.Size(fastMsg.Interface())

	require.Equal(t, expected, result.Size)
}

func testMarshal(t *rapid.T) {
	slowMsg := getRapidMsg(t)
	fastMsg := slowMsg.ProtoReflect()

	methods := fastMsg.ProtoMethods()

	result, err := methods.Marshal(protoiface.MarshalInput{
		NoUnkeyedLiterals: struct{}{},
		Message:           fastMsg,
		Buf:               nil,
		Flags:             0,
	})
	require.NoError(t, err)

	canonical, err := proto.MarshalOptions{Deterministic: true}.Marshal(fastMsg.Interface())
	require.NoError(t, err)

	require.Equal(t, canonical, result.Buf)
}

func testUnmarshal(t *rapid.T) {
	slowMsg := getRapidMsg(t)
	fastMsg := slowMsg.ProtoReflect()

	proto.Marshal(fastMsg.Interface())

	methods := fastMsg.ProtoMethods()

	testCases := []struct {
		name      string
		unmarshal func(bz []byte, msg *A)
		marshal   func(msg proto.Message) []byte
	}{
		{
			name: "proto.Marshal pulsar Unmarshal",
			marshal: func(msg proto.Message) []byte {
				bz, err := proto.Marshal(msg)
				require.NoError(t, err)
				return bz
			},
			unmarshal: func(bz []byte, msg *A) {
				_, err := methods.Unmarshal(protoiface.UnmarshalInput{
					NoUnkeyedLiterals: struct{}{},
					Message:           msg.ProtoReflect(),
					Buf:               bz,
					Flags:             0,
					Resolver:          nil,
				})
				require.NoError(t, err)
			},
		},
		{
			name: "pulsar.Marshal, proto.Unmarshal",
			marshal: func(msg proto.Message) []byte {
				res, err := methods.Marshal(protoiface.MarshalInput{
					NoUnkeyedLiterals: struct{}{},
					Message:           msg.ProtoReflect(),
					Buf:               nil,
					Flags:             0,
				})
				require.NoError(t, err)
				return res.Buf
			},
			unmarshal: func(bz []byte, msg *A) {
				err := proto.Unmarshal(bz, msg)
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Logf("Test: %s\n", tc.name)
		aa := A{}
		bz := tc.marshal(fastMsg.Interface())
		tc.unmarshal(bz, &aa)
		fastAA := fastReflection_A(aa)
		underlying := fastMsg.(*fastReflection_A)
		CheckEqual(t, fastAA, *underlying)
	}
}

func CheckEqual(t *rapid.T, a1, a2 fastReflection_A) {
	require.Equal(t, a1.Enum, a2.Enum)
	require.Equal(t, a1.SomeBoolean, a2.SomeBoolean)
	require.Equal(t, a1.INT32, a2.INT32)
	require.Equal(t, a1.SINT32, a2.SINT32)
	require.Equal(t, a1.UINT32, a2.UINT32)
	require.Equal(t, a1.INT64, a2.INT64)
	require.Equal(t, a1.SING64, a2.SING64)
	require.Equal(t, a1.UINT64, a2.UINT64)
	require.Equal(t, a1.SFIXED32, a2.SFIXED32)
	require.Equal(t, a1.FIXED32, a2.FIXED32)
	require.Equal(t, a1.FLOAT, a2.FLOAT)
	require.Equal(t, a1.SFIXED64, a2.SFIXED64)
	require.Equal(t, a1.FIXED64, a2.FIXED64)
	require.Equal(t, a1.DOUBLE, a2.DOUBLE)
	require.Equal(t, a1.STRING, a2.STRING)
	if len(a1.BYTES) == 0 {
		require.True(t, len(a2.BYTES) == 0)
	} else {
		require.Equal(t, a1.BYTES, a2.BYTES)
	}
	require.Equal(t, a1.MESSAGE.X, a2.MESSAGE.X)
	switch a1o := a1.ONEOF.(type) {
	case *A_ONEOF_B:
		a2o := a2.ONEOF.(*A_ONEOF_B)
		require.Equal(t, a1o.ONEOF_B.X, a2o.ONEOF_B.X)
	case *A_ONEOF_STRING:
		a2o := a2.ONEOF.(*A_ONEOF_STRING)
		require.Equal(t, a1o.ONEOF_STRING, a2o.ONEOF_STRING)
	}
	if len(a1.LIST_ENUM) == 0 {
		require.True(t, len(a2.LIST_ENUM) == 0)
	} else {
		require.Equal(t, a1.LIST_ENUM, a2.LIST_ENUM)
	}
	for i := range a1.LIST {
		a1list := a1.LIST
		a2list := a2.LIST
		require.Equal(t, a1list[i].X, a2list[i].X)
	}
}

func getRapidMsg(t *rapid.T) A {
	return A{
		Enum:        Enumeration(rapid.IntRange(0, 1).Draw(t, "enum").(int)),
		SomeBoolean: rapid.Bool().Draw(t, "SomeBool").(bool),
		INT32:       rapid.Int32().Draw(t, "INT32").(int32),
		SINT32:      rapid.Int32().Draw(t, "SINT32").(int32),
		UINT32:      rapid.Uint32().Draw(t, "UINT32").(uint32),
		INT64:       rapid.Int64().Draw(t, "INT64").(int64),
		SING64:      rapid.Int64().Draw(t, "SING64").(int64),
		UINT64:      rapid.Uint64().Draw(t, "UINT64").(uint64),
		SFIXED32:    rapid.Int32().Draw(t, "SFIXED32").(int32),
		FIXED32:     rapid.Uint32().Draw(t, "FIXED32").(uint32),
		FLOAT:       rapid.Float32().Draw(t, "FLOAT").(float32),
		SFIXED64:    rapid.Int64().Draw(t, "SFIXED64").(int64),
		FIXED64:     rapid.Uint64().Draw(t, "FIXED64").(uint64),
		DOUBLE:      rapid.Float64().Draw(t, "DOUBLE").(float64),
		STRING:      rapid.String().Draw(t, "STRING").(string),
		BYTES:       rapid.SliceOf(rapid.Byte()).Draw(t, "byte slice").([]byte),
		MESSAGE:     genMessageB.Draw(t, "MESSAGE").(*B),
		LIST:        rapid.SliceOf(genMessageB).Draw(t, "LIST").([]*B),
		ONEOF:       genOneOf.Draw(t, "one of").(isA_ONEOF),
		LIST_ENUM:   rapid.SliceOf(genEnumSlice).Draw(t, "slice enum").([]Enumeration),
	}
}

var genEnumSlice = rapid.Custom(func(t *rapid.T) Enumeration {
	n := rapid.Int32Range(0, 1).Draw(t, "int32").(int32)
	return Enumeration(n)
})

var genOneOf = rapid.Custom(func(t *rapid.T) isA_ONEOF {
	oneof := rapid.OneOf(genOneOfB, genOneOfString).Draw(t, "oneof").(isA_ONEOF)
	return oneof
})

var genOneOfB = rapid.Custom(func(t *rapid.T) *A_ONEOF_B {
	return &A_ONEOF_B{ONEOF_B: genMessageB.Draw(t, "message B in one of").(*B)}
})

var genOneOfString = rapid.Custom(func(t *rapid.T) *A_ONEOF_STRING {
	return &A_ONEOF_STRING{ONEOF_STRING: rapid.StringN(1, -1, -1).Draw(t, "string in one of").(string)}
})

var genMessageB = rapid.Custom(func(t *rapid.T) *B {
	msg := B{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		X:             rapid.String().Draw(t, "X").(string),
	}
	return &msg
})
