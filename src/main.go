package main

import "fmt"

func main() {
	// Array INMUTABLE
	fmt.Println("ARRAY")
	var array [4]int8
	// se auto llena con los zerovalues
	fmt.Println(array)
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	fmt.Println("SLICE")
	// Slice MUTABLE
	// muy similar a Array pero
	// No se indica la cantidad de valores que tendra
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos de Slice
	fmt.Println(slice[0])  // imprime el valor del indice 0
	fmt.Println(slice[:3]) // imprime el valor hasta el indice 3
	// el numero despuesde : es exclusivo
	fmt.Println(slice[2:4]) // imprime los valores entre el indice 2 y 4
	// el primer numero es inclusivo y el segundo exclusivo
	fmt.Println(slice[4:]) // imprime del indice 4 en delante

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)
	// appen nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
