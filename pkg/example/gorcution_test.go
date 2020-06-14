package demo

import (
	"math/rand"
	"testing"
)

import "sync/atomic"

import "time"

import "fmt"

type readOp struct {
	key     int
	readRes chan int
}

type writeOp struct {
	key      int
	val      int
	writeRes chan bool
}

func TestGoC(t *testing.T) {

	var readOps uint64 = 0
	var writeOps uint64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.readRes <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.writeRes <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:     rand.Intn(5),
					readRes: make(chan int),
				}
				reads <- read
				<-read.readRes
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:      rand.Intn(5),
					val:      rand.Intn(100),
					writeRes: make(chan bool),
				}
				writes <- write
				<-write.writeRes
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Microsecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Println("read ops: ", readOpsFinal)
	fmt.Println("write ops: ", writeOpsFinal)
}
