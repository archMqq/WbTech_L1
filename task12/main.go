package main

import "fmt"

func getUnique(sl []string) []string {
	uniq, res := make(map[string]struct{}), make([]string, 0)

	for _, v := range sl {
		if _, ok := uniq[v]; !ok {
			uniq[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getUnique(sl))
}
