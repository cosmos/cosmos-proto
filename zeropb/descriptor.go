package zeropb

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Registry struct {
	messages map[protoreflect.FullName]*MessageDescriptor
}

func NewRegistry() *Registry {
	return &Registry{
		messages: map[protoreflect.FullName]*MessageDescriptor{},
	}
}

func (r *Registry) MessageDescriptor(descriptor protoreflect.MessageDescriptor) (*MessageDescriptor, error) {
	md := &MessageDescriptor{
		MessageDescriptor: descriptor,
		fieldsByName:      map[protoreflect.Name]*FieldDescriptor{},
		fieldsByNumber:    map[protoreflect.FieldNumber]*FieldDescriptor{},
	}

	var lastFieldNum protoreflect.FieldNumber
	currentOffset := uint32(0)
	n := descriptor.Fields().Len()
	fields := make([]protoreflect.FieldDescriptor, n)
	for i := 0; i < n; i++ {
		fields[i] = descriptor.Fields().Get(i)
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Number() < fields[j].Number()
	})

	for _, field := range fields {
		num := field.Number()
		if num-lastFieldNum != 1 {
			return nil, fmt.Errorf("field numbers must be consecutive, but %d is followed by %d in descriptor %s", lastFieldNum, num, descriptor.FullName())
		}
		size, err := r.fieldSize(field)
		if err != nil {
			return nil, err
		}

		fd := &FieldDescriptor{
			offset:          currentOffset,
			size:            size,
			FieldDescriptor: field,
		}

		md.fieldsByName[field.Name()] = fd
		md.fieldsByNumber[num] = fd
		md.fields = append(md.fields, fd)

		currentOffset += size
		lastFieldNum = num
	}

	return md, nil
}

type MessageDescriptor struct {
	protoreflect.MessageDescriptor
	fields         []*FieldDescriptor
	fieldsByName   map[protoreflect.Name]*FieldDescriptor
	fieldsByNumber map[protoreflect.FieldNumber]*FieldDescriptor
}

func (z *MessageDescriptor) FieldOffset(field protoreflect.FieldDescriptor) (uint32, error) {
	fd, ok := z.fieldsByName[field.Name()]
	if !ok {
		return 0, fmt.Errorf("field %s not found", field.Name())
	}
	return fd.offset, nil
}

func (r *Registry) fieldSize(field protoreflect.FieldDescriptor) (uint32, error) {
	switch field.Kind() {
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

type FieldDescriptor struct {
	protoreflect.FieldDescriptor
	offset uint32
	size   uint32
}

type FieldDescriptors struct {
	protoreflect.FieldDescriptors
	fds        []protoreflect.FieldDescriptor
	byName     map[protoreflect.Name]protoreflect.FieldDescriptor
	byJSONName map[string]protoreflect.FieldDescriptor
	byTextName map[string]protoreflect.FieldDescriptor
	byNumber   map[protoreflect.FieldNumber]protoreflect.FieldDescriptor
}

func NewFieldDescriptors(parent protoreflect.FieldDescriptors, overrides []protoreflect.FieldDescriptor) *FieldDescriptors {
	fds := &FieldDescriptors{
		FieldDescriptors: parent,
		fds:              overrides,
		byName:           map[protoreflect.Name]protoreflect.FieldDescriptor{},
		byJSONName:       map[string]protoreflect.FieldDescriptor{},
		byTextName:       map[string]protoreflect.FieldDescriptor{},
		byNumber:         map[protoreflect.FieldNumber]protoreflect.FieldDescriptor{},
	}

	for _, fd := range overrides {
		fds.byName[fd.Name()] = fd
		fds.byJSONName[fd.JSONName()] = fd
		fds.byTextName[fd.TextName()] = fd
		fds.byNumber[fd.Number()] = fd
	}

	return fds
}

func (f FieldDescriptors) Len() int {
	return len(f.fds)
}

func (f FieldDescriptors) Get(i int) protoreflect.FieldDescriptor {
	return f.fds[i]
}

func (f FieldDescriptors) ByName(s protoreflect.Name) protoreflect.FieldDescriptor {
	return f.byName[s]
}

func (f FieldDescriptors) ByJSONName(s string) protoreflect.FieldDescriptor {
	return f.byJSONName[s]
}

func (f FieldDescriptors) ByTextName(s string) protoreflect.FieldDescriptor {
	return f.byTextName[s]
}

func (f FieldDescriptors) ByNumber(n protoreflect.FieldNumber) protoreflect.FieldDescriptor {
	return f.byNumber[n]
}

var _ protoreflect.FieldDescriptors = FieldDescriptors{}
