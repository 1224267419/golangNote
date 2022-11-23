package main

import (
	"fmt"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NunTwo int
}
type Res struct {
	Num int
}

func main() {
	req := Req{NumOne: 1, NunTwo: 2} //传入数量不正确的参数容易导致panic,最好不要这样做
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888") //创建连接
	if err != nil {
		fmt.Println(err)
	}
	//client.Call("Server.Add", req, &res)    //同步操作
	ca := client.Go("Server.Add", req, &res, nil) //异步操作,done参数取nil会给10长度的缓冲区
	//fmt.Println("在服务端处理好之前我可以执行其他操作")
	//<-ca.Done //参考之前异步的部分

	for {
		select {
		case <-ca.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("我等你处理完")
		}
	}

	fmt.Println(res)
}
