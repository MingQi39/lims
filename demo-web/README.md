# demo-web

React SPA for the [LIMS Stack Demo](../README.md): Vite + TanStack Query + Zustand, talking to `demo-server`.

## Features

- Sign-in page (JWT in localStorage)  
- Notes list: search, pagination, create, edit (`updated_at` optimistic lock), delete  

## Run locally

```bash
cp .env.example .env
pnpm install
pnpm dev
```

Open the dev server URL and sign in with `demo` / `demo123`.
