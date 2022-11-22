package main

//tcp客户端
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9090") //通过端口号,返回TCP端点地址
	fmt.Println("TCP端点地址:", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //可以理解为tcp传输的管道
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin) //通过bufio创建带缓存的输入
	for {
		bytes, _, _ := reader.ReadLine() //读取bufio.NewReader(os.Stdin)的一行(其实就是自己在shell写的)
		conn.Write(bytes)                //向管道写入

		buf := make([]byte, 1024) //创建[]byte
		n, err := conn.Read(buf)  //通过[]byte来接收管道的信息
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(conn.RemoteAddr().String() + string(buf[0:n])) //输出管道信息
	}
}
