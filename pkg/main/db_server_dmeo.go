package main

import (
	"fmt"
	"net/http"
)

type database map[string]string

func (db database) list(w http.ResponseWriter, request *http.Request) {
	if nil == db {
		fmt.Fprintln(w, "no data in db")
	}

	for key, value := range db {
		fmt.Fprintf(w, "key:%s, value:%s", key, value)
	}

}

func main() {
	db := database{"apple": "100", "banana": "200"}
	http.HandleFunc("/list", db.list)
	http.ListenAndServe("localhost:8000", nil)
}
