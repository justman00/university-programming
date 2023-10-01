package chatgpt

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

type chatgptClient struct {
	openaiApiKey  string
	openaiBaseUrl string
	openaiModel   string
}

type Client interface {
	Question(ctx context.Context, question *Question) (*Answer, error)
}

func New(openaiApiKey string, openaiBaseUrl string, openaiModel string) Client {
	return &chatgptClient{
		openaiApiKey:  openaiApiKey,
		openaiBaseUrl: openaiBaseUrl,
		openaiModel:   openaiModel,
	}
}

type Question struct {
	HumanChatMessage  string
	SystemChatMessage string
}

type Answer struct {
	AnswerText string
}

func (chatgptClient *chatgptClient) Question(ctx context.Context, question *Question) (*Answer, error) {
	if len(strings.TrimSpace(question.HumanChatMessage)) < 10 {
		return nil, fmt.Errorf("the text '%s' has to be at least 10 characters long", question.HumanChatMessage)
	}

	chat, err := openai.NewChat(
		openai.WithModel(chatgptClient.openaiModel),
		openai.WithAPIType(openai.APITypeAzure),
		openai.WithBaseURL(chatgptClient.openaiBaseUrl),
		openai.WithToken(chatgptClient.openaiApiKey),
		openai.WithEmbeddingModel("text-embedding-ada-002"),
	)
	if err != nil {
		return nil, fmt.Errorf("openai client cannot be created: %s", err)
	}

	messages := []schema.ChatMessage{
		schema.SystemChatMessage{Content: question.SystemChatMessage},
		schema.HumanChatMessage{Content: question.HumanChatMessage},
	}
	answer, err := chat.Call(ctx, messages, llms.WithTemperature(0.0))
	if err != nil {
		return nil, fmt.Errorf("call openai: %s", err)
	}

	return &Answer{
		AnswerText: answer.Content,
	}, nil
}
