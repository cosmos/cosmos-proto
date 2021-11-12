package fuzz

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"math"
	"pgregory.net/rapid"
)

const (
	MaxBytesArraySize = 100
	MaxListLength     = 50
	MaxDepthDefault   = 1
)

func MessageWithDepth(t *rapid.T, messageType protoreflect.MessageType, fillAll bool, invalidValues bool, maxDepth int) protoreflect.Message {
	return genMessage(t, messageType, fillAll, invalidValues, 0, maxDepth)
}

// Message generates a random protoreflect.Message given its protoreflect.MessageType
// if fillAll is true, then all the fields will be filled
// if invalidValues is true, then the genMessage will contain invalid values
func Message(t *rapid.T, messageType protoreflect.MessageType, fillAll bool, invalidValues bool) protoreflect.Message {
	return genMessage(t, messageType, fillAll, invalidValues, 0, MaxDepthDefault)
}

func genMessage(t *rapid.T, messageType protoreflect.MessageType, fillAll bool, invalidValues bool, depth, maxDepth int) protoreflect.Message {
	if depth == maxDepth {
		return protoreflect.ValueOfMessage(messageType.New()).Message()
	}
	m := messageType.New()
	md := messageType.Descriptor()

	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		// skip if we're not filling all fields and if rapid says not to generate the field
		if !fillAll && !rapid.Bool().Draw(t, fmt.Sprintf("generate random field genValue for %s", fd.FullName())).(bool) {
			continue
		}
		value := genValue(t, fd, messageType, fillAll, invalidValues, depth, maxDepth)
		m.Set(fd, value)
	}

	return m
}

// genValue generates a random protoreflect.Value given its protoreflect.FieldDescriptor
// TODO(fdymylja): how to properly handle oneof fields?
func genValue(t *rapid.T,
	fd protoreflect.FieldDescriptor,
	messageType protoreflect.MessageType,
	fillAll bool,
	invValues bool,
	depth,
	maxDepth int) protoreflect.Value {

	if invValues {
		invalid := rapid.Bool().Draw(t, fmt.Sprintf("generate invalid genValue for %s", fd.FullName())).(bool)
		if invalid {
			return protoreflect.Value{}
		}
	}

	// check if list or map
	switch {
	case fd.IsList():
		listValue := genList(t, fd, messageType, fillAll, invValues, depth, maxDepth)
		return protoreflect.ValueOfList(listValue)
	case fd.IsMap():
		mapValue := genMap(t, fd, messageType, fillAll, invValues, depth, maxDepth)
		return protoreflect.ValueOfMap(mapValue)
	default:
		return genSingularValue(t, fd, messageType, fillAll, invValues, depth, maxDepth)
	}
}

func genMap(t *rapid.T, fd protoreflect.FieldDescriptor, messageType protoreflect.MessageType, all bool, values bool, depth, maxDepth int) protoreflect.Map {
	keyDesc := fd.MapKey()
	valueDesc := fd.MapValue()

	mapValue := messageType.New().NewField(fd).Map()

	// random list length
	length := rapid.IntRange(0, MaxListLength).Draw(t, "map length for "+string(fd.FullName())).(int)

	for i := 0; i < length; i++ {
		keyValue := genSingularValue(t, keyDesc, nil, all, values, depth, maxDepth)
		var valueValue protoreflect.Value
		switch valueDesc.Kind() {
		case protoreflect.MessageKind:
			valueType := mapValue.NewValue().Message().Type()
			valueValue = protoreflect.ValueOfMessage(genMessage(t, valueType, all, values, depth+1, maxDepth))
		default:
			valueValue = genSingularValue(t, valueDesc, nil, all, values, depth, maxDepth)
		}

		mapValue.Set(protoreflect.MapKey(keyValue), valueValue)
	}

	return mapValue
}

func genList(t *rapid.T, fd protoreflect.FieldDescriptor, messageType protoreflect.MessageType, all bool, values bool, depth, maxDepth int) protoreflect.List {
	list := messageType.New().NewField(fd).List()
	// random list length
	length := rapid.IntRange(0, MaxListLength).Draw(t, "list length for "+string(fd.FullName())).(int)
	for i := 0; i < length; i++ {
		switch fd.Kind() {
		case protoreflect.MessageKind:
			mType := list.NewElement().Message().Type()
			listValue := genMessage(t, mType, all, values, depth+1, maxDepth)
			list.Append(protoreflect.ValueOfMessage(listValue))
		default:
			list.Append(genSingularValue(t, fd, messageType, all, values, depth, maxDepth))
		}
	}

	return list
}

// genSingularValue generates a random protoreflect.Value for the specified protoreflect.FieldDescriptor
func genSingularValue(t *rapid.T, fd protoreflect.FieldDescriptor, messageType protoreflect.MessageType, fillAll bool, invValues bool, depth, maxDepth int) protoreflect.Value {
	switch fd.Kind() {
	// bool kind
	case protoreflect.BoolKind:
		value := rapid.Bool().Draw(t, label(fd)).(bool)
		return protoreflect.ValueOfBool(value)
	// int32 kinds
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		value := rapid.Int32().Draw(t, label(fd)).(int32)
		return protoreflect.ValueOfInt32(value)
	// int64 kinds
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		value := rapid.Int64().Draw(t, label(fd)).(int64)
		return protoreflect.ValueOfInt64(value)
	// uint32 kinds
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		value := rapid.Uint32().Draw(t, label(fd)).(uint32)
		return protoreflect.ValueOfUint32(value)
	// uint64 kinds
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		value := rapid.Uint64().Draw(t, label(fd)).(uint64)
		return protoreflect.ValueOfUint64(value)
	// float32 kind
	case protoreflect.FloatKind:
		value := rapid.Float32Max(math.MaxFloat32).Draw(t, label(fd)).(float32)
		return protoreflect.ValueOfFloat32(value)
	// float64 kind
	case protoreflect.DoubleKind:
		value := rapid.Float64().Draw(t, label(fd)).(float64)
		return protoreflect.ValueOfFloat64(value)
	// string kind
	case protoreflect.StringKind:
		value := rapid.String().Draw(t, label(fd)).(string)
		return protoreflect.ValueOfString(value)
	// bytes kind
	case protoreflect.BytesKind:
		value := randomBytes(t, fd)
		return protoreflect.ValueOfBytes(value)
	// enum kind
	case protoreflect.EnumKind:
		enumIndex := rapid.IntRange(0, fd.Enum().Values().Len()-1).Draw(t, "random enum index for "+string(fd.FullName())).(int)
		enum := fd.Enum().Values().Get(enumIndex)
		return protoreflect.ValueOfEnum(enum.Number())
	// message kinds
	case protoreflect.MessageKind:
		embeddedMessageType := messageType.New().NewField(fd).Message().Type()
		messageValue := genMessage(t, embeddedMessageType, fillAll, invValues, depth+1, maxDepth)
		return protoreflect.ValueOfMessage(messageValue)
	default:
		panic(fmt.Errorf("cannot handle: %s", fd.Kind()))
	}
}

func randomBytes(t *rapid.T, fd protoreflect.FieldDescriptor) []byte {
	length := rapid.IntRange(0, MaxBytesArraySize).Draw(t, "maximum bytes length for "+string(fd.FullName())).(int)

	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rapid.ByteMax(math.MaxUint8).Draw(t, fmt.Sprintf("genValue of byte in %s at index %d", fd.FullName(), i)).(byte)
		bytes[i] = b
	}

	return bytes
}

func label(fd protoreflect.FieldDescriptor) string {
	return fmt.Sprintf("random %s for %s", fd.Kind(), fd.FullName())
}
