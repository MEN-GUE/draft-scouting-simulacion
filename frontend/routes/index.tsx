import HistogramaCampana from "../islands/HistogramaCampana.tsx";
import RadarProspectos from "../islands/RadarProspectos.tsx";

export default function Home() {
  return (
    <div class="min-h-screen bg-gray-50 p-8 font-sans">
      <header class="max-w-6xl mx-auto mb-12 text-center">
        <h1 class="text-4xl font-extrabold text-gray-900 mb-2">⚽ Dashboard de Visorías</h1>
        <p class="text-gray-600 text-lg">Plataforma de análisis probabilístico para fuerzas básicas</p>
      </header>

      <main class="max-w-6xl mx-auto space-y-16">
        
        {/* SECCIÓN 1: Resultados de la Binomial Negativa (Prospectos) */}
        <section>
          <div class="mb-6">
            <h2 class="text-2xl font-bold text-gray-800">1. Reportes de Ojeador (Prospectos Élite)</h2>
            <p class="text-gray-500 text-sm mt-1">
              Resultados de la simulación iterativa (Binomial Negativa) hasta encontrar los 3 objetivos principales.
            </p>
          </div>
          <RadarProspectos />
        </section>

        {/* SECCIÓN 2: Validación del Modelo Normal (Histograma) */}
        <section>
          <div class="mb-6">
            <h2 class="text-2xl font-bold text-gray-800">2. Validación Estadística de Generadores</h2>
            <p class="text-gray-500 text-sm mt-1">
              Evidencia empírica de la Campana de Gauss utilizando el Método Polar (Box-Muller) procesado por el backend en Go.
            </p>
          </div>
          <HistogramaCampana />
        </section>

      </main>
      
      <footer class="max-w-6xl mx-auto mt-16 text-center text-gray-400 text-sm border-t pt-6">
        Proyecto 1: Modelación y Simulación - Universidad del Valle de Guatemala
      </footer>
    </div>
  );
}