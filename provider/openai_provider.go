package provider

import (
	"context"
	"fmt"
	"ollamaproxy/model"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
	BaseProvider
	config     model.OllamaProxyConfig
	client     *openai.Client
	modelCache []model.OllamaModel
}

func (p *OpenAIProvider) SetConfig(config model.OllamaProxyConfig) {
	p.config = config
	var key string
	if config.Key != nil {
		key = *config.Key
	}
	if key == "" {
		panic("OpenAIProvider config: key is required")
	}
	cfg := openai.DefaultConfig(key)
	if config.Endpoint != "" {
		cfg.BaseURL = config.Endpoint
	}
	p.client = openai.NewClientWithConfig(cfg)
}

func (p *OpenAIProvider) GetTags() ([]model.OllamaModel, error) {
	if len(p.modelCache) > 0 {
		return p.modelCache, nil
	}
	ctx := context.Background()
	resp, err := p.client.ListModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("OpenAI ListModels: %w", err)
	}
	var models []model.OllamaModel
	for _, m := range resp.Models {
		models = append(models, model.OllamaModel{
			Name:       m.ID,
			Model:      m.ID,
			ModifiedAt: time.Now().Format(time.RFC3339),
			Size:       0,
			Digest:     "",
			Details:    nil,
		})
	}
	filtered := p.FilterModels(models, p.config)
	p.modelCache = filtered
	return filtered, nil
}

func (p *OpenAIProvider) Chat(req model.OllamaChatRequest) (<-chan model.OllamaChatResponse, error) {
	ch := make(chan model.OllamaChatResponse)

	go func() {
		defer close(ch)

		openaiReq := openai.ChatCompletionRequest{
			Model:  req.Model,
			Stream: req.Stream,
		}
		for _, msg := range req.Messages {
			var role string
			switch strings.ToLower(msg.Role) {
			case "system", "assistant", "user":
				role = strings.ToLower(msg.Role)
			default:
				role = "user"
			}
			openaiReq.Messages = append(openaiReq.Messages, openai.ChatCompletionMessage{
				Role:    role,
				Content: msg.Content,
			})
		}

		ctx := context.Background()
		stream, err := p.client.CreateChatCompletionStream(ctx, openaiReq)
		if err != nil {
			return
		}
		defer stream.Close()
		var contentBuilder strings.Builder
		for {
			response, err := stream.Recv()
			if err != nil {
				break // stream ended or error
			}
			for _, choice := range response.Choices {
				content := choice.Delta.Content
				contentBuilder.WriteString(content)
				ch <- model.OllamaChatResponse{
					Model:     req.Model,
					CreatedAt: time.Now().Format(time.RFC3339),
					Message: model.OllamaChatMessage{
						Role:    "assistant",
						Content: content,
					},
					Done: false,
				}
			}
		}
		ch <- model.OllamaChatResponse{
			Model:     req.Model,
			CreatedAt: time.Now().Format(time.RFC3339),
			Message: model.OllamaChatMessage{
				Role:    "assistant",
				Content: contentBuilder.String(),
			},
			Done: true,
		}
	}()

	return ch, nil
}
