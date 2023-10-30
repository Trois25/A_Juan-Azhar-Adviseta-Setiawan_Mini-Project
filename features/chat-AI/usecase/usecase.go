package usecase

import (
	"context"
	"fmt"

	"event_ticket/features/chat-AI/dto"

	"github.com/sashabaranov/go-openai"
)

type AnimeRecomendationUsecase interface {
	AnimeRecomendation(request dto.RequestData, key string) (string, error)
}

type Usecase struct{}

func NewUseCase() AnimeRecomendationUsecase {
	return &Usecase{}
}

func (uc *Usecase) getCompletionMessages(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (us *Usecase) AnimeRecomendation(request dto.RequestData, key string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(key)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content : "You must give anime recommendation for user with information about the anime like synopsis, MAL rating,release date and total episode of the anime you recommended",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Anime recommendation with Genre : %s and description : %s with the year of release : %s", request.Genre, request.Description, request.Year),
		},
	}

	response, err := us.getCompletionMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := response.Choices[0].Message.Content
	return answer, nil
}
