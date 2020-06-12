package main

import "fmt"

//ByteCounter byte count
type ByteCounter int

// WriteNoPoint Write No Point
func (c ByteCounter) WriteNoPoint(p []byte) (int, error) {
	fmt.Println(&c)
	c += ByteCounter(len(p))
	return len(p), nil
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {

	b := []byte{'1', '2', '3', '4'}
	bc := ByteCounter(0)

	fmt.Printf("%v\n", b)
	fmt.Println(bc)

	bc.Write(b)
	fmt.Println(bc)

	bc.Write(b)
	fmt.Println(bc)

	bcn := ByteCounter(0)
	bcn.WriteNoPoint(b)
	bcn.WriteNoPoint(b)
	fmt.Println(bcn)
}
