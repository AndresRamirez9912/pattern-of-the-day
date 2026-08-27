package challenge

import (
	"math/rand"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// randomTopics is a pool of example topics used to fill in a challenge
// request's Topic when the caller doesn't provide one, so challenges keep
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

// randomTargetsByType is a pool of example targets per challenge type, used
// to fill in a challenge request's Target when the caller doesn't provide
// one. Unlike Type, Target has no fixed enum — these are just reasonable
// starting points, ordered roughly from basic to advanced within each type.
var randomTargetsByType = map[domain.ChallengeType][]string{
	domain.DesignPatternsChallengeType: {
		// Básico
		"singleton", "factory-method", "builder", "adapter", "decorator",
		"strategy", "observer", "template-method", "facade", "iterator",
		// Intermedio
		"abstract-factory", "prototype", "composite", "proxy", "command",
		"state", "chain-of-responsibility", "mediator", "memento", "null-object",
		// Avanzado
		"bridge", "flyweight", "visitor", "interpreter", "dependency-injection",
		"repository", "unit-of-work", "specification", "cqrs", "event-sourcing",
		"pipeline", "object-pool",
	},
	domain.TerraformChallengeType: {
		// Básico
		"recursos básicos y providers", "variables y outputs", "tipos de datos y validación de variables",
		"archivos de estado (state) local", "formato y organización de archivos .tf",
		"uso de tags y metadatos en recursos", "interpolación de strings",
		"dependencias implícitas y explícitas (depends_on)", "contadores con count",
		"estructuras condicionales básicas",
		// Intermedio
		"for_each y colecciones", "módulos reutilizables", "locals y expresiones",
		"data sources", "manejo de estado remoto (remote state)", "workspaces",
		"providers múltiples y alias", "funciones integradas (built-in functions)",
		"output entre módulos", "variables sensibles (sensitive)",
		// Avanzado
		"importación de recursos existentes (terraform import)", "lifecycle rules",
		"state locking y backends remotos", "módulos versionados y registries",
		"dynamic blocks", "provisioners", "gestión de drift (terraform plan -refresh-only)",
		"políticas con Sentinel/OPA", "migraciones de estado (terraform state mv/rm)",
		"multi-entorno con workspaces o directorios separados", "composición avanzada de módulos condicionales",
	},
	domain.DataAnalyticsChallengeType: {
		// Básico
		"lectura y escritura de archivos CSV", "tipos de datos y conversión de columnas",
		"filtrado de filas y columnas", "manejo de valores nulos o faltantes",
		"estadística descriptiva básica (media, mediana, moda)", "ordenamiento y ranking de datos",
		"renombrado y selección de columnas", "combinación de datasets (concatenación)",
		"detección y eliminación de duplicados", "visualización básica de datos",
		// Intermedio
		"agregaciones y agrupamiento (groupby)", "joins entre datasets",
		"pivoteo de tablas (pivot/melt)", "limpieza y normalización de texto",
		"detección de outliers", "series de tiempo y resampleo",
		"generación de reportes automatizados", "validación de esquemas de datos",
		"manejo de datos categóricos (encoding)", "muestreo de datos (sampling)",
		// Avanzado
		"pipelines de transformación de datos (ETL)", "procesamiento de datos en lotes",
		"análisis de correlación entre variables", "reducción de dimensionalidad",
		"detección de anomalías con métodos estadísticos", "feature engineering para modelos",
		"optimización de memoria en datasets grandes", "procesamiento de datos en streaming",
		"versionado y trazabilidad de datasets", "orquestación de pipelines de datos",
		"control de calidad de datos (data quality checks)", "integración con fuentes de datos externas",
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
