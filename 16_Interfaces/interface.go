package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	razorpayPaymentGe := razorpay{}
	razorpayPaymentGe.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment via razorpay", amount)
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(100)
}
