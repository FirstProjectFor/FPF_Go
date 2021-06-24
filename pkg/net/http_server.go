package main

import (
	"fmt"
	"github.com/xiaotian/demo/pkg/net/common"
	"io/ioutil"
	"net/http"
)

type HttpServer struct {
}

func (hs *HttpServer) start() {
	err := http.ListenAndServe(common.ListenAddr, hs)
	if err != nil {
		panic(err)
	}
}

func (hs *HttpServer) ServeHTTP(response http.ResponseWriter, request *http.Request) () {
	fmt.Println("accept request")
	_, err := response.Write([]byte("ping"))
	if err != nil {
		panic(err)
	}
}

func main() {
	server := HttpServer{}
	go server.start()

	response, err := http.Get(common.RequestUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bodyByte))
}
