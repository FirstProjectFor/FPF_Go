package demo

import (
	"sync"
	"testing"
)

import "fmt"

var cache = Cache{
	mapping: make(map[string]string),
}

// Cache struct define
type Cache struct {
	sync.Mutex
	mapping map[string]string
}

// LookUp get cache value
func lookUp(key string) (value string) {
	cache.Lock()
	value = cache.mapping[key]
	cache.Unlock()
	return
}

func TestCacheStruct(t *testing.T) {
	cache.mapping["name"] = "xiaotian"
	fmt.Println(cache.mapping["name"])
}
