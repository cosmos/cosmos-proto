package protoc

import (
	generator2 "github.com/cosmos/cosmos-proto/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	generator2.RegisterFeature("protoc", func(gen *generator2.GeneratedFile, plugin *protogen.Plugin) generator2.FeatureGenerator {
		return protocGenGoFeature{
			Plugin:        plugin,
			GeneratedFile: gen,
			once:          false,
		}
	})
}

type protocGenGoFeature struct {
	*protogen.Plugin
	*generator2.GeneratedFile
	once bool
}

func (pg protocGenGoFeature) GenerateFile(file *protogen.File, plugin *protogen.Plugin) bool {
	GenerateFile(plugin, file, pg.GeneratedFile)
	return pg.once
}

func (pg protocGenGoFeature) GenerateHelpers() {} //noop
