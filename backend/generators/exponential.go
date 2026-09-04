package generators

import (
	"math"
	"math/rand"
)

// ExponencialTransformadaInversa genera un valor aleatorio de una distribución exponencial.
// Utiliza el método de la transformada inversa: x = -(1/lambda) * ln(1-U).
func ExponencialTransformadaInversa(lambda float64) float64 {
	// rand.Float64() genera un número seudoaleatorio U en el rango [0.0, 1.0)
	u := rand.Float64()
	
	// Evitamos el ln(0) en el caso teóricamente posible pero improbable de que 1.0 - u == 0
	if u == 1.0 {
		u = 0.9999999999
	}
	
	return -(1.0 / lambda) * math.Log(1.0-u)
}