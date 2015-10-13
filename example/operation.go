//Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
package main
import (
	"fmt"
//	"time"
)
func main() {
	messages := make(chan string ,5)
	signals := make(chan bool)
	//Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}


	//A non-blocking send works similarly.
	msg := "hi"
//	time.Sleep(time.Second * 1)
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	//We can use multiple cases above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	messages <- msg
}
/*
$ go run non-blocking-channel-operations.go
no message received
no message sent
no activity
*/
//Next example: Closing Channels.

