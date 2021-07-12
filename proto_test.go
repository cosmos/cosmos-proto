package proto_test

import (
	p "github.com/cosmos/cosmos-proto/proto"
	"github.com/stretchr/testify/require"
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
	b := p.Bar{Baz: "Boop"}
	bz, err := b.Marshal()
	require.NoError(t, err)
	require.Greater(t, 0, len(bz))
}