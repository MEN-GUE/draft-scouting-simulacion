import { useEffect, useRef } from "preact/hooks";
import { Chart, registerables } from "https://esm.sh/chart.js@4.4.1";

Chart.register(...registerables);

const PlayerRadar = ({ player }: { player: any }) => {
  const chartRef = useRef<HTMLCanvasElement>(null);
  const chartInstance = useRef<Chart | null>(null);

  useEffect(() => {
    if (!chartRef.current) return;
    if (chartInstance.current) chartInstance.current.destroy();

    chartInstance.current = new Chart(chartRef.current, {
      type: "radar",
      data: {
        labels: ["Vel", "Res", "Fuerza", "Control", "Pase", "Tiro", "Visión", "Comp"],
        datasets: [{
          label: `OVR: ${player.PromedioGeneral.toFixed(1)}`,
          data: [
            player.Atributos.Velocidad, player.Atributos.Resistencia, player.Atributos.Fuerza,
            player.Atributos.ControlBalon, player.Atributos.PaseCorto, player.Atributos.Definicion,
            player.Atributos.Vision, player.Atributos.Compostura
          ],
          backgroundColor: "rgba(34, 197, 94, 0.2)",
          borderColor: "rgba(21, 128, 61, 1)",
          pointBackgroundColor: "rgba(21, 128, 61, 1)",
          borderWidth: 2,
        }]
      },
      options: {
        responsive: true,
        scales: {
          r: {
            angleLines: { color: 'rgba(0, 0, 0, 0.1)' },
            grid: { color: 'rgba(0, 0, 0, 0.1)' },
            pointLabels: { font: { size: 10 } },
            ticks: { display: false, min: 0, max: 100 }
          }
        },
        plugins: { legend: { position: 'bottom' } }
      }
    });

    return () => { if (chartInstance.current) chartInstance.current.destroy(); };
  }, [player]);

  return <canvas ref={chartRef}></canvas>;
};

export default function RadarProspectos({ prospectos }: { prospectos: any[] }) {
  if (!prospectos || prospectos.length === 0) return null;

  return (
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      {prospectos.map((p) => (
        <div key={p.ID} class="bg-white p-4 rounded-xl border border-gray-200 shadow-sm">
          <div class="text-center mb-2 border-b pb-2">
            <span class="text-xs font-bold tracking-widest text-green-600 uppercase">Fichaje</span>
            <h3 class="text-lg font-bold text-gray-800">Aspirante #{p.ID}</h3>
          </div>
          <PlayerRadar player={p} />
        </div>
      ))}
    </div>
  );
}