package mypc

import "fmt"

// estructura publica
type Pc struct {
	ram   uint32 // atributos privados
	disk  uint32
	brand string
}

// funcion publica
func New(rm, dsk uint32, brn string) Pc {
	customPc := Pc{ram: rm, disk: dsk, brand: brn}
	return customPc
}

func (myPC Pc) FormatPrint() {
	fmt.Printf("Esta pc marca %s cuenta con una ram de %dGB y un disco de %dGB.\n", myPC.brand, myPC.ram, myPC.disk)
}

func (myPC *Pc) DuplicateRAM() {
	myPC.ram *= 2
}

// Getter & Setter
func (myPc Pc) GetRam() uint32 {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm uint32) {
	myPc.ram = rm
}
