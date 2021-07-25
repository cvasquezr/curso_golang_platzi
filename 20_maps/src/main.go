package main

import "fmt"

func main() {
	//Los maps son llave valor en python se llaman diccionarios
	m := make(map[string]int)

	m["Carlos"] = 31
	m["Luciana"] = 5

	fmt.Println(m)

	//Recorre map
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	//Encontrar un valor
	value := m["Carlos"]
	fmt.Println(value)

	//Llave que no existe
	valor, ok := m["Charlie"]
	fmt.Println(valor, ok)
}
