package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	start := time.Now() //获取当前时间
	waitGroup.Add(5)    //waitGroup计数+5,负数就是-
	for i := 0; i < 5; i++ {
		go func() { //匿名函数
			defer waitGroup.Done() //waitGroup计数-1
			time.Sleep(time.Second)
			fmt.Println("done")
		}() //这里5个线程同步运行,因此只用了1s,删掉go就要用5s
	}

	waitGroup.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
}

/*
done
done
done
done
done
1.000306089
*/
