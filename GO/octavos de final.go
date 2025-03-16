package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Semilla para números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Lista de partidos
	partidos := [][]string{
		{"Paises Bajos", "Corea del Sur"}, {"Senegal", "Portugal"}, {"Inglaterra", "Suiza"},
		{"USA", "Brazil"}, {"Argentina", "Croacia"}, {"Polonia", "Marruecos"},
		{"Francia", "España"}, {"Australia", "Japon"},
	}

	// Mapa para almacenar resultados
	resultados := make(map[int]int)

	// Slice para los equipos que avanzan
	var equiposPasan []string

	// Iterar sobre los partidos
	for i, partido := range partidos {
		equipo1, equipo2 := partido[0], partido[1]
		resultado := rand.Intn(3) // 0 = pierde equipo1, 1 = empate, 2 = gana equipo1

		// Guardar el resultado en el mapa
		resultados[i] = resultado

		// Determinar el equipo que avanza
		if resultado == 2 {
			fmt.Printf("%s gana contra %s\n", equipo1, equipo2)
			equiposPasan = append(equiposPasan, equipo1)
		} else if resultado == 1 {
			fmt.Printf("Empate entre %s y %s\n", equipo1, equipo2)
		} else {
			fmt.Printf("%s pierde contra %s\n", equipo1, equipo2)
			equiposPasan = append(equiposPasan, equipo2)
		}
	}

	// Mostrar los resultados finales
	fmt.Println("\nResultados finales:", resultados)

	// Mostrar los equipos que avanzan
	fmt.Println("\nEquipos que pasan a la siguiente ronda:")
	for _, equipo := range equiposPasan {
		fmt.Println(equipo)
	}
}
