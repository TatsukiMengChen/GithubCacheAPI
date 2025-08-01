# GithubCacheAPI

GithubCacheAPI 是一个用于代理 Github RESTful API 并进行缓存的服务，解决单个 token 每小时请求次数限制问题。

## 功能简介
- 代理 Github API 请求，支持所有 HTTP 方法
- 对 GET 请求结果进行 Redis 缓存，缓存有效期 1~24 小时
- 1 小时内重复 GET 请求直接返回缓存，24 小时内优先代理，失败则返回缓存
- 非 GET 请求仅代理，不缓存，且必须由用户通过 Authorization header 传入 token，不使用环境变量 token
- 支持 Docker 和 docker-compose 部署

## 目录结构
```
internal/
  cache/    # Redis 缓存逻辑
  github/   # Github API 代理逻辑
  handler/  # Gin 路由处理
  router/   # 路由分组注册
  proxy/    # 启动入口
api/        # 可扩展 API 层
config/     # 配置相关
pkg/        # 可复用包
scripts/    # 脚本
```

## 快速开始
1. 配置环境变量
   复制 `.env.example` 为 `.env`，填写 Github Token 和 Redis 地址。

2. 启动 Redis 服务（可用 docker-compose）

3. 构建并运行
```bash
docker-compose up --build
```

4. 访问接口
- 代理 Github API: `GET /proxy/repos/{owner}/{repo}`
- 其他 Github API 路径均支持
- 非 GET 请求需在请求头中传入 `Authorization: Bearer <your_token>`，否则无法访问


## 环境变量
- `GITHUB_TOKENS`：多个 Github 个人访问令牌，逗号分隔（如 `token1,token2,token3`），系统自动负载均衡使用。
  > **安全提醒：** 生成 Github Token 时请根据实际需求勾选所需 API 权限（如 repo、read:user、public_repo 等），避免授予过高权限，以防止接口滥用和安全风险（如被恶意使用或泄露）。否则部分接口可能无法访问或存在安全隐患。
- `REDIS_ADDR`：Redis 服务地址，默认 `localhost:6379`

## 贡献与扩展
- 路由分组可在 `internal/router/router.go` 扩展
- 新功能建议按 Go 项目规范分层

## License
MIT
