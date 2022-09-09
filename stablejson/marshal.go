package stablejson

import (
	"bytes"
	"io"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func Marshal(message proto.Message) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := MarshalOptions{}.Marshal(message, buf)
	return buf.Bytes(), err
}

type MarshalOptions struct {
	// HexBytes specifies whether bytes fields should be marshaled as upper-case
	// hex strings. If set to false, bytes fields will be encoded as standard
	// base64 strings as specified by the official proto3 JSON mapping.
	HexBytes bool

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Marshal by default encodes the provided proto.Message as spec-compliant
// proto3 JSON with the following restrictions which ensure a deterministic encoding:
// - fields are ordered based on field number
// - map fields are ordered
//
//   - alphabetically for string keys
//
//   - in numeric order for numeric keys
//
//   - false first for boolean keys
//
//   - durations have either 0 or 9 fractional digits depending on whether any fractional digits are needed
//
//   - timestamps have either 0 or 9 fractional digits depending on whether any fractional digits are needed
//
//   - floats and doubles always have the minimum number of trailing digits possible, although these types should
//     generally be avoided in deterministic applications
func (opts MarshalOptions) Marshal(message proto.Message, writer io.Writer) error {
	return opts.marshalMessage(message.ProtoReflect(), writer)
}
