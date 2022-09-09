package rapidproto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"pgregory.net/rapid"
)

func Message[T proto.Message](x T, options ...GeneratorOption) *rapid.Generator[T] {
	msgType := x.ProtoReflect().Type()
	return rapid.Custom(func(t *rapid.T) T {
		msg := msgType.New()
		fields := msg.Descriptor().Fields()
		n := fields.Len()
		for i := 0; i < n; i++ {
			field := fields.Get(i)
			if !rapid.Bool().Draw(t, fmt.Sprintf("gen-%s", field.Name())) {
				continue
			}

			setFieldValue(t, msg, field)
		}

		return msg.Interface().(T)
	})
}

type GeneratorOption interface {
	private()
}

func setFieldValue(t *rapid.T, msg protoreflect.Message, field protoreflect.FieldDescriptor) {
	switch field.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		msg.Set(field, protoreflect.ValueOfInt32(rapid.Int32().Draw(t, string(field.Name()))))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		msg.Set(field, protoreflect.ValueOfUint32(rapid.Uint32().Draw(t, string(field.Name()))))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		msg.Set(field, protoreflect.ValueOfInt64(rapid.Int64().Draw(t, string(field.Name()))))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		msg.Set(field, protoreflect.ValueOfUint64(rapid.Uint64().Draw(t, string(field.Name()))))
	default:
		panic("TODO")
	}
}
