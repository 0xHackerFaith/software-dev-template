# Go Gin Minimal

Go + Gin 最简 HTTP API 脚手架。

## 包含内容

- `cmd/server`: HTTP 服务入口
- `internal/http`: Gin 路由
- `GET /`: 服务基本信息
- `GET /healthz`: 健康检查
- `router_test.go`: 路由测试

## 本地启动

```bash
go mod tidy
go run ./cmd/server
```

默认监听 `:8080`。需要换端口时：

```bash
PORT=3000 go run ./cmd/server
```

## 测试

```bash
go test ./...
```

