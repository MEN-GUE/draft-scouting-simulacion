package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MEN-GUE/draft-scouting-simulacion/backend/generators"
	"github.com/MEN-GUE/draft-scouting-simulacion/backend/simulation"
)

// handlerSimulacion recibe la petición HTTP del frontend, ejecuta el motor
// probabilístico en tiempo real y devuelve el estado completo del campamento en JSON.
func handlerSimulacion(w http.ResponseWriter, r *http.Request) {
	// 1. Configuración de cabeceras CORS para permitir peticiones desde localhost:8000 (Deno Fresh)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Interceptar peticiones pre-flight de CORS del navegador
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 2. Configuración de Parámetros Matemáticos y del Sistema
	objetivoElite := 3
	tasaLlegadas := 50.0 // lambda: aspirantes/hora
	mu := 50.0           // Media (OVR)
	sigma := 15.0        // Desviación estándar

	// 3. Inicialización y Ejecución del Motor (Binomial Negativa)
	campamento := simulation.NewCampamento(objetivoElite, tasaLlegadas, mu, sigma)

	// Utilizamos el algoritmo validado como el más apto según nuestras pruebas locales
	campamento.Ejecutar(generators.NormalPolarBoxMuller)

	fmt.Printf("✅ Simulación finalizada (API): %d aspirantes evaluados en %.2f horas.\n",
		campamento.JugadoresEvaluados, campamento.TiempoTotal)

	// 4. Serialización y respuesta al cliente
	// Go codificará la estructura `Campamento` completa a JSON automáticamente
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(campamento); err != nil {
		http.Error(w, "Error interno codificando los datos de la simulación", http.StatusInternalServerError)
		log.Printf("❌ Error de codificación JSON: %v\n", err)
	}
}

func main() {
	// Definimos la ruta de nuestra API REST
	http.HandleFunc("/api/simular", handlerSimulacion)

	fmt.Println("🚀 Servidor Go (Motor de Simulación) iniciado correctamente.")
	fmt.Println("📡 Escuchando peticiones HTTP en http://localhost:8080 ...")

	// Levantamos el servidor en el puerto 8080 para evitar conflictos con el 8000 de Deno
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("❌ Error crítico iniciando el servidor HTTP: ", err)
	}
}
