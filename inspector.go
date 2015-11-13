package generator

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
)

/*
StructData contains data about this struct that is used during the code
generation, such as the struct name and package name such that we could
generate things like

	package foo

	type BarSearchRequest struct {
		// ...
	}

StructData also collects the Fields contained with the struct such that we
can generate a list of fields that are approved for search or create or update
requests like

 	type BarField int

	const (
		Baz BarField = iota
		Boom
		Biff
		Bang
	)
*/
type StructData struct {
	Name          string    // the name of the struct we're looking for
	PkgImportPath string    // the import path for this type
	File          *ast.File // parsed file containing our struct def
	Fields        []Field   // the fields def
	inType        bool      // private helper for when we are walking the AST
}

/*
Field contains the parsed metadata about a struct field, like it's Name
and a map of it's Tags.
*/
type Field struct {
	Name string            // name of the field
	Tags map[string]string // tags on the field
}

/*
GatherStructData finds the correct file containing the struct in question,
parses out the pertinents, and then returns the struct data to be used for
generating files.
*/
func GatherStructData(structName string) (*StructData, error) {
	return nil, nil
}

/*
Searches all the files in this package directory, parses them in to go/ast
(see http://golang.org/pkg/go/ast/#File) and returns the file that contains
the struct name we're looking for.
*/
func findFile(structName string) (*ast.File, error) {
	pkg, err := build.Default.ImportDir(".", 0)

	if err != nil {
		return nil, fmt.Errorf("Cannot process directory . : %s", err.Error())
	}

	// build a list of all *.go files in this package directory
	var filenames []string
	filenames = append(filenames, pkg.GoFiles...)

	// search files for the specified struct definition
	fs := token.NewFileSet()

	for _, name := range filenames {
		f, err := parser.ParseFile(fs, name, nil, 0)

		if err != nil {
			return nil, fmt.Errorf("Cannot parse file %s, %s", name, err.Error())
		}

		if containsStruct(f, structName) {
			return f, nil
		}
	}

	return nil, fmt.Errorf("Type %s not found in any package files", structName)
}

/*
Check to see if a file contains the struct definition we seek.
*/
func containsStruct(f *ast.File, structName string) bool {
	hasIt := false

	ast.Inspect(f, func(n ast.Node) bool {
		tspec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if tspec.Name.String() == structName {
			hasIt = true
		}

		return true
	})

	return hasIt
}
