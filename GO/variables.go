package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := 4

	// Multiplicación
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: *")
	fmt.Println("La segunda variable es:", b)
	c := a * b
	fmt.Println("El resultado es:", c)
	fmt.Println("Tipo de dato:", reflect.TypeOf(c))

	// División
	fmt.Println("\nLa primera variable es:", a)
	fmt.Println("El signo de la operación es: /")
	fmt.Println("La segunda variable es:", b)
	cDiv := float64(a) / float64(b) // Conversión a float para obtener decimales
	fmt.Println("El resultado es:", cDiv)
	fmt.Println("Tipo de dato:", reflect.TypeOf(cDiv))

	// Suma
	fmt.Println("\nLa primera variable es:", a)
	fmt.Println("El signo de la operación es: +")
	fmt.Println("La segunda variable es:", b)
	c = a + b
	fmt.Println("El resultado es:", c)
	fmt.Println("Tipo de dato:", reflect.TypeOf(c))

	// Resta
	fmt.Println("\nLa primera variable es:", a)
	fmt.Println("El signo de la operación es: -")
	fmt.Println("La segunda variable es:", b)
	c = a - b
	fmt.Println("El resultado es:", c)
	fmt.Println("Tipo de dato:", reflect.TypeOf(c))

	// Potencia (Go no tiene operador **, usamos math.Pow)
	fmt.Println("\nLa primera variable es:", a)
	fmt.Println("El signo de la operación es: **")
	fmt.Println("La segunda variable es:", b)
	potencia := float64(a)
	for i := 1; i < b; i++ {
		potencia *= float64(a)
	}
	fmt.Println("El resultado es:", potencia)
	fmt.Println("Tipo de dato:", reflect.TypeOf(potencia))

	// Módulo
	fmt.Println("\nLa primera variable es:", a)
	fmt.Println("El signo de la operación es: %")
	fmt.Println("La segunda variable es:", b)
	c = a % b
	fmt.Println("El resultado es:", c)
	fmt.Println("Tipo de dato:", reflect.TypeOf(c))
}
