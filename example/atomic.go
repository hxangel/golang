//The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
package main
import "fmt"
import "time"
import "sync/atomic"
import "runtime"
func main() {
	//We’ll use an unsigned integer to represent our (always-positive) counter.
	var ops uint64 = 0
	//To simulate concurrent updates, we’ll start 50 goroutines that each increment the counter about once a millisecond.
	for i := 0; i < 100; i++ {
		go func() {
			for {
				//To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
				atomic.AddUint64(&ops, 1)
				//Allow other goroutines to proceed.
				runtime.Gosched()
			}
		}()
	}
	//Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)
	//In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value into opsFinal via LoadUint64. As above we need to give this function the memory address &ops from which to fetch the value.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
//Running the program shows that we executed about 40,000 operations.
/*
$ go run atomic-counters.go
ops: 40200
*/
//Next we’ll look at mutexes, another tool for managing state.
//Next example: Mutexes.