package payments

import "fmt"

type payment struct {
	paymentProvessor PaymentProcessor // bridge
}

func NewPayment() *payment {
	return &payment{
		paymentProvessor: &CreditCardProcessor{}, // default
	}
}

func (p *payment) Create() error {
	fmt.Println("payment created")

	return nil
}

func (p *payment) Process() error {
	p.paymentProvessor.ProcessPayment(0)

	return nil
}

func (p *payment) Cancel() error {
	fmt.Println("payment canceled")

	return nil
}

func (p *payment) Refund() error {
	fmt.Println("payment refunded")

	return nil
}

func (p *payment) Get() error {
	fmt.Println("payment get")

	return nil
}

// Bridge is a structural design pattern that divides business logic or huge class into separate class hierarchies that can be developed independently.
// https://refactoring.guru/design-patterns/bridge/go/example#example-0
func (p *payment) SetPaymentProcessor(paymentProcessor PaymentProcessor) {
	p.paymentProvessor = paymentProcessor
}
