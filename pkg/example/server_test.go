package demo

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	err := http.ListenAndServe("127.0.0.1:8686", &helloHandle{})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour * 10000)
}

type helloHandle struct {
}

func (h *helloHandle) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	go func() {
		host := request.Host
		remoteAddr := request.RemoteAddr
		info := fmt.Sprintf("host:%s,remoteAddr:%s", host, remoteAddr)
		fmt.Println(info)
		_, err := response.Write([]byte(info))
		if err != nil {
			fmt.Println(err)
		}
	}()
}
