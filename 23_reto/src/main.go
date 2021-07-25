package main

import (
	"fmt"

	pc "github.com/cvasquezr/curso_golang_platzi/23_reto/src/pc"
)

func main() {
	myPc := pc.New(8, 240, "Asus")
	myPc.PrintPc()
	myPc.DuplicateRam()
	fmt.Println(myPc)
}
