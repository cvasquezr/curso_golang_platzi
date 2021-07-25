package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)
	fmt.Println(myCar.brand)
	fmt.Println(myCar.year)

	//Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
	fmt.Println(otherCar.brand)
}
