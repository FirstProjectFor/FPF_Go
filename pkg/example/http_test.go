package demo

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {

	md5.Sum([]byte(""))

	sum := [md5.Size]byte{95, 196, 20, 126, 7, 180, 155, 110, 125, 97, 120, 77, 48, 253, 209, 113}
	var times int64 = 0

	for {
		client := &http.Client{}
		request, err := http.NewRequest("GET", "http://www.suo.im/5W8Cqo", nil)
		if err != nil {
			panic(err)
		}
		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		responseDataByte, err := ioutil.ReadAll(response.Body)
		md5Sum := md5.Sum(responseDataByte)
		fmt.Println()
		fmt.Println(times)
		fmt.Println(md5Sum)
		if sum != md5Sum {
			fmt.Println(response.Request.URL.String())
			break
		}
		times = times + 1
		time.Sleep(time.Second * 1)
	}
}
