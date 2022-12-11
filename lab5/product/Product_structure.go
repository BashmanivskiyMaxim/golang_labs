package product

import (
	"lab5/currency"
	"strconv"
)

type Product struct {
	Name     string
	Cost     int
	Quantity int
	Producer string
	Weight   int
	Money    currency.Currency
}

// NewProduct Конструктор

func NewProduct() *Product {
	const (
		defaultName     = "Name"
		defaultCost     = 1
		defaultQuantity = 1
		defaultProducer = "Producer"
		defaultWeight   = 1
	)
	defaultMoney := currency.Currency{Name: "name", ExRate: 1.1}

	return &Product{
		Name:     defaultName,
		Cost:     defaultCost,
		Quantity: defaultQuantity,
		Producer: defaultProducer,
		Weight:   defaultWeight,
		Money:    defaultMoney,
	}
}

func (e *Product) SetProduct(Name string, Cost int, Quantity int, Producer string, Weight int, NameVal string, ExRate float64) {
	e.SetName(Name)
	e.SetCost(Cost)
	e.SetQuantity(Quantity)
	e.SetProducer(Producer)
	e.SetWeight(Weight)
	e.SetMoney(NameVal, ExRate)
}

func (e *Product) GetName() string {
	return e.Name
}
func (e *Product) GetCost() int {
	return e.Cost
}
func (e *Product) GetQuantity() int {
	return e.Quantity
}
func (e *Product) GetProducer() string {
	return e.Producer
}
func (e *Product) GetWeight() int {
	return e.Weight
}

func (e *Product) SetName(name string) {
	e.Name = name
}
func (e *Product) SetMoney(name string, exrate float64) {
	e.Money.Name = name
	e.Money.ExRate = exrate
}

func (e *Product) SetCost(cost int) {
	e.Cost = cost
}
func (e *Product) SetQuantity(quantity int) {
	e.Quantity = quantity
}
func (e *Product) SetProducer(producer string) {
	e.Producer = producer
}
func (e *Product) SetWeight(weight int) {
	e.Weight = weight
}

func (e *Product) GetPriceIn() float64 {
	return float64(e.Cost) * e.Money.ExRate
}

func (e *Product) GetTotalPrice() float64 {
	return float64(e.Cost) * e.Money.ExRate * float64(e.Quantity)
}

func (e *Product) GetTotalWeight() string {
	return strconv.Itoa(e.Quantity*e.Weight) + " кг"
}
