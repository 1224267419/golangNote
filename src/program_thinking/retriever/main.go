package main

import (
	"fmt"
	mock "program_thinking/retriever/mockretirver"
	real2 "program_thinking/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake imooc,com"} //拷贝,值传递
	fmt.Println(download(r))                     //同样是download函数,通过mock.Retriever定义变量实现了Get()函数的转变

	r = &real2.Retriever{} //指针传递,
	fmt.Println(download(r))

	fmt.Println(r) //输出 {空格(UserAgent string) 0s(TimeOut   time.Duration)}
	r = &real2.Retriever{ //改成指针接收者
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Println(r)
}
