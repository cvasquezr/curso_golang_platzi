package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	//FOR while
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	//For forever
	//counterForEver := 0
	/*for {
		counterForEver++
		fmt.Println(counterForEver)
	}*/

	//For range, parecido al foreach
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("Posicion %d numero par: %d \n", i, par)
	}

	//For inverter
	maxNumber := 10
	for maxNumber > 0 {
		fmt.Println("El numero es:", maxNumber)
		maxNumber--
	}
}
