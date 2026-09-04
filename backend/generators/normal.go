package generators

import (
	"math"
	"math/rand"
)

// NormalTransformadaInversa genera un valor con Distribución Normal usando una aproximación numérica.
// Se emplea la aproximación racional de Abramowitz y Stegun para la inversa de la CDF Normal.
func NormalTransformadaInversa(mu, sigma float64) float64 {
	u := rand.Float64()
	
	// Constantes de la aproximación racional
	c0, c1, c2 := 2.515517, 0.802853, 0.010328
	d1, d2, d3 := 1.432788, 0.189269, 0.001308

	var z float64

	// La aproximación divide la campana en dos mitades (u < 0.5 y u >= 0.5)
	if u < 0.5 {
		t := math.Sqrt(-2.0 * math.Log(u))
		z = -(t - ((c0 + c1*t + c2*t*t) / (1.0 + d1*t + d2*t*t + d3*t*t*t)))
	} else {
		t := math.Sqrt(-2.0 * math.Log(1.0-u))
		z = t - ((c0 + c1*t + c2*t*t) / (1.0 + d1*t + d2*t*t + d3*t*t*t))
	}

	// Transformamos el valor Z estándar a la media y desviación requerida (mu + Z*sigma)
	return mu + z*sigma
}

// NormalPolarBoxMuller genera un valor con Distribución Normal usando el Método Polar (Box-Muller).
// Utiliza funciones trigonométricas sobre dos variables uniformes.
func NormalPolarBoxMuller(mu, sigma float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()

	// Protección contra math.Log(0)
	if u1 == 0 {
		u1 = 1e-15
	}

	// Box-Muller genera dos valores Z independientes, pero para optimizar retornaremos solo Z0
	z0 := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
	
	return mu + z0*sigma
}