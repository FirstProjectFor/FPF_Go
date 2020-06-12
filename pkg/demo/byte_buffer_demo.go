package main

import "bytes"

import "fmt"

func main() {
	fmt.Println(intToString([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(comm("123"))
	fmt.Println(comm("123456789"))
}

func intToString(numbers []int) string {

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range numbers {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comm(number string) string {
	l := len(number)
	if l <= 3 {
		return number
	}

	var buf bytes.Buffer
	split := ""
	for i := 0; i < l; {
		buf.WriteString(split)
		buf.WriteString(number[i : i+3])
		split = ","
		i += 3
	}

	return buf.String()
}
