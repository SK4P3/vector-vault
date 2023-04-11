package main

import (
	"net/http"
	"vector-vault/vectorstore"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	godotenv.Load(".env")

	router.GET("/query/text", searchText)

	router.Run("localhost:8080")

	// fmt.Println(store.Dumps())
}

func searchText(c *gin.Context) {

	store := vectorstore.PersistentVectorStore("./mockStore/")

	var user_query string = c.Query("query")
	var user_embedding = getOpenAIEmbedding(user_query)

	result := store.SimilaritySearch(user_embedding, 1)

	c.IndentedJSON(http.StatusOK, result)
}
