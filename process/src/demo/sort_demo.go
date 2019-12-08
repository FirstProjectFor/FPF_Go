package main

import "sort"

import "fmt"

func main() {

	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println("strings: ", strs)

	ints := []int{3, 2, 1, 5, 6, 7, 3, 2, 6}
	sort.Ints(ints)
	fmt.Println("ints:, ", ints)

	hasSorted := sort.IntsAreSorted(ints)

	fmt.Println("has sorted: ", hasSorted)
}
