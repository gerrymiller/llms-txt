package parser

func GenerateManifest(path string) (string, error) {
	// TODO: walk directory, parse OpenAPI/Markdown, extract metadata
	// For now, return a placeholder
	return `# llms.txt - Agent Discovery Manifest

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
`, nil
}
