# O&M Platform (MVP Scaffold)

This repository provides a front-end/back-end separated scaffold for an O&M platform.

## Current scope

- Vue 3 + TypeScript admin UI shell with route map for all MVP pages.
- Go HTTP API server with in-memory storage for:
  - Hosts
  - Packs
  - Projects
  - Graph save/load/validate
  - Install config save/load/validate
  - Generated output preview
- Graph validation rules aligned with topology constraints:
  - nginx -> frontend/backend
  - frontend -> backend
  - backend -> db/redis

## Run

### Server

```bash
cd server
go run ./cmd/om-server
```

### Admin UI

```bash
cd web/admin-ui
npm install
npm run dev
```

## API examples

```bash
curl -X POST http://localhost:8081/api/projects \
  -H 'Content-Type: application/json' \
  -d '{"id":"customer-a-prod","name":"Customer A Production","environment":"prod","description":"demo"}'
```

```bash
curl -X PUT http://localhost:8081/api/projects/customer-a-prod/graph \
  -H 'Content-Type: application/json' \
  -d '{"nodes":[],"edges":[]}'
```
