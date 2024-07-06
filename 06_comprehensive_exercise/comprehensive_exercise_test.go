package comprehensiveexercise

import (
	"reflect"
	"testing"
)

// 测试 GetHouseInfo 方法
func TestGetHouseInfo(t *testing.T) {
	manager := NewManager()

	// 调用 GetHouseInfo
	house, err := manager.GetHouseInfo(1)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// 预期结果
	expectedBasicInfo := []string{"1号恒大楼盘", "2单元", "3楼"}
	expectedPriceHistory := []int{1, 2, 3}

	// 验证 BasicInfo
	if !reflect.DeepEqual(house.BasicInfo, expectedBasicInfo) {
		t.Errorf("expected BasicInfo %v, but got %v", expectedBasicInfo, house.BasicInfo)
	}

	// 验证 PriceHistory
	if !reflect.DeepEqual(house.PriceHistory, expectedPriceHistory) {
		t.Errorf("expected PriceHistory %v, but got %v", expectedPriceHistory, house.PriceHistory)
	}
}
