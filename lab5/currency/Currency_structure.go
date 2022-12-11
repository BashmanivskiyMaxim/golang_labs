package currency

import (
	"fmt"
)

type Currency struct {
	Name   string
	ExRate float64
}

// Конструктор

func NewCurrency() *Currency {
	const (
		defaultName   = "Name"
		defaultExRate = 1.1
	)

	return &Currency{
		Name:   defaultName,
		ExRate: defaultExRate,
	}
}

func (e *Currency) SetCurrency(Name string, ExRate float64) {
	e.SetName(Name)
	e.SetExRate(ExRate)
}

//Get - маетоди

func (e *Currency) GetName() string {
	return e.Name
}

func (e *Currency) GetExRate() float64 {
	return e.ExRate
}

//Set - методи

func (e *Currency) SetName(name string) {
	e.Name = name
}

func (e *Currency) SetExRate(exRate float64) {
	e.ExRate = exRate
}

func (e *Currency) GetNameIn() string {
	return "Ім'я - " + e.Name
}

func (e *Currency) GetInfo() string {
	return "Ім'я - " + e.Name + "\n" + "Курс - " + fmt.Sprintf("%f", e.ExRate)
}
