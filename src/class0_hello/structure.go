package main

import "fmt"

type book struct {
	name    string
	athor   string
	price   int
	country //可以简写也可以把名字加上(结构体嵌套)
}
type country struct {
	where string
	area  int
}

func (b *book) fun(name string) (res string) { //结构体方法不在结构体里面,需要额外定义
	fmt.Println(b.name)
	return b.athor
}
func (country *country) fun1() (res string) {
	return country.where
}
func (country *country) fun() (res string) {
	return country.where
}

func main() {
	var book1 book
	var book2 book

	book3 := new(book) //new创建指针类型
	book3.name = "susan"
	//*book3.name="改一下"//报错,
	(*book3).name = "改一下" //.的优先级比&更高

	book1.name = "书1"
	book1.athor = "作者1"
	book1.price = 100
	book2.name = "书2"
	book2.athor = "作者2"
	book2.price = 30
	book2.where = "中国"
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book2.price)
	fmt.Println(book1.name)
	fmt.Println(book3.name)
	var struct_pointer *book //结构体指针
	struct_pointer = &book1
	fmt.Println(struct_pointer.name)

	var number []int
	fmt.Println((number == nil))
	fmt.Println(book2.area)          //直接调用元素的 元素
	fmt.Println(book2.fun("11"))     //直接用的book的方法
	fmt.Println(book2.country.fun()) //调用的country的方法
	fmt.Println(book2.fun1())        //调用的country的方法

	fmt.Println(book2.fun(book2.name)) //结构体方法,用输出接收返回值

}
