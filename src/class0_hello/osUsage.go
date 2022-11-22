package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func checkRead() {
	f, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_RDWR, 0777) //os.O_CREATE|os.O_RDWR:不存在文件就创建,给读写权限
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		b := make([]byte, 16) //调整byte的数量,从而得到更多字符(乱码是因为没用加载整个中文字符,3字节)
		n, err := f.Read(b)
		f.Seek(-3, 1) //每次读取后从当前位置向前移动3byte,但是一直向前移结束不了
		//fmt.Println(b)
		if err != nil {
			fmt.Println(err) //结束了返回EOF
			return
		}
		fmt.Println(string(b), n)
	}
}
func checkWrite() {
	f, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_RDWR, 0777) //os.O_CREATE|os.O_RDWR:不存在文件就创建,给读写权限
	if err != nil {
		fmt.Println(err)
		return
	}
	//for {
	//	b := make([]byte, 16) //调整byte的数量,从而得到更多字符(乱码是因为没用加载整个中文字符,3字节)
	//	f.Write([]byte("123456"))
	//	time.Sleep(1 * time.Second)
	//	//fmt.Println(b)
	//	if err != nil {
	//		fmt.Println(err) //结束了返回EOF
	//		return
	//	}
	//	fmt.Println(string(b))
	//}

	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(f)
	n := 0
	for {
		n++
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		writer.WriteString(strconv.Itoa(n) + " " + str) //读取完,加上n直接写入缓存
	}
	f.Seek(0, 0) //光标放回开始值

	writer.Flush() //缓存写入文件

}

func checkBufio() {
	f, err := os.OpenFile("./1.txt", os.O_CREATE|os.O_RDWR, 0777) //os.O_CREATE|os.O_RDWR:不存在文件就创建,给读写权限
	if err != nil {
		fmt.Println(err)
		return
	}

	////示例1
	//reader := bufio.NewReader(f)
	//for {//一直写入
	//	res, err := reader.ReadString('\n') //每次到\n就停止
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(res)
	//}

	//示例2
	//reader := bufio.NewReader(f)
	//for {
	//	res, isPrefix, err := reader.ReadLine() //每次读一行,读完找下一行
	//	if err != nil {
	//		fmt.Println(err, isPrefix)
	//		return
	//	}
	//	fmt.Println(string(res), isPrefix)
	//}

	//示例3

	res, err := ioutil.ReadAll(f) //每次到\n就停止
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))

	os.ReadFile("./1.txt") //文件打开并 只读

}

func main() {
	//io.Seeker.Seek()//光标位置,用于读和写
	//os.Open()
	//checkRead()
	checkWrite() //写入
	//checkBufio()
}
