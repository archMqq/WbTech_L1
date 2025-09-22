package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func infinityData(channel chan<- int) {
	for i := 0; ; i++ {
		channel <- i
	}
}

func worker(id int, channel <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case val, ok := <-channel:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			os.Stdout.WriteString(fmt.Sprintf("worker %d read info: %d\n", id, val))
		case <-ctx.Done():
			fmt.Println("stop signal")
			return
		}
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()

	go infinityData(channel)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i+1, channel, wg, ctx)
	}

	wg.Wait()
	fmt.Println("all workers done")
}
