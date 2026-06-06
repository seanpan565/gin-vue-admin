---
tool: cursor
---

# 商城管理后台

## 项目结构

- `server/`：Go 后端（管理后台 `/mall/*` + C 端 `/site/auth/*`）
- `web/`：Vue 3 管理后台前端

## 开发约定

- 商城业务：`model/mall/`、`service/mall/`、`api/admin/mall/`、`api/siteapi/auth/`
- 管理路由：`router/admin/`，C 端路由：`router/siteapi/`
- 本地：MySQL `gva`、Redis、后端 `:8888`、前端 `:8080`
