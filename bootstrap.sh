#!/bin/bash

echo "📁 Creating folders..."
mkdir -p cmd/server
mkdir -p internal/{agent,config,executor,middleware,registry,router,tools,tracing,validator}
mkdir -p api schema test

echo "📝 Creating placeholder files..."
touch internal/{agent,config,executor,middleware,registry,router,tools,tracing,validator}/.keep

cat <<EOF > cmd/server/main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("✅ MCP Agent Runtime starting on :8080")
    http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("MCP Agent Runtime OK"))
    }))
}
EOF

echo "✅ Scaffold complete. Run with: go run ./cmd/server"
