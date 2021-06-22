package proto_test

import (
	p "github.com/cosmos/cosmos-proto/proto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	hey := p.Hello{World: "hey!!"}
	bz, err := hey.Marshal()
	require.NoError(t, err)
	require.Greater(t, len(bz), 0)

	var x p.Hello
	x.Unmarshal(bz)
	require.NotNil(t, x)
	require.Equal(t, x.World, hey.World)
}
