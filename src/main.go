package main

import (
	pc "course_go/src/mypc"
	"fmt"
)

func main() {
	a := 50
	b := &a
	// el & es para accder a le direccion
	// de memoria
	fmt.Println("b", b)
	// y el * es para acceder al valor
	// de la direccino de memoria
	fmt.Println("*b", *b)
	*b = 100
	// cualquier modificacion a la direccion de memoria
	// se modificara en todas sus referencias
	fmt.Println("a", a)

	// ejemplo en un modulo
	myPc := pc.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()

	// usando el output Personalizado
	fmt.Println(myPc)
}
