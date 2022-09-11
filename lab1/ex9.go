package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false, не ініціалізована змінна bool типу буде мати значення false
	fmt.Println("second = ", second)      // false, не ініціалізована змінна bool типу буде мати значення false
	fmt.Println("third  = ", third)       // true, ініціалізована змінна bool типу має значення true
	fmt.Println("fourth = ", fourth)      // false, змінна записана у короткій формі, та ініціалізована як протилежне значення до змінної third
	fmt.Println("fifth  = ", fifth, "\n") // true, ініціалізована змінна bool типу має значення true

	fmt.Println("!true  = ", !true)        // false, логічний оператор ні,
	fmt.Println("!false = ", !false, "\n") // true, якщо приклад оцінюється як хибний, повертає істину

	fmt.Println("true && true   = ", true && true)         // true, опираючись на таблицю істини для(&&(і)) true && true = true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true, опираючись на таблицю істини для(||(або)) true && true = true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true, істина 3 більше 2
	fmt.Println("2 > 3  = ", 2 > 3)        // false, хиба 2 менше 3
	fmt.Println("3 < 3  = ", 3 < 3)        // false, хиба 3 не менше 3
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true, істина 3 менше рівне 3
	fmt.Println("3 > 3  = ", 3 > 3)        // false, хиба 3 не більше 3
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true, істина 3 більше рівне 3
	fmt.Println("2 == 3 = ", 2 == 3)       // false, хиба 2 не рівне 3
	fmt.Println("3 == 3 = ", 3 == 3)       // true, істина 3 рівне 3
	fmt.Println("2 != 3 = ", 2 != 3)       // true, істина 2 не дорівнює 3
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false, хиба, 3 дорівнює 3

	//Задание.
	//1. Пояснить результаты операций
}
