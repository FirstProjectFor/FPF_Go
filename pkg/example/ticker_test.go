package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("ticket: ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("ticker stop ", time.Now().Format("2006-01-02 15:04:05"))
}
