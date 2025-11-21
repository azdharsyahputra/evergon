import { Play, Square, RefreshCcw, Folder, Plus } from "lucide-react";
import PhpLogo from "../assets/services/php-fpm.svg";

interface PhpVersion {
  version: string;
  path: string;
  status: "running" | "stopped";
}

const mockPhp: PhpVersion[] = [
  { version: "PHP 7.4", path: "php_versions/php74/", status: "running" },
  { version: "PHP 8.0", path: "php_versions/php80/", status: "stopped" },
  { version: "PHP 8.2", path: "php_versions/php82/", status: "running" },
];

export default function PhpVersions() {
  return (
    <div className="space-y-10">

    {/* Page Header */}
    <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg">
      <h1 className="text-4xl font-extrabold tracking-tight">
        PHP Versions
      </h1>
      <p className="text-indigo-100 text-lg mt-2">
        Manage installed PHP versions for Evergon Engine.
      </p>
    </section>

      {/* PHP List */}
      <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">
        {mockPhp.map((php, idx) => (
          <PhpVersionCard key={idx} php={php} />
        ))}

        {/* ADD NEW VERSION CARD */}
        <AddPhpCard />
      </div>
    </div>
  );
}

// =============================
// CARD: PHP Version
// =============================

function PhpVersionCard({ php }: { php: PhpVersion }) {
  const isRunning = php.status === "running";

  return (
    <div className="bg-white/90 backdrop-blur-xl border rounded-2xl shadow-sm hover:shadow-xl transition p-6 relative overflow-hidden">

      {/* Gradient Bar */}
      <div
        className={`absolute top-0 left-0 w-1 h-full ${
          isRunning
            ? "bg-gradient-to-b from-green-400 to-green-600"
            : "bg-gradient-to-b from-red-400 to-red-600"
        }`}
      />

      {/* Header */}
      <div className="flex items-start justify-between">
        <div className="flex items-center gap-4">
          <div className="w-14 h-14 rounded-xl bg-gray-50 overflow-hidden shadow flex items-center justify-center">
            <img src={PhpLogo} alt="php" className="w-full h-full object-contain" />
          </div>

          <div>
            <h3 className="font-bold text-xl text-gray-900">{php.version}</h3>
            <div className="flex items-center gap-2 text-gray-500 text-sm mt-1">
              <Folder size={15} />
              {php.path}
            </div>
          </div>
        </div>

        <span
          className={`px-3 py-1 rounded-full text-sm font-semibold ${
            isRunning ? "bg-green-100 text-green-700" : "bg-red-100 text-red-700"
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
          <button className="px-4 py-2 bg-red-600 text-white rounded-lg flex items-center gap-2 hover:bg-red-700 transition">
            <Square size={18} />
            Stop
          </button>
        ) : (
          <button className="px-4 py-2 bg-green-600 text-white rounded-lg flex items-center gap-2 hover:bg-green-700 transition">
            <Play size={18} />
            Start
          </button>
        )}

        <button className="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg flex items-center gap-2 hover:bg-gray-300 transition">
          <RefreshCcw size={18} />
          Restart
        </button>
      </div>

    </div>
  );
}

// =============================
// CARD: Add New PHP Version
// =============================

function AddPhpCard() {
  return (
    <button className="group bg-white/70 backdrop-blur-xl border-2 border-dashed border-gray-300 rounded-2xl p-6 shadow-sm hover:shadow-lg hover:border-indigo-400 transition flex flex-col items-center justify-center text-center gap-4">
      
      <div className="w-16 h-16 rounded-xl bg-gray-100 group-hover:bg-indigo-50 flex items-center justify-center transition">
        <Plus size={32} className="text-gray-500 group-hover:text-indigo-600" />
      </div>

      <div>
        <p className="font-semibold text-gray-800 group-hover:text-indigo-700">
          Add New PHP Version
        </p>
        <p className="text-sm text-gray-500 mt-1">
          Install and register another PHP runtime.
        </p>
      </div>
    </button>
  );
}
