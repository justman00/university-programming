package cmd

import (
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/justman00/teza-de-licenta/internal/importers/trustpilot"
	"github.com/justman00/teza-de-licenta/internal/tasks"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type EnrolClient struct {
	EnrolClientID string `json:"client_id"`
	Source        string `json:"source"`
}

func ServeCMD() *cobra.Command {
	var serveCMD = &cobra.Command{
		Use:   "serve",
		Short: "`serve` este comanda care porneste serverul pentru teza de licenta. Acesta include pornirea serverului pentru afisarea si procesarea datelor.",
		Run: func(cmd *cobra.Command, args []string) {
			workerClient := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
			defer workerClient.Close()

			trustpilotClient := trustpilot.NewTrustpilotClient(os.Getenv("TRUSTPILOT_API_KEY"))

			e := echo.New()
			g := e.Group("/api")

			mon := asynqmon.New(asynqmon.Options{
				RootPath: "/monitoring/tasks",
				RedisConnOpt: asynq.RedisClientOpt{
					Addr: os.Getenv("REDIS_ADDR"),
				},
			})
			e.Any("/monitoring/tasks/*", echo.WrapHandler(mon))

			g.GET("/health", func(c echo.Context) error {
				return c.String(http.StatusOK, "Welcome to teza-de-licenta!")
			})

			g.POST("/enrol-client", func(c echo.Context) error {
				ctx := c.Request().Context()

				enrolClient := new(EnrolClient)
				if err := c.Bind(enrolClient); err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
				}

				reviews, err := trustpilotClient.GetReviews(ctx, enrolClient.EnrolClientID)
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
						Rating:    review.Stars,
					})
					if err != nil {
						logrus.Errorf("failed to create task: %v", err)

						return echo.NewHTTPError(http.StatusInternalServerError, "failed to create task")
					}

					if _, err := workerClient.EnqueueContext(ctx, task, asynq.Queue("reviews")); err != nil {
						logrus.Errorf("failed to enqueue task: %v", err)

						return echo.NewHTTPError(http.StatusInternalServerError, "failed to enqueue task")
					}
				}

				return c.String(http.StatusOK, "Hello, World!")
			})

			e.Logger.Fatal(e.Start(":8080"))
		},
	}

	return serveCMD
}
