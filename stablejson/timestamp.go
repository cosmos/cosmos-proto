package stablejson

import (
	"fmt"
	io "io"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	secondsName protoreflect.Name = "seconds"
	nanosName   protoreflect.Name = "nanos"
)

func marshalTimestamp(writer io.Writer, message protoreflect.Message) error {
	// PROTO3 SPEC:
	// Uses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits.
	// Offsets other than "Z" are also accepted.

	fields := message.Descriptor().Fields()
	secondsField := fields.ByName(secondsName)
	if secondsField == nil {
		return fmt.Errorf("expected seconds field")
	}

	nanosField := fields.ByName(nanosName)
	if nanosField == nil {
		return fmt.Errorf("expected nanos field")
	}

	seconds := message.Get(secondsField).Int()
	nanos := message.Get(nanosField).Int()
	if nanos < 0 {
		return fmt.Errorf("nanos must be non-negative on timestamp %v", message)
	}

	t := time.Unix(seconds, nanos).UTC()
	var str string
	if nanos == 0 {
		str = t.Format(time.RFC3339)
	} else {
		str = t.Format(time.RFC3339Nano)
	}

	_, err := writer.Write([]byte(str))
	return err
}
