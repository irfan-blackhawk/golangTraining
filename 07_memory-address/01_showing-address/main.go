package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a- memmory addres", &a)
	fmt.Printf("%d \n", &a)
}