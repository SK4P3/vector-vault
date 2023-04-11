package vectorstore

import "errors"

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
