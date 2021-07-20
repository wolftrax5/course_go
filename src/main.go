package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println(text, " Es palindromo")
	} else {
		fmt.Println(text, " No es palindromo")
	}
}

func main() {
	var slice = []string{"hola", "que", "haces"}
	// de esta forma iteramos obteniendo el indice y el valor
	// ya que range retorna estos 2 valores
	for i, varlor := range slice {
		fmt.Println("index:", i, "value:", varlor)
	}
	// con _ (piso) escapamos el uso de indice el cual no usamos
	for _, varlor := range slice {
		fmt.Println("value:", varlor)
	}
	// con _ (piso) escapamos el uso de indice el cual no usamos
	for i := range slice {
		fmt.Println("index:", i)
	}

	isPalindrome("ama")
}
