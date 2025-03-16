package main

import (
	"fmt"
)

func main() {
	var figura string
	fmt.Println("Seleccione la figura para calcular su área: ")
	fmt.Println("\n1 para rombo\n2 para rectángulo\n3 para cuadrado\n4 para trapecio\n")
	fmt.Scanln(&figura)

	const Pi = 3.1416
	const Const = 2

	switch figura {
	case "1":
		var Dmayor, Dmenor float64
		fmt.Print("Ingrese el valor de la diagonal mayor: ")
		fmt.Scanln(&Dmayor)
		fmt.Print("Ingrese el valor de la diagonal menor: ")
		fmt.Scanln(&Dmenor)
		area := (Dmayor * Dmenor) / float64(Const)
		fmt.Println("El área del rombo es:", area)

	case "2":
		var Largo, Ancho float64
		fmt.Print("Ingrese el valor del largo del rectángulo: ")
		fmt.Scanln(&Largo)
		fmt.Print("Ingrese el valor del ancho del rectángulo: ")
		fmt.Scanln(&Ancho)
		area := Largo * Ancho
		fmt.Println("El área del rectángulo es:", area)

	case "3":
		var Lado float64
		fmt.Print("Ingrese el valor del lado del cuadrado: ")
		fmt.Scanln(&Lado)
		area := Lado * Lado
		fmt.Println("El área del cuadrado es:", area)

	case "4":
		var Bmayor, Bmenor, H float64
		fmt.Print("Ingrese el valor de la base mayor: ")
		fmt.Scanln(&Bmayor)
		fmt.Print("Ingrese el valor de la base menor: ")
		fmt.Scanln(&Bmenor)
		fmt.Print("Ingrese la altura del trapecio: ")
		fmt.Scanln(&H)
		area := ((Bmayor + Bmenor) * H) / 2
		fmt.Println("El área del trapecio es:", area)

	default:
		fmt.Println("Opción no válida")
	}
}
