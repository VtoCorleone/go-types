package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sayHello(s string) string {
	return "Hello " + s
}

func workingWithBools() {
	fmt.Println("\nBOOLS")
	var b bool
	fmt.Println(b)

	b = true
	fmt.Println(b)
}

func workingWithNumbers() {
	fmt.Println("\nNUMBERS")
	var num int = 4
	fmt.Println(num)

	var fl float32 = 0.11
	fmt.Println(fl)
}

func workingWithArrays() {
	fmt.Println("\nARRAYS")
	var beatles [4]string
	beatles[0] = "John"
	beatles[1] = "Paul"
	beatles[2] = "Ringo"
	beatles[3] = "George"
	fmt.Println(beatles)
}

func workingWithReflection() {
	fmt.Println("\nREFLECTION")
	var s string = "string"
	var i int = 10
	var f float32 = 1.2

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}

func workingWithTypeConversions() {
	fmt.Println("\nTYPE CONVERSIONS")
	var s string = "true"

	b, err := strconv.ParseBool(s)
	fmt.Println(b)
	fmt.Println(err)

	s = strconv.FormatBool(true)
	fmt.Println(s)
}

func main() {
	fmt.Println(sayHello("Vto"))

	workingWithBools()
	workingWithNumbers()
	workingWithArrays()
	workingWithReflection()
	workingWithTypeConversions()
}
