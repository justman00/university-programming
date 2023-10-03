package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the postgres driver
)

type DB interface {
	InsertReview(ctx context.Context, review *InsertReviewParams) error
	GetReviews(ctx context.Context, params *GetReviewsParams) ([]*Review, error)
	CreateEntity(ctx context.Context, params *CreateEntityParams) (*Entity, error)
}

type dbImpl struct {
	*sqlx.DB
}

func New() (DB, error) {
	var err error

	// Connect to the database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	// Run the database migrations
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("create migration instance: %w", err)
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) && err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &dbImpl{db}, nil
}

type Review struct {
	ID              uuid.UUID `db:"id"`
	Rating          int       `db:"rating"`
	Source          string    `db:"source"`
	Review          string    `db:"review"`
	Analysis        string    `db:"analysis"`
	OriginalPayload string    `db:"original_payload"`
	EntityID        string    `db:"entity_id"`
	ReviewCreatedAt time.Time `db:"review_created_at"`
	ReviewUpdatedAt time.Time `db:"review_updated_at"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type InsertReviewParams struct {
	ID              uuid.UUID `db:"id"`
	Rating          int       `db:"rating"`
	Source          string    `db:"source"`
	Review          string    `db:"review"`
	Analysis        string    `db:"analysis"`
	OriginalPayload string    `db:"original_payload"`
	EntityID        string    `db:"entity_id"`
	ReviewCreatedAt time.Time `db:"review_created_at"`
	ReviewUpdatedAt time.Time `db:"review_updated_at"`
}

func (db *dbImpl) InsertReview(ctx context.Context, review *InsertReviewParams) error {
	query := `
        INSERT INTO reviews (id, rating, source, review, analysis, original_payload, review_created_at, review_updated_at, entity_id)
        VALUES (:id, :rating, :source, :review, :analysis, :original_payload, :review_created_at, :review_updated_at, :entity_id)
    `
	_, err := db.NamedExec(query, review)
	if err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}

type GetReviewsParams struct {
	TopicClassification *string `db:"topic_classification"`
	Emotion             *string `db:"emotion"`
	Sentiment           *string `db:"sentiment"`
	Source              *string `db:"source"`
	Limit               int     `db:"limit"`
}

func (db *dbImpl) GetReviews(ctx context.Context, params *GetReviewsParams) ([]*Review, error) {
	query := `
	SELECT *
	FROM reviews
	WHERE ($1::text is NULL OR source = $1)
		AND ($2::text is NULL OR analysis->>'emotion' = $2)
		AND ($3::text is NULL OR analysis->>'sentiment' = $3)
		AND ($4::text is NULL OR analysis->'topic_classification' @> to_jsonb(string_to_array($4, ',')::text[]))
	ORDER BY review_created_at DESC
	LIMIT $5;
	`

	reviews := []*Review{}
	err := db.SelectContext(ctx, &reviews, query,
		params.Source,
		params.Emotion,
		params.Sentiment,
		params.TopicClassification,
		params.Limit,
	)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

type Entity struct {
	ID         string    `db:"id"`
	ExternalID string    `db:"external_id"`
	Name       string    `db:"name"`
	Source     string    `db:"source"`
	CreatedAt  time.Time `db:"created_at"`
}

type CreateEntityParams struct {
	ExternalID string `db:"external_id"`
	Name       string `db:"name"`
	Source     string `db:"source"`
}

func (db *dbImpl) CreateEntity(ctx context.Context, params *CreateEntityParams) (*Entity, error) {
	query := `
        INSERT INTO entities (external_id, name, source)
        VALUES ($1, $2, $3)
        RETURNING id, external_id, name, source, created_at
    `
	entity := &Entity{}

	if err := db.QueryRowContext(ctx,
		query,
		params.ExternalID,
		params.Name,
		params.Source,
	).Scan(
		&entity.ID,
		&entity.ExternalID,
		&entity.Name,
		&entity.Source,
		&entity.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("insert entity: %w", err)
	}

	return entity, nil
}
