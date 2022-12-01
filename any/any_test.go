package any_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/cosmos/cosmos-proto/any"
	"github.com/cosmos/cosmos-proto/testpb"
)

func TestAny(t *testing.T) {
	value := &testpb.A{SomeBoolean: true}

	dst1 := &anypb.Any{}
	err := any.MarshalFrom(dst1, value, proto.MarshalOptions{})
	require.NoError(t, err)
	require.Equal(t, "/A", dst1.TypeUrl) // Make sure there's no "type.googleapis.com/" prefix.

	dst2, err := any.New(value)
	require.NoError(t, err)
	require.Equal(t, "/A", dst2.TypeUrl) // Make sure there's no "type.googleapis.com/" prefix.

	// Round trip.
	newValue, err := anypb.UnmarshalNew(dst2, proto.UnmarshalOptions{})
	require.NoError(t, err)
	diff := cmp.Diff(value, newValue, protocmp.Transform())
	require.Empty(t, diff)
}
