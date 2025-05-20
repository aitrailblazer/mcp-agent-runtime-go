package router

type PromptRequest struct {
	Prompt  string                 `json:"prompt"`
	Tools   []string               `json:"tools"`
	Context map[string]interface{} `json:"context"`
}
