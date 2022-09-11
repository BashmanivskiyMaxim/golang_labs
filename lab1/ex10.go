package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"

	symb := 'Ї'
	fmt.Printf("Code '%c' - %d\n", symb, symb)

	//rune - тип даних, який використовує кодові точки стандарту Unicode
	//для представлення елемента чергами байтів значенням.
}
