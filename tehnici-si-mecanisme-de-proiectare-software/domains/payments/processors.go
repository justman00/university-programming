package payments

import (
	"fmt"
)

// Existing system's payment processor interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	ID() string
}

// Existing system's credit card processor
type CreditCardProcessor struct{}

func (p *CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}

func (p *CreditCardProcessor) ID() string {
	return "credit_card"
}

// PayPal API that we want to adapt to our PaymentProcessor interface
type PayPal struct{}

func (p *PayPal) SendPayment(amount float64) error {
	fmt.Printf("Sending PayPal payment of $%.2f\n", amount)
	return nil
}

// Adapter is a structural design pattern, which allows incompatible objects to collaborate.
// https://refactoring.guru/design-patterns/adapter/go/example#example-0
// PayPalAdapter adapts the PayPal API to the PaymentProcessor interface
type PayPalAdapter struct {
	PayPal *PayPal
}

func (a *PayPalAdapter) ProcessPayment(amount float64) error {
	return a.PayPal.SendPayment(amount)
}

func (a *PayPalAdapter) ID() string {
	return "paypal"
}
