package vectorstore

// it would take ~1.000.000 Years for 18 quintillion liters of water to fall down the Niagara Falls

type DataEntry struct {
	Title   string
	Content string
	Vector  []float32
}

type vectorEntry struct {
	dataIdx int
	Vector  []float32
}

type VectorStore struct {
	config  VectorStoreConfig
	data    []DataEntry
	vectors []vectorEntry
}

type VectorStoreConfig struct {
	Name           string
	PersistantPath string
}
