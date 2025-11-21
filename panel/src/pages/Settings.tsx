import { Cog, Monitor, Info, Wrench } from "lucide-react";
import type { ReactNode } from "react";

export default function Settings() {
  return (
    <div className="space-y-10">

      {/* Title */}
      <section className="bg-gradient-to-r from-indigo-600 to-indigo-500 rounded-2xl p-8 text-white shadow-lg mb-6">
        <h1 className="text-4xl font-extrabold tracking-tight">
          Settings
        </h1>
        <p className="text-indigo-100 text-lg mt-2">
          Manage Evergon preferences, engine configuration, and UI options.
        </p>
      </section>

      {/* Settings Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">

        {/* General Settings */}
        <SettingCard
          icon={<Cog size={22} className="text-indigo-600" />}
          title="General Settings"
          description="Global settings and default behavior of Evergon Panel."
        >
          <SettingRow label="Workspace Path" value="/home/azdhar/evergon" />
          <SettingRow label="Auto-Start Engine" value="Enabled" />
          <SettingRow label="Default PHP Version" value="PHP 8.2" />
        </SettingCard>

        {/* Engine Settings */}
        <SettingCard
          icon={<Wrench size={22} className="text-indigo-600" />}
          title="Engine Configuration"
          description="Manage Evergon Engine runtime behavior."
        >
          <SettingRow label="Engine API Port" value="9999" />
          <SettingRow label="Service Timeout" value="5000 ms" />
          <SettingRow label="Error Logging" value="Enabled" />
        </SettingCard>

        {/* UI Preferences */}
        <SettingCard
          icon={<Monitor size={22} className="text-indigo-600" />}
          title="UI Preferences"
          description="Customize the look & feel of Evergon Panel."
        >
          <SettingRow label="Theme" value="Light" />
          <SettingRow label="Sidebar Mode" value="Expanded" />
          <SettingRow label="Animations" value="Enabled" />
        </SettingCard>

        {/* About Evergon */}
        <SettingCard
          icon={<Info size={22} className="text-indigo-600" />}
          title="About Evergon"
          description="Information about this version of Evergon."
        >
          <SettingRow label="Version" value="v0.1.0" />
          <SettingRow label="Engine Build" value="Go-based Core Engine" />
          <SettingRow label="Status" value="Under Development" />
        </SettingCard>

      </div>
    </div>
  );
}

// ===============================================
// REUSABLE COMPONENTS
// ===============================================

// Main Card
function SettingCard({
  icon,
  title,
  description,
  children,
}: {
  icon: ReactNode;
  title: string;
  description: string;
  children: ReactNode;
}) {
  return (
    <div className="bg-white border rounded-2xl shadow-sm p-6 hover:shadow-md transition">

      <div className="flex items-center gap-3 mb-4">
        <div className="p-2 bg-indigo-50 rounded-xl">{icon}</div>
        <div>
          <h3 className="font-semibold text-gray-900">{title}</h3>
          <p className="text-gray-500 text-sm">{description}</p>
        </div>
      </div>

      <div className="space-y-3">
        {children}
      </div>
    </div>
  );
}

// Row Component
function SettingRow({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex items-center justify-between py-2 px-3 bg-gray-50 rounded-lg">
      <span className="text-gray-600 font-medium">{label}</span>
      <span className="text-gray-900">{value}</span>
    </div>
  );
}
