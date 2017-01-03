//Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it.
package main

import "time"
import "fmt"

func main() {

	Daemon()
	return
	//Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.
	ticker := time.NewTicker(time.Millisecond * 400)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func Daemon() {
	//一个小时抓取一次
	tick_get := time.NewTicker(time.Second * 8)
	//两个小时候检查一次
	tick_check := time.NewTicker(11 * time.Second)

	go func() {
		for {
			select {
			case <-tick_get.C:
				fmt.Println("get")
			case <-tick_check.C:
				fmt.Println("check")
			}
		}

	}()
	fmt.Println("start")
	time.Sleep(time.Hour)
}
//When we run this program the ticker should tick 3 times before we stop it.
/*
$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
*/
//Next example: Worker Pools.