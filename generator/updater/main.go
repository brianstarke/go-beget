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
	structName = flag.String("struct", "", "name of the struct to generate an updater for")
	tableName  = flag.String("table", "", "SQL table name if you want to generate SQLUpdater")
	impls      = flag.String("impls", "", "comma separated list of implementations you'd like to generate (only 'sql' and 'gin' available at this time)")

	logPrefix = "[" +
		ansi.Color("go-beget", "154") +
		"/" +
		ansi.Color("updater", "111") + "] "
)

/*
UpdateableField holds the parsed values from the tags present on the struct's
fields.

If there are no `db` or `json` tags provided, we will just use the Field Name
value for those purposes.
*/
type UpdateableField struct {
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
	UpdateableFields []UpdateableField
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(logPrefix)

	flag.Parse()

	if len(*structName) == 0 {
		log.Fatal("`struct` must be specified")
	}

	log.Printf("Generating updater for %s", ansi.Color(*structName, "155+b"))

	pkg, _ := build.Default.ImportDir("../", 0)

	data, err := generator.GatherStructData(*structName)

	if err != nil {
		log.Fatalf("Failed getting struct data for %s: %s", *structName, err.Error())
	}

	// pull out the updateable fields
	udpateableFields := gatherUpdateableFields(data.Fields)

	tmplData := TemplateData{
		PackageName:      pkg.Name,
		TypeImport:       data.PkgImportPath,
		SearchImport:     strings.Replace(data.PkgImportPath, "/types", "/update", 1),
		TypeName:         data.Name,
		TableName:        *tableName,
		UpdateableFields: udpateableFields,
	}

	createUpdateRequest(tmplData)

	if len(*impls) > 0 {
		// contains works for now, since there aren't any conflicts yet

		if strings.Contains(*impls, "sql") {
			createSQLUpdater(tmplData)
		}

		if strings.Contains(*impls, "gin") {
			createGinHandlers(tmplData)
		}
	}
}

func gatherUpdateableFields(fields []generator.Field) []UpdateableField {
	var updateableFields []UpdateableField

	for _, field := range fields {
		bTags, ok := field.Tags["beget"]

		// the field may have other tags instead of `beget`, just pass through
		if !ok {
			continue
		}

		// is this a searchable field
		if strings.Contains(bTags, "update") {
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

			updateableFields = append(updateableFields, UpdateableField{
				Name:     field.Name,
				DbName:   dbName,
				JSONName: jsonName,
			})
		}
	}

	return updateableFields
}

func createUpdateRequest(tmplData TemplateData) {
	t, err := templates.Asset("templates/updateRequest.tmpl")

	if err != nil {
		panic(err)
	}

	updateRequestTmpl, err := template.New("updateRequest").Parse(string(t))

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

	createUpdateDirectory()

	output := fmt.Sprintf("../update/%sUpdateRequest.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("UpdateRequest generated %s", ansi.Color(output, "155+b"))
}

func createSQLUpdater(tmplData TemplateData) {
	t, err := templates.Asset("templates/sqlUpdater.tmpl")

	tmpl, err := template.New("sqlUpdater").Parse(string(t))

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

	output := fmt.Sprintf("../update/%sSQLUpdater.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SQLUpdater generated %s", ansi.Color(output, "155+b"))
}

func createGinHandlers(tmplData TemplateData) {
	t, err := templates.Asset("templates/ginUpdater.tmpl")

	tmpl, err := template.New("ginUpdater").Parse(string(t))

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

	output := fmt.Sprintf("../update/%sGinUpdater.go", strings.ToLower(tmplData.TypeName))
	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GinUpdater generated %s", ansi.Color(output, "155+b"))
}

func createUpdateDirectory() error {
	if _, err := os.Stat("../update"); os.IsNotExist(err) {
		log.Printf("%s does not exist, I'll create it", ansi.Color("../update", "155+b"))

		os.Mkdir("../update", 0777)
	}

	return nil
}
