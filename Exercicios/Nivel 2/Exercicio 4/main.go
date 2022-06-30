package main

import "fmt"

func main() {
	x := 200
	y := 0

	fmt.Printf(" %b\t %#x\t %d\t\n", x, x, x)
	y = x << 1
	fmt.Printf(" %b\t %#x\t %d", y, y, y)
}
