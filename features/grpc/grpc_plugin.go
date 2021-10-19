// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package grpc

import (
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "1.1.0-vtproto"

var requireUnimplementedAlways = true
var requireUnimplemented = &requireUnimplementedAlways

func init() {
	generator.RegisterFeature("grpc", func(gen *generator.GeneratedFile, _ *protogen.Plugin) generator.FeatureGenerator {
		return &grpc{gen}
	})
}

type grpc struct {
	*generator.GeneratedFile
}

func (g *grpc) GenerateFile(file *protogen.File, _ *protogen.Plugin) bool {
	if len(file.Services) == 0 {
		return false
	}

	generateFileContent(nil, file, g.GeneratedFile)
	return true
}

func (g *grpc) GenerateHelpers() {}
