package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func swap(a, b *int) {
	(*a) ^= (*b)
	(*b) ^= (*a)
	(*a) ^= (*b)
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 1, 546
	fmt.Println(swap2(a, b))
	swap(&a, &b)
	fmt.Println(a, b)

	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
