package stablejson

import (
	"fmt"
	io "io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func marshalDuration(writer io.Writer, message protoreflect.Message) error {
	// PROTO3 SPEC:
	// Generated output always contains 0, 3, 6, or 9 fractional digits, depending on required precision, followed by
	// the suffix "s". Accepted are any fractional digits (also none) as long as they fit into nano-seconds precision
	// and the suffix "s" is required.

	fields := message.Descriptor().Fields()
	secondsField := fields.ByName(secondsName)
	if secondsField == nil {
		return fmt.Errorf("expected seconds field")
	}

	seconds := message.Get(secondsField).Int()
	_, err := fmt.Fprintf(writer, "%d", seconds)
	if err != nil {
		return err
	}

	nanosField := fields.ByName(nanosName)
	if nanosField == nil {
		return fmt.Errorf("expected nanos field")
	}

	nanos := message.Get(nanosField).Int()
	if nanos != 0 {
		if nanos > 0 {
			if seconds < 0 {
				return fmt.Errorf("seconds and nanos must be of the same sign for duration %v", message)
			}

			_, err := fmt.Fprintf(writer, ".%d", nanos)
			if err != nil {
				return err
			}
		} else {
			if seconds > 0 {
				return fmt.Errorf("seconds and nanos must be of the same sign for duration %v", message)
			}

			_, err := fmt.Fprintf(writer, ".%d", -nanos)
			if err != nil {
				return err
			}
		}
	}

	_, err = writer.Write([]byte("s"))
	return err
}
