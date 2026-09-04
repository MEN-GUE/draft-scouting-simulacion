# ⚽ Proyecto 1: Simulación de Campamento de Scouting (Draft)
**Curso:** CC2017 - Modelación y Simulación  
**Descripción:** Simulación estadística de un campamento de reclutamiento juvenil para una franquicia de fútbol. El sistema evalúa a miles de aspirantes y filtra el talento utilizando distribuciones de probabilidad y algoritmos de generación de variables aleatorias.

## 📋 Arquitectura del Sistema y Decisiones de Diseño

### 1. El Sistema
El sistema simula un día completo de visorías. Los aspirantes (jugadores) llegan al complejo deportivo, se registran y son sometidos a una evaluación de atributos físicos, técnicos y mentales. El motor de simulación filtrará a los jugadores hasta encontrar a los prospectos "élite" que cumplan con los rigurosos estándares de la academia.

### 2. Modelo de Aleatoriedad (Distribuciones)
El sistema depende de tres fuentes principales de aleatoriedad, cada una fundamentada teóricamente:

*   **Llegada de Aspirantes:** 
    *   *Distribución:* **Poisson / Exponencial**. 
    *   *Justificación:* El volumen de jóvenes que llegan por hora se comporta como un proceso de Poisson. El tiempo exacto entre la llegada de un jugador y el siguiente se modela con una distribución Exponencial.
*   **Atributos del Jugador (Evaluación 1-99):** 
    *   *Distribución:* **Normal** ($\mu = 50$, $\sigma = 15$). 
    *   *Justificación:* Las capacidades y el talento humano se concentran alrededor de una media estadística. Los atributos a evaluar son:
        *   **Físicos:** Velocidad, Resistencia, Fuerza.
        *   **Técnicos:** Control de Balón, Pase Corto, Definición.
        *   **Mentales:** Visión, Compostura.
    *   *Condición "Élite":* Promedio general > 82 o al menos dos atributos clave > 90.
*   **Criterio de Parada (Búsqueda de Talento):** 
    *   *Distribución:* **Binomial Negativa**. 
    *   *Justificación:* El campamento finaliza al asegurar el fichaje de **3 prospectos élite**. La distribución permite determinar cuántos jóvenes (ensayos) fue necesario evaluar hasta acumular exactamente estos 3 éxitos.

### 3. Métodos de Generación de Variables a Comparar
Para cumplir con los requerimientos académicos, se implementarán y compararán los siguientes algoritmos matemáticos en el motor de simulación:

*   **Tiempos de llegada (Exponencial):**
    *   Método de la Transformada Inversa.
*   **Atributos de los jugadores (Normal):**
    *   Método de la Transformada Inversa (mediante aproximación numérica).
    *   Método Polar (Box-Muller) / Aceptación-Rechazo.

### 4. Métricas y Criterios de Comparación
El proyecto calculará las siguientes métricas del sistema:
*   **Tasa de Retención (Embudo):** % de jugadores aceptados vs. descartados.
*   **Tiempo de Reclutamiento:** Cantidad de aspirantes procesados hasta encontrar a las 3 joyas de la cantera.
*   **Gráficos de Radar:** Perfil estadístico visual de los jugadores élite reclutados.

**Evaluación Técnica (Benchmarking):**
Se realizará una comparativa del **Tiempo de Ejecución** (milisegundos) y la **Precisión del Ajuste Estadístico** (histogramas) generados al procesar arreglos masivos de jugadores (ej. 10,000 entidades) utilizando la Transformada Inversa frente al Método Polar.

## 🛠️ Estructura del Repositorio
*   `/backend`: Motor de simulación concurrente (generación de variables, cálculos matemáticos y benchmarks).
*   `/frontend`: Interfaz de usuario y visualización de datos (dashboard, embudos y gráficos de radar).
*   `/docs`: Informe técnico, validaciones estadísticas y documentos adicionales.
