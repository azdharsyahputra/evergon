import {
  Play,
  Square,
  RefreshCcw,
  SignalHigh,
  SignalLow,
} from "lucide-react";

// Image assets
import NginxLogo from "../assets/services/nginx.png";
import MysqlLogo from "../assets/services/mysql.png";
import PhpFpmLogo from "../assets/services/php-fpm.svg";
import PostgresLogo from "../assets/services/postgres.svg";

/* ====================================
   SERVICE DEFINITIONS
==================================== */

const services = [
  {
    name: "Nginx",
    desc: "HTTP server, routing & virtual hosts.",
    status: "running",
    color: "indigo",
    logo: NginxLogo,
    cpu: "3%",
    mem: "42 MB",
  },
  {
    name: "PHP-FPM",
    desc: "PHP execution engine for dynamic apps.",
    status: "running",
    color: "emerald",
    logo: PhpFpmLogo,
    cpu: "4%",
    mem: "71 MB",
  },
  {
    name: "MySQL",
    desc: "Relational database engine.",
    status: "stopped",
    color: "sky",
    logo: MysqlLogo,
    cpu: "0%",
    mem: "0 MB",
  },
  {
    name: "PostgreSQL",
    desc: "High-performance SQL database.",
    status: "running",
    color: "violet",
    logo: PostgresLogo,
    cpu: "2%",
    mem: "55 MB",
  },
] as const;

/* ======================================
   PAGE
====================================== */

export default function Services() {
  return (
    <div className="space-y-12">

      <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg mb-6">
        <h1 className="text-4xl font-extrabold tracking-tight">
          Services
        </h1>
        <p className="text-indigo-100 text-lg mt-2">
          Monitor and control all Evergon Engine system services.
        </p>
      </section>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {services.map((s, idx) => (
          <ServiceCard key={idx} {...s} />
        ))}
      </div>
    </div>
  );
}

/* ======================================
   COMPONENT: ServiceCard (with logos)
====================================== */

function ServiceCard({
  name,
  desc,
  status,
  logo,
  color,
  cpu,
  mem,
}: {
  name: string;
  desc: string;
  status: "running" | "stopped";
  logo: string;
  color: string;
  cpu: string;
  mem: string;
}) {
  const isRunning = status === "running";

  return (
    <div className="relative overflow-hidden bg-white/90 backdrop-blur-xl border rounded-2xl shadow-sm hover:shadow-xl transition p-6">

      {/* Gradient Strip */}
      <div
        className={`absolute top-0 left-0 h-full w-1 bg-gradient-to-b from-${color}-500 to-${color}-300`}
      />

      {/* Top */}
      <div className="flex items-start justify-between">
        <div className="flex items-center gap-4">
          <div className="w-14 h-14 rounded-xl bg-gray-50 flex items-center justify-center shadow-sm overflow-hidden">
            <img src={logo} alt={`${name} logo`} className="w-full h-full object-contain" />
          </div>

          <div>
            <h2 className="text-xl font-bold text-gray-900">{name}</h2>
            <p className="text-gray-500 text-sm">{desc}</p>
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

      {/* Metrics */}
      <div className="flex items-center gap-8 mb-5">
        <div className="flex items-center gap-2">
          {isRunning ? (
            <SignalHigh className="text-green-600" size={18} />
          ) : (
            <SignalLow className="text-red-600" size={18} />
          )}
          <div>
            <p className="text-xs text-gray-500">CPU</p>
            <p className="font-semibold text-gray-800">{cpu}</p>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <div className="w-2 h-2 bg-indigo-500 rounded-full" />
          <div>
            <p className="text-xs text-gray-500">Memory</p>
            <p className="font-semibold text-gray-800">{mem}</p>
          </div>
        </div>
      </div>

      {/* Actions */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-3">
          {isRunning ? (
            <button className="flex items-center gap-2 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition">
              <Square size={16} />
              Stop
            </button>
          ) : (
            <button className="flex items-center gap-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition">
              <Play size={16} />
              Start
            </button>
          )}

          <button className="flex items-center gap-2 px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition">
            <RefreshCcw size={16} />
            Restart
          </button>
        </div>

        <p className={`text-sm font-medium ${isRunning ? "text-green-600" : "text-red-600"}`}>
          {isRunning ? "Active" : "Offline"}
        </p>
      </div>
    </div>
  );
}
