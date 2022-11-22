package main

//TCP服务端
import (
	"fmt"
	"net"
)

func handleConnection(conn *net.TCPConn) {

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //见客户端
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(conn.RemoteAddr().String() + string(buf[0:n]))
		str := "收到了:" + string(buf[0:n])
		conn.Write([]byte(str)) //向管道写入
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9090") //创建tcp地址(选择tcp类型,端口号
	if err != nil {
		fmt.Println(err)
		return
	}
	ln, err := net.ListenTCP("tcp", tcpAddr) //监听tcp
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.AcceptTCP() //循环接收访问
		if err != nil {
			fmt.Println(err)
			// handle error
		}
		go handleConnection(conn)
	}

}
