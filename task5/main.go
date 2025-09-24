package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func writeChannel(ctx context.Context, ch chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("writing ended")
			return
		default:
			ch <- i
		}
	}
}

func readChannel(ctx context.Context, ch <-chan int) {
	for val := range ch {
		select {
		case <-ctx.Done():
			fmt.Println("reading ended")
			return
		default:
			fmt.Printf("read value: %d\n", val)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no needed 'N' arg")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("'N' has to be int")
	}

	//program stops after 'N' seconds
	ctx, stop := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer stop()

	ch := make(chan int)
	go writeChannel(ctx, ch)
	go readChannel(ctx, ch)

	time.Sleep(time.Duration(n+1) * time.Second)
}
