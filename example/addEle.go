package main

import (
	"fmt"
)
type m map[string]string

func main() {
	var x m
	x.set("abc","1")
	x.set("def","x")
fmt.Println(x)
}


func (p *m) set(k string, v string) {
	if *p == nil {
		*p = m{k: v}
		return
	}

	(*p)[k] = v
}