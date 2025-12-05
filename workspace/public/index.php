<?php
// Config data
$appName   = "Evergon";
$tagline   = "Local Web Dev Environment, Reimagined.";
$subtitle  = "Multi-project. Multi-runtime. Zero-config. Untuk developer yang ingin fokus ngoding, bukan ngurus server.";
$ctaMain   = "Download Evergon";
$ctaAlt    = "Lihat Dokumentasi";
$version   = "v1.0.0";

$navLinks = [
    "Features" => "#features",
    "How it works" => "#workflow",
    "Stack" => "#stack",
    "Use cases" => "#usecases",
    "FAQ" => "#faq",
];

$features = [
    [
        "title" => "Multi-project ready",
        "desc"  => "Scan dan daftarkan banyak project sekaligus, Evergon akan urus port, host, dan routing secara otomatis."
    ],
    [
        "title" => "Multi-PHP runtime",
        "desc"  => "Jalankan project PHP 7.4, 8.0, 8.1, 8.2, 8.3 berdampingan tanpa konflik konfigurasi global."
    ],
    [
        "title" => "Nginx + PHP-FPM orchestration",
        "desc"  => "Evergon mengatur Nginx dan PHP-FPM, regenerate vhost, dan menghubungkan project ke runtime yang tepat."
    ],
    [
        "title" => "Per-project engine",
        "desc"  => "Set stack per project: PHP, Node, Go, atau mix. Satu project, satu konfigurasi yang jelas."
    ],
    [
        "title" => "Zero manual path config",
        "desc"  => "Tidak perlu lagi utak-atik path di config. Evergon punya resolver yang mengerti struktur workspace kamu."
    ],
    [
        "title" => "Realtime status panel",
        "desc"  => "Pantau status service (Nginx, PHP-FPM, database) dan project langsung dari panel modern."
    ],
];

$workflow = [
    "Install Evergon" => "Extract dan jalankan Evergon engine. Workspace dan folder penting akan dibuat otomatis.",
    "Tambahkan PHP runtime" => "Taruh PHP portable di folder php_versions, Evergon akan mendeteksi dan menyiapkan mapping.",
    "Scan workspace" => "Point workspace ke folder projekmu. Evergon akan memindai dan mendaftarkan tiap project.",
    "Set stack per project" => "Tentukan runtime untuk tiap project: versi PHP, port, dan base URL semuanya otomatis.",
    "Start services" => "Nyalakan Nginx, PHP-FPM, database, lalu akses project lewat domain lokal yang sudah disiapkan."
];

$stack = [
    "Core engine" => "Go-based HTTP engine untuk mengatur proses, port, dan komunikasi dengan panel.",
    "Web server" => "Nginx portable dengan auto vhost generator untuk tiap project.",
    "PHP runtime" => "Beberapa versi PHP portable (7.4+ sampai 8.x) dalam satu folder terstruktur.",
    "Database" => "Integrasi siap pakai untuk MySQL/MariaDB, Postgres, dan lain-lain (via service terpisah).",
    "Panel" => "Dashboard React/Tailwind atau HTML statis yang berkomunikasi ke engine via API.",
];

$useCases = [
    "Fullstack PHP" => "Cocok untuk Laravel, CodeIgniter, WordPress, atau project monolith yang butuh banyak versi PHP.",
    "API + SPA" => "Pisahkan backend (PHP/Go) dan frontend (React/Vue) per project dengan port dan domain lokal berbeda.",
    "Team onboarding" => "Developer baru cukup clone repo dan jalankan Evergon, tanpa harus setup manual environment.",
    "Experiment sandbox" => "Coba runtime baru tanpa merusak konfigurasi server global di sistem operasi kamu.",
];

$faq = [
    [
        "q" => "Evergon itu apa, sebenarnya?",
        "a" => "Evergon adalah local web development environment yang mengelola Nginx, PHP-FPM, multi-PHP runtime, dan multi-project dalam satu paket, tanpa ketergantungan Docker."
    ],
    [
        "q" => "Apakah saya tetap bisa pakai Docker?",
        "a" => "Bisa. Evergon tidak menggantikan Docker sepenuhnya. Evergon cocok untuk developer yang ingin setup cepat tanpa container, atau untuk project yang kurang cocok dijalankan di Docker."
    ],
    [
        "q" => "Apakah Evergon hanya untuk PHP?",
        "a" => "Fokus awalnya PHP, tapi arsitektur Evergon dirancang agar bisa menambahkan engine lain (Go, Node, dll) secara per project."
    ],
    [
        "q" => "Apakah butuh instalasi rumit?",
        "a" => "Tidak. Cukup extract, jalankan engine, dan buka panel. Resolver akan membuat folder kerja dan konfigurasi dasar untukmu."
    ],
];
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title><?= htmlspecialchars($appName) ?> – Local Web Dev Environment</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="<?= htmlspecialchars($subtitle) ?>">
    <style>
        :root {
            --bg: #050816;
            --bg-soft: #070b1f;
            --bg-soft-2: #0b1024;
            --accent: #4f46e5;
            --accent-soft: rgba(79, 70, 229, 0.15);
            --accent-strong: #a855f7;
            --text: #f9fafb;
            --text-soft: #a5b4fc;
            --text-muted: #9ca3af;
            --border-subtle: rgba(148, 163, 184, 0.25);
            --radius-lg: 24px;
            --radius-md: 18px;
            --radius-pill: 999px;
            --shadow-soft: 0 18px 60px rgba(15, 23, 42, 0.9);
            --shadow-soft-2: 0 20px 80px rgba(15, 23, 42, 0.9);
            --blur-card: blur(18px);
            --gradient-hero: radial-gradient(circle at top, #4f46e580, transparent 55%), radial-gradient(circle at bottom, #0f766e4d, transparent 55%), linear-gradient(135deg, #020617, #020617);
        }

        * { box-sizing: border-box; }

        body {
            margin: 0;
            min-height: 100vh;
            font-family: system-ui, -apple-system, BlinkMacSystemFont, "Inter", sans-serif;
            background: radial-gradient(circle at top, #1e293b 0, #020617 60%);
            color: var(--text);
        }

        a { text-decoration: none; color: inherit; }

        .page {
            min-height: 100vh;
            color: var(--text);
        }

        .noise-overlay {
            pointer-events: none;
            position: fixed;
            inset: 0;
            opacity: 0.18;
            mix-blend-mode: soft-light;
            background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 1600 800' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3' stitchTiles='noStitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.35'/%3E%3C/svg%3E");
            z-index: -1;
        }

        .max-w-6xl {
            max-width: 76rem;
            margin: 0 auto;
            padding: 0 1.5rem;
        }

        header {
            position: sticky;
            top: 0;
            z-index: 40;
            backdrop-filter: blur(14px);
            background: linear-gradient(to bottom, rgba(15, 23, 42, 0.92), rgba(15, 23, 42, 0.6), transparent);
            border-bottom: 1px solid rgba(15, 23, 42, 0.4);
        }

        .nav {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 1rem 0;
            gap: 1.5rem;
        }

        .brand {
            display: flex;
            align-items: center;
            gap: 0.75rem;
        }

        .brand-logo {
            width: 40px;
            height: 40px;
            border-radius: 14px;
            background: radial-gradient(circle at 10% 0, #a855f7 0, #4f46e5 35%, #22c55e 100%);
            position: relative;
            overflow: hidden;
            box-shadow: 0 10px 35px rgba(79, 70, 229, 0.7);
        }

        .brand-logo-orbit {
            position: absolute;
            inset: 5px;
            border-radius: 999px;
            border: 1px solid rgba(15, 23, 42, 0.65);
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .brand-logo-core {
            width: 14px;
            height: 14px;
            border-radius: 999px;
            border: 2px solid rgba(15, 23, 42, 0.9);
            box-shadow: 0 0 0 3px rgba(15, 23, 42, 0.35);
            background: conic-gradient(from 0deg, #e5e7eb, #a5b4fc, #22c55e, #e5e7eb);
        }

        .brand-text-main {
            font-weight: 600;
            letter-spacing: 0.02em;
            font-size: 1.25rem;
        }

        .brand-text-sub {
            font-size: 0.75rem;
            color: var(--text-muted);
        }

        .nav-links {
            display: flex;
            align-items: center;
            gap: 1rem;
            font-size: 0.9rem;
            color: var(--text-muted);
        }

        .nav-links a {
            padding: 0.4rem 0.75rem;
            border-radius: 999px;
            border: 1px solid transparent;
        }

        .nav-links a:hover {
            color: var(--text);
            border-color: rgba(148, 163, 184, 0.45);
            background: radial-gradient(circle at top, rgba(79, 70, 229, 0.26), transparent);
        }

        .nav-cta {
            display: flex;
            align-items: center;
            gap: 0.75rem;
        }

        .btn {
            display: inline-flex;
            align-items: center;
            justify-content: center;
            gap: 0.4rem;
            padding: 0.6rem 1.1rem;
            border-radius: var(--radius-pill);
            font-size: 0.85rem;
            font-weight: 500;
            border: 1px solid transparent;
            cursor: pointer;
            transition: transform 0.13s ease, box-shadow 0.13s ease, background 0.13s ease, border-color 0.13s ease;
            background: transparent;
            color: var(--text);
        }

        .btn-primary {
            background: linear-gradient(135deg, var(--accent), var(--accent-strong));
            border-color: rgba(59, 130, 246, 0.4);
            box-shadow: 0 12px 35px rgba(79, 70, 229, 0.7);
        }

        .btn-primary:hover {
            transform: translateY(-1px);
            box-shadow: 0 16px 45px rgba(79, 70, 229, 0.85);
        }

        .btn-ghost {
            border-color: rgba(148, 163, 184, 0.5);
            background: radial-gradient(circle at top left, rgba(148, 163, 184, 0.35), transparent);
            color: var(--text-soft);
        }

        .btn-ghost:hover {
            background: radial-gradient(circle at top, rgba(79, 70, 229, 0.32), rgba(15, 23, 42, 0.96));
            color: var(--text);
        }

        .badge {
            display: inline-flex;
            align-items: center;
            gap: 0.4rem;
            padding: 0.25rem 0.7rem;
            border-radius: 999px;
            font-size: 0.7rem;
            color: var(--text-soft);
            border: 1px solid rgba(148, 163, 184, 0.4);
            background: radial-gradient(circle at top, rgba(79, 70, 229, 0.28), rgba(15, 23, 42, 0.95));
        }

        .badge-dot {
            width: 7px;
            height: 7px;
            border-radius: 999px;
            background: radial-gradient(circle, #22c55e, #16a34a);
            box-shadow: 0 0 0 4px rgba(34, 197, 94, 0.25);
        }

        main {
            padding-bottom: 5rem;
        }

        .hero {
            background: var(--gradient-hero);
            padding: 3.5rem 0 3rem;
        }

        .hero-grid {
            display: grid;
            grid-template-columns: minmax(0, 1.1fr) minmax(0, 1fr);
            gap: 3rem;
            align-items: center;
        }

        .hero-title {
            font-size: clamp(2.5rem, 4vw, 3.25rem);
            line-height: 1.05;
            letter-spacing: -0.05em;
            margin: 1rem 0 0.75rem;
        }

        .hero-title span {
            background: linear-gradient(135deg, #a5b4fc, #c4b5fd, #22c55e);
            -webkit-background-clip: text;
            background-clip: text;
            color: transparent;
        }

        .hero-subtitle {
            max-width: 34rem;
            font-size: 0.98rem;
            color: var(--text-muted);
            line-height: 1.6;
        }

        .hero-actions {
            display: flex;
            flex-wrap: wrap;
            gap: 0.85rem;
            margin: 1.75rem 0 1rem;
        }

        .hero-meta {
            display: flex;
            flex-wrap: wrap;
            gap: 1.5rem;
            font-size: 0.75rem;
            color: var(--text-muted);
        }

        .hero-meta-item {
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }

        .hero-meta-dot {
            width: 16px;
            height: 16px;
            border-radius: 999px;
            border: 1px solid rgba(148, 163, 184, 0.5);
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 0.55rem;
        }

        .hero-card {
            position: relative;
            border-radius: var(--radius-lg);
            padding: 1.4rem;
            background: radial-gradient(circle at top left, rgba(79, 70, 229, 0.23), rgba(15, 23, 42, 0.98));
            border: 1px solid rgba(148, 163, 184, 0.3);
            box-shadow: var(--shadow-soft-2);
            overflow: hidden;
        }

        .hero-card-grid {
            display: grid;
            grid-template-columns: 1.1fr 0.9fr;
            gap: 1.1rem;
        }

        .hero-card-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 0.75rem;
            margin-bottom: 0.9rem;
        }

        .hero-card-title {
            font-size: 0.9rem;
            font-weight: 500;
            color: var(--text-soft);
        }

        .chip {
            display: inline-flex;
            align-items: center;
            gap: 0.4rem;
            padding: 0.25rem 0.75rem;
            border-radius: 999px;
            font-size: 0.7rem;
            border: 1px solid rgba(148, 163, 184, 0.4);
            background: rgba(15, 23, 42, 0.8);
        }

        .chip-dot {
            width: 7px;
            height: 7px;
            border-radius: 999px;
        }

        .chip-dot.ok {
            background: #22c55e;
            box-shadow: 0 0 0 4px rgba(34, 197, 94, 0.3);
        }

        .chip-dot.warn {
            background: #f97316;
            box-shadow: 0 0 0 4px rgba(249, 115, 22, 0.26);
        }

        .chip-dot.idle {
            background: #64748b;
        }

        .hero-services {
            display: grid;
            grid-template-columns: repeat(3, minmax(0, 1fr));
            gap: 0.7rem;
            font-size: 0.75rem;
        }

        .service-pill {
            padding: 0.55rem 0.6rem;
            border-radius: 14px;
            border: 1px solid rgba(148, 163, 184, 0.4);
            background: radial-gradient(circle at top, rgba(15, 23, 42, 0.85), rgba(15, 23, 42, 0.98));
        }

        .service-pill-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 0.35rem;
        }

        .service-pill-name {
            font-size: 0.75rem;
        }

        .status-dot {
            width: 8px;
            height: 8px;
            border-radius: 999px;
        }

        .status-dot.running {
            background: #22c55e;
            box-shadow: 0 0 0 3px rgba(34, 197, 94, 0.25);
        }

        .status-dot.stopped {
            background: #ef4444;
            box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.25);
        }

        .service-pill-meta {
            display: flex;
            justify-content: space-between;
            font-size: 0.7rem;
            color: var(--text-muted);
        }

        .hero-card-activity {
            font-size: 0.75rem;
            border-radius: 16px;
            padding: 0.75rem;
            border: 1px dashed rgba(148, 163, 184, 0.5);
            background: radial-gradient(circle at top left, rgba(15, 23, 42, 0.75), rgba(15, 23, 42, 0.98));
        }

        .hero-card-activity-lines {
            margin: 0.35rem 0;
            font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
            font-size: 0.7rem;
            color: #e5e7eb;
            line-height: 1.5;
        }

        .hero-card-activity-lines span.ok { color: #22c55e; }
        .hero-card-activity-lines span.info { color: #38bdf8; }

        .hero-badge-strip {
            margin-top: 1rem;
            display: flex;
            flex-wrap: wrap;
            gap: 0.5rem;
            font-size: 0.7rem;
            color: var(--text-muted);
        }

        .hero-badge-strip span {
            padding: 0.35rem 0.7rem;
            border-radius: 999px;
            border: 1px solid rgba(148, 163, 184, 0.4);
            background: rgba(15, 23, 42, 0.9);
        }

        .section {
            padding: 3.25rem 0 0;
        }

        .section-header {
            display: flex;
            align-items: flex-end;
            justify-content: space-between;
            gap: 1.5rem;
            margin-bottom: 2rem;
        }

        .section-title {
            font-size: 1.35rem;
            font-weight: 600;
        }

        .section-kicker {
            font-size: 0.75rem;
            letter-spacing: 0.18em;
            text-transform: uppercase;
            color: var(--text-muted);
        }

        .section-subtitle {
            font-size: 0.9rem;
            color: var(--text-muted);
            max-width: 32rem;
        }

        .pill-row {
            display: flex;
            flex-wrap: wrap;
            gap: 0.65rem;
            font-size: 0.75rem;
        }

        .pill-row span {
            padding: 0.4rem 0.7rem;
            border-radius: 999px;
            border: 1px solid rgba(148, 163, 184, 0.4);
            color: var(--text-muted);
        }

        .features-grid {
            display: grid;
            grid-template-columns: repeat(3, minmax(0, 1fr));
            gap: 1rem;
        }

        .feature-card {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px solid rgba(148, 163, 184, 0.35);
            background: radial-gradient(circle at top left, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            box-shadow: var(--shadow-soft);
            font-size: 0.9rem;
        }

        .feature-icon {
            width: 28px;
            height: 28px;
            border-radius: 999px;
            display: inline-flex;
            align-items: center;
            justify-content: center;
            font-size: 0.9rem;
            margin-bottom: 0.6rem;
            background: radial-gradient(circle at top left, rgba(79, 70, 229, 0.8), rgba(139, 92, 246, 0.7));
            box-shadow: 0 12px 25px rgba(79, 70, 229, 0.65);
        }

        .feature-title {
            font-weight: 500;
            margin-bottom: 0.35rem;
        }

        .feature-desc {
            color: var(--text-muted);
            font-size: 0.83rem;
            line-height: 1.55;
        }

        .workflow-grid {
            display: grid;
            grid-template-columns: 1.15fr 1fr;
            gap: 1.6rem;
        }

        .timeline {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px solid rgba(148, 163, 184, 0.4);
            background: radial-gradient(circle at top, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            box-shadow: var(--shadow-soft);
        }

        .timeline-step {
            display: grid;
            grid-template-columns: auto minmax(0, 1fr);
            gap: 0.75rem;
            padding: 0.65rem 0;
        }

        .timeline-step:not(:last-child) {
            border-bottom: 1px dashed rgba(148, 163, 184, 0.4);
        }

        .step-index {
            width: 26px;
            height: 26px;
            border-radius: 999px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 0.75rem;
            font-weight: 500;
            background: radial-gradient(circle at top, rgba(79, 70, 229, 0.9), rgba(79, 70, 229, 0.6));
            box-shadow: 0 10px 25px rgba(79, 70, 229, 0.7);
        }

        .step-title {
            font-size: 0.9rem;
            font-weight: 500;
        }

        .step-desc {
            font-size: 0.8rem;
            color: var(--text-muted);
            margin-top: 0.2rem;
        }

        .workflow-aside {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px dashed rgba(148, 163, 184, 0.5);
            background: radial-gradient(circle at top, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            font-size: 0.8rem;
            color: var(--text-muted);
        }

        .stack-grid {
            display: grid;
            grid-template-columns: repeat(2, minmax(0, 1fr));
            gap: 1rem;
        }

        .stack-card {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px solid rgba(148, 163, 184, 0.35);
            background: radial-gradient(circle at top left, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            font-size: 0.85rem;
        }

        .stack-title {
            font-weight: 500;
            margin-bottom: 0.35rem;
        }

        .stack-desc {
            color: var(--text-muted);
        }

        .usecases-grid {
            display: grid;
            grid-template-columns: repeat(2, minmax(0, 1fr));
            gap: 1rem;
        }

        .usecase-card {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px solid rgba(148, 163, 184, 0.35);
            background: radial-gradient(circle at top left, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            font-size: 0.85rem;
        }

        .usecase-title {
            font-weight: 500;
            margin-bottom: 0.35rem;
        }

        .usecase-desc {
            color: var(--text-muted);
        }

        .faq-grid {
            display: grid;
            grid-template-columns: repeat(2, minmax(0, 1fr));
            gap: 1rem;
        }

        .faq-item {
            border-radius: var(--radius-md);
            padding: 1rem;
            border: 1px solid rgba(148, 163, 184, 0.35);
            background: radial-gradient(circle at top left, rgba(15, 23, 42, 0.9), rgba(15, 23, 42, 0.98));
            font-size: 0.85rem;
        }

        .faq-q {
            font-weight: 500;
            margin-bottom: 0.3rem;
        }

        .faq-a {
            color: var(--text-muted);
        }

        .cta-strip {
            margin-top: 3.5rem;
            border-radius: var(--radius-lg);
            padding: 1.25rem 1.35rem;
            border: 1px solid rgba(148, 163, 184, 0.45);
            background: radial-gradient(circle at left, rgba(79, 70, 229, 0.4), rgba(15, 23, 42, 0.98));
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 1.5rem;
            box-shadow: 0 18px 55px rgba(15, 23, 42, 0.9);
        }

        .cta-strip-title {
            font-size: 0.95rem;
            font-weight: 500;
        }

        .cta-strip-sub {
            font-size: 0.8rem;
            color: var(--text-muted);
        }

        footer {
            padding: 1.75rem 0 2.25rem;
            font-size: 0.75rem;
            color: var(--text-muted);
        }

        .footer-row {
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 1rem;
            border-top: 1px solid rgba(15, 23, 42, 0.85);
            padding-top: 1.25rem;
        }

        .footer-links {
            display: flex;
            flex-wrap: wrap;
            gap: 0.7rem;
        }

        .footer-links a {
            color: var(--text-muted);
        }

        .footer-links a:hover {
            color: var(--text-soft);
        }

        @media (max-width: 960px) {
            .hero-grid {
                grid-template-columns: minmax(0, 1fr);
            }
            .hero-card {
                margin-top: 2rem;
            }
            .features-grid {
                grid-template-columns: repeat(2, minmax(0, 1fr));
            }
            .workflow-grid {
                grid-template-columns: minmax(0, 1fr);
            }
            .stack-grid,
            .usecases-grid,
            .faq-grid {
                grid-template-columns: minmax(0, 1fr);
            }
            header {
                backdrop-filter: blur(10px);
            }
            .nav-links {
                display: none;
            }
        }

        @media (max-width: 640px) {
            .features-grid {
                grid-template-columns: minmax(0, 1fr);
            }
            .cta-strip {
                flex-direction: column;
                align-items: flex-start;
            }
            .hero-actions {
                flex-direction: column;
                align-items: flex-start;
            }
            .hero-card-grid {
                grid-template-columns: minmax(0, 1fr);
            }
        }
    </style>
</head>
<body>
<div class="page">
    <div class="noise-overlay"></div>

    <header>
        <div class="max-w-6xl">
            <nav class="nav">
                <div class="brand">
                    <div class="brand-logo">
                        <div class="brand-logo-orbit">
                            <div class="brand-logo-core"></div>
                        </div>
                    </div>
                    <div>
                        <div class="brand-text-main"><?= htmlspecialchars($appName) ?></div>
                        <div class="brand-text-sub">Local Dev Environment</div>
                    </div>
                </div>
                <div class="nav-links">
                    <?php foreach ($navLinks as $label => $href): ?>
                        <a href="<?= $href ?>"><?= htmlspecialchars($label) ?></a>
                    <?php endforeach; ?>
                </div>
                <div class="nav-cta">
                    <span style="font-size:0.7rem;color:var(--text-muted);">Current build <?= htmlspecialchars($version) ?></span>
                    <a href="#download" class="btn btn-primary">
                        <span>Download</span>
                        <span style="font-size:0.9rem;">↓</span>
                    </a>
                </div>
            </nav>
        </div>
    </header>

    <main>
        <section class="hero">
            <div class="max-w-6xl hero-grid">
                <div>
                    <div class="badge">
                        <span class="badge-dot"></span>
                        <span>Engine-controlled Nginx + PHP-FPM</span>
                    </div>

                    <h1 class="hero-title">
                        <?= htmlspecialchars($appName) ?> is how<br>
                        <span>local web dev should feel.</span>
                    </h1>

                    <p class="hero-subtitle">
                        <?= htmlspecialchars($subtitle) ?>
                    </p>

                    <div class="hero-actions">
                        <a href="#download" class="btn btn-primary">
                            <span><?= htmlspecialchars($ctaMain) ?></span>
                            <span style="font-size:0.9rem;">⇣</span>
                        </a>
                        <a href="#stack" class="btn btn-ghost">
                            <span><?= htmlspecialchars($ctaAlt) ?></span>
                        </a>
                    </div>

                    <div class="hero-meta">
                        <div class="hero-meta-item">
                            <div class="hero-meta-dot">∞</div>
                            <span>Multi-project, multi-runtime</span>
                        </div>
                        <div class="hero-meta-item">
                            <div class="hero-meta-dot">0</div>
                            <span>Global config editing</span>
                        </div>
                        <div class="hero-meta-item">
                            <div class="hero-meta-dot">✓</div>
                            <span>Ready for PHP dev teams</span>
                        </div>
                    </div>
                </div>

                <div>
                    <div class="hero-card">
                        <div class="hero-card-grid">
                            <div>
                                <div class="hero-card-header">
                                    <div class="hero-card-title">Evergon Engine · Services</div>
                                    <div class="chip">
                                        <span class="chip-dot ok"></span>
                                        <span>All services healthy</span>
                                    </div>
                                </div>
                                <div class="hero-services">
                                    <div class="service-pill">
                                        <div class="service-pill-header">
                                            <span class="service-pill-name">Nginx</span>
                                            <span class="status-dot running"></span>
                                        </div>
                                        <div class="service-pill-meta">
                                            <span>HTTP: 8080</span>
                                            <span>config: /nginx/conf/</span>
                                        </div>
                                    </div>
                                    <div class="service-pill">
                                        <div class="service-pill-header">
                                            <span class="service-pill-name">PHP-FPM 8.3</span>
                                            <span class="status-dot running"></span>
                                        </div>
                                        <div class="service-pill-meta">
                                            <span>pool: evergon</span>
                                            <span>port: 9083</span>
                                        </div>
                                    </div>
                                    <div class="service-pill">
                                        <div class="service-pill-header">
                                            <span class="service-pill-name">MySQL</span>
                                            <span class="status-dot stopped"></span>
                                        </div>
                                        <div class="service-pill-meta">
                                            <span>status: stopped</span>
                                            <span>click to start</span>
                                        </div>
                                    </div>
                                </div>
                                <div class="hero-badge-strip">
                                    <span>Auto vhost generation</span>
                                    <span>Per-project PHP version</span>
                                    <span>Port conflict safety</span>
                                </div>
                            </div>

                            <div>
                                <div class="hero-card-activity">
                                    <div style="display:flex;align-items:center;justify-content:space-between;gap:0.5rem;">
                                        <span>Recent engine log</span>
                                        <span class="chip">
                                            <span class="chip-dot idle"></span>
                                            <span>idle</span>
                                        </span>
                                    </div>
                                    <div class="hero-card-activity-lines">
                                        <div><span class="ok">[ok]</span> nginx started at <span class="info">:8080</span></div>
                                        <div><span class="ok">[ok]</span> php-fpm[83] listening on <span class="info">:9083</span></div>
                                        <div><span class="ok">[ok]</span> testapp registered at <span class="info">http://testapp.local</span></div>
                                        <div><span class="ok">[ok]</span> auto-resolved root: <span class="info">workspace/www/testapp</span></div>
                                    </div>
                                    <div style="display:flex;justify-content:flex-end;margin-top:0.4rem;font-size:0.7rem;color:var(--text-muted);">
                                        <span>engine uptime: 02h 31m</span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div style="position:absolute;inset:0;pointer-events:none;mix-blend-mode:screen;opacity:0.28;">
                            <div style="position:absolute;right:-40px;top:-40px;width:180px;height:180px;border-radius:999px;background:radial-gradient(circle,#4f46e5,transparent);"></div>
                            <div style="position:absolute;left:-60px;bottom:-60px;width:220px;height:220px;border-radius:999px;background:radial-gradient(circle,#22c55e,transparent);"></div>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section id="features" class="section">
            <div class="max-w-6xl">
                <div class="section-header">
                    <div>
                        <div class="section-kicker">Why <?= htmlspecialchars($appName) ?></div>
                        <div class="section-title">Built for real-world PHP workflows.</div>
                    </div>
                    <p class="section-subtitle">
                        Evergon lahir dari rasa lelah mengatur server lokal yang rapuh. Ia menggantikan
                        tumpukan config manual dengan engine yang mengerti cara developer bekerja.
                    </p>
                </div>

                <div class="features-grid">
                    <?php foreach ($features as $index => $feature): ?>
                        <article class="feature-card">
                            <div class="feature-icon">
                                <?= $index + 1 ?>
                            </div>
                            <div class="feature-title"><?= htmlspecialchars($feature['title']) ?></div>
                            <p class="feature-desc"><?= htmlspecialchars($feature['desc']) ?></p>
                        </article>
                    <?php endforeach; ?>
                </div>
            </div>
        </section>

        <section id="workflow" class="section">
            <div class="max-w-6xl workflow-grid">
                <div>
                    <div class="section-header">
                        <div>
                            <div class="section-kicker">How it works</div>
                            <div class="section-title">From zero to running stack in minutes.</div>
                        </div>
                    </div>

                    <div class="timeline">
                        <?php $i = 1; ?>
                        <?php foreach ($workflow as $title => $desc): ?>
                            <div class="timeline-step">
                                <div class="step-index"><?= $i ?></div>
                                <div>
                                    <div class="step-title"><?= htmlspecialchars($title) ?></div>
                                    <div class="step-desc"><?= htmlspecialchars($desc) ?></div>
                                </div>
                            </div>
                            <?php $i++; ?>
                        <?php endforeach; ?>
                    </div>
                </div>

                <aside class="workflow-aside">
                    <div style="font-weight:500;margin-bottom:0.4rem;">Opinionated, but on your side.</div>
                    <p style="margin:0 0 0.7rem;">
                        Evergon sengaja membatasi model konfigurasi: satu project, satu stack. Tidak ada
                        konfigurasi global yang abu-abu. Ini membuat panel dan engine bisa
                        memecahkan masalah port, vhost, dan runtime tanpa kejutan.
                    </p>
                    <p style="margin:0 0 0.7rem;">
                        Hasilnya: onboarding lebih cepat, context lebih jelas, dan environment
                        yang bisa dipindah dengan mudah ke laptop lain dalam satu folder terstruktur.
                    </p>
                    <div class="pill-row">
                        <span>Portable-first</span>
                        <span>Engine-managed runtime</span>
                        <span>Config as data</span>
                    </div>
                </aside>
            </div>
        </section>

        <section id="stack" class="section">
            <div class="max-w-6xl">
                <div class="section-header">
                    <div>
                        <div class="section-kicker">Under the hood</div>
                        <div class="section-title">A small engine that orchestrates a serious stack.</div>
                    </div>
                    <p class="section-subtitle">
                        Evergon tidak mencoba menjadi hypervisor. Ia hanya ingin memastikan bahwa
                        setiap service yang kamu perlu untuk development bisa start, stop, dan dipantau dengan jelas.
                    </p>
                </div>

                <div class="stack-grid">
                    <?php foreach ($stack as $title => $desc): ?>
                        <article class="stack-card">
                            <div class="stack-title"><?= htmlspecialchars($title) ?></div>
                            <p class="stack-desc"><?= htmlspecialchars($desc) ?></p>
                        </article>
                    <?php endforeach; ?>
                </div>
            </div>
        </section>

        <section id="usecases" class="section">
            <div class="max-w-6xl">
                <div class="section-header">
                    <div>
                        <div class="section-kicker">Use cases</div>
                        <div class="section-title">Designed for the setups you actually run.</div>
                    </div>
                </div>

                <div class="usecases-grid">
                    <?php foreach ($useCases as $title => $desc): ?>
                        <article class="usecase-card">
                            <div class="usecase-title"><?= htmlspecialchars($title) ?></div>
                            <p class="usecase-desc"><?= htmlspecialchars($desc) ?></p>
                        </article>
                    <?php endforeach; ?>
                </div>
            </div>
        </section>

        <section id="faq" class="section">
            <div class="max-w-6xl">
                <div class="section-header">
                    <div>
                        <div class="section-kicker">FAQ</div>
                        <div class="section-title">Questions you might ask before switching.</div>
                    </div>
                </div>

                <div class="faq-grid">
                    <?php foreach ($faq as $item): ?>
                        <article class="faq-item">
                            <div class="faq-q"><?= htmlspecialchars($item['q']) ?></div>
                            <p class="faq-a"><?= htmlspecialchars($item['a']) ?></p>
                        </article>
                    <?php endforeach; ?>
                </div>

                <div id="download" class="cta-strip">
                    <div>
                        <div class="cta-strip-title">Ready to try <?= htmlspecialchars($appName) ?>?</div>
                        <div class="cta-strip-sub">
                            Download build terbaru dan jalankan di workspace kamu. Tidak ada installer agresif,
                            hanya satu folder yang bisa kamu bawa ke mana saja.
                        </div>
                    </div>
                    <div style="display:flex;flex-wrap:wrap;gap:0.75rem;">
                        <a href="#" class="btn btn-primary">
                            <span>Download for Windows</span>
                            <span style="font-size:0.9rem;">.zip</span>
                        </a>
                        <a href="#" class="btn btn-ghost">
                            <span>Download for Linux</span>
                            <span style="font-size:0.9rem;">.tar.gz</span>
                        </a>
                    </div>
                </div>
            </div>
        </section>
    </main>

    <footer>
        <div class="max-w-6xl footer-row">
            <div>© <?= date('Y') ?> <?= htmlspecialchars($appName) ?>. Built for local web developers.</div>
            <div class="footer-links">
                <a href="#features">Features</a>
                <a href="#stack">Stack</a>
                <a href="#faq">FAQ</a>
            </div>
        </div>
    </footer>
</div>
</body>
</html>
