package service

import (
	"errors"
	"ollamaproxy/config"
	"ollamaproxy/model"
	"ollamaproxy/provider"
	"strings"
)

type providerWithConfig struct {
	Config   model.OllamaProxyConfig
	Provider provider.Provider
}

type OllamaProxyService struct {
	providers []providerWithConfig
}

func NewOllamaProxyService(configLoader *config.ConfigLoader) *OllamaProxyService {
	proxies, err := configLoader.LoadConfig()
	if err != nil {
		panic(err)
	}
	providers := make([]providerWithConfig, 0, len(proxies))
	for _, proxy := range proxies {
		var prov provider.Provider
		switch proxy.Implementation {
		case model.OpenAIApiV1:
			prov = &provider.OpenAIProvider{}
		default:
			continue
		}
		prov.SetConfig(proxy)
		providers = append(providers, providerWithConfig{Config: proxy, Provider: prov})
	}
	return &OllamaProxyService{providers: providers}
}

func (s *OllamaProxyService) GetTags() ([]model.OllamaModel, error) {
	var all []model.OllamaModel
	for _, pc := range s.providers {
		tags, err := pc.Provider.GetTags()
		if err != nil {
			continue
		}
		all = append(all, tags...)
	}
	return all, nil
}

func (s *OllamaProxyService) Chat(req model.OllamaChatRequest) (<-chan model.OllamaChatResponse, error) {
	provCfg := s.routeModelByIdentifier(req.Model)
	if provCfg == nil {
		return nil, errors.New("model must be in 'provider@modelName' format")
	}
	parts := strings.SplitN(req.Model, "@", 2)
	if len(parts) != 2 {
		return nil, errors.New("model must be in 'provider@modelName' format")
	}
	req.Model = parts[1]
	return provCfg.Provider.Chat(req)
}

func (s *OllamaProxyService) routeModelByIdentifier(model string) *providerWithConfig {
	identifier := strings.SplitN(model, "@", 2)[0]
	for i, p := range s.providers {
		if p.Config.Identifier == identifier {
			return &s.providers[i]
		}
	}
	return nil
}
