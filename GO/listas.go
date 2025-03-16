package main

import (
	"fmt"
)

func main() {
	// Lista de días de la semana
	lista := []interface{}{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	
	// Mostrar el elemento en la posición 4 (índice 4)
	fmt.Println(lista[4]) // Traerá "Viernes"

	// Imprimir la lista completa
	fmt.Println(lista)

	// Imprimir los primeros tres elementos
	fmt.Println(lista[0:3])

	// Lista con diferentes tipos de datos, incluyendo una lista anidada
	listaCompleja := []interface{}{
		"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", 1, 2, 3,
		[]interface{}{"Santiago", 0.2, 2.25, true},
	}

	// Imprimir los primeros cuatro elementos y los primeros tres elementos de la lista anidada
	fmt.Println(listaCompleja[0:4], listaCompleja[9].([]interface{})[0:3])
}
