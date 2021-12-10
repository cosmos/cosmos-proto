package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/cosmos/cosmos-proto/features/fastreflection"
	_ "github.com/cosmos/cosmos-proto/features/grpc"
	_ "github.com/cosmos/cosmos-proto/features/pool"
	_ "github.com/cosmos/cosmos-proto/features/protoc"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type ObjectSet map[protogen.GoIdent]bool

func (o ObjectSet) String() string {
	return fmt.Sprintf("%#v", o)
}

func (o ObjectSet) Set(s string) error {
	idx := strings.LastIndexByte(s, '.')
	if idx < 0 {
		return fmt.Errorf("invalid object name: %q", s)
	}

	ident := protogen.GoIdent{
		GoImportPath: protogen.GoImportPath(s[0:idx]),
		GoName:       s[idx+1:],
	}
	o[ident] = true
	return nil
}

func main() {
	var features string
	poolable := make(ObjectSet)

	var f flag.FlagSet
	f.Var(poolable, "pool", "use memory pooling for this object")
	f.StringVar(&features, "features", "all", "list of features to generate (separated by '+')")

	protogen.Options{ParamFunc: f.Set}.Run(func(plugin *protogen.Plugin) error {
		processedMessages := make(map[protoreflect.FullName]struct{})
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}
			for _, message := range file.Messages {
				rewriteMessageField(message, processedMessages)
			}
		}
		return generateAllFiles(plugin, strings.Split(features, "+"), poolable)
	})
}

var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func generateAllFiles(plugin *protogen.Plugin, featureNames []string, poolable ObjectSet) error {
	ext := &generator.Extensions{Poolable: poolable}
	gen, err := generator.NewGenerator(plugin.Files, featureNames, ext)
	if err != nil {
		return err
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		gf := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".pulsar.go", file.GoImportPath)
		gf.P("package ", file.GoPackageName)
		if !gen.GenerateFile(plugin, gf, file) {
			gf.Skip()
		}
	}

	// plugin.SupportedFeatures = SupportedFeatures
	return nil
}

var reservedFieldNames = map[string]struct{}{
	"Descriptor":   {},
	"Type":         {},
	"New":          {},
	"Interface":    {},
	"Range":        {},
	"Has":          {},
	"Clear":        {},
	"Get":          {},
	"Set":          {},
	"Mutable":      {},
	"NewField":     {},
	"WhichOneof":   {},
	"GetUnknown":   {},
	"SetUnknown":   {},
	"IsValid":      {},
	"ProtoMethods": {},
}

func rewriteMessageField(message *protogen.Message, processed map[protoreflect.FullName]struct{}) {
	// skip already processed messages, useful for recursive messages
	if _, done := processed[message.Desc.FullName()]; done {
		return
	}
	// skip map entries
	if message.Desc.IsMapEntry() {
		return
	}

	for _, field := range message.Fields {
		_, reserved := reservedFieldNames[field.GoName]
		if !reserved {
			continue
		}
		log.Printf("Message %s contains the reserved field name %s which conflicts with protoreflect.Message interface implementation.\nThis field will be suffixed with an underscore '_'.\nIf you can change the message field name, please do so.\nIn a future iteration of pulsar we will make a breaking change to this practice in order to be compliant with field naming of the original golang protobuf implementation.", message.Desc.FullName(), field.Desc.FullName())
		field.GoName = field.GoName + "_"
	}
	processed[message.Desc.FullName()] = struct{}{}

	for _, nestedMessage := range message.Messages {
		rewriteMessageField(nestedMessage, processed)
	}
}
