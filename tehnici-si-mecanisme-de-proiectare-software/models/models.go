package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ClientModels struct {
	DB *sql.DB
}

type Client struct  {
	ID string
	Name string
	Email string
	Location string
	WorkingHours string
	TimePerTable int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type workingHours struct {
	From string `json:"from"`
	To string `json:"to"`
}

func (c *ClientModels) Create(client Client) error {
	if client.ID == "" {
		client.ID = uuid.New().String()
	}

	splitWorkingHours := strings.Split(client.WorkingHours, "-")
	workingHours := workingHours{
		From: splitWorkingHours[0],
		To: splitWorkingHours[1],
	}

	workingHoursJSON, err := json.Marshal(workingHours)
	if err != nil {
		return fmt.Errorf("failed to marshal working hours: %w", err)
	}

	query := `
		INSERT INTO clients (id, name, email, slug, location, working_hours, time_per_table) VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, err = c.DB.Exec(
		query, 
		client.ID, 
		client.Name, 
		client.Email, 
		createSlug(client.Name), 
		client.Location, 
		workingHoursJSON, 
		client.TimePerTable,
	)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}

func createSlug(input string) string {
	// Convert to lowercase
	input = strings.ToLower(input)
	
	// Replace spaces with hyphens
	input = strings.Replace(input, " ", "-", -1)

	// Remove non-alphanumeric and non-hyphen characters
	reg, _ := regexp.Compile("[^a-z0-9-]+")
	slug := reg.ReplaceAllString(input, "")

	return slug
}

type BookingModels struct {
	DB *sql.DB
}

type Booking struct {
	ID string
	ClientID string
	StartTime time.Time
	EndTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}