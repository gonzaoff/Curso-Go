package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {

	//hola mundo
	fmt.Println("Hola,Go")
	/*
		fmt.Println("Hola,Go")
		fmt.Println("Hola,Go")
		fmt.Println("Hola,Go")
		fmt.Println("Hola,Go")
	*/

	fmt.Println("Hola,Go")

	//var/const nameVar tipo = valor
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

	//Concatenando variable entera a float para no perder precision
	var myFloat = 6.5
	fmt.Println(float64(myInt) + myFloat)

	//Variable boolean
	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declarada e inicializada
	myString3 := "Esto es un string"
	fmt.Println(myString3)

	//Constantes
	const myConst = "Esto es una constante"
	fmt.Println(myConst)
}

func condiEIteracion() {

	myInt := 10
	myString := "Hola"
	//inicializacion de condicial IF
	if myInt == 11 && myString == "Hola" {
		fmt.Println("correcto")
	} else if myInt == 10 || myString == "Hola" {
		fmt.Println("la prueba de la condicion OR es correcta")
	} else {
		fmt.Println("incorrecto")
	}

	//var/const Array[cantElementos]Tipo

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	//No agrega valores ademas de los establecidos
	// myArray[3] = 1
	// fmt.Println(myArray[3])

	myArray[0] = 12

	fmt.Println(myArray)
	fmt.Println(myArray[1])
	fmt.Println(myArray[0])

	// Map
	//make(map[clave]valor)

	myMap := make(map[string]int)
	// myMap["Gonzalo"] = 22
	// myMap["Dario"] = 23
	// myMap["Escobar"] = 24
	// mySecondMap := map[string]int{"Marcos": 36, "Roberto": 20}
	// fmt.Println(myMap)
	// fmt.Println(myMap["Gonzalo"])
	// fmt.Println(mySecondMap)

	//Lists

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(123)
	fmt.Println(myList.Back().Value)

	// Bucles
	//Recorre una array item por item arrancando por el valor de "i" en el rango de "myArray"
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	//Recorro el diccionario de modo clave,valor
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//recorro una lista verificando el indice y el valor
	for index, value := range myArray {
		fmt.Println(index, value)
	}
}
