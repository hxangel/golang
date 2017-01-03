package main

import (
	"fmt"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()

}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

func TimerS()  {
	MyColor("red")
}
func timer1() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}

func timer2() {
	timer2 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer2.C:
			testTimer2()
		}
	}
}
func timer(){
	timer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer.C:
			for i := 1; i < 6; i++ {
				go fmt.Println(fmt.Sprintf("http://www.kuaidaili.com/proxylist/%d", i))
			}
		}
	}
}
func main() {
//	go timer1()
//	timer2()
//	timer()
	TimerS()
}
