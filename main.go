package main

import (
	ana "clase-go/ejercicios"
	"fmt"
)

var (
	a int
	b bool
	c string
	d float32
	e byte
)

func main() {
	resultado := ana.Sumar(1, 6)
	fmt.Println(resultado)
}

// modificadores de acceso
// 	cuando la incial de una funcion o variable es mayuscula, es publico
// 	cuando la incial de una funcion o variable es minuscula, es privado

// el valor por defecto de un boolean en go es false
//
