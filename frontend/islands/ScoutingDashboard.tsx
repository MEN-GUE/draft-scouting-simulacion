import { useState } from "preact/hooks";
import HistogramaCampana from "./HistogramaCampana.tsx";
import RadarProspectos from "./RadarProspectos.tsx";

export default function ScoutingDashboard() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [showEliteOnly, setShowEliteOnly] = useState(false);

  const runSimulation = async () => {
    setLoading(true);
    try {
      const res = await fetch("http://localhost:8080/api/simular");
      const json = await res.json();
      setData(json);
    } catch (error) {
      console.error("Error conectando con el backend en Go:", error);
      alert("Error: Asegúrate de que el servidor Go esté corriendo en el puerto 8080.");
    } finally {
      setLoading(false);
    }
  };

  const displayedPlayers = data 
    ? (showEliteOnly ? data.jugadoresElite : data.todosLosJugadores)
    : [];

  return (
    <div class="space-y-8">
      {/* Panel de Control */}
      <div class="bg-white p-6 rounded-xl shadow-md border border-gray-200 flex flex-col md:flex-row justify-between items-center">
        <div>
          <h2 class="text-2xl font-bold text-gray-800">Motor de Simulación</h2>
          <p class="text-gray-500">Distribución Binomial Negativa & Normal Polar (Box-Muller)</p>
        </div>
        <button 
          onClick={runSimulation} 
          disabled={loading}
          class="mt-4 md:mt-0 px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-bold rounded-lg transition-colors disabled:opacity-50"
        >
          {loading ? "Simulando..." : "▶ Ejecutar Nueva Simulación"}
        </button>
      </div>

      {data && !loading && (
        <>
          {/* Métricas Rápidas */}
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 text-center">
              <p class="text-gray-500 font-semibold">Jugadores Evaluados (Ensayos)</p>
              <p class="text-4xl font-black text-blue-600">{data.jugadoresEvaluados}</p>
            </div>
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 text-center">
              <p class="text-gray-500 font-semibold">Tiempo Simulado</p>
              <p class="text-4xl font-black text-gray-800">{data.tiempoTotal.toFixed(1)} <span class="text-xl">hrs</span></p>
            </div>
            <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 text-center">
              <p class="text-gray-500 font-semibold">Prospectos Élite Reclutados</p>
              <p class="text-4xl font-black text-green-600">{data.objetivoElite}</p>
            </div>
          </div>

          {/* Gráficas Visuales */}
          <RadarProspectos prospectos={data.jugadoresElite} />
          <HistogramaCampana jugadores={data.todosLosJugadores} />

          {/* Auditoría / Tabla de Datos */}
          <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
            <div class="flex justify-between items-center mb-4 border-b pb-4">
              <h3 class="text-xl font-bold text-gray-800">Auditoría de Aspirantes</h3>
              <label class="flex items-center space-x-2 cursor-pointer">
                <input 
                  type="checkbox" 
                  checked={showEliteOnly} 
                  onChange={(e) => setShowEliteOnly((e.target as HTMLInputElement).checked)}
                  class="form-checkbox h-5 w-5 text-blue-600"
                />
                <span class="text-gray-700 font-semibold">Filtrar solo Élite</span>
              </label>
            </div>
            
            <div class="overflow-x-auto max-h-96">
              <table class="w-full text-sm text-left text-gray-500">
                <thead class="text-xs text-gray-700 uppercase bg-gray-50 sticky top-0">
                  <tr>
                    <th class="px-6 py-3">ID</th>
                    <th class="px-6 py-3">Tiempo Llegada</th>
                    <th class="px-6 py-3">OVR</th>
                    <th class="px-6 py-3">Velocidad</th>
                    <th class="px-6 py-3">Tiro</th>
                    <th class="px-6 py-3">Visión</th>
                    <th class="px-6 py-3">Estado</th>
                  </tr>
                </thead>
                <tbody>
                  {displayedPlayers.map((j: any) => (
                    <tr key={j.ID} class="bg-white border-b hover:bg-gray-50">
                      <td class="px-6 py-4 font-bold text-gray-900">#{j.ID}</td>
                      <td class="px-6 py-4">{j.TiempoLlegada.toFixed(2)} h</td>
                      <td class="px-6 py-4 font-bold">{j.PromedioGeneral.toFixed(1)}</td>
                      <td class="px-6 py-4">{j.Atributos.Velocidad}</td>
                      <td class="px-6 py-4">{j.Atributos.Definicion}</td>
                      <td class="px-6 py-4">{j.Atributos.Vision}</td>
                      <td class="px-6 py-4">
                        {j.EsElite 
                          ? <span class="bg-green-100 text-green-800 px-2 py-1 rounded font-bold">Fichado</span> 
                          : <span class="bg-red-100 text-red-800 px-2 py-1 rounded">Rechazado</span>}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </>
      )}
    </div>
  );
}