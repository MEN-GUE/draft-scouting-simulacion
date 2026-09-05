import { useEffect, useRef } from "preact/hooks";
import { Chart, registerables } from "https://esm.sh/chart.js@4.4.1";

// Registramos todos los componentes internos de Chart.js
Chart.register(...registerables);

export default function HistogramaCampana() {
  const chartRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // 1. Lectura del archivo generado por el backend en Go
        const res = await fetch("/data/distribucion.csv");
        if (!res.ok) throw new Error("No se pudo cargar el CSV");
        const text = await res.text();
        
        // 2. Parseo del CSV
        const rows = text.split("\n").slice(1); // Omitimos la cabecera
        const values: number[] = [];
        
        rows.forEach(row => {
          const cols = row.split(",");
          if (cols.length === 2) {
            values.push(parseFloat(cols[1]));
          }
        });

        // 3. Agrupación de datos (Binning) para el Histograma
        // La media teórica es 50 y la desviación 15. Calculamos rangos de 2 puntos.
        const bins = new Array(30).fill(0);
        const minBin = 20; 
        const binSize = 2;

        values.forEach(val => {
          const index = Math.floor((val - minBin) / binSize);
          if (index >= 0 && index < bins.length) {
            bins[index]++;
          }
        });

        const labels = bins.map((_, i) => `${minBin + i * binSize}-${minBin + (i + 1) * binSize}`);

        // 4. Renderizado del Gráfico
        if (chartRef.current) {
          new Chart(chartRef.current, {
            type: "bar",
            data: {
              labels: labels,
              datasets: [{
                label: "Volumen de Aspirantes",
                data: bins,
                backgroundColor: "rgba(59, 130, 246, 0.7)", // Azul de Tailwind (blue-500)
                borderColor: "rgba(37, 99, 235, 1)",
                borderWidth: 1,
                barPercentage: 1.0,
                categoryPercentage: 1.0 // Elimina espacios para que parezca un histograma real
              }]
            },
            options: {
              responsive: true,
              plugins: {
                title: { 
                  display: true, 
                  text: "Validación Estadística: Distribución Normal de la Valoración General (OVR)",
                  font: { size: 16 }
                },
                legend: { display: false }
              },
              scales: {
                y: { title: { display: true, text: "Frecuencia (Cantidad de Jugadores)" } },
                x: { title: { display: true, text: "Rango de Valoración Promedio" } }
              }
            }
          });
        }
      } catch (error) {
        console.error("Error procesando los datos del histograma:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <div class="w-full max-w-4xl mx-auto p-6 bg-white shadow-md rounded-xl border border-gray-200">
      <canvas ref={chartRef}></canvas>
    </div>
  );
}