package main

import "fmt"

func quickSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}

	pivot := sl[len(sl)/2]

	less, equal, bigger := []int{}, []int{}, []int{}

	for _, v := range sl {
		if v > pivot {
			bigger = append(bigger, v)
		} else if v < pivot {
			less = append(less, v)
		} else {
			equal = append(equal, v)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(bigger)...)
}

func main() {
	tc := []int{3, 6, 8, 10, 1, 2, 1}

	fmt.Printf("Исходный: %v\n", tc)
	fmt.Printf("Отсортированный: %v\n\n", quickSort(tc))
}
