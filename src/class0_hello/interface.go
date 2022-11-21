package main

import "fmt"

type animal interface {
	Eat()
	Run()
}
type cat struct {
	Name string
	Sex  bool
}

func (c cat) Run() {
	fmt.Println(c.Name, "跑")
}
func (c cat) Eat() {
	fmt.Println(c.Name, "吃")
}

type dog struct {
	Name string
}

func (c dog) Run() {
	fmt.Println(c.Name, "跑")
}
func (c dog) Eat() {
	fmt.Println(c.Name, "吃")
} //猫,狗必须符合条件才能实现接口
func (c dog) La() {
	fmt.Println(c.Name, "拉")
} //dog多了不同与

// 空接口
func myfun(a interface{}) {
	fmt.Println(a)
}

func main() {
	var a animal

	c := cat{
		Name: "Tom",
		Sex:  false,
	}
	d := dog{
		Name: "旺财",
	}

	a = c   //a作为接口类型,能直接被符合条件的类型赋值
	a.Run() //可以通过接口调用结构体方法,但不能调用值

	a = d //也是可行的,但只能调用animal有的方法即可
	//a.Name="tom"//不能这样改,因为a只有属于c的方法的部分,并未含有属性的部分

	myfun("1")
	myfun(2)
	myfun(true) //可以接收多种类型

}
