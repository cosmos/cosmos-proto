package cosmos_proto_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	cosmos_proto "github.com/cosmos/cosmos-proto"
	"github.com/cosmos/cosmos-proto/testpb"
	"github.com/stretchr/testify/require"
)

func TestAny(t *testing.T) {
	any := &cosmos_proto.Any{}
	value := &testpb.A{}
	err := cosmos_proto.MarshalFrom(any, value, proto.MarshalOptions{})
	require.NoError(t, err)

	// Make sure there's no "type.googleapis.com/" prefix.
	require.Equal(t, "A", any.TypeUrl)
}
