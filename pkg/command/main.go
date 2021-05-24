/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/xiaotian/demo/pkg/command/cmd"

//go:generate go run main.go help

//go:generate go run main.go help echo

//go:generate go run main.go help version

//go:generate go run main.go help version version1

//go:generate go run main.go help version version2

//go:generate go run main.go

//go:generate go run main.go name age

//go:generate go run main.go echo

//go:generate go run main.go version

//go:generate go run main.go version version1

//go:generate go run main.go version version2

func main() {
	cmd.Execute()
}
