package main

import "fmt"

func main() {
	// Declaracion de Constantes
	const PI float64 = 3.14

	fmt.Println("Valor de PI", PI)
	// Declaracion de Variables Enteras
	base := 12          // los : nos indican que no a sido declarada
	var altura int = 14 // se declara junto con su tipo y asinga
	var area int        // solo se declara

	fmt.Println("base:", base, "altura:", altura, "area:", area)

	// Zero values
	var a int     // 0
	var b float64 // 0
	var c string  // blank space
	var d bool    // false

	fmt.Println(a, b, c, d)

	// area Cuadrado
	const baseCuadrado = 10 // go asigna el valor y el tipo en vase al valor
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)
}
