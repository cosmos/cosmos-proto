package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/cosmos/cosmos-proto/internal/zpbgen"
)

func main() {
	protogen.Options{}.Run(zpbgen.GoPluginRunner)
}
