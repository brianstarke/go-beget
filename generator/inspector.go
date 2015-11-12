package generator

import "go/ast"

// GeneratorInfo is a thing
type GeneratorInfo struct {
	PkgImportPath string    // the import path for this type
	TypeName      string    // name of the type we're looking for
	File          *ast.File // parsed file containing our struct def
	Fields        []Field   // the fields def
	inType        bool      // private helper for when we are walking the AST
}

// Field is a thing
type Field struct {
	Name string
	Tags map[string]string
}
