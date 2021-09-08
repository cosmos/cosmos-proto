package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

const (
	flagFunc   = "func"
	flagFile   = "file"
	flagFiles  = "files"
	flagSaveAs = "saveAs"
	flagDelete = "D"
)

func getDeleteFuncCmd() *cobra.Command {
	var deleteFuncCmd = &cobra.Command{
		Use:   "delete -func <function name> -file <path/to/file>",
		Short: "delete a function from a go src file",
		Long:  `delete a function from the Golang source file AST. Must pass a -func and -file flag.`,
		Run: func(cmd *cobra.Command, args []string) {
			fileName, err := cmd.Flags().GetString(flagFile)
			if err != nil {
				panic(err)
			}

			fnName, err := cmd.Flags().GetString(flagFunc)
			if err != nil {
				panic(err)
			}

			deleteFunc(fnName, fileName)
		},
	}

	deleteFuncCmd.Flags().String(flagFunc, "", "the function to delete")
	deleteFuncCmd.Flags().String(flagFile, "", "the file to look in")

	deleteFuncCmd.MarkFlagRequired(flagFunc)
	deleteFuncCmd.MarkFlagRequired(flagFile)

	return deleteFuncCmd
}

func deleteFunc(funcName, fileName string) {
	err := removeFunc(fileName, funcName)
	if err != nil {
		panic(err)
	}
}

func removeFunc(path, function string) error {
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
	err = os.WriteFile(path, buf.Bytes(), 0766)
	if err != nil {
		return err
	}

	return nil
}

func getPatchCmd() *cobra.Command {
	var patchFileCmd = &cobra.Command{
		Use:   "patch -saveAs <fileName.go> -files <somefile.go> <anotherfile.go> -D=<boolean>",
		Short: "merge a set of files into a source file",
		Long:  "merge the files passed in by -files and -saveAs the given filename. Pass in -D=true to delete the files passed in by -files.",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := cmd.Flags().GetStringSlice(flagFiles)
			if err != nil {
				panic(err)
			}
			saveAs, err := cmd.Flags().GetString(flagSaveAs)
			if err != nil {
				panic(err)
			}
			del, err := cmd.Flags().GetBool(flagDelete)

			patchCmd(saveAs, del, files...)

		},
	}

	patchFileCmd.Flags().StringSlice(flagFiles, []string{}, "the files to merge together")
	patchFileCmd.Flags().String(flagSaveAs, "", "the name of the file to save the merged files as")

	patchFileCmd.MarkFlagRequired(flagFiles)
	patchFileCmd.MarkFlagRequired(flagSaveAs)

	return patchFileCmd
}

func patchCmd(saveAs string, deleteFiles bool, files ...string) {
	if len(files) < 1 {
		panic("patchCmd requires at least 2 files to merge")
	}
	file, fset, err := patch(files[0], files[1:]...)
	if err != nil {
		panic(err)
	}
	bz := getBytesFromFile(fset, file)
	os.WriteFile(saveAs, bz.Bytes(), 0766)
	if deleteFiles {
		for _, f := range files {
			err := os.Remove(f)
			if err != nil {
				panic(err)
			}
		}
	}
}

func patch(src string, paths ...string) (*ast.File, *token.FileSet, error) {
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
