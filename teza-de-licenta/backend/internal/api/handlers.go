package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/justman00/teza-de-licenta/internal/db"
	"github.com/justman00/teza-de-licenta/internal/importers/trustpilot"
	"github.com/justman00/teza-de-licenta/internal/tasks"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type handler struct {
	dbInstance       db.DB
	trustpilotClient *trustpilot.TrustpilotClient
	workerClient     *asynq.Client
}

func New(dbInstance db.DB, trustpilotClient *trustpilot.TrustpilotClient, workerClient *asynq.Client) *handler {
	return &handler{dbInstance: dbInstance, trustpilotClient: trustpilotClient, workerClient: workerClient}
}

type Review struct {
	ID                 string   `json:"id"`
	Rating             int      `json:"rating"`
	Contents           string   `json:"contents"`
	TopiClassification []string `json:"topic_classification"`
	Sentiment          string   `json:"sentiment"`
	Emotion            string   `json:"emotion"`
	Source             string   `json:"source"`
	CreatedAt          string   `json:"created_at"`
}

type Analysis struct {
	TopicClassification []string `json:"topic_classification"`
	Sentiment           string   `json:"sentiment"`
	Emotion             string   `json:"emotion"`
}

func (h *handler) GetReviewsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var source *string
	if s := c.QueryParam("source"); s != "" {
		source = &s
	}

	var limit int
	if l := c.QueryParam("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid limit")
		}

		limit = parsedLimit
	} else {
		limit = 1000
	}

	var topicClassification *string
	if tc := c.QueryParam("topic_classification"); tc != "" {
		topicClassification = &tc
	}

	var emotion *string
	if e := c.QueryParam("emotion"); e != "" {
		emotion = &e
	}

	var sentiment *string
	if s := c.QueryParam("sentiment"); s != "" {
		sentiment = &s
	}

	reviews, err := h.dbInstance.GetReviews(ctx, &db.GetReviewsParams{
		Limit:               limit,
		Source:              source,
		TopicClassification: topicClassification,
		Emotion:             emotion,
		Sentiment:           sentiment,
	})
	if err != nil {
		return fmt.Errorf("failed to get reviews: %w", err)
	}

	formattedReviews := make([]*Review, 0, len(reviews))
	for _, review := range reviews {
		var analysis *Analysis
		if review.Analysis != "" {
			if err := json.Unmarshal([]byte(review.Analysis), &analysis); err != nil {
				return fmt.Errorf("failed to unmarshal analysis: %w", err)
			}
		}

		formattedReviews = append(formattedReviews, &Review{
			ID:                 review.ID.String(),
			Rating:             review.Rating,
			Contents:           review.Review,
			TopiClassification: analysis.TopicClassification,
			Sentiment:          analysis.Sentiment,
			Emotion:            analysis.Emotion,
			Source:             review.Source,
			CreatedAt:          review.ReviewCreatedAt.String(),
		})
	}

	return c.JSON(http.StatusOK, formattedReviews)
}

type EnrolClient struct {
	EnrolClientID string `json:"client_id"`
	Source        string `json:"source"`
	Name          string `json:"name"`
}

func (h *handler) EnrolClientHandler(c echo.Context) error {
	ctx := c.Request().Context()

	enrolClient := new(EnrolClient)
	if err := c.Bind(enrolClient); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	entity, err := h.dbInstance.CreateEntity(ctx, &db.CreateEntityParams{
		ExternalID: enrolClient.EnrolClientID,
		Name:       enrolClient.Name,
		Source:     enrolClient.Source,
	})
	if err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}

	reviews, err := h.trustpilotClient.GetReviews(ctx, enrolClient.EnrolClientID)
	if err != nil {
		logrus.Errorf("failed to get reviews: %v", err)

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get reviews")
	}

	for _, review := range reviews {
		task, err := tasks.NewReviewSubmittedTask(tasks.TypeReviewSubmittedPayload{
			CreatedAt: review.CreatedAt,
			ReviewID:  review.ID,
			URL:       review.URL(),
			Contents:  review.Text,
			Source:    enrolClient.Source,
			EntityID:  entity.ID,
			Rating:    review.Stars,
		})
		if err != nil {
			logrus.Errorf("failed to create task: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create task")
		}

		if _, err := h.workerClient.EnqueueContext(ctx, task, asynq.Queue("reviews")); err != nil {
			logrus.Errorf("failed to enqueue task: %v", err)

			return echo.NewHTTPError(http.StatusInternalServerError, "failed to enqueue task")
		}
	}

	return c.JSON(http.StatusOK, entity)
}
