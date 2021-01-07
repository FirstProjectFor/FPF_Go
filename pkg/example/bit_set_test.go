package demo

import (
	"fmt"
	"github.com/willf/bitset"
	"testing"
)

func TestBitSet(t *testing.T) {

	flagMap := bitset.New(100)

	flagMap.Set(1)
	flagMap.Set(1024*1024)
	fmt.Println(flagMap.DumpAsBits())


}
