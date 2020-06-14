package demo

import (
	"testing"
	"time"
)

import "fmt"

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
