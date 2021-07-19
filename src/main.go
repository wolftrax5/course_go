package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

// se declara que el argumento a y b son del mismo typo de dato
// tambien puede ser a int, b int, pero no se recomienda
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// declaramos el tipo que se retornara
func returnValue(a int) int {
	return a * 2
}

// retornas mas de 1 dato
func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func dobleTipeReturn() (c string, d int) {
	return "hello", 200
}

func main() {
	normalFunction("Hello World")
	tripleArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("value", value)
	// usando el orden de retorno asingamos los resultados
	value1, value2 := dobleReturn(2)
	fmt.Println("value1:", value1, "value2", value2)
	// suponiendo que solo necesites un solo valor
	// podemos usar  _ el piso que descarta el valor
	mensage, _ := dobleTipeReturn()
	fmt.Println("Solo el mensaje:", mensage)
	_, valor := dobleTipeReturn()
	fmt.Println("Solo el valor:", valor)
}
