package main

import "fmt"

func sliceIntersection[T comparable](a []T, b []T) []T {
	res := make([]T, 0)
	hash := make(map[T]struct{})

	for _, val := range a {
		hash[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := hash[val]; ok {
			res = append(res, val)
		}
	}

	return res
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 4}

	fmt.Println(sliceIntersection(a, b))
}
