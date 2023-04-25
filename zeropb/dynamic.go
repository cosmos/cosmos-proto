package zeropb

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func NewMessageType(descriptor *MessageDescriptor) protoreflect.MessageType {
	return &messageType{descriptor}
}

type messageType struct {
	desc *MessageDescriptor
}

func (m messageType) New() protoreflect.Message {
	return &Message{}
}

func (m messageType) Zero() protoreflect.Message {
	//TODO implement me
	panic("implement me")
}

func (m messageType) Descriptor() protoreflect.MessageDescriptor {
	return m.desc
}

type Message struct {
	typ     messageType
	data    *BufferContext
	ext     map[protoreflect.FieldNumber]protoreflect.FieldDescriptor
	unknown protoreflect.RawFields
}

var (
	_ protoreflect.Message      = (*Message)(nil)
	_ protoreflect.ProtoMessage = (*Message)(nil)
	_ protoiface.MessageV1      = (*Message)(nil)
)

func (m *Message) ProtoReflect() protoreflect.Message {
	return m
}

func (m *Message) Reset() {
	m.ext = map[protoreflect.FieldNumber]protoreflect.FieldDescriptor{}
	m.unknown = nil
	panic("TODO")
}

func (m *Message) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (m *Message) ProtoMessage() {}

func (m *Message) Descriptor() protoreflect.MessageDescriptor {
	return m.typ.desc
}

func (m *Message) Type() protoreflect.MessageType {
	return m.typ
}

func (m *Message) New() protoreflect.Message {
	return m.typ.New()
}

func (m *Message) Interface() protoreflect.ProtoMessage {
	return m
}

func (m *Message) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	for _, field := range m.typ.desc.fields {
		if !f(field, m.Get(field)) {
			break
		}
	}
}

func (m *Message) Has(descriptor protoreflect.FieldDescriptor) bool {
	field := m.typ.desc.fieldsByNumber[descriptor.Number()]
	if field == nil {
		return false
	}

	if field.IsMap() {
		return false
	}

	if field.IsList() {
		panic("TODO dynamic list")
	}

	switch field.Kind() {
	case protoreflect.BoolKind:
		return m.data.ReadBool(field.offset)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return m.data.ReadInt32(field.offset) != 0
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return m.data.ReadInt64(field.offset) != 0
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return m.data.ReadUint32(field.offset) != 0
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return m.data.ReadUint64(field.offset) != 0
	case protoreflect.FloatKind:
		return m.data.ReadFloat32(field.offset) != 0
	case protoreflect.DoubleKind:
		return m.data.ReadFloat64(field.offset) != 0
	case protoreflect.StringKind:
		return m.data.ReadString(field.offset) != ""
	case protoreflect.BytesKind:
		return m.data.ReadBytes(field.offset) != nil
	case protoreflect.EnumKind:
		return m.data.ReadUint8(field.offset) != 0
	case protoreflect.MessageKind, protoreflect.GroupKind:
		panic("TODO dynamic message")
	default:
		panic("unknown kind")
	}
}

func (m *Message) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	field := m.typ.desc.fieldsByNumber[descriptor.Number()]
	if field == nil {
		return protoreflect.Value{}
	}

	if field.IsMap() {
		return protoreflect.Value{}
	}

	if field.IsList() {
		return protoreflect.ValueOfList(&dynamicList{
			desc: field,
			data: m.data.ResolvePointer(field.offset),
		})
	}

	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(m.data.ReadBool(field.offset))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(m.data.ReadInt32(field.offset))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(m.data.ReadInt64(field.offset))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(m.data.ReadUint32(field.offset))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(m.data.ReadUint64(field.offset))
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(m.data.ReadFloat32(field.offset))
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(m.data.ReadFloat64(field.offset))
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(m.data.ReadString(field.offset))
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes(m.data.ReadBytes(field.offset))
	case protoreflect.MessageKind, protoreflect.GroupKind:
		panic("TODO dynamic message")
	default:
		panic("unknown kind")
	}
}

func (m *Message) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	field := m.typ.desc.fieldsByNumber[descriptor.Number()]
	if field == nil {
		panic("unknown field")
	}

	if field.IsMap() {
		panic("maps aren't supported")
	}

	if field.IsList() {
		panic("TODO dynamic list")
	}

	switch field.Kind() {
	case protoreflect.BoolKind:
		m.data.SetBool(field.offset, value.Bool())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		m.data.SetInt32(field.offset, int32(value.Int()))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		m.data.SetInt64(field.offset, value.Int())
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		m.data.SetUint32(field.offset, uint32(value.Uint()))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		m.data.SetUint64(field.offset, value.Uint())
	case protoreflect.FloatKind:
		m.data.SetFloat32(field.offset, float32(value.Float()))
	case protoreflect.DoubleKind:
		m.data.SetFloat64(field.offset, value.Float())
	case protoreflect.StringKind:
		m.data.SetString(field.offset, value.String())
	case protoreflect.BytesKind:
		m.data.SetBytes(field.offset, value.Bytes())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		panic("TODO dynamic message")
	default:
		panic("unknown kind")
	}
}

func (m *Message) Clear(descriptor protoreflect.FieldDescriptor) {
	field := m.typ.desc.fieldsByNumber[descriptor.Number()]
	if field == nil {
		panic("unknown field")
	}

	if field.IsMap() {
		panic("maps aren't supported")
	}

	if field.IsList() {
		panic("TODO dynamic list")
	}

	switch field.Kind() {
	case protoreflect.BoolKind:
		m.data.SetBool(field.offset, false)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		m.data.SetInt32(field.offset, 0)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		m.data.SetInt64(field.offset, 0)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		m.data.SetUint32(field.offset, 0)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		m.data.SetUint64(field.offset, 0)
	case protoreflect.FloatKind:
		m.data.SetFloat32(field.offset, 0)
	case protoreflect.DoubleKind:
		m.data.SetFloat64(field.offset, 0)
	case protoreflect.EnumKind:
		m.data.SetUint8(field.offset, 0)
	case protoreflect.StringKind:
		m.data.ClearPointer(field.offset)
	case protoreflect.BytesKind:
		m.data.ClearPointer(field.offset)
	default:
		panic("unknown kind")
	}
}

func (m *Message) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (m *Message) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (m *Message) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	//TODO implement me
	panic("implement me")
}

func (m *Message) GetUnknown() protoreflect.RawFields {
	return m.unknown
}

func (m *Message) SetUnknown(fields protoreflect.RawFields) {
	m.unknown = fields
}

func (m *Message) IsValid() bool {
	return true
}

func (m *Message) ProtoMethods() *protoiface.Methods {
	return nil
}

func (m *Message) Unmarshal(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
	//TODO implement me
	panic("implement me")
}

type dynamicList struct {
	desc *FieldDescriptor
	data *BufferContext
}

func (d dynamicList) Len() int {
	return int(d.data.ReadUint16(0))
}

func (d dynamicList) Get(i int) protoreflect.Value {
	if i < 0 || i >= d.Len() {
		return protoreflect.Value{}
	}
	panic("TODO implement me")
}

func (d dynamicList) Set(i int, value protoreflect.Value) {
	//TODO implement me
	panic("implement me")
}

func (d dynamicList) Append(value protoreflect.Value) {
	//TODO implement me
	panic("implement me")
}

func (d dynamicList) AppendMutable() protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (d dynamicList) Truncate(i int) {
	//TODO implement me
	panic("implement me")
}

func (d dynamicList) NewElement() protoreflect.Value {
	//TODO implement me
	panic("implement me")
}

func (d dynamicList) IsValid() bool {
	//TODO implement me
	panic("implement me")
}
