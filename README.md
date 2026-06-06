# 商城管理后台

基于 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 精简改造的本地学习/测试项目，前后端分离。原框架的代码生成、MCP、自动插件等能力已移除，业务以商城会员模块为主。

| 端 | 目录 | 说明 |
|---|---|---|
| 管理后台 API | `server/api/admin/` | JWT + Casbin，如 `/mall/member/*` |
| C 端 API | `server/api/siteapi/auth/` | 会员注册/登录，Header `x-member-token` |
| 管理后台前端 | `web/` | Vue 3 + Element Plus |
| 共享业务层 | `server/model/mall/`、`server/service/mall/` | 商城领域模型与服务 |

后端详细说明见 [server/README.md](server/README.md)。

## 环境要求

- Go >= 1.24
- Node.js >= 18
- MySQL 5.7+（InnoDB）
- Redis（建议开启：验证码、限流、多点登录）

## 配置

| 文件 | 说明 |
|------|------|
| `server/config.yaml` | 本地开发主配置 |
| `server/config.example.yaml` | 生产敏感项与环境变量对照 |
| `web/.env.development` | 前端开发环境（端口、代理） |
| `web/.env.production` | 前端生产构建（API 前缀、错误上报密钥） |

本地默认数据库（`server/config.yaml`）：

- MySQL 库名 `gva`，用户 `root` / `123456`
- Redis `127.0.0.1:6379`

生产环境请用环境变量覆盖密钥，参见下方 [安全与监控](#安全与监控)。

## 本地启动

### 1. 准备数据库

```bash
mysql -uroot -p123456 -e "CREATE DATABASE IF NOT EXISTS gva DEFAULT CHARSET utf8mb4;"
```

确保 MySQL、Redis 已启动。首次访问若库未初始化，打开 `http://127.0.0.1:8080` 会进入初始化向导。

### 2. 启动后端

```bash
cd server
go mod tidy
go run .
```

| 项 | 地址 |
|---|---|
| API | http://127.0.0.1:8888 |
| Swagger | http://127.0.0.1:8888/swagger/index.html（`security.swagger-enable: true` 时） |
| 健康检查 | http://127.0.0.1:8888/health |
| Prometheus | http://127.0.0.1:8888/metrics（`security.metrics-enable: true` 时） |

### 3. 启动前端

```bash
cd web
npm install
npm run serve
```

默认端口：**8080**（`VITE_BASE_API=/api`，由 Vite 代理到后端 8888）。

### 4. 登录管理后台

| 项 | 值 |
|---|---|
| 地址 | http://127.0.0.1:8080 |
| 账号 | `admin` |
| 密码 | `123456` |

侧边栏 **商城管理 → 会员管理** 可查看 C 端注册的会员。

## C 端接口

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/base/captcha` | 图形验证码 |
| POST | `/site/auth/register` | 注册 |
| POST | `/site/auth/login` | 登录 |
| POST | `/site/auth/logout` | 登出（服务端拉黑 token） |
| GET | `/site/auth/profile` | 个人信息（需 token） |

登录成功后，后续请求 Header 携带：`x-member-token: <token>`。

默认 `security.member-captcha-enable: true`，注册/登录需带 `captchaId`、`captcha`。本地调试可在 `config.yaml` 设为 `false`。

### curl 示例（含验证码）

```bash
# 1. 获取验证码
curl -s -X POST http://127.0.0.1:8888/base/captcha -H 'Content-Type: application/json' -d '{}'
# 响应中的 captchaId、captchaLength；验证码图片为 base64，需人工识别 captcha 字段

# 2. 注册
curl -X POST http://127.0.0.1:8888/site/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"mobile":"13800138000","password":"123456","nickname":"测试","captchaId":"<id>","captcha":"<code>"}'
```

## 常用命令

```bash
# Swagger 文档
cd server && swag init

# 前端生产构建
cd web && npm run build

# 后端编译
cd server && go build -o server .

# Makefile 容器构建（Go 1.24 镜像）
make build-server   # 仅后端
make build-web      # 仅前端
```

## 目录结构

```text
.
├── server/                     # Go 后端
│   ├── api/admin/              # 管理后台接口
│   ├── api/siteapi/auth/       # C 端会员认证
│   ├── router/admin/           # 管理路由
│   ├── router/siteapi/         # C 端路由
│   ├── model/mall/             # 商城模型
│   ├── service/mall/           # 商城服务
│   ├── middleware/             # 安全、限流、监控中间件
│   ├── source/system/          # 初始化种子（菜单、API、Casbin）
│   ├── config.yaml
│   └── README.md               # 后端专项文档
├── web/                        # 管理后台前端
│   └── src/view/mall/          # 商城页面
├── Makefile                    # Docker 构建脚本
└── mall-admin.code-workspace   # VS Code / Cursor 工作区
```

## VS Code / Cursor 工作区

打开 `mall-admin.code-workspace`，使用 **Both (Backend & Frontend)** 可同时调试前后端。

## 安全与监控

### 已内置能力

| 类别 | 配置/说明 |
|------|-----------|
| 全局限流 | `system.iplimit-count` / `iplimit-time`；无 Redis 时自动降级内存限流 |
| 认证限流 | `security.auth-limit-count` / `auth-limit-time`（登录、注册、错误上报） |
| SQL 注入防护 | GORM 参数化 + 排序字段白名单 |
| 上传白名单 | `security.upload-allowed-exts`，静态目录拦截脚本扩展名 |
| JWT | 后台 Audience `GVA`；会员 `GVA-MALL`；禁用账号 token 失效 |
| 会员安全 | 可选验证码、登出 token 黑名单 |
| 操作审计 | 登录/改密等请求体密码脱敏 |
| 健康检查 | `GET /health`（DB / Redis 状态） |
| 指标 | `GET /metrics`（Prometheus，可选 `metrics-auth-token`） |
| Swagger | `security.swagger-enable` 开关 |

### 生产环境

**环境变量（后端）：**

```bash
export GVA_JWT_SIGNING_KEY="$(uuidgen)"
export GVA_MYSQL_PASSWORD="your-db-password"
export GVA_REDIS_PASSWORD="your-redis-password"
export GVA_SECURITY_SWAGGER_ENABLE=false
export GVA_SECURITY_METRICS_AUTH_TOKEN="your-metrics-token"
export GVA_SECURITY_ERROR_REPORT_KEY="your-error-report-key"
```

**前端构建（`web/.env.production`）：**

```bash
VITE_BASE_API=/api
VITE_FILE_API=/api
VITE_BASE_PATH=https://your-admin-domain.com   # 改成真实域名
VITE_ERROR_REPORT_KEY=your-error-report-key     # 与后端 security.error-report-key 一致
```

**`config.yaml` 生产建议：**

- `security.swagger-enable: false`
- `security.metrics-auth-token`、`security.error-report-key` 设随机值
- `system.disable-auto-migrate: true`
- `system.use-redis: true`

**日志与告警：** 应用日志在 `server/log/`（Zap）。可对接 Loki/ELK，配合 Prometheus + Grafana 监控 `/metrics`（QPS、延迟、5xx）。

### 权限与种子数据

修改路由或 `server/source/system/` 种子后，**已有数据库不会自动更新**。请在管理后台：

**超级管理员 → API 管理 → 同步 API → 确认同步**

否则可能出现新接口无权限或残留废弃 API。

### 跨域部署

本地开发靠 Vite 代理，无需 CORS。若前后端分域名部署，需在 `server/initialize/router.go` 启用 `middleware.CorsByRules()`，并在 `config.yaml` 的 `cors` 中配置白名单（需包含 `x-member-token` 等头）。

## 说明

本项目仅供个人学习与本地测试，不包含完整生产部署方案。Docker 参考配置见 `server/config.docker.yaml` 与 `Makefile`，上线前务必替换默认密钥与地址。
