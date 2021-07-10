package main

import (
	"fmt"
	"math"
)

const pi = 3.141516

func normalFunction(message string) {
	fmt.Println(message)
}

func sum(a int, b int) {
	result := a + b
	fmt.Println("The result of the sum is:", result)
}

func subtract(a, b int) {
	result := a - b
	fmt.Println("The result of the subtraction is:", result)
}

func multiplication(a, b int) int {
	return a * b
}

//Retornar mas de un valor
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaCircle(radio float64) float64 {
	return pi * (radio * radio)
}

func areaSquare(side float64) float64 {
	return math.Pow(side, 2)
}

func main() {
	normalFunction("Hello world")
	sum(5, 6)
	value := multiplication(5, 8)
	fmt.Printf("The result of the operation is: %d\n", value)
	fmt.Println("The result of the operation is:", multiplication(5, 5))

	//Retornar mas de un valor
	value1, value2 := doubleReturn(2)
	fmt.Println("The value 1 and 2 are:", value1, value2)

	//Si solo necesito un elemento de los dos
	value3, _ := doubleReturn(2)
	fmt.Println("The value 3 is:", value3)

	//Area of a circle
	fmt.Println("The area of the circle is:", areaCircle((3)))

	//Area of a square or rectangule
	fmt.Println("The area of the square is:", areaSquare((2)))

}
