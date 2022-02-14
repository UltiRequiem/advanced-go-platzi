package main

import "fmt"

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

// This requires an argument so it does not satisfy the Payment interface
func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using Bank Account %d...\n", bankAccount)
}

type BankPaymentAdapter struct {
	bankAccount int
	BankPayment *BankPayment
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash...")
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	myBank := &BankPaymentAdapter{5, &BankPayment{}}
	ProcessPayment(myBank)
}
