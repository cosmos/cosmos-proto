package main

import (
	"fmt"

	"github.com/cosmos/cosmos-proto/features/fastreflection"
	"google.golang.org/protobuf/compiler/protogen"
)

// TODO(fdymylja): remove me when pulsar codegen is fixed

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			genFile := plugin.NewGeneratedFile(fmt.Sprintf("%s.fasteflect.go", f.GeneratedFilenamePrefix), f.GoImportPath)
			genFile.P("package ", f.GoPackageName)
			for _, msg := range f.Messages {
				fastreflection.GenProtoMessage(f, genFile, msg)
			}
		}

		return nil
	})

}
