package main

import "io/ioutil"
import "fmt"

func eval(a, b int, op string) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		println(result)
	}
}
func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename) //ioutil.ReadFile返回两个参数,要用两个变量接收
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
}
