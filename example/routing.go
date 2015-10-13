// main.go
package main

import (
	"fmt"
//	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 1; i++ {

		fmt.Println("-------",s,"-------")
//		runtime.Gosched()
		fmt.Println("~~~~~~~",s,"~~~~~~~")
		fmt.Println("+++++++",s,"+++++++")


	}
}

func main() {
	go say("hello1")
	go say("hello2")
	go say("hello3")
    fmt.Println("----- start exe------")
	say("hello0")
	var t =time.NewTimer(time.Millisecond*1)
	<-t.C
//	var input string
//	fmt.Scanln(&input)
//	fmt.Println("done")
}