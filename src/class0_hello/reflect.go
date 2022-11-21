package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Name string
	Age  int
	Sex  bool
}

type Student1 struct {
	Class int
	User1
}

func (s Student1) sing(name string) {
	fmt.Println("学生", s.Name, "会唱歌")
}

func (u User1) SayName(name string) {
	fmt.Println("我叫", name)
}

func check1(inter interface{}) {
	t := reflect.TypeOf(inter)          //获取接口类型的值
	v := reflect.ValueOf(inter)         //获取接口数据的值
	for i := 0; i < t.NumField(); i++ { //t.NumField()获取接口类型的数量
		fmt.Println(v.Field(i)) //v.Field(i)获取结构体第i个值
	}
	fmt.Println(t, v)
	fmt.Println(v.FieldByIndex([]int{1, 0})) //v.FieldByIndex按层级取值,参数必须是int型数组
	fmt.Println(v.FieldByName("Class"))      //v.FieldByIndex按名称取值,参数必须是值的名称
	ty := t.Kind()                           //通过 reflect.TypeOf(inter).Kind()得到数据类型
	//fmt.Println(ty)           //t就是结构体
	if ty == reflect.Struct { //reflect包带有多种类型,均可用于判断
		fmt.Println("我是结构体struct")
	}
}
func check2(inter interface{}) {
	v := reflect.ValueOf(inter) //获取接口数据的值
	fmt.Println("接下来改变原始数据")

	e := v.Elem() //类似于指针,e为结构体原始数据

	e.FieldByName("Age").SetInt(4) //先通过.FieldByName()找到数值,再通过.SetInt()修改数值
}
func check3(inter interface{}) {
	v := reflect.ValueOf(inter) //获取接口数据的值
	fmt.Println("接下来反射方法")

	m := v.Method(0) //调用第0个方法
	//m = v.MethodByName("sing")//名称调用方法
	m.Call([]reflect.Value{reflect.ValueOf("bin")})
	//
}

func main() {
	u := User1{
		Name: "学生",
		Age:  18,
		Sex:  true,
	}
	s := Student1{
		Class: 3,
		User1: u,
	}
	check1(s)
	check2(&s)
	fmt.Println("修改后S的Age是", s.Age)
	fmt.Println("修改后u的Age是", u.Age) //u的age没变
	check3(s)
}
