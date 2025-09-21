package main

import (
	"fmt"
	"log"
	"sync"
)

type ConqArray struct {
	array [5]int
	mutex *sync.Mutex
	wg    *sync.WaitGroup
}

func (ca *ConqArray) setSquare(pos int) error {
	defer ca.wg.Done()

	ca.mutex.Lock()
	defer ca.mutex.Unlock()

	if pos < 0 || pos >= len(ca.array) {
		log.Println("pos bigger than array len or less than zero")
	}

	value := ca.array[pos]
	ca.array[pos] = value * value

	return nil
}

func main() {
	ca := ConqArray{mutex: &sync.Mutex{}, wg: &sync.WaitGroup{}}
	ca.array = [5]int{2, 4, 6, 8, 10}

	for i := range ca.array {
		ca.wg.Add(1)
		go ca.setSquare(i)
	}

	ca.wg.Wait()

	fmt.Println(ca.array)
}
