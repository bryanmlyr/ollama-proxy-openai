package model

type OllamaProxyImplementation string

const (
	OllamaApiV1 OllamaProxyImplementation = "OLLAMA_API_V1"
	OpenAIApiV1 OllamaProxyImplementation = "OPENAI_API_V1"
)

type OllamaProxyConfig struct {
	Identifier     string                    `json:"identifier"`
	Implementation OllamaProxyImplementation `json:"implementation"`
	Models         []string                  `json:"models"`
	Endpoint       string                    `json:"endpoint"`
	Key            *string                   `json:"key"`
}
