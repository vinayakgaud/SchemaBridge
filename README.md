## What SchemaBridge Solves

API systems rarely evolve cleanly.

Teams have existing REST APIs, JSON payloads, OpenAPI specifications,
GraphQL schemas, protobuf definitions, and undocumented contracts.

SchemaBridge aims to provide a bridge between these realities.

It will:

- Observe existing API contracts and data
- Infer missing or ambiguous schema information
- Explain why an inference was made
- Assign confidence to inferred decisions
- Detect incompatible contract changes
- Translate contracts between API technologies
- Support controlled migration instead of manual rewrites

SchemaBridge is not intended to blindly convert one format into another.
It is designed to understand the source contract first, explain its reasoning,
and then produce a target contract with explicit confidence and evidence.

## Planned Interfaces

SchemaBridge will eventually be available through:

1. **CLI**
2. **VS Code Extension**
3. **Chromium Extension**
4. **Web Application / PWA**

All interfaces will use the same underlying SchemaBridge capabilities.

## Current Implementation

### CLI

The CLI currently provides:

```bash
sb analyze <input>
```

It can:

- Accept an input file
- Read the file through the core file-handling layer
- Analyze JSON structure
- Produce a recursive observation model

### JSON Structural Analysis

The current JSON observation layer understands:

- Primitive values
- Objects
- Nested objects
- Arrays
- Arrays of primitives
- Arrays of objects
- Mixed arrays
- Arrays containing objects with different shapes
- Null values
- Basic nullability observations
- Invalid JSON

The current pipeline is:

```text
Input File
    │
    ▼
CLI
    │
    ▼
Core File Handling
    │
    ▼
JSON Structure Analysis
    │
    ▼
Observation Model
```

The observation layer describes **what was observed**.

It does not yet attempt to intelligently infer the final API contract.

## Observation Model

The observation model represents JSON as a recursive structure:

```text
ObservationModel
└── Root
    ├── Type
    ├── Nullable
    ├── Fields
    │   └── FieldObservation
    │       ├── Name
    │       └── Node
    └── ArrayElement
        └── ObservationNode
```

This model will become the foundation for the future inference engine.

## Development

### Requirements

- Go
- Task
- Lefthook

### Common Commands

Run SchemaBridge in development mode:

```bash
task dev -- analyze ./examples/json/users.json
```

Format source code:

```bash
task format
```

Run unit tests:

```bash
task test:unit
```

Run all tests:

```bash
task test
```

Run static analysis:

```bash
task vet
```

Run all quality checks:

```bash
task check
```

Clean Go build and test caches:

```bash
task clean
```

`task clean` is intentionally explicit and is not executed automatically by
pre-commit hooks or CI.

## Quality Gates

Local commits use Lefthook to run:

```text
Format Check
     │
     ▼
  Go Vet
     │
     ▼
 Unit Tests
```

GitHub Actions independently validates changes through CI before they are
merged.

## Project Status

**Early development / architecture phase**

### Completed

- Repository structure
- Go module setup
- Cobra CLI foundation
- `analyze` command
- Core file handling
- Taskfile development workflow
- Lefthook pre-commit checks
- CI foundation
- Recursive JSON observation model
- JSON structural analysis
- Unit test foundation

### Next

- Improve observation semantics
- Build the inference engine
- Introduce confidence scoring
- Add evidence and explanation
- Define the canonical contract representation
- Begin contract translation
- Add GraphQL, Protobuf, and OpenAPI capabilities

## Security

See [SECURITY.md](SECURITY.md) for the security policy.

## License

SchemaBridge is licensed under the MIT License.
See [LICENSE](LICENSE).
