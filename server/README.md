# mall-admin / server

Go 后端：管理后台 API（JWT + Casbin）与 C 端会员 API（`x-member-token`）。

## 项目结构

```text
server/
├── api/
│   ├── admin/          # 管理后台 Handler（system / example / mall）
│   └── siteapi/        # C 端 Handler（会员认证等）
├── config/             # config.yaml 配置结构体
├── core/               # Viper、Zap、HTTP Server 启动
├── docs/               # Swagger 生成物
├── global/             # 全局单例（DB、Redis、配置、日志）
├── initialize/         # 路由、Gorm、Redis、插件、健康检查等初始化
├── middleware/         # JWT、Casbin、限流、安全头、Prometheus、操作审计
├── model/
│   ├── mall/           # 商城领域模型
│   ├── system/         # 系统模型
│   └── example/        # 示例模块
├── plugin/             # 邮件、公告等插件
├── router/
│   ├── admin/          # 管理端路由（system / example / mall）
│   └── siteapi/        # C 端路由
├── service/            # 业务逻辑
├── source/             # 数据库种子数据（菜单、API、Casbin 等）
├── utils/              # 工具（JWT、上传、验证码、校验等）
├── config.yaml         # 本地主配置
├── config.example.yaml # 生产敏感项模板（片段）
└── config.docker.yaml  # Docker 构建参考配置
```

| 目录 | 说明 |
|------|------|
| `api/admin/` | 需 `x-token` + Casbin 权限的管理接口 |
| `api/siteapi/` | C 端公开/会员接口，不走 Casbin |
| `router/admin/` | 与 `api/admin` 对应的路由注册 |
| `source/system/` | 首次初始化写入的菜单、API、角色权限种子 |
| `middleware/` | 全局安全与观测中间件入口 |
| `initialize/router.go` | Gin 总路由、限流、Swagger、静态文件、metrics |

## 快速启动

```bash
cd server
go mod tidy
go run .
```

默认监听 **8888**。首次无库时，访问前端会引导执行 `POST /init/initdb`。

## 配置说明

| 文件 | 用途 |
|------|------|
| `config.yaml` | 本地开发主配置 |
| `config.example.yaml` | 生产敏感项与环境变量对照 |
| `config.docker.yaml` | 容器镜像参考（需按环境覆盖密钥） |

### 常用配置项

```yaml
system:
  addr: 8888
  router-prefix: ""       # 全局 API 前缀，本地一般为空
  use-redis: true         # 建议开启（验证码、多点登录、限流精度）
  iplimit-count: 15000    # 全局限流：单 IP 窗口内最大请求数
  iplimit-time: 3600      # 全局限流窗口（秒）

jwt:
  signing-key: "..."      # 生产请用环境变量 GVA_JWT_SIGNING_KEY

security:
  swagger-enable: true    # 生产设为 false
  metrics-enable: true    # Prometheus /metrics
  member-captcha-enable: true
  auth-limit-count: 30    # 登录/注册限流
  auth-limit-time: 60
```

### 环境变量覆盖（生产推荐）

```bash
export GVA_JWT_SIGNING_KEY="$(uuidgen)"
export GVA_MYSQL_PASSWORD="your-db-password"
export GVA_REDIS_PASSWORD="your-redis-password"
export GVA_SECURITY_SWAGGER_ENABLE=false
export GVA_SECURITY_METRICS_AUTH_TOKEN="your-metrics-token"
export GVA_SECURITY_ERROR_REPORT_KEY="your-error-report-key"
```

## 路由约定

| 类型 | 前缀/路径 | 鉴权 |
|------|-----------|------|
| 管理后台 | `/user/*`、`/menu/*`、`/api/*`、`/mall/member/*` 等 | `x-token` + Casbin |
| C 端会员 | `/site/auth/*` | 公开或 `x-member-token` |
| 基础 | `/base/login`、`/base/captcha` | 公开（登录有限流） |
| 初始化 | `/init/initdb`、`/init/checkdb` | 公开（仅库未初始化时） |
| 观测 | `/health`、`/metrics` | 公开（metrics 可配 token） |

前端开发时 Vite 将 `/api` 代理到 `8888` 并去掉前缀，因此后端 `router-prefix` 保持空字符串即可。

## 安全与监控

| 能力 | 实现位置 |
|------|----------|
| 全局限流 | `middleware.DefaultLimit()`，Redis 不可用时降级内存限流 |
| 认证限流 | `middleware.AuthRateLimit()`，用于登录/注册/错误上报 |
| 上传白名单 | `utils/upload/validate.go` |
| 静态文件防护 | `initialize/static_secure.go` |
| JWT Audience | 后台 `GVA`、会员 `GVA-MALL` |
| 操作日志脱敏 | `middleware/operation.go` |
| Panic 恢复 | `middleware/error.go` → `sys_errors` |
| Prometheus | `middleware/prometheus.go` → `GET /metrics` |

Swagger 由 `security.swagger-enable` 控制，默认开发开启。

## 种子数据与权限

初始化数据在 `source/system/`（`api.go`、`casbin.go`、`menu.go` 等）。

**已有数据库不会自动应用种子变更**。修改种子或路由后，请在管理后台：

**超级管理员 → API 管理 → 同步 API → 确认同步**

否则可能出现新接口无权限（如 `freshCasbin`）或残留废弃 API。

## 常用命令

```bash
# 编译
go build -o server .

# 生成 Swagger
swag init

# 运行测试
go test ./...
```

## 插件

- `plugin/email`：邮件发送
- `plugin/announcement`：公告示例（`/info/*`）

插件在 `initialize/plugin.go` 中安装，路由挂载到与管理端相同的 Gin 引擎。
