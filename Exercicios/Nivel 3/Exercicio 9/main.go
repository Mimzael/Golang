package main

import "fmt"

func main() {
	var EsporteFavorito string = "Futebol"

	switch EsporteFavorito {
	case "Tenis":
		fmt.Println("Esporte favorito é Tenis")
	case "Futebol":
		fmt.Println("Esporte favorito é Futebol")
	case "Box":
		fmt.Println("Esporte favorito é Box")
	}

}
