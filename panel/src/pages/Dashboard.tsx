import {
  Server,
  Code2,
  FolderDot,
  Activity,
  Cpu,
  Gauge,
  BarChart3,
} from "lucide-react";
import type { ReactNode } from "react";

export default function Dashboard() {
  return (
    <div className="space-y-12">

      {/* HERO SECTION */}
      <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg">
        <h1 className="text-4xl font-extrabold tracking-tight">
          Welcome to Evergon
        </h1>
        <p className="text-indigo-100 text-lg mt-2">
          Local development engine & server manager powered by Go.
        </p>

        <div className="mt-6 flex items-center gap-6">
          <div className="flex items-center gap-2 text-indigo-100">
            <Activity size={20} className="text-green-300" />
            Engine Status: <span className="font-semibold">Running</span>
          </div>

          <div className="flex items-center gap-2 text-indigo-100">
            <Cpu size={20} className="text-yellow-300" />
            CPU Load: <span className="font-semibold">12%</span>
          </div>

          <div className="flex items-center gap-2 text-indigo-100">
            <Gauge size={20} className="text-blue-200" />
            Memory: <span className="font-semibold">32%</span>
          </div>
        </div>
      </section>

      {/* TOP STATS */}
      <section className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-6">

        <StatCard
          title="Engine Status"
          value="Running"
          icon={<Activity size={26} />}
          status="running"
        />

        <StatCard
          title="PHP Versions"
          value="3 Installed"
          icon={<Code2 size={26} />}
        />

        <StatCard
          title="Active Projects"
          value="7 Projects"
          icon={<FolderDot size={26} />}
        />

        <StatCard
          title="Nginx Status"
          value="Running"
          icon={<Server size={26} />}
          status="running"
        />
      </section>

      {/* ANALYTICS SECTION */}
      <section className="space-y-6">
        <h2 className="text-xl font-semibold text-gray-800 flex items-center gap-2">
          <BarChart3 size={20} className="text-indigo-500" />
          System Analytics
        </h2>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

          <AnalyticsCard
            title="CPU Usage (Last 30s)"
            metric="12%"
            color="bg-indigo-500"
          />

          <AnalyticsCard
            title="Memory Usage"
            metric="32%"
            color="bg-green-500"
          />
        </div>
      </section>

      {/* COMING SOON */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold text-gray-900">
          Coming Soon
        </h2>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

          <ComingCard
            icon={<Activity size={20} />}
            title="Service Controls"
            subtitle="Start/Stop PHP-FPM, MySQL, PostgreSQL & Nginx."
          />

          <ComingCard
            icon={<FolderDot size={20} />}
            title="Project Scanner"
            subtitle="Auto-detect Laravel, CI4, React, Next.js & Node."
          />
        </div>
      </section>
    </div>
  );
}

/* ============================
   COMPONENTS
============================ */

function StatCard({
  title,
  value,
  icon,
  status,
}: {
  title: string;
  value: string;
  icon: ReactNode;
  status?: "running" | "stopped";
}) {
  return (
    <div className="bg-white rounded-2xl p-6 shadow-sm hover:shadow-md transition group border">
      <div className="flex items-center justify-between">
        <p className="text-gray-500 text-sm font-medium">{title}</p>

        <div className="p-2 rounded-xl bg-indigo-50 text-indigo-600 group-hover:bg-indigo-100 transition">
          {icon}
        </div>
      </div>

      {status ? (
        <div className="mt-3">
          <span
            className={`px-3 py-1 rounded-full text-sm font-semibold ${
              status === "running"
                ? "bg-green-100 text-green-700"
                : "bg-red-100 text-red-700"
            }`}
          >
            {value}
          </span>
        </div>
      ) : (
        <h2 className="mt-3 text-2xl font-bold text-gray-900">{value}</h2>
      )}
    </div>
  );
}

function AnalyticsCard({
  title,
  metric,
  color,
}: {
  title: string;
  metric: string;
  color: string;
}) {
  return (
    <div className="bg-white rounded-2xl p-6 shadow-sm hover:shadow-md transition border">
      <p className="text-gray-500 text-sm font-medium">{title}</p>

      <h2 className="text-3xl font-bold text-gray-900 mt-2">{metric}</h2>

      {/* pseudo sparkline */}
      <div className="mt-4">
        <div className={`h-2 rounded-full ${color} opacity-80 w-full`}></div>
      </div>
    </div>
  );
}

function ComingCard({
  icon,
  title,
  subtitle,
}: {
  icon: ReactNode;
  title: string;
  subtitle: string;
}) {
  return (
    <div className="bg-white p-6 border rounded-2xl shadow-sm hover:shadow-md transition group">
      <h3 className="font-semibold text-gray-800 flex items-center gap-2">
        <span className="text-indigo-500">{icon}</span>
        {title}
      </h3>
      <p className="text-gray-500 text-sm mt-1">{subtitle}</p>
    </div>
  );
}
