package blurhash

import "math"

func sRGBToLinear(value int64) float64 {
	v := float64(value / 255.0)
	if v <= 0.04045 {
		return v / 12.92
	} else {
		return math.Pow((v+0.055)/1.055, 2.4)
	}
}

func linearTosRGB(value float64) int64 {
	v := math.Max(0, math.Min(1, value))
	if v <= 0.0031308 {
		return int64(v*12.92*255 + 0.5)
	} else {
		return int64((1.055*math.Pow(v, 1/2.4)-0.055)*255 + 0.5)
	}
}

func signPow(value float64, exp float64) float64 {
	return math.Copysign(math.Pow(math.Abs(value), exp), value)
}

func max(values [][]float64, fromInclusive int32, endExclusive int32) float64 {
	result := math.Inf(-1)
	for i := fromInclusive; i < endExclusive; i++ {
		for j := 0; j < len(values[i]); j++ {
			value := values[i][j]
			if value > result {
				result = value
			}
		}
	}
	return result
}
