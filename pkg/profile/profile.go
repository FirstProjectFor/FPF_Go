package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const kb = 1 << 10
const mb = kb << 10

//go:generate go tool pprof -inuse_space http://127.0.0.1:8080/debug/pprof/heap
//go:generate go tool pprof -h
//go:generate go tool pprof -sample_index=inuse_space http://127.0.0.1:8080/debug/pprof/heap
//go:generate go tool pprof -sample_index=inuse_space -http=:8081 http://127.0.0.1:8080/debug/pprof/heap
//go:generate go tool pprof -sample_index=cpu http://127.0.0.1:8080/debug/pprof/profile?seconds=30
//go:generate go tool pprof -sample_index=cpu -http=:8081 http://127.0.0.1:8080/debug/pprof/profile?seconds=30
//go:generate go tool pprof -sample_index=samples http://127.0.0.1:8080/debug/pprof/profile?seconds=30
//go:generate go tool pprof -sample_index=samples -http=:8081 http://127.0.0.1:8080/debug/pprof/profile?seconds=30
//go:generate go tool pprof http://127.0.0.1:8080/debug/pprof/goroutine
//go:generate go tool pprof -http=:8081 http://127.0.0.1:8080/debug/pprof/goroutine

func main() {

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	root := &node{
		value: generateData(mb),
		next:  nil,
	}
	go makeChain(root)
	slowRoot := &node{
		value: generateData(mb),
		next:  nil,
	}
	go makeChainSlow(slowRoot)

	time.Sleep(time.Hour * 100)
	fmt.Println(root)
	fmt.Println(slowRoot)
}

func makeChain(n *node) {
	n.next = &node{
		value: generateData(mb),
		next:  nil,
	}
	time.Sleep(200 * time.Millisecond)
	makeChain(n.next)
}

func makeChainSlow(n *node) {
	n.next = &node{
		value: generateData(mb),
		next:  nil,
	}
	time.Sleep(300 * time.Millisecond)
	makeChainSlow(n.next)
}

type node struct {
	value []byte
	next  *node
}

func generateData(size int) []byte {
	result := make([]byte, 0, size)
	rand.Read(result)
	return result
}
