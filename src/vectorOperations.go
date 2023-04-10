package main

import (
	"errors"
	"sync"
	"time"
)

func dotProduct(a []float32, b []float32) (float32, error) {
	var lengthA int = len(a)

	if lengthA != len(b) {
		return -10, errors.New("Error: Vectors don't have same dimension")
	}

	var product float32 = 0

	for i := 0; i < lengthA; i++ {
		product = product + a[i]*b[i]

	}

	return product, nil
}

func getTopNVectors(query []float32, n int) []similarityEntry {

	defer timeTrack(time.Now(), "Top Vectors")

	similarityChan := make(chan similarityEntry, 1)
	var similarityList []similarityEntry

	var wg sync.WaitGroup

	wg.Add(1)
	go func(queryVector []float32, vectors []vectorEntry, wg *sync.WaitGroup, sampleChan chan similarityEntry) {
		defer wg.Done()
		defer close(sampleChan)

		var w sync.WaitGroup
		w.Add(len(vectors))

		for _, vector := range vectors {
			go calculateSililarity(&w, queryVector, vector, sampleChan)
		}

		w.Wait()
	}(query, vectorEntries, &wg, similarityChan)

	wg.Add(1)
	go collectResult(&wg, similarityChan, &similarityList)
	wg.Wait()

	var similarityLength = len(similarityList)
	var mostSimilarVectors = quickSortSimilarityStart(similarityList)[similarityLength-n : len(similarityList)]

	return mostSimilarVectors
}

func partition(arr []similarityEntry, low, high int) ([]similarityEntry, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].Similarity < pivot.Similarity {
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
