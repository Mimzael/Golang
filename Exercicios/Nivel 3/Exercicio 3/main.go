package main

import "fmt"

func main() {
	nascido := 2005
	data := 2022
	for nascido <= data {
		fmt.Println(nascido)
		nascido++
	}
}
