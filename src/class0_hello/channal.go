package main

import "fmt"

func main() {
	c1 := make(chan int, 2) //设置通道缓冲区(可以先进去,再等人来取
	c2 := make(chan int)    //无缓冲区
	c1 <- 11                //向缓冲区存入值
	c1 <- 1
	go func() {
		c2 <- 1 //等在c2处直至别的协程需要c2
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c1) //如果只存了一个,上面全取出来,就啥都没有了;存了两个则按队列排序,先进先出

	fmt.Println(<-c2) //需要c2时就去找输入c2的协程(注意,main是主线程

	fmt.Println("下一例程")
	fmt.Println("协程的范例(建议调试看")

	c3 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {

			c3 <- i //等在c3处直至别的协程需要c3

		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c3) //建议调试看看流程,蛮有意思的
	}

	fmt.Println("下一例程")
	fmt.Println("仅可读可写管道")

	c4 := make(chan int, 4)
	var readc <-chan int = c4  //仅可读
	var writec chan<- int = c4 //仅可写
	writec <- 1
	fmt.Println(<-readc)

	writec <- 11
	fmt.Println(<-c4) //一样可以通过c4进行读和写

	fmt.Println("下一例程")
	fmt.Println("close")

	c5 := make(chan int, 4)
	c5 <- 1
	c5 <- 2
	c5 <- 3
	c5 <- 4
	fmt.Println(<-c5)
	//c5 <- 1//c5关了,不能再写
	close(c5)           //关闭管道只能停止写入,但是仍能读取
	for v := range c5 { //使用for输出管道所有剩余值时,必须先close(),否则会报错,可以试试
		fmt.Println(v)
	}
}
