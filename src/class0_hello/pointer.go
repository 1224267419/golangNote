package main

import "fmt"

func swap(a, b int) (int, int) {
	c := a
	a = b
	b = c
	fmt.Println(a, b)
	return a, b
}
func swap1(a, b *int) { //引用传递
	*a, *b = *b, *a
}
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	var a int = 2
	var pa *int = &a //指针
	*pa = 3          //指针解引用
	fmt.Println(a)
	var c, d int = 3, 4
	swap(c, d)
	fmt.Println(c, d)
	swap1(&c, &d) //指针传参
	fmt.Println(c, d)
	c, d = swap(c, d) //值互换
	fmt.Println(c, d)
}
