package main

import (
	"fmt"
)

func main() {
	var a, b, c string
	var d, e int

	fmt.Print("Escriba sus nombres completos: ")
	fmt.Scanln(&a)

	fmt.Print("Escriba sus Apellidos completos: ")
	fmt.Scanln(&b)

	fmt.Print("Escriba su profesión: ")
	fmt.Scanln(&c)

	fmt.Print("Escriba su año de nacimiento: ")
	fmt.Scanln(&d)

	e = 2025 - d
	fmt.Printf("El (La) %s %s %s tiene %d años\n", c, a, b, e)
}
