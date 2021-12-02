package protoc

import (
	"github.com/cosmos/cosmos-proto/cosmos_proto"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"strings"
)

func isCustomType(field *protogen.Field) bool {
	fd := field.Desc
	t := proto.GetExtension(fd.Options(), cosmos_proto.E_GoType).(string)
	return t != ""
}

func customFieldType(g *generator.GeneratedFile, field *protogen.Field) string {
	if field.Desc.IsList() || field.Desc.IsMap() {
		panic("invalid custom type") // TODO(fdymylja): better err msg
	}
	t := proto.GetExtension(field.Desc.Options(), cosmos_proto.E_GoType).(string)
	switch {
	case !strings.Contains(t, "."):
		return "*" + t
	default:
		s := strings.Split(t, ".")
		typeName := s[len(s)-1]
		pkg := protogen.GoImportPath(strings.Join(s[0:len(s)-1], "."))
		g.Import(pkg)
		return "*" + g.QualifiedGoIdent(pkg.Ident(typeName))
	}
}
