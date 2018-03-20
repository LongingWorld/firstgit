package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b //常量声明省略值时，默认和之前一个值的字面相同 此时b=iota
		c = "iota"
		d
		e = 2018
		f
		g = iota
		h
	)
	fmt.Println(a, b, c, d, e, f, g, h)
}
