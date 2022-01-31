![Github CI/CD](https://img.shields.io/github/workflow/status/asavt7/lrucache/build)
![Go Report](https://goreportcard.com/badge/github.com/asavt7/lrucache)
![Repository Top Language](https://img.shields.io/github/languages/top/asavt7/lrucache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/asavt7/lrucache)
![Github Repository Size](https://img.shields.io/github/repo-size/asavt7/lrucache)
![Github Open Issues](https://img.shields.io/github/issues/asavt7/lrucache)
![Lines of code](https://img.shields.io/tokei/lines/github/asavt7/lrucache)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/asavt7/lrucache)
![GitHub contributors](https://img.shields.io/github/contributors/asavt7/lrucache)

# [LRU cache](https://lk.rebrainme.com/golang-advanced/task/508)

## Getting started

```shell
go get github.com/asavt7/lrucache
```

## Examples

```go
//init cache with max keys size = 100
var cache lrucache.LRUCache = lrucache.NewLRUCache(100)

k := "0"
if v, ok := cache.Get(k); !ok {
//do dome work
v := "payload"
cache.Add(k, v)
}

// rm key in cache
cache.Remove(k)


```