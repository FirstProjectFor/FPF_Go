package demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	text := " Hello GO! "
	fmt.Println("Count o", strings.Count(text, "o"))
	fmt.Println("Upper", strings.ToUpper(text))
	fmt.Println("TrimSpace", strings.TrimSpace(text))
	fmt.Println("Has Prefixx", strings.HasPrefix(text, " "))
	fmt.Println("Join", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("repeat", strings.Repeat("as", 5))
	fmt.Println("replace", strings.Replace("foo", "o", "O", 1))
	fmt.Println("replace", strings.Replace("foo", "o", "O", 2))
	fmt.Println("replace", strings.Replace("foo", "o", "O", 3))
	fmt.Println("Index", strings.Index("foo", "o"))
	fmt.Println("len", len(text))

	fmt.Println("replace", strings.ReplaceAll("name\rname", "\r", ""))
	fmt.Println("replace", strings.ReplaceAll("name\n\n\nname", "\n", ""))
	fmt.Println("replace", strings.ReplaceAll("name\tname", "\t", ""))
	fmt.Println("replace", strings.ReplaceAll("name\\name", "\\", ""))
	fmt.Println("replace", strings.ReplaceAll("name\"name", "\"", ""))
}

