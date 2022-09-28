package task3

import "errors"

func main() {
	Add(1, 2, 3)
	FindMin(1, 2, 4)
	FindAverage(1, 2, 3)
	FindX(3, -15)
}

func Add(x1 float64, x2 float64, x3 float64) float64 {
	total := x1 + x2 + x3
	return total
}

func FindMin(vars ...int) int {
	min := vars[0]

	for _, i := range vars { // _,i значення діапазону
		if min > i {
			min = i
		}
	}
	return min
}

func FindAverage(x1 float64, x2 float64, x3 float64) float64 {
	return (x1 + x2 + x3) / 3
}

//3x = 24
//0x= 0

func FindX(a float64, b float64) (float64, error) {
	switch {
	case a != 0:
		return -(b / a), errors.New("")
	case a == 0 && b != 0:
		return 0, errors.New("рівняння коренів не має")
	case a == 0 && b == 0:
		return 0, errors.New("рівняння має бескінечну к-сть коренів")
	default:
		return 0, errors.New("помилка")
	}
}
