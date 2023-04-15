package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"vector-index/openai"
	"vector-index/vectorstore"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// todo refactor api to use only http response
type HttpResponse struct {
	Message string
}

type ErrorResponse struct {
	Message string
}

var store *vectorstore.VectorStore

func main() {

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	godotenv.Load(".env")

	store = vectorstore.PersistentVectorStore(os.Getenv("STORE_LOCATION"))

	router.GET("/api/health", health)
	router.GET("/api/stats", getIndexData)

	router.POST("/api/embedding/ada", getEmbeddingAda2)
	router.GET("/api/query/text", searchText)

	router.GET("/api/entry/list", listEntries)

	router.GET("/api/entry", getEntry)
	router.POST("/api/entry", addEntry)
	router.DELETE("/api/entry", removeEntry)
	router.PUT("/api/entry", updateEntry)

	router.Run("localhost:8080")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, "Im Ok! Thanks for asking!")
}

func getIndexData(c *gin.Context) {
	data := store.IndexData()
	c.JSON(http.StatusOK, data)
}

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

func getEntry(c *gin.Context) {
	idx, err := strconv.Atoi(c.Query("idx"))
	entry, err := store.GetEntry(idx)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	c.JSON(http.StatusOK, entry)
}

type AddEntryRequest struct {
	// todo make array
	Entry vectorstore.DataEntry
}

type AddEntryResponse struct {
	Idx int
}

func addEntry(c *gin.Context) {
	var body AddEntryRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	added := store.Insert(body.Entry)
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

type GetEmbeddingRequest struct {
	Content string
}

type GetEmbeddingResponse struct {
	Embedding []float32
}

func getEmbeddingAda2(c *gin.Context) {
	var body GetEmbeddingRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	user_embedding := openai.EmbeddingAda2(body.Content)
	c.JSON(http.StatusOK, GetEmbeddingResponse{Embedding: user_embedding})
}

func searchTextB(c *gin.Context) {
	var user_query string = c.Query("query")
	var user_embedding = openai.EmbeddingAda2(user_query)

	result := store.SimilaritySearch(user_embedding, 1)

	c.JSON(http.StatusOK, result)
}
