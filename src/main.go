package main

import "fmt"

/*
Si los structs que tenemos en el código tienen métodos
que hacen algo en común
(Cálculos, obtener data, etc),
es posible ejecutar éstos métodos usando una interfaz,
de esta forma evitamos hacer código por cada struct.
*/
type figuras2D interface {
	area() float64 // tine que ser exactamente le nombre de la funcion
}
type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	// ser repite el nombre pero no el tipo de dato acetado
	// o la logica de ejecucion
	return c.base * c.base
}
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectanculo := rectangulo{base: 2, altura: 4}
	/*
		el mejor punto de usar una interface
		es cuando los structs comparten una misma funcion
		para su reutilizacion
	*/
	calcular(myCuadrado)
	calcular(myRectanculo)

	// listas de interfaces
	// por defecto go no hacepta slice de diferente tipos
	// pero con interfaces podemos hacerlos
	myInterface := []interface{}{"Hola", 121, 4.90}
	// los ... destricitran cada uno delos elementos
	fmt.Println(myInterface...)
}
