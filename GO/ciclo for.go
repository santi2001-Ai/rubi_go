package main

import (
	"fmt"
	"math"
)

func main() {
	// Capturar el primer número
	fmt.Println("Ingrese el primer valor:")
	var A int
	fmt.Scanln(&A)

	// Capturar el segundo número (exponente)
	fmt.Println("Ingrese el segundo valor:")
	var C int
	fmt.Scanln(&C)

	// Calcular la potencia
	valor := math.Pow(float64(A), float64(C))

	// Imprimir el resultado
	fmt.Printf("La potencia de %d sobre %d es: %.0f\n", A, C, valor)
}
