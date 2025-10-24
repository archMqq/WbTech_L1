package main

import "fmt"

func DeleteItem[T any](slice []T, index int) []T {
	copy(slice[index:], slice[index+1:])
	var zero T
	slice[len(slice)-1] = zero

	return slice[:len(slice)-1]
}

func main() {
	s := []int{1, 2, 3}
	s = DeleteItem(s, 1)
	fmt.Println(s)
}
