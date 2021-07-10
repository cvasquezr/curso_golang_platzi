package main

import "fmt"

func main() {
	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("It's pair")
	default:
		fmt.Println("It's odd")
	}

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("It's pair")
	default:
		fmt.Println("It's odd")
	}

	//Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Value is greater than 100")
	case value < 0:
		fmt.Println("Value is less than 0")
	default:
		fmt.Println("No condition")
	}

}
