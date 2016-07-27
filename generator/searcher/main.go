package main

import (
	"bytes"
	"flag"
	"fmt"

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
Field holds the parsed values from the tags present on the struct's
fields.

If there are no `db` or `json` tags provided, we will just use the Field Name
value for those purposes.
*/
type Field struct {
	Name     string // the name of the field in Go
	DbName   string // the name of the field in the db
	JSONName string // the name of the field in json
}

/*
TemplateData is comprised of the metadata needed during generation of both
the SearchRequest and the Searcher types.
*/
type TemplateData struct {
	PackageName string
	TypeName    string
	TableName   string
	Fields      []Field
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(logPrefix)

	flag.Parse()

	if len(*structName) == 0 {
		log.Fatal("`struct` must be specified")
	}

	log.Printf("Generating searcher for %s", ansi.Color(*structName, "155+b"))

	data, err := generator.GatherStructData(*structName)

	if err != nil {
		log.Fatalf("Failed getting struct data for %s: %s", *structName, err.Error())
	}

	// pull out the searchable fields
	fields := gatherFields(data.Fields)

	pathParts := strings.Split(data.PkgImportPath, "/")

	tmplData := TemplateData{
		PackageName: pathParts[len(pathParts)-1],
		TypeName:    data.Name,
		TableName:   *tableName,
		Fields:      fields,
	}

	createSearcher(tmplData)
}

func gatherFields(f []generator.Field) []Field {
	var fields []Field

	for _, field := range f {

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

		fields = append(fields, Field{
			Name:     field.Name,
			DbName:   dbName,
			JSONName: jsonName,
		})
	}

	return fields
}

func createSearcherEnums(tmplData TemplateData) {
	if _, err := os.Stat("../searchRequestEnums.go"); !os.IsNotExist(err) {
		return
	}

	t, err := templates.Asset("templates/searchRequestEnums.tmpl")

	if err != nil {
		panic(err)
	}

	searchRequestTmpl, err := template.New("searchRequestEnums").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = searchRequestTmpl.Execute(buf, tmplData)

	if err != nil {
		panic(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("searchRequestEnums.go")
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("Generated searchRequestEnums [%s]", ansi.Color(output, "155+b"))
}

func createSearcher(tmplData TemplateData) {
	createSearcherEnums(tmplData)

	t, err := templates.Asset("templates/searchRequest.tmpl")

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

	output := fmt.Sprintf("%sSearchRequest.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SearchRequest generated [%s]", ansi.Color(output, "155+b"))
}
