package demo

import (
	"fmt"
	"testing"
)

import "sort"

// ByLength string array
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

func TestSortCustom(t *testing.T) {
	fruits := []string{"banana", "apple", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

func TestSort(t *testing.T) {

	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println("strings: ", strs)

	ints := []int{3, 2, 1, 5, 6, 7, 3, 2, 6}
	sort.Ints(ints)
	fmt.Println("ints:, ", ints)

	hasSorted := sort.IntsAreSorted(ints)

	fmt.Println("has sorted: ", hasSorted)
}
