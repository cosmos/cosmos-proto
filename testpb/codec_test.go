package testpb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

var testA = &A{
	Enum:        Enumeration_Two,
	SomeBoolean: true,
	INT32:       2,
	SINT32:      3,
	UINT32:      4,
	INT64:       5,
	SING64:      6,
	UINT64:      7,
	SFIXED32:    8,
	FIXED32:     9,
	FLOAT:       10.1,
	SFIXED64:    11,
	FIXED64:     12,
	DOUBLE:      13,
	STRING:      "fourteen",
	BYTES:       []byte("fifteen"),
	MESSAGE:     &B{X: "something"},
	MAP:         map[string]*B{"a": &B{X: "aa"}},
	LIST:        []*B{{X: "list"}},
	ONEOF:       &A_ONEOF_B{ONEOF_B: &B{X: "ONEOF"}},
	LIST_ENUM:   []Enumeration{Enumeration_One},
}

func TestMarshal(t *testing.T) {
	dynA := dynamicpb.NewMessage(md_A)
	// dynB := dynamicpb.NewMessage(md_B)

	dynA.Set(fd_A_enum, protoreflect.ValueOfEnum(protoreflect.EnumNumber(testA.Enum)))
	dynA.Set(fd_A_some_boolean, protoreflect.ValueOfBool(testA.SomeBoolean))
	dynA.Set(fd_A_INT32, protoreflect.ValueOfInt32(testA.INT32))
	dynA.Set(fd_A_SINT32, protoreflect.ValueOfInt32(testA.SINT32))
	dynA.Set(fd_A_UINT32, protoreflect.ValueOfUint32(testA.UINT32))
	dynA.Set(fd_A_INT64, protoreflect.ValueOfInt64(testA.INT64))
	dynA.Set(fd_A_SING64, protoreflect.ValueOfInt64(testA.SING64))
	dynA.Set(fd_A_UINT64, protoreflect.ValueOfUint64(testA.UINT64))
	dynA.Set(fd_A_SFIXED32, protoreflect.ValueOfInt32(testA.SFIXED32))
	dynA.Set(fd_A_FIXED32, protoreflect.ValueOfUint32(testA.FIXED32))
	dynA.Set(fd_A_FLOAT, protoreflect.ValueOfFloat32(testA.FLOAT))
	dynA.Set(fd_A_SFIXED64, protoreflect.ValueOfInt64(testA.SFIXED64))
	dynA.Set(fd_A_FIXED64, protoreflect.ValueOfUint64(testA.FIXED64))
	dynA.Set(fd_A_DOUBLE, protoreflect.ValueOfFloat64(testA.DOUBLE))
	dynA.Set(fd_A_STRING, protoreflect.ValueOfString(testA.STRING))
	dynA.Set(fd_A_BYTES, protoreflect.ValueOfBytes(testA.BYTES))
	dynA.Set(fd_A_MESSAGE, protoreflect.ValueOfMessage(testA.MESSAGE.ProtoReflect()))
	dynMap := dynA.Mutable(fd_A_MAP).Map()
	dynMap.Set(protoreflect.MapKey(protoreflect.ValueOfString("a")), protoreflect.ValueOfMessage(testA.MAP["a"].ProtoReflect()))
	dynA.Mutable(fd_A_LIST).List().AppendMutable().Message().Set(fd_B_x, protoreflect.ValueOfString(testA.LIST[0].X))
	dynA.Set(fd_A_ONEOF_B, protoreflect.ValueOfMessage(testA.ONEOF.(*A_ONEOF_B).ONEOF_B.ProtoReflect()))
	dynA.Mutable(fd_A_LIST_ENUM).List().Append(protoreflect.ValueOfEnum((protoreflect.EnumNumber)(Enumeration_One)))

	got, err := proto.MarshalOptions{Deterministic: true}.Marshal(testA)
	require.NoError(t, err)

	expected, err := proto.MarshalOptions{Deterministic: true}.Marshal(dynA)
	require.NoError(t, err)

	require.Equal(t, expected, got)

	zpbBuf := make([]byte, 64*1024)
	n, err := testA.MarshalZeroPB(zpbBuf)
	require.NoError(t, err)
	var msg2 A
	n2, err := msg2.UnmarshalZeroPB(zpbBuf[:n])
	require.NoError(t, err)
	require.Equal(t, n, n2)
	zpbGot, err := proto.MarshalOptions{Deterministic: true}.Marshal(&msg2)
	require.NoError(t, err)
	require.Equal(t, expected, zpbGot)
}

func BenchmarkMarshalA(b *testing.B) {
	protoOpts := proto.MarshalOptions{}
	b.Run("proto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			protoOpts.Marshal(testA)
		}
	})
	zpbBuf := make([]byte, 64*1024)
	b.Run("zeropb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testA.MarshalZeroPB(zpbBuf)
		}
	})
	n, err := testA.MarshalZeroPB(zpbBuf)
	require.NoError(b, err)
	data := make([]byte, n)
	testA.MarshalZeroPB(data)
	b.Run("copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(zpbBuf, data)
			testA.MarshalZeroPB(zpbBuf)
		}
	})
}

func BenchmarkUnmarshalA(b *testing.B) {
	protoA, err := proto.MarshalOptions{}.Marshal(testA)
	require.NoError(b, err)
	b.Run("proto", func(b *testing.B) {
		var a A
		for i := 0; i < b.N; i++ {
			proto.Unmarshal(protoA, &a)
		}
	})
	zpbBuf := make([]byte, 64*1024)
	n, err := testA.MarshalZeroPB(zpbBuf)
	require.NoError(b, err)
	b.Run("zeropb", func(b *testing.B) {
		var testA A
		for i := 0; i < b.N; i++ {
			testA.UnmarshalZeroPB(zpbBuf[:n])
		}
	})
}
