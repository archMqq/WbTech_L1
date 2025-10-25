package main

import (
	"fmt"
	"unicode"
)

func isUnique(s string) bool {
	sl := []rune(s) //можно байт слайс, но в задании нет ограничения на латиницу
	symbols := make(map[rune]struct{})
	for _, v := range sl {
		symb := unicode.ToLower(v)
		if _, ok := symbols[symb]; ok {
			return false
		}
		symbols[symb] = struct{}{}
	}

	return true
}

func main() {
	a, b, c := "abcdefgh", "abcdA", "abcda"

	fmt.Println(isUnique(a), isUnique(b), isUnique(c))
}
