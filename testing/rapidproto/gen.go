package rapidproto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
	"pgregory.net/rapid"
)

func MessageGenerator[T proto.Message](x T, options GeneratorOptions) *rapid.Generator[T] {
	msgType := x.ProtoReflect().Type()
	return rapid.Custom(func(t *rapid.T) T {
		msg := msgType.New()

		options.setFields(t, msg, 0)

		return msg.Interface().(T)
	})
}

type GeneratorOptions struct {
}

const depthLimit = 10

func (opts GeneratorOptions) setFields(t *rapid.T, msg protoreflect.Message, depth int) {
	// to avoid stack overflow we limit the depth of nested messages
	if depth > depthLimit {
		return
	}

	descriptor := msg.Descriptor()
	fullName := descriptor.FullName()
	switch fullName {
	case timestampFullName:
		opts.genTimestamp(t, msg)
	case durationFullName:
		opts.genDuration(t, msg)
	default:
		fields := descriptor.Fields()
		n := fields.Len()
		for i := 0; i < n; i++ {
			field := fields.Get(i)
			if !rapid.Bool().Draw(t, fmt.Sprintf("gen-%s", field.Name())) {
				continue
			}

			opts.setFieldValue(t, msg, field, depth)
		}
	}
}

const (
	timestampFullName protoreflect.FullName = "google.protobuf.Timestamp"
	durationFullName                        = "google.protobuf.Duration"
)

func (opts GeneratorOptions) setFieldValue(t *rapid.T, msg protoreflect.Message, field protoreflect.FieldDescriptor, depth int) {
	name := string(field.Name())
	kind := field.Kind()

	if field.IsList() {

		list := msg.Mutable(field).List()
		n := rapid.IntRange(0, 10).Draw(t, fmt.Sprintf("%sN", name))
		for i := 0; i < n; i++ {
			if kind == protoreflect.MessageKind || kind == protoreflect.GroupKind {
				opts.setFields(t, list.AppendMutable().Message(), depth+1)
			} else {
				list.Append(opts.genScalarFieldValue(t, field, fmt.Sprintf("%s%d", name, i)))
			}
		}

	} else if field.IsMap() {

		m := msg.Mutable(field).Map()
		n := rapid.IntRange(0, 10).Draw(t, fmt.Sprintf("%sN", name))
		for i := 0; i < n; i++ {
			keyField := field.MapKey()
			valueField := field.MapValue()
			valueKind := valueField.Kind()
			key := opts.genScalarFieldValue(t, keyField, fmt.Sprintf("%s%d-key", name, i))
			if valueKind == protoreflect.MessageKind || valueKind == protoreflect.GroupKind {
				opts.setFields(t, m.Mutable(key.MapKey()).Message(), depth+1)
			} else {
				value := opts.genScalarFieldValue(t, valueField, fmt.Sprintf("%s%d-key", name, i))
				m.Set(key.MapKey(), value)
			}
		}

	} else {

		if kind == protoreflect.MessageKind || kind == protoreflect.GroupKind {
			opts.setFields(t, msg.Mutable(field).Message(), depth+1)
		} else {
			msg.Set(field, opts.genScalarFieldValue(t, field, name))
		}

	}
}

func (opts GeneratorOptions) genScalarFieldValue(t *rapid.T, field protoreflect.FieldDescriptor, name string) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(rapid.Int32().Draw(t, name))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(rapid.Uint32().Draw(t, name))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(rapid.Int64().Draw(t, name))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(rapid.Uint64().Draw(t, name))
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(rapid.Bool().Draw(t, name))
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes(rapid.SliceOf(rapid.Byte()).Draw(t, name))
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(rapid.Float32().Draw(t, name))
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(rapid.Float64().Draw(t, name))
	case protoreflect.EnumKind:
		enumValues := field.Enum().Values()
		val := rapid.Int32Range(0, int32(enumValues.Len()-1)).Draw(t, name)
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(val))
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(rapid.String().Draw(t, name))
	default:
		t.Fatalf("unexpected %v", field)
		return protoreflect.Value{}
	}
}

const (
	secondsName protoreflect.Name = "seconds"
	nanosName                     = "nanos"
)

func (opts GeneratorOptions) genTimestamp(t *rapid.T, msg protoreflect.Message) {
	seconds := rapid.Int64Range(-9999999999, 9999999999).Draw(t, "seconds")
	nanos := rapid.Int32Range(0, 999999999).Draw(t, "nanos")
	setSecondsNanosFields(t, msg, seconds, nanos)
}

func (opts GeneratorOptions) genDuration(t *rapid.T, msg protoreflect.Message) {
	seconds := rapid.Int64Range(0, 315576000000).Draw(t, "seconds")
	nanos := rapid.Int32Range(0, 999999999).Draw(t, "nanos")
	setSecondsNanosFields(t, msg, seconds, nanos)
}

func setSecondsNanosFields(t *rapid.T, message protoreflect.Message, seconds int64, nanos int32) {
	fields := message.Descriptor().Fields()

	secondsField := fields.ByName(secondsName)
	assert.Assert(t, secondsField != nil)
	message.Set(secondsField, protoreflect.ValueOfInt64(seconds))

	nanosField := fields.ByName(nanosName)
	assert.Assert(t, nanosField != nil)
	message.Set(nanosField, protoreflect.ValueOfInt32(nanos))

	return
}
