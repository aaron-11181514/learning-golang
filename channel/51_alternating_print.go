package channel

import "fmt"

// 使用兩個 goroutine 交替打印數字和字母：
// 1, A, 2, B, 3, C...
// 要求使用 channel 進行協調
// 並實現優雅的終止機制
func alternatingPrint() {
	numChan := make(chan struct{})
	letterChan := make(chan struct{})
	doneChan := make(chan struct{})
	go func() {
		for i := 0; i <= 10-1; i++ {
			<-numChan
			fmt.Sprint(i)
			letterChan <- struct{}{}
		}
		doneChan <- struct{}{}
	}()

	go func() {
		for i := 'A'; i <= 'J'; i++ {
			<-letterChan
			fmt.Sprint(i)
			numChan <- struct{}{}
		}
		doneChan <- struct{}{}

	}()
	numChan <- struct{}{}

	<-doneChan
	<-doneChan
	// 不需要顯示的關閉通道，因為當 chan 沒有被寫入跟取出後會進入閒置階段，閒置一段時間，會被 golang GC 清除
	// 通常在當作傳遞信號使用的 chan 不需要主動關閉
	// close(numChan)
	// close(letterChan)

	// 也不一定需要顯示的關閉
	// 如果對已經關閉的 chan 再次關閉會 panic
	// 對已經關閉得 chan 寫入，會 panic
	close(doneChan)

}
