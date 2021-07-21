package main

import "fmt"

// kyword para hacer estrucutras de datos
type car struct {
	brand string
	year  int16
}

func main() {
	var myCar = car{brand: "Ford", year: 2019}
	otherCar := car{brand: "Nissan", year: 2021}
	fmt.Println(myCar)
	fmt.Printf("myCar type: %T \n", myCar)
	fmt.Println(otherCar)
	fmt.Printf("otherCar type: %T \n", otherCar)
	// Tambien podemos instanciar de vorma vacia
	// tomara los zero values en donde no asignemos algo
	var nextCar car
	nextCar.brand = "Ferrari"
	fmt.Println(nextCar)
	fmt.Printf("nextCar type: %T \n", nextCar)
}
