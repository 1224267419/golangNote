package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 读取 key 和 value
	for key, value := range pow {
		fmt.Println(key, value)
	}

	// 读取 key
	for key := range pow {
		fmt.Println(key)
	}

	// 读取 value
	for _, value := range pow {
		fmt.Println(value)
	}
}
