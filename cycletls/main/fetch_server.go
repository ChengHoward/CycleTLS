package main

import (
	"io"
	"log"
	"net/http"
	"runtime"

	"github.com/ChengHoward/CycleTLS/cycletls"
	"github.com/ChengHoward/CycleTLS/cycletls/imitate"
	"github.com/gin-gonic/gin"
)

// FetchRequest 请求结构体
type FetchRequest struct {
	URL             string            `json:"url" binding:"required"`
	Method          string            `json:"method"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	Proxy           string            `json:"proxy"`
	Timeout         int               `json:"timeout"`
	DisableRedirect bool              `json:"disableRedirect"`
	UserAgent       string            `json:"userAgent"`
	Ja3             string            `json:"ja3"`
	Cookies         []cycletls.Cookie `json:"cookies"`
}

// FetchResponse 响应结构体
type FetchResponse struct {
	OK      bool              `json:"ok"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Error   string            `json:"error,omitempty"`
}

func main() {
	// 设置最大CPU核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化CycleTLS客户端（全局单例，支持高并发）
	client := cycletls.Init()

	// 创建gin路由（生产环境建议使用gin.ReleaseMode）
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 添加CORS中间件（如果需要跨域访问）
	r.Use(corsMiddleware())

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 主要的fetch端点
	r.POST("/fetch", func(c *gin.Context) {
		handleFetch(c, client)
	})

	// 启动服务器
	port := ":8800"
	log.Printf("Fetch服务启动在端口 %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// handleFetch 处理fetch请求
func handleFetch(c *gin.Context, client cycletls.CycleTLS) {
	var req FetchRequest

	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, FetchResponse{
			OK:    false,
			Error: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 设置默认值
	if req.Method == "" {
		req.Method = "GET"
	}
	if req.Timeout == 0 {
		req.Timeout = 30 // 默认30秒超时
	}
	if req.Headers == nil {
		req.Headers = make(map[string]string)
	}

	// 构建CycleTLS选项
	options := cycletls.Options{
		Headers:         req.Headers,
		Body:            req.Body,
		Proxy:           req.Proxy,
		Timeout:         req.Timeout,
		DisableRedirect: req.DisableRedirect,
		UserAgent:       req.UserAgent,
		Ja3:             req.Ja3,
		Cookies:         req.Cookies,
	}

	// 如果没有指定Ja3，使用Firefox指纹（可以根据需要改为Chrome等）
	if options.Ja3 == "" {
		imitate.Firefox(&options)
	}
	// 执行请求
	resp, err := client.Do(req.URL, options, req.Method)

	println(req.Method, req.URL, resp.Status)

	// 处理错误（在关闭Body之前检查）
	if err != nil {
		status := resp.Status
		if resp.Body != nil {
			resp.Body.Close()
		}
		c.JSON(http.StatusOK, FetchResponse{
			OK:     false,
			Status: status,
			Error:  err.Error(),
		})
		return
	}

	// 确保响应体被关闭
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusOK, FetchResponse{
			OK:     false,
			Status: resp.Status,
			Error:  "读取响应体失败: " + err.Error(),
		})
		return
	}

	// 解码响应体（处理gzip、br等压缩）
	contentEncoding := resp.Headers["Content-Encoding"]
	var decodedBody string
	if contentEncoding != "" {
		decodedBody = cycletls.DecompressBody(bodyBytes, []string{contentEncoding}, nil)
	} else {
		// 如果没有Content-Encoding，检查Content-Type来决定是否需要base64编码
		contentType := resp.Headers["Content-Type"]
		if contentType != "" {
			decodedBody = cycletls.DecompressBody(bodyBytes, nil, []string{contentType})
		} else {
			decodedBody = string(bodyBytes)
		}
	}

	// 返回成功响应
	c.JSON(http.StatusOK, FetchResponse{
		OK:      true,
		Status:  resp.Status,
		Headers: resp.Headers,
		Body:    decodedBody,
	})
}

// corsMiddleware CORS中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
