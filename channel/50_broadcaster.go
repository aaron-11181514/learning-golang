package channel

import (
	"fmt"
	"sync"
)

// 這是一個廣播機制
type Broadcaster struct {
	Subscribers      map[*Subscriber]struct{} // 用 map 在新增刪除的時候，效率更好
	AddSubscriber    chan *Subscriber
	DeleteSubscriber chan *Subscriber
	Broadcast        chan string // 這是為了異步處理，發送 message 的程式，只需要將 string 放入 chan 就可以了
	Close            chan struct{}
	mu               sync.Mutex // 因為需要對 map 作操作，只需要宣告結構就好，不需要初始化
}
type Subscriber struct {
	ID      string
	Message chan string // 用一個 chan 來接 message，這樣其他程式可以專心傳送 message 進來就好
}

func (b *Broadcaster) addSubscriber(Subscriber *Subscriber) {
	b.AddSubscriber <- Subscriber
}
func (b *Broadcaster) deleteSubscriber(Subscriber *Subscriber) {
	b.DeleteSubscriber <- Subscriber
}
func (b *Broadcaster) broadcast(s string) {
	b.Broadcast <- s
}
func (b *Broadcaster) close() {
	b.Close <- struct{}{}
}
func (b *Broadcaster) Run() {

	for { // 需要用 for 包住是因為，不包的話，變成只執行一次 select 的多選一個結果後，就不執行了
		// 用 for 包起來就會當任一個 chan 收到訊號的時候就會作某些事情
		select {
		case subscriber := <-b.AddSubscriber:
			b.mu.Lock()
			b.Subscribers[subscriber] = struct{}{}
			b.mu.Unlock()
		case subscriber := <-b.DeleteSubscriber:

			b.mu.Lock()
			delete(b.Subscribers, subscriber)
			close(subscriber.Message) // 也要關掉這個 subscriber 的 chan
			b.mu.Unlock()

		case message := <-b.Broadcast:
			b.mu.Lock()
			// 這邊為什麼要使用副本是因為把 map 鎖上，然後快速建立副本，快速釋放鎖，這樣可以把 map 的操作權讓出去，不用因為要分發消息，而讓 map 卡住
			// 為什麼副本要用 slice ，而不是也是使用 map ，因為這個副本的用途是拿來遍歷的，而創造 slice 的消耗是比較小的
			// map 底層因為還有 hash 表，所以消耗比較大
			// 遍歷的速度上 slice 也是高於 map ，因為 slice 是按索引訪問元素， map 是遍歷 hash 處理 key-value pair
			// 當然在其他情況， map 可能就會有其優勢了 1. 需要快速查找特定訂閱者 2. 需要避免重複 3. 需要存儲附加數據
			copySubscriber := make([]*Subscriber, 0, len(b.Subscribers)) // 長度 0 但是容量給定大小，這樣就不用擴容
			for subscriber := range b.Subscribers {
				copySubscriber = append(copySubscriber, subscriber)
			}
			b.mu.Unlock()

			for _, subscribe := range copySubscriber {
				select {
				// 如果 subscribe 的 message 通道可以塞入 message 那就塞入
				case subscribe.Message <- message:

					// 如果不行的話，就丟失這個訊息
				default:
					fmt.Sprintf("this message loss")
					// 因為與其廣播系統因為這一則訊息而停擺，不如接受丟失這則訊息的損失
				}
			}
		case <-b.Close:
			b.mu.Lock()
			for subscribe := range b.Subscribers {
				close(subscribe.Message)
			}
			b.Subscribers = nil //
			b.mu.Unlock()
			close(b.AddSubscriber)
			close(b.DeleteSubscriber)
			close(b.Broadcast)
		}
	}

}
