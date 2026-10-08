package main

import "fmt"

type Owner struct {
	Name  string
	AccID string
}

type BankAccount struct {
	Owner
	Balance float64
}

func (ba BankAccount) balance() string {
	return fmt.Sprintf("%.2f", ba.Balance)
}

func main() {
	account := BankAccount{
		Owner:   Owner{Name: "Ibragim", AccID: "ACC-001"},
		Balance: 125000.50,
	}

	fmt.Println(account.Name)
	fmt.Println(account.AccID)
	fmt.Println(account.balance())
}
