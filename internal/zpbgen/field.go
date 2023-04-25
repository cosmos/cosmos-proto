package zpbgen

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (g goGen) getExpr(fd *zeroCopyFieldDescriptor) string {
	offset := fd.offset
	switch fd.Desc.Kind() {
	case protoreflect.BoolKind:
		return fmt.Sprint("m.ctx.ReadBool(", offset, ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return fmt.Sprint("int32(m.ctx.ReadUint32(", offset, "))")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return fmt.Sprint("m.ctx.ReadUint32(", offset, ")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return fmt.Sprint("int64(m.ctx.ReadUint64(", offset, "))")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return fmt.Sprint("m.ctx.ReadUint64(", offset, ")")
	case protoreflect.FloatKind:
		return fmt.Sprint("float32(m.ctx.ReadUint32(", offset, "))")
	case protoreflect.DoubleKind:
		return fmt.Sprint("float64(m.ctx.ReadUint64(", offset, "))")
	default:
		panic(fmt.Errorf("unhandled kind: %v", fd.Desc.Kind()))
	}
}

func (g goGen) hasExpr(fd *zeroCopyFieldDescriptor) string {
	if fd.Desc.IsList() {
		panic("unhandled kind: list")
	}

	if fd.Desc.IsMap() {
		panic("unhandled kind: map")
	}

	switch fd.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind, protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return fmt.Sprint(g.getExpr(fd), " != 0")
	case protoreflect.BoolKind:
		return g.getExpr(fd)
	case protoreflect.StringKind, protoreflect.BytesKind:
		panic("unhandled kind: string, bytes")
	case protoreflect.MessageKind, protoreflect.GroupKind:
		panic("unhandled kind: message, group")
	case protoreflect.EnumKind:
		panic("unhandled kind: enum")
	default:
		panic(fmt.Errorf("unhandled kind: %v", fd.Desc.Kind()))
	}
}
