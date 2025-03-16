package main

import (
	"fmt"
)

func main() {
	// Definir los arreglos
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	c := []int{}

	// Multiplicación de los elementos de las listas
	for i := 0; i < len(a); i++ {
		c = append(c, a[i]*b[i])
	}
	fmt.Println(c)

	// Llamada a la función mostrarTexto
	mostrarTexto()

	// Llamada a la función multiplicar
	multiplicar()

	// Variables globales
	aGlobal := 5
	bGlobal := 8
	multiplicarConVariables(aGlobal, bGlobal)

	// Llamada a la función con return
	resultado := multiplicarConReturn()
	fmt.Println(resultado + 5)

	// Validar dato
	aGlobal = 5
	dato := validarDato(aGlobal)
	fmt.Println(dato)
}

// Definir función sin parámetros
func mostrarTexto() {
	fmt.Println("hola")
}

// Función que multiplica valores fijos
func multiplicar() {
	a := 5
	b := 8
	fmt.Println(a * b)
}

// Función que usa variables externas (pasadas como argumentos)
func multiplicarConVariables(a, b int) {
	fmt.Println(a * b)
}

// Función que devuelve un valor
func multiplicarConReturn() int {
	a := 5
	b := 8
	return a * b
}

// Función que devuelve un booleano
func validarDato(a int) bool {
	if a == 5 {
		return true
	}
	return false
}
