import Logo from "../assets/evergon.png";

export default function Header() {

  return (
    <header className="h-16 px-6 bg-white/80 backdrop-blur-md border-b shadow-sm flex items-center justify-between sticky top-0 z-50">

      {/* LEFT */}
      <div className="flex items-center gap-4">
        {/* Logo */}
        <div className="w-10 h-10 rounded-xl bg-indigo-50 flex items-center justify-center shadow-sm overflow-hidden">
          <img
            src={Logo}
            alt="Evergon Logo"
            className="w-7 h-7 object-contain"
          />
        </div>

        {/* Brand */}
        <div>
          <h1 className="text-xl font-bold tracking-tight text-gray-900">
            Evergon
          </h1>
          <p className="text-xs text-gray-400 -mt-1">Local Dev Engine</p>
        </div>
      </div>

      {/* RIGHT */}
      <div className="flex items-center gap-5">

      </div>

    </header>
  );
}
