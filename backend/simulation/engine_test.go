package simulation

import (
	"testing"

	"github.com/MEN-GUE/draft-scouting-simulacion/backend/generators"
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/models"
)

// BenchmarkTransformadaInversa mide el rendimiento de la aproximación polinomial de Abramowitz y Stegun.
// Al ejecutar las pruebas (go test -bench=. -benchmem), el compilador aumentará iterativamente b.N 
// para obtener una medición precisa de los nanosegundos requeridos por operación (ns/op)[cite: 18].
func BenchmarkTransformadaInversa(b *testing.B) {
	// Reportar asignaciones de memoria es útil para verificar que nuestra generación en tiempo lineal O(n)
	// no sufra de fugas de memoria o heap allocations innecesarias en el bucle principal.
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Evaluamos el costo computacional de generar 8 variables aleatorias normales usando 
		// la aproximación polinomial racional[cite: 18].
		_ = models.NewJugador(
			i,
			0.0,
			generators.NormalTransformadaInversa,
			50.0, // Media OVR
			15.0, // Varianza / Desviación estándar
		)
	}
}

// BenchmarkBoxMuller mide el rendimiento del método polar utilizando funciones trigonométricas (seno/coseno).
// Esta prueba contrastará el coste de las instrucciones de punto flotante en el procesador frente a 
// las simples sumas y multiplicaciones polinomiales del método anterior[cite: 18].
func BenchmarkBoxMuller(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Evaluamos el costo computacional de Box-Muller.
		// Ejecutar nativamente en un entorno Linux moderno asegura que los contadores de ciclos 
		// del host provean métricas de hardware fidedignas para la sustentación matemática[cite: 18].
		_ = models.NewJugador(
			i,
			0.0,
			generators.NormalPolarBoxMuller,
			50.0, 
			15.0, 
		)
	}
}