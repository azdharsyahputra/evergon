# Evergon (v0.1 – Foundation Release)

Evergon is a modern, portable, and developer-centric local web environment — aiming to provide a clean and modular alternative to tools like Laragon.  
**Version 0.1 marks the first foundation release**, where the core engine is built and the system architecture is established.

This version focuses on engine-level infrastructure.  
Future versions will include the full panel, PHP runtime integration, additional services, and installer tooling.

---

## 🚧 Project Status — v0.1 (Early Foundation)

This release contains:

### ✔ Fully Implemented
- Core Go Engine  
- Internal API structure  
- Process Manager (base)  
- Nginx Manager (early implementation)  
- Nginx Portable Runtime  
- Configuration Loader  
- Basic Project Scanner  

### 🚧 In Progress
- PHP Manager (runtime switching)  
- Panel UI (React/Tailwind scaffolding only)  
- Auto vhost generator  
- Service lifecycle handling  
- Engine-panel communication  

### ❌ Not Implemented Yet (Coming Soon)
- MySQL portable runtime  
- PHP-FPM orchestration  
- SSL support  
- Full installer pipeline  
- Add-on / plugin system  

> **Evergon is not ready for production use.**  
> This version is intended for architecture validation and local development experiments.

---

# 📐 Architecture Overview

Save the image to:

```
docs/images/architecture.png
```

Then the README displays it here:

<p align="center">
  <img src="docs/images/architecture.png" width="820">
</p>

The system flows horizontally:

**Panel → Engine → Runtime (Nginx, PHP Versions, Projects)**  
Engine acts as the central orchestrator.

---

## 🧭 Project Structure

```
evergon/
  engine/
    cmd/evergon-engine/
    internal/
      api/
      process/
      manager/
      scanner/
      config/
      util/
    go.mod

  panel/
  admin/
  php_versions/
  nginx/
    build/
    portable/
  nginx_template/
  installer/
  docs/
    images/
  public/
  www/
  README.md
```

---

## 🚀 Getting Started (v0.1)

### 1. Clone Repo
```bash
git clone https://github.com/yourusername/evergon.git
cd evergon
```

### 2. Build Engine
```bash
cd engine
go build -o evergon-engine ./cmd/evergon-engine
```

### 3. Run Engine
```bash
./evergon-engine
```

### 4. Panel (not yet functional)
Panel UI structure exists but not wired to backend.

---

## 🛠 Configuration

Engine config is located here:

```
engine/internal/config/config.go
```

Handles:
- Runtime paths  
- Nginx executable  
- PHP version slots (future)  
- Project scanning rules  

---

## 🗺 Roadmap (Next Versions)

### v0.2
- Full PHP Manager  
- Panel ↔ Engine communication  
- Automatic Vhost Generator  
- Service lifecycle management  
- Basic installation script

### v0.3
- PHP-FPM integration  
- MySQL portable runtime  
- Project templates (Laravel / CI4)  
- SSL generation (self-signed)

### v0.4
- Windows installer (.exe)  
- Plugin/add-on system  
- Multi-engine option (Nginx/Caddy)

### v1.0
- Full stable release  
- Complete control panel  
- Production-ready workflows  

---

## 📄 License  
MIT License

---

## ✍ Author  
**Muhammad Azdhar Syahputra** — Creator & Lead Developer (Evergon)

