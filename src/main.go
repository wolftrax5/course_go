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

	// Aritmetica
	x := 10
	y := 50
	// Suma
	result := x + y // usamos los : por que a sido declarada anteriormente
	fmt.Println(x, "+", y, "=", result)
	//resta
	result = x - y // no usamos los : ya fue declarada
	fmt.Println(x, "-", y, "=", result)
	// multiplicacion
	result = x * y
	fmt.Println(x, "*", y, "=", result)
	// Divicion
	result = y / x
	fmt.Println(y, "/", x, "=", result)
	// Modulo (reciduo)
	result = y % x
	fmt.Println(y, "/", x, "Modulo(residuo):", result)
	// incremental
	x++
	fmt.Println("incremental", x)
	// decremental
	x--
	fmt.Println("decremental", x)

	height := 40
	var rArea = base * height
	fmt.Println("Area del rectangulo: ", rArea)

	higher := 90
	less := 45
	size := 55
	var tArea = ((higher + less) * size / 2)
	fmt.Println("Area del trapecio: ", tArea)

	radio := 45.5
	const pi = 3.14
	var cArea = (radio * radio) * pi
	fmt.Println("Area del circulo: ", cArea)

}
