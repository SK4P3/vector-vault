package vectors

import "math"

func (v *FloatVector) DotProduct(v2 FloatVector) float64 {
	var dotProduct float64 = 0
	for idx, f := range *v {
		dotProduct += f * v2[idx]
	}
	return dotProduct
}

func (v *FloatVector) Magnitude() float64 {
	var sum float64 = 0
	for _, f := range *v {
		sum += f * f
	}
	return math.Sqrt(sum)
}

func (v *FloatVector) Normalize() FloatVector {
	vector := make(FloatVector, len(*v))
	magnitude := v.Magnitude()
	for idx, f := range *v {
		vector[idx] = f / magnitude
	}
	return vector
}

func (v *FloatVector) CosineSimilarity(v2 FloatVector) float64 {
	dotProduct := v.DotProduct(v2)
	magnitudes := v.Magnitude() * v2.Magnitude()
	return dotProduct / magnitudes
}

func (v *FloatVector) CosineSimilarityUnitVector(v2 FloatVector) float64 {
	return v.DotProduct(v2)
}
