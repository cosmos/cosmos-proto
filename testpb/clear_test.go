package testpb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestClear(t *testing.T) {
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")

	t.Run("field is cleared", func(t *testing.T) {
		m := &A{BYTES: []byte("test")}

		m.ProtoReflect().Clear(fd)

		require.Nil(t, m.BYTES)
	})

	t.Run("unknown field descriptor", func(t *testing.T) {
		m := &A{}

		require.Panics(t, func() {
			m.ProtoReflect().Clear((&anypb.Any{}).ProtoReflect().Descriptor().Fields().Get(0))
		})
	})
}
