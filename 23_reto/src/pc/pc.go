package pc

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(ram, disk int, brand string) Pc {
	myPc := Pc{ram: ram, disk: disk, brand: brand}
	return myPc
}

func (myPc Pc) PrintPc() {
	fmt.Printf("Esta es un computador de marca %s con una ram de %dGB y con %dGB de espacio en disco", myPc.brand, myPc.ram, myPc.disk)
	fmt.Println("")
}

func (myPc Pc) String() string {
	return fmt.Sprintf("Esta es un computador de marca %s con una ram de %dGB y con %dGB de espacio en disco", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRam() {
	myPc.ram = myPc.ram * 2
}
