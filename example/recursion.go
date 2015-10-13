//Go by Example: Recursion
//Go supports recursive functions. Here’s a classic factorial example.
package main
import "fmt"
//This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}
func main() {
	fmt.Println(fact(7))
}
/*
$ go run recursion.go
5040
*/
//Next example: Pointers.