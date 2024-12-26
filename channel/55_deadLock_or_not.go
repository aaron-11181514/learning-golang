package channel

import "fmt"

// 問：這段程式碼會正常執行還是發生死鎖？為什麼？
func DeadLockOrNot() {
	ch := make(chan int)
	done := make(chan bool)

	for i := 0; i < 3; i++ {
		go func() {
			<-ch
			done <- true
		}()
	}

	close(ch)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)

	// 會發生死鎖
	// go routine 可以想像成一個方塊放在旁邊，可能先執行也可能後執行
	// close ch 不會有問題
	// 從關閉的 chan 取東西，也不會有問題，如果東西可以取了，會給零值
	// 但是取出零值後，馬上會寫 true 到 done 通道裡面，因為 done 通道是無緩衝通道，所以會阻塞
	// 有可能這個時候執行到 18 行然後有沒有阻塞，然後又執行 goroutine ，不過也可能沒有那麼剛好
	// 所以大機率會永久阻塞，變為死鎖
}
