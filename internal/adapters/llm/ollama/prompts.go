package ollama

import (
	"fmt"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// refactoringGuruBaseURL is the Spanish refactoring.guru catalog of design
// patterns. Every domain.Pattern slug (e.g. "facade", "factory-method")
// matches the URL segment refactoring.guru uses for that pattern, so we can
// point the model at the canonical description of whichever pattern it's
// working with as grounding context.
const refactoringGuruBaseURL = "https://refactoring.guru/es/design-patterns/"

// patternReferenceURL builds the refactoring.guru reference URL for a pattern.
func patternReferenceURL(p domain.Pattern) string {
	return refactoringGuruBaseURL + string(p)
}

// challengeSystemPrompt instructs the model to design a Go-specific exercise
// and answer with a strict JSON object.
const challengeSystemPrompt = `Eres un instructor senior de Go (Golang) especializado en enseñar patrones de diseño mediante ejercicios prácticos de programación.

Tu tarea es diseñar un reto que solo se pueda resolver correctamente implementando, en Go idiomático, el patrón de diseño solicitado. Go no tiene clases ni herencia: el reto debe plantearse en términos de interfaces, structs y composición, nunca de clases abstractas o jerarquías de herencia.

Requisitos de la descripción del reto:
- Plantea un escenario realista y concreto (no una definición abstracta del patrón).
- Enumera los requisitos funcionales que la solución en Go debe cumplir.
- Deja explícito qué debe exponer la solución (p. ej. "define una interfaz que exponga el método X", "el struct Y debe poder intercambiarse por otra implementación sin cambiar el código que lo usa"), sin dictar la implementación exacta — quien resuelva el reto debe seguir tomando decisiones de diseño.
- No incluyas código de la solución ni pseudocódigo.
- El código de la solución debe poder escribirse usando solo la librería estándar de Go.

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional, sin explicaciones y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"name": "<nombre corto y descriptivo del reto>", "description": "<descripción completa del reto, puede tener varios párrafos>"}`

// challengeUserPrompt builds the user-turn content for a challenge generation request.
func challengeUserPrompt(req ports.ChallengeGenerationRequest) string {
	return fmt.Sprintf(
		"Genera un reto sobre el tema %q, dificultad %s, de tipo %s, enfocado en el patrón %s.\n"+
			"Documentación de referencia del patrón (úsala para asegurar que la estructura y el propósito del reto sean fieles al patrón, no la cites textualmente): %s",
		req.Topic,
		req.Difficulty,
		req.Type,
		req.Pattern,
		patternReferenceURL(req.Pattern),
	)
}

// clueSystemPrompt instructs the model to give a Go-specific, progressive
// hint and answer with a strict JSON object.
const clueSystemPrompt = `Eres un asistente que da pistas progresivas para retos de patrones de diseño que se resuelven en Go (Golang).

Reglas:
- La pista debe ayudar a avanzar en el diseño sin revelar la solución completa ni incluir código.
- Habla siempre en términos de construcciones de Go (interfaces, structs, métodos, composición, embedding). Nunca menciones "clases", "clases abstractas" ni "herencia": Go no las tiene y esos términos confundirían al usuario.
- Cada pista nueva debe ser más específica que las anteriores, en función de cuántas pistas ya se entregaron: la primera pista debe ser conceptual (qué problema resuelve el patrón aquí), y las siguientes deben acercarse cada vez más a la estructura concreta en Go (qué interfaz o composición se necesita).

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"clue": "<texto de la pista, una o dos oraciones>"}`

// clueUserPrompt builds the user-turn content for a clue generation request.
func clueUserPrompt(challenge *domain.Challenge) string {
	return fmt.Sprintf(
		"Reto: %s\nDescripción: %s\nPatrón objetivo: %s\nDocumentación de referencia del patrón: %s\nPistas ya entregadas: %d\n\nGenera la pista número %d, en términos de Go, más específica que las anteriores.",
		challenge.Name,
		challenge.Description,
		challenge.Pattern,
		patternReferenceURL(challenge.Pattern),
		len(challenge.Clues),
		len(challenge.Clues)+1,
	)
}

// feedbackSystemPrompt instructs the model to review a Go solution and
// answer with a strict JSON object.
const feedbackSystemPrompt = `Eres un revisor de código senior especializado en Go (Golang) y en patrones de diseño.

Vas a recibir la descripción de un reto, el patrón de diseño objetivo, y el código en Go que un usuario entregó como solución. Evalúa:
- Si el patrón de diseño fue implementado correctamente y cumple su propósito (no basta con que el código funcione, debe seguir la estructura del patrón).
- Si el código usa Go idiomático: interfaces pequeñas, composición en vez de intentar simular herencia, manejo explícito de errores, nombres idiomáticos, uso de punteros solo cuando corresponde.
- Si cumple los requisitos funcionales del reto.

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"score": <entero de 0 a 100>, "summary": "<resumen de una o dos oraciones sobre la calidad de la solución>", "suggestions": ["<sugerencia concreta de mejora>", "..."]}

Incluye al menos una sugerencia en la lista, incluso si la solución es correcta y completa.`

// feedbackUserPrompt builds the user-turn content for a solution evaluation request.
func feedbackUserPrompt(challenge *domain.Challenge, solutionCode string) string {
	return fmt.Sprintf(
		"Reto: %s\nDescripción: %s\nPatrón objetivo: %s\nDocumentación de referencia del patrón: %s\n\nSolución entregada en Go:\n```go\n%s\n```",
		challenge.Name,
		challenge.Description,
		challenge.Pattern,
		patternReferenceURL(challenge.Pattern),
		solutionCode,
	)
}
