package comprehensiveexercise

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
综合练习：获取房产的数据
*/

type House struct {
	Id           int
	BasicInfo    []string // 基本信息
	PriceHistory []int    // 价格历史
}

type Response struct {
	data map[string]any
	err  error
}

var (
	KeyBasicInfo    = "BasicInfo"
	KeyPriceHistory = "PriceHistory"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (h *Manager) GetHouseInfo(id int) (*House, error) {
	resCh := make(chan Response, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go h.getBasicInfo(id, resCh, wg)
	go h.getPriceHistory(id, resCh, wg)
	wg.Wait()
	close(resCh)

	resMap := make(map[string]any, 3)
	for res := range resCh {
		if res.err != nil {
			return &House{}, res.err
		}

		for key, val := range res.data {
			resMap[key] = val
		}
	}

	house := &House{
		Id:           id,
		BasicInfo:    (resMap[KeyBasicInfo]).([]string),
		PriceHistory: (resMap[KeyPriceHistory]).([]int),
	}
	return house, nil
}

func (h *Manager) getBasicInfo(id int, resCh chan Response, wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	basicResCh := make(chan []string)
	go func(basicResCh chan []string) {
		time.Sleep(time.Second * 1) // 模拟请求1s
		fmt.Sprintln("get basic info success, id=$d", id)

		basicResCh <- []string{"1号恒大楼盘", "2单元", "3楼"}
	}(basicResCh)

	select {
	case <-ctx.Done():
		fmt.Sprintln("get basic info timeout")
		resCh <- Response{
			data: map[string]any{},
			err:  fmt.Errorf("get basic info timeout, id=%d", id),
		}
		wg.Done()
	case res := <-basicResCh:
		resCh <- Response{
			data: map[string]any{KeyBasicInfo: res},
			err:  nil,
		}
		wg.Done()
	}
}

func (h *Manager) getPriceHistory(id int, resCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2) // 模拟请求2s
	fmt.Sprintln("get price history success, id=$d", id)
	resCh <- Response{
		data: map[string]any{KeyPriceHistory: []int{1, 2, 3}},
		err:  nil,
	}
	wg.Done()
}
