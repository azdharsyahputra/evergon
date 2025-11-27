import { BrowserRouter, Routes, Route } from "react-router-dom";
import MainLayout from "./layouts/MainLayout";

import Dashboard from "./pages/Dashboard";
import Services from "./pages/Services";
import PhpVersions from "./pages/PhpVersions";
import Settings from "./pages/Settings";
import Projects from "./pages/Projects";
import ProjectConfig from "./pages/ProjectConfig";

export default function App() {
  return (
    <BrowserRouter>
      <MainLayout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/services" element={<Services />} />
          <Route path="/php" element={<PhpVersions />} />
          <Route path="/settings" element={<Settings />} />
          <Route path="/projects" element={<Projects />} />
          <Route path="/projects/:name" element={<ProjectConfig />} />
        </Routes>
      </MainLayout>
    </BrowserRouter>
  );
}
