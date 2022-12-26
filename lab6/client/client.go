package client

import (
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	DepositActivity func(client *Client)
	CreditActivity  func(client *Client)

	name          string
	surname       string
	accountNumber string
	cDeposit      chan int64
	cCredit       chan int64
}

func (c *Client) CDeposit() chan int64 {
	return c.cDeposit
}

func (c *Client) CDepositValue() int64 {
	return <-c.cDeposit
}

func (c *Client) SetCDeposit(cDeposit chan int64) {
	c.cDeposit = cDeposit
}

func (c *Client) CCredit() chan int64 {
	return c.cCredit
}

func (c *Client) CCreditValue() int64 {
	return <-c.cCredit
}

func (c *Client) SetCCredit(cCredit chan int64) {
	c.cCredit = cCredit
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) Surname() string {
	return c.surname
}

func (c *Client) SetSurname(surname string) {
	c.surname = surname
}

func (c *Client) AccountNumber() string {
	return c.accountNumber
}

func (c *Client) SetAccountNumber(accountNumber string) {
	c.accountNumber = accountNumber
}

func NewClient(name string, surname string, accountNumber string, cDeposit int64, cCredit int64) *Client {
	cDepositCh := make(chan int64)
	cCreditCh := make(chan int64)
	go func() {
		cDepositCh <- cDeposit
		cCreditCh <- cCredit
	}()
	return &Client{
		name:          name,
		surname:       surname,
		accountNumber: accountNumber,
		cDeposit:      cDepositCh,
		cCredit:       cCreditCh,

		CreditActivity: func(client *Client) {
			credit := <-client.CCredit()
			creditDynamic := credit
			for creditDynamic > 0 {
				res := creditDynamic - int64(rand.Intn(int(credit/8)))
				client.CCredit() <- res
				creditDynamic = res
				time.Sleep(time.Second)
			}
			fmt.Println(client.Name() + " " + client.Surname() + " кредит виплачено")
			return
		},
		DepositActivity: func(client *Client) {
			deposit := <-client.CDeposit()
			depositDynamic := deposit
			for depositDynamic > 0 {
				res := depositDynamic - int64(rand.Intn(int(deposit/8)))
				client.CDeposit() <- res
				depositDynamic = res
				time.Sleep(time.Second)
			}
			fmt.Println(client.Name() + " " + client.Surname() + " депосит знято")
			return
		},
	}
}
