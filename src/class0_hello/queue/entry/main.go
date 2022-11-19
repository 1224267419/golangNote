package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.Queue{1} //q push 和pop后已经不是原来的值了,但是传出的值是指针,
	//所以可以仍可以使用q来进行操作
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty()) //判断是否为空
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
