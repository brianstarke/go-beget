package update

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/updater
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// NewThingUpdateHandler will return a handler for ThingUpdateRequests.
//
// The handler assumes that the ThingUpdateRequest will be PATCHed using
// JSON.
func NewThingUpdateHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ur ThingUpdateRequest
		err := c.BindJSON(&ur)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = NewSQLThingUpdater(db).Update(ur)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
		return
	}
}
