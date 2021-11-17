package pool

import (
	"fmt"

	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	generator.RegisterFeature("pool", func(gen *generator.GeneratedFile, plugin *protogen.Plugin) generator.FeatureGenerator {
		return &pool{GeneratedFile: gen}
	})
}

type pool struct {
	*generator.GeneratedFile
	once bool
}

var _ generator.FeatureGenerator = (*pool)(nil)

func (p *pool) GenerateHelpers() {}

func (p *pool) GenerateFile(file *protogen.File, plugin *protogen.Plugin) bool {
	for _, message := range file.Messages {
		p.message(message)
	}
	return p.once
}

func (p *pool) message(message *protogen.Message) {
	for _, nested := range message.Messages {
		p.message(nested)
	}

	if message.Desc.IsMapEntry() || !p.ShouldPool(message) {
		return
	}

	p.once = true
	ccTypeName := message.GoIdent

	p.P(`var vtprotoPool_`, ccTypeName, ` = `, p.Ident("sync", "Pool"), `{`)
	p.P(`New: func() interface{} {`)
	p.P(`return &`, message.GoIdent, `{}`)
	p.P(`},`)
	p.P(`}`)

	p.P(`func (x *`, ccTypeName, `) ResetVT() {`)
	var saved []*protogen.Field
	for _, field := range message.Fields {
		fieldName := field.GoName

		if field.Desc.IsList() {
			switch field.Desc.Kind() {
			case protoreflect.MessageKind, protoreflect.GroupKind:
				if p.ShouldPool(field.Message) {
					p.P(`for _, mm := range x.`, fieldName, `{`)
					p.P(`mm.ResetVT()`)
					p.P(`}`)
				}
			}
			p.P(fmt.Sprintf("f%d", len(saved)), ` := x.`, fieldName, `[:0]`)
			saved = append(saved, field)
		} else {
			switch field.Desc.Kind() {
			case protoreflect.MessageKind, protoreflect.GroupKind:
				if p.ShouldPool(field.Message) {
					p.P(`x.`, fieldName, `.ReturnToVTPool()`)
				}
			case protoreflect.BytesKind:
				p.P(fmt.Sprintf("f%d", len(saved)), ` := x.`, fieldName, `[:0]`)
				saved = append(saved, field)
			}
		}
	}

	p.P(`x.Reset()`)
	for i, field := range saved {
		p.P(`x.`, field.GoName, ` = `, fmt.Sprintf("f%d", i))
	}
	p.P(`}`)

	p.P(`func (x *`, ccTypeName, `) ReturnToVTPool() {`)
	p.P(`if x != nil {`)
	p.P(`x.ResetVT()`)
	p.P(`vtprotoPool_`, ccTypeName, `.Put(m)`)
	p.P(`}`)
	p.P(`}`)

	p.P(`func `, ccTypeName, `FromVTPool() *`, ccTypeName, `{`)
	p.P(`return vtprotoPool_`, ccTypeName, `.Get().(*`, ccTypeName, `)`)
	p.P(`}`)
}
