package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	//Задание.
	//1. Создайте 2 переменные  разных типов. Выпоните арифметические операции. Результат вывести

	var a int8 = 120
	var b int32 = 456456456

	fmt.Printf("a + b = %-5d = %.16b\n", int32(a)+b, int32(a)+b)
	fmt.Printf("a - b = %-5d = %.16b\n", int32(a)-b, int32(a)-b)
	fmt.Printf("a * b = %-5d = %.16b\n", int64(a)*int64(b), int64(a)*int64(b))
	fmt.Printf("b / a = %f", float32(b)/float32(a))

}
