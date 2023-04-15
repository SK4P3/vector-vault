package vectorstore

import (
	"errors"
	"fmt"
	"reflect"
	"vector-index/utils"
)

type IndexData struct {
	Name            string
	Entries         int
	Size            string
	VectorType      string
	VectorDimension int
}

func EmptyVectorStore(storeConfig VectorStoreConfig) *VectorStore {
	store := VectorStore{
		config:  storeConfig,
		data:    []DataEntry{},
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
		data:    []DataEntry{},
		vectors: []vectorEntry{},
	}

	store.loadPersistentData()

	return &store
}

func (store *VectorStore) IndexData() IndexData {
	size, _ := utils.DirSize(store.config.PersistantPath)
	data := IndexData{
		Name:            store.config.Name,
		Entries:         len(store.data),
		Size:            fmt.Sprint(size, " bytes"),
		VectorType:      fmt.Sprint(reflect.TypeOf(store.data[0].Vector)),
		VectorDimension: len(store.data[0].Vector),
	}
	return data
}

func (store *VectorStore) ListData(from int, to int) ([]DataEntry, error) {
	length := len(store.data)
	if from > length || to > length {
		return nil, errors.New("Error: Index out of bounds!")
	}
	return store.data[from:to], nil
}

func (store *VectorStore) GetEntry(idx int) (DataEntry, error) {
	length := len(store.data) - 1
	if idx > length {
		return DataEntry{}, errors.New("Error: Index out of bounds!")
	}
	return store.data[idx], nil
}

func (store *VectorStore) Insert(entry DataEntry) int {
	store.data = append(store.data, entry)
	return len(store.data) - 1
}

func (store *VectorStore) Remove(index int) {
	store.data = append(store.data[:index], store.data[index+1:]...)
}

func (store *VectorStore) SimilaritySearch(query []float32, n int) []DataEntry {
	result := []DataEntry{}
	for _, entry := range store.getTopNVectors(query, n) {
		result = append(result, store.data[entry.dataIdx])
	}
	return result
}

func (store *VectorStore) Dumps() string {
	storeString := "Vectorstore: " + store.config.Name + "\n"
	storeString += "Data: \n"
	storeString += fmt.Sprintf("%#v", store.data)
	storeString += "Vectors: \n"
	storeString += fmt.Sprintf("%#v", store.vectors)
	return storeString
}
