package search

/*
	GENERATED CODE, YOUR EDITS ARE FUTILE

	generated from Thing
	using http://github.com/brianstarke/go-beget/generator/searcher
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// NewThingSearchHandler will return a handler for ThingSearchRequests.
//
// The handler assumes that the ThingSearchRequest will be POSTed using
// JSON.
func NewThingSearchHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sr ThingSearchRequest
		err := c.BindJSON(&sr)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		results, err := NewSQLThingSearcher(db).Search(sr)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
		return
	}
}

// NewThingCountHandler will return a handler for ThingSearchRequests.
//
// The handler assumes that the ThingSearchRequest will be POSTed using
// JSON.
func NewThingCountHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sr ThingSearchRequest
		err := c.BindJSON(&sr)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		result, err := NewSQLThingSearcher(db).Count(sr)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"count": result})
		return
	}
}

// NewThingGetByIDHandler will return a handler for Thing GET requests.
//
// The handler assumes that you have constructed a GET route with a path param of
// `:id` (i.e. `/api/v1/:id`) for the ID and also that the ID field of Thing
// is called ID.
func NewThingGetByIDHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if len(id) == 0 {
			c.JSON(200, gin.H{"error": "No ID provided"})
			return
		}

		result, err := NewSQLThingSearcher(db).GetByField(ThingID, id)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
		return
	}
}
