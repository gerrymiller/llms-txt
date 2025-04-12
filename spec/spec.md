# llms.txt Specification (v0.1 Draft)

`llms.txt` is a proposed open specification to support agentic discovery on the modern web. It provides a lightweight, machine-readable manifest file—placed at the root of a domain—that informs AI agents about the APIs, tools, documentation, and semantics available for integration or interaction.

Inspired by `robots.txt`, `manifest.json`, and `sitemap.xml`, this file allows both general-purpose and task-specific AI agents to quickly understand how to interact with a system.

---

## 🧩 File Format

- **Type:** Plaintext or YAML
- **Location:** Placed at the root of the domain (e.g., `https://example.com/llms.txt`)
- **Encoding:** UTF-8
- **Extension:** `.txt` (though `.yaml` or `.json` may be used in future versions)

---

## 🔖 Top-Level Fields

| Field             | Type    | Description                                                                 |
|------------------|---------|-----------------------------------------------------------------------------|
| `openapi`         | string  | URL to OpenAPI spec or Swagger file                                        |
| `mcp_endpoint`    | string  | URL to a Model Context Protocol (MCP) endpoint                             |
| `a2a_manifest`    | string  | URL to an Agent-to-Agent capabilities manifest                             |
| `content_schema`  | string  | URL to a schema file describing available document/content structure       |
| `semantic_index`  | string  | URL to a vector index or embedding endpoint                               |
| `capabilities`    | list    | List of high-level agent-callable actions supported                        |
| `preferred_roles` | list    | Agent roles best suited to interact with this resource                     |

---

## 📌 File Location

By convention, `llms.txt` should be placed at the root of the domain:

```plaintext
https://example.com/llms.txt
```

Alternatively, the file may be served from the `.well-known/` path:

```plaintext
https://example.com/.well-known/llms.txt
```

Supporting both (via redirect or alias) is encouraged for compatibility with future agent standards and crawlers.

---

## 📄 Example `llms.txt`

```yaml
# llms.txt - Agent Discovery Manifest
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

---

## 🚀 Design Goals

- Be trivially readable by both humans and machines
- Support progressive enhancement (i.e., work with just `openapi`, but grow as needed)
- Provide the **minimum viable agent onboarding layer**
- Serve as a pointer layer into more complex standards like MCP, A2A, OpenAPI, etc.

---

## 🔄 Versioning

This is **v0.1 (draft)**. Future versions may:

- Support JSON/YAML syntax formally
- Add `auth_methods`, `rate_limits`, or `agent_contact` fields
- Reference remote `.well-known` resources

---

## 📬 Feedback

Open issues or PRs at [https://github.com/gerrymiller/llms-txt](https://github.com/gerrymiller/llms-txt)

Want to contribute to the future of agentic interoperability? Join us.
