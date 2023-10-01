package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/justman00/teza-de-licenta/internal/db"
	"github.com/justman00/teza-de-licenta/internal/services/chatgpt"
	"github.com/justman00/teza-de-licenta/internal/tasks"
)

func WorkerCMD() *cobra.Command {
	var serveCMD = &cobra.Command{
		Use:   "worker",
		Short: "`worker` este comanda care porneste procesul ce va prelucra datele in background.",
		Run: func(cmd *cobra.Command, args []string) {
			dbInstance, err := db.New()
			if err != nil {
				logrus.Fatalf("failed to create db instance: %v", err)
			}

			chatgptClient := chatgpt.New(os.Getenv("OPENAI_API_KEY"), os.Getenv("OPENAI_BASE_URL"), os.Getenv("OPENAI_MODEL"))
			taskProcessor := tasks.NewReviewProcessor(chatgptClient, dbInstance)

			// start the asynq worker
			srv := asynq.NewServer(
				asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")},
				asynq.Config{
					// Specify how many concurrent workers to use
					Concurrency: 1,
					// Optionally specify multiple queues with different priority.
					Queues: map[string]int{
						"reviews": 10,
					},
					Logger: logrus.New(),
					// See the godoc for other configuration options
					BaseContext: func() context.Context {
						return cmd.Context()
					},
					StrictPriority:  true,
					ShutdownTimeout: 10 * time.Second,
				},
			)

			// mux maps a type to a handler
			mux := asynq.NewServeMux()
			mux.Handle(tasks.TypeReviewSubmitted, taskProcessor)

			if err := srv.Run(mux); err != nil {
				log.Fatalf("failed to start worker server: %v", err)
			}
		},
	}

	return serveCMD
}
