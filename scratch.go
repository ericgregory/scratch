package main

import (
	"fmt"
)

func init() {
	fmt.Println("Let's get going")
}

//go:generate wit-bindgen tiny-go wit --out-dir=gen --gofmt
func main() {}
