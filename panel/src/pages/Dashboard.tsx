import {
  Server,
  Code2,
  FolderDot,
  Activity,
  Cpu,
  Gauge,
  BarChart3,
} from "lucide-react";
import { useEffect, useState } from "react";
import type { ReactNode } from "react";

/* =====================================================
   API UTILS
===================================================== */

// POST untuk action (start/stop)
async function callAPI(path: string) {
  const res = await fetch(`http://127.0.0.1:9090${path}`, { method: "POST" });
  if (!res.ok) throw new Error(await res.text());
  return res.text();
}

// GET untuk status
async function getStatus(path: string) {
  const res = await fetch(`http://127.0.0.1:9090${path}`);
  if (!res.ok) throw new Error(await res.text());
  return res.text();
}

// NGINX
const startNginx = () => callAPI("/nginx/start");
const stopNginx = () => callAPI("/nginx/stop");
const reloadNginx = () => callAPI("/nginx/reload");

// PHP
const startPHP = (root: string) =>
  callAPI(`/php/start?root=${encodeURIComponent(root)}`);
const stopPHP = () => callAPI("/php/stop");

// STATUS
const fetchPhpStatus = () => getStatus("/php/status");
const fetchNginxStatus = () => getStatus("/nginx/status");

/* =====================================================
   DASHBOARD PAGE
===================================================== */

export default function Dashboard() {
  const [phpStatus, setPhpStatus] = useState<"running" | "stopped">("stopped");
  const [nginxStatus, setNginxStatus] = useState<"running" | "stopped">(
    "stopped"
  );

  const [loadingPHP, setLoadingPHP] = useState(false);
  const [loadingNginx, setLoadingNginx] = useState(false);

  /* ----------------------------------
      AUTO LOAD + AUTO REFRESH STATUS
  ----------------------------------- */
  useEffect(() => {
    async function loadStatus() {
      try {
        const php = await fetchPhpStatus();
        const ngx = await fetchNginxStatus();

        setPhpStatus(php.includes("running") ? "running" : "stopped");
        setNginxStatus(ngx.includes("running") ? "running" : "stopped");
      } catch (err) {
        console.error("status error:", err);
      }
    }

    // Load awal
    loadStatus();

    // Auto refresh setiap 3 detik
    const timer = setInterval(loadStatus, 3000);
    return () => clearInterval(timer);
  }, []);

  const [phpVersions, setPhpVersions] = useState<{version: string}[]>([]);
  const [currentPhpVersion, setCurrentPhpVersion] = useState<string>("");

  useEffect(() => {
    async function loadPhpVersions() {
      const res = await fetch("http://127.0.0.1:9090/php/version/list");
      const versions = await res.json();
      setPhpVersions(versions);

      const cur = await fetch("http://127.0.0.1:9090/php/version/current");
      const curJson = await cur.json();
      setCurrentPhpVersion(curJson.php_version);
    }

    loadPhpVersions();
  }, []);

  async function changePhpVersion(version: string) {
    await fetch(`http://127.0.0.1:9090/php/version/set?version=${version}`, {
      method: "POST",
    });
    setCurrentPhpVersion(version);
  }

  return (
    <div className="space-y-12">

      {/* HERO */}
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

      {/* SERVICE CARDS */}
      <section className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-6">

      <StatCard
        title="PHP Version"
        value={`PHP ${currentPhpVersion}`}
        icon={<Code2 size={26} className="text-indigo-600" />}
        actions={
          <div className="w-full">
            <label className="block text-xs font-medium text-gray-500 mb-1">
              Select Version
            </label>

            <div className="relative">
              <select
                value={currentPhpVersion}
                onChange={(e) => changePhpVersion(e.target.value)}
                className="
                  w-full appearance-none
                  rounded-xl border border-gray-300
                  bg-white px-3 py-2 pr-10
                  text-sm font-medium text-gray-800
                  shadow-sm
                  focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500
                  hover:border-gray-400
                  transition
                "
              >
                {phpVersions.map((v) => (
                  <option key={v.version} value={v.version}>
                    PHP {v.version}
                  </option>
                ))}
              </select>

              <span
                className="
                  pointer-events-none absolute inset-y-0 right-3
                  flex items-center text-gray-500
                "
              >
                ▼
              </span>
            </div>
          </div>
        }
      />

        {/* PHP STATUS */}
        <StatCard
          title="PHP-FPM"
          value={phpStatus === "running" ? "Running" : "Stopped"}
          icon={<Code2 size={26} />}
          status={phpStatus}
          actions={
            <div className="flex gap-2">
              {phpStatus === "running" ? (
                <button
                  onClick={async () => {
                    setLoadingPHP(true);
                    try {
                      await stopPHP();
                      setPhpStatus("stopped");
                    } finally {
                      setLoadingPHP(false);
                    }
                  }}
                  disabled={loadingPHP}
                  className="px-3 py-1 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 disabled:opacity-50"
                >
                  {loadingPHP ? "..." : "Stop"}
                </button>
              ) : (
                <button
                  onClick={async () => {
                    setLoadingPHP(true);
                    try {
                      await startPHP("/home/azdhar/evergon/www");
                      setPhpStatus("running");
                    } finally {
                      setLoadingPHP(false);
                    }
                  }}
                  disabled={loadingPHP}
                  className="px-3 py-1 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 disabled:opacity-50"
                >
                  {loadingPHP ? "..." : "Start"}
                </button>
              )}
            </div>
          }
        />

        {/* PROJECTS */}
        <StatCard
          title="Active Projects"
          value="7 Projects"
          icon={<FolderDot size={26} />}
        />

        {/* NGINX STATUS */}
        <StatCard
          title="Nginx"
          value={nginxStatus === "running" ? "Running" : "Stopped"}
          icon={<Server size={26} />}
          status={nginxStatus}
          actions={
            <div className="flex gap-2">
              {nginxStatus === "running" ? (
                <>
                  <button
                    onClick={async () => {
                      setLoadingNginx(true);
                      try {
                        await stopNginx();
                        setNginxStatus("stopped");
                      } finally {
                        setLoadingNginx(false);
                      }
                    }}
                    disabled={loadingNginx}
                    className="px-3 py-1 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 disabled:opacity-50"
                  >
                    {loadingNginx ? "..." : "Stop"}
                  </button>

                  <button
                    onClick={async () => {
                      setLoadingNginx(true);
                      try {
                        await reloadNginx();
                      } finally {
                        setLoadingNginx(false);
                      }
                    }}
                    disabled={loadingNginx}
                    className="px-3 py-1 bg-gray-200 text-gray-700 text-sm rounded-lg hover:bg-gray-300 disabled:opacity-50"
                  >
                    {loadingNginx ? "..." : "Reload"}
                  </button>
                </>
              ) : (
                <button
                  onClick={async () => {
                    setLoadingNginx(true);
                    try {
                      await startNginx();
                      setNginxStatus("running");
                    } finally {
                      setLoadingNginx(false);
                    }
                  }}
                  disabled={loadingNginx}
                  className="px-3 py-1 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 disabled:opacity-50"
                >
                  {loadingNginx ? "..." : "Start"}
                </button>
              )}
            </div>
          }
        />

      </section>

      {/* ANALYTICS */}
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

      {/* Coming Soon */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold text-gray-900">Coming Soon</h2>

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

/* =====================================================
   COMPONENTS
===================================================== */

function StatCard({
  title,
  value,
  icon,
  status,
  actions,
}: {
  title: string;
  value: string;
  icon: ReactNode;
  status?: "running" | "stopped";
  actions?: ReactNode;
}) {
  return (
    <div className="bg-white rounded-2xl p-6 border shadow-sm hover:shadow-md transition group">
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

      {actions && <div className="mt-4">{actions}</div>}
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
    <div className="bg-white rounded-2xl p-6 border shadow-sm hover:shadow-md transition">
      <p className="text-gray-500 text-sm font-medium">{title}</p>
      <h2 className="text-3xl font-bold text-gray-900 mt-2">{metric}</h2>
      <div className="mt-4">
        <div className={`h-2 w-full rounded-full ${color} opacity-80`} />
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
