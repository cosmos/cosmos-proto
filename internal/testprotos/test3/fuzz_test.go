package test3

import (
	"github.com/cosmos/cosmos-proto/internal/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"pgregory.net/rapid"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("marshal unmarshal", rapid.MakeCheck(func(t *rapid.T) {
		mType := (&TestAllTypes{}).ProtoReflect().Type()
		msg := fuzz.Message(t, mType, true, false)

		msgBytes, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg.Interface())
		require.NoError(t, err)

		uMsg := mType.New()
		err = proto.UnmarshalOptions{}.Unmarshal(msgBytes, uMsg.Interface())
		require.NoError(t, err)
		cmpOpt := protocmp.Transform()
		diff := cmp.Diff(uMsg.Interface(), msg.Interface(), cmpOpt)
		require.Emptyf(t, diff, "non matching messages\n%s", diff)
	}))
}

func TestOneofMarshalUnmarshal(t *testing.T) {
	msg1 := &TestAllTypes{OneofField: &TestAllTypes_OneofEnum{OneofEnum: NestedEnum_FOO}}
	b, err := proto.Marshal(msg1)
	require.NoError(t, err)

	msg2 := &TestAllTypes{}
	require.NoError(t, proto.Unmarshal(b, msg2))
	log.Printf("%s %s", msg2, msg1)

	dynMsg := dynamicpb.NewMessage(md_TestAllTypes)
	dynEnum := dynamicpb.NewEnumType(NestedEnum_FOO.Descriptor())

	dynMsg.Set(fd_TestAllTypes_oneof_enum, protoreflect.ValueOfEnum(dynEnum.New(NestedEnum_FOO.Number()).Number()))

	dynB, err := proto.Marshal(dynMsg)
	require.NoError(t, err)
	t.Logf("%s", dynMsg)

	dynMsg2 := dynamicpb.NewMessage(md_TestAllTypes)
	require.NoError(t, proto.Unmarshal(dynB, dynMsg2))

	t.Logf("%v", proto.Equal(dynMsg, dynMsg2))
}
