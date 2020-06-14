package demo

import (
	"io/ioutil"
	"testing"
)

import "fmt"

import "os"

func check(e error) {
	if nil != e {
		panic(e)
	}
}

func TestWriteFile(t *testing.T) {
	d1 := []byte("hello \n go\n")
	err := ioutil.WriteFile("/tmp/data", d1, 0777)
	check(err)

	data, err := ioutil.ReadFile("/tmp/data")
	check(err)
	fmt.Println("data: ", string(data))

	f, err := os.Create("/tmp/data1")
	check(err)
	defer f.Close()

	n, err := f.WriteString("Hello Hello")
	check(err)
	fmt.Println("write numbers: ", n)
	f.Sync()

	data2, err2 := ioutil.ReadFile("/tmp/data1")
	check(err2)
	fmt.Println("data2: ", string(data2))

}
