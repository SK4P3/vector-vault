package vectors

func (v *GptEmbeddingVector) DotProduct(v2 GptEmbeddingVector) float64 {
	var dotProduct float64 = 0
	for idx, f := range *v {
		dotProduct += f * v2[idx]
	}
	return dotProduct
}

func (v *GptEmbeddingVector) CosineSimilarity(v2 GptEmbeddingVector) float64 {
	return v.DotProduct(v2)
}
