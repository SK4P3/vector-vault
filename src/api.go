package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	godotenv.Load(".env")

	router.GET("/query/text", searchText)

	router.Run("localhost:8080")
}

func calculateSililarity(wg *sync.WaitGroup, queryVector []float32, storeVector vectorEntry, ch chan similarityEntry) {
	defer wg.Done()
	var similarity, _ = dotProduct(queryVector, storeVector.Vector)
	ch <- similarityEntry{ID: storeVector.ID, Similarity: similarity}
}

func collectResult(wg *sync.WaitGroup, channel chan similarityEntry, similarityList *[]similarityEntry) {
	defer wg.Done()
	for s := range channel {
		*similarityList = append(*similarityList, s)
	}
}

func searchText(c *gin.Context) {
	var user_query string = c.Query("query")
	var user_embedding = getOpenAIEmbedding(user_query)

	var topVectors = getTopNVectors(user_embedding, 1)

	c.IndentedJSON(http.StatusOK, topVectors)
}
