package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	fmt.Println("****************************")

	myPc := pc{ram: 16, disk: 200, brand: "asus"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRam()

	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
