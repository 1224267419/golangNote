package main

import "fmt"

func useful1() {
	const (
		a = 1 << (2 * iota) //iota照常计数,
		b                   //1 << 2
		c                   //1 << 4
		d                   //1 << 6
	)

	println(a, b, c, d)

}
func useful2() {
	const (
		i = 1 << iota
		j = 3 << iota
		k //3 <<2
		l //3 <<3
	)

	println(i, j, k, l)

}
func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	useful1()
	useful2()

}
