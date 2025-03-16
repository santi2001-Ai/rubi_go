package main

import (
	"fmt"
)

// Definición de la estructura Usuario
type Usuario struct {
	Nombre string
	Pin    int
	Saldo  float64
}

// Definición de la estructura Banco
type Banco struct {
	Usuarios []Usuario
}

// Método para autenticar a un usuario
func (b *Banco) Autenticar(nombre string, pin int) bool {
	for _, usuario := range b.Usuarios {
		if usuario.Nombre == nombre {
			if usuario.Pin == pin {
				fmt.Println("¡Estas Logeado!")
				return true
			} else {
				fmt.Println("Pin o nombre incorrecto")
				return false
			}
		}
	}
	fmt.Println("El usuario no existe")
	return false
}

// Método para retirar dinero
func (b *Banco) SacarDinero(nombre string, saldo float64) {
	for i := range b.Usuarios {
		if b.Usuarios[i].Nombre == nombre {
			if b.Usuarios[i].Saldo < saldo {
				fmt.Println("Saldo insuficiente")
				break
			} else {
				b.Usuarios[i].Saldo -= saldo
				fmt.Println("El saldo disponible es:", b.Usuarios[i].Saldo)
			}
		}
	}
}

func main() {
	// Creación de usuarios
	ana := Usuario{"Ana", 9872, 450}
	pablo := Usuario{"Pablo", 5555, 200}
	rodrigo := Usuario{"Rodrigo", 2222, 900}

	// Creación del banco con los usuarios
	banco := Banco{Usuarios: []Usuario{ana, pablo, rodrigo}}

	for {
		fmt.Println("Bienvenido al Banco, por favor, identifíquese.")
		fmt.Println("Introduzca el nombre:")
		var nombre string
		fmt.Scanln(&nombre)

		fmt.Println("Introduzca el PIN:")
		var pin int
		fmt.Scanln(&pin)

		if banco.Autenticar(nombre, pin) {
			for {
				fmt.Println("Por favor, elija una opción:")
				fmt.Println("1. Sacar dinero")
				fmt.Println("2. Terminar sesión")

				var opcion int
				fmt.Scanln(&opcion)

				if opcion == 1 {
					fmt.Println("Introduce cantidad a sacar:")
					var saldo float64
					fmt.Scanln(&saldo)
					banco.SacarDinero(nombre, saldo)
				} else if opcion == 2 {
					fmt.Println("Saliendo del sistema.")
					break
				} else {
					fmt.Println("Opción incorrecta. Por favor, introduzca otra opción.")
				}
			}
		} else {
			fmt.Println("Usuario no autenticado. Por favor, introduzca nombre y PIN correctos.")
		}
	}
}
