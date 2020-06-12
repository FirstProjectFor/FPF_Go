package main

import "sync"

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
func LookUp(key string) (value string) {
	cache.Lock()
	value = cache.mapping[key]
	cache.Unlock()
	return
}

func main() {
	cache.mapping["name"] = "xiaotian"
	fmt.Println(cache.mapping["name"])
}
