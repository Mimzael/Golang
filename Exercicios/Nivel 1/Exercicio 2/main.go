package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {

	s := fmt.Sprintln(x, y, z)
	fmt.Println(s)
	
}
