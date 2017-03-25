package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter the the length in meters")
	fmt.Scan(&meters)
	yards := meters * metersToYards;
	// %v	the value in a default format
	fmt.Printf("meters %v to yards %v ", meters, yards)
}