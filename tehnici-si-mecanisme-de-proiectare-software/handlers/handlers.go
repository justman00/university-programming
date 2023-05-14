package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/models"
	"github.com/uptrace/bunrouter"
)

type handler struct {
	ClientModels *models.ClientModels
}

func NewHandler(clientModels *models.ClientModels) *handler {
	return &handler{
		ClientModels: clientModels,
	}
}

type CreateClientRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Location string `json:"location"`
	WorkingHours string `json:"working_hours"`
	TimePerTable int `json:"time_per_table"` // in minutes
}

func (h *handler) CreateClient(w http.ResponseWriter, req bunrouter.Request) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	fmt.Println(string(b))

	var createClientRequest CreateClientRequest
	if err := json.Unmarshal(b, &createClientRequest); err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	
	if err := h.ClientModels.Create(models.Client{
		Name: createClientRequest.Name,
		Email: createClientRequest.Email,
		Location: createClientRequest.Location,
		WorkingHours: createClientRequest.WorkingHours,
		TimePerTable: createClientRequest.TimePerTable,
	}); err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, "Created")
}

func (h *handler) CreateReservation(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}

func (h *handler) GetReservations(w http.ResponseWriter, req bunrouter.Request) error {

	return nil
}