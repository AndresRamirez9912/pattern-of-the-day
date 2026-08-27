package usecases

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// randomTopics is a small pool of example topics used to fill in the
// --topic flag when the caller doesn't provide one, so challenges keep
// some variety without requiring a topic to be picked by hand every time.
var randomTopics = []string{
	"gestión de conexiones a base de datos",
	"sistema de notificaciones multicanal",
	"procesamiento de pagos",
	"generación de reportes en distintos formatos",
	"sistema de caché en memoria",
	"validación de formularios",
	"exportación de datos a distintos formatos de archivo",
	"sistema de logging configurable",
	"manejo de reintentos ante fallos de red",
	"integración con proveedores de autenticación externos",
}

// randomTargetsByType is a small pool of example targets per challenge
// type, used to fill in the --target flag when the caller doesn't provide
// one, so challenges keep some variety without requiring a target to be
// picked by hand every time. Unlike Type, Target has no fixed enum — these
// are just reasonable starting points.
var randomTargetsByType = map[domain.ChallengeType][]string{
	domain.DesignPatternsChallengeType: {
		"factory-method", "abstract-factory", "builder", "prototype", "singleton",
		"adapter", "bridge", "composite", "decorator", "facade", "flyweight", "proxy",
		"chain-of-responsibility", "command", "iterator", "mediator", "memento",
		"observer", "state", "strategy", "template-method", "visitor",
	},
	domain.TerraformChallengeType: {
		"módulos reutilizables", "manejo de estado remoto", "variables y outputs",
		"workspaces", "providers múltiples", "data sources", "lifecycle rules",
		"importación de recursos existentes",
	},
	domain.DataAnalyticsChallengeType: {
		"limpieza de datos", "agregaciones y agrupamiento", "detección de outliers",
		"series de tiempo", "joins entre datasets", "generación de reportes",
		"pipelines de transformación de datos",
	},
}

// randomTopic returns a random topic from the example pool.
func randomTopic() string {
	return randomTopics[rand.Intn(len(randomTopics))]
}

// randomTarget returns a random example target for the given challenge type.
func randomTarget(t domain.ChallengeType) string {
	targets := randomTargetsByType[t]
	if len(targets) == 0 {
		return ""
	}
	return targets[rand.Intn(len(targets))]
}

// randomChallengeType returns a random supported challenge type.
func randomChallengeType() domain.ChallengeType {
	types := domain.AllChallengeTypes()
	return types[rand.Intn(len(types))]
}

// parseDifficulty validates that s is one of the supported difficulties.
func parseDifficulty(s string) (domain.Difficulty, error) {
	difficulty := domain.Difficulty(strings.ToLower(strings.TrimSpace(s)))

	for _, valid := range domain.AllDifficulties() {
		if difficulty == valid {
			return difficulty, nil
		}
	}

	return "", fmt.Errorf("dificultad inválida %q, debe ser una de: easy, medium, hard", s)
}

// isValidChallengeType reports whether t is one of the supported challenge types.
func isValidChallengeType(t domain.ChallengeType) bool {
	for _, valid := range domain.AllChallengeTypes() {
		if t == valid {
			return true
		}
	}
	return false
}
