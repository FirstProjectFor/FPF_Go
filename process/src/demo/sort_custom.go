package main

import "fmt"

import "sort"

type ByLength []string

// Len
func (s ByLength) Len() int {
	return len(s)
}

//123.
func (s ByLength) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"banana", "apple", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
