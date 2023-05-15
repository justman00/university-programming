package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/models"
	"github.com/uptrace/bunrouter"
)

type handler struct {
	clientModels *models.ClientModels
	bookingModels *models.BookingModels
}

func NewHandler(clientModels *models.ClientModels, bookingModels *models.BookingModels) *handler {
	return &handler{
		clientModels: clientModels,
		bookingModels: bookingModels,
	}
}

type CreateClientRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Location string `json:"location"`
	WorkingHours string `json:"working_hours"`
	TimePerTable int `json:"time_per_table"` // in minutes
	Tables int `json:"tables"`
}

func (h *handler) CreateClient(w http.ResponseWriter, req bunrouter.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	var createClientRequest CreateClientRequest
	if err := json.Unmarshal(b, &createClientRequest); err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	
	if err := h.clientModels.Create(models.Client{
		Name: createClientRequest.Name,
		Email: createClientRequest.Email,
		Location: createClientRequest.Location,
		WorkingHours: createClientRequest.WorkingHours,
		TimePerTable: createClientRequest.TimePerTable,
		Tables: createClientRequest.Tables,
	}); err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, "Created")
}

type CreateReservationRequest struct {
	ClientID string `json:"client_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}

func (h *handler) CreateReservation(w http.ResponseWriter, req bunrouter.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	var createReservationRequest CreateReservationRequest
	if err := json.Unmarshal(b, &createReservationRequest); err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	if err := h.bookingModels.Create(models.Booking{
		ClientID: createReservationRequest.ClientID,
		StartTime: createReservationRequest.StartTime,
		EndTime: createReservationRequest.EndTime,
	}); err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	// TODO:
	// get available tables
	// check if there are any available tables
	// if there are no available tables, return error
	// if there are available tables, pick one
	// create booking

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, "Created")
}

func (h *handler) GetReservations(w http.ResponseWriter, req bunrouter.Request) error {
	query := req.URL.Query()

	clientID := query.Get("client_id")
	if clientID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("client_id is required")
	}

	date := query.Get("date")
	if date == "" {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("date is required")
	}

	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("failed to parse date: %w", err)
	}

	bookings, err := h.bookingModels.GetByClientID(clientID, time)
	if err != nil {
		return fmt.Errorf("failed to get reservations: %w", err)
	}
	
	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, bookings)
}