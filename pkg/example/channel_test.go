package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestForChannel(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "e1"
	queue <- "e2"

	close(queue)

	for element := range queue {
		fmt.Println("element: ", element)
	}
}

func TestChannel(t *testing.T) {

	naturals := make(chan int64)
	squares := make(chan int64)

	var x int64 = 0

	go func() {
		for ; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	go func() {
		for s := range squares {
			fmt.Println(s)
		}
	}()

	time.Sleep(time.Second * 5)
	close(naturals)
	close(squares)
}

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1: " + msg1)
		case msg2 := <-c2:
			fmt.Println("msg2: " + msg2)
		}
	}
}

func TestCloseChannel(t *testing.T) {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("start job: ", j)
			} else {
				fmt.Println("all job has finished!")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job i: ", i)
		time.Sleep(time.Second * 1)
	}

	close(jobs)

	fmt.Println("send all jobs")
	<-done
}

func TestUnBlock(t *testing.T) {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("msg: " + msg)
	default:
		fmt.Println("default")
	}

	select {
	case messages <- "msg2":
		fmt.Println("send message to messages")
	default:
		fmt.Println("un send message to messages")
	}

	select {
	case msg := <-messages:
		fmt.Println("accept message from messages: ", msg)
	case signal := <-signals:
		fmt.Println("accept signal from signales: ", signal)
	default:
		fmt.Println("not accept any message")
	}

	time.Sleep(time.Second * 2)
}

func TestTimeOut(t *testing.T) {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "msg1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("msg1: " + msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("get msg1 time out!")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "msg2"
	}()

	select {
	case msg2 := <-c1:
		fmt.Println("msg2: " + msg2)
	case <-time.After(time.Second * 2):
		fmt.Println("get msg1 time out!")
	}
}

func TestPlus(t *testing.T) {
	const times = 100000

	firstChannel := make(chan int)

	var before, next chan int = nil, firstChannel

	for i := 0; i < times; i++ {
		before, next = next, make(chan int)
		go Translate(before, next)
	}

	next <- 0
	result := <-firstChannel

	fmt.Println(result)
}

func Translate(left, right chan int) {
	left <- 1 + <-right
}
