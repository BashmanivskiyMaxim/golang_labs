package Server

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	pageHeader = `<!DOCTYPE html>
	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <title>Title</title>
	</head>`

	pageBodyTask1 = `<body>
		<h1>Завдання 1</h1>`

	pageBodyTask2 = `<body>
		<h1>Завдання 2</h1>`

	form2 = `<form action="/task2" method="POST">
        <label>
            Введіть значення через кому:
            <input type="text" name="numbers">
        </label>
        <input type="submit">

    </form>`

	form1 = `<form action="/task1" method="POST">
        <label>
            A
            <input type="text" name="numbersA">
        </label>
        <label>
            B
            <input type="text" name="numbersB">
        </label>
        <label>
            C
            <input type="text" name="numbersC">
        </label>
        
        <input type="submit">

    </form>`
	pageFooter = `</body></html`

	result = `<p>Результат: </p>`
)

type Solution string

var HttpSolution Solution = "Завдання 1"

func FindSol(a string, b string, c string) string {
	aParse, _ := strconv.ParseFloat(a, 32)
	bParse, _ := strconv.ParseFloat(b, 32)
	cParse, _ := strconv.ParseFloat(c, 32)

	disc := (bParse * bParse) - (4 * aParse * cParse)

	if disc > 0 {
		res1 := (-bParse + math.Sqrt(disc)) / (2 * aParse)
		res2 := (-bParse - math.Sqrt(disc)) / (2 * aParse)
		return fmt.Sprintf(`<p>Корінь 1 = %s Корінь 2 = %s</p`, fmt.Sprintf("%.2f", res1), fmt.Sprintf("%.2f", res2))
	} else {
		if disc < 0 {
			return fmt.Sprintf(`<p>Коренів немає!'</p`)

		} else {
			if disc == 0 {
				res3 := -bParse / (2 * aParse)
				return fmt.Sprintf(`<p>Корінь = %s </p`, fmt.Sprintf("%.2f", res3))
			}
		}
	}
	return "Error"
}

func CalculateArr(val []string) string {
	valSlice := make([]float64, len(val))
	var v string
	var i int
	for i, v = range val {
		valSlice[i], _ = strconv.ParseFloat(v, 64)
	}

	var sumNegative float64 = 0
	var product float64 = 1

	for j := 0; j < len(valSlice); j++ {
		if valSlice[j] < 0 {
			sumNegative = sumNegative + valSlice[j]
		}
		product = product * valSlice[j]
	}
	return fmt.Sprintf(`<p>Зріз: %s Сума від'ємних елементів = %s Добуток елементів = %s</p`,
		fmt.Sprintf("%.2f", valSlice), fmt.Sprintf("%.2f", sumNegative), fmt.Sprintf("%.2f", product))

}

func StartServ() {

	http.HandleFunc("/task2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, pageHeader, pageBodyTask2, form2)
		if r.Method == "GET" {
			arr := r.URL.Query()
			values := arr["arr"]
			CalculateArr(values)
			fmt.Fprintf(w, CalculateArr(values))

		}
		if r.Method == "POST" {
			err := r.ParseForm()
			post := r.PostForm
			if err != nil {
				fmt.Fprintf(w, result, err)
				return
			}
			pnumbers := post.Get("numbers")
			fmt.Println(pnumbers)

			numbers := strings.Split(pnumbers, ",")
			numbersSlice := numbers[:]
			fmt.Fprintf(w, CalculateArr(numbersSlice))

		}

		fmt.Fprint(w, "\n", pageFooter)
	})

	http.HandleFunc("/task1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, pageHeader, pageBodyTask1, form1)

		if r.Method == "GET" {
			a := r.URL.Query().Get("a")
			b := r.URL.Query().Get("b")
			c := r.URL.Query().Get("c")
			FindSol(a, b, c)
			fmt.Fprintf(w, FindSol(a, b, c))
		}

		if r.Method == "POST" {
			err := r.ParseForm()
			post := r.PostForm
			if err != nil {
				fmt.Fprintf(w, result, err)
				return
			}
			pnumbersA := post.Get("numbersA")
			pnumbersB := post.Get("numbersB")
			pnumbersC := post.Get("numbersC")

			fmt.Fprintf(w, FindSol(pnumbersA, pnumbersB, pnumbersC))

		}
		fmt.Fprint(w, "\n", pageFooter)
	})

	http.ListenAndServe("Localhost: 80", nil)
}
