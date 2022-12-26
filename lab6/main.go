package main

import (
	"bufio"
	"fmt"
	"lab6/bank"
	"lab6/client"
	"os"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func FindClient(c *client.Client, info string) {
	if c.AccountNumber() == info || c.Surname() == info {
		var deposit int64
		var credit int64
		select {
		case deposit = <-c.CDeposit():
		case credit = <-c.CCredit():
		}
		fmt.Printf("ім'я - %s\n", c.Name())
		fmt.Printf("фамілія - %s\n", c.Surname())
		fmt.Printf("номер - %s\n", c.AccountNumber())
		fmt.Printf("кредит - %d\n", credit)
		fmt.Printf("депозит - %d\n", deposit)
	}
}

func CreateClient() (string, string, string, int64) {
	scannerC := bufio.NewScanner(os.Stdin)

	var name string
	var surName string
	accountNumber := "777"
	var summ int64

	fmt.Print("Ім'я клієнта: ")
	scannerC.Scan()
	name = scannerC.Text()
	fmt.Print("Фамілія клієнта: ")
	scannerC.Scan()

	surName = scannerC.Text()
	fmt.Print("Сума: ")
	scannerC.Scan()
	summ, _ = strconv.ParseInt(scannerC.Text(), 10, 0)
	if summ < 1 {
		fmt.Println("Помилка введення суми!")
		panic("Помилка введення")
	}
	return name, surName, accountNumber, summ
}

func HandleErrorBank(bankNum int64, banks int64) {
	if bankNum < 0 || bankNum > banks-1 {
		fmt.Println("Помилка вибору! Оберіть банк зі списку")
		panic("Помилка вибору")
	}
}

func main() {
	var banks []*bank.Bank
	flag := true

	scanner := bufio.NewScanner(os.Stdin)
	for flag == true {
		time.Sleep(time.Second / 2)
		fmt.Print("\nСтворити банк - 1" +
			"\nСтворити клієнта для роботи кредитом - 2" +
			"\nСтворити клієнта для роботи депозитом - 3" +
			"\nВивести клієнта за фамілією - 4" +
			"\nВивести клієнта за номером рахунку - 5" +
			"\nВийти - 0\n")
		scanner.Scan()
		w := scanner.Text()
		switch w {
		case "0":
			flag = false
		case "1":
			var name string
			var bankMoney int64

			fmt.Print("Назва банку: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Гроші банку: ")
			scanner.Scan()
			bankMoney, _ = strconv.ParseInt(scanner.Text(), 10, 0)
			if bankMoney < 1 {
				fmt.Println("Помилка введення суми!")
				panic("Помилка введення")
			}
			bank := bank.NewBank(name, bankMoney)
			banks = append(banks, bank)

			go bank.BankCredit(bank)
			go bank.BankDeposit(bank)

		case "2":
			if len(banks) < 1 {
				fmt.Println("Банків не створено!")
				break
			} else {
				name, surName, accountNumber, credit := CreateClient()
				var bankNum int64
				fmt.Println("Оберіть банк: ")
				for i, bank := range banks {
					fmt.Printf("%s - %d", bank.Name(), i)
				}
				fmt.Println("\nБанк: ")
				scanner.Scan()
				bankNum, _ = strconv.ParseInt(scanner.Text(), 10, 0)
				HandleErrorBank(bankNum, int64(len(banks)))
				var currentClient *client.Client
				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					client := client.NewClient(name, surName, accountNumber, 0, credit)
					newClients := append(banks[bankNum].Clients(), client)
					banks[bankNum].SetClients(newClients)
					currentClient = client
					return
				}()
				waitGroup.Wait()
				go func() {
					client := currentClient
					client.CreditActivity(client)
					return
				}()
			}
		case "3":
			if len(banks) < 1 {
				fmt.Println("Банків не створено!")
				break
			} else {
				name, surName, accountNumber, deposit := CreateClient()
				var bankNum int64
				fmt.Println("Оберіть банк: ")
				for i, bank := range banks {
					fmt.Printf("%s - %d", bank.Name(), i)
				}
				fmt.Println("\nБанк: ")
				scanner.Scan()
				bankNum, _ = strconv.ParseInt(scanner.Text(), 10, 0)
				HandleErrorBank(bankNum, int64(len(banks)))
				var currentClient *client.Client
				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					client := client.NewClient(name, surName, accountNumber, deposit, 0)
					newClients := append(banks[bankNum].Clients(), client)
					banks[bankNum].SetClients(newClients)
					currentClient = client
					return
				}()
				waitGroup.Wait()
				go func() {
					client := currentClient
					client.DepositActivity(client)
					return
				}()
			}
		case "4":
			if len(banks) < 1 {
				fmt.Println("Банків не створено!")
				break
			} else {
				var bankNum int64
				var surname string

				fmt.Println("Оберіть банк: ")
				for i, bank := range banks {
					fmt.Printf("%s - %d", bank.Name(), i)
				}
				fmt.Println("\nБанк: ")
				scanner.Scan()
				bankNum, _ = strconv.ParseInt(scanner.Text(), 10, 0)
				HandleErrorBank(bankNum, int64(len(banks)))
				fmt.Print("Фамілія клієнта: ")
				scanner.Scan()
				surname = scanner.Text()
				go func() {
					for _, client := range banks[bankNum].Clients() {
						FindClient(client, surname)
					}
					return
				}()
				time.Sleep(time.Second)
			}
		case "5":
			if len(banks) < 1 {
				fmt.Println("Банків не створено!")
				break
			} else {
				var bankNum int64
				var accountNumber string

				fmt.Println("Оберіть банк: ")
				for i, bank := range banks {
					fmt.Printf("%s - %d", bank.Name(), i)
				}
				fmt.Println("\nБанк: ")
				scanner.Scan()
				bankNum, _ = strconv.ParseInt(scanner.Text(), 10, 0)
				HandleErrorBank(bankNum, int64(len(banks)))
				fmt.Print("Номер облікового запису: ")
				scanner.Scan()
				accountNumber = scanner.Text()
				go func() {
					for _, client := range banks[bankNum].Clients() {
						FindClient(client, accountNumber)
					}
					return
				}()
				time.Sleep(time.Second)
			}
		default:
			fmt.Println("Помилка операції")
		}
	}

}
