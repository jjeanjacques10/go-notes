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

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exit...")
	default:
		fmt.Println("Command not found")
	}

}
