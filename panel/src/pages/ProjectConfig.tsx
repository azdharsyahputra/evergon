import { useParams, Link } from "react-router-dom";
import { useEffect, useState } from "react";
import { ArrowLeft, ServerCog, Plug } from "lucide-react";

interface ProjectConfigData {
  php_version: string;
  php_port: string;
}

interface PhpVersionItem {
  version: string;
  path: string;
}

export default function ProjectConfig() {
  const { name } = useParams();
  const [config, setConfig] = useState<ProjectConfigData>({
    php_version: "",
    php_port: "",
  });

  const [phpVersions, setPhpVersions] = useState<PhpVersionItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    loadConfig();
    loadPhpVersions();
  }, []);

  async function loadConfig() {
    if (!name) return;

    const res = await fetch(
      `http://127.0.0.1:9090/php/project/get?project=${name}`
    );
    const data = await res.json();

    setConfig({
      php_version: data.php_version || "",
      php_port: data.php_port || "",
    });

    setLoading(false);
  }

  async function loadPhpVersions() {
    const res = await fetch(`http://127.0.0.1:9090/php/version/list`);
    const data = await res.json();
    setPhpVersions(data);
  }

  async function saveConfig() {
    if (!name) return;
    setSaving(true);

    await fetch(
      `http://127.0.0.1:9090/php/project/set?project=${name}&version=${config.php_version}&port=${config.php_port}`
    );

    setSaving(false);
  }

  if (loading) {
    return (
      <div className="p-10 text-gray-600 text-lg">
        Loading configuration...
      </div>
    );
  }

  return (
    <div className="space-y-10 p-10">

      {/* Header Section */}
      <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg flex items-center justify-between">
        <div>
          <h1 className="text-4xl font-extrabold tracking-tight">
            {name}
          </h1>
          <p className="text-indigo-100 text-lg mt-2">
            Project configuration & runtime settings
          </p>
        </div>

        <Link
          to="/projects"
          className="bg-white/20 hover:bg-white/30 transition px-4 py-2 rounded-lg flex items-center gap-2 text-white"
        >
          <ArrowLeft size={18} />
          Back to Projects
        </Link>
      </section>

      {/* Editable Config */}
      <section className="bg-white/90 backdrop-blur-xl border rounded-2xl shadow p-8 space-y-6">
        <h2 className="text-2xl font-bold text-gray-900 flex items-center gap-3">
          <ServerCog size={22} className="text-indigo-600" />
          Runtime Configuration
        </h2>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

          {/* PHP Version Dropdown */}
          <div>
            <label className="text-gray-600 text-sm">PHP Version</label>
            <select
              value={config.php_version}
              onChange={(e) =>
                setConfig({ ...config, php_version: e.target.value })
              }
              className="w-full mt-1 border rounded-lg p-3 bg-gray-50 text-gray-800"
            >
              <option value="">Select Version</option>
              {phpVersions.map(v => (
                <option key={v.version} value={v.version}>
                  {v.version}
                </option>
              ))}
            </select>
          </div>

          {/* Port Input */}
          <div>
            <label className="text-gray-600 text-sm">Assigned Port</label>
            <div className="flex items-center gap-2 mt-1">
              <input
                type="text"
                value={config.php_port}
                onChange={(e) =>
                  setConfig({ ...config, php_port: e.target.value })
                }
                className="w-full border rounded-lg p-3 bg-gray-50 text-gray-800"
              />
              <Plug size={18} className="text-indigo-600" />
            </div>
          </div>
        </div>

        {/* Save Button */}
        <button
          onClick={saveConfig}
          disabled={saving}
          className={`px-6 py-3 rounded-lg text-white font-semibold transition 
            ${saving ? "bg-gray-400" : "bg-indigo-600 hover:bg-indigo-700"}`}
        >
          {saving ? "Saving..." : "Save Changes"}
        </button>
      </section>
    </div>
  );
}
