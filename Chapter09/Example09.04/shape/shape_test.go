package shape

import "testing"

func TestGetName(t *testing.T) {
	triangle := Triangle{Base: 15.5, Height: 20.1}
	rectangle := Rectangle{Length: 20, Width: 10}
	square := Square{Side: 10}

	if name := GetName(triangle); name != "三角形" {
		t.Errorf("%T 形狀錯誤: %v", triangle, name)
	}
	if name := GetName(rectangle); name != "長方形" {
		t.Errorf("%T 形狀錯誤: %v", rectangle, name)
	}
	if name := GetName(square); name != "正方形" {
		t.Errorf("%T 形狀錯誤: %v", square, name)
	}
}

func TestGetArea(t *testing.T) {
	triangle := Triangle{Base: 15.5, Height: 20.1}
	rectangle := Rectangle{Length: 20, Width: 10}
	square := Square{Side: 10}

	if value := GetArea(triangle); value != 155.775 {
		t.Errorf("%T 面積錯誤: %v", triangle, value)
	}
	if value := GetArea(rectangle); value != 200 {
		t.Errorf("%T 面積錯誤: %v", rectangle, value)
	}
	if value := GetArea(square); value != 100 {
		t.Errorf("%T 面積錯誤: %v", square, value)
	}
}
