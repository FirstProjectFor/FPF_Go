package link

import (
	_ "github.com/xiaotian/demo/pkg/link/inner"
	_ "unsafe"
)

//go:linkname helloName github.com/xiaotian/demo/pkg/link/inner.helloImpl
func helloName() string
