package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	//res.Header()      //得到头部标签
	//res.Write([]byte("我收到了给你返回")) //向管道写入,直接去http://localhost:8080/test看就行
	//可以通过调试查看 Header类,里面有很多http的信息
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了给你返回GET"))
	case "POST":
		//res.Write([]byte("我收到了给你返回POST"))//写入了会提前产生状态码.具体看源代码就知道了
		b, _ := io.ReadAll(req.Body) //通过req.Body 请求body内容
		//header := res.Header()
		fmt.Println(req.Header["Test"])
		//header["test"] = []string{"test1", "test2"} //通过header设置头map
		res.WriteHeader(http.StatusLengthRequired) //设置返回状态码
		//json.Unmarshal()//通过这个吧b变成json
		res.Write(b)
		break
	}
}

func main() {
	//http.Handle()     //已经声明好的方法
	fmt.Println("打开 " + "http://localhost:8080/test")
	http.HandleFunc("/test", handler) //把handle创建进入默认路由器中
	http.ListenAndServe(":8080", nil) //启动服务

}
