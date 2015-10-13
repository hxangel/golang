package main
import "fmt"
import "math"
//const declares a constant value.
//常量声明与赋值
const strings string = "constant"
//常量操作
func main() {
	fmt.Println(strings)
	//A const statement can appear anywhere a var statement can.
	//常量可以任何任何声明变量的地方声明
	const n = 500000000
	//Constant expressions perform arithmetic with arbitrary precision.
	//常量表达式可执行任意精度的算术运算
	const d = 3e20 / n
	fmt.Println(d)
	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	//一个数字类型的常量没有类型，直到分配一个数据类型，如显示式给予
	fmt.Println(int64(d))
	//A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	//一个数字在结构中需要的情况下会自动分配一个类型，如变量分配或者函数调用如Tan传递一个float64
	fmt.Println(math.Tan(d))
}
