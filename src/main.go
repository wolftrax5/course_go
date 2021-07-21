package main

// nombre a un lado de una importacion funciona
// como un alias
import (
	"fmt"

	pk "course_go/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrai"
	myCar.Year = 1999
	fmt.Println(myCar)

	pk.PrintMessage()
}
