package test3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestInternalNotRegistered(t *testing.T) {
	_, err := protoregistry.GlobalTypes.FindMessageByName((&TestAllTypes{}).ProtoReflect().Descriptor().FullName())
	require.Error(t, err)

	_, err = protoregistry.GlobalFiles.FindFileByPath("internal/testprotos/test3/test_import.proto")
	require.Error(t, err)
}
