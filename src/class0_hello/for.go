package main

import (
	"fmt"
)

func for_for() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func for_while() { //返回二进制值
	sum := 0
	i := 1
	for i <= 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func forever() {
	a := 1
	for {
		a++
		fmt.Println(a)
		if a > 1000 {
			break //go的break,同c里面的
		}
	}
}
func range_for() {

	strings := []string{"google", "runoob"}

	for i, s := range strings {

		fmt.Println(i, s)

	}
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	for _, value := range numbers { //用_来获取不需要的参数,避免报错
		fmt.Printf("value is: %d\n", value)
	}
}
func break_for() {
	a := 0

	fmt.Println("我在A")
A: //用A标注循环,再用break
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 7 {
			break A
		}
	}
	return
}

func main() {
	for_for()   //类比c的for
	for_while() //类比c的while
	range_for() //range的用法
	break_for() //标签标记循环与break
	//forever()   //死循环
}
