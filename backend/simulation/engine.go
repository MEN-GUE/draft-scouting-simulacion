package simulation

import (
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/generators"
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/models"
)

// Campamento mantiene el estado global y las métricas estocásticas del draft de visorías.
// Esta estructura actúa como el entorno unificado donde ocurren todos los eventos[cite: 17].
type Campamento struct {
	JugadoresEvaluados int               // Variable X de la Binomial Negativa (total de ensayos antes de alcanzar r éxitos)[cite: 17].
	TiempoTotal        float64           // Reloj global asíncrono (acumulación continua de tiempos exponenciales)[cite: 17].
	JugadoresElite     []*models.Jugador // Slice dinámico que almacena únicamente los "éxitos" encontrados (los futuros canteranos)[cite: 17].
	ObjetivoElite      int               // Criterio de parada r para la Binomial Negativa[cite: 17].
	TasaLlegadas       float64           // Parámetro lambda para el proceso de Poisson / Tiempos Exponenciales[cite: 17].
	MediaAtributos     float64           // Parámetro mu (μ) para la distribución Normal[cite: 17].
	DesvAtributos      float64           // Parámetro sigma (σ) para la distribución Normal[cite: 17].
	HistorialPromedios []float64         // Colección estadística para registrar la valoración general de TODO el volumen de aspirantes[cite: 17].
}

// NewCampamento inicializa el entorno de simulación en memoria.
// Se pre-asigna capacidad a los slices para evitar el costo de re-localización en RAM 
// durante los ciclos estocásticos intensivos.
func NewCampamento(objetivo int, tasaLlegadas, mu, sigma float64) *Campamento {
	return &Campamento{
		JugadoresEvaluados: 0,
		TiempoTotal:        0.0,
		JugadoresElite:     make([]*models.Jugador, 0, objetivo),
		ObjetivoElite:      objetivo,
		TasaLlegadas:       tasaLlegadas,
		MediaAtributos:     mu,
		DesvAtributos:      sigma,
		// Se pre-asigna un buffer inicial (ej. 10,000) basado en una estimación teórica 
		// para optimizar el rendimiento de la escritura en el heap[cite: 17].
		HistorialPromedios: make([]float64, 0, 10000), 
	}
}

// Ejecutar orquesta el bucle de eventos discretos que simula el campamento de fútbol.
// Al inyectar GeneradorNormal como dependencia, logramos aislar la evaluación del modelo deportivo 
// de la complejidad del algoritmo generador (Transformada Inversa o Box-Muller),
// permitiendo una comparación algorítmica impecable[cite: 17].
func (c *Campamento) Ejecutar(generadorAtributos models.GeneradorNormal) {
	// BUCLE BINOMIAL NEGATIVA: 
	// Repetimos la evaluación del jugador de forma independiente hasta que el 
	// slice JugadoresElite alcance la longitud definida en ObjetivoElite (r)[cite: 17].
	for len(c.JugadoresElite) < c.ObjetivoElite {
		
		// 1. Modelado Continuo de Llegadas (Proceso de Markov / Poisson)
		// Generamos el delta de tiempo exacto que transcurrió desde la llegada del jugador anterior[cite: 17].
		deltaTiempo := generators.ExponencialTransformadaInversa(c.TasaLlegadas)
		
		// El tiempo global avanza asíncronamente sumando variables aleatorias independientes[cite: 17].
		c.TiempoTotal += deltaTiempo
		
		// 2. Conteo de Ensayos (Variable X)
		c.JugadoresEvaluados++
		
		// 3. Generación Física, Técnica y Mental (Distribución Normal)
		// Se invoca al constructor, el cual ejecutará el algoritmo generador 8 veces por aspirante[cite: 17].
		aspirante := models.NewJugador(
			c.JugadoresEvaluados, 
			c.TiempoTotal, 
			generadorAtributos, 
			c.MediaAtributos, 
			c.DesvAtributos,
		)
		
		// Acumulamos el dato empírico para la validación estadística y posterior exportación CSV[cite: 17].
		c.HistorialPromedios = append(c.HistorialPromedios, aspirante.PromedioGeneral)

		// 4. Verificación del Criterio de Éxito (Ensayo de Bernoulli)
		// Si el jugador posee el OVR requerido o atributos especializados > 90, se firma para la academia[cite: 17].
		if aspirante.EsElite {
			c.JugadoresElite = append(c.JugadoresElite, aspirante)
		}
		
		// Garbage Collection (GC):
		// Los punteros de los jugadores que resultaron en `EsElite == false` quedan sin referencia.
		// El GC de Go reciclará esta memoria eficientemente, lo cual es crítico para no saturar 
		// la memoria del host durante simulaciones masivas de gran escala[cite: 17].
	}
}