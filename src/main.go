package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

// la funcion main esta corriendo dentro de
// una go rutine, una ves ejecutada muere
func main() {
	/*
		Un WaitGroup espera a que una colección de goroutines
		termine su ejecución.
		Para esto se una la WaitGroup.Add()
		( wg.add(1) en el ejemplo de la clase).
		El número entero indica el número de goroutines
		que debe esperar para finalizar
		la ejecución de la goroutine principal.

		Cada vez que una goroutine termina
		su ejecución, llama el método Done().
		Esto hace que el contador del WaitGroup se
		reduzca.
		Cuando el contador llegue a zero la rutina principal
		continuará su ejecución.

		La función wait() bloquea la rutina principal
		hasta que todas las demás rutinas del grupo hayan terminado.
	*/
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)
	// para manejar las go runites es necesario
	// usar la keyword go antes de la funcion
	go say("World", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
