package generator

import "go/ast"

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

// Field is a thing
type Field struct {
	Name string            // name of the field
	Tags map[string]string // tags on the field
}
