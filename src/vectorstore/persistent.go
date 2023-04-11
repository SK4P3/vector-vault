package vectorstore

import (
	"encoding/gob"
	"fmt"
	"os"
)

func (store *VectorStore) loadPersistentData() {

	dataFile, err := os.Open(store.config.PersistantPath + "data.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configFile, err := os.Open(store.config.PersistantPath + "config.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&store.data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configDecoder := gob.NewDecoder(configFile)
	err = configDecoder.Decode(&store.config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile.Close()
	configFile.Close()

	store.loadVectors()
}

func (store *VectorStore) loadVectors() {
	for index, data := range store.data {
		store.vectors = append(store.vectors, vectorEntry{dataIdx: index, Vector: data.Vector})
	}
}

func (store *VectorStore) Persist() {

	if store.config.PersistantPath == "" {
		fmt.Println("No persistent storage Path Specified")
		os.Exit(1)
	}

	os.MkdirAll(store.config.PersistantPath, os.ModeAppend)

	dataFile, err := os.Create(store.config.PersistantPath + "data.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configFile, err := os.Create(store.config.PersistantPath + "config.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataEncoder := gob.NewEncoder(dataFile)
	configEncoder := gob.NewEncoder(configFile)

	dataEncoder.Encode(store.data)
	configEncoder.Encode(store.config)

	dataFile.Close()
	configFile.Close()
}
