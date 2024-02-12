package test3

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/testing/prototest"
)

func TestCompliance(t *testing.T) {
	prototest.Message{}.Test(t, (&TestAllTypes{}).ProtoReflect().Type())
}

func TestSourceCodeInfo(t *testing.T) {
	descriptor := (&ForeignMessage{}).ProtoReflect().Descriptor()
	sourceInfos := descriptor.ParentFile().SourceLocations()

	if !strings.Contains(sourceInfos.ByDescriptor(descriptor).LeadingComments, "This comment is for testing source code info comments") {
		t.Errorf("LeadingComments not found in source info")
	}
}
