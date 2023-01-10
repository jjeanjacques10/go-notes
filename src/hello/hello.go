package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Barros"
	age := 23
	var version float32 = 1.1
	fmt.Println("Hello World! Mr.", name)
	fmt.Println("You are ", age, " years old") // Print age equals 0
	fmt.Println("This software version is ", version)

	fmt.Println("Get variable type: ", reflect.TypeOf(age))
}
