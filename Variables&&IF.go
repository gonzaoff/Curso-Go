package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Forma de crear una variable definida
	var myString string = "Hola, declare que soy un string."
	fmt.Println(myString)

	//Forma de crear una variable indefinida
	var myString2 = 6
	fmt.Println(myString2)

	myString2 = 6 + 6
	fmt.Println(myString2)

	// myString2 = "Hola" error de tipo <-- es tipado estatico, no se puede modificar el tipo de variable

	var myInt int = 7
	myInt = myInt + 4
	fmt.Println(myInt - 1)

	//Concatenando variables de distinto tipo
	fmt.Println(myString, myInt)
	//Mostrando tipo de variable
	fmt.Println(reflect.TypeOf(myString))

	var myFloat = 6.5
	fmt.Println(float64(myInt) + myFloat)

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declarada e inicializada
	myString3 := "Esto es un string"
	fmt.Println(myString3)

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	myInt = 10
	myString = "Hola"

	if myInt == 11 && myString == "Hola" {
		fmt.Println("correcto")
	} else if myInt == 10 || myString == "Hola" {
		fmt.Println("la prueba de la condicion OR es correcta")
	} else {
		fmt.Println("incorrecto")
	}

}
