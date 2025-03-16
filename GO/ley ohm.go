package main

import (
	"fmt"
)

func main() {
	var voltaje, resistencia int
	var intensidad float64

	fmt.Print("Ingrese el valor del voltaje del circuito: ")
	fmt.Scanln(&voltaje)

	fmt.Print("Ingrese el valor de la resistencia del circuito: ")
	fmt.Scanln(&resistencia)

	// Cálculo de la intensidad (amperaje)
	intensidad = float64(voltaje) / float64(resistencia)

	// Imprimir el resultado
	fmt.Printf("Al conectar un resistor de R %d ohm a una fuente de V %d voltios, circulará una corriente de %.2f amperios\n", resistencia, voltaje, intensidad)
}
