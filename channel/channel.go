package channel

import (
	"fmt"
)

/*
用goroutine和channel协同工作完成
- 开启一个writeData协程，向管道intChan中写入50个整数
- 开启一个readData协程，从管道intChan中读取writeData写入的数据
- 注意：读写操作的是同一个管道的数据。主线程需要等待读写的协程完成才能退出。
*/

func goroutineWithChannel() {
	intChan := make(chan int, 50)
	existChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, existChan)

	for {
		if <-existChan {
			break
		}
	}
	fmt.Println("end main")
}

func writeData(intChan chan int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("写入: %d\n", i)
		// time.Sleep(time.Second)
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取: %d\n", num)
	}
	exitChan <- true
	close(exitChan)
}
