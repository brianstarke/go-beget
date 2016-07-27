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
	structName = flag.String("struct", "", "name of the struct to generate a searcher for")
	tableName  = flag.String("table", "", "SQL table name")

	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("searcher", "159") + "] "
)

/*
SearchableField holds the parsed values from the tags present on the struct's
fields.

If there are no `db` or `json` tags provided, we will just use the Field Name
value for those purposes.
*/
type SearchableField struct {
	Name     string // the name of the field in Go
	DbName   string // the name of the field in the db
	JSONName string // the name of the field in json
}

/*
TemplateData is comprised of the metadata needed during generation of both
the SearchRequest and the Searcher types.
*/
type TemplateData struct {
	PackageName      string
	TypeImport       string
	SearchImport     string
	TypeName         string
	TableName        string
	SearchableFields []SearchableField
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(logPrefix)

	flag.Parse()

	if len(*structName) == 0 {
		log.Fatal("`struct` must be specified")
	}

	log.Printf("Generating searcher for %s", ansi.Color(*structName, "155+b"))

	pkg, _ := build.Default.ImportDir("../", 0)

	data, err := generator.GatherStructData(*structName)

	if err != nil {
		log.Fatalf("Failed getting struct data for %s: %s", *structName, err.Error())
	}

	// pull out the searchable fields
	searchableFields := gatherSearchableFields(data.Fields)

	tmplData := TemplateData{
		PackageName:      pkg.Name,
		TypeImport:       data.PkgImportPath,
		SearchImport:     strings.Replace(data.PkgImportPath, "/types", "/search", 1),
		TypeName:         data.Name,
		TableName:        *tableName,
		SearchableFields: searchableFields,
	}

	createSearchDirectory()
	createSearcher(tmplData)
}

func gatherSearchableFields(fields []generator.Field) []SearchableField {
	var searchableFields []SearchableField

	for _, field := range fields {
		bTags, ok := field.Tags["beget"]

		// the field may have other tags instead of `beget`, just pass through
		if !ok {
			continue
		}

		// is this a searchable field
		if strings.Contains(bTags, "search") {
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

			searchableFields = append(searchableFields, SearchableField{
				Name:     field.Name,
				DbName:   dbName,
				JSONName: jsonName,
			})
		}
	}

	return searchableFields
}

func createSearcherEnums() {
	if _, err := os.Stat("../search/searchEnums.go"); !os.IsNotExist(err) {
		return
	}

	t, err := templates.Asset("templates/searcher_enums.tmpl")

	if err != nil {
		panic(err)
	}

	searchRequestTmpl, err := template.New("searcher_enums").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = searchRequestTmpl.Execute(buf, nil)

	if err != nil {
		panic(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("../search/searcherEnums.go")
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("Generated search_enums [%s]", ansi.Color(output, "155+b"))
}

func createSearcher(tmplData TemplateData) {
	createSearcherEnums()

	t, err := templates.Asset("templates/searcher.tmpl")

	tmpl, err := template.New("searcher").Parse(string(t))

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

	output := fmt.Sprintf("../search/%sSearcher.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Searcher generated [%s]", ansi.Color(output, "155+b"))
}

func createSearchDirectory() error {
	if _, err := os.Stat("../search"); os.IsNotExist(err) {
		log.Printf("%s does not exist, I'll create it", ansi.Color("../search", "155+b"))

		os.Mkdir("../search", 0777)
	}

	return nil
}
