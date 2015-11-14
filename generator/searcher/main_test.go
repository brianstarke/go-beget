package main

import (
	"testing"

	"github.com/brianstarke/go-begetter/generator"
	"github.com/stretchr/testify/assert"
)

func TestGatherSearchableFields(t *testing.T) {
	assert := assert.New(t)

	fields := []generator.Field{
		{
			Name: "Foo",
			Tags: map[string]string{
				"beget": "search,create,update",
				"db":    "foo",
				"json":  "FOO",
			},
		},
		{
			Name: "Bar",
			Tags: map[string]string{
				"beget": "create",
				"db":    "bar",
				"json":  "BAR",
			},
		},
		{
			Name: "Baz",
			Tags: map[string]string{
				"beget": "search",
			},
		},
	}

	f := gatherSearchableFields(fields)

	assert.NotNil(f)
	assert.Len(f, 2)

	assert.Equal("Foo", f[0].Name)
	assert.Equal("FOO", f[0].JSONName)
	assert.Equal("foo", f[0].DbName)

	assert.Equal("Baz", f[1].Name)
	assert.Equal("Baz", f[1].JSONName)
	assert.Equal("Baz", f[1].DbName)
}
