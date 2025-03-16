package main

import (
	"fmt"
)

func main() {
	// Variables booleanas
	a := true
	b := false
	fmt.Println(a && b) // Equivalente a "a and b" en Python

	// Variables enteras
	a = 2
	b = 3
	c := 4
	d := 10

	// Operaciones lógicas
	fmt.Println(a == b && c < d)  // Evaluación de condiciones con AND
	fmt.Println(!(a == b) && c > d) // Uso de NOT y AND
}
