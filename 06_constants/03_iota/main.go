package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota //3
)

/**
* https://golang.org/ref/spec#Iota
* Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.
* It is reset to 0 whenever the reserved word const appears in the source and increments after each ConstSpec.
* It can be used to construct a set of related constants:
*/

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}