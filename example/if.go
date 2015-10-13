package main
import "fmt"
func main() {
	//	Here’s a basic example.
	//  简单实例
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//	You can have an if statement without an else.
	//  if语句可以不带else
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//  A statement can precede conditionals; any variables declared in this statement are available in all branches.
	//  在条件之前优先声明变量，任意在此变量在所有的分支语句都都可以使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
//Note that you don’t need parentheses around conditions in Go, but that the braces are required.
//注意：你不需要在条件上加括号，但是语句体上是一定要加上的
//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
//Go中没有三元运算表达式，所以得用if条件语句进行替换