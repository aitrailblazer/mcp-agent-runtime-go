package types

type PromptRequest struct {
	Prompt  string                 `json:"prompt"`
	Tools   []string               `json:"tools"`
	Context map[string]interface{} `json:"context"`
}

type PromptResponse struct {
	Status  string                 `json:"status"`
	Prompt  string                 `json:"prompt"`
	Tools   []string               `json:"tools"`
	Context map[string]interface{} `json:"context"`
}
