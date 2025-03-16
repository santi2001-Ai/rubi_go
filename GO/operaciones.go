package main

import (
	"fmt"
)

func main() {
	var A, B int

	// Solicitar valores al usuario
	fmt.Print("Ingrese el primer valor: ")
	fmt.Scan(&A)
	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scan(&B)

	// Operaciones matemáticas
	suma := A + B
	resta := A - B
	multi := A * B
	div := float64(A) / float64(B) // Conversión a float para obtener división real

	// Imprimir resultados
	fmt.Println("La suma de los números es:", suma)
	fmt.Println("La resta de los números es:", resta)
	fmt.Println("La multiplicación de los números es:", multi)
	fmt.Println("La división de los números es:", div)

	// Comparaciones
	fmt.Println("Los números son iguales:", A == B)

	if A > B {
		fmt.Println("El número menor es:", B)
		fmt.Println("El número mayor es:", A)
	} else {
		fmt.Println("El número menor es:", A)
		fmt.Println("El número mayor es:", B)
	}
}
