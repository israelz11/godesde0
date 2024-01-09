package main

import (
	"fmt"

	"github.com/israelz11/godesde0/variables"
)

func main() {

	//variables.MostrarEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConvierteATexto(1599)

	fmt.Println(estado)
	fmt.Println(texto)
}
