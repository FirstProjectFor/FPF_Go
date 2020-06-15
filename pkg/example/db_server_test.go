package demo

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type database map[string]string

func (db database) list(w http.ResponseWriter, request *http.Request) {
	if nil == db {
		fprintf, _ := fmt.Fprintf(w, "no data in db")
		fmt.Println(fprintf)
	}

	for key, value := range db {
		fprintf, _ := fmt.Fprintf(w, "key:%s, value:%s", key, value)
		fmt.Println(fprintf)
	}
}

func TestServer(t *testing.T) {
	db := database{"apple": "100", "banana": "200"}
	http.HandleFunc("/list", db.list)
	http.ListenAndServe("localhost:8000", nil)
	time.Sleep(time.Second * 10)
}
