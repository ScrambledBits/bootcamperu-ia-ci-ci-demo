# cicd-talk-demo

App Astro mínima usada como punto de partida del demo en vivo de la charla
"El uso de la IA generativa para crear flujos de CI/CD" (Bootcamperu, mayo 2026).

## Qué hay aquí

Una app Astro vacía de CI/CD. Durante el demo, Claude generará en vivo:

- `Dockerfile` multi-stage
- `.github/workflows/ci.yml`
- `.gitlab-ci.yml`

Esos tres archivos NO existen en este repo a propósito.

## Estructura

```
.
├── README.md
├── package.json
├── astro.config.mjs
├── tsconfig.json
├── .gitignore
├── .dockerignore
├── public/
└── src/
    ├── layouts/Base.astro
    ├── components/Hello.astro
    └── pages/index.astro
```

## Setup local

Requisitos: Node 22+ y npm.

```bash
npm install
npm run dev
```

Abre http://localhost:4321 y deberías ver "Hola, Bootcamperu".

## Scripts

| Comando            | Qué hace                              |
|--------------------|---------------------------------------|
| `npm run dev`      | Levanta dev server en localhost:4321  |
| `npm run build`    | Build estático en `dist/`             |
| `npm run preview`  | Sirve el build de `dist/`             |

## Antes del primer push

Configura los secrets y permisos en GitHub y GitLab. La guía completa está en
`presenter-materials/SECRETS.md` (fuera de este repo). Sin esa configuración,
el job `publish` del workflow falla con errores de auth.

## Licencia

Material didáctico de Bootcamperu. Uso libre con atribución.
