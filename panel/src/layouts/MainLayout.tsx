import type { ReactNode } from "react";
import Sidebar from "../components/Sidebar";
import Header from "../components/Header";

export default function MainLayout({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-screen w-full bg-gray-100">
      <Sidebar />

      <div className="flex flex-col flex-1 w-full">
        <Header />
        <main className="p-6 w-full">{children}</main>
      </div>
    </div>
  );
}
