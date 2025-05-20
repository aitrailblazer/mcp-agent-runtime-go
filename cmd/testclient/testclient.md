# `testclient` — Local HTTP Test Driver for MCP Agent Runtime

The `testclient` is a minimal Go-based HTTP client for invoking the `/invoke` endpoint of your MCP server using local JSON files.

---

## ✅ Location

```bash
cmd/testclient/main.go
```

---

## 🚀 Usage

### 1. Start the MCP server

From project root:

```bash
go run ./cmd/server
```

---

### 2. Run the test client

```bash
go run ./cmd/testclient   --url http://localhost:8080/invoke   --input-file ./test/prompt.json
```

---

### 🧪 Example `prompt.json`

File: `test/prompt.json`

```json
{
  "prompt": "Hello, MCP Agent Runtime!",
  "tools": [],
  "context": {}
}
```

---

### ✅ Expected Output

```
Status: 200 OK

Response:
{
  "status": "received",
  "prompt": "Hello, MCP Agent Runtime!",
  "tools": [],
  "context": {}
}
```

---

## 💡 Why This Client Exists

This `testclient` is faster and simpler than `mcpcurl` for local dev:

- No stdio proxy needed
- Works directly with `POST /invoke`
- Accepts standard JSON
- Easy to debug and extend

Use it to validate handler behavior before integrating tool dispatch, schema validation, or executor chaining.

