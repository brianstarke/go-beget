package main

import (
	"bytes"
	"flag"
	"fmt"

	"go/build"
	"go/format"

	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/brianstarke/go-beget/generator"
	"github.com/brianstarke/go-beget/templates"
	"github.com/mgutz/ansi"
)

var (
	structName = flag.String("struct", "", "name of the struct to generate a creator for")
	tableName  = flag.String("table", "", "SQL table name if you want to generate SQLCreator")
	impls      = flag.String("impls", "", "comma separated list of implementations you'd like to generate (only 'sql' and 'gin' available at this time)")

	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("creator", "168") + "] "
)

/*
CreatableField holds the parsed values from the tags present on the struct's
fields.

If there are no `db` or `json` tags provided, we will just use the Field Name
value for those purposes.
*/
type CreatableField struct {
	Name     string // the name of the field in Go
	DbName   string // the name of the field in the db
	JSONName string // the name of the field in json
}

/*
TemplateData is comprised of the metadata needed during generation of both
the SearchRequest and the Searcher types.
*/
type TemplateData struct {
	PackageName     string
	TypeImport      string
	SearchImport    string
	TypeName        string
	TableName       string
	CreatableFields []CreatableField
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(logPrefix)

	flag.Parse()

	if len(*structName) == 0 {
		log.Fatal("`struct` must be specified")
	}

	log.Printf("Generating creator for %s", ansi.Color(*structName, "155+b"))

	pkg, _ := build.Default.ImportDir("../", 0)

	data, err := generator.GatherStructData(*structName)

	if err != nil {
		log.Fatalf("Failed getting struct data for %s: %s", *structName, err.Error())
	}

	// pull out the searchable fields
	creatableFields := gatherCreatableFields(data.Fields)

	tmplData := TemplateData{
		PackageName:     pkg.Name,
		TypeImport:      data.PkgImportPath,
		SearchImport:    strings.Replace(data.PkgImportPath, "/types", "/search", 1),
		TypeName:        data.Name,
		TableName:       *tableName,
		CreatableFields: creatableFields,
	}

	if len(*impls) > 0 {
		// contains works for now, since there aren't any conflicts yet

		if strings.Contains(*impls, "sql") {
			createSQLCreator(tmplData)
		}

		if strings.Contains(*impls, "gin") {
			createGinHandlers(tmplData)
		}
	}
}

func gatherCreatableFields(fields []generator.Field) []CreatableField {
	var creatableFields []CreatableField

	for _, field := range fields {
		bTags, ok := field.Tags["beget"]

		// the field may have other tags instead of `beget`, just pass through
		if !ok {
			continue
		}

		// is this a searchable field
		if strings.Contains(bTags, "create") {
			var dbName, jsonName string

			// if no db tag we will just use the field name as the db field name
			dbTags, ok := field.Tags["db"]

			if !ok {
				dbName = field.Name
			} else {
				// strip out omitempty and commas
				dbName = strings.Replace(dbTags, ",", "", -1)
				dbName = strings.Replace(dbName, "omitempty", "", -1)
			}

			// same thing for json tags
			jsonTags, ok := field.Tags["json"]

			if !ok {
				jsonName = field.Name
			} else {
				// strip out omitempty and commas
				jsonName = strings.Replace(jsonTags, ",", "", -1)
				jsonName = strings.Replace(jsonName, "omitempty", "", -1)
			}

			creatableFields = append(creatableFields, CreatableField{
				Name:     field.Name,
				DbName:   dbName,
				JSONName: jsonName,
			})
		}
	}

	return creatableFields
}

func createSQLCreator(tmplData TemplateData) {
	t, err := templates.Asset("templates/sqlCreator.tmpl")

	tmpl, err := template.New("sqlCreator").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = tmpl.Execute(buf, tmplData)

	if err != nil {
		log.Fatal(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		log.Fatal(err)
	}

	createCreateDirectory()

	output := fmt.Sprintf("../create/%sSQLCreator.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SQLCreator generated %s", ansi.Color(output, "155+b"))
}

func createGinHandlers(tmplData TemplateData) {
	t, err := templates.Asset("templates/ginCreator.tmpl")

	tmpl, err := template.New("ginCreator").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = tmpl.Execute(buf, tmplData)

	if err != nil {
		log.Fatal(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		log.Fatal(err)
	}

	createCreateDirectory()

	output := fmt.Sprintf("../create/%sGinCreator.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GinCreator generated %s", ansi.Color(output, "155+b"))
}

func createCreateDirectory() error {
	if _, err := os.Stat("../create"); os.IsNotExist(err) {
		log.Printf("%s does not exist, I'll create it", ansi.Color("../create", "155+b"))

		os.Mkdir("../create", 0777)
	}

	return nil
}
