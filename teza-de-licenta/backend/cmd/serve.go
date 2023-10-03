package cmd

import (
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/justman00/teza-de-licenta/internal/api"
	"github.com/justman00/teza-de-licenta/internal/db"
	"github.com/justman00/teza-de-licenta/internal/importers/trustpilot"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ServeCMD() *cobra.Command {
	var serveCMD = &cobra.Command{
		Use:   "serve",
		Short: "`serve` este comanda care porneste serverul pentru teza de licenta. Acesta include pornirea serverului pentru afisarea si procesarea datelor.",
		Run: func(cmd *cobra.Command, args []string) {
			dbInstance, err := db.New()
			if err != nil {
				logrus.Fatalf("failed to create db instance: %v", err)
			}

			workerClient := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
			defer workerClient.Close()

			trustpilotClient := trustpilot.NewTrustpilotClient(os.Getenv("TRUSTPILOT_API_KEY"))
			h := api.New(dbInstance, trustpilotClient, workerClient)

			e := echo.New()
			g := e.Group("/api")

			mon := asynqmon.New(asynqmon.Options{
				RootPath: "/monitoring/tasks",
				RedisConnOpt: asynq.RedisClientOpt{
					Addr: os.Getenv("REDIS_ADDR"),
				},
			})

			g.Use(middleware.Logger())

			e.Any("/monitoring/tasks/*", echo.WrapHandler(mon))
			g.GET("/health", func(c echo.Context) error {
				return c.String(http.StatusOK, "Welcome to teza-de-licenta!")
			})
			g.GET("/reviews", h.GetReviewsHandler)
			g.POST("/enrol-client", h.EnrolClientHandler)

			e.Logger.Fatal(e.Start(":8080"))
		},
	}

	return serveCMD
}
