package main

import (
	"context"
	"fmt"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func timeTrack(start time.Time, name string) {
	fmt.Printf("%s took %.12fs \n", name, time.Since(start).Seconds())
}

func getOpenAIEmbedding(query string) []float32 {
	defer timeTrack(time.Now(), "Get Embedding")

	client := openai.NewClient(os.Getenv("OPENAPI_KEY"))

	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: []string{query},
			Model: openai.AdaEmbeddingV2,
			User:  "test",
		},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	return resp.Data[0].Embedding
}
