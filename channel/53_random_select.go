package channel

import "fmt"

func RandomSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("ch1 received")
	case <-ch2:
		fmt.Println("ch2 received")
	}
	// 問：每次執行輸出是否相同？為什麼？
	// 不一定
	// 因為在 select 當中的順序是隨機的，如果 case1, case2 同時滿足，那會隨機挑一個執行，然後這個 select 就結束了
	// 如果要一直執行，要用 for 把 select 包起來

}
