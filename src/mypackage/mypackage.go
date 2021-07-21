package mypackage

// al crearla de forma capitalizada la primera letra
// se indica que es de accesso publico
// es buena practica dejar siempre un comentario de lo que es y hace

// CarPublic Car con acesso publico
type CarPublic struct {
	Brand string
	Year  int16
}

type carPrivate struct {
	brand string
	year  int16
}

// PrintMessage imprime un mensaje
func PrintMessage() {
	println("Message")
}

func privateMessage() {
	println("Message Private")
}

// Los modificadores de Acceso tambien pueden ser aplicados
// a funciones , structs, slices, maps,
// e incluso a los tipos de datos dentr de estas
