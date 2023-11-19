package is_prime

func initChan(intChan chan int, num int) {
	for i := 1; i <= num; i++ {
		intChan <- i
	}
	close(intChan)
}

func isPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		var flag = true
		for j := 2; j < num; j++ {
			if num%j == 0 {
				flag = false
				continue
			}
		}

		if flag {
			primeChan <- num
			// fmt.Printf("num %d is prime\n", i)
		}
	}

	exitChan <- true
}
