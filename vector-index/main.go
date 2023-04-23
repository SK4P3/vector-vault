package main

import (
	"fmt"
	"net/http"
	"os"
	"vector-index/openai"
	"vector-index/vectorstore"

	_ "vector-index/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// todo refactor api to use only http response
type HttpResponse struct {
	Message string
}

type ErrorResponse struct {
	Message string
}

type GetEmbeddingRequest struct {
	Content string
}

type GetEmbeddingResponse struct {
	Embedding []float32
}

var store *vectorstore.VectorStore

func main() {

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	godotenv.Load(".env")

	store = vectorstore.PersistentVectorStore(os.Getenv("STORE_LOCATION"))

	router.GET("/health", health)
	router.GET("/stats", getIndexData)

	router.POST("/embedding/ada", getEmbeddingAda2)

	v1 := router.Group("/")
	{
		Entires(v1.Group("/entry"))
		Queries(v1.Group("/query"))
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8081")
}

// HealthCheck
// @Summary Check if Index is online
// @Description Ping Index for a simple Helath Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {string} Im Ok! Thanks for asking!
// @Router /health [get]
func health(c *gin.Context) {
	c.JSON(http.StatusOK, "Im Ok! Thanks for asking!")
}

// IndexData
// @Summary Get some basic Index data
// @Description Get Index data including Name, Entries, Size, VectorType, VectorDimension
// @Tags Stats
// @Accept json
// @Produce json
// @Success 200 {object} vectorstore.IndexData
// @Router /stats [get]
func getIndexData(c *gin.Context) {
	data := store.IndexData()
	c.JSON(http.StatusOK, data)
}

// IndexData
// @Summary Get OpenAI ada2 Embedding
// @Description Get the vector representation of a given Input string accordint to the OpenAI ada2 embedding model
// @Tags OpenAI
// @Accept json
// @Produce json
// @Param request body GetEmbeddingRequest true "Strig to convert to Embedding"
// @Success 200 {object} GetEmbeddingResponse
// @Router /embedding/ada [post]
func getEmbeddingAda2(c *gin.Context) {
	var body GetEmbeddingRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Error: " + fmt.Sprint(err)})
		return
	}
	user_embedding := openai.EmbeddingAda2(body.Content)
	c.JSON(http.StatusOK, GetEmbeddingResponse{Embedding: user_embedding})
}
