package main
import "fmt"
import "time"
//Switch statements express conditionals across many branches.
//switch条件语句穿过很多分支条件
func main() {
	//Here’s a basic switch.
	//	基本switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	//在同一个case语句中多个表达式使用逗号（,）隔开，同样可以使用可选得default作为默认
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	//switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	//没有表达式的switch语句是一种替换if／else逻辑的表达方式，这里也表现出case表达式可以为非常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}