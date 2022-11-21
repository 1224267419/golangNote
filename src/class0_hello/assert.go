package main

import "fmt"

type User struct {
	name string
	age  int
	sex  bool
}

type Student struct {
	User
}

func (u User) SayName(name string) {
	fmt.Println("我叫", name)
}

func check(v interface{}) {
	v.(User).SayName(v.(User).name) //断言后能调用结构体元素和方法
	//v.SayName(v.(User).age)//v没有SayName的方法,但是User有

	switch v.(type) { //获取接口类型(直接赋值也可),只能在switch中用
	case User:
		fmt.Println("我是user")
	case Student:
		fmt.Println("我是student")
	}
	//fmt.Println(u.(type))

}

func main() {
	u := User{
		name: "学生",
		age:  18,
		sex:  true,
	}
	check(u)
}
