package main

import (
	"fmt"
	"github.com/xiaotian/demo/pkg/net/common"
	"io/ioutil"
	"net/http"
)



type HttpsServer struct {
}

func (hs *HttpsServer) start() {
	err := http.ListenAndServeTLS(common.HttpsListenAddr, common.CertFile, common.KeyFile, hs)
	if err != nil {
		panic(err)
	}
}

func (hs *HttpsServer) ServeHTTP(response http.ResponseWriter, request *http.Request) () {
	response.Header().Set("Content-Type", "text/plain")
	_, err := response.Write([]byte("ping\n"))
	if err != nil {
		panic(err)
	}
}

func main() {
	server := HttpsServer{}
	go server.start()

	response, err := http.Get(common.HttpsRequestUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bodyByte))
}
