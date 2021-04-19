package demo

import (
	"os"
	"testing"
)

import "os/signal"

func TestSignal(t *testing.T) {
	signals := make(chan os.Signal, 1)

	//注册通道用于接受信号
	signal.Notify(signals)

	//s := <-signals

	//fmt.Println("signal: ", s)
}
