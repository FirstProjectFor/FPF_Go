package demo

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)

	fmt.Println("ops: ", opsFinal)
}
