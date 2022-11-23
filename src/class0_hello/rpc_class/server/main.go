package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}
type Req struct {
	NumOne int
	NunTwo int
}
type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res) error { //第二个参数必须是指针
	res.Num = req.NunTwo + req.NumOne
	fmt.Println("模拟操作时间")
	time.Sleep(3 * time.Second)
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8888")
	if e != nil {
		fmt.Println("报错")
	}
	http.Serve(l, nil)
}
