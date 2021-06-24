package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
}

func (hr *HttpRequest) get() {
	response, err := http.Get("https://note.sunfeilong.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)
	bodyByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bodyByte))
}

func main() {
	request := HttpRequest{}
	request.get()
}
