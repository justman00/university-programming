package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeReviewSubmitted = "review:submitted"
)

type TypeReviewSubmittedPayload struct {
	CreatedAt time.Time
	ReviewID  string
	URL       string
	Contents  string
	Source    string
	Rating    int
}

func NewReviewSubmittedTask(submittedReview TypeReviewSubmittedPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(submittedReview)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeReviewSubmitted, payload), nil
}

// ReviewProcessor implements asynq.Handler interface.
type ReviewProcessor struct {
	// ... fields for struct
}

func (processor *ReviewProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p TypeReviewSubmittedPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal review submitted: %v: %w", err, asynq.SkipRetry)
	}

	fmt.Println("ReviewProcessor: received task:", p.Contents)

	return nil
}

func NewReviewProcessor() *ReviewProcessor {
	return &ReviewProcessor{}
}
