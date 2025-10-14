package main

import (
	"sync"
	"sync/atomic"
)

type Incrementor struct {
	value atomic.Uint32
}

func (inc *Incrementor) Increase() {
	inc.value.Add(1)
}

func main() {
	inc := Incrementor{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			inc.Increase()
			wg.Done()
		}()

	}

	wg.Wait()
}
