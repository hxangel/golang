package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	t := time.Tick(time.Second*3)

	go func() {
		for {
			select {
			case <-t:
				watching()
			}
		}
	}()

	for i := 0; i < 100000000; i++ {
		c <- true
		go worker1(i, c)
	}

	fmt.Println("Done")
}

func watching() {
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
}

func worker1(i int, c chan bool) {
	//fmt.Println("worker", i)
	time.Sleep(100 * time.Microsecond)
	<-c
}