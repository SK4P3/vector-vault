package vectorstore

import (
	"fmt"
	"sync"
	"time"
	"vector-index/utils"
)

type similarityEntry struct {
	dataIdx    int
	Similarity float32
}

func calculateSimilarity(wg *sync.WaitGroup, queryVector []float32, storeVector vectorEntry, ch chan similarityEntry) {
	defer wg.Done()
	var similarity, _ = dotProduct(queryVector, storeVector.Vector)
	ch <- similarityEntry{dataIdx: storeVector.dataIdx, Similarity: similarity}
}

func collectResult(wg *sync.WaitGroup, channel chan similarityEntry, similarityList *[]similarityEntry) {
	defer wg.Done()
	for s := range channel {
		*similarityList = append(*similarityList, s)
	}
}

func (store *VectorStore) getTopNVectors(query []float32, n int) []similarityEntry {

	defer utils.TimeTrack(time.Now(), "Top Vectors by query")

	similarityChan := make(chan similarityEntry, 1)
	var similarityList []similarityEntry

	var wg sync.WaitGroup

	wg.Add(1)
	go func(queryVector []float32, vectors []vectorEntry, wg *sync.WaitGroup, resultChan chan similarityEntry) {
		defer wg.Done()
		defer close(resultChan)

		var w sync.WaitGroup
		w.Add(len(vectors))

		for _, vector := range vectors {
			go calculateSimilarity(&w, queryVector, vector, resultChan)
		}

		w.Wait()
	}(query, store.vectors, &wg, similarityChan)

	wg.Add(1)
	go collectResult(&wg, similarityChan, &similarityList)
	wg.Wait()

	var mostSimilarVectors = quickSortSimilarityStart(similarityList)[:n]

	fmt.Println(mostSimilarVectors)

	return mostSimilarVectors
}

func partition(arr []similarityEntry, low, high int) ([]similarityEntry, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].Similarity > pivot.Similarity {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSortSimilarity(arr []similarityEntry, low, high int) []similarityEntry {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortSimilarity(arr, low, p-1)
		arr = quickSortSimilarity(arr, p+1, high)
	}
	return arr
}

func quickSortSimilarityStart(arr []similarityEntry) []similarityEntry {
	return quickSortSimilarity(arr, 0, len(arr)-1)
}
