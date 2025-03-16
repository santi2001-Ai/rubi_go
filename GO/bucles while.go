package main

import (
	"fmt"
)

func main() {
	// Bucle while en Go (simulado con for)
	contador := 0
	for contador < 30 {
		contador++
		fmt.Println("Iteración", contador)
	}

	// Bucle con if-else en Go
	for {
		fmt.Println("Introduzca un valor mayor a 10")
		var a int
		fmt.Scan(&a) // Equivalente a input() en Python

		if a > 10 {
			fmt.Println("Es correcto")
		} else {
			fmt.Println("Es incorrecto, pruebe de nuevo")
			break // Rompe el bucle como en Python
		}
	}
}
