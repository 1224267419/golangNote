package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := new(http.Client)
	//request, err := http.NewRequest("GET", "http://localhost:8080/test", nil)                             //方法,路径,body	构造请求
	request1, err := http.NewRequest("POST", "http://localhost:8080/test", bytes.NewBuffer([]byte("{test:我是客户端}"))) //方法,路径,body	构造请求
	//得到的内容去看服务端就知道了
	request1.Header["test"] = []string{"tset1", "test233"}

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(request1) //用request取得请求后的结果
	if err != nil {
		fmt.Println(err)
	}
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b)) //打印客户端接收到的东西
}
