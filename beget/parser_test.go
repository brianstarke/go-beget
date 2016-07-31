package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Searchable string `beget:"search" db:"searchable" json:"searchable"`
	Createable string `beget:"create" db:"creatable" json:"creatable"`
	Updateable string `beget:"update" db:"updateable" json:"updateable"`
	Allable    string `beget:"search,create,update" db:"allable" json:"allable"`
	NotParsed  string
}

func TestFindFile(t *testing.T) {
	assert := assert.New(t)

	// try to find TestStruct contained within this test file
	f, err := findFile("TestStruct")

	assert.NoError(err, "should not error when looking for TestStruct")
	assert.NotNil(f, "should return a file")
}

func TestGetImportPath(t *testing.T) {
	assert := assert.New(t)

	// try to find the import path of this test file
	path, err := getImportPath()

	assert.NoError(err, "should not error when trying to find the import path")
	assert.Equal("github.com/brianstarke/go-beget/beget", path)
}

func TestParseStructFields(t *testing.T) {
	assert := assert.New(t)

	// find the TestStruct contained within this test file
	f, err := findFile("TestStruct")

	assert.NoError(err, "should not error when looking for TestStruct")
	assert.NotNil(f, "should return a file")

	// try to find valid fields
	fields := parseStructFields(f, "TestStruct")

	assert.NoError(err, "should not error when parsing struct fields")
	assert.Len(fields, 4)

	assert.Equal("Searchable", fields[0].Name)
	assert.Equal("search", fields[0].Tags["beget"])

	assert.Equal("Createable", fields[1].Name)
	assert.Equal("create", fields[1].Tags["beget"])

	assert.Equal("Updateable", fields[2].Name)
	assert.Equal("update", fields[2].Tags["beget"])

	assert.Equal("Allable", fields[3].Name)
	assert.Equal("search,create,update", fields[3].Tags["beget"])
}
