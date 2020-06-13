package main

import "fmt"

import "strings"

func main() {
	var strArr = []string{"peach", "apple", "pear", "plum"}
	fmt.Println("Index of apple: ", Index(strArr, "apple"))
	fmt.Println("Include pear ? ", Include(strArr, "pear"))
	fmt.Println("any start with 'p'", Any(strArr, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println("any start with 'm'", Any(strArr, func(s string) bool {
		return strings.HasPrefix(s, "m")
	}))
	fmt.Println("all start with 'p'", All(strArr, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println("map to Upper'", Map(strArr, func(s string) string {
		return strings.ToUpper(s)
	}))

	fmt.Println("map to Upper'", Map(strArr, strings.ToUpper))
}

// Index index
func Index(vs []string, t string) int {
	for i, v := range vs {
		if t == v {
			return i
		}
	}
	return -1
}

// Include include
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any any
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// ALl all
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter filter
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// MAP map
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
