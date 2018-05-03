package main

import (
	"fmt"
	"unsafe"
)

// 常量
func main() {
	const length int = 10
	const width int = 5
	var area int

	area = length * width
	fmt.Println("5 * 10 =", area)
	//println(length, width)

	const (
		a = 1
		b = iota
		c
		d = "emmmm.."
		e
		f = 100
		g
		h = iota
		i
	)

	const (
		j = iota
		k
	)
	// := 左边的值不用是赋值过得变量名，且不需要类型 系统会自动判断
	m := 100
	fmt.Println(a,b,c,d,e,f,g,h,i,j,k)
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(m)
}