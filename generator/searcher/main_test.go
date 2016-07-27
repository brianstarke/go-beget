package main

import (
	"testing"

	"github.com/brianstarke/go-beget/generator"
	"github.com/stretchr/testify/assert"
)

func TestGatherFields(t *testing.T) {
	assert := assert.New(t)

	fields := []generator.Field{
		{
			Name: "Foo",
			Tags: map[string]string{
				"db":   "foo",
				"json": "FOO",
			},
		},
		{
			Name: "Bar",
			Tags: map[string]string{
				"db":   "bar",
				"json": "BAR",
			},
		},
		{
			Name: "Baz",
			Tags: map[string]string{},
		},
	}

	f := gatherFields(fields)

	assert.NotNil(f)
	assert.Len(f, 3)

	assert.Equal("Foo", f[0].Name)
	assert.Equal("FOO", f[0].JSONName)
	assert.Equal("foo", f[0].DbName)

	assert.Equal("Bar", f[1].Name)
	assert.Equal("BAR", f[1].JSONName)
	assert.Equal("bar", f[1].DbName)

	assert.Equal("Baz", f[2].Name)
	assert.Equal("Baz", f[2].JSONName)
	assert.Equal("Baz", f[2].DbName)
}
