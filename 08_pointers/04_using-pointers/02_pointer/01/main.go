package main

import (
	"fmt"
)

func zero(z *int)  {
	*z = 0
}

func main()  {
	a := 10

	fmt.Println(a)

	zero(&a)
	fmt.Println(a)
}