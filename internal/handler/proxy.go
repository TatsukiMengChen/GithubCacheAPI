package handler

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/cache"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/github"
)

func ProxyHandler(ca *cache.Cache) gin.HandlerFunc {
    return func(c *gin.Context) {
        path := c.Param("path")
        // 去除 path 前的 proxy 前缀（如 /proxy/repos/... => /repos/...）
        if len(path) > 0 && path[0] == '/' {
            path = path[1:]
        }
        query := c.Request.URL.RawQuery
        cacheKey := fmt.Sprintf("github:%s?%s", path, query)
        ctx := context.Background()

        var val string
        if c.Request.Method == http.MethodGet {
            v, err := ca.Get(ctx, cacheKey)
            if err == nil {
                ttl, _ := ca.TTL(ctx, cacheKey)
                if ttl > 0 && ttl <= time.Hour {
                    c.Data(http.StatusOK, "application/json", []byte(v))
                    return
                }
            }
            val = v
        }

        var resp *http.Response
        var err error
        if c.Request.Method == http.MethodGet {
            resp, err = github.ProxyRequest(c.Request.Method, path, query, c.Request.Body)
        } else {
            // 非 GET 方法必须用户自行传入 token
            userToken := c.GetHeader("Authorization")
            if userToken == "" {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Non-GET requests require Authorization header with token"})
                return
            }
            resp, err = github.ProxyRequestWithToken(c.Request.Method, path, query, c.Request.Body, userToken)
        }
        if err == nil && resp.StatusCode < 500 {
            defer resp.Body.Close()
            body, _ := io.ReadAll(resp.Body)
            if c.Request.Method == http.MethodGet {
                ca.Set(ctx, cacheKey, body, 24*time.Hour)
                c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
                return
            } else {
                // 非 GET 方法只做代理，不读写缓存
                c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
                return
            }
        }

        if err == nil {
            defer resp.Body.Close()
        }
        if c.Request.Method == http.MethodGet && val != "" {
            c.Data(http.StatusOK, "application/json", []byte(val))
            return
        }
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "github api unavailable and no cache"})
    }
}
