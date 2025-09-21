package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func infinityData(channel chan<- int) {
	for i := 0; ; i++ {
		channel <- i
	}
}

func worker(id int, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range channel {
		os.Stdout.WriteString(fmt.Sprintf("worker %d read info: %d\n", id, val))
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no args")
	}

	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan int)
	wg := &sync.WaitGroup{}

	go infinityData(channel)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i+1, channel, wg)
	}

	wg.Wait()
	fmt.Println("all workers done")
}
