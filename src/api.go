package main

import (
	"fmt"
	"net/http"
	"vector-vault/vectors"
	"vector-vault/vectorstore"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()

	// godotenv.Load(".env")

	// router.GET("/query/text", searchText)

	// router.Run("localhost:8080")

	var vector1 = vectors.GeoVector{}
	vector1[0] = 47.384350
	vector1[1] = 13.462740

	var vector2 = vectors.GeoVector{}
	vector2[0] = 47.384148
	vector2[1] = 13.462639

	store := vectorstore.PersistentVectorStore("./mockStore/")

	fmt.Println(vector1.CosineSimilarity(vector2))
	fmt.Println(len(store.GetByKey(0).Vector))
}

// todo use vector type
func searchText(c *gin.Context) {

	store := vectorstore.PersistentVectorStore("./mockStore/")

	var user_query string = c.Query("query")
	var user_embedding = getOpenAIEmbedding(user_query)

	result := store.SimilaritySearch(user_embedding, 1)

	c.IndentedJSON(http.StatusOK, result)
}
