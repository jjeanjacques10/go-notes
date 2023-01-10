package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const MONITORING = 3
const CHECK_DELAY = 5

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
	//sites := []string{"https://www.alura.com.br", "https://jjeanjacques10.github.io", "https://www.caelum.com.br"}

	sites := readFile()

	fmt.Println(sites)

	for i := 0; i < MONITORING; i++ {
		for _, site := range sites {
			checkSite(site)
		}
		time.Sleep(CHECK_DELAY * time.Second)
	}
}

func checkSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Website ", site, " has been displayed successfully")
	} else {
		fmt.Println("Website ", site, " has a problem. Status Code:", resp.StatusCode)
	}
}

func readFile() []string {
	var sites []string

	//file, err := os.Open("sites.txt")
	file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(file))
	return sites
}
