# Evergon

Evergon is a multi-component platform designed to manage environments,
versions, and web services with an extensible architecture. The project
is organized into several core modules:

## Project Structure

    evergon/
      engine/
        cmd/evergon-engine/main.go
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
      nginx_template/
      installer/
      docs/

## Overview

### Engine

The core backend responsible for: - Environment scanning - Process
management - Service orchestration - Configuration handling - Exposing
internal APIs

### Panel

Frontend panel for interacting with the engine, managing services, and
monitoring system state.

### Admin

Administrative utilities and configurations.

### PHP Versions

Pre-packaged PHP versions used for environment provisioning.

### Nginx Template

Template files for generating Nginx server configurations dynamically.

### Installer

Installation system for distributing Evergon as a packaged executable.

### Docs

Documentation and internal technical notes.

## Features

-   Automated scanning of services and environments
-   Runtime process manager
-   Modular engine architecture
-   Web-based control panel
-   Dynamic Nginx configuration generation
-   Multi-version PHP support

## Requirements

-   Go 1.21+
-   Linux environment recommended
-   Admin privileges for full system access

## Setup

### 1. Clone the Repository

    git clone https://github.com/yourusername/evergon.git
    cd evergon

### 2. Build Engine

    cd engine
    go build -o evergon-engine ./cmd/evergon-engine

### 3. Run Engine

    ./evergon-engine

### 4. Access Panel

Open browser:

    http://localhost:9090

## License

MIT License

## Author

Ajar