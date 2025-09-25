package main

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mp map[K]V
	mx *sync.RWMutex
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		mp: make(map[K]V),
		mx: &sync.RWMutex{},
	}
}

func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()

	value, exists := sm.mp[key]

	return value, exists
}

func (sm *SafeMap[K, V]) Set(key K, val V) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	sm.mp[key] = val
}

func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	delete(sm.mp, key)
}

func main() {
	sm := NewSafeMap[string, int]()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Get(fmt.Sprintf("key%d", i))
		}(i)
	}

	wg.Wait()
	fmt.Println("Тест завершен")
}
