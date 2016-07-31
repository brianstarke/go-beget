package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	One   string `db:"field_one" json:"fieldOne"`
	Two   string `db:"field_two" json:"fieldTwo"`
	Three string
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
	assert.Len(fields, 3)

	assert.Equal("One", fields[0].Name)
	assert.Equal("fieldOne", fields[0].JSONName)
	assert.Equal("field_one", fields[0].DbName)

	assert.Equal("Two", fields[1].Name)
	assert.Equal("fieldTwo", fields[1].JSONName)
	assert.Equal("field_two", fields[1].DbName)

	assert.Equal("Three", fields[2].Name)
	assert.Equal("three", fields[2].JSONName)
	assert.Equal("three", fields[2].DbName)
}
