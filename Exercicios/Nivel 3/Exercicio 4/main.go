package main

import "fmt"

func main() {
	nascido := 2005
	data := 3000
	for {
		if nascido > data {
			break
		} else {
			fmt.Println(nascido)
			nascido++
		}
	}
}
