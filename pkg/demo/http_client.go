package main

import "net/http"

import "fmt"

import "bufio"

func main() {

	url := "http://note.sunfeilong.com"
	res, err := http.Get(url)
	if nil != err {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("response state: ", res.Status)

	scanner := bufio.NewScanner(res.Body)

	for i := 0; i < 100 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if error := scanner.Err(); nil != error {
		panic(error)
	}
}
