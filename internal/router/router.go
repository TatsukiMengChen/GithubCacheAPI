package router

import (
    "github.com/gin-gonic/gin"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/cache"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/handler"
)

func Register(r *gin.Engine, ca *cache.Cache) {
    proxy := r.Group("/proxy")
    proxy.Any("/*path", handler.ProxyHandler(ca))
}
