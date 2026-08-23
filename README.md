# SchemaBridge

A bridge between JSON &lt;-> GraphQL &lt;-> gRPC

## SchemaBridge is a contract-translation and schema-intelligence layer that bridges different API contracts by observing existing API data, inferring the missing schema, explaining the evidence behind its decisions, identifying incompatibilities, and allowing engineers to generate or migrate that contract into another API technology.

               EXISTING REALITY
                       │
             JSON / REST / OpenAPI
                       │
                       ▼
              ┌─────────────────┐
              │   SchemaBridge  │
              │                 │
              │ Parse           │
              │ Infer           │
              │ Normalize       │
              │ Compare         │
              │ Validate        │
              │ Explain         │
              └────────┬────────┘
                       │
              ┌────────┼─────────┐
              ▼        ▼         ▼
          GraphQL   Protobuf   OpenAPI
                       │
                       ▼
                    Migration

We'll be building

1. CLI Tool
2. VSCode Extension
3. Chromium Extension
4. Webapp (PWA)
