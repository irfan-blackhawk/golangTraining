package main

import "fmt"

const (
	_ = iota // 0
	b = iota * 10 // 1 x 10
	c = iota * 20 // 2 x 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
}