package main

import "fmt"

var aa = 3
var bb = 4 //包内全局变量

var (
	cc = 10
	dd = 5 //括号省略多个var
)

func variable() {
	var a int
	var b string

	fmt.Printf("%d  %q\n", a, b)
}

func variableInit() {
	var a, b int = 3, 4
	var c string = "abc"
	fmt.Print(a, b, c)
}

func variableInit1() {
	var a, b, c, d = 3, 4, true, "def"
	q, w, e, r := 4, 2, false, "dasd"
	s, f := 1, 2
	fmt.Print(
		a, b, c, d, q, w, e, r, s, f)
}

func main() {
	fmt.Println("hello world")
	variable() //变量定义
	variableInit()
	variableInit1() //一次定义多个变量的两种方法
}
