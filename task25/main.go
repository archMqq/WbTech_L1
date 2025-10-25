package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	//по заданию вроде как требовалось реализовать именно так (канал + горутина),
	// но неэффективно, поскольку создается доп канал
	timer := time.NewTimer(duration)
	ch := make(chan struct{})

	go func() {
		<-timer.C
		ch <- struct{}{}
		close(ch)
	}()

	<-ch
}

func sleepSmpl(duration time.Duration) { //эквивалентен коду выше, но не создает лишний канал
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	sleep(time.Second * 3)
	fmt.Println("heheh")

	sleepSmpl(time.Second * 3)
	fmt.Println("geg")
}
