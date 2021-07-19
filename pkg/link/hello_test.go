package link

import (
	"fmt"
	"math"
	"testing"
)

func TestHello(t *testing.T) {
	var maxUint64 int32 = math.MaxInt32
	fmt.Println(maxUint64)

	maxUint64 = maxUint64 + 1
	fmt.Println(maxUint64)
}
