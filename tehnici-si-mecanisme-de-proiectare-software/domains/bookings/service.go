package bookings

import (
	"fmt"
	"time"

	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/domains/payments"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/messaging"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/models"
)

type payment interface {
	Create() error
	Process() error
	SetPaymentProcessor(paymentProcessor payments.PaymentProcessor)
}

type Service struct {
	payment       payment
	clientModels  *models.ClientModels
	bookingModels *models.BookingModels
	sender        messaging.Sender
}

type CreateReservationRequest struct {
	ClientID    string    `json:"client_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TableNumber int       `json:"table_number"`
	TableType   string    `json:"table_type"`
}

// Facade is a structural design pattern that provides a simplified (but limited) interface to a complex system of classes, library or framework.
// https://refactoring.guru/design-patterns/facade/go/example#example-0
func NewService(clientModels *models.ClientModels, bookingModels *models.BookingModels, sender messaging.Sender) *Service {
	return &Service{
		payment:       payments.NewPayment(),
		clientModels:  clientModels,
		bookingModels: bookingModels,
		sender:        sender,
	}
}

func (s *Service) CreateBooking(createReservationRequest CreateReservationRequest) error {
	// get client
	client, err := s.clientModels.GetByID(createReservationRequest.ClientID)
	if err != nil {
		return fmt.Errorf("failed to get client: %w", err)
	}

	restaurant, err := models.GetRestaurantFactory(client.Type)
	if err != nil {
		return fmt.Errorf("failed to get restaurant factory: %w", err)
	}

	if createReservationRequest.TableType == "square" {
		createReservationRequest.TableType = restaurant.CreateSquareTable().GetInfo()
	} else if createReservationRequest.TableType == "round" {
		createReservationRequest.TableType = restaurant.CreateRoundTable().GetInfo()
	} else {
		return fmt.Errorf("invalid table type, only allowed round or square")
	}

	if err := s.bookingModels.Create(models.Booking{
		ClientID:    createReservationRequest.ClientID,
		StartTime:   createReservationRequest.StartTime,
		EndTime:     createReservationRequest.EndTime,
		TableNumber: createReservationRequest.TableNumber,
		TableType:   createReservationRequest.TableType,
	}); err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	// set payment processor as liked by client
	paypalPaymentProcessor := &payments.PayPalAdapter{PayPal: &payments.PayPal{}}
	s.payment.SetPaymentProcessor(paypalPaymentProcessor)

	if err := s.takePayment(); err != nil {
		// if fails, retry payment
		s.payment.SetPaymentProcessor(&payments.CreditCardProcessor{}) // usage of the bridge

		if err := s.takePayment(); err != nil {
			return fmt.Errorf("failed to take payment: %w", err)
		}
	}
	
	go s.notify(client.Email, fmt.Sprintf("Reservation created for table with id %v and of type: %v", createReservationRequest.TableNumber, createReservationRequest.TableType))

	return nil
}

func (s *Service) takePayment() error {
	if err := s.payment.Create(); err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	if err := s.payment.Process(); err != nil {
		return fmt.Errorf("failed to process payment: %w", err)
	}

	return nil
}

func (s *Service) notify(to, message string) error {
	err := s.sender.Send(to, message)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to send communication to client %v: %w", to, err))
	}

	return nil
}
