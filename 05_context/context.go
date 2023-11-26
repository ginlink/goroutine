package my_context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("canceled")
				return
			default:
				fmt.Println("still working")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	// 3秒的时候取消
	cancel()
	time.Sleep(time.Second * 1)
	fmt.Println("done")
}

func ContextChainedCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx1, _ := context.WithCancel(ctx)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("job1 canceled")
				wg.Done()
				return
			default:
				fmt.Println("job1 still working")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("job2 canceled")
				wg.Done()
				return
			default:
				fmt.Println("job2 still working")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	// 3秒后再取消
	cancel()

	wg.Wait()
	fmt.Println("done")
}
