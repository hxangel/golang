//A goroutine is a lightweight thread of execution.
package main
import "fmt"
func f(from string) {
	for i := 1; i <= 100; i++ {
		if i%10==0{
			fmt.Println();
		}
		fmt.Print(from, ":", i)
	}
}
func main() {
	//Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
	f("d")
	//To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("e")
	//You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	//Our two function calls are running asynchronously in separate goroutines now, so execution falls through to here. This Scanln code requires we press a key before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
//When we run this program, we see the output of the blocking call first, then the interleaved output of the two gouroutines. This interleaving reflects the goroutines being run concurrently by the Go runtime.
/*
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done
*/
//Next we’ll look at a complement to goroutines in concurrent Go programs: channels.
//Next example: Channels.