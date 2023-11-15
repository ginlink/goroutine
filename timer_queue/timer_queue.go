package timer_queue

import (
	"fmt"
	"sync"
	"time"
)

func TimerQueue() {
	var group sync.WaitGroup
	group.Add(1)

	go func() {
		ticker := time.NewTicker(time.Second * 1)

		defer ticker.Stop()
		defer group.Done()

		count := 0
		for {
			t := <-ticker.C
			fmt.Println("now: ", t.Format("2006-01-02 15:04:05"))
			count++
			if count >= 2 {
				break
			}
		}
	}()

	group.Wait()
	fmt.Println("main end")
}
