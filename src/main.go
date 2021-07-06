package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	// Declaración de constantes otra manera
	const pi2 = 3.141516
	fmt.Println(pi, pi2)
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("Hola mundo")

	//Declaración e vairables enteras
	//Incializa sin declarar variable
	base := 12
	//Declara e inicializa
	var altura int = 14
	//Declara, pero no se inicializa
	var area int
	area = base * altura
	fmt.Println(area)

	//Zero value, valor por defecto
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
