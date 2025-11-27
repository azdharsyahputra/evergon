import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import ProjectCard from "../components/ProjectCard";

interface ProjectItem {
  name: string;
  phpVersion: string;
  phpPort: string;
  status: "running" | "stopped";
  framework: "laravel" | "ci" | "php";
}

const logoMap: Record<ProjectItem["framework"], string> = {
  laravel: "/icons/laravel.svg",
  ci: "/icons/ci.svg",
  php: "/icons/php.svg"
};

export default function Projects() {
  const [projects, setProjects] = useState<ProjectItem[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    loadProjects();
  }, []);

  async function loadProjects() {
    const base = await fetch("http://127.0.0.1:9090/projects").then(r => r.json());

    const enriched = await Promise.all(
      base.map(async (p: any) => {
        const cfg = await fetch(
          `http://127.0.0.1:9090/php/project/get?project=${p.name}`
        ).then(r => r.json());

        const status = await fetch(
          `http://127.0.0.1:9090/php/project/status?project=${p.name}`
        ).then(r => r.json());

        return {
          name: p.name,
          phpVersion: cfg.php_version || "-",
          phpPort: cfg.php_port || "-",
          status: status.running ? "running" : "stopped",
          framework: detectFramework(p.name)
        };
      })
    );

    setProjects(enriched);
  }

  function detectFramework(name: string): ProjectItem["framework"] {
    const n = name.toLowerCase();
    if (n.includes("laravel")) return "laravel";
    if (n.includes("ci")) return "ci";
    return "php";
  }

  return (
    <div className="space-y-10 p-10">
      <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg">
        <h1 className="text-4xl font-extrabold tracking-tight">
          Projects
        </h1>
        <p className="text-indigo-100 text-lg mt-2">
          Manage detected PHP projects in Evergon Workspace.
        </p>
      </section>

      <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">
        {projects.map(p => (
          <ProjectCard
            key={p.name}
            name={p.name}
            phpVersion={p.phpVersion}
            phpPort={p.phpPort}
            status={p.status}
            logo={logoMap[p.framework]}
            onStart={() =>
              fetch(`http://127.0.0.1:9090/php/project/start?project=${p.name}`)
                .then(loadProjects)
            }
            onStop={() =>
              fetch(`http://127.0.0.1:9090/php/project/stop?project=${p.name}`)
                .then(loadProjects)
            }
            onOpen={() => window.open(`http://127.0.0.1:${p.phpPort}`, "_blank")}
            onConfigure={() => navigate(`/projects/${p.name}`)}
          />
        ))}
      </div>
    </div>
  );
}
