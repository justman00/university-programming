package payments

import "fmt"

type payment struct {
	paymentProcessors []PaymentProcessor
}

func NewPayment() *payment {
	return &payment{
		paymentProcessors: []PaymentProcessor{
			&CreditCardProcessor{},
			&PayPalAdapter{PayPal: &PayPal{}},
		},
	}
}

func (p *payment) Create() error {
	fmt.Println("payment created")

	return nil
}

func (p *payment) Process(method string) error {
	for _, processor := range p.paymentProcessors {
		if processor.ID() == method {
			processor.ProcessPayment(100)
		}
	}

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
