package main

import "os"

import "os/signal"

import "syscall"

import "fmt"

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//注册通道用于接受信号
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	func() {
		s := <-signals
		fmt.Println("signal: ", s)
		done <- true
	}()

	<-done
	fmt.Println("system exit!")
}
