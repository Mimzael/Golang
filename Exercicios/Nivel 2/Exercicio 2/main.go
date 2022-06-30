package main

import "fmt"
 
func main(){
	a := (7 == 7)
	b := (7 != 7)
	c := (7 <= 7)
	d := (7 < 2)
	e := (7 >= 7)
	f := (7 > 8)
	fmt.Printf(" %v\n %v\n %v\n %v\n %v\n %v\n", a, b, c, d, e, f)
}