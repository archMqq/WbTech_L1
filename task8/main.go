package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needed arg")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("needed int arg")
		return
	}

	newN := n ^ 1
	fmt.Printf("New int: %d", newN)
}
