package main

import "fmt"

func main() {
	// Definir la variable
	a := 2

	// Estructura condicional simple
	if a == 2 {
		fmt.Println("a vale 2")
	} else {
		fmt.Println("a es diferente de 2")
	}

	// Definir variables
	Nombre := "Santiago"
	Edad := 23
	Pais := "Colombia"

	// Evaluar condiciones
	if Nombre == "Santiago" && Edad == 23 && Pais == "Colombia" {
		fmt.Printf("Su nombre es %s, tiene %d y es de %s\n", Nombre, Edad, Pais)
	} else if Nombre == "Santiago" && Edad != 23 {
		fmt.Println("Su nombre es Santiago y no tiene 23 años")
	} else if Nombre != "Santiago" && Edad == 23 {
		fmt.Println("Su nombre no es Santiago y tiene 23 años")
	} else {
		fmt.Println("No se llama Santiago ni tiene 23 años")
	}
}
