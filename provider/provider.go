package provider

import "ollamaproxy/model"

type Provider interface {
	SetConfig(config model.OllamaProxyConfig)
	GetTags() ([]model.OllamaModel, error)
	Chat(req model.OllamaChatRequest) (<-chan model.OllamaChatResponse, error)
	FilterModels(models []model.OllamaModel, config model.OllamaProxyConfig) []model.OllamaModel
}

type BaseProvider struct{}

func (bp *BaseProvider) FilterModels(models []model.OllamaModel, config model.OllamaProxyConfig) []model.OllamaModel {
	if len(config.Models) == 0 {
		return []model.OllamaModel{}
	}
	filtered := []model.OllamaModel{}
	for _, m := range models {
		for _, allowed := range config.Models {
			if m.Name == allowed || m.Model == allowed {
				m.Name = config.Identifier + "@" + m.Name
				m.Model = config.Identifier + "@" + m.Model
				filtered = append(filtered, m)
			}
		}
	}
	return filtered
}
