package proxy

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/cache"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/router"
)

func getEnv(key, def string) string {
    v := os.Getenv(key)
    if v == "" {
        return def
    }
    return v
}

func Run() {
    redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
    ca := cache.New(redisAddr)
    r := gin.Default()
    router.Register(r, ca)
    r.Run(":8080")
}
