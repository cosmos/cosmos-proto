package testpb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasBytes(t *testing.T) {
	fd := (&A{}).ProtoReflect().Descriptor().Fields().ByName("BYTES")
	t.Run("nil bytes", func(t *testing.T) {
		m := &A{BYTES: nil}

		require.False(t, m.slowProtoReflect().Has(fd))
		require.False(t, m.ProtoReflect().Has(fd))
	})

	t.Run("0 len bytes", func(t *testing.T) {
		m := &A{BYTES: []byte{}}

		require.False(t, m.slowProtoReflect().Has(fd))
		require.False(t, m.ProtoReflect().Has(fd))
	})
}
