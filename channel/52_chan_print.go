package channel

import "fmt"

func chanPrint() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 對於已經關閉且緩衝區已經空的 chan 取值，並不會產生 panic ，而是會取得該型別的零值
	// 像是 int 就會是 0 
	// 是對已經關閉的 chan 進行寫入，才會 panic
}
