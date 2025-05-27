package model

type OllamaModelDetails struct {
	ParentModel       string   `json:"parent_model"`
	Format            string   `json:"format"`
	Family            string   `json:"family"`
	Families          []string `json:"families"`
	ParameterSize     string   `json:"parameter_size"`
	QuantizationLevel string   `json:"quantization_level"`
}

type OllamaModel struct {
	Name       string              `json:"name"`
	Model      string              `json:"model"`
	ModifiedAt string              `json:"modified_at"`
	Size       int64               `json:"size"`
	Digest     string              `json:"digest"`
	Details    *OllamaModelDetails `json:"details,omitempty"`
}

type OllamaModelListResponse struct {
	Models []OllamaModel `json:"models"`
}
