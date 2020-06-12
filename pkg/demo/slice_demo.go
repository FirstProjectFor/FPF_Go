package main

import "fmt"

func main() {

	s := make([]string, 5)
	fmt.Println("empty s : ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s : ", s)
	fmt.Println("length os s: ", len(s))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")

	fmt.Println("s : ", s)
	fmt.Println("length os s: ", len(s))

	fmt.Println("s[2,3] : ", s[2:3])

}
