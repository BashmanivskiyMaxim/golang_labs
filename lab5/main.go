package main

import (
	"encoding/json"
	"fmt"
	"lab5/currency"
	"lab5/product"
	"os"
	"sort"
	"strconv"
)

func GetProductsInfo(prod []product.Product) string {
	sort.SliceStable(prod, func(i, j int) bool {
		return prod[i].GetPriceIn() < prod[j].GetPriceIn()
	})
	out1, _ := json.Marshal(prod[0])
	out2, _ := json.Marshal(prod[len(prod)-1])
	return string(out1) + " - найдешевший товар \n" + string(out2) + " - найдорожчий товар"
}

func PrintProducts(prod []product.Product) {
	for j := 0; j < len(prod); j++ {
		fmt.Printf("[%d]", j)
		PrintProduct(prod[j])
	}
}

func PrintProduct(prod product.Product) {
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("Назва продукту: %s\n", prod.Name)
	fmt.Printf("Ціна продукту: %d\n", prod.Cost)
	fmt.Printf("К-сть продуктів на складі: %d\n", prod.Quantity)
	fmt.Printf("Назва компанії виробника продукту: %s\n", prod.Producer)
	fmt.Printf("Назва валюти: %s\n", prod.Money.Name)
	fmt.Printf("Курс валюти: %f\n", prod.Money.ExRate)
	fmt.Println("--------------------------------------------------")
}

func ReadProductsArray() []product.Product {
	var s string
	var i int

	fmt.Println("Введіть кількість товарів: ")

	_, err := fmt.Scan(&s)
	i, err = strconv.Atoi(s)

	if err != nil {
		fmt.Println("Введіть число!")
	} else {

		products := make([]product.Product, i)

		var name string
		var cost int
		var quantity int
		var producer string
		var weight int
		var nameVal string
		var exrate float64

		for j := 0; j < i; j++ {
			fmt.Print("Введіть назву продукту: ")
			_, err = fmt.Scan(&name)

			fmt.Print("Введіть ціну продукту: ")
			_, err = fmt.Scan(&cost)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			fmt.Print("Введіть к-сть продуктів на складі: ")
			_, err = fmt.Scan(&quantity)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			fmt.Print("Введіть назву компанії виробника продукту: ")
			_, err = fmt.Scan(&producer)

			fmt.Print("Введіть вагу продукту: ")
			_, err = fmt.Scan(&weight)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			fmt.Print("Введіть назву валюти: ")
			_, err = fmt.Scan(&nameVal)

			fmt.Print("Введіть курс валюти: ")
			_, err = fmt.Scan(&exrate)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			products[j] = product.Product{Name: name, Cost: cost, Quantity: quantity,
				Producer: producer, Weight: weight, Money: currency.Currency{Name: nameVal, ExRate: exrate}}
			fmt.Println("----------------------------------------------------------------")
		}
		return products
	}

	return []product.Product{}
}

func main() {

	//var curr1 = currency.NewCurrency()
	//curr1.SetName("Prikol")
	//fmt.Println(curr1)

	var prod2 = product.Product{Name: "asdf1", Cost: 65, Quantity: 78, Producer: "Qwerty", Weight: 78, Money: currency.Currency{Name: "UAH", ExRate: 78.67}}
	var prod3 = product.Product{Name: "asdf2", Cost: 35, Quantity: 78, Producer: "Qwerty", Weight: 78, Money: currency.Currency{Name: "UAH", ExRate: 78.67}}
	var prod4 = product.Product{Name: "asdf3", Cost: 15, Quantity: 78, Producer: "Qwerty", Weight: 78, Money: currency.Currency{Name: "UAH", ExRate: 78.67}}
	var prod5 = product.Product{Name: "asdf4", Cost: 95, Quantity: 78, Producer: "Qwerty", Weight: 78, Money: currency.Currency{Name: "UAH", ExRate: 78.67}}

	//products := [4]product.Product{prod2, prod3, prod4, prod5}

	var prod1 = product.NewProduct()
	prod1.SetProduct("Mouse", 1000, 245, "Razer", 2, "UAH", 40.5)
	fmt.Println(prod1.GetTotalPrice())
	fmt.Println(prod1.GetTotalWeight())
	fmt.Println(prod1.GetPriceIn())

	//ReadProductsArray()

	PrintProducts([]product.Product{prod2, prod3, prod4, prod5})
	PrintProduct(prod2)

	fmt.Println(GetProductsInfo([]product.Product{prod2, prod3, prod4, prod5}))
}
