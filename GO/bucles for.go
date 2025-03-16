package main

import "fmt"

func main() {
    // Lista de nombres
    nombres := []string{"Santiago", "Hans", "Jhon", "Juan Pablo"}

    for _, nombre := range nombres {
        fmt.Println(nombre)
    }

    // Lista de diccionarios en Go (usamos slices de maps)
    personas := []map[string]interface{}{
        {"nombre": "Santiago", "edad": 23},
        {"nombre": "Hans", "edad": 27},
        {"nombre": "Jhon", "edad": 41},
        {"nombre": "Juan Pablo", "edad": 23},
    }

    for _, persona := range personas {
        fmt.Println(persona["nombre"], persona["edad"])
    }
}
