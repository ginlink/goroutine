package wait_group

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("job 1 down")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("job 2 down")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All job down")
}
