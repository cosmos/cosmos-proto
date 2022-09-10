package stablejson

import (
	"encoding/json"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (opts MarshalOptions) marshalMap(fieldDescriptor protoreflect.FieldDescriptor, value protoreflect.Map, writer io.Writer) error {
	_, err := writer.Write([]byte("{"))
	if err != nil {
		return err
	}

	allKeys := make([]protoreflect.MapKey, 0, value.Len())
	value.Range(func(key protoreflect.MapKey, _ protoreflect.Value) bool {
		allKeys = append(allKeys, key)
		return true
	})

	// TODO sort

	valueField := fieldDescriptor.MapValue()
	first := true
	for _, key := range allKeys {
		if !first {
			_, err = writer.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		first = false

		err = opts.marshalMapKey(key, writer)
		if err != nil {
			return err
		}

		_, err = writer.Write([]byte(":"))
		if err != nil {
			return err
		}

		err = opts.marshalSingleValue(valueField, value.Get(key), writer)
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte("}"))
	return err
}

func (opts MarshalOptions) marshalMapKey(key protoreflect.MapKey, writer io.Writer) error {
	switch key := key.Interface().(type) {
	case int32, int64, uint32, uint64:
		_, err := fmt.Fprintf(writer, `"%d"`, key)
		return err
	case string:
		return json.NewEncoder(writer).Encode(key)
	case bool:
		_, err := fmt.Fprintf(writer, `"%t"`, key)
		return err
	default:
		return fmt.Errorf("unexpected map key type %T", key)
	}
}
