# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

Demo app for the talk "El uso de la IA generativa para crear flujos de CI/CD" (Bootcamperu, mayo 2026). It's a minimal Astro 5 static site used as the live-demo starting point. The three CI/CD artifacts — `Dockerfile`, `.github/workflows/ci.yml`, and `.gitlab-ci.yml` — are intentionally absent and are meant to be generated live during the talk.

## Commands

```bash
npm install          # install dependencies (Node 22+ required)
npm run dev          # dev server at http://localhost:4321
npm run build        # static build → dist/
npm run preview      # serve the dist/ build locally
```

There are no test or lint scripts defined.

## Architecture

Single-page Astro 5 app with `output: 'static'` and strict TypeScript. Component hierarchy:

```
pages/index.astro  →  layouts/Base.astro (HTML shell, <slot />)
                   →  components/Hello.astro (greeting, accepts name prop)
```

`astro.config.mjs` sets `site: 'https://example.com'` — update this when targeting a real deployment host.

## CI/CD Generation Context

When generating the three missing files, keep in mind:

- `.dockerignore` already excludes `node_modules`, `dist`, `.astro`, `.git`, `.github`, `.gitlab-ci.yml`, and the presenter markdown files.
- The build produces a fully static `dist/` directory — no Node runtime is needed at serve time, so a multi-stage Dockerfile (build stage → nginx/static server stage) is the right pattern.
- Secrets and registry permissions must be configured in GitHub and GitLab before the `publish` job will succeed (see `presenter-materials/SECRETS.md`, outside this repo).
