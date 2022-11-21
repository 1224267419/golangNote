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
func checkmap() {
	//m := map[int]int{} //同时对标准map进行读写,使用读写锁以支持异步处理map
	m := &sync.Map{}
	/*
		go func() {
			for {
				m[1] = 1
			}
		}()
		go func() {
			for {
				fmt.Println(m[1])
			}
		}()
	*/
	go func() {
		for {
			m.Store(1, 1) //存
			//m.Delete(1)//删
			m.Store(2, 3)
			m.LoadOrStore(2, 4) //存或读
			//m.LoadAndDelete()//见名知义,读完就删
		}
	}()
	go func() {
		for {
			fmt.Println(m.Load(1))
			fmt.Println(m.Load(2))
			fmt.Println(m.LoadOrStore(3, 6))

		} //不再报错,因为异步map

		m.Range(func(key, value any) bool {
			fmt.Println(key, value)
			time.Sleep(1 * time.Second)
			return true //仅当此处为true时才继续运行,否则阻塞(本身跟range差不多,得到所有key和value
		})
	}()
}

func rlockFun(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("+1s")
	time.Sleep(time.Second)
	lock.Unlock()
}
func checkPool() {
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	p.Put(5)
	p.Put(77)
	p.Put(11)
	for i := 0; i < 10; i++ {
		fmt.Println(p.Get()) //池子随机取值,没顺序之分,取完值就的nil
		//fmt.Println(p.Get())

	}
}
func checkCodn() {
	co := sync.NewCond(&sync.Mutex{}) //函数只能接收锁
	/*
		co.L.Lock()                       //调用锁
		co.Wait() //锁与解锁之间可以等待
		co.L.Unlock() //调用解锁
		co.Signal()    //解除一个等待(Wait)
		co.Broadcast() //解除所有锁

	*/
	go func() {
		co.L.Lock()
		fmt.Println("lock1")
		co.Wait()
		fmt.Println("unlock1")

		co.L.Unlock()
	}()
	go func() {
		co.L.Lock()
		fmt.Println("lock2")
		co.Wait()
		fmt.Println("unlock2")
		co.L.Unlock()
	}()
	time.Sleep(1 * time.Second) //等待解锁
	//co.Broadcast()              //解锁全部
	co.Signal() //只能解锁一个,而且是随机解锁一个
	//关于flag的用法看Signal()的例程吧
	time.Sleep(10 * time.Second)
}
func main() {
	//syncClass()
	//fmt.Println("下一部分,读写锁")
	//rwsyncClass()
	//checkmap()
	//checkPool()
	checkCodn()
	time.Sleep(1000000000)
}
