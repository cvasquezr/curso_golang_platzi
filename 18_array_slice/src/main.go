package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 3
	fmt.Println(array)

	//Tamaño del array len
	fmt.Println(len(array))

	//Capacidad maxima del array cap
	fmt.Println(cap(array))

	//Slice, similar a los array, pero no se indica el tamaño
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//Metodos slice
	fmt.Println((slice[0]))
	fmt.Println((slice[:3]))
	fmt.Println((slice[2:4]))
	fmt.Println((slice[4:]))

	//Apend
	slice = append(slice, 8)
	fmt.Println(slice)

	//Apend para agregar una nueva lista
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
