package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)
	// b referencing a
	var b = &a;
	fmt.Println(b);
}