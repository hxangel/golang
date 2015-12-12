package main
import (
	"fmt"
)
type key struct {
	X int
	Y int
}
type point struct {
	key
	Z int
}

type xxoo struct {
	X int
	Z int
}

func main() {
	m := make(map[key]int)
	p1 := point{key{1, 2}, 3}
	p2 := point{key{3, 4}, 1}
	p3 := point{key{1, 2}, 2}
	m[p1.key] = 120
	m[p2.key] = 340
	fmt.Println(p1.key)
	fmt.Println(m[p1.key])
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}