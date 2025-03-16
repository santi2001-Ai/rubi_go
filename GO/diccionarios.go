package main

import (
	"fmt"
)

func main() {
	// Creación de un diccionario (map en Go)
	coche := map[string]string{
		"marca":  "Chana",
		"color":  "blanco",
		"modelo": "sedan",
		"placa":  "DAD190",
	}

	// Imprimir el color del coche
	fmt.Println(coche["color"])

	// Cambiar el color
	coche["color"] = "Negro"
	fmt.Println(coche["color"])

	// Imprimir la marca del coche
	fmt.Println(coche["marca"])

	// Cambiar la marca
	coche["marca"] = "Renault"
	fmt.Println(coche["marca"])

	// Imprimir el diccionario completo
	fmt.Println(coche)
}

