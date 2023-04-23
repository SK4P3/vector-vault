package main

import (
	"fmt"
	"net/http"
	"strconv"
	"vector-index/vectorstore"

	"github.com/gin-gonic/gin"
)

type AddEntryRequest struct {
	Entry vectorstore.DataEntry
}

type AddEntryResponse struct {
	Idx int
}

func Entires(g *gin.RouterGroup) {
	g.GET("/list", listEntries)
	g.GET("", getEntry)
	g.POST("", addEntry)
	g.DELETE("", removeEntry)
	g.PUT("", updateEntry)
}

// ListEntries
// @Summary List entries in the given range
// @Description Get the data entries in the Vector store from 'from' to 'to'
// @Tags Entry
// @Accept json
// @Produce json
// @Param from query int true "from idx"
// @Param to query int true "to idx"
// @Success 200 {object} []vectorstore.DataEntry
// @Router /entry/list [get]
func listEntries(c *gin.Context) {
	from, err := strconv.Atoi(c.Query("from"))
	to, err := strconv.Atoi(c.Query("to"))
	data, err := store.ListData(from, to+1)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetEntry
// @Summary Get a specific Entry
// @Description Get Entry at index i from the store data
// @Tags Entry
// @Accept json
// @Produce json
// @Param idx query int true "Entry idx"
// @Success 200 {object} vectorstore.DataEntry
// @Router /entry [get]
func getEntry(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	entry, err := store.GetEntry(idx)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, entry)
}

// AddEntry
// @Summary Add an entry to the vector store
// @Description Append a data entry to the vector store and save the updated index version to persistent storage
// @Tags Entry
// @Accept json
// @Produce json
// @Param request body AddEntryRequest true "Entry to add"
// @Success 200 {object} AddEntryResponse
// @Router /entry [post]
func addEntry(c *gin.Context) {
	var body AddEntryRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	added := store.Insert(body.Entry)
	c.JSON(http.StatusOK, added)
	store.Persist()
}

// DeleteEntry
// @Summary Remove a specific Entry
// @Description Remove Entry at index i from the store and save the updated index version to persistent storage
// @Tags Entry
// @Accept json
// @Produce json
// @Param idx query int true "Entry idx to delete"
// @Success 200 {string} Removed
// @Router /entry [delete]
func removeEntry(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	store.Remove(idx)
	c.JSON(http.StatusOK, "Removed")
	store.Persist()
}

func updateEntry(c *gin.Context) {
	// update single entry with index a
}
