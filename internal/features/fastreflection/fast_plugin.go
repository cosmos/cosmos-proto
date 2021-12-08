package fastreflection

import (
	generator2 "github.com/cosmos/cosmos-proto/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	generator2.RegisterFeature("fast", func(gen *generator2.GeneratedFile, _ *protogen.Plugin) generator2.FeatureGenerator {
		return fastReflectionFeature{
			GeneratedFile: gen,
			Stable:        false,
			once:          false,
		}
	})
}

type fastReflectionFeature struct {
	*generator2.GeneratedFile
	Stable, once bool
}

type fastGenerator struct {
	*generator2.GeneratedFile
	file    *protogen.File
	message *protogen.Message

	Stable   bool
	typeName string
	err      error
}

func newGenerator(f *protogen.File, g *generator2.GeneratedFile, message *protogen.Message) *fastGenerator {
	return &fastGenerator{
		GeneratedFile: g,
		file:          f,
		message:       message,
		Stable:        true,
		typeName:      fastReflectionTypeName(message),
		err:           nil,
	}
}

func (g fastReflectionFeature) GenerateFile(file *protogen.File, _ *protogen.Plugin) bool {
	for _, msg := range file.Messages {
		GenProtoMessage(file, g.GeneratedFile, msg)
	}
	return true // only do this once
}

func (g fastReflectionFeature) GenerateHelpers() {
	// no helpers needed here yet
}
