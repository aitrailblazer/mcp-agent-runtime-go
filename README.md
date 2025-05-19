# mcp-agent-runtime-go

A modular Go runtime for the **Model Context Protocol (MCP)** — designed for intelligent agent coordination, model selection, prompt optimization, and lifecycle-aware execution.

This SDK helps you build systems that:

* Coordinate multiple models (e.g., GPT-4o, Claude, Gemini)
* Dynamically generate and execute agents
* Route tasks based on trust, capability, or intent
* Manage prompt formatting and lifecycle control

---

## 🧠 Why this exists

Modern AI infrastructure is **agentic**: prompts now flow through models, memory, tools, and reasoning layers. MCP offers a universal protocol for managing that flow.

This Go implementation provides:

* **Concurrency-safe, production-ready execution**
* **Symbolically annotated agent flows**
* **Pluggable model and tool support**
* **Session lifecycle and halting controls**

---

## 🔧 Key Features

* `SessionCoordinator` for request routing, lifecycle hooks
* `ModelRouter` for dynamic model selection
* `AgentGenerator` for building runtime agents on demand
* `PromptOptimizerStep` to transform user input into optimized prompts
* `HaltingPolicy` to bound execution
* `DescriptorHash()` to lock and validate execution configuration

---

## 📦 Installation

```bash
go get github.com/modelcontextprotocol/mcp-agent-runtime-go
```

---

## 📁 Project Structure

```
/core/               # Protocol definitions, session types
/model/              # ModelSpec, Registry, Router logic
/agentgen/           # Dynamic agent generation
/promptopt/          # Prompt optimization steps and templates
/ext/                # Optional extensions (e.g., memory, auth)
/transport/          # HTTP/SSE/WebSocket integrations
roadmap.md           # Strategic plan for contributors
README.md            # This file
```

---

## 🗺️ Roadmap Overview

See [`roadmap.md`](./roadmap.md) for full scoring and implementation priority.

| Core Module         | Priority | Description                               |
| ------------------- | -------- | ----------------------------------------- |
| SessionLifecycle    | 10       | Session entry points and lifecycle events |
| ModelRouter         | 10       | Multi-model dispatch logic                |
| AgentGenerator      | 10       | Runtime agent composition and execution   |
| HaltingPolicy       | 9        | Safe runtime control                      |
| PromptOptimizerStep | 9        | Transform intent into optimized prompts   |
| DescriptorHash()    | 10       | Lock agent and routing structure          |

---

## 🤝 Contributing

This repo is open for discussion and design iteration.

* Start with [GitHub Discussion #364](https://github.com/orgs/modelcontextprotocol/discussions/364)
* See [roadmap.md](./roadmap.md) for priorities
* Use structured PRs based on core module folders

---

## 📜 License

MIT License — open use encouraged with attribution.
