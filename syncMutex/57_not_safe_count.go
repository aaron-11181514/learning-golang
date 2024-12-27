package syncmutex

import (
	"fmt"
	"time"
)

// 有一個全局變量 counter int，兩個 goroutine 分別執行 1000 次加一操作。請問以下程式碼的輸出結果是否一定是 2000？為什麼？

func NotSafeCount() {
	var counter int

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(counter)

	// 不一定，因為對於 counter 的操作，並不是併發安全的
	// 實際上的操作分了三個步驟，1. 複製取值 2. ++ 3. 賦值回去
	// 可能同時複製到 100 這個值，一起對他+1，這樣對於 global 來說就是少 1
}
