package main

import (
	"fmt"
)

func main() {
	name := "Barros"
	var version float32 = 1.1
	fmt.Println("Hello World! Mr.", name)
	fmt.Println("This software version is ", version)

	fmt.Println("1- Start Monit")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)

	fmt.Println("The selected command was ", command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing logs...")
	} else if command == 0 {
		fmt.Println("Exit...")
	} else {
		fmt.Println("Command not found")
	}

}
