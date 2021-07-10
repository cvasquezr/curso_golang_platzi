package main

import "fmt"

func isPair(number int) bool {
	return number % 2 == 0
}
func validateUser(username, password string) bool {
	return username == "carlos" && password == "123456"
}

func main() {
	valor1 := 1
	valor2 := 5

	if valor1 == 1 {
		fmt.Println("El valor1 es igual a 1")
	} else {
		fmt.Println("El valor1 es diferente a 1")
	}

	if valor1 == 1 && valor2 == 5 {
		fmt.Println("Correcto")
	}

	if valor1 < valor2 {
		fmt.Println("El valor1 es menor al valor")
	}

	if isPair(6) {
		fmt.Println("Then number is pair")
	} else {
		fmt.Println("The number is ood")
	}

	if validateUser("carlos", "123456") {
		fmt.Println("The username is correct")
	} else {
		fmt.Println("Then username is incorrect")
	}
}
