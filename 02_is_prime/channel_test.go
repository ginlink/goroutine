package is_prime

import (
	"fmt"
	"testing"
)

func TestIsPrime(m *testing.T) {
	num := 100
	childNum := 10
	intChan := make(chan int, num)
	primeChan := make(chan int, num)
	exitChan := make(chan bool, childNum)

	go initChan(intChan, num)

	for i := 0; i <= childNum; i++ {
		go isPrime(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i <= childNum; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		fmt.Println("素数：", res)
	}
}
