# pattern-of-the-day

the-pattern-of-the-day/
│
├── cmd/
│ └── pattern-of-the-day/
│ └── main.go
│
├── internal/
│ │
│ ├── domain/
│ │ ├── challenge.go
│ │ ├── clue.go
│ │ ├── feedback.go
│ │ ├── pattern.go
│ │ └── errors.go
│ │
│ ├── application/
│ │ ├── challenge/
│ │ │ ├── create.go
│ │ │ ├── get.go
│ │ │ └── reveal.go
│ │ │
│ │ ├── clue/
│ │ │ └── generate.go
│ │ │
│ │ └── feedback/
│ │ └── generate.go
│ │
│ ├── ports/
│ │ ├── llm.go
│ │ ├── challenge_repository.go
│ │ ├── clue_repository.go
│ │ ├── feedback_repository.go
│ │ └── filesystem.go
│ │
│ ├── adapters/
│ │ │
│ │ ├── llm/
│ │ │ ├── openai/
│ │ │ ├── ollama/
│ │ │ └── anthropic/
│ │ │
│ │ ├── persistence/
│ │ │ └── mysql/
│ │ │ ├── challenge_repository.go
│ │ │ ├── clue_repository.go
│ │ │ └── feedback_repository.go
│ │ │
│ │ └── filesystem/
│ │ └── local.go
│ │
│ └── prompts/
│ ├── challenge_generation.txt
│ ├── clue_generation.txt
│ └── feedback_generation.txt
│
├── migrations/
│ ├── 001_create_challenges.sql
│ ├── 002_create_clues.sql
│ └── 003_create_feedback.sql
│
├── config/
│ └── config.go
│
├── challenges/
│ └── .gitkeep
│
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── .env.example
