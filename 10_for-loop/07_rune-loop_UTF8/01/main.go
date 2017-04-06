package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Print(i, " - ", string(i), " - ", []byte(string(i)))
		fmt.Println()
	}

	foo := "a"

	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
