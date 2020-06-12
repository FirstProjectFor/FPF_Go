package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
	"fmt"
)

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

// LookUp get value from cache
func LookUp(key string) string {
	mu.Lock()
	result := mapping[key]
	mu.Unlock()
	return result
}

// ADD add key value to cache
func ADD(key string, value string) string {
	mu.Lock()
	oldValue := mapping[key]
	mapping[key] = value
	mu.Unlock()
	return oldValue
}

func main() {

	go func(des string) {
		for {
			number := rand.Intn(1000)
			key := fmt.Sprintf("%d", number)
			value := fmt.Sprintf("%d", number)
			ADD(key, value)
			log.Printf("%s, key: %s, value:%s", des, key, value)
		}
	}("set")

	go func(des string) {
		for {
			number := rand.Intn(1000)
			key := fmt.Sprintf("%d", number)
			value := LookUp(key)
			log.Printf("%s, key: %s, value:%s", des, key, value)
		}
	}("set")

	waitTime := time.Second * 100
	time.Sleep(waitTime)
}
