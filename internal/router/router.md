# `router.InvokeHandler` — Why Structured Request Decoding Is Recommended

The current implementation of `/invoke` in `mcp-agent-runtime-go` decodes incoming JSON into a `PromptRequest` struct and returns a structured response.

This approach aligns with modern agentic API practices and is recommended for the following reasons:

---

## ✅ 1. Model Context Protocol (MCP) Compliance

The [Model Context Protocol (MCP)](https://modelcontextprotocol.io/specification/2025-03-26) defines:

- A canonical `PromptRequest` structure
- Tool-aware execution logic
- Schema-bound request/response lifecycles

Your `InvokeHandler` using:
```go
type PromptRequest struct {
  Prompt  string                 `json:"prompt"`
  Tools   []string               `json:"tools"`
  Context map[string]interface{} `json:"context"`
}
```
...ensures your runtime is protocol-compliant.

---

## ✅ 2. API Standards: OpenAPI & REST

OpenAPI design guidelines encourage:
- Struct-based JSON parsing
- Clear field contracts (typed)
- JSON responses with traceable status and content

Structured decoding enables future OpenAPI compatibility, versioning, and validation layers.

---

## ✅ 3. Go Ecosystem Best Practices

Go’s standard library encourages:
```go
json.NewDecoder(r.Body).Decode(&struct)
```
This avoids unsafe reflection, enables IDE support, and allows for easy unit testing with `httptest`.

---

## ✅ 4. Observability and Logging

Decoded structs allow:
- Logging individual fields (e.g., `req.Prompt`)
- Tracing user inputs
- Tagging request IDs or session tokens per invocation

---

## ✅ 5. Agentic Routing and Lifecycle Hooks

Once parsed, `PromptRequest` can be:
- Routed to an agent or tool
- Passed into a session
- Transformed by a `PromptOptimizer`

This scaffolds the full MCP lifecycle.

---

## ✅ 6. Testability with `mcpcurl` and Clients

Struct decoding enables:
- Compatibility with `mcpcurl --input-file prompt.json`
- Test-driven workflows with `go test` and `httptest.NewRecorder`

---

## Summary

✅ Decoding into a `PromptRequest` struct is the minimum viable standard for:
- Safety
- Compatibility
- Future extension
- Agentic design alignment

