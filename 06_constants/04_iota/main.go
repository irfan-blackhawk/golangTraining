package main

import "fmt"

/**
Auto increment with single assignment
 */

const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}