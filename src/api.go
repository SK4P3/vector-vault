package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"vector-vault/openai"
	"vector-vault/vectorstore"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	Message string
}

var store *vectorstore.VectorStore

func main() {
	router := gin.Default()

	godotenv.Load(".env")

	store = vectorstore.PersistentVectorStore(os.Getenv("STORE_LOCATION"))

	router.GET("/api/stats", getIndexData)

	router.GET("/api/query/text", searchText)

	router.GET("/api/entry/list", listEntries)

	router.GET("/api/entry", getEntry)
	router.POST("/api/entry", addEntry)
	router.DELETE("/api/entry", removeEntry)
	router.PUT("/api/entry", updateEntry)

	router.Run("localhost:8080")
}

func getIndexData(c *gin.Context) {
	data := store.IndexData()
	c.JSON(http.StatusOK, data)
}

func listEntries(c *gin.Context) {
	from, err := strconv.Atoi(c.Query("from"))
	to, err := strconv.Atoi(c.Query("to"))
	data, err := store.ListData(from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, data)
}

func getEntry(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	entry, err := store.GetEntry(idx)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, entry)
}

func addEntry(c *gin.Context) {
	var entryToAdd vectorstore.DataEntry
	if err := c.BindJSON(&entryToAdd); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	added := store.Insert(entryToAdd)
	c.JSON(http.StatusOK, added)
}

func removeEntry(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	store.Remove(idx)
	c.JSON(http.StatusOK, "Removed")
}

func updateEntry(c *gin.Context) {
	// update single entry with index a
}

func searchText(c *gin.Context) {

}

func searchTextB(c *gin.Context) {
	var user_query string = c.Query("query")
	var user_embedding = openai.EmbeddingAda2(user_query)

	result := store.SimilaritySearch(user_embedding, 1)

	c.JSON(http.StatusOK, result)
}
