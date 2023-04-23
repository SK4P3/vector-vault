package main

import (
	"fmt"
	"net/http"
	"strconv"
	"vector-index/openai"
	"vector-index/vectorstore"

	"github.com/gin-gonic/gin"
)

type GetTopNEntriesRequest struct {
	QueryVector []float32
	TopN        int
}

type GetTopNEntriesResponse struct {
	Entries []vectorstore.DataEntry
}

func Queries(g *gin.RouterGroup) {
	g.POST("/embedding", getTopNEntries)
	g.GET("/text", searchText)
}

// QueryEmbedding
// @Summary Query vector store using an embedding vector
// @Description Query vector store using an embedding vector and get top N results besed on cosine similarity, faster as there is no extra requst to the embedding api
// @Tags Query
// @Accept json
// @Produce json
// @Param request body GetTopNEntriesRequest true "Query Vector, and TopN Results"
// @Success 200 {object} GetTopNEntriesResponse
// @Router /query/embedding [post]
func getTopNEntries(c *gin.Context) {
	var body GetTopNEntriesRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}

	var user_embedding = body.QueryVector
	var topn = body.TopN

	result := store.SimilaritySearch(user_embedding, topn)

	c.JSON(http.StatusOK, result)
}

// QueryText
// @Summary Query vector store using a string
// @Description Query vector store using a string that will be converted to an embedding vector and get top N results besed on cosine similarity
// @Tags Query
// @Accept json
// @Produce json
// @Param query query string true "query string"
// @Param topn query int true "top n results"
// @Success 200 {object} []vectorstore.DataEntry
// @Router /query/text [get]
func searchText(c *gin.Context) {
	var user_query string = c.Query("query")
	topn, _ := strconv.Atoi(c.Query("topn"))
	var user_embedding = openai.EmbeddingAda2(user_query)

	result := store.SimilaritySearch(user_embedding, topn)

	c.JSON(http.StatusOK, result)
}
