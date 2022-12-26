package bank

import (
	"lab6/client"
	"time"
)

type Bank struct {
	name      string
	bankMoney int64
	deposit   int64
	credit    int64
	clients   []*client.Client

	BankDeposit func(bank *Bank)
	BankCredit  func(bank *Bank)
}

func (b *Bank) Name() string {
	return b.name
}

func (b *Bank) SetName(name string) {
	b.name = name
}

func (b *Bank) BankMoney() int64 {
	return b.bankMoney
}

func (b *Bank) SetBankMoney(bankMoney int64) {
	b.bankMoney = bankMoney
}

func (b *Bank) Deposit() int64 {
	return b.deposit
}

func (b *Bank) SetDeposit(deposit int64) {
	b.deposit = deposit
}

func (b *Bank) Credit() int64 {
	return b.credit
}

func (b *Bank) SetCredit(credit int64) {
	b.credit = credit
}

func (b *Bank) Clients() []*client.Client {
	return b.clients
}

func (b *Bank) SetClients(clients []*client.Client) {
	b.clients = clients
}

func NewBank(name string, bankMoney int64) *Bank {
	return &Bank{
		name:      name,
		bankMoney: bankMoney,
		deposit:   0,
		credit:    0,
		clients:   []*client.Client{},

		BankDeposit: func(bank *Bank) {
			for {
				var deposit int64
				for _, client := range bank.Clients() {
					deposit += <-client.CDeposit()
				}
				bank.SetDeposit(deposit)
				time.Sleep(time.Second)

			}
		},
		BankCredit: func(bank *Bank) {
			for {
				var credit int64
				for _, client := range bank.Clients() {
					credit += <-client.CCredit()
				}
				bank.SetCredit(credit)
				time.Sleep(time.Second)
			}
		},
	}
}
