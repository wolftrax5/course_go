package main

import "fmt"

func main() {
	// con make podemos crear diccionarios
	// u otros tipos de variables
	// map[typo de la llave]tipo de valor
	m := make(map[string]int)

	m["Jose"] = 13
	m["Alex"] = 20
	// en las impreciones por consola
	// no se separan por , sino por espacios
	fmt.Println(m) // map[Alex:20 Jose:13]

	// Recorrido de un map
	for i, v := range m {
		fmt.Println(i, v)
	}
	// Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)
	// a un valor no encontrado usara el Zero Value
	value = m["Josep"]
	fmt.Println(value)
	// podemos declarar otra variable
	// nos retorna un booleando
	// indicando que la llave/valor exite en el Map
	value2, ok := m["Josep"]
	fmt.Println(value2, ok)

}
