package models

import (
	"math"
)

// Atributos agrupa las características del aspirante en una escala de 1 a 99.
type Atributos struct {
	// Físicos
	Velocidad   int
	Resistencia int
	Fuerza      int

	// Técnicos
	ControlBalon int
	PaseCorto    int
	Definicion   int

	// Mentales
	Vision     int
	Compostura int
}

// Jugador representa una entidad dentro del campamento de visorías.
type Jugador struct {
	ID              int
	TiempoLlegada   float64   // Generado vía distribución Exponencial
	Atributos       Atributos // Generados vía distribución Normal
	PromedioGeneral float64   // Valoración global (OVR)
	EsElite         bool      // Determina el "éxito" para la Binomial Negativa
}

// clamp recorta el valor generado por la campana de Gauss para asegurar
// que se mantenga en los límites realistas del simulador deportivo [1, 99].
func clamp(val float64) int {
	v := int(math.Round(val))
	if v < 1 {
		return 1
	}
	if v > 99 {
		return 99
	}
	return v
}