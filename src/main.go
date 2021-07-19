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

func esPar(num int) bool {
	result := num%2 == 0
	if result {
		fmt.Println(num, "es par")
	} else {
		fmt.Println(num, "NO es par")
	}
	return result
}

func validUser(username, pass string) bool {
	return username == "kuro" && pass == "kuro"
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

	// FOR
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// For while
	var counter int8 = 0
	// El for continuara hasta cumplir la condicion
	for counter <= 5 {
		fmt.Println("counter:", counter)
		counter++ // actualizamos la condicion de paro
	}
	// For Forever
	/*
		Si no se agregan la condicion correra por siempre
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever ++
		}
	*/
	// Fore range
	// cuando se tiene una colección de un objeto
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	// IF
	const condition1 int8 = 1
	const condition2 int8 = 2

	if condition1 == 1 {
		fmt.Println("ES 1")
	} else {
		fmt.Println("No es 1")
	}

	// with and
	if condition1 == 1 && condition2 == 2 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}
	// with or
	if condition1 == 0 || condition2 == 2 {
		fmt.Println("Verdad")
	}

	// SWITCH
	// Estructura condicional - Switch
	// Si tienes una variable y conoses los valores
	// de esta variable este es tu mejor uso
	const dia int8 = 4

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}
	// otra forma
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es impar")
	}
	// swich sin condicion
	// Si vas a revisar multiples aspectos o codiciones detenrminadas
	option := 200
	switch {
	case option > 100:
		fmt.Println("Es mayor que 100")
	case option < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No Condition")
	}
}
