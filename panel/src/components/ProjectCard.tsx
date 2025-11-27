import { Play, Square, Globe, Settings } from "lucide-react";

interface ProjectCardProps {
  name: string;
  phpVersion: string;
  phpPort: string;
  status: "running" | "stopped";
  logo: string;
  onStart: () => void | Promise<void>;
  onStop: () => void | Promise<void>;
  onOpen: () => void;
  onConfigure: () => void | Promise<void>;
}

export default function ProjectCard({
  name,
  phpVersion,
  phpPort,
  status,
  logo,
  onStart,
  onStop,
  onOpen,
  onConfigure
}: ProjectCardProps) {

  const isRunning = status === "running";

  return (
    <div className="bg-white/90 backdrop-blur-xl border rounded-2xl shadow-sm hover:shadow-xl transition p-6 relative overflow-hidden">

      {/* Accent Bar */}
      <div
        className={`absolute top-0 left-0 w-1 h-full ${
          isRunning
            ? "bg-gradient-to-b from-green-400 to-green-600"
            : "bg-gradient-to-b from-gray-400 to-gray-600"
        }`}
      />

      {/* Header */}
      <div className="flex items-start justify-between">
        <div className="flex items-center gap-4">
          <div className="w-14 h-14 rounded-xl bg-gray-50 overflow-hidden shadow flex items-center justify-center">
            <img src={logo} alt="logo" className="w-full h-full object-contain" />
          </div>

          <div>
            <h3 className="font-bold text-xl text-gray-900">{name}</h3>
            <div className="flex items-center gap-2 text-gray-500 text-sm mt-1">
              PHP {phpVersion} • Port {phpPort}
            </div>
          </div>
        </div>

        <span
          className={`px-3 py-1 rounded-full text-sm font-semibold ${
            isRunning ? "bg-green-100 text-green-700" : "bg-gray-200 text-gray-600"
          }`}
        >
          {isRunning ? "Running" : "Stopped"}
        </span>
      </div>

      {/* Divider */}
      <div className="border-t my-5" />

      {/* Buttons */}
      <div className="flex items-center gap-3">
        {isRunning ? (
          <button
            onClick={onStop}
            className="px-4 py-2 bg-red-600 text-white rounded-lg flex items-center gap-2 hover:bg-red-700 transition"
          >
            <Square size={18} />
            Stop
          </button>
        ) : (
          <button
            onClick={onStart}
            className="px-4 py-2 bg-green-600 text-white rounded-lg flex items-center gap-2 hover:bg-green-700 transition"
          >
            <Play size={18} />
            Start
          </button>
        )}

        <button
          onClick={onOpen}
          className="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg flex items-center gap-2 hover:bg-gray-300 transition"
        >
          <Globe size={18} />
          Open
        </button>

        <button
          onClick={onConfigure}
          className="px-4 py-2 bg-indigo-600 text-white rounded-lg flex items-center gap-2 hover:bg-indigo-700 transition"
        >
          <Settings size={18} />
          Configure
        </button>
      </div>

    </div>
  );
}
