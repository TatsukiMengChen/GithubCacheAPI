package main

import (
    "github.com/TatsukiMengChen/GithubCacheAPI/config"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/proxy"
    "github.com/TatsukiMengChen/GithubCacheAPI/internal/github"
)

func main() {
    config.LoadEnv()
    github.InitTokens()
    proxy.Run()
}
