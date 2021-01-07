package demo

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpSuoIO(t *testing.T) {

	md5.Sum([]byte(""))

	var times int64 = 0

	for {
		client := &http.Client{}
		request, err := http.NewRequest("GET", "http://suo.im/6cJRIs", nil)
		if err != nil {
			panic(err)
		}
		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		responseDataByte, err := ioutil.ReadAll(response.Body)
		md5Sum := md5.Sum(responseDataByte)
		realUrl := response.Request.URL.String()
		fmt.Println(times)
		fmt.Println(md5Sum)
		fmt.Println(realUrl)
		if realUrl != "http://www.baidu.com" {
			fmt.Println(realUrl)
			fmt.Println(string(responseDataByte))
			break
		}
		times = times + 1
	}
}
