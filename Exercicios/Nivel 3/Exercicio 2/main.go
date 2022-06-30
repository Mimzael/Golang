package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for u := 1; u <= 3; u++ {
			fmt.Printf("\t%#U\n", u)
		}
	}
}
