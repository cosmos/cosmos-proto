package testpb

import (
	"google.golang.org/protobuf/testing/prototest"
	"testing"
)

func TestCompliance(t *testing.T) {
	prototest.Message{}.Test(t, (&A{}).ProtoReflect().Type())
}
