package main

import "fmt"

func main() {
	x := 13 % 3;
	println(x)
	if x == 1 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	for i := 0; i < 70; i++ {
		if i % 2 == 1 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}