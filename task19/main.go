package main

import "fmt"

func reverse(input string) string {
	r := []rune(input)
	var res []rune

	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}

	return string(res)
}

func main() {
	str := "dsgsшувгыфаiuefs"

	fmt.Println(reverse(str))
}
