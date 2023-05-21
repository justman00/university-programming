package payments

import "fmt"

type payment struct{}

func NewPayment() *payment {
	return &payment{}
}

func (p *payment) Create() error {
	fmt.Println("payment created")

	return nil
}

func (p *payment) Process() error {
	fmt.Println("payment get")

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
