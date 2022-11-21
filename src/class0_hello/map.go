package main

import "fmt"

func exam() {
	// 声明变量，默认 map 是 nil
	var book map[string]int
	book = map[string]int{} //缺一不可,因为map没有没有初始化,找不到地址

	book1 := make(map[string]int)
	book2 := map[string]int{} //	这两开箱即用,不需要初始化
	fmt.Println(book1["1"])
	fmt.Println(book2["2"])

	book["1元"] = 1
	book["2元"] = 2
	book["3元"] = 3
	book["4元"] = 4

	fmt.Println(book["1元"]) //assignment to entry in nil map,map赋值前要先初始化

	//用ok确认元素是否存在
	price, ok := book["3元"]
	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("这本书不存在")
	}
	fmt.Println(book)
	delete(book, "1元") //删除key
	fmt.Println(book)

	for k, v := range book {
		fmt.Println(k, ":", v)
	}
}

func main() {
	exam()
}
