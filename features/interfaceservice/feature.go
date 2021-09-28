package interfaceservice

import (
	cosmos_proto "github.com/cosmos/cosmos-proto"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func init() {
	generator.RegisterFeature("interfaceservice", func(gen *generator.GeneratedFile, plugin *protogen.Plugin) generator.FeatureGenerator {
		return &interfaceService{GeneratedFile: gen}
	})
}

type interfaceService struct {
	*generator.GeneratedFile
}

func (i interfaceService) GenerateFile(file *protogen.File, plugin *protogen.Plugin) bool {
	for _, svc := range file.Services {
		options := svc.Desc.Options()
		xt := proto.GetExtension(options, cosmos_proto.E_InterfaceService).(*cosmos_proto.InterfaceService)
		if xt == nil {
			continue
		}
		err := (&gen{file: file, plugin: plugin, svc: svc, desc: xt, GeneratedFile: i.GeneratedFile}).generate()
		if err != nil {
			plugin.Error(err)
			panic(err) // TODO: do we need to panic?
		}
	}

	return true
}

func (i interfaceService) GenerateHelpers() {

}
