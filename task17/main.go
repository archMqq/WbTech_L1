package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func binSearch[T Number](sl []T, searchable T) int {
	low, high := 0, len(sl)-1

	for low <= high {
		guess := (low + high) / 2
		val := sl[guess]

		switch {
		case val == searchable:
			return guess
		case val > searchable:
			high = guess - 1
		case val < searchable:
			low = guess + 1
		}
	}

	return -1
}

func main() {
	sl := []float64{1, 46, 453, 34576, 325252}

	fmt.Println(binSearch(sl, 1))
}
