import { NavLink } from "react-router-dom";
import { Home, ServerCog, Code, Settings, FolderDot } from "lucide-react";

const nav = [
  { name: "Dashboard", to: "/", icon: Home },
  { name: "Services", to: "/services", icon: ServerCog },
  { name: "Projects", to: "/projects", icon: FolderDot },
  { name: "PHP Versions", to: "/php", icon: Code },
  { name: "Settings", to: "/settings", icon: Settings },
];

export default function Sidebar() {
  return (
    <div className="w-64 bg-white border-r shadow-sm flex flex-col">
      <div className="p-4 text-xl font-semibold tracking-wide">
        Evergon Panel
      </div>

      <nav className="flex-1 px-3 space-y-1">
        {nav.map((item) => {
          const Icon = item.icon;
          return (
            <NavLink
              key={item.name}
              to={item.to}
              className={({ isActive }) =>
                `flex items-center gap-3 px-4 py-3 rounded-md transition 
                ${
                  isActive
                    ? "bg-indigo-100 text-indigo-700 font-medium"
                    : "text-gray-700 hover:bg-gray-100"
                }`
              }
            >
              <Icon size={18} />
              {item.name}
            </NavLink>
          );
        })}
      </nav>
    </div>
  );
}
