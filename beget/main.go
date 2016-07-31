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

	"github.com/brianstarke/go-beget/templates"
	"github.com/mgutz/ansi"
)

var (
	structName = flag.String("struct", "", "name of the struct to generate a searcher for")
	tableName  = flag.String("table", "", "SQL table name")
	omitInsert = flag.String("omitFromInsert", "", "Comma separated list of fields to omit from insert in the create function")

	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("generator", "159") + "] "
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

	if len(*structName) == 0 || len(*tableName) == 0 {
		log.Fatal("`struct` and `tableName` must be specified")
	}

	log.Printf("Generating searcher for %s", ansi.Color(*structName, "155+b"))

	data, err := gatherStructData(*structName)

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

	createSearch(tmplData)
	createCreate(tmplData)
}

func gatherFields(f []structField) []Field {
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

func createSearchEnums(tmplData TemplateData) {
	if _, err := os.Stat("../searchEnums.go"); !os.IsNotExist(err) {
		return
	}

	t, err := templates.Asset("templates/searchEnums.tmpl")

	if err != nil {
		panic(err)
	}

	searchRequestTmpl, err := template.New("searchEnums").Parse(string(t))

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

	output := fmt.Sprintf("searchEnums.go")
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("Generated searchRequestEnums [%s]", ansi.Color(output, "155+b"))
}

func createSearch(tmplData TemplateData) {
	createSearchEnums(tmplData)

	t, err := templates.Asset("templates/search.tmpl")

	tmpl, err := template.New("search").Parse(string(t))

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

	output := fmt.Sprintf("%sSearch.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Search generated [%s]", ansi.Color(output, "155+b"))
}

func createCreate(tmplData TemplateData) {
	t, err := templates.Asset("templates/create.tmpl")

	tmpl, err := template.New("create").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	// strip out ID, CreatedAt, and anything listed in omitFromInsert flag
	var createableFields []Field
	omitFields := map[string]struct{}{
		"ID":        {},
		"CreatedAt": {},
	}

	if len(*omitInsert) > 0 {
		for _, s := range strings.Split(*omitInsert, ",") {
			omitFields[s] = struct{}{}
		}
	}

	for _, f := range tmplData.Fields {
		if _, ok := omitFields[f.Name]; !ok {
			createableFields = append(createableFields, f)
		}
	}

	tmplData.Fields = createableFields

	err = tmpl.Execute(buf, tmplData)

	if err != nil {
		log.Fatal(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("%sCreate.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Create generated [%s]", ansi.Color(output, "155+b"))
}

func createUpdateRequest(tmplData TemplateData) {
	t, err := templates.Asset("templates/update.tmpl")

	if err != nil {
		panic(err)
	}

	updateRequestTmpl, err := template.New("update").Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = updateRequestTmpl.Execute(buf, tmplData)

	if err != nil {
		panic(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("%sUpdateRequest.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("UpdateRequest generated %s", ansi.Color(output, "155+b"))
}
