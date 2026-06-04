# demo-server

Go API for the [LIMS Stack Demo](../README.md): Gin + GORM + SQLite, with Handler → Service → Repository layering and `/api/v1` response envelope.

## Endpoints

- `POST /api/v1/auth/login` — JWT sign-in  
- `GET/POST /api/v1/notes` — list and create notes  
- `GET/PUT/DELETE /api/v1/notes/:id` — read, update (`updated_at` optimistic lock), delete  

## Run locally

```bash
cp .env.example .env
go run ./cmd/server
```

Default user: `demo` / `demo123` (seeded on first startup).

See `.env.example` for configuration.
