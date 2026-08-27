package ollama

import (
	"fmt"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// refactoringGuruBaseURL is the Spanish refactoring.guru catalog of design
// patterns. For design-patterns challenges, the Target string (e.g.
// "facade", "factory-method") matches the URL segment refactoring.guru uses
// for that pattern, so we can point the model at the canonical description
// of whichever pattern it's working with as grounding context.
const refactoringGuruBaseURL = "https://refactoring.guru/es/design-patterns/"

// patternReferenceURL builds the refactoring.guru reference URL for a design pattern target.
func patternReferenceURL(target string) string {
	return refactoringGuruBaseURL + target
}

// typeGuidance returns type-specific instructions describing what solving a
// challenge of this type actually involves — since "target" means something
// different depending on the subject area.
func typeGuidance(t domain.ChallengeType) string {
	switch t {
	case domain.TerraformChallengeType:
		return "El reto debe resolverse escribiendo código Terraform (HCL) idiomático. \"target\" es el concepto de Terraform que se está practicando (por ejemplo: módulos, manejo de estado remoto, variables y outputs, workspaces, providers múltiples)."
	case domain.DataAnalyticsChallengeType:
		return "El reto debe resolverse escribiendo código Go que procese o analice datos. \"target\" es la técnica de análisis de datos que se está practicando (por ejemplo: limpieza de datos, agregaciones, detección de outliers, series de tiempo)."
	case domain.DesignPatternsChallengeType:
		return "El reto debe resolverse implementando, en Go idiomático (interfaces, structs y composición — Go no tiene clases ni herencia), el patrón de diseño indicado en \"target\"."
	default:
		return ""
	}
}

// challengeSystemPrompt instructs the model to design a practical exercise
// and answer with a strict JSON object.
const challengeSystemPrompt = `Eres un instructor senior que diseña retos prácticos de programación para que otras personas los resuelvan y aprendan.

Requisitos de la descripción del reto:
- Plantea un escenario realista y concreto (no una definición abstracta del tema).
- Enumera los requisitos funcionales que la solución debe cumplir.
- Deja explícito qué debe exponer o lograr la solución, sin dictar la implementación exacta — quien resuelva el reto debe seguir tomando decisiones de diseño.
- No incluyas código de la solución ni pseudocódigo.

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional, sin explicaciones y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"name": "<nombre corto y descriptivo del reto>", "description": "<descripción completa del reto, puede tener varios párrafos>"}`

// challengeUserPrompt builds the user-turn content for a challenge generation request.
func challengeUserPrompt(req ports.ChallengeGenerationRequest) string {
	prompt := fmt.Sprintf(
		"Genera un reto sobre el tema %q, dificultad %s, de tipo %s, con target %q.\n%s",
		req.Topic,
		req.Difficulty,
		req.Type,
		req.Target,
		typeGuidance(req.Type),
	)

	if req.Type == domain.DesignPatternsChallengeType {
		prompt += fmt.Sprintf("\nDocumentación de referencia del patrón (úsala para asegurar que la estructura y el propósito del reto sean fieles al patrón, no la cites textualmente): %s", patternReferenceURL(req.Target))
	}

	return prompt
}

// clueSystemPrompt instructs the model to give a progressive hint and
// answer with a strict JSON object.
const clueSystemPrompt = `Eres un asistente que da pistas progresivas para retos de programación práctica.

Reglas:
- La pista debe ayudar a avanzar en la solución sin revelarla por completo ni incluir código.
- Cada pista nueva debe ser más específica que las anteriores, en función de cuántas pistas ya se entregaron: la primera pista debe ser conceptual (qué problema hay que resolver), y las siguientes deben acercarse cada vez más a la estructura concreta de la solución.

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"clue": "<texto de la pista, una o dos oraciones>"}`

// clueUserPrompt builds the user-turn content for a clue generation request.
func clueUserPrompt(challenge *domain.Challenge) string {
	prompt := fmt.Sprintf(
		"Reto: %s\nDescripción: %s\nTipo: %s\nTarget: %s\n%s\nPistas ya entregadas: %d\n\nGenera la pista número %d, más específica que las anteriores.",
		challenge.Name,
		challenge.Description,
		challenge.Type,
		challenge.Target,
		typeGuidance(challenge.Type),
		len(challenge.Clues),
		len(challenge.Clues)+1,
	)

	if challenge.Type == domain.DesignPatternsChallengeType {
		prompt += fmt.Sprintf("\nDocumentación de referencia del patrón: %s", patternReferenceURL(challenge.Target))
	}

	return prompt
}

// feedbackSystemPrompt instructs the model to review a submitted solution
// and answer with a strict JSON object.
const feedbackSystemPrompt = `Eres un revisor de código senior que evalúa soluciones entregadas para retos de programación práctica.

Vas a recibir la descripción de un reto, su tipo y target, y el código que un usuario entregó como solución. Evalúa:
- Si la solución cumple el propósito del target indicado (no basta con que el código funcione, debe abordar específicamente lo que el reto pedía).
- Si el código es idiomático para la tecnología del reto y sigue buenas prácticas.
- Si cumple los requisitos funcionales del reto.

Responde EXCLUSIVAMENTE con un objeto JSON válido, sin texto adicional y sin bloques de markdown. El JSON debe tener exactamente esta forma:
{"score": <entero de 0 a 100>, "summary": "<resumen de una o dos oraciones sobre la calidad de la solución>", "suggestions": ["<sugerencia concreta de mejora>", "..."]}

Incluye al menos una sugerencia en la lista, incluso si la solución es correcta y completa.`

// feedbackUserPrompt builds the user-turn content for a solution evaluation request.
func feedbackUserPrompt(challenge *domain.Challenge, solutionCode string) string {
	prompt := fmt.Sprintf(
		"Reto: %s\nDescripción: %s\nTipo: %s\nTarget: %s\n%s",
		challenge.Name,
		challenge.Description,
		challenge.Type,
		challenge.Target,
		typeGuidance(challenge.Type),
	)

	if challenge.Type == domain.DesignPatternsChallengeType {
		prompt += fmt.Sprintf("\nDocumentación de referencia del patrón: %s", patternReferenceURL(challenge.Target))
	}

	prompt += fmt.Sprintf("\n\nSolución entregada:\n```\n%s\n```", solutionCode)

	return prompt
}
