package main

import (
	"bytes"
	"flag"
	"fmt"
	"sync"

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
TemplateData is comprised of the metadata needed during generation of both
the SearchRequest and the Searcher types.
*/
type TemplateData struct {
	PackageName  string
	TypeName     string
	TableName    string
	Fields       []Field
	InsertFields []Field
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(logPrefix)

	flag.Parse()

	if len(*structName) == 0 || len(*tableName) == 0 {
		log.Fatal("`struct` and `tableName` must be specified")
	}

	log.Printf("Generating methods for %s", ansi.Color(*structName, "155+b"))

	data, err := gatherStructData(*structName)

	if err != nil {
		log.Fatalf("Failed getting struct data for %s: %s", *structName, err.Error())
	}

	pathParts := strings.Split(data.PkgImportPath, "/")

	tmplData := TemplateData{
		PackageName: pathParts[len(pathParts)-1],
		TypeName:    data.Name,
		TableName:   *tableName,
		Fields:      data.Fields,
	}

	// strip out ID, CreatedAt, and anything listed in omitFromInsert flag
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
			tmplData.InsertFields = append(tmplData.InsertFields, f)
		}
	}

	var wg sync.WaitGroup
	wg.Add(4)

	if _, err := os.Stat("../search.generated.go"); os.IsNotExist(err) {
		wg.Add(1)
		go generateFromTemplate("enums", tmplData, &wg)
	}

	go generateFromTemplate("fields", tmplData, &wg)
	go generateFromTemplate("search", tmplData, &wg)
	go generateFromTemplate("create", tmplData, &wg)
	go generateFromTemplate("update", tmplData, &wg)

	wg.Wait()
}

func generateFromTemplate(tmplName string, tmplData TemplateData, wg *sync.WaitGroup) {
	defer wg.Done()

	t, err := templates.Asset(fmt.Sprintf("templates/%s.tmpl", tmplName))

	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(tmplName).Parse(string(t))

	if err != nil {
		panic(err)
	}

	b := []byte{}
	buf := bytes.NewBuffer(b)

	err = tmpl.Execute(buf, tmplData)

	if err != nil {
		panic(err)
	}

	outputBytes, err := format.Source(buf.Bytes())

	if err != nil {
		panic(err)
	}

	var output string

	if tmplName == "enums" {
		output = "search.generated.go"
	} else {
		output = fmt.Sprintf(
			"%s%s.generated.go",
			strings.ToLower(tmplData.TypeName),
			strings.Title(tmplName))
	}

	err = ioutil.WriteFile(output, outputBytes, 0644)

	if err != nil {
		panic(err)
	}

	log.Printf("%s generated \t[%s]", tmplName, ansi.Color(output, "155+b"))
}
