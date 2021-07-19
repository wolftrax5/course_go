package main

import "fmt"

func main() {
	// Defer
	/*
		** defer: ** ejecuta la parte de código especificada al final
		de la ejecución de la función o método.
		Como buena práctica usar solo una sentencia “defer”
		dentro de un bloque.
		Ejemplo: abres conexion a una BD un defear que cierre la conexion
		Cuando abres un archivo como defer para cerrarlo.
	*/
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	// Continue y Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
			// Se utiliza segun cuando una condicion dada
			// se interesa que se continua como un error controlado
		}
		// Break
		if i == 8 {
			fmt.Println("Break")
			// rompera el ciclo
			break
		}
	}
}
