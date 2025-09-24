package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// base out - ending after complition
func basicOut() {
	for i := 0; i < 10; i++ {
		fmt.Println("Working...")
	}
	fmt.Println("goroutine end")
}

// condition out - ending when condition is met
func conditionOut() {
	for i := 0; i < 10; i++ {
		fmt.Println("Working...")
		if i == 5 {
			fmt.Println("condition is met")
			return
		}
	}
}

// channel out - ending when smth is sent to channel
func chanOut(ch <-chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("got message")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(400 * time.Millisecond)
		}
	}
}

// context out - ending when cancel func is called
func ctxOut(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("got context stop")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(400 * time.Millisecond)
		}
	}
}

// channel closing out - ending when closing channel is called
func chanCloseOut(ch <-chan int) {
	for val := range ch {
		fmt.Printf("got value: %d\n", val)
	}
	fmt.Println("channel was closed")
}

// runtime.Goexit out
func goexitOut() {
	defer fmt.Println("goexit was called")

	go func() {
		time.Sleep(500 * time.Microsecond)
		runtime.Goexit()
	}()
}

// timeout - ending when time is out
func timeoutExit() {
	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("time is out")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("basic out")
	go basicOut()
	time.Sleep(2 * time.Second)

	fmt.Println("condition out")
	go conditionOut()
	time.Sleep(2 * time.Second)

	fmt.Println("chan out")
	ch := make(chan struct{})
	go chanOut(ch)
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
	time.Sleep(2 * time.Second)

	fmt.Println("context out")
	ctx, stop := context.WithCancel(context.Background())
	go ctxOut(ctx)
	time.Sleep(1 * time.Second)
	stop()
	time.Sleep(2 * time.Second)

	fmt.Println("chan closing out")
	channel := make(chan int)
	go chanCloseOut(channel)
	channel <- 1
	channel <- 2
	close(channel)

	fmt.Println("goexit out")
	go goexitOut()
	time.Sleep(2 * time.Second)

	fmt.Println("timeout")
	go timeoutExit()
	time.Sleep(3 * time.Second)
}
