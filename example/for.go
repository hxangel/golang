package main
import "fmt"
//    for循环
func main() {
	//	The most basic type, with a single condition.
	//	最基本类型，单条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//	A classic initial/condition/after for loop.
	//	一个典型的整数循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//	for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	//	没有条件的循环会一直重复直到遇到中断或者 封闭 函数返回
	for {
		fmt.Println("loop")
		break
	}
}