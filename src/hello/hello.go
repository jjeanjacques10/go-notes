package main

import (
	"fmt"
	"os"
)

func main() {

	showHelloWorld()
	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exit...")
		os.Exit(0)
	default:
		fmt.Println("Command not found")
		os.Exit(-1)
	}
}

func showHelloWorld() {
	name := "Barros"
	var version float32 = 1.1
	fmt.Println("Hello World! Mr.", name)
	fmt.Println("This software version is ", version)
}

func showMenu() {
	fmt.Println("1- Start Monit")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)

	fmt.Println("The selected command was ", commandRead)
	return commandRead
}
