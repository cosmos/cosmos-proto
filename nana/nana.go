package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {

	var fNames string
	patcher := flag.NewFlagSet("sew", flag.ExitOnError)
	saveAs := patcher.String("saveAs", "patched.go", "the file to write to disk")
	patcher.StringVar(&fNames, "files", "", "the files to sew together")
	remove := patcher.Bool("D", false, "delete files after patching")

	deleter := flag.NewFlagSet("delete", flag.ExitOnError)
	funcName := deleter.String("func", "", "function to patch -- example: -func nameOfFunc")
	fileName := deleter.String("file", "", "the path to the file to delete -- example: -file path/to/file")

	if len(os.Args) < 3 {
		fmt.Println("not enough arguments passed")
		os.Exit(51)
	}

	if os.Args[1] == "delete" {
		deleter.Parse(os.Args[2:])
		fmt.Println("deleting ", *fileName)
		delete(*fileName, *funcName)
	} else if os.Args[1] == "sew" {
		patcher.Parse(os.Args[2:])
		args := patcher.Args()
		fmt.Println(*saveAs)
		fmt.Println(args)
		patch(*saveAs, *remove, args...)
	}

	fmt.Printf("command '%s' not recognized", os.Args[1])
	os.Exit(50)

}

func patch(saveAs string, deleteFiles bool, files ...string) {
	if len(files) < 1 {
		panic("patch requires at least 2 files to merge")
	}
	file, fset, err := Patch(files[0], files[1:]...)
	if err != nil {
		panic(err)
	}
	bz := getBytesFromFile(fset, file)
	os.WriteFile(saveAs, bz.Bytes(), 0766)
	if deleteFiles {
		for _, f := range files {
			os.Remove(f)
		}
	}
}

func delete(funcName, fileName string) {
	err := RemoveFunction(fileName, funcName)
	if err != nil {
		panic(err)
	}
}

func RemoveFunction(path, function string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	found := false
	for i, v := range f.Decls {
		switch t := v.(type) {
		case *ast.FuncDecl:
			if t.Name.String() == function {
				found = true

				// move method to end
				if i != len(f.Decls)-1 {
					f.Decls[i] = f.Decls[len(f.Decls)-1]
				}

				// drop the last element
				f.Decls = f.Decls[:len(f.Decls)-1]
			}
		}
	}

	if !found {
		return fmt.Errorf("func %s not found", function)
	}

	buf := getBytesFromFile(fset, f)
	err = os.Remove(path)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, buf.Bytes(), 0776)
	if err != nil {
		return err
	}

	return nil
}

func Patch(src string, paths ...string) (*ast.File, *token.FileSet, error) {
	srcFset := token.NewFileSet()
	srcF, err := parser.ParseFile(srcFset, src, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	importMap := make(map[string]*ast.ImportSpec)
	for _, imp := range srcF.Imports {
		importMap[imp.Path.Value] = imp
	}

	for _, path := range paths {
		f, err := parser.ParseFile(srcFset, path, nil, parser.ParseComments)
		if err != nil {
			return nil, nil, err
		}

		importMap = patchImports(f.Imports, importMap)

		err = patchScope(srcF.Scope, f.Scope)
		if err != nil {
			return nil, nil, err
		}

		mergedDecls, err := patchDecls(srcF.Decls, f.Decls)
		if err != nil {
			return nil, nil, err
		}

		srcF.Decls = mergedDecls
	}

	return srcF, srcFset, nil
}

func patchImports(imports []*ast.ImportSpec, importMap map[string]*ast.ImportSpec) map[string]*ast.ImportSpec {
	for _, imp := range imports {
		iport, ok := importMap[imp.Path.Value]
		if !ok {
			importMap[imp.Path.Value] = iport
		}
	}

	return importMap
}

func patchScope(src, dst *ast.Scope) error {
	for k, v := range dst.Objects {
		if _, ok := src.Objects[k]; ok {
			return fmt.Errorf("duplicate scope found %s", k)
		}
		src.Objects[k] = v
	}
	return nil
}

func patchDecls(src, dst []ast.Decl) ([]ast.Decl, error) {
	patched := make([]ast.Decl, 0, len(src)+len(dst))
	patched = append(patched, src...)

	dstImports, ok := dst[0].(*ast.GenDecl)
	if !ok {
		return nil, errors.New("destination did not have imports")
	}
	srcImports, ok := src[0].(*ast.GenDecl)
	if !ok {
		return nil, errors.New("src did not have imports")
	}
	srcImports.Specs = append(srcImports.Specs, dstImports.Specs...)

	src[0] = srcImports
	dst = dst[1:]
	src = append(src, dst...)

	return src, nil
}

func getBytesFromFile(set *token.FileSet, file *ast.File) *bytes.Buffer {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, set, file)
	if err != nil {
		panic(err)
	}
	return &buf
}
