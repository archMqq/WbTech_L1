package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func chanWrite(n int) []<-chan int {
	channels := make([]<-chan int, n)
	writeChannels := make([]chan int, n)
	for i := 0; i < n; i++ {
		ch := make(chan int, 10000)
		channels[i] = ch
		writeChannels[i] = ch
	}

	go func() {
		for i := 0; i < 10000; i++ {
			for j, ch := range writeChannels {
				if j == 0 {
					safeSend(ch, i)
				} else {
					res := i
					for k := 0; k < j; k++ {
						res *= i
					}
					safeSend(ch, res)
				}
			}
		}

		for i := 0; i < n; i++ {
			close(writeChannels[i])
		}
	}()

	return channels
}

func safeSend(ch chan<- int, value int) {
	select {
	case ch <- value:
	default:
		log.Println("channel blocked skipping value:", value)
	}
}

func reorderToOut(ctx context.Context, ch <-chan int) {
loop:
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				log.Println("channel closed")
				break loop
			} else {
				os.Stdout.WriteString(fmt.Sprintf("Got value: %d\n", val))
			}
		case <-ctx.Done():
			break loop
		}
	}
}

func main() {
	channels := chanWrite(2)

	ctx, stop := context.WithTimeout(context.Background(), 5*time.Minute)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		reorderToOut(ctx, channels[1])
	}()

	wg.Wait()
}
