package main
import "fmt"
//In Go, an array is a numbered sequence of elements of a specific length.
//在Go里面 数组是一个指定长度的相同类型的连续元素的序列
func main() {
	//Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
//这里创建一个5个int类型的数组，数组的元素类型和长度都是数组类型的一部分，默认会赋零值
	var a [5]int
	fmt.Println("emp:", a)
	//We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//The builtin len returns the length of an array.
//使用内置长度len返回数组元素个数
	fmt.Println("len:", len(a))
	//Use this syntax to declare and initialize an array in one line.
//使用这种语法定义以及初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
//数组类型是一维数组，但是可以组合多维数组结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
//Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
//You’ll see slices much more often than arrays in typical Go. We’ll look at slices next.