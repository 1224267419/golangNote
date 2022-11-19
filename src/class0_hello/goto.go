package main

import "fmt"

func main() {
	a := 0

	fmt.Println("我在A")
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 7 {
			goto B
		}
	}
B: //跳至标签处

	fmt.Println("我在B")

}
