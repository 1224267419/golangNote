package main

import (
	"fmt"
	"program_thinking/urlretriever"
	"program_thinking/urlretriever/testing"
)

/*
func retrieve(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body) //暂时不用err参数,以后再说错误处理
	return string(bytes)
}
*/

// 说明retriever是符合以下条件的所有类型
type retriever interface {
	Get(string) string
}

func getRetrievert() retriever {
	return testing.Retriever{} //此处仍使用testing.Retriever{},因为testing.Retriever符合retriever的定义
}
func getRetriever() retriever {
	return urlretriever.Retriever{}
}

func main() {
	//fmt.Printf("%s\n", retrieve("https://www.imooc.com"))
	//retriever := urlretriever.Retriever{}               //创建对象
	//retriever := getRetriever()                         //用函数来代替直接使用结构体,体现了main函数和Retriever结构体的解耦
	//var retriever urlretriever.Retriever = getRetriever() //体现了表明类型,好看(虽然对编译器来说和上一个一样
	var retrieve retriever = getRetriever()      //真实数据
	var retrieveTest retriever = getRetrievert() //测试数据

	fmt.Println(retrieve.Get("https://www.imooc.com"))     //真实数据
	fmt.Println(retrieveTest.Get("https://www.imooc.com")) //测试数据

}
