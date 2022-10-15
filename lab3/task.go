package main

import (
	"fmt"
	"math"
)

func main() {
	task()
}

const N int64 = 30000

func task() {
	var x0 uint64
	var a uint64 = 16807
	var c uint64 = 0
	var k int = 30000

	var arr [N]uint64
	var arr1 [250]int64

	var x float64 = 1
	var m = uint64(math.Pow(2, 31) - 1)

	fmt.Print("Введіть початкове значення x: ")
	fmt.Scanf("%d", &x0)

	for i := 0; i < k; i++ {
		x0 = (x0*a + c) % m
		x = float64(x0) / 2147483647 * 200
		arr[i] = uint64(x)
		fmt.Printf(" %d", arr[i])
	}
	fmt.Print("\n")
	fmt.Print("Частота інтервалів появи випадкових величин:\n")
	var f int = 0
	var g int = 0
	var arr2 [200]float64

	for h := 0; h < 200; h++ {
		for j := 0; j < k; j++ {
			if arr[j] == uint64(h) {
				f++
			}
		}
		fmt.Printf("  (%d) | %d\n", h, f)
		arr1[h] = int64(f)
		f = 0
	}
	fmt.Printf("\n")
	fmt.Print("Статистична ймовірність появи випадкових величин:\n")
	for h := 0; h < 200; h++ {
		for j := 0; j < k; j++ {
			if arr[j] == uint64(h) {
				g++
			}
		}
		fmt.Printf("  (%d) | %f\n", h, float64(g)/30000)
		arr2[h] = float64(g) / 30000
		g = 0
	}
	var Expvalue float64 = 0
	for j := 0; j < 200; j++ {
		Expvalue = float64(arr[j])*arr2[j] + Expvalue
	}
	fmt.Printf("\n")
	fmt.Printf("Математичне сподівання появи випадкових величин: %f \n", Expvalue)

	var variance float64 = 0
	for j := 0; j < k; j++ {
		variance = math.Pow(float64(arr[j])-Expvalue, 2)*float64(arr[j]) + variance
	}
	fmt.Printf("Дисперсія випадкових велечин: %f \n", variance)

	var deviation float64 = math.Sqrt(variance)
	fmt.Printf("Середньоквадратичне відхилення випадкових величин: %f \n", deviation)

}
