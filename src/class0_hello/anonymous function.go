package main

import (
	"fmt"
	"math"
)

func getSequence() func(int) (int, int) {
	i := 0
	return func(num int) (int, int) { //闭包函数+1
		i += 1
		return i, num //直接调用i,用匿名函数作为getSequence()函数的返回值
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence() //这里的参数给getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber(9)) //闭包的传参,这里参数给匿名函数
	fmt.Println(nextNumber(9))
	fmt.Println(nextNumber(9))

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1(9))
	fmt.Println(nextNumber1(9))
	//不同变量的运算过程不共通

	//例子2
	hypot := func(x, y float64) float64 { //匿名函数,直接由变量接收
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance)
}
