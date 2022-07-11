package main

import (
	rt "curso_golang_platzi/retos/strucs_y_punteros/src/package"
	"fmt"
)

func main() {
	myPc := rt.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}
