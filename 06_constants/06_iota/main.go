package main

import "fmt"

// x << 1 is the same as x*2 and x >> 1 is the same as x/2 but truncated towards negative infinity.
const (
	_ = iota // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary \t \t decimal")

	fmt.Println("%b\t", KB)
	fmt.Println("%d\t", KB)

	fmt.Println("%b\t", MB)
	fmt.Println("%d\t", MB)

	fmt.Println("%b\t", GB)
	fmt.Println("%d\t", GB)

	fmt.Println("%b\t", TB)
	fmt.Println("%d\t", TB)
}