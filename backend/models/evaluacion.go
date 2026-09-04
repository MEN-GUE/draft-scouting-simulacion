package models

// GeneradorNormal define la firma de las funciones algorítmicas del paquete /generators.
type GeneradorNormal func(mu, sigma float64) float64

// NewJugador es el constructor que inicializa un aspirante.
// Recibe el ID, el tiempo generado exponencialmente y la función matemática
// para calcular la Distribución Normal de sus atributos.
func NewJugador(id int, tiempoLlegada float64, generador GeneradorNormal, mu, sigma float64) *Jugador {
	j := &Jugador{
		ID:            id,
		TiempoLlegada: tiempoLlegada,
	}
	j.generarAtributos(generador, mu, sigma)
	j.evaluarCategoria()
	return j
}

// generarAtributos invoca el algoritmo matemático inyectado para generar
// los 8 atributos aleatorios independientes utilizando la media y desviación estándar.
func (j *Jugador) generarAtributos(generador GeneradorNormal, mu, sigma float64) {
	j.Atributos = Atributos{
		Velocidad:    clamp(generador(mu, sigma)),
		Resistencia:  clamp(generador(mu, sigma)),
		Fuerza:       clamp(generador(mu, sigma)),
		ControlBalon: clamp(generador(mu, sigma)),
		PaseCorto:    clamp(generador(mu, sigma)),
		Definicion:   clamp(generador(mu, sigma)),
		Vision:       clamp(generador(mu, sigma)),
		Compostura:   clamp(generador(mu, sigma)),
	}

	// Calcular la valoración general (Promedio)
	total := j.Atributos.Velocidad + j.Atributos.Resistencia + j.Atributos.Fuerza +
		j.Atributos.ControlBalon + j.Atributos.PaseCorto + j.Atributos.Definicion +
		j.Atributos.Vision + j.Atributos.Compostura

	j.PromedioGeneral = float64(total) / 8.0
}

// evaluarCategoria define si el jugador cumple los requisitos para ser el
// éxito buscado por la Distribución Binomial Negativa.
func (j *Jugador) evaluarCategoria() {
	// Condición 1: Promedio General (OVR) sobresaliente
	if j.PromedioGeneral > 82.0 {
		j.EsElite = true
		return
	}

	// Condición 2: Talento altamente especializado (al menos 2 atributos >= 90)
	atributosSobresalientes := 0
	valores := []int{
		j.Atributos.Velocidad, j.Atributos.Resistencia, j.Atributos.Fuerza,
		j.Atributos.ControlBalon, j.Atributos.PaseCorto, j.Atributos.Definicion,
		j.Atributos.Vision, j.Atributos.Compostura,
	}

	for _, val := range valores {
		if val >= 90 {
			atributosSobresalientes++
		}
	}

	if atributosSobresalientes >= 2 {
		j.EsElite = true
		return
	}

	j.EsElite = false
}