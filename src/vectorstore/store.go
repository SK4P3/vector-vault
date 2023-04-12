package vectorstore

import "fmt"

func EmptyVectorStore(storeConfig VectorStoreConfig) *VectorStore {
	store := VectorStore{
		config:  storeConfig,
		data:    []dataEntry{{Title: "Test Title", Content: "Test Content", Vector: []float32{0.420}}},
		vectors: []vectorEntry{},
	}

	return &store
}

func PersistentVectorStore(path string) *VectorStore {
	store := VectorStore{
		config: VectorStoreConfig{
			Name:           "",
			PersistantPath: path,
		},
		data:    []dataEntry{},
		vectors: []vectorEntry{},
	}

	store.loadPersistentData()

	return &store
}

func (store *VectorStore) Insert(entry dataEntry) {
	store.data = append(store.data, entry)
}

func (store *VectorStore) Remove(index int) {
	store.data = append(store.data[:index], store.data[index+1:]...)
}

func (store *VectorStore) SimilaritySearch(query []float32, n int) []dataEntry {
	result := []dataEntry{}
	for _, entry := range store.getTopNVectors(query, n) {
		result = append(result, store.data[entry.dataIdx])
	}
	return result
}

func (store *VectorStore) GetByKey(key int) dataEntry {
	return store.data[key]
}

func (store *VectorStore) Dumps() string {
	storeString := "Vectorstore: " + store.config.Name + "\n"
	storeString += "Data: \n"
	storeString += fmt.Sprintf("%#v", store.data)
	storeString += "Vectors: \n"
	storeString += fmt.Sprintf("%#v", store.vectors)
	return storeString
}
