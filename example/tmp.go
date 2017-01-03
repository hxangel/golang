package main

import (
	"time"
	"fmt"
	"math/rand"
)

var Checking chan bool

func main() {
	rand.Seed(time.Now().UnixNano())
	skip := rand.Intn(100);
	fmt.Println(skip);
	Chenx()
	return
	rand.Seed(time.Now().UnixNano())
	//skip := rand.Intn(10000);
	fmt.Println(skip);
	return
	Check()
	return
	t, _ := time.Parse("2006-01-02 15:04:05", "2016-11-30 22:20:07")
	if t.Add(time.Hour).Before(time.Now()) {
		fmt.Println(t);
	}
}
func Chenx() {
	a := 0
	for {
		a++

		if a > 10 {

			break
		}
		fmt.Println(a)

	}
}
func Chen() {
	ch := make(chan bool, 2)
	jobs := make(chan int, 80)
	//full := make(chan bool)
	done := 0
	go func() {
		for {
			j, more := <-jobs
			if more {
				go func() {
					//time.Sleep(time.Microsecond*1500)
					time.Sleep(time.Microsecond * 1000)

					fmt.Println(j)
					//if j % 10 == 0 {
					//	fmt.Println(j)
					//} else {
					//	//SpiderLoger.I("Check Faild [", j.Host, "]")
					//}
					//SpiderLoger.I("Checked [", done, "]")
					done++
					// 回收chan
					<-ch
				}()
				ch <- true
			} else {
				//time.Sleep(time.Second * 1)
				fmt.Println("FULL")
				//满号
				//full <- true
			}
		}

	}()

	for i := 0; i <= 25; i++ {
		jobs <- i
	}
	time.Sleep(time.Second * 10)
	<-Checking

	//close(jobs)


}

func Check() {

	// Query All
	//mc := ms.Bulk()

	fmt.Println("Start checking proxys")

	ch := make(chan bool, 2)
	jobs := make(chan int)
	done := make(chan bool)
	//var wg sync.WaitGroup
	//wg.Add(1)
	go func() {
		for {
			j, more := <-jobs
			if more {
				go func() {
					if j % 50 == 0 {
						fmt.Print("H")
					}
					//回收ch
					<-ch
				}()
				//占ch坑
				ch <- true
			} else {
				//for i := 0; i < 10; i++ {
				//	<-ch
				//}
				//time.Sleep(time.Second * 1)
				fmt.Println("End checking proxys count[", "HHHHH", "]")
				//没有坑可以用了
				done <- true
			}
		}
		//defer wg.Done()

	}()
	//最后一次检测时间在8小时以前记录
	for p := 0; p < 10; p++ {
		jobs <- p
	}
	close(jobs)
	<-done
	//wg.Wait()
	fmt.Println("End checking proxys count[", "EEEEEEEEE", "]")
	//	fmt.Println("sent all jobs")
	//We await the worker using the synchronization approach we saw earlier.


}
