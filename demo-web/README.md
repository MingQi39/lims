# demo-web

从 [LIMS](../lims-electron) 提炼的个人全栈 Demo 前端：React + Vite + TanStack Query + Zustand，对接 `demo-server`。

## 功能

- 登录页（JWT 存 localStorage）
- 笔记列表：搜索、分页、新建、编辑（`updated_at` 乐观锁）、删除

## 快速开始

```bash
cp .env.example .env
pnpm install
pnpm dev
```

浏览器打开 Vite 提示的地址，使用 `demo` / `demo123` 登录。

## 与完整 LIMS 的关系

不含 Electron、权限矩阵、可配置表格、Excel 编辑器等。完整 UI 见 `lims-electron`。
