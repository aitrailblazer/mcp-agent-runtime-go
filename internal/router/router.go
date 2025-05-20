// Package router provides HTTP routing functionality for the MCP Agent Runtime
package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aitrailblazer/mcp-agent-runtime-go/internal/types"
)

// HealthHandler handles the /healthz endpoint for health checks
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "MCP Agent Runtime OK")
}

// PingHandler handles the /ping endpoint for basic connectivity tests
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

// InvokeHandler handles the /invoke endpoint for agent operations
func InvokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := types.PromptResponse{
		Status:  "received",
		Prompt:  req.Prompt,
		Tools:   req.Tools,
		Context: req.Context,
	}

	json.NewEncoder(w).Encode(resp)
}

// SchemaHandler serves the MCP schema.json
func SchemaHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./schema/schema.json")
}
