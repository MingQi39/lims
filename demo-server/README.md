# demo-server

从 [LIMS](../lims-server) 提炼的个人全栈 Demo 后端：Gin + GORM + SQLite，展示 Handler → Service → Repository 分层与 `/api/v1` 统一响应。

## 功能

- `POST /api/v1/auth/login` — JWT 登录
- `GET/POST /api/v1/notes` — 笔记列表与创建
- `GET/PUT/DELETE /api/v1/notes/:id` — 详情、更新（`updated_at` 乐观锁）、删除

## 快速开始

```bash
cp .env.example .env
go run ./cmd/server
```

默认账号：`demo` / `demo123`（首次启动自动建库并种子用户）。

环境变量见 `.env.example`。

## 与完整 LIMS 的关系

本目录为**开源展示用精简版**，不含租户、权限矩阵、Worker、Excel/PDF 等能力。完整实现见仓库内 `lims-server`。
