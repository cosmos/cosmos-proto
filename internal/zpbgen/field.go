package zpbgen

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-proto/zeropb"
)

func (g goGen) getExpr(fd *zeropb.FieldDescriptor, receiverName string) string {
	offset := fd.offset
	switch fd.Desc.Kind() {
	case protoreflect.BoolKind:
		return fmt.Sprint(receiverName, ".ctx.ReadBool(", offset, ")")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return fmt.Sprint("int32(", receiverName, ".ctx.ReadUint32(", offset, "))")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return fmt.Sprint(receiverName, ".ctx.ReadUint32(", offset, ")")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return fmt.Sprint("int64(", receiverName, ".ctx.ReadUint64(", offset, "))")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return fmt.Sprint(receiverName, ".ctx.ReadUint64(", offset, ")")
	case protoreflect.FloatKind:
		return fmt.Sprint("float32(", receiverName, ".ctx.ReadUint32(", offset, "))")
	case protoreflect.DoubleKind:
		return fmt.Sprint("float64(", receiverName, ".ctx.ReadUint64(", offset, "))")
	default:
		return fmt.Sprintf("panic(\"unhandled kind: %v\")", fd.Desc.Kind())
	}
}

func (g goGen) hasExpr(fd *zeropb.FieldDescriptor, receiverName string) string {
	if fd.Desc.IsList() {
		return fmt.Sprintf("panic(\"unhandled kind: list\")")
	}

	if fd.Desc.IsMap() {
		return fmt.Sprintf("panic(\"unhandled kind: map\")")
	}

	switch fd.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind, protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return fmt.Sprint(g.getExpr(fd, receiverName), " != 0")
	case protoreflect.BoolKind:
		return g.getExpr(fd, receiverName)
	default:
		return fmt.Sprintf("panic(\"unhandled kind: %v\")", fd.Desc.Kind())
	}
}

func protoReflectValueExpr(field protoreflect.FieldDescriptor, expr string) string {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return fmt.Sprintf("protoreflect.ValueOfBool(%s)", expr)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return fmt.Sprintf("protoreflect.ValueOfInt32(%s)", expr)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return fmt.Sprintf("protoreflect.ValueOfUint32(%s)", expr)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return fmt.Sprintf("protoreflect.ValueOfInt64(%s)", expr)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return fmt.Sprintf("protoreflect.ValueOfUint64(%s)", expr)
	case protoreflect.FloatKind:
		return fmt.Sprintf("protoreflect.ValueOfFloat32(%s)", expr)
	case protoreflect.DoubleKind:
		return fmt.Sprintf("protoreflect.ValueOfFloat64(%s)", expr)
	case protoreflect.StringKind:
		return fmt.Sprintf("protoreflect.ValueOfString(%s)", expr)
	case protoreflect.BytesKind:
		return fmt.Sprintf("protoreflect.ValueOfBytes(%s)", expr)
	case protoreflect.EnumKind:
		return fmt.Sprintf("protoreflect.ValueOfEnum(%s)", expr)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return fmt.Sprintf("protoreflect.ValueOfMessage(%s)", expr)
	default:
		panic(fmt.Sprintf("unhandled kind: %v", field.Kind()))
	}
}
