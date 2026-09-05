import { useEffect, useRef, useState } from "preact/hooks";
import { Chart, registerables } from "https://esm.sh/chart.js@4.4.1";

Chart.register(...registerables);

// Interfaz para el tipado de los datos que vienen del backend en Go
interface Atributos {
  Velocidad: number; Resistencia: number; Fuerza: number;
  ControlBalon: number; PaseCorto: number; Definicion: number;
  Vision: number; Compostura: number;
}

interface Jugador {
  ID: number;
  TiempoLlegada: number;
  Atributos: Atributos;
  PromedioGeneral: number;
  EsElite: boolean;
}

// Sub-componente que renderiza un gráfico de radar individual
const PlayerRadar = ({ player }: { player: Jugador }) => {
  const chartRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    if (chartRef.current) {
      const chart = new Chart(chartRef.current, {
        type: "radar",
        data: {
          labels: [
            "Velocidad", "Resistencia", "Fuerza", 
            "Control Balón", "Pase Corto", "Definición", 
            "Visión", "Compostura"
          ],
          datasets: [{
            label: `Valoración: ${player.PromedioGeneral.toFixed(1)}`,
            data: [
              player.Atributos.Velocidad, player.Atributos.Resistencia, player.Atributos.Fuerza,
              player.Atributos.ControlBalon, player.Atributos.PaseCorto, player.Atributos.Definicion,
              player.Atributos.Vision, player.Atributos.Compostura
            ],
            backgroundColor: "rgba(34, 197, 94, 0.2)", // Verde tailwind (green-500)
            borderColor: "rgba(21, 128, 61, 1)", // Verde oscuro (green-700)
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
              pointLabels: { font: { size: 11, family: 'sans-serif' }, color: '#374151' },
              ticks: { display: false, min: 0, max: 100, stepSize: 20 }
            }
          },
          plugins: {
            legend: { position: 'bottom' }
          }
        }
      });
      // Limpieza del canvas para evitar bugs si Deno Fresh recarga el componente
      return () => chart.destroy();
    }
  }, [player]);

  return <canvas ref={chartRef}></canvas>;
};

// Componente principal exportado
export default function RadarProspectos() {
  const [players, setPlayers] = useState<Jugador[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/data/resultados.json")
      .then(res => res.json())
      .then(data => {
        setPlayers(data);
        setLoading(false);
      })
      .catch(err => {
        console.error("Error cargando prospectos:", err);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div class="text-center text-gray-500 py-10">Generando reportes de ojeador...</div>;
  }

  return (
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      {players.map((p) => (
        <div key={p.ID} class="bg-white p-6 rounded-xl border border-gray-200 shadow-sm hover:shadow-md transition-shadow">
          <div class="text-center mb-4 border-b pb-2">
            <span class="text-xs font-bold tracking-widest text-green-600 uppercase">Fichaje Confirmado</span>
            <h3 class="text-xl font-bold text-gray-800 mt-1">Aspirante #{p.ID}</h3>
            <p class="text-sm text-gray-500">
              Evaluado a las {p.TiempoLlegada.toFixed(1)} hrs
            </p>
          </div>
          <PlayerRadar player={p} />
        </div>
      ))}
    </div>
  );
}