# llms.txt

> **A lightweight, agent-readable manifest format for surfacing your content, APIs, and documentation to LLM-based agents.**

`llms.txt` is a proposed open standard that acts as a machine-readable index for AI agents—similar in spirit to `robots.txt`, but designed for the age of autonomous AI. It complements protocols like MCP (Model Context Protocol) and A2A (Agent-to-Agent), providing a bootstrap layer for AI agent discovery and interoperability.

---

## 🚀 What It Does

- Enables AI agents to discover available APIs, data sources, and content schemas
- Offers a static, easy-to-implement discovery surface for LLM-based agents
- Works alongside dynamic protocols like MCP and A2A to improve agent UX
- Helps developers build "agent-ready" platforms and tools

## 📂 Example `llms.txt`

```yaml
# llms.txt
openapi: /api/openapi.json
mcp_endpoint: https://api.example.com/mcp
a2a_manifest: /a2a/capabilities.json
content_schema: /docs/ontology.json
semantic_index: /vectors/index.json
capabilities:
  - data.read
  - api.call
  - doc.search
preferred_roles:
  - "integration_assistant"
  - "compliance_checker"
```

## 🛠 Tooling

This repo includes a CLI tool (`llmsgen`) written in Go to automatically generate a valid `llms.txt` file from your existing documentation and OpenAPI specs.

```bash
$ llmsgen ./docs
```

## 📁 Project Structure

```plaintext
llms-txt/
├── cmd/llmsgen/               # CLI tool
├── internal/parser/           # Markdown/OpenAPI parser logic
├── spec/
│   ├── spec.md                # Draft spec for llms.txt
│   ├── llms.schema.json       # JSON Schema for validation (optional)
│   └── examples/              # Real-world examples
├── README.md
├── LICENSE
├── .gitignore
├── go.mod
```

## 👥 Contributors

- [@gerrymiller](https://github.com/gerrymiller) — Author and initiator of the `llms.txt` standard

## 🧪 Next Steps

- [ ] Finalize spec (see `spec/spec.md`)
- [ ] Populate examples (Redox, Particle Health)
- [ ] Add OpenAPI + Markdown parser
- [ ] Blog post: *Introducing `llms.txt` — The Agentic Web’s Missing Bootstrap Layer*

---

> Want to help define the agentic discovery layer of the internet? Jump in, fork it, and contribute.

---
