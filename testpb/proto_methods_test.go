package testpb

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/dynamicpb"
	"math"
	"pgregory.net/rapid"
	"testing"
)

func TestNegativeZero(t *testing.T) {

	testCases := []struct {
		name  string
		value float64
	}{
		{
			name:  "negative 0",
			value: math.Copysign(0, -1),
		},
		{
			name:  "negative float",
			value: -0.420,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := A{}
			a.DOUBLE = tc.value

			dyn := dynamicpb.NewMessage(md_A)
			dyn.Set(fd_A_DOUBLE, a.ProtoReflect().Get(fd_A_DOUBLE))
			bz, err := proto.Marshal(dyn)
			require.NoError(t, err)

			bz2, err := proto.Marshal(a.ProtoReflect().Interface())
			require.NoError(t, err)

			bz3, err := a.ProtoReflect().ProtoMethods().Marshal(protoiface.MarshalInput{Message: a.ProtoReflect()})
			require.NoError(t, err)

			require.Equal(t, bz, bz2)
			require.Equal(t, bz, bz3.Buf)
		})
	}
}

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
		require.True(t, proto.Equal(fastAA.Interface(), underlying.Interface()))
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
