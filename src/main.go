package main

import (
	"fmt"
	"math"
)

func main() {
	// Variables, constantes y zero values
	var_const_zval()

	// Operadores aritmeticos
	arithmeticOperator()

	// Package fmt
	packageFmt()

	// call functions
	normalFunction("Hola mundo")
	multipleArguments(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, _ := doubleReturn(2)
	fmt.Println("Value1: ", value1)
	// fmt.Println("Value2: ", value2)

	// ciclos - bucles
	// For
	conditionalFor()
	fmt.Printf("\n")
	// For while
	forWhile()
	fmt.Printf("\n")
	// For reverse
	reverseFor()

}

func conditionalFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWhile() {
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
}

func reverseFor() {
	for i := 9; i >= 0; i-- {
		fmt.Println(i)
	}
}

func normalFunction(message string) {
	fmt.Println(message)
}

func multipleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func var_const_zval() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)
}

func arithmeticOperator() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Retos:
	// Calcular area rectangulo, trapecio y circulo
	// Area rectangulo
	h := 10
	b := 30
	areaRectangulo := h * b
	fmt.Println("Area rectangulo: ", areaRectangulo)

	// Area trapecio
	AB := 10
	CD := 20
	h = 5
	areaTrapecio := (AB + CD) * h / 2
	fmt.Println("Area trapecio: ", areaTrapecio)

	// Area circulo
	r := 2
	areaCirculo := math.Pi * math.Pow(float64(r), 2)
	fmt.Println("Area trapecio: ", areaCirculo)

}

func packageFmt() {

	// Println
	helloMessage := "hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
