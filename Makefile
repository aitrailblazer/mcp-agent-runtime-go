# Makefile for mcp-agent-runtime-go

# Help menu
help:
	@echo ""
	@echo "📘 MCP Agent Runtime Makefile Commands"
	@echo "---------------------------------------"
	@echo "make run-server-bg    # ✅ Start the MCP server in the background (recommended)"
	@echo "make stop-server      # 🛑 Stop the background MCP server"
	@echo "make test-invoke      # 🧪 Run HTTP testclient against /invoke using prompt.json"
	@echo "make build-mcpcurl    # 🔧 Build the mcpcurl CLI tool (STDIO)"
	@echo "make test-mcpcurl     # 🧪 Run mcpcurl test using STDIO server"
	@echo "make run-server       # 🐞 Run the MCP server in foreground"
	@echo "make help             # 📖 Show this help message"
	@echo ""

# 🐞 Run the MCP server in the foreground (debugging/manual)
run-server:
	go run ./cmd/server

# ✅ Recommended: Run the server in the background with nohup
run-server-bg:
	@echo "🚀 Starting server in background on port 8080"
	nohup go run ./cmd/server > server.log 2>&1 & echo $$! > server.pid

# 🛑 Stop the background server
stop-server:
	@echo "🛑 Stopping server..."
	-@kill `cat server.pid` 2>/dev/null || true
	@rm -f server.pid

# 🧪 Run HTTP test client using prompt.json
test-invoke:
	@echo "📤 Sending test request to /invoke (HTTP)..."
	go run ./cmd/testclient \
		--url http://localhost:8080/invoke \
		--input-file ./test/prompt.json

# 🔧 Build the mcpcurl CLI tool inside tools/mcpcurl
build-mcpcurl:
	@echo "🔧 Building mcpcurl from tools/mcpcurl..."
	cd tools/mcpcurl && go mod tidy && go build -o mcpcurl

# 🧪 Run test using mcpcurl in STDIO mode
test-mcpcurl:
	@echo "🧪 Testing using mcpcurl (stdio)..."
	./tools/mcpcurl/mcpcurl \
		--stdio-server-cmd "go run ./cmd/stdio" \
		--input-file ./test/prompt.json
