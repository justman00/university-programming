package messaging

import "fmt"

// Template Method is a behavioral design pattern that allows you to defines a skeleton of an algorithm in a base class and let subclasses override the steps without changing the overall algorithm’s structure.
// https://refactoring.guru/design-patterns/template-method/go/example#example-0
type Sender interface {
	Send(to string, message string) error
}

// Simple Factory design pattern
// https://refactoring.guru/design-patterns/factory-method/go/example#example-0
func NewSender(kind string) (Sender, error) {
	if kind == "email" {
		return &emailSender{}, nil
	} else if kind == "sms" {
		return &smsSender{}, nil
	}

	return nil, fmt.Errorf("invalid sender kind: %s", kind)
}

type emailSender struct{}

func (e *emailSender) Send(to string, message string) error {
	fmt.Println(fmt.Sprintf("sending the following email content: %s to %s", message, to))
	return nil
}

type smsSender struct{}

func (e *smsSender) Send(to string, message string) error {
	fmt.Println(fmt.Sprintf("sending the following sms content: %s to %s", message, to))
	return nil
}
