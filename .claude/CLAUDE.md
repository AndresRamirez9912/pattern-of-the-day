# Pattern of the Day

Personal Go project that generates practice challenges (Terraform,
design-patterns, data-analytics) via a local Ollama LLM, and evaluates
submitted solutions with feedback. It doubles as a study project for Clean
Architecture / Ports & Adapters — architectural consistency matters as much
as the feature itself.

## Layout

- `internal/domain/` — core types and business rules (Challenge, Attempt,
  Clue, Feedback, User). No dependency on anything else in the project.
- `internal/ports/` — interfaces the application layer needs fulfilled
  (repositories, LLM provider, file writer, logger).
- `internal/application/` — use cases, one package per resource (challenge,
  clue, attempt, feedback, user). Orchestrates ports; never imports an adapter.
- `internal/adapters/inbound/` — driving adapters: `cli/` (cobra commands),
  `rest/` (net/http handlers + router). Translate external input into a call
  on a use case; no business logic lives here.
- `internal/adapters/outbound/` — driven adapters: `persistence/` (SQLite via
  sqlc), `llm/ollama/`, `filesystem/` (markdown output).
- `internal/app/` — composition root: `app.go`/`services.go` wire every
  adapter into its use case and build `App{Services, Logger, Ctx}`.
- `cmd/pattern-of-the-day/` — binary entry point and cobra command tree.

## Behavior rules

1. Comments are 1-2 lines max. Only go longer when it's genuinely necessary
   to explain a non-obvious constraint — never to restate what the code does.
2. Don't collapse an error check into its declaration
   (`if err := f(); err != nil`). Always declare on one line, check on the
   next:
   ```go
   result, err := f()
   if err != nil {
       return err
   }
   ```
3. Every change must fit the existing ports & adapters layering: business
   logic belongs in `application`, not in an adapter; a new inbound adapter
   (CLI, REST) must call an existing use case rather than reimplement its
   logic; a new outbound need becomes a port in `internal/ports` before an
   adapter implements it.
