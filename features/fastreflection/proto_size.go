package fastreflection

import (
	"fmt"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"strings"
)

func (g *fastGenerator) genSizeMethod() {

	g.P(`size := func(input `, protoifacePkg.Ident("SizeInput"), ") ", protoifacePkg.Ident("SizeOutput"), " {")
	g.P("x := input.Message.Interface().(*", g.message.GoIdent, ")")
	g.P(`if x == nil {`)
	g.P(`return `, protoifacePkg.Ident("SizeOutput"), "{ ")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Size: 0,")
	g.P("}")
	g.P(`}`)
	g.P(`var n int`)
	g.P(`var l int`)
	g.P(`_ = l`)
	oneofs := make(map[string]struct{})
	for _, field := range g.message.Fields {
		oneof := field.Oneof != nil && !field.Oneof.Desc.IsSynthetic()
		if !oneof {
			g.field(true, field)
		} else {
			fieldName := field.Oneof.GoName
			if _, ok := oneofs[fieldName]; !ok {
				// TODO: here we usually cast to the vtprotobuf oneof interface, but we'll unroll manually instead.
				oneofs[fieldName] = struct{}{}
				g.P("switch x := x.", fieldName, ".(type) {")
				for _, ooField := range field.Oneof.Fields {

					g.P("case *", ooField.GoIdent, ": ")
					g.P("if x == nil {")
					g.P("break")
					g.P("}")
					g.field(true, ooField)
				}
				g.P("}")
			}
		}
	}

	// last thing to do
	g.P(`if x.unknownFields != nil {`)
	g.P(`n+=len(x.unknownFields)`)
	g.P(`}`)
	g.P(`return `, protoifacePkg.Ident("SizeOutput"), "{ ")
	g.P("NoUnkeyedLiterals: struct{}{},")
	g.P("Size: n,")
	g.P("}") // TODO: return sizeinput not n
	g.P(`}`)
	g.P()
}

func (g *fastGenerator) field(proto3 bool, field *protogen.Field) {
	fieldname := field.GoName
	nullable := field.Message != nil || (field.Oneof != nil && field.Oneof.Desc.IsSynthetic())
	repeated := field.Desc.Cardinality() == protoreflect.Repeated
	if repeated {
		g.P(`if len(x.`, fieldname, `) > 0 {`)
	} else if nullable {
		g.P(`if x.`, fieldname, ` != nil {`)
	}

	packed := field.Desc.IsPacked()
	wireType := generator.ProtoWireType(field.Desc.Kind())
	fieldNumber := field.Desc.Number()
	if packed {
		wireType = protowire.BytesType
	}
	key := generator.KeySize(fieldNumber, wireType)
	switch field.Desc.Kind() {
	case protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		if packed {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(len(x.`, fieldname, `)*8))`, `+len(x.`, fieldname, `)*8`)
		} else if repeated {
			g.P(`n+=`, strconv.Itoa(key+8), `*len(x.`, fieldname, `)`)
		} else if proto3 && !nullable {
			g.P(`if x.`, fieldname, ` != 0 {`)
			g.P(`n+=`, strconv.Itoa(key+8))
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key+8))
		}
	case protoreflect.DoubleKind:
		if packed {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(len(x.`, fieldname, `)*8))`, `+len(x.`, fieldname, `)*8`)
		} else if repeated {
			g.P(`n+=`, strconv.Itoa(key+8), `*len(x.`, fieldname, `)`)
		} else if proto3 && !nullable {
			g.P(`if x.`, fieldname, ` != 0 || `, mathPackage.Ident("Signbit"), `(x.`, fieldname, `) {`)
			g.P(`n+=`, strconv.Itoa(key+8))
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key+8))
		}
	case protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		if packed {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(len(x.`, fieldname, `)*4))`, `+len(x.`, fieldname, `)*4`)
		} else if repeated {
			g.P(`n+=`, strconv.Itoa(key+4), `*len(x.`, fieldname, `)`)
		} else if proto3 && !nullable {
			g.P(`if x.`, fieldname, ` != 0 {`)
			g.P(`n+=`, strconv.Itoa(key+4))
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key+4))
		}
	case protoreflect.FloatKind:
		if packed {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(len(x.`, fieldname, `)*4))`, `+len(x.`, fieldname, `)*4`)
		} else if repeated {
			g.P(`n+=`, strconv.Itoa(key+4), `*len(x.`, fieldname, `)`)
		} else if proto3 && !nullable {
			g.P(`if x.`, fieldname, ` != 0 || `, mathPackage.Ident("Signbit"), `(float64(x.`, fieldname, `)) {`)
			g.P(`n+=`, strconv.Itoa(key+4))
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key+4))
		}
	case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
		if packed {
			g.P(`l = 0`)
			g.P(`for _, e := range x.`, fieldname, ` {`)
			g.P(`l+=`, runtimePackage.Ident("Sov"), `(uint64(e))`)
			g.P(`}`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(l))+l`)
		} else if repeated {
			g.P(`for _, e := range x.`, fieldname, ` {`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(e))`)
			g.P(`}`)
		} else if nullable {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(*x.`, fieldname, `))`)
		} else if proto3 {
			g.P(`if x.`, fieldname, ` != 0 {`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(x.`, fieldname, `))`)
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(x.`, fieldname, `))`)
		}
	case protoreflect.BoolKind:
		if packed {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(len(x.`, fieldname, `)))`, `+len(x.`, fieldname, `)*1`)
		} else if repeated {
			g.P(`n+=`, strconv.Itoa(key+1), `*len(x.`, fieldname, `)`)
		} else if proto3 && !nullable {
			g.P(`if x.`, fieldname, ` {`)
			g.P(`n+=`, strconv.Itoa(key+1))
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key+1))
		}
	case protoreflect.StringKind:
		if repeated {
			g.P(`for _, s := range x.`, fieldname, ` { `)
			g.P(`l = len(s)`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
			g.P(`}`)
		} else if nullable {
			g.P(`l=len(*x.`, fieldname, `)`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
		} else if proto3 {
			g.P(`l=len(x.`, fieldname, `)`)
			g.P(`if l > 0 {`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
			g.P(`}`)
		} else {
			g.P(`l=len(x.`, fieldname, `)`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
		}
	case protoreflect.GroupKind:
		panic(fmt.Errorf("size does not support group %v", fieldname))
	case protoreflect.MessageKind:
		if field.Desc.IsMap() {
			fieldKeySize := generator.KeySize(field.Desc.Number(), generator.ProtoWireType(field.Desc.Kind()))
			keyKeySize := generator.KeySize(1, generator.ProtoWireType(field.Message.Fields[0].Desc.Kind()))
			valueKeySize := generator.KeySize(2, generator.ProtoWireType(field.Message.Fields[1].Desc.Kind()))
			g.P(`for k, v := range x.`, fieldname, ` { `)
			g.P(`_ = k`)
			g.P(`_ = v`)
			sum := []string{strconv.Itoa(keyKeySize)}

			switch field.Message.Fields[0].Desc.Kind() {
			case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
				sum = append(sum, `8`)
			case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
				sum = append(sum, `4`)
			case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
				sum = append(sum, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Sov")), `(uint64(k))`))
			case protoreflect.BoolKind:
				sum = append(sum, `1`)
			case protoreflect.StringKind, protoreflect.BytesKind:
				sum = append(sum, `len(k)`, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Sov")), `(uint64(len(k)))`))
			case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
				sum = append(sum, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Soz")), `(uint64(k))`))
			}

			switch field.Message.Fields[1].Desc.Kind() {
			case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, strconv.Itoa(8))
			case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, strconv.Itoa(4))
			case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Sov")), `(uint64(v))`))
			case protoreflect.BoolKind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `1`)
			case protoreflect.StringKind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `len(v)`, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Sov")), `(uint64(len(v)))`))
			case protoreflect.BytesKind:
				g.P(`l = `, strconv.Itoa(valueKeySize), ` + len(v)+`, runtimePackage.Ident("Sov"), `(uint64(len(v)))`)
				sum = append(sum, `l`)
			case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, fmt.Sprintf("%s%s", g.QualifiedGoIdent(runtimePackage.Ident("Soz")), `(uint64(v))`))
			case protoreflect.MessageKind:
				g.P(`l = 0`)
				g.P(`if v != nil {`)
				g.messageSize("v", field.Message.Fields[1].Message)
				g.P(`}`)
				g.P(`l += `, strconv.Itoa(valueKeySize), `+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
				sum = append(sum, `l`)
			}
			g.P(`mapEntrySize := `, strings.Join(sum, "+"))
			g.P(`n+=mapEntrySize+`, fieldKeySize, `+`, runtimePackage.Ident("Sov"), `(uint64(mapEntrySize))`)
			g.P(`}`)
		} else if field.Desc.IsList() {
			g.P(`for _, e := range x.`, fieldname, ` { `)
			g.messageSize("e", field.Message)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
			g.P(`}`)
		} else {
			g.messageSize("x."+fieldname, field.Message)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
		}
	case protoreflect.BytesKind:
		if repeated {
			g.P(`for _, b := range x.`, fieldname, ` { `)
			g.P(`l = len(b)`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
			g.P(`}`)
		} else if proto3 {
			g.P(`l=len(x.`, fieldname, `)`)
			g.P(`if l > 0 {`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
			g.P(`}`)
		} else {
			g.P(`l=len(x.`, fieldname, `)`)
			g.P(`n+=`, strconv.Itoa(key), `+l+`, runtimePackage.Ident("Sov"), `(uint64(l))`)
		}
	case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		if packed {
			g.P(`l = 0`)
			g.P(`for _, e := range x.`, fieldname, ` {`)
			g.P(`l+=`, runtimePackage.Ident("Soz"), `(uint64(e))`)
			g.P(`}`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Sov"), `(uint64(l))+l`)
		} else if repeated {
			g.P(`for _, e := range x.`, fieldname, ` {`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Soz"), `(uint64(e))`)
			g.P(`}`)
		} else if nullable {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Soz"), `(uint64(*x.`, fieldname, `))`)
		} else if proto3 {
			g.P(`if x.`, fieldname, ` != 0 {`)
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Soz"), `(uint64(x.`, fieldname, `))`)
			g.P(`}`)
		} else {
			g.P(`n+=`, strconv.Itoa(key), `+`, runtimePackage.Ident("Soz"), `(uint64(x.`, fieldname, `))`)
		}
	default:
		panic("not implemented")
	}
	if repeated || nullable {
		g.P(`}`)
	}
}

func (g *fastGenerator) messageSize(varName string, message *protogen.Message) {
	g.P(`l = `, protoPkg.Ident("Size"), "(", varName, ")")
}
