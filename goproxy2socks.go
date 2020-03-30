package main

import (
    "github.com/elazarl/goproxy"
    "golang.org/x/net/proxy"
    "log"
    "net/http"
    "os"
)

func main() {
    server := goproxy.NewProxyHttpServer()
    server.Verbose = true
    const proxyUrl = "socks5://127.0.0.1:1080"
    os.Setenv("ALL_PROXY", proxyUrl)
    server.ConnectDial = proxy.FromEnvironment().Dial
    log.Println("Starting listening as http proxy on port 8080 forwarding to socks port 1080")
    os.Setenv("HTTP_PROXY", proxyUrl)
    log.Fatal(http.ListenAndServe("[::1]:8080", server))
}
