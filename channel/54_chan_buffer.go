package channel

import "fmt"

// 問：這段程式碼會輸出什麼？為什麼？
func ChanBuffer() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("Value sent")
	default:
		fmt.Println("Would block")
	}

	select {
	case v := <-ch:
		fmt.Printf("Received: %d\n", v)
	default:
		fmt.Println("No value available")
	}

	// Would block
	// Received 1

	// 有緩衝跟沒有緩衝的差別
	// 如果沒有緩衝
	// ch := make(chan int) // 無緩衝 channel
	// ch <- 1 // 這裡會阻塞，程式無法繼續執行

	// 那要怎麼作才可以呢？
	// 使用 goroutine 來讓程式往下跑
	// go func() {
	// 	fmt.Println(<-ch) // 接收值 1，解除阻塞
	// }()
	// ch <- 1 // 傳送值 1，因為有接收者，傳送成功

	// 要在阻塞前，建立好一接，一收
	// 如果有緩衝的話，舉例來說有5個緩衝
	// 那前五個寫入的動作，就可以直接寫在程式碼當中，像是一般的寫法
	// 因為不會造成阻塞行為

	// buffer 1 就是長度 1 可以先塞入一個值
}
