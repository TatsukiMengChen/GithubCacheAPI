package github

import (
    "io"
    "net/http"
    "os"
    "strings"
    "sync/atomic"
)

var (
    tokens []string
    tokenIdx uint32
)


// 手动初始化 tokens，需在 .env 加载后调用
func InitTokens() {
    env := os.Getenv("GITHUB_TOKENS")
    tokens = nil
    for _, t := range strings.Split(env, ",") {
        t = strings.TrimSpace(t)
        if t != "" {
            tokens = append(tokens, t)
        }
    }
}

func nextToken() string {
    if len(tokens) == 0 {
        return ""
    }
    idx := atomic.AddUint32(&tokenIdx, 1)
    return tokens[int(idx)%len(tokens)]
}


func ProxyRequest(method, path, query string, body io.Reader) (*http.Response, error) {
    url := "https://api.github.com/" + path
    if query != "" {
        url += "?" + query
    }
    req, _ := http.NewRequest(method, url, body)
    req.Header.Set("Accept", "application/vnd.github+json")
    token := nextToken()
    if token != "" {
        req.Header.Set("Authorization", "Bearer "+token)
    }
    return http.DefaultClient.Do(req)
}

// 非 GET 方法，用户必须自行传入 token
func ProxyRequestWithToken(method, path, query string, body io.Reader, userToken string) (*http.Response, error) {
    url := "https://api.github.com/" + path
    if query != "" {
        url += "?" + query
    }
    req, _ := http.NewRequest(method, url, body)
    req.Header.Set("Accept", "application/vnd.github+json")
    if userToken != "" {
        // 支持 "Bearer xxx" 或直接 token
        if strings.HasPrefix(userToken, "Bearer ") {
            req.Header.Set("Authorization", userToken)
        } else {
            req.Header.Set("Authorization", "Bearer "+userToken)
        }
    }
    return http.DefaultClient.Do(req)
}
