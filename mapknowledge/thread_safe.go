package mapknowledge

import (
	"fmt"
	"sync"
)

// 請解釋為什麼 map 在 Go 中是非線程安全的？如何實現線程安全的 map？
// map is designed for efficient key-value storage but lacks built-in mechanisms thread safe
// if want to make a map thread safe,

// 1. sync.Mutex and sync.RWMutex
// granular control over lock boundaries
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func ThreadSafe1() {
	var sm = SafeMap{}

	var wg sync.WaitGroup
	wg.Add(100)

	// concurrent writes
	for i := 0; i <= 100-1; i++ {
		go func() {
			sm.mu.Lock()
			defer sm.mu.Unlock()
			sm.m[i] = i * 100
		}()
	}

	wg.Wait()
}

// 2. sync.Map
// high performance scenarios
func TreadSafe2() {
	sm := sync.Map{}

	var wg sync.WaitGroup

	wg.Add(100)
	// concurrent writes
	for i := 0; i <= 100-1; i++ {
		go func() {

			sm.Store(i, i)
		}()
	}
	sm.Range(func(key, value any) bool {
		fmt.Printf("%s: %v\n", key, value)
		return true
	})
}
