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

type Client struct {
	ID           string
	Name         string
	Email        string
	Location     string
	WorkingHours string
	TimePerTable int
	Tables       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Type         string
	Slug         string
}

type workingHours struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (c *ClientModels) Create(client Client) error {
	if client.ID == "" {
		client.ID = uuid.New().String()
	}

	splitWorkingHours := strings.Split(client.WorkingHours, "-")
	workingHours := workingHours{
		From: splitWorkingHours[0],
		To:   splitWorkingHours[1],
	}

	workingHoursJSON, err := json.Marshal(workingHours)
	if err != nil {
		return fmt.Errorf("failed to marshal working hours: %w", err)
	}

	query := `
		INSERT INTO clients (id, name, email, slug, tables, location, working_hours, time_per_table, type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err = c.DB.Exec(
		query,
		client.ID,
		client.Name,
		client.Email,
		client.Slug,
		client.Tables,
		client.Location,
		workingHoursJSON,
		client.TimePerTable,
		client.Type,
	)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}

func (c *ClientModels) GetByID(id string) (Client, error) {
	var client Client

	query := `
		SELECT id, name, email, slug, tables, location, working_hours, time_per_table, tables, type FROM clients WHERE id = $1;
	`
	err := c.DB.QueryRow(query, id).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Slug,
		&client.Tables,
		&client.Location,
		&client.WorkingHours,
		&client.TimePerTable,
		&client.Tables,
		&client.Type,
	)
	if err != nil {
		return Client{}, fmt.Errorf("failed to get client by id: %w", err)
	}

	return client, nil
}

func CreateSlug(input string) string {
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
	ID          string
	ClientID    string
	StartTime   time.Time
	EndTime     time.Time
	TableNumber int
	TableType   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *BookingModels) Create(booking Booking) error {
	if booking.ID == "" {
		booking.ID = uuid.New().String()
	}

	query := `
		INSERT INTO bookings (id, client_id, start_time, end_time, table_number, type) VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := b.DB.Exec(
		query,
		booking.ID,
		booking.ClientID,
		booking.StartTime,
		booking.EndTime,
		booking.TableNumber,
		booking.TableType,
	)
	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}

	return nil
}

func (b *BookingModels) GetByClientID(clientID string, date time.Time) ([]Booking, error) {
	query := `
		SELECT id, client_id, start_time, end_time, created_at, updated_at, table_number, type FROM bookings WHERE client_id = $1 AND start_time::date = $2::date;
	`
	rows, err := b.DB.Query(query, clientID, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.ClientID, &booking.StartTime, &booking.EndTime, &booking.CreatedAt, &booking.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}
