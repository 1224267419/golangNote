package main

import (
	"fmt"
	"sync"
	"time"
)

func syncClass() {
	l := &sync.Mutex{}
	//l := new(sync.Mutex)//两者是一致的
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
	for {
	} //阻塞程序,显示结果
	/*
	   互斥锁,一把锁把其他程序锁死,
	   直到解锁后才可以执行,所以跟顺序执行的效果一致
	*/
}

func lockFun1(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("-1s")
	time.Sleep(time.Second)
	lock.Unlock()
}
func lockFun(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("-1s")
	time.Sleep(time.Second)
	lock.Unlock()
}

func rwsyncClass() {
	l := &sync.RWMutex{} //读写锁
	go lockFun1(l)
	go rlockFun(l) //读锁只会阻塞其他锁,不阻塞同类锁,写锁同理
	go rlockFun(l)
	for {
	} //阻塞程序,显示结果
	/*
	   互斥锁,一把锁把其他程序锁死,
	   直到解锁后才可以执行,所以跟顺序执行的效果一致
	*/
}
func rlockFun(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("+1s")
	time.Sleep(time.Second)
	lock.Unlock()
}
func main() {
	//syncClass()
	fmt.Println("下一部分,读写锁")
	rwsyncClass()
}
