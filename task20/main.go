package main

import (
	"fmt"
	"strings"
)

// Реверс строки через дополнительный слайс
func reverseStrWords(str string) string {
	slice := strings.Split(str, " ")

	if len(slice) == 1 {
		return str
	} else {
		max := len(slice) - 1
		for i := 0; i < len(slice)/2; i++ {
			max -= i
			tm := slice[i]
			slice[i] = slice[max]
			slice[max] = tm
		}
	}

	return strings.Join(slice, " ")
}

// ------
// Реверс слов без доп слайсов (только преобразование строки в слайс)
func reverseStrWordsWithoutSlice(str string) string {
	rev := []rune(str) //возможно использовать тип byte, но только если имеется гарантия использования только ASCII символов
	left, right := 0, len(str)-1

	for left < right {
		rev[left], rev[right] = rev[right], rev[left]
		left++
		right--
	}

	start := 0
	for start < len(rev) {
		end := start
		for end < len(rev) && rev[end] != ' ' {
			end++
		}

		left, right := start, end-1
		for left < right {
			rev[left], rev[right] = rev[right], rev[left]
			left++
			right--
		}
		start = end + 1
	}
	return string(rev)
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverseStrWordsWithoutSlice(str))
	fmt.Println(reverseStrWords(str))
}
