/*
 * 利用信道做定时器
 */

package main

import (
	"fmt"
	"strconv"
)







func main() {
	go func() { fmt.Println("xxxxx") }()

	x := fmt.Sprintf("%0" + strconv.Itoa(2) + "x", 104 + 88)
	fmt.Println(x)

}