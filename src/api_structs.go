package main

type dataEntry struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"artist"`
}

type vectorEntry struct {
	ID     string    `json:"id"`
	Vector []float32 `json:"vector"`
}

type similarityEntry struct {
	ID         string
	Similarity float32
}
