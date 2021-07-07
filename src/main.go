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

	//DECLARACIÓN DE VARIABLES

	//AreaCuadrao
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//SUMA
	result := x + y
	fmt.Println("Suma", result)

	//RESTA
	result1 := y - x
	fmt.Println("Resta", result1)

	//MULTIPLICACIÓN
	result2 := x * y
	fmt.Println("Multiplicación", result2)

	//DIVISIÓN
	result3 := y / x
	fmt.Println("División", result3)

	//MODULO
	result4 := y % x
	fmt.Println("Modulo", result4)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Incremental:", x)
}
