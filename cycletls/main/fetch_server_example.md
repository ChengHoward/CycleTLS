# Fetch服务使用说明

## 启动服务

```bash
cd cycletls/main
go run fetch_server.go
```

服务默认运行在 `http://localhost:8080`

## API接口

### POST /fetch

发送HTTP请求的接口。

#### 请求示例

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://tls.browserleaks.com/json",
    "method": "GET",
    "headers": {
      "Accept": "application/json"
    },
    "timeout": 30
  }'
```

#### 请求参数

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| url | string | 是 | 请求的URL地址 |
| method | string | 否 | HTTP方法，默认为GET |
| headers | object | 否 | 请求头，键值对格式 |
| body | string | 否 | 请求体（字符串格式） |
| proxy | string | 否 | 代理地址，格式：http://host:port 或 socks5://host:port |
| timeout | int | 否 | 超时时间（秒），默认30秒 |
| disableRedirect | bool | 否 | 是否禁用重定向 |
| userAgent | string | 否 | 自定义User-Agent |
| ja3 | string | 否 | JA3指纹字符串，如果不提供则使用Firefox指纹 |
| cookies | array | 否 | Cookie数组 |

#### 响应格式

```json
{
  "ok": true,
  "status": 200,
  "headers": {
    "Content-Type": "application/json",
    "Content-Length": "1234"
  },
  "body": "响应体内容（已自动解码）",
  "error": ""
}
```

#### 响应字段

| 字段 | 类型 | 说明 |
|------|------|------|
| ok | bool | 请求是否成功 |
| status | int | HTTP状态码 |
| headers | object | 响应头 |
| body | string | 响应体（已自动解码gzip、br、deflate等压缩格式） |
| error | string | 错误信息（如果ok为false） |

## 完整请求示例

### 1. 简单GET请求

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://httpbin.org/get"
  }'
```

### 2. POST请求带Body

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://httpbin.org/post",
    "method": "POST",
    "headers": {
      "Content-Type": "application/json"
    },
    "body": "{\"key\":\"value\"}"
  }'
```

### 3. 使用代理

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://httpbin.org/ip",
    "proxy": "http://127.0.0.1:1080"
  }'
```

### 4. 自定义JA3指纹

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://tls.browserleaks.com/json",
    "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513,29-23-24,0",
    "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
  }'
```

### 5. 带Cookie的请求

```bash
curl -X POST http://localhost:8080/fetch \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://httpbin.org/cookies",
    "cookies": [
      {
        "name": "session",
        "value": "abc123",
        "domain": "httpbin.org"
      }
    ]
  }'
```

## 健康检查

```bash
curl http://localhost:8080/health
```

## 性能优化

- 服务使用全局CycleTLS客户端实例，支持高并发
- 自动设置最大CPU核心数
- 响应体自动解码（gzip、br、deflate、zstd等）
- 支持连接复用

## 注意事项

1. 如果不提供`ja3`参数，服务会自动使用Firefox指纹
2. 响应体会自动根据`Content-Encoding`头进行解码
3. 如果请求失败，`ok`字段为`false`，错误信息在`error`字段中
4. 超时时间建议根据实际需求设置，避免过长导致资源占用



