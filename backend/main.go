package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	// Nota: Verifica que este path coincida exactamente con la declaración de tu go.mod
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/generators"
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/simulation"
)

// main es el punto de entrada de la aplicación. Orquesta la simulación completa
// del campamento de reclutamiento para la franquicia, generando miles de perfiles
// deportivos para encontrar a los talentos ideales para la academia de desarrollo[cite: 16].
func main() {
	fmt.Println("⚽ Iniciando Simulación de Draft de Fuerzas Básicas...")
	fmt.Println("⚙️  Procesando distribuciones probabilísticas...")

	// 1. Configuración de Parámetros Matemáticos y del Sistema
	// Se busca reclutar a 3 prospectos que cumplan con la condición de élite[cite: 16].
	objetivoElite := 3
	
	// lambda: 50 aspirantes por hora. Define la tasa del proceso de Poisson[cite: 16].
	tasaLlegadas := 50.0 
	
	// Parámetros de la campana de Gauss para escalar los atributos en un rango realista (1-99)[cite: 16].
	mu := 50.0           
	sigma := 15.0        

	// 2. Inicialización del Motor de Simulación (Distribución Binomial Negativa)
	// Se reserva memoria para el historial estadístico para evitar realojamientos de memoria costosos[cite: 16].
	campamento := simulation.NewCampamento(objetivoElite, tasaLlegadas, mu, sigma)

	// 3. Ejecución del Ciclo Principal
	// Inyectamos el algoritmo Polar de Box-Muller. La justificación de esta elección
	// frente a la Transformada Inversa se fundamentará en el análisis del benchmark[cite: 16].
	campamento.Ejecutar(generators.NormalPolarBoxMuller)

	// Resumen de métricas operativas por consola
	fmt.Printf("\n📊 RESULTADOS DE LA SIMULACIÓN\n")
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("Jugadores evaluados en total (Ensayos): %d\n", campamento.JugadoresEvaluados)
	fmt.Printf("Tiempo total del campamento simulado:   %.2f horas\n", campamento.TiempoTotal)
	fmt.Printf("Tasa de retención (Élite / Total):      %.4f%%\n", (float64(objetivoElite)/float64(campamento.JugadoresEvaluados))*100)
	fmt.Printf("--------------------------------------------------\n")

	// 4. Persistencia de Datos para Renderizado Visual y Defensa del Proyecto
	exportarResultadosJSON(campamento)
	exportarDistribucionCSV(campamento.HistorialPromedios)
}

// exportarResultadosJSON serializa las estructuras complejas de los prospectos seleccionados.
// Este JSON alimentará directamente el frontend en Deno Fresh para renderizar los 
// gráficos de radar (Spider Charts) que ilustran el perfil técnico del jugador[cite: 16].
func exportarResultadosJSON(c *simulation.Campamento) {
	archivoJSON, err := os.Create("resultados.json")
	if err != nil {
		log.Fatalf("❌ Error crítico creando resultados.json: %v", err)
	}
	defer archivoJSON.Close()

	// Convertimos el slice de punteros a un formato JSON identado.
	encoder := json.NewEncoder(archivoJSON)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c.JugadoresElite); err != nil {
		log.Fatalf("❌ Error de I/O codificando JSON: %v", err)
	}
	fmt.Println("✅ resultados.json generado exitosamente (Data para perfiles de ojeador).")
}

// exportarDistribucionCSV escribe los promedios en un formato tabular delimitado.
// Este archivo actúa como la evidencia empírica irrefutable requerida por la rúbrica 
// para demostrar que los miles de jugadores procesados forman una campana de Gauss[cite: 16].
func exportarDistribucionCSV(promedios []float64) {
	archivoCSV, err := os.Create("distribucion.csv")
	if err != nil {
		log.Fatalf("❌ Error crítico creando distribucion.csv: %v", err)
	}
	defer archivoCSV.Close()

	writer := csv.NewWriter(archivoCSV)
	defer writer.Flush()

	// Escribir la cabecera estándar de datos
	if err := writer.Write([]string{"JugadorID", "PromedioGeneral"}); err != nil {
		log.Fatalf("❌ Error escribiendo cabecera CSV: %v", err)
	}

	// Escribir los miles de registros generados en tiempo O(n)
	for i, promedio := range promedios {
		fila := []string{
			strconv.Itoa(i + 1),
			fmt.Sprintf("%.4f", promedio),
		}
		if err := writer.Write(fila); err != nil {
			log.Fatalf("❌ Error de I/O escribiendo fila %d en CSV: %v", i, err)
		}
	}
	fmt.Println("✅ distribucion.csv generado exitosamente (Data para validación de la Distribución Normal).")
}