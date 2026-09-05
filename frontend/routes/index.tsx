import ScoutingDashboard from "../islands/ScoutingDashboard.tsx";

export default function Home() {
  return (
    <div class="min-h-screen bg-gray-50 p-8 font-sans">
      <header class="max-w-6xl mx-auto mb-10 text-center">
        <h1 class="text-5xl font-black text-gray-900 mb-3 tracking-tight">⚽ Dashboard de Visorías</h1>
        <p class="text-gray-600 text-lg font-medium">Plataforma interactiva de simulación y reclutamiento juvenil</p>
      </header>

      <main class="max-w-6xl mx-auto">
        <ScoutingDashboard />
      </main>
      
      <footer class="max-w-6xl mx-auto mt-16 text-center text-gray-400 text-sm border-t pt-6">
        Proyecto 1: Modelación y Simulación - Universidad del Valle de Guatemala
      </footer>
    </div>
  );
}