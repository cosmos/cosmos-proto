package rapidproto

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test1(t *testing.T) {
	gen := MessageGenerator(&timestamppb.Timestamp{})
	for i := 0; i < 100; i++ {
		x := gen.Example(i)
		t.Logf("%+v", x)
	}
}
