package create

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/creator
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	types "github.com/brianstarke/go-beget/example/types"
)

// NewThingCreateHandler will return a handler for Thing creates.
//
// The handler assumes that a Thing will be PUT using JSON.
func NewThingCreateHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t types.Thing
		err := c.BindJSON(&t)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		id, err := NewSQLThingCreator(db).Create(t)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"id": id})
		return
	}
}
