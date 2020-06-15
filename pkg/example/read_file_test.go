package demo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func main(t *testing.T) {

	data, err := ioutil.ReadFile("d://test.txt")
	check(err)
	fmt.Println(string(data))

	f, err := os.Open("d://test.txt")
	check(err)
	defer f.Close()

	buf1 := make([]byte, 15)
	f.Read(buf1)
	fmt.Println("first line:", string(buf1))
	o2, err := f.Seek(2, 1)
	fmt.Println("seek to: ", o2)
	buf2 := make([]byte, 3)
	n2, err := f.Read(buf2)

	check(err)
	fmt.Println("read bytes: ", n2)
	fmt.Println("second line: ", string(buf2))
}
