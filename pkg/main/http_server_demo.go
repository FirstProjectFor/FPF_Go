package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello go server")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v:%v", name, h)
		}
	}
}

func waitSignal(d chan<- bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigs
	fmt.Println("accept signal: ", sig)
	d <- true
}

func main() {
	done := make(chan bool, 1)
	fmt.Println("start simple http server!")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
	<-done
	fmt.Println("start simple http server!")
}
