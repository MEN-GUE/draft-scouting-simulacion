import { useEffect, useRef } from "preact/hooks";
import { Chart, registerables } from "https://esm.sh/chart.js@4.4.1";

Chart.register(...registerables);

export default function HistogramaCampana({ jugadores }: { jugadores: any[] }) {
  const chartRef = useRef<HTMLCanvasElement>(null);
  const chartInstance = useRef<Chart | null>(null);

  useEffect(() => {
    if (!chartRef.current || jugadores.length === 0) return;

    // Destruir la gráfica anterior si existe (vital para evitar bugs al re-simular)
    if (chartInstance.current) {
      chartInstance.current.destroy();
    }

    const bins = new Array(35).fill(0);
    const minBin = 20; 
    const binSize = 2;

    jugadores.forEach(j => {
      const index = Math.floor((j.PromedioGeneral - minBin) / binSize);
      if (index >= 0 && index < bins.length) {
        bins[index]++;
      }
    });

    const labels = bins.map((_, i) => `${minBin + i * binSize}-${minBin + (i + 1) * binSize}`);

    chartInstance.current = new Chart(chartRef.current, {
      type: "bar",
      data: {
        labels: labels,
        datasets: [{
          label: "Volumen de Aspirantes",
          data: bins,
          backgroundColor: "rgba(59, 130, 246, 0.7)",
          borderColor: "rgba(37, 99, 235, 1)",
          borderWidth: 1,
          barPercentage: 1.0,
          categoryPercentage: 1.0 
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: { display: false }
        },
        scales: {
          y: { title: { display: true, text: "Frecuencia" } },
          x: { title: { display: true, text: "Rango de Valoración Promedio" } }
        }
      }
    });

    return () => {
      if (chartInstance.current) chartInstance.current.destroy();
    };
  }, [jugadores]);

  return (
    <div class="w-full bg-white p-4 rounded-xl border border-gray-200 shadow-sm">
      <h3 class="text-lg font-bold text-gray-800 text-center mb-4">Validación: Campana de Gauss (OVR)</h3>
      <canvas ref={chartRef}></canvas>
    </div>
  );
}