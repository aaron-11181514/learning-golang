package syncmutex

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 以下程式碼，的輸出會照順序嗎？還有 i 的值會改動嗎？還是會印出同一個值呢
func Clousure() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	time.Sleep(time.Second)

	wg.Wait()
}

// 不會照順序，因為 goroutine 是沒有順序的
// 理論上會是同一個 i 的值，因為是共享外面的 i
// 不過我順序而已試了很多次，都是會 0~4 不過不同
// 最好的方式是將 i 當作參數傳入各自的 goroutine
