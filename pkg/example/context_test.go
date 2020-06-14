package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithDeadLine(t *testing.T) {
	fmt.Println("testWithDeadLine start ...")
	afterFiveSeconds := time.Now().Add(time.Second + 5)
	ctx, _ := context.WithDeadline(context.Background(), afterFiveSeconds)
	times := 1000
	for i := 0; i < times; i++ {
		go printInfo(ctx, i)
	}
	fmt.Println("testWithDeadLine wait end ...")
	time.Sleep(time.Second * 6)
	fmt.Println("testWithDeadLine end ...")
}

func TestWithCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	times := 1000
	for i := 0; i <= times; i++ {
		go printInfo(ctx, i)
	}
	time.Sleep(time.Second * 10)
	cancelFunc()
	time.Sleep(time.Second * 1)
}

func printInfo(ctx context.Context, num int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine stop, num: ", num)
			return
		default:
			i++
			fmt.Println("goroutine num: ", num, "run times:", i)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
