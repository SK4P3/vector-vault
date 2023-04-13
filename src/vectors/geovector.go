package vectors

import "math"

func (v *GeoVector) DotProduct(v2 GeoVector) float64 {
	var dotProduct float64 = v[0]*v2[0] + v[1]*v2[1]
	return dotProduct
}

func (v *GeoVector) Magnitude() float64 {
	var sum float64 = v[0]*v[0] + v[1]*v[1]
	return math.Sqrt(sum)
}

func (v *GeoVector) CosineSimilarity(v2 GeoVector) float64 {
	dotProduct := v.DotProduct(v2)
	magnitudes := v.Magnitude() * v2.Magnitude()
	return dotProduct / magnitudes
}
