package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showHelloWorld()
	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonit()
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

func startMonit() {
	fmt.Println("Monitoring...")
	var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://jjeanjacques.com"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	url := "https://www.alura.com.br"
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("Website ", url, " has been displayed successfully")
	} else {
		fmt.Println("Website ", url, " has a problem. Status Code:", resp.StatusCode)
	}
}
