package simulation

import (
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/generators"
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/models"
)

// Campamento mantiene el estado global y las métricas estocásticas del draft de visorías.
// Al exportar los campos (con mayúscula inicial), el paquete encoding/json podrá
// serializar la estructura completa automáticamente para enviarla como respuesta de la API.
type Campamento struct {
	JugadoresEvaluados int               `json:"jugadoresEvaluados"` // Variable X de la Binomial Negativa
	TiempoTotal        float64           `json:"tiempoTotal"`        // Reloj global asíncrono
	JugadoresElite     []*models.Jugador `json:"jugadoresElite"`     // Slice con los "éxitos" encontrados
	TodosLosJugadores  []*models.Jugador `json:"todosLosJugadores"`  // Almacena el objeto completo de TODOS los aspirantes evaluados
	ObjetivoElite      int               `json:"objetivoElite"`      // Criterio de parada r
	TasaLlegadas       float64           `json:"tasaLlegadas"`       // Parámetro lambda (Poisson/Exponencial)
	MediaAtributos     float64           `json:"mediaAtributos"`     // Parámetro mu (Distribución Normal)
	DesvAtributos      float64           `json:"desvAtributos"`      // Parámetro sigma (Distribución Normal)
}

// NewCampamento inicializa el entorno de simulación en memoria.
func NewCampamento(objetivo int, tasaLlegadas, mu, sigma float64) *Campamento {
	return &Campamento{
		JugadoresEvaluados: 0,
		TiempoTotal:        0.0,
		JugadoresElite:     make([]*models.Jugador, 0, objetivo),
		// Pre-asignamos un buffer inicial para optimizar el rendimiento de escritura en RAM
		TodosLosJugadores: make([]*models.Jugador, 0, 2000),
		ObjetivoElite:     objetivo,
		TasaLlegadas:      tasaLlegadas,
		MediaAtributos:    mu,
		DesvAtributos:     sigma,
	}
}

// Ejecutar orquesta el bucle de eventos discretos que simula el campamento de fútbol.
func (c *Campamento) Ejecutar(generadorAtributos models.GeneradorNormal) {
	// BUCLE BINOMIAL NEGATIVA:
	// Repetimos la evaluación del jugador de forma independiente hasta que el
	// slice JugadoresElite alcance la longitud definida en ObjetivoElite (r).
	for len(c.JugadoresElite) < c.ObjetivoElite {

		// 1. Modelado Continuo de Llegadas (Proceso de Markov / Poisson)
		deltaTiempo := generators.ExponencialTransformadaInversa(c.TasaLlegadas)
		c.TiempoTotal += deltaTiempo

		// 2. Conteo de Ensayos (Variable X)
		c.JugadoresEvaluados++

		// 3. Generación Física, Técnica y Mental (Distribución Normal)
		aspirante := models.NewJugador(
			c.JugadoresEvaluados,
			c.TiempoTotal,
			generadorAtributos,
			c.MediaAtributos,
			c.DesvAtributos,
		)

		// Almacenamos al aspirante para enviar el registro completo al cliente (Frontend)
		c.TodosLosJugadores = append(c.TodosLosJugadores, aspirante)

		// 4. Verificación del Criterio de Éxito (Ensayo de Bernoulli)
		if aspirante.EsElite {
			c.JugadoresElite = append(c.JugadoresElite, aspirante)
		}
	}
}
