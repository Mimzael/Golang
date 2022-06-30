package main

import "fmt"

func main() {
	var x int = 1
	var y int = 2

	if x > y {
		fmt.Println("x é maior que y")
	} else if y > x {
		fmt.Println("y é maior que x")
	} else {
		fmt.Println("mesmo valor")
	}
}
