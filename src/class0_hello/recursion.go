package main

import "fmt"

func recursion(i int) (result int) {
	if i > 0 {
		result = i * recursion(i-1)
		return result
	}
	return 1
}

func main() {
	fmt.Println(recursion(5))
}
