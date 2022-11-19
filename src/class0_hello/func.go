package main

import (
	"fmt"
	"strconv"
)

func bino(num int) (result string) { //返回二进制值

	for ; num > 0; num /= 2 {
		lsb := num % 2
		result = strconv.Itoa(lsb) + result

	}
	return //result在开始时已经定义好了,结束时也会自动返回
}

func apply(op func(int) string, a int) string { //函数可以嵌套函数,名字在前,类型在后
	return op(a)
}

func sum(values ...int) { //传入不定长的int数组
	sum := 0
	for i := range values {
		sum += values[i]

	}
	fmt.Println(sum)
}

func main() {
	apply(bino, 5)
	sum(1, 2, 3, 4, 5)

	//匿名函数例子
	b := func(int) int {
		return 123123
	}
	fmt.Println(b(1))
}
