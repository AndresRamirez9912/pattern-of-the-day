# Pattern of the Day

## What is this?

A personal project to practice and evaluate skills during study, by generating
practical challenges tailored to a topic and difficulty level. It currently
covers three subject areas: **Terraform**, **design patterns**, and **data
analytics**. A locally-running LLM (via [Ollama](https://ollama.com)) writes
the challenge, generates progressive clues on request, and evaluates
submitted solutions with feedback.

Beyond that, this project is also a hands-on exercise in **Clean
Architecture / Ports & Adapters (hexagonal architecture)** — the design
choices here favor learning and enforcing that separation as much as they
favor shipping the feature itself.

## Architecture

The project follows a Ports & Adapters (hexagonal) architecture: dependencies
always point inward, toward the domain. Adapters depend on ports and use
cases; nothing in `domain` or `application` knows an adapter exists.

```
internal/
├── domain/         Core business types and rules: Challenge, Attempt, Clue,
│                   Feedback, User, domain errors. Depends on nothing else.
├── ports/          Interfaces the application layer needs fulfilled —
│                   repositories, the LLM provider, the file writer, the logger.
├── application/    Use cases, one package per resource (challenge, clue,
│                   attempt, feedback, user). Each orchestrates ports to
│                   fulfill one business operation, unaware of which adapter
│                   sits behind a port.
├── adapters/
│   ├── inbound/    Driving adapters — translate an external trigger into a
│   │   │           call on a use case, never containing business logic.
│   │   ├── cli/    Cobra commands.
│   │   └── rest/   HTTP handlers + router (net/http, no external framework).
│   └── outbound/   Driven adapters — fulfill a port with real infrastructure.
│       ├── persistence/   SQLite via sqlc.
│       ├── llm/ollama/    Ollama-backed LLMProvider implementation.
│       └── filesystem/    Writes generated content as markdown files
│                          (challenge.md, clues.md, attempt-N-feedback.md).
└── app/            Composition root — opens the DB connection, wires every
                    adapter into its use case (services.go), and exposes
                    App{Services, Logger, Ctx} to whichever inbound adapter
                    is running (app.go, config/).

cmd/pattern-of-the-day/   Binary entry point: main.go + the cobra command tree
                          (root, database migrations, start-server).
migrations/               goose SQL migrations.
dockerfile                Multi-stage build producing a ~13MB scratch image.
```

The rule that keeps this consistent: a port is a technology-agnostic
interface owned by whichever layer needs the capability (usually
`application`). CLI and REST are just two different inbound adapters calling
the exact same use cases — neither one duplicates business logic of its own.

## Running it

There are two ways to interact with the app, both driving the same use cases
underneath.

### 1. CLI

```sh
go build -o patternd ./cmd/pattern-of-the-day
./patternd database up

./patternd use-cases users create <username> <email>
./patternd use-cases challenge create <username> <difficulty> [type] \
  --topic "..." --target "..." --out ./my-challenges
./patternd use-cases clue create <challenge-id> --out ./my-challenges
./patternd use-cases feedback create <attempt-id> <solution-path> --out ./my-challenges
```

### 2. REST

```sh
./patternd start-server
```

Listens on `app.rest_addr` from `config.yaml` (default `0.0.0.0:8080`).

| Method | Path                       | Notes                                                    |
|--------|----------------------------|-----------------------------------------------------------|
| POST   | `/users`                   | |
| GET    | `/users/{user_id}`         | |
| POST   | `/challenges`              | |
| GET    | `/challenges/{challenge_id}` | Returns the challenge with its clues, attempts, and feedbacks |
| POST   | `/clues`                   | |
| POST   | `/feedback`                | `multipart/form-data`: `attempt_id` field + one or more files under `files` |
| GET    | `/health`                  | Liveness check |

Both ways of running the app need `config.yaml` at the repo root, and a
running Ollama instance reachable at `llm.base_url` for anything that
generates or evaluates content (challenge, clue, feedback).
