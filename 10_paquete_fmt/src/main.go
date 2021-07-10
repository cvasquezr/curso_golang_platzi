package main

import "fmt"

func main() {
	//Declaración sencilla
	helloMessage := "Hello"
	worldMessage := "World"

	//Println, agrega un salto de linea
	fmt.Println(helloMessage, worldMessage)

	//Printf, Es un rpint al cual le puedes especificar el tipo de objeto que le vas a dar
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// v, cuando no se sabe el tipo de dato
	fmt.Printf("%v tiene más de %d cursos\n", nombre, cursos)

	//Sprintf, No imprime nada en consola, simplemente lo guuarda como un string
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Saber tipo de dato
	fmt.Printf("helloMessage: %T", helloMessage)
}
