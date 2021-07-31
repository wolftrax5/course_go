package main

import "fmt"

// indicamos que este canal eso solo de entrada
// lo contrario seria c <-chan
func say(text string, c chan<- string) {
	// c <- esto indica que vamos a ingresar un dato
	// en este canal
	c <- text

}

func main() {
	c := make(chan string, 1)
	fmt.Println("HEllo")
	go say("Bye", c)
	// <-c esto indica que vamos a obtener la salida
	// un dato en este canal
	fmt.Println(<-c)
}
