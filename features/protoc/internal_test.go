package protoc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInternalPackage(t *testing.T) {
	require.True(t, isInternalPackage("internal"))
	require.True(t, isInternalPackage("example/internal"))
	require.True(t, isInternalPackage("example.com/internal"))
	require.True(t, isInternalPackage("example.com/internal/foo"))
	require.False(t, isInternalPackage("example"))
	require.False(t, isInternalPackage("example.com"))
	require.False(t, isInternalPackage("example.com/false"))
}
