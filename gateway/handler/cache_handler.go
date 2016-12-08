package handler

import (
    "fmt"
    "net/http"
)

type CacheHandler struct {}

// Retrieve an instance of HTTP Cache handler
func NewCacheHandler() *CacheHandler {
    return &CacheHandler{}
}

func (cache *CacheHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    fmt.Println("Cached HTTP request: ")
    fmt.Println(r.URL.String())
    next(rw, r)
}
