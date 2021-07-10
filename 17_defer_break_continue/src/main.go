package main

import "fmt"

func main() {
	//Defer, defer deja que se ejecute antes de terminar la función
	//Por ejemplo una funcion que se conecte a la base de datos, con defer
	//esta se cerraria, se puede usar mas de un defer por funcion, pero se recomienda uno por funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {

		//continue
		if i == 2 {
			fmt.Println("Es dos")
			continue
		}

		if i == 7 {
			fmt.Println("Detenido")
			break
		}

		fmt.Println(i)
	}
}
