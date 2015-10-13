package main
import "fmt"
//基本数据类型
func main() {
//	字符串类型 通过‘＋’连接
	fmt.Println("go" + "lang")
//	整型和浮点类型
	fmt.Println("1+1 =", 1 + 1)
	fmt.Println("7.0/3.0 =", 7.0 / 3.0)
//	布尔类型与布尔操作
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(0==0.00)
}