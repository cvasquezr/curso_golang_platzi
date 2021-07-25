package main

import (
	"fmt"

	pk "github.com/cvasquezr/curso_golang_platzi/22_modificadores/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Fiat"
	myCar.Year = 2021
	fmt.Println(myCar)

	pk.PrintMessage("Hello")
}
