package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
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
type structData struct {
	Name          string    // the name of the struct we're looking for
	PkgImportPath string    // the import path for this type
	File          *ast.File // parsed file containing our struct def
	Fields        []Field   // the fields def
}

// Field holds struct field metadata
type Field struct {
	Name     string // the name of the field in Go
	DbName   string // the name of the field in the db
	JSONName string // the name of the field in json
}

/*
GatherStructData finds the correct file containing the struct in question,
parses out the pertinents, and then returns the struct data to be used for
generating files.
*/
func gatherStructData(structName string) (*structData, error) {
	file, err := findFile(structName)

	if err != nil {
		return nil, err
	}

	importPath, err := getImportPath()

	if err != nil {
		return nil, err
	}

	fields := parseStructFields(file, structName)

	structData := &structData{
		Name:          structName,
		PkgImportPath: importPath,
		File:          file,
		Fields:        fields,
	}

	return structData, nil
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
	filenames = append(filenames, pkg.TestGoFiles...)

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

/*
Attempts to find the import path we will use to import the struct type in to
the generated files.
*/
func getImportPath() (string, error) {
	pwd, err := filepath.Abs(".")

	if err != nil {
		return "", fmt.Errorf("Error getting absolute filepath: %s", err.Error())
	}

	pwd, err = filepath.EvalSymlinks(pwd)

	if err != nil {
		return "", fmt.Errorf("Error evaluating filepath symlinks: %s", err.Error())
	}

	startPos := strings.LastIndex(pwd, "src/")

	if startPos == -1 {
		return "", fmt.Errorf("Couldn't figure out what the package import path is from %s", pwd)
	}

	return pwd[startPos+4 : len(pwd)], nil
}

/*
Parses the fields from the specified type
*/
func parseStructFields(file *ast.File, structName string) []Field {
	var withinStruct bool
	var fields []Field

	ast.Inspect(file, func(n ast.Node) bool {
		switch t := n.(type) {

		case *ast.TypeSpec:
			// If we're within the struct we're looking for, set the withinStruct
			// flag so we start parsing fields
			if t.Name.String() == structName {
				withinStruct = true
			}
			return true

		case *ast.StructType:
			if !withinStruct {
				return true
			}

			for _, field := range t.Fields.List {
				var v string

				if field.Tag != nil {
					v = field.Tag.Value
				}

				fields = append(fields, Field{
					Name:     field.Names[0].String(),
					DbName:   parseDbName(field.Names[0].String(), v),
					JSONName: parseJSONName(field.Names[0].String(), v),
				})
			}
		}

		return true
	})

	return fields
}

// TODO refactor these

// parse the db field name out of the tag, if there's no db tag just
// lowercase the fieldName
func parseDbName(fieldName, tagValue string) string {
	for _, t := range strings.Fields(strings.Replace(tagValue, "`", "", -1)) {
		s := strings.Split(t, ":")

		if s[0] == "db" {
			v := strings.Replace(s[1], `"`, "", -1)
			v = strings.Replace(v, ",", "", -1)
			return strings.Replace(v, "omitempty", "", -1)
		}
	}

	return strings.ToLower(fieldName)
}

// parse the db field name out of the tag, if there's no db tag just
// lowercase the fieldName
func parseJSONName(fieldName, tagValue string) string {
	for _, t := range strings.Fields(strings.Replace(tagValue, "`", "", -1)) {
		s := strings.Split(t, ":")

		if s[0] == "json" {
			v := strings.Replace(s[1], `"`, "", -1)
			v = strings.Replace(v, ",", "", -1)
			return strings.Replace(v, "omitempty", "", -1)
		}
	}

	return strings.ToLower(fieldName)
}
