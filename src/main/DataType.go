package main

import (
	"fmt"
)

// 默认是false
var result bool = true
var f32 float32
var f64 float64
var name string = "abcd"
var x int = 1
var y int = x

// 简短声明方式，左侧的变量不能是声明过的
var c = 10

//全局变量可以声明但不适用，局部变量不可以
func main() {
	fmt.Println(name, c, result, f32, f64)
	fmt.Println(x, y)
	x = 2
	fmt.Println(x, y)
	fmt.Println(&x)
	var r1 *int = &x
	var r2 *int = r1
	x = 10000
	r1 = &x
	fmt.Println(r1, r2)
}
