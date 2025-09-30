package main

import (
	"fmt"
	"math"
)

func setSlices(sl []float64) map[float64][]float64 {
	res := make(map[float64][]float64, 0)

	calcRef := func(ref float64) float64 {
		if ref < 0 {
			return 10 * math.Ceil(ref/10)
		} else {
			return 10 * math.Floor(ref/10)
		}
	}

	for _, val := range sl {
		ref := calcRef(val)

		if res[ref] == nil {
			res[ref] = make([]float64, 0)
		}
		res[ref] = append(res[ref], val)
	}

	return res
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := setSlices(temps)
	fmt.Println(result)
}
