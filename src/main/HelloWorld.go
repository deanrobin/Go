package main

import "fmt"

/**

*/

// 全局变量的声明和赋值
var word = "hello"

// 常量定义
const a = 1

// 一般类型声明
type newType int

// 结构的声明
type gopher struct {

}

// 接口的声明
type golang interface {

}

// 函数名首字母小写为private 首字母大写为public
func main() {
	var word = "hello world!"

	fmt.Println(word, a)
}