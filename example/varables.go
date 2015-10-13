package main
import "fmt"
func main() {
	//var declares 1 or more variables.
	//使用var声明一个（多个）变量
	var a string = "initial"
	fmt.Println(a)
	//You can declare multiple variables at once.
	//一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)
	//Go will infer the type of initialized variables.
	//Go 会推断变量初始化的类型
	var d = true
	fmt.Println(d)
	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	//未赋值变量将初始化一个零值,如int 零值为0
	var e int
	fmt.Println(e)
	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case.
	//定义变量简化格式，根据值类型得到变量的类型
	f := "short"
	fmt.Println(f)
}