# `test/prompt.json` — Sample Request for /invoke

This file contains a valid `PromptRequest` payload that can be used to test the MCP Agent Runtime via the `/invoke` endpoint.

---

## 📄 File: `test/prompt.json`

```json
{
  "prompt": "Hello, MCP Agent Runtime!",
  "tools": [],
  "context": {}
}
```

---

## 🧪 How to Use

Start your server:
```bash
go run ./cmd/server
```

Run the test client:
```bash
go run ./cmd/testclient   --url http://localhost:8080/invoke   --input-file ./test/prompt.json
```

---

## ✅ What It Tests

- MCP JSON decoding via `PromptRequest`
- Route wiring of `/invoke`
- Response formatting
- Compatibility with `mcpcurl` or any compliant MCP client

---

## 🌱 Next Steps

You can duplicate or modify this file to test:
- Tool execution logic
- Prompt optimization paths
- Context propagation

