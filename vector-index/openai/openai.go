package openai

import (
	"context"
	"fmt"
	"os"
	"time"
	"vector-index/utils"

	openaiAip "github.com/sashabaranov/go-openai"
)

func EmbeddingAda2(query string) []float32 {
	defer utils.TimeTrack(time.Now(), "Get Embedding")

	client := openaiAip.NewClient(os.Getenv("OPENAPI_KEY"))

	resp, err := client.CreateEmbeddings(
		context.Background(),
		openaiAip.EmbeddingRequest{
			Input: []string{query},
			Model: openaiAip.AdaEmbeddingV2,
			User:  "test",
		},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	return resp.Data[0].Embedding
}
