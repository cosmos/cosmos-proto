package testpb

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type slowPrWrapper struct {
	x interface {
		slowProtoReflect() protoreflect.Message
	}
}

func (x slowPrWrapper) ProtoReflect() protoreflect.Message {
	return x.x.slowProtoReflect()
}

func TestMarshal(t *testing.T) {
	msg := &A{
		Enum:        Enumeration_Two,
		SomeBoolean: true,
		INT32:       2,
		SINT32:      3,
		UINT32:      4,
		INT64:       5,
		SING64:      6,
		UINT64:      7,
		SFIXED32:    8,
		FIXED32:     9,
		FLOAT:       10.1,
		SFIXED64:    11,
		FIXED64:     12,
		DOUBLE:      13,
		STRING:      "fourteen",
		BYTES:       []byte("fifteen"),
		MESSAGE:     &B{X: "something"},
		MAP:         map[string]*B{"a": &B{X: "aa"}},
		LIST:        []*B{{X: "list"}},
		ONEOF:       &A_ONEOF_B{ONEOF_B: &B{X: "ONEOF"}},
		LIST_ENUM:   []Enumeration{Enumeration_One},
	}

	spw := slowPrWrapper{x: msg}

	got, err := proto.MarshalOptions{Deterministic: true}.Marshal(msg)
	require.NoError(t, err)

	expected, err := proto.MarshalOptions{Deterministic: true}.Marshal(spw)
	require.NoError(t, err)

	require.Equal(t, expected, got)
}
