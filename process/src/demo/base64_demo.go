package main

import "encoding/base64"

import "fmt"

func main() {

	data := "abc123!?$*&()'-=@~"
	se := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(se)

	sed, _ := base64.StdEncoding.DecodeString(se)
	fmt.Println(string(sed))

	seu := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(seu)

	sedu, _ := base64.URLEncoding.DecodeString(se)
	fmt.Println(string(sedu))
}
