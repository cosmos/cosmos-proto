package zpbgen

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type zeroCopyMessageDescriptor struct {
	message        *protogen.Message
	fields         []*zeroCopyFieldDescriptor
	fieldsByName   map[protoreflect.Name]*zeroCopyFieldDescriptor
	fieldsByNumber map[protoreflect.FieldNumber]*zeroCopyFieldDescriptor
}

func newZeroCopyDescriptor(message *protogen.Message) (*zeroCopyMessageDescriptor, error) {
	md := &zeroCopyMessageDescriptor{
		message:        message,
		fieldsByName:   map[protoreflect.Name]*zeroCopyFieldDescriptor{},
		fieldsByNumber: map[protoreflect.FieldNumber]*zeroCopyFieldDescriptor{},
	}

	var lastFieldNum protoreflect.FieldNumber
	currentOffset := uint32(0)
	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Desc.Number() < fields[j].Desc.Number()
	})

	for _, field := range fields {
		num := field.Desc.Number()
		if num-lastFieldNum != 1 {
			return nil, fmt.Errorf("field numbers must be consecutive, but %d is followed by %d in message %s", lastFieldNum, num, message.Desc.FullName())
		}
		size, err := fieldSize(field)
		if err != nil {
			return nil, err
		}

		fd := &zeroCopyFieldDescriptor{
			offset: currentOffset,
			size:   size,
			Field:  field,
		}

		md.fieldsByName[field.Desc.Name()] = fd
		md.fieldsByNumber[num] = fd
		md.fields = append(md.fields, fd)

		currentOffset += size
		lastFieldNum = num
	}

	return md, nil
}

func (z *zeroCopyMessageDescriptor) fieldOffset(field *protogen.Field) (uint32, error) {
	fd, ok := z.fieldsByName[field.Desc.Name()]
	if !ok {
		return 0, fmt.Errorf("field %s not found", field.Desc.Name())
	}
	return fd.offset, nil
}

func fieldSize(field *protogen.Field) (uint32, error) {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return 1, nil
	case protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.FloatKind, protoreflect.Fixed32Kind:
		return 4, nil
	case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind, protoreflect.DoubleKind, protoreflect.Fixed64Kind:
		return 8, nil
	case protoreflect.StringKind:
		return 2, nil
	case protoreflect.BytesKind:
		return 2, nil
	default:
		//return 0, fmt.Errorf("unhandled field kind: %v", field.Desc.Kind())
		//fmt.Printf("unhandled field kind: %v\n", field.Desc.Kind())
		return 0, nil
	}
}

type zeroCopyFieldDescriptor struct {
	offset uint32
	size   uint32
	*protogen.Field
}
