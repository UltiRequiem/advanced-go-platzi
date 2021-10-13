package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

// Adapter --------
type BankPaymentAdapter struct {
	bankAccount int
	BankPayment *BankPayment
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

// -----------------

func (CashPayment) Pay() {
	fmt.Println("Payment using cash...")
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using Bank Account %d...\n", bankAccount)
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	myBank := &BankPaymentAdapter{5, &BankPayment{}}
	ProcessPayment(myBank)
}
