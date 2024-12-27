package syncmutex

import (
	"fmt"
	"sync"
)

// 使用 Mutex 實現一個簡單的計數器，支持並發安全的增加和讀取操作
type SafeCounterStruct struct {
	value int
	mu    sync.Mutex
}

func (s *SafeCounterStruct) Increase() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value++
}

func (s *SafeCounterStruct) GetValue() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

func SafeCounter() {
	sc := &SafeCounterStruct{}

	var wg sync.WaitGroup

	for i := 0; i <= 100-1; i++ {
		go func() {
			wg.Add(1)
			for i := 0; i <= 1000-1; i++ {
				sc.Increase()

			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", sc.GetValue())
}

// 使用 mutex 鎖的效率比 chan 高，因為 chan 還需要有自己的邏輯需要處理
// 像是需要使用
