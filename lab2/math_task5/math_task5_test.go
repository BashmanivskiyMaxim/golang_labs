package math_task5

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	x := Add(1, 2, -3)
	res := 0.0
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}

func TestMin(t *testing.T) {
	x := FindMin(1, 2, -3)
	res := -3
	if x != res {
		t.Errorf("Тест не пройдений! Результат %d, а повинен бути %d", x, res)
	}
}

func TestMin1(t *testing.T) {
	x := FindMin(124, 45, 43)
	res := 43
	if x != res {
		t.Errorf("Тест не пройдений! Результат %d, а повинен бути %d", x, res)
	}
}

func TestAverage(t *testing.T) {
	x := FindAverage(124, 45, 47)
	var res float64 = 72
	if x != res {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res)
	}
}

func TestFindX1(t *testing.T) {
	x, y := FindX(3, 15)
	var res1 float64 = -5
	var res2 error = errors.New("")
	if x != res1 && y != res2 {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res1)
	}
}

func TestFindX2(t *testing.T) {
	x, y := FindX(3, 6)
	var res1 float64 = -2
	var res2 error = errors.New("")
	if x != res1 && y != res2 {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f", x, res1)
	}
}

func TestFindX3(t *testing.T) {
	x, y := FindX(0, -7)
	var res1 float64 = 0
	var res2 error = errors.New("рівняння коренів не має")
	if x != res1 && y != res2 {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f, %s", x, res1, res2)
	}
}

func TestFindX4(t *testing.T) {
	x, y := FindX(0, 0)
	var res1 float64 = 0
	var res2 error = errors.New("рівняння має бескінечну к-сть коренів")
	if x != res1 && y != res2 {
		t.Errorf("Тест не пройдений! Результат %f, а повинен бути %f, %s", x, res1, res2)
	}
}
