package main
import (
	"fmt"
)
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()


	for w := range ch {
		fmt.Println("fmt print", w)
		if w > 5 {
			//break // 在这里break循环也可以
			close(ch)
		}
	}
	fmt.Println("after range or close ch!")
}

